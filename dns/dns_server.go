package dns

import (
	"encoding/json"
	"github.com/Azer0s/alexandria/communication"
	"github.com/Azer0s/alexandria/dns/cfg"
	"github.com/Azer0s/alexandria/dns/enums/message_type"
	"github.com/Azer0s/alexandria/dns/enums/opcode"
	"github.com/Azer0s/alexandria/dns/enums/response_code"
	"github.com/Azer0s/alexandria/dns/protocol"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
)

func getByQuestion(q protocol.DNSQuestion) []protocol.DNSResourceRecord {
	return cfg.GetAnswer(strings.Join(q.Labels, "."), q.Type)

	//TODO: Figure out DNS pointers
	//Okay...so apparently to get DNS pointers to work,
	//one has to insert the size of the FQDN before the
	//pointer, then the label and then the pointer
	//So basically the same as for normal labels
	//Which means that if the end of the request FQDN
	//matches with our response FQDN, we can cut the end
	//and replace it with a pointer

	//And I could've known all that without searching
	//for hours by just reading the RFC 🤷🏻‍️

	//TODO: Rework labels & resource data (sometimes it needs to be formatted as a label)
	//TODO: Figure out SOA for NS requests
}

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

		answers := make([]protocol.DNSResourceRecord, 0)
		for _, question := range pdu.Questions {
			answers = append(answers, getByQuestion(question)...)
		}

		response := protocol.DNSPDU{
			Header: protocol.DNSHeader{
				Identifier:                     pdu.Header.Identifier,
				Flags:                          0,
				TotalQuestions:                 pdu.Header.TotalQuestions,
				TotalAnswerResourceRecords:     uint16(len(answers)),
				TotalAuthorityResourceRecords:  0,
				TotalAdditionalResourceRecords: 0,
			},
			Questions:                 pdu.Questions,
			AnswerResourceRecords:     answers,
			AuthorityResourceRecords:  make([]protocol.DNSResourceRecord, 0),
			AdditionalResourceRecords: make([]protocol.DNSResourceRecord, 0),
			Flags: protocol.DNSFlags{
				QueryResponse:       message_type.Response,
				OpCode:              opcode.Query,
				AuthoritativeAnswer: false,
				Truncated:           false,
				RecursionDesired:    false,
				RecursionAvailable:  false,
				ResponseCode:        response_code.NoError,
			},
		}

		b, err := response.Bytes()
		if err != nil {
			log.WithFields(log.Fields{
				"client": addr.String(),
				"error":  err.Error(),
			}).Warnf("Error while converting DNS response to bytes")

			return []byte{}
		}

		return b
	})
}
