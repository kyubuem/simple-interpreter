package main

import (
	"github.com/kyubuem/interpreter/parser"
)

func main() {
	p := parser.NewSimpleParser("14 + 2 * 3 - 6 / 2")
	if _, err := p.Parse(); err != nil {
		panic(err)
	}
}
