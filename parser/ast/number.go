package ast

import (
	"github.com/kyubuem/interpreter/lexer/token"
)

func NewNumber(token token.Token) AST {
	return &Number{
		token: token,
		value: token.Value(),
	}
}

type Number struct {
	token token.Token
	value string
}

func (n Number) Accept(v Visitor) {
	v.VisitNumber(n)
}

func (n Number) GetValue() string {
	return n.value
}
