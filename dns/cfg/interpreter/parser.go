package interpreter

import (
	"github.com/Azer0s/alexandria/dns/enums/fields"
	"github.com/Azer0s/alexandria/dns/enums/record_type"
	"log"
	"net"
	"strconv"
	"strings"
)

func getResourceTypeByIdentifier(identifier string, line int) fields.RecordType {
	identifier = strings.ToLower(identifier)

	if identifier == "cname" {
		return record_type.CNAME
	} else if identifier == "a" {
		return record_type.A
	} else if identifier == "txt" {
		return record_type.TXT
	} else {
		log.Fatalf("Unexpected identifier %s! Line %d", identifier, line)
	}

	return 0
}

type KV struct {
	Name       string
	Value      string
	Native     bool
	Recursive  bool
	IP         bool
	TimeToLive uint32
}

func parseKV(tokens []Token, i *int, canRecurse bool) []KV {
	var token Token
	next(&token, &tokens, i)

	if token.Type == OPEN_BRACE && canRecurse {
		kvs := make([]KV, 0)
		for token.Type != CLOSE_BRACE {
			kvs = append(kvs, parseKV(tokens, i, false)...)
			token = tokens[*i]
		}

		return kvs
	} else if token.Type == STRING || token.Type == NATIVE {
		kv := KV{}
		if token.Type == NATIVE {
			kv.Native = true
		}

		kv.Name = token.Value

		next(&token, &tokens, i)

		if token.Type == NUMBER {
			ttl, _ := strconv.Atoi(token.Value)
			kv.TimeToLive = uint32(ttl)
		}

		if token.Type != ARROW {
			log.Fatalf("Expected '=>', got %s! Line %d", token.Value, token.Line)
		}

		next(&token, &tokens, i)

		if token.Type == IN {
			kv.Recursive = true
			next(&token, &tokens, i)
		} else if strings.ToLower(token.Value) == "ip" {
			kv.IP = true
			next(&token, &tokens, i)
		}

		if token.Type != STRING {
			log.Fatalf("Expected string, got %s! %d", token.Value, token.Line)
		}

		kv.Value = token.Value
		return []KV{kv}
	} else if token.Type == ARROW {
		log.Fatalf("Expected identifier or string! Line %d", token.Line)
	}

	return nil
}

func parseZone(tokens []Token, i *int, defaultTtl uint32) Zone {
	var token Token
	next(&token, &tokens, i)

	if token.Type != STRING {
		log.Fatalf("Expected string, got %s! Line %d", token.Value, token.Line)
	}

	name := token.Value

	next(&token, &tokens, i)

	if token.Type != OPEN_BRACE {
		log.Fatalf("Expected '{', got %s! Line %d", token.Value, token.Line)
	}

	braces := 1

	entries := make([]Entry, 0)
	zones := make([]Zone, 0)

	for braces != 0 {
		next(&token, &tokens, i)

		if token.Type == CLOSE_BRACE {
			braces--
		} else if token.Type == IDENTIFIER {
			entryType := getResourceTypeByIdentifier(token.Value, token.Line)
			kvs := parseKV(tokens, i, true)

			for _, kv := range kvs {
				bytes := make([]byte, 0)

				if kv.IP {
					ip := net.ParseIP(kv.Value)
					if ip == nil {
						log.Fatalf("Invalid IP %s!", kv.Value)
					}

					if strings.Contains(kv.Value, ":") {
						bytes = ip.To16()
					} else {
						bytes = ip.To4()
					}
				} else {
					bytes = []byte(kv.Value)
				}

				if kv.TimeToLive == 0 {
					kv.TimeToLive = defaultTtl
				}

				entries = append(entries, Entry{
					Type:       entryType,
					Recursive:  kv.Recursive,
					Native:     kv.Native,
					Name:       kv.Name,
					Value:      bytes,
					TimeToLive: kv.TimeToLive,
				})
			}
		} else if token.Type == ZONE {
			zones = append(zones, parseZone(tokens, i, defaultTtl))
		} else {
			log.Fatalf("Unexpected %s! %d", token.Value, token.Line)
		}
	}

	return Zone{
		FQDN:    name,
		Entries: entries,
		Zones:   zones,
	}
}

func next(token *Token, tokens *[]Token, i *int) {
	*i++
	*token = (*tokens)[*i]
}

func ParseConfig(tokens []Token, ttl uint32) []Zone {
	zones := make([]Zone, 0)

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token.Type != ZONE {
			log.Fatalf("Expected 'zone', got %s! Line %d", token.Value, token.Line)
		}

		zones = append(zones, parseZone(tokens, &i, ttl))
	}

	return zones
}
