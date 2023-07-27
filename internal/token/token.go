package token

// Token
type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(typ TokenType, literal string) Token {
	return Token{Type: typ, Literal: literal}
}
