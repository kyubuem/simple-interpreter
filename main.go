package main

import (
	"github.com/kyubuem/interpreter/parser"
)

func main() {
	p := parser.NewSimpleParser("5-1")
	if _, err := p.Parse(); err != nil {
		panic(err)
	}
}
