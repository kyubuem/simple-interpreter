package parser

import (
	"fmt"

	"github.com/kyubuem/interpreter/lexer"
	"github.com/kyubuem/interpreter/lexer/token"
	"github.com/kyubuem/interpreter/parser/ast"
	"github.com/pkg/errors"
)

func NewSimpleParser(text string) *SimpleParser {
	parser := &SimpleParser{
		lexer: lexer.NewSimpleLexer(([]byte(text))),
	}
	t, _ := parser.lexer.Next()
	parser.curr = t

	return parser
}

type SimpleParser struct {
	lexer lexer.Lexer
	curr  token.Token
}

func (s *SimpleParser) Parse() (*ast.AST, error) {
	return s.parse()
}

func (s *SimpleParser) parse() (*ast.AST, error) {
	result := s.expr()
	fmt.Println(result)
	return nil, nil
}

func (s *SimpleParser) eat(t token.TokenType) error {
	if s.curr.Type() == t {
		var err error
		if s.curr, err = s.lexer.Next(); err != nil {
			return err
		}
		return nil
	}
	return errors.Errorf("cannot matched requested token type($s:%s)", s.curr.Type().String(), t.String())
}

func (s *SimpleParser) factor() int {
	t := s.curr
	s.eat(token.Integer)
	value, _ := t.ToInt()
	return value
}

func (s *SimpleParser) term() int {
	result := s.factor()
	for s.curr.Type() == token.Mul || s.curr.Type() == token.Div {
		t := s.curr
		if t.Type() == token.Mul {
			s.eat(token.Mul)
			result = result * s.factor()
		} else if t.Type() == token.Div {
			s.eat(token.Div)
			result = result / s.factor()
		}
	}
	return result
}

func (s *SimpleParser) expr() int {
	result := s.term()
	for s.curr.Type() == token.Plus || s.curr.Type() == token.Minus {
		t := s.curr
		if t.Type() == token.Plus {
			s.eat(token.Plus)
			result = result + s.term()
		} else if t.Type() == token.Minus {
			s.eat(token.Minus)
			result = result - s.term()
		}
	}
	return result
}
