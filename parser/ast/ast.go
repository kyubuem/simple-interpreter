package ast

type AST interface {
	visit() int
}
