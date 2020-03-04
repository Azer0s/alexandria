package protocol

import "github.com/Azer0s/alexandria/dns/enums/fields"

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
