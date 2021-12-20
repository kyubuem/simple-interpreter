package ast

import "github.com/kyubuem/interpreter/lexer/token"

func NewBinaryOp(left, right AST, token token.Token) AST {
	return &BinaryOp{
		left:  left,
		right: right,
		token: token,
	}
}

type BinaryOp struct {
	left  AST
	right AST
	token token.Token
}

func (b BinaryOp) Accept(v Visitor) {
	v.VisitBinaryOp(b)
}

func (b *BinaryOp) Left() AST {
	return b.left
}

func (b *BinaryOp) Right() AST {
	return b.right
}

func (b BinaryOp) Op() token.TokenType {
	return b.token.Type()
}
