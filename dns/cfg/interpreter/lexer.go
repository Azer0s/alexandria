package interpreter

import (
	"log"
	"strconv"
)

type TokenType uint8

const (
	IDENTIFIER  TokenType = 0
	STRING      TokenType = 1
	NUMBER      TokenType = 2
	IN          TokenType = 3
	NATIVE      TokenType = 4
	ZONE        TokenType = 5
	ARROW       TokenType = 6
	OPEN_BRACE  TokenType = 7
	CLOSE_BRACE TokenType = 8
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
}

func isLetter(ch uint8) bool {
	if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') {
		return false
	}
	return true
}

func isNumber(ch uint8) bool {
	if _, err := strconv.Atoi(string(ch)); err == nil {
		return true
	}

	return false
}

func set(i *int, ch *uint8, cfg string, inc int) {
	*i += inc
	*ch = cfg[*i]
}

func LexConfig(config string) []Token {
	line := 1

	tokens := make([]Token, 0)

	for i := 0; i < len(config); i++ {
		buffer := make([]byte, 0)
		var tokenType TokenType

		ch := config[i]

		if isLetter(ch) {
			tokenType = IDENTIFIER

			for isLetter(ch) {
				buffer = append(buffer, ch)
				set(&i, &ch, config, 1)
			}
			set(&i, &ch, config, -1)

			str := string(buffer)
			if str == "in" {
				tokenType = IN
			} else if str == "native" {
				tokenType = NATIVE
			} else if str == "zone" {
				tokenType = ZONE
			}
		} else if isNumber(ch) {
			tokenType = NUMBER

			for isNumber(ch) {
				buffer = append(buffer, ch)
				set(&i, &ch, config, 1)
			}
			set(&i, &ch, config, -1)
		} else if ch == '=' {
			set(&i, &ch, config, 1)

			if ch == '>' {
				tokenType = ARROW
			} else {
				log.Fatalf("Error parsing config! Expected '>'! Line %d", line)
			}
			buffer = []byte("=>")
		} else if ch == '{' {
			tokenType = OPEN_BRACE
			buffer = []byte("{")
		} else if ch == '}' {
			tokenType = CLOSE_BRACE
			buffer = []byte("}")
		} else if ch == '"' {
			tokenType = STRING
			set(&i, &ch, config, 1)

			for ch != '"' {
				buffer = append(buffer, ch)
				set(&i, &ch, config, 1)
			}
		} else if ch == '\n' {
			line++
		}

		if ch != ' ' && ch != '\n' && ch != '\t' {
			tokens = append(tokens, Token{
				Type:  tokenType,
				Value: string(buffer),
				Line:  line,
			})
		}
	}

	return tokens
}
