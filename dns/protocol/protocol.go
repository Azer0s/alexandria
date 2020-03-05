package protocol

import (
	"bytes"
	"encoding/binary"
	"github.com/Azer0s/alexandria/dns/enums/fields"
)

// RE: Labels
// RFC1035: "Domain names in messages are expressed in terms of a sequence
// of labels. Each label is represented as a one octet length field followed
// by that number of octets.  Since every domain name ends with the null label
// of the root, a domain name is terminated by a length byte of zero."

// DNSHeader describes the DNS header as documented in RFC 1035
type DNSHeader struct {
	Identifier                     uint16 `json:"identifier"`
	Flags                          uint16 `json:"flags"`
	TotalQuestions                 uint16 `json:"num_questions"`
	TotalAnswerResourceRecords     uint16 `json:"num_answers"`
	TotalAuthorityResourceRecords  uint16 `json:"num_authority"`
	TotalAdditionalResourceRecords uint16 `json:"num_additional"`
}

// DNSResourceRecord describes individual records in the request and response of the DNS payload body
type DNSResourceRecord struct {
	Labels             []string           `json:"labels"`
	Type               fields.RecordType  `json:"type"`
	Class              fields.RecordClass `json:"class"`
	TimeToLive         uint32             `json:"ttl"`
	ResourceDataLength uint16             `json:"rd_length"`
	ResourceData       []byte             `json:"rd"`
}

// DNSQuestion describes individual records in the request and response of the DNS payload body
type DNSQuestion struct {
	Labels []string           `json:"labels"`
	Type   fields.RecordType  `json:"type"`
	Class  fields.RecordClass `json:"class"`
}

// DNSPDU describes the DNS protocol data unit as documented in RFC 1035
type DNSPDU struct {
	Header                    DNSHeader           `json:"header"`
	Questions                 []DNSQuestion       `json:"questions"`
	AnswerResourceRecords     []DNSResourceRecord `json:"answers"`
	AuthorityResourceRecords  []DNSResourceRecord `json:"authority"`
	AdditionalResourceRecords []DNSResourceRecord `json:"additional"`
}

func writeLabels(responseBuffer *bytes.Buffer, labels []string) error {
	for _, label := range labels {
		labelLength := len(label)
		labelBytes := []byte(label)

		responseBuffer.WriteByte(byte(labelLength))
		responseBuffer.Write(labelBytes)
	}

	err := responseBuffer.WriteByte(byte(0))

	return err
}

func writeResourceRecords(buffer *bytes.Buffer, rrs []DNSResourceRecord) error {
	for _, rr := range rrs {
		err := writeLabels(buffer, rr.Labels)
		if err != nil {
			return err
		}

		err = binary.Write(buffer, binary.BigEndian, rr.Type)
		if err != nil {
			return err
		}

		err = binary.Write(buffer, binary.BigEndian, rr.Class)
		if err != nil {
			return err
		}

		err = binary.Write(buffer, binary.BigEndian, rr.TimeToLive)
		if err != nil {
			return err
		}

		err = binary.Write(buffer, binary.BigEndian, rr.ResourceDataLength)
		if err != nil {
			return err
		}

		err = binary.Write(buffer, binary.BigEndian, rr.ResourceData)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pdu DNSPDU) Bytes() ([]byte, error) {
	var responseBuffer = new(bytes.Buffer)

	err := binary.Write(responseBuffer, binary.BigEndian, &pdu.Header)

	if err != nil {
		return nil, err
	}

	for _, question := range pdu.Questions {
		err := writeLabels(responseBuffer, question.Labels)
		if err != nil {
			return nil, err
		}

		err = binary.Write(responseBuffer, binary.BigEndian, question.Type)
		if err != nil {
			return nil, err
		}

		err = binary.Write(responseBuffer, binary.BigEndian, question.Class)
		if err != nil {
			return nil, err
		}
	}

	err = writeResourceRecords(responseBuffer, pdu.AnswerResourceRecords)
	if err != nil {
		return nil, err
	}

	err = writeResourceRecords(responseBuffer, pdu.AuthorityResourceRecords)
	if err != nil {
		return nil, err
	}

	err = writeResourceRecords(responseBuffer, pdu.AdditionalResourceRecords)
	if err != nil {
		return nil, err
	}

	return responseBuffer.Bytes(), nil
}
