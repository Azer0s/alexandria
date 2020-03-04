package opcode

import "github.com/Azer0s/alexandria/dns/enums/fields"

const (
	Query  fields.OpCode = 0 // RFC 1035
	Status fields.OpCode = 2 // RFC 1035
	Notify fields.OpCode = 4 // RFC 1996
	Update fields.OpCode = 5 // RFC 2136
)
