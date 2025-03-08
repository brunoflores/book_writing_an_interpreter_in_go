package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"

	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var fromString = map[byte]TokenType{
	'=': ASSIGN,
	';': SEMICOLON,
	'(': LPAREN,
	')': RPAREN,
	',': COMMA,
	'+': PLUS,
	'{': LBRACE,
	'}': RBRACE,
	0:   EOF,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(ch byte) Token {
	return Token{Type: fromString[ch], Literal: string(ch)}
}
