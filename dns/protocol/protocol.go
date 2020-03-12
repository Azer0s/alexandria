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

// https://www.cloudflare.com/learning/dns/dns-records

// DNSHeader describes the DNS header as documented in RFC 1035
type DNSHeader struct {
	Identifier                     uint16 `json:"identifier"`
	Flags                          uint16 `json:"flags"`
	TotalQuestions                 uint16 `json:"num_questions"`
	TotalAnswerResourceRecords     uint16 `json:"num_answers"`
	TotalAuthorityResourceRecords  uint16 `json:"num_authority"`
	TotalAdditionalResourceRecords uint16 `json:"num_additional"`
}

// DNSFlags describe the flags of the DNS header
type DNSFlags struct {
	QueryResponse       fields.MessageType  `json:"query_response"`
	OpCode              fields.OpCode       `json:"op_code"`
	AuthoritativeAnswer bool                `json:"authoritative_answer"`
	Truncated           bool                `json:"truncated"`
	RecursionDesired    bool                `json:"recursion_desired"`
	RecursionAvailable  bool                `json:"recursion_available"`
	ResponseCode        fields.ResponseCode `json:"response_code"`
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
	Flags                     DNSFlags            `json:"flags"`
	Questions                 []DNSQuestion       `json:"questions"`
	AnswerResourceRecords     []DNSResourceRecord `json:"answers"`
	AuthorityResourceRecords  []DNSResourceRecord `json:"authority"`
	AdditionalResourceRecords []DNSResourceRecord `json:"additional"`
}

func (flags DNSFlags) Uint16() uint16 {
	result := uint16(0)

	if flags.QueryResponse {
		result |= uint16(0b1000_0000_0000_0000)
	}

	result |= (uint16(flags.OpCode) << 11) & uint16(0b01111_1000_0000_0000)

	if flags.AuthoritativeAnswer {
		result |= uint16(0b0000_0100_0000_0000)
	}

	if flags.Truncated {
		result |= uint16(0b0000_0010_0000_0000)
	}

	if flags.RecursionDesired {
		result |= uint16(0b0000_0001_0000_0000)
	}

	if flags.RecursionAvailable {
		result |= uint16(0b0000_0000_1000_0000)
	}

	result |= uint16(uint8(flags.ResponseCode) & uint8(0b0000_1111))

	return result
}

func writeLabels(responseBuffer *bytes.Buffer, labels []string) error {
	//If the label is nil, we just insert a DNS pointer to the request FQDN position (byte 13)
	if labels == nil {
		_, err := responseBuffer.Write([]byte{0xc0, 0x0c})
		return err
	}

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

	pdu.Header.Flags = pdu.Flags.Uint16()

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
