package lexer

import "github.com/kyubuem/interpreter/lexer/token"

type Lexer interface {
	Next() (token.Token, error)
}
