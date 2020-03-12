package fields

// https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml

// MessageType The DNS QueryResponse flag
type MessageType bool

// OpCode The DNS Opcode
type OpCode uint8 //4 bit

//RecordClass This is the DNS Class
type RecordClass uint16

//RecordType This is the DNS ResourceRecord (RR) Type
type RecordType uint16

// ResponseCode This is the DNS RCODE
type ResponseCode uint8 //4 bit
