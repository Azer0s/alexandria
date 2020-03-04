package communication

import (
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
)

type UDPServer struct {
	Hostname string
	Port     int
}

type UDPHandler func(addr net.Addr, buf []byte) []byte

func (s UDPServer) Start(handler UDPHandler) {
	log.Infof("Listening on %s:%d/udp", s.Hostname, s.Port)
	go func() {
		pc, err := net.ListenPacket("udp", ":8333")
		if err != nil {
			log.Fatal(err)
		}
		defer pc.Close()

		for {
			buf := make([]byte, 1024)
			n, addr, err := pc.ReadFrom(buf)
			if err != nil {
				continue
			}
			go func() {
				log.WithFields(log.Fields{
					"client": addr.String(),
				}).Debugf("Received %d bytes", n)

				buf = buf[:n]

				log.WithFields(log.Fields{
					"client": addr.String(),
				}).Trace(func() string {
					return strings.ReplaceAll(fmt.Sprintf("\n%s", hex.Dump(buf)), "\n", "\n\t")
				}())

				result := handler(addr, buf)
				n, err := pc.WriteTo(result, addr)

				if err != nil {
					log.WithFields(log.Fields{
						"client": addr.String(),
					}).Warnf("Error while replying")
				}

				log.WithFields(log.Fields{
					"client": addr.String(),
				}).Debugf("Sent %d bytes", n)
			}()
		}
	}()
}
