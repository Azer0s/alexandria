package dns

import (
	"github.com/Azer0s/alexandria/communication"
	"net"
)

func StartDnsUdp(hostname string, port int){
	server := communication.UDPServer{
		Hostname: hostname,
		Port: port,
	}

	server.Start(func(addr net.Addr, buf []byte) []byte {
		return []byte("Hello\n")
	})
}