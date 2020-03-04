package launchctrl

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

const (
	LOGLEVEL   = "LOG_LEVEL"
	LOGJSON    = "LOG_JSON"
	PRINTTITLE = "PRINT_TITLE"
	HOSTNAME   = "HOSTNAME"
	PORT       = "PORT"

	ENV_INVALID_STRING = "Environment variable %s not set or invalid, aborting!"
)

type Config struct {
	LogLevel   log.Level `json:"log_level"`
	LogJson    bool      `json:"log_json"`
	PrintTitle bool      `json:"print_title"`
	Hostname   string    `json:"hostname"`
	Port       int       `json:"port"`
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
