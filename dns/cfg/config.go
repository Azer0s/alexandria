package cfg

import (
	"github.com/Azer0s/alexandria/dns/cfg/interpreter"
	"github.com/Azer0s/alexandria/dns/enums/fields"
	"github.com/Azer0s/alexandria/dns/enums/record_class"
	"github.com/Azer0s/alexandria/dns/protocol"
	"github.com/Azer0s/alexandria/launchctrl/env"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
)

func Parse(file string) interpreter.Zones {
	b, _ := ioutil.ReadFile(file)
	str := string(b)

	dir, err := filepath.Abs(filepath.Dir(file))
	if err != nil {
		log.Fatal(err)
	}

	str, ttl := interpreter.DoPreprocessing(str, dir)
	zones := interpreter.ParseConfig(interpreter.LexConfig(str), ttl)
	return zones
}

var zoneMap = make(interpreter.ZoneMap, 0)

func getFQDNName(entry interpreter.Entry, zoneFqdn string) string {
	fqdnName := zoneFqdn

	if !entry.Native {
		fqdnName = entry.Name + "." + fqdnName
	}

	return fqdnName
}

func zoneToDnsPdu(zone interpreter.Zone, parentZone string) interpreter.ZoneMap {
	result := make(interpreter.ZoneMap, 0)

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

func GetAnswer(fqdn string, rt fields.RecordType) []protocol.DNSResourceRecord {
	return zoneMap[fqdn][rt]
}

func ConfigureZones(cfg env.Config) {
	zones := make([]interpreter.Zones, 0)
	for _, config := range cfg.Configs {
		zones = append(zones, Parse(config))
	}

	for _, zone := range zones {
		for _, subzone := range zone {
			for k, v := range zoneToDnsPdu(subzone, "") {
				zoneMap[k] = v
			}
		}
	}

	//TODO: Store ZoneMap _somewhere efficient_
}
