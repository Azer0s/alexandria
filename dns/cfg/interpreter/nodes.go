package interpreter

import "github.com/Azer0s/alexandria/dns/enums/fields"

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
