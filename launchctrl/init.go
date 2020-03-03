package launchctrl

import (
	"github.com/Azer0s/alexandria/dns"
	log "github.com/sirupsen/logrus"
)

func Startup() {
	log.Info("Starting DNS Server")
	dns.StartDnsUdp("localhost", 1053)
}
