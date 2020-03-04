package dns

import (
	"encoding/json"
	"github.com/Azer0s/alexandria/communication"
	log "github.com/sirupsen/logrus"
	"net"
)

func StartDnsUdp(hostname string, port int) {
	server := communication.UDPServer{
		Hostname: hostname,
		Port:     port,
	}

	server.Start(func(addr net.Addr, buf []byte) []byte {
		pdu, err := ParseDnsPdu(buf)

		if err != nil {
			log.WithFields(log.Fields{
				"client": addr.String(),
				"error":  err.Error(),
			}).Warnf("Error while parsing DNS PDU")

			return []byte{}
		}

		log.WithFields(log.Fields{
			"client":     addr.String(),
			"request_id": pdu.Header.Identifier,
		}).Infof("Handling DNS request")

		log.WithFields(log.Fields{
			"client":     addr.String(),
			"request_id": pdu.Header.Identifier,
		}).Trace(func() string {
			b, err := json.Marshal(pdu)

			if err != nil {
				return ""
			}

			return string(b)
		}())

		return []byte("Hello\n")
	})
}
