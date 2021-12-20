package interpreter

import (
	"fmt"
	"strconv"

	"github.com/kyubuem/interpreter/lexer/token"
	"github.com/kyubuem/interpreter/parser/ast"
)

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

type Interpreter struct {
	value int
}

func (i *Interpreter) VisitBinaryOp(b ast.BinaryOp) {
	fmt.Println("VisitBinaryOp")
	b.Left().Accept(i)
	left := i.value

	b.Right().Accept(i)
	right := i.value

	switch b.Op() {
	case token.Plus:
		i.value = left + right
	case token.Minus:
		i.value = left - right
	case token.Mul:
		i.value = left * right
	case token.Div:
		i.value = left / right
	}
}

func (i *Interpreter) VisitNumber(n ast.Number) {
	fmt.Println("VisitNumber")
	i.value, _ = strconv.Atoi(n.GetValue())
}

func (i Interpreter) Result() int {
	return i.value
}
