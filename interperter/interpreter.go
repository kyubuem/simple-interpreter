package interpreter

import "github.com/kyubuem/interpreter/parser/ast"

type Interpreter interface {
	Interpret(*ast.AST) (interface{}, error)
}
