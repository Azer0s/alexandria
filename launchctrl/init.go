package launchctrl

import (
	"github.com/Azer0s/alexandria/dns"
	log "github.com/sirupsen/logrus"
)

func Startup(cfg Config) {
	log.Info("Starting DNS Server")
	dns.StartDnsUdp(cfg.Hostname, cfg.Port)
}
