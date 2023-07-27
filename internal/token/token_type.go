package token

type TokenType rune

const (
	EOF TokenType = iota - 1
	ILLEGAL

	// identifier + literal
	SECTION
	DESC // when closing angle bracket
	IDENT

	// keywords
	OPTIONS
	ARGUMENTS
	HEAD
	COMMANDS

	// types
	STRING // when find string
	BOOL   // when find bool

	COLON
	FLOAT
	NUMERIC
	HYPHEN
	LSQBRACKET
	RSQBRACKET
	ELLIPSIS
	COMMA
	INTERR // when found ?
	ASSIGN
	GT
)

// all available keywords
var keywords = map[string]TokenType{
	"commands":  COMMANDS,
	"head":      HEAD,
	"arguments": ARGUMENTS,
	"options":   OPTIONS,
}

// AsTokenType return keyword token type if correspond, IDENT either
func AsTokenType(str string) TokenType {
	if t, ok := keywords[str]; ok {
		return t
	}
	return IDENT
}
