package parser

import "github.com/kyubuem/interpreter/parser/ast"

type Parser interface {
	Parse() (*ast.AST, error)
}
