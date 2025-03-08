package token

import (
	"strconv"
)

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	BANG     = "BANG"
	ASTERISK = "ASTERISK"
	SLASH    = "/"

	LT = "LT"
	GT = "GT"

	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Type string

var symbol = map[string]Type{
	"=": ASSIGN,
	";": SEMICOLON,
	"(": LPAREN,
	")": RPAREN,
	",": COMMA,
	"+": PLUS,
	"{": LBRACE,
	"}": RBRACE,
	"!": BANG,
	"-": MINUS,
	"*": ASTERISK,
	"/": SLASH,
	"<": LT,
	">": GT,
}

var keyword = map[string]Type{
	"let": LET,
	"fn":  FUNCTION,
}

type Token struct {
	Symbol  Type
	Keyword Type
	Ident   string
	Literal int
}

func New(word string) Token {
	if s, ok := symbol[word]; ok {
		return Token{Symbol: s}
	} else if k, ok := keyword[word]; ok {
		return Token{Keyword: k}
	} else {
		i, err := strconv.Atoi(word)
		if err != nil {
			return Token{Ident: word}
		}
		return Token{Literal: i}
	}
}
