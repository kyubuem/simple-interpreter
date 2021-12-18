package ast

import "github.com/kyubuem/interpreter/lexer/token"

type BinOp struct {
	left  *Num
	token token.Token
	right *Num
}

func (b *BinOp) visit() int {
	return 0
}
