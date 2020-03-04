package record_class

import "github.com/Azer0s/alexandria/dns/enums/fields"

const (
	Internet fields.RecordClass = 1 // RFC 1035
	Chaos    fields.RecordClass = 3
	Hesiod   fields.RecordClass = 4
	None     fields.RecordClass = 254 // RFC 2136
	Any      fields.RecordClass = 255 // RFC 1035
)
