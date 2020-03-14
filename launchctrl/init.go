package launchctrl

import (
	"github.com/Azer0s/alexandria/dns"
	"github.com/Azer0s/alexandria/launchctrl/env"
	log "github.com/sirupsen/logrus"
)

func Startup(cfg env.Config) {
	log.Info("Starting DNS Server")
	dns.StartDnsUdp(cfg.Hostname, cfg.Port)
}
