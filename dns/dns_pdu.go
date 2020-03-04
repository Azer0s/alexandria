package dns

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/Azer0s/alexandria/dns/protocol"

	"github.com/Azer0s/alexandria/dns/enums/fields"
)

func parseHeader(buffer *bytes.Buffer) (protocol.DNSHeader, error) {
	var header protocol.DNSHeader
	err := binary.Read(buffer, binary.BigEndian, &header)

	if err != nil {
		return protocol.DNSHeader{}, errors.New("couldn't read DNS header")
	}

	return header, nil
}

// RFC1035: "Domain names in messages are expressed in terms of a sequence
// of labels. Each label is represented as a one octet length field followed
// by that number of octets.  Since every domain name ends with the null label
// of the root, a domain name is terminated by a length byte of zero."
func readLabels(buffer *bytes.Buffer) ([]string, error) {
	labels := make([]string, 0)

	b, err := buffer.ReadByte()

	for ; b != 0 && err == nil; b, err = buffer.ReadByte() {
		length := int(b)
		labelBytes := buffer.Next(length)
		labels = append(labels, string(labelBytes))
	}

	return labels, err
}

func parseBody(header protocol.DNSHeader, buffer *bytes.Buffer) (protocol.DNSPDU, error) {
	pdu := protocol.DNSPDU{}

	questions := make([]protocol.DNSQuestion, 0)
	for i := 0; i < int(header.TotalQuestions); i++ {
		question := protocol.DNSQuestion{}

		labels, err := readLabels(buffer)
		question.Labels = labels

		if err != nil {
			return protocol.DNSPDU{}, err
		}

		question.Type = fields.RecordType(binary.BigEndian.Uint16(buffer.Next(2)))
		question.Class = fields.RecordClass(binary.BigEndian.Uint16(buffer.Next(2)))

		questions = append(questions, question)
	}

	//TODO: Answers, AuthorityResourceRecords

	additional := make([]protocol.DNSResourceRecord, 0)
	for i := 0; i < int(header.TotalAdditionalResourceRecords); i++ {
		rr := protocol.DNSResourceRecord{}

		labels, err := readLabels(buffer)
		rr.Labels = labels

		if err != nil {
			return protocol.DNSPDU{}, err
		}

		rr.Type = fields.RecordType(binary.BigEndian.Uint16(buffer.Next(2)))
		rr.Class = fields.RecordClass(binary.BigEndian.Uint16(buffer.Next(2)))
		rr.TimeToLive = binary.BigEndian.Uint32(buffer.Next(4))
		rr.ResourceDataLength = binary.BigEndian.Uint16(buffer.Next(2))
		rr.ResourceData = buffer.Next(int(rr.ResourceDataLength))

		additional = append(additional, rr)
	}

	pdu.Questions = questions
	//TODO: Answers, AuthorityResourceRecords
	pdu.AdditionalResourceRecords = additional

	pdu.Header = header

	return pdu, nil
}

func ParseDnsPdu(buf []byte) (protocol.DNSPDU, error) {
	buffer := bytes.NewBuffer(buf)

	header, err := parseHeader(buffer)
	if err != nil {
		return protocol.DNSPDU{}, err
	}

	pdu, err := parseBody(header, buffer)
	if err != nil {
		return protocol.DNSPDU{}, err
	}

	return pdu, nil
}
