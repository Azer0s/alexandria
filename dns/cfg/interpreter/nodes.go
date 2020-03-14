package interpreter

import (
	"github.com/Azer0s/alexandria/dns/enums/fields"
	"github.com/Azer0s/alexandria/dns/protocol"
)

type Zone struct {
	FQDN    string
	Entries []Entry
	Zones   []Zone
}

type Entry struct {
	Type       fields.RecordType
	TimeToLive uint32
	Recursive  bool
	Native     bool
	Name       string
	Value      []byte
}

type Zones []Zone
type ZoneMap map[string]map[fields.RecordType][]protocol.DNSResourceRecord
