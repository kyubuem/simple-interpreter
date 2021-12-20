package ast

type Visitor interface {
	VisitBinaryOp(BinaryOp)
	VisitNumber(Number)
}
