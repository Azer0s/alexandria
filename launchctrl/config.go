package launchctrl

import (
	"encoding/json"
	dnscfg "github.com/Azer0s/alexandria/dns/cfg"
	"github.com/Azer0s/alexandria/dns/cfg/interpreter"
	"github.com/Azer0s/alexandria/dns/enums/fields"
	"github.com/Azer0s/alexandria/dns/enums/record_class"
	"github.com/Azer0s/alexandria/dns/protocol"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

type ZoneMap map[string]map[fields.RecordType][]protocol.DNSResourceRecord

const (
	LOGLEVEL   = "LOG_LEVEL"
	LOGJSON    = "LOG_JSON"
	PRINTTITLE = "PRINT_TITLE"
	HOSTNAME   = "HOSTNAME"
	PORT       = "PORT"
	CONFIGS    = "CONFIGS"

	ENV_INVALID_STRING = "Environment variable %s not set or invalid, aborting!"
)

type Config struct {
	LogLevel   log.Level `json:"log_level"`
	LogJson    bool      `json:"log_json"`
	PrintTitle bool      `json:"print_title"`
	Hostname   string    `json:"hostname"`
	Port       int       `json:"port"`
	Configs    []string  `json:"configs"`
}

func GetConfig() Config {
	cfg := Config{}

	logLevel, err := log.ParseLevel(os.Getenv(LOGLEVEL))
	if err != nil {
		log.Fatalf(ENV_INVALID_STRING, LOGLEVEL)
	}
	cfg.LogLevel = logLevel

	logJson, err := strconv.ParseBool(os.Getenv(LOGJSON))
	if err != nil {
		log.Fatalf(ENV_INVALID_STRING, LOGJSON)
	}
	cfg.LogJson = logJson

	printTitle, err := strconv.ParseBool(os.Getenv(PRINTTITLE))
	if err != nil {
		log.Fatalf(ENV_INVALID_STRING, PRINTTITLE)
	}
	cfg.PrintTitle = printTitle

	hostname := os.Getenv(HOSTNAME)
	if hostname == "" {
		hostname = "localhost"
	}
	cfg.Hostname = hostname

	port, err := strconv.Atoi(os.Getenv(PORT))
	if err != nil {
		log.Fatalf(ENV_INVALID_STRING, PORT)
	}
	cfg.Port = port

	cfg.Configs = strings.Split(os.Getenv(CONFIGS), ",")

	return cfg
}

func ConfigureLog(cfg Config) {
	log.SetLevel(cfg.LogLevel)

	if cfg.LogJson {
		log.SetFormatter(&log.JSONFormatter{})
	}

	log.Trace(func() string {
		b, err := json.Marshal(cfg)

		if err != nil {
			return ""
		}

		return string(b)
	}())
}

func getFQDNName(entry interpreter.Entry, zoneFqdn string) string {
	fqdnName := zoneFqdn

	if !entry.Native {
		fqdnName = entry.Name + "." + fqdnName
	}

	return fqdnName
}

func zoneToDnsPdu(zone interpreter.Zone, parentZone string) ZoneMap {
	result := make(ZoneMap, 0)

	if parentZone != "" {
		zone.FQDN = zone.FQDN + "." + parentZone
	}

	for _, entry := range zone.Entries {
		fqdnName := getFQDNName(entry, zone.FQDN)

		if result[fqdnName] == nil {
			result[fqdnName] = make(map[fields.RecordType][]protocol.DNSResourceRecord, 0)
		}

		if result[fqdnName][entry.Type] == nil {
			result[fqdnName][entry.Type] = make([]protocol.DNSResourceRecord, 0)
		}
	}

	for _, entry := range zone.Entries {
		fqdnName := getFQDNName(entry, zone.FQDN)

		result[fqdnName][entry.Type] = append(result[fqdnName][entry.Type], protocol.DNSResourceRecord{
			Labels:             nil,
			Type:               entry.Type,
			Class:              record_class.Internet,
			TimeToLive:         entry.TimeToLive,
			ResourceDataLength: uint16(len(entry.Value)),
			ResourceData:       entry.Value,
		})
	}

	for _, z := range zone.Zones {
		for k, v := range zoneToDnsPdu(z, zone.FQDN) {
			result[k] = v
		}
	}

	return result
}

func ConfigureZones(cfg Config) {
	zones := make([]dnscfg.Zones, 0)
	for _, config := range cfg.Configs {
		zones = append(zones, dnscfg.Parse(config))
	}

	zoneMap := make(ZoneMap, 0)

	for _, zone := range zones {
		for _, subzone := range zone {
			for k, v := range zoneToDnsPdu(subzone, "") {
				zoneMap[k] = v
			}
		}
	}

	//TODO: Store ZoneMap _somewhere efficient_
}
