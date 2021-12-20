package main

import (
	"fmt"

	interpreter "github.com/kyubuem/interpreter/interperter"
	"github.com/kyubuem/interpreter/parser"
)

func main() {
	p := parser.NewSimpleParser("7 * 6 - 10")
	if node, err := p.Parse(); err != nil {
		panic(err)
	} else {
		i := interpreter.NewInterpreter()
		node.Accept(i)
		fmt.Println(i.Result())
	}
}
