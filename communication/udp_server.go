package communication

import (
	log "github.com/sirupsen/logrus"
	"net"
)

type UDPServer struct {
	Hostname string
	Port     int
}

type UDPHandler func(addr net.Addr, buf []byte)[]byte

func (s UDPServer) Start(handler UDPHandler) {
	log.Debugf("Listening on %s:%d/udp", s.Hostname, s.Port)
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
				log.Tracef("Received %d bytes from client %s", n, addr.String())
				result := handler(addr, buf[:n])
				n, err := pc.WriteTo(result, addr)

				if err != nil {
					log.Debugf("Error while replying to client %s", addr.String())
				}
				log.Tracef("Sent %d bytes to client %s", n, addr.String())
			}()
		}
	}()
}