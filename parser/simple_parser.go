package parser

import (
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

func (s *SimpleParser) Parse() (ast.AST, error) {
	return s.parse()
}

func (s *SimpleParser) parse() (ast.AST, error) {
	return s.expr(), nil
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

func (s *SimpleParser) factor() ast.AST {
	t := s.curr
	var node ast.AST
	if t.Type() == token.Integer {
		s.eat(token.Integer)
		return ast.NewNumber(t)
	} else if t.Type() == token.Lparen {
		s.eat(token.Lparen)
		node = s.expr()
		s.eat(token.Rparen)
	}
	return node
}

func (s *SimpleParser) term() ast.AST {
	node := s.factor()
	for s.curr.Type() == token.Mul || s.curr.Type() == token.Div {
		t := s.curr
		if t.Type() == token.Mul {
			s.eat(token.Mul)
		} else if t.Type() == token.Div {
			s.eat(token.Div)
		}
		node = ast.NewBinaryOp(node, s.factor(), t)
	}
	return node
}

func (s *SimpleParser) expr() ast.AST {
	node := s.term()
	for s.curr.Type() == token.Plus || s.curr.Type() == token.Minus {
		t := s.curr
		if t.Type() == token.Plus {
			s.eat(token.Plus)
		} else if t.Type() == token.Minus {
			s.eat(token.Minus)
		}
		node = ast.NewBinaryOp(node, s.term(), t)
	}
	return node
}
