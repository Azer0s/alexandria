package record_type

import "github.com/Azer0s/alexandria/dns/enums/fields"

const (
	// A a host address
	A fields.RecordType = 1 // RFC 1035, Section 3.4.1

	// NS an authoritative name server
	NS fields.RecordType = 2 // RFC 1035, Section 3.3.11

	// MD a mail destination (OBSOLETE - use MX)
	MD fields.RecordType = 3 // RFC 1035, Section 3.3.4 (obsolete)

	// MF a mail forwarder (OBSOLETE - use MX)
	MF fields.RecordType = 4 // RFC 1035, Section 3.3.5 (obsolete)

	// CNAME the canonical name for an alias
	CNAME fields.RecordType = 5 // RFC 1035, Section 3.3.1

	// SOA marks the start of a zone of authority
	SOA fields.RecordType = 6 // RFC 1035, Section 3.3.13

	// MB a mailbox domain name (EXPERIMENTAL)
	MB fields.RecordType = 7 // RFC 1035, Section 3.3.3

	// MG a mail group member (EXPERIMENTAL)
	MG fields.RecordType = 8 // RFC 1035, Section 3.3.6

	// MR a mail rename domain name (EXPERIMENTAL)
	MR fields.RecordType = 9 // RFC 1035, Section 3.3.8

	// NULL a null RR (EXPERIMENTAL)
	NULL fields.RecordType = 10 // RFC 1035, Section 3.3.10

	// WKS a well known service description
	WKS fields.RecordType = 11 // RFC 1035, Section 3.4.2 (deprecated)

	// PTR a domain name pointer
	PTR fields.RecordType = 12 // RFC 1035, Section 3.3.12

	// HINFO host information
	HINFO fields.RecordType = 13 // RFC 1035, Section 3.3.2

	// MINFO mailbox or mail list information
	MINFO fields.RecordType = 14 // RFC 1035, Section 3.3.7

	// MX mail exchange
	MX fields.RecordType = 15 // RFC 1035, Section 3.3.9

	// TXT text strings
	TXT fields.RecordType = 16 // RFC 1035, Section 3.3.14

	// RP for Responsible Person
	RP fields.RecordType = 17 // RFC 1183, Section 2.2

	// AFSDB for AFS Data Base location
	AFSDB fields.RecordType = 18 // RFC 1183, Section 1

	// X25 for X.25 PSDN address
	X25 fields.RecordType = 19 // RFC 1183, Section 3.1

	// ISDN for ISDN address
	ISDN fields.RecordType = 20 // RFC 1183, Section 3.2

	// RT for Route Through
	RT fields.RecordType = 21 // RFC 1183, Section 3.3

	// NSAP for NSAP address, NSAP style A record
	NSAP fields.RecordType = 22 // RFC 1706, Section 5

	// NSAP_PTR	for domain name pointer, NSAP style
	NSAP_PTR fields.RecordType = 23 // RFC 1348 (obsolete)

	// SIG for security signature
	SIG fields.RecordType = 24 // RFC 2535, Section 4.1

	// KEY for security key
	KEY fields.RecordType = 25 // RFC 2535, Section 3.1

	// PX X.400 mail mapping information
	PX fields.RecordType = 26 // RFC 2163,

	// GPOS Geographical Position
	GPOS fields.RecordType = 27 // RFC 1712 (obsolete)

	// AAAA IP6 Address
	AAAA fields.RecordType = 28 // RFC 1886, Section 2.1

	// LOC Location Information
	LOC fields.RecordType = 29 // RFC 1876

	// NXT Next Domain (OBSOLETE)
	NXT fields.RecordType = 30 // RFC 2535, Section 5.2 obsoleted by RFC3755

	// EID Endpoint Identifier
	EID fields.RecordType = 31 // draft-ietf-nimrod-dns-xx.txt

	// NIMLOC Nimrod Locator
	NIMLOC fields.RecordType = 32 // draft-ietf-nimrod-dns-xx.txt

	// SRV Server Selection
	SRV fields.RecordType = 33 // RFC 2052

	// ATMA ATM Address
	ATMA fields.RecordType = 34 // ???

	// NAPTR Naming Authority Pointer
	NAPTR fields.RecordType = 35 // RFC 2168

	// KX Key Exchanger
	KX fields.RecordType = 36 // RFC 2230

	// CERT CERT
	CERT fields.RecordType = 37 // RFC 2538

	// DNAME DNAME
	DNAME fields.RecordType = 39 // RFC 2672

	// OPT OPT
	OPT fields.RecordType = 41 // RFC 2671

	// APL APL
	APL fields.RecordType = 42 // RFC 3123

	// DS Delegation Signer
	DS fields.RecordType = 43 // RFC 4034

	// SSHFP SSH Key Fingerprint
	SSHFP fields.RecordType = 44 // RFC 4255

	// IPSECKEY IPSECKEY
	IPSECKEY fields.RecordType = 45 // RFC 4025

	// RRSIG RRSIG
	RRSIG fields.RecordType = 46 // RFC 4034

	// NSEC NSEC
	NSEC fields.RecordType = 47 // RFC 4034

	// DNSKEY DNSKEY
	DNSKEY fields.RecordType = 48 // RFC 4034

	// DHCID DHCID
	DHCID fields.RecordType = 49 // RFC 4701

	// NSEC3 NSEC3
	NSEC3 fields.RecordType = 50 // RFC 5155

	// NSEC3PARAM NSEC3PARAM
	NSEC3PARAM fields.RecordType = 51 // RFC 5155

	// TLSA TLSA
	TLSA fields.RecordType = 52 // RFC 6698

	// HIP Host Identity Protocol
	HIP fields.RecordType = 55 // RFC 5205

	// CDS Child DS
	CDS fields.RecordType = 59 // RFC 7344

	// CDNSKEY DNSKEY(s) the Child wants reflected in DS
	CDNSKEY fields.RecordType = 60 // RFC 7344

	// SPF SPF
	SPF fields.RecordType = 99 // RFC 7208

	// UINFO UINFO
	UINFO = 100 // IANA-Reserved

	// UID UID
	UID = 101 // IANA-Reserved

	// GID GID
	GID = 102 // IANA-Reserved

	// UNSPEC UNSPEC
	UNSPEC = 103 // IANA-Reserved

	// TKEY Transaction Key
	TKEY fields.RecordType = 249 // RFC 2930

	// TSIG Transaction Signature
	TSIG fields.RecordType = 250 // RFC 2931

	// IXFR	incremental transfer
	IXFR fields.RecordType = 251 // RFC 1995

	// AXFR transfer of an entire zone
	AXFR fields.RecordType = 252 // RFC 1035

	// MAILB mailbox-related RRs (MB, MG or MR)
	MAILB fields.RecordType = 253 // RFC 1035 (MB, MG, MR)

	// MAILA mail agent RRs (OBSOLETE - see MX)
	MAILA fields.RecordType = 254 // RFC 1035 (obsolete - see MX)

	// ANY A request for some or all records the server has available
	ANY fields.RecordType = 255 // RFC 1035

	// URI URI
	URI fields.RecordType = 256 // RFC 7553

	// CAA Certification Authority Restriction
	CAA fields.RecordType = 257 // RFC 6844

	// DLV DNSSEC Lookaside Validation (OBSOLETE)
	DLV fields.RecordType = 32769 // RFC 4431 (informational)
)
