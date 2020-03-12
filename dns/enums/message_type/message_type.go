package message_type

import (
	"github.com/Azer0s/alexandria/dns/enums/fields"
)

const (
	Query    fields.MessageType = false // RFC 1035
	Response fields.MessageType = true  // RFC 1035
)
