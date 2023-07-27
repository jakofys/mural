package lexer

import (
	"bytes"

	"github.com/jakofys/mural/internal/token"
)

type lexer struct {
	buf *bytes.Reader
}

func (l *lexer) NextToken() token.Token {
	for ch := l.NextToken(){
		switch ch{
		case token.LSQBRACKET:
			return token.NewToken(token.LSQBRACKET,ch)
		}
	}
}
