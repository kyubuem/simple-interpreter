package lexer

import (
	"unicode"

	"github.com/kyubuem/interpreter/lexer/token"
	"github.com/pkg/errors"
)

func NewSimpleLexer(text []byte) *SimpleLexer {
	return &SimpleLexer{
		text: text,
		curr: text[0],
		pos:  0,
	}
}

type SimpleLexer struct {
	text []byte
	curr byte
	pos  int
}

func (s *SimpleLexer) Next() (token.Token, error) {
	return s.next()
}

func (s *SimpleLexer) next() (token.Token, error) {
	for s.curr != 0 {
		if unicode.IsSpace(rune(s.curr)) {
			s.skipWhitespace()
			continue
		}

		if unicode.IsDigit(rune(s.curr)) {
			return token.New(token.Integer, s.integer()), nil
		}

		if s.curr == '+' {
			s.advance()
			return token.New(token.Plus, "+"), nil
		}

		if s.curr == '-' {
			s.advance()
			return token.New(token.Minus, "-"), nil
		}

		if s.curr == '*' {
			s.advance()
			return token.New(token.Mul, "*"), nil
		}

		if s.curr == '/' {
			s.advance()
			return token.New(token.Div, "/"), nil
		}

		if s.curr == '(' {
			s.advance()
			return token.New(token.Lparen, "("), nil
		}

		if s.curr == ')' {
			s.advance()
			return token.New(token.Rparen, ")"), nil
		}

		return token.Token{}, errors.Errorf("Cannot tokenize for requested character(%c)", s.curr)
	}

	return token.New(token.Eof, ""), nil
}

func (s *SimpleLexer) advance() {
	s.pos += 1
	if s.pos > len(s.text)-1 {
		s.curr = 0
		return
	}
	s.curr = s.text[s.pos]
}

func (s *SimpleLexer) skipWhitespace() {
	for s.curr != 0 && unicode.IsSpace(rune(s.curr)) {
		s.advance()
	}
}

func (s *SimpleLexer) integer() string {
	result := make([]byte, 0)
	for s.curr != 0 && unicode.IsDigit(rune(s.curr)) {
		result = append(result, s.curr)
		s.advance()
	}
	return string(result)
}
