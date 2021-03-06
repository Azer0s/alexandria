package dns

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/Azer0s/alexandria/dns/enums/fields"
	"github.com/Azer0s/alexandria/dns/protocol"
)

func parseHeader(buffer *bytes.Buffer) (protocol.DNSHeader, error) {
	var header protocol.DNSHeader
	err := binary.Read(buffer, binary.BigEndian, &header)

	if err != nil {
		return protocol.DNSHeader{}, errors.New("couldn't read DNS header")
	}

	return header, nil
}

func readLabels(buffer *bytes.Buffer) ([]string, error) {
	labels := make([]string, 0)

	b, err := buffer.ReadByte()

	for ; b != 0 && err == nil; b, err = buffer.ReadByte() {
		length := int(b)
		//TODO: Parse DNS pointer
		labelBytes := buffer.Next(length)
		labels = append(labels, string(labelBytes))
	}

	return labels, err
}

func parseResourceRecords(num int, buffer *bytes.Buffer) ([]protocol.DNSResourceRecord, error) {
	rrs := make([]protocol.DNSResourceRecord, 0)
	for i := 0; i < num; i++ {
		rr := protocol.DNSResourceRecord{}

		labels, err := readLabels(buffer)
		rr.Labels = labels

		if err != nil {
			return make([]protocol.DNSResourceRecord, 0), err
		}

		rr.Type = fields.RecordType(binary.BigEndian.Uint16(buffer.Next(2)))
		rr.Class = fields.RecordClass(binary.BigEndian.Uint16(buffer.Next(2)))
		rr.TimeToLive = binary.BigEndian.Uint32(buffer.Next(4))
		rr.ResourceDataLength = binary.BigEndian.Uint16(buffer.Next(2))
		rr.ResourceData = buffer.Next(int(rr.ResourceDataLength))

		rrs = append(rrs, rr)
	}

	return rrs, nil
}

func parseBody(header protocol.DNSHeader, buffer *bytes.Buffer) (protocol.DNSPDU, error) {
	pdu := protocol.DNSPDU{}
	pdu.Header = header

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

	answers, err := parseResourceRecords(int(header.TotalAnswerResourceRecords), buffer)
	if err != nil {
		return protocol.DNSPDU{}, err
	}

	authority, err := parseResourceRecords(int(header.TotalAuthorityResourceRecords), buffer)
	if err != nil {
		return protocol.DNSPDU{}, err
	}

	additional, err := parseResourceRecords(int(header.TotalAdditionalResourceRecords), buffer)
	if err != nil {
		return protocol.DNSPDU{}, err
	}

	pdu.Questions = questions
	pdu.AnswerResourceRecords = answers
	pdu.AuthorityResourceRecords = authority
	pdu.AdditionalResourceRecords = additional

	return pdu, nil
}

func parseFlags(flags uint16) protocol.DNSFlags {
	result := protocol.DNSFlags{}

	if (flags >> 15) == 1 {
		result.QueryResponse = true
	}

	result.OpCode = fields.OpCode(uint8(flags>>11) & uint8(0b0000_1111))

	if (flags<<5)>>15 == 1 {
		result.AuthoritativeAnswer = true
	}

	if (flags<<6)>>15 == 1 {
		result.Truncated = true
	}

	if (flags<<7)>>15 == 1 {
		result.RecursionDesired = true
	}

	if (flags<<8)>>15 == 1 {
		result.RecursionAvailable = true
	}

	result.ResponseCode = fields.ResponseCode(uint8(flags) & uint8(0b0000_1111))

	return result
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

	pdu.Flags = parseFlags(header.Flags)

	return pdu, nil
}
