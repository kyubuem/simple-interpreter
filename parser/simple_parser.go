package parser

import (
	"fmt"

	"github.com/kyubuem/interpreter/lexer"
	"github.com/kyubuem/interpreter/lexer/token"
	"github.com/kyubuem/interpreter/parser/ast"
	"github.com/pkg/errors"
)

func NewSimpleParser(text string) *SimpleParser {
	return &SimpleParser{
		lexer: lexer.NewSimpleLexer([]byte(text)),
	}
}

type SimpleParser struct {
	lexer lexer.Lexer
	curr  token.Token
}

func (s *SimpleParser) Parse() (*ast.AST, error) {
	return s.parse()
}

func (s *SimpleParser) parse() (*ast.AST, error) {
	curr, err := s.lexer.Next()
	if err != nil {
		return nil, err
	}
	s.curr = curr

	left := s.curr
	s.eat(token.Integer)

	op := s.curr
	if op.Type() == token.Plus {
		s.eat(token.Plus)
	} else {
		s.eat(token.Minus)
	}

	right := s.curr
	s.eat(token.Integer)

	leftVal, err := left.ToInt()
	if err != nil {
		return nil, err
	}
	rightVal, err := right.ToInt()
	if err != nil {
		return nil, err
	}
	if op.Type() == token.Plus {
		fmt.Println(leftVal + rightVal)
		return nil, nil
	}
	fmt.Println(leftVal - rightVal)
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
