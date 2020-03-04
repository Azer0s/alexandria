package response_code

import "github.com/Azer0s/alexandria/dns/enums/fields"

const (
	// NoError No Error
	NoError fields.ResponseCode = 0 // RFC 1035

	// FormErr Format Error
	FormErr fields.ResponseCode = 1 // RFC 1035

	// ServFail Server Failure
	ServFail fields.ResponseCode = 2 // RFC 1035

	// NXDomain Non-Existent Domain
	NXDomain fields.ResponseCode = 3 // RFC 1035

	// NotImp Not Implemented
	NotImp fields.ResponseCode = 4 // RFC 1035

	// Refused Query Refused
	Refused fields.ResponseCode = 5 // RFC 1035

	// YXDomain Name Exists when it should not
	YXDomain fields.ResponseCode = 6 // RFC 2136

	// YXRRSet RR Set Exists when it should not
	YXRRSet fields.ResponseCode = 7 // RFC 2136

	// NXRRSet RR Set that should exist does not
	NXRRSet fields.ResponseCode = 8 // RFC 2136

	// NotAuth Not Authorized
	NotAuth fields.ResponseCode = 9 // RFC 2136

	// NotZone Name not contained in zone
	NotZone fields.ResponseCode = 10 // RFC 2136
)
