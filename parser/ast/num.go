package ast

import "github.com/kyubuem/interpreter/lexer/token"

type Num struct {
	token token.Token
	value int
}

func (n *Num) visit() int {
	return 0
}
