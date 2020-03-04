package message_type

import (
	"github.com/Azer0s/alexandria/dns/enums/fields"
)

const (
	Query    fields.MessageType = 0 // RFC 1035
	Response fields.MessageType = 1 // RFC 1035
)
