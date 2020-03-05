package dns

import (
	"encoding/json"
	"github.com/Azer0s/alexandria/communication"
	"github.com/Azer0s/alexandria/dns/enums/record_class"
	"github.com/Azer0s/alexandria/dns/enums/record_type"
	"github.com/Azer0s/alexandria/dns/protocol"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
)

func getByQuestion(q protocol.DNSQuestion) []protocol.DNSResourceRecord {
	fqdn := strings.Join(q.Labels, ".")

	var answers = make([]protocol.DNSResourceRecord, 0)

	if fqdn == "google.com" && q.Type == record_type.A && q.Class == record_class.Internet {
		googleRecord := protocol.DNSResourceRecord{
			Labels:             []string{"google", "com"},
			Type:               record_type.A,
			Class:              record_class.Internet,
			TimeToLive:         31337,
			ResourceData:       []byte{216, 58, 207, 46}, // ipv4 address
			ResourceDataLength: 4,
		}

		answers = append(answers, googleRecord)
	}

	if fqdn == "google.com" && q.Type == record_type.NS && q.Class == record_class.Internet {

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
		b := []byte("ns1")

		ns1 := []byte{byte(len(b))}
		ns1 = append(ns1, b...)
		ns1 = append(ns1, []byte{0xc0, 0x0c}...)

		//TODO: Fix flags
		answers = append(answers, protocol.DNSResourceRecord{
			Labels:             nil,
			Type:               record_type.NS,
			Class:              record_class.Internet,
			TimeToLive:         36493,
			ResourceData:       ns1,
			ResourceDataLength: uint16(len(ns1)),
		})
	}

	return answers
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
				Flags:                          1 << 15, //Response flag
				TotalQuestions:                 pdu.Header.TotalQuestions,
				TotalAnswerResourceRecords:     uint16(len(answers)),
				TotalAuthorityResourceRecords:  0,
				TotalAdditionalResourceRecords: 0,
			},
			Questions:                 pdu.Questions,
			AnswerResourceRecords:     answers,
			AuthorityResourceRecords:  make([]protocol.DNSResourceRecord, 0),
			AdditionalResourceRecords: make([]protocol.DNSResourceRecord, 0),
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
