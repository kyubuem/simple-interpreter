package token

import (
	"fmt"
	"strconv"
)

type TokenType int

const (
	Integer TokenType = iota
	Plus
	Minus
	Mul
	Div
	Eof
)

func (t TokenType) String() string {
	var tokenType = [...]string{
		"Integer",
		"Plus",
		"Eof",
	}
	return tokenType[int(t)%len(tokenType)]
}

func New(_type TokenType, _value string) Token {
	return Token{_type, _value}
}

type Token struct {
	tokenType  TokenType
	tokenValue string
}

func (t Token) Type() TokenType {
	return t.tokenType
}

func (t Token) Value() string {
	return t.tokenValue
}

func (t Token) ToInt() (int, error) {
	return strconv.Atoi(t.Value())
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%s, %s)", t.tokenType.String(), t.tokenValue)
}
