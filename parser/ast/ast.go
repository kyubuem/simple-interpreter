package ast

type AST interface {
	Accept(Visitor)
}
