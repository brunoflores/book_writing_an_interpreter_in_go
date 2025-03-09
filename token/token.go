package token

import (
	"strconv"
)

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	BANG     = "BANG"
	ASTERISK = "ASTERISK"
	SLASH    = "/"

	LT     = "LT"
	GT     = "GT"
	EQ     = "EQ"
	NOT_EQ = "NOT_EQ"

	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type Ty interface {
	ty()
}

type Symbol struct {
	Id string
}

func (s Symbol) ty() {}
func (s Symbol) equal(this Symbol) bool {
	return s.Id == this.Id
}

type Keyword struct {
	Name string
}

func (k Keyword) ty() {}
func (k Keyword) equal(this Keyword) bool {
	return k.Name == this.Name
}

type Ident struct {
	Name string
}

func (i Ident) ty() {}
func (i Ident) equal(this Ident) bool {
	return i.Name == this.Name
}

type LiteralString struct {
	String string
}

func (s LiteralString) ty() {}
func (s LiteralString) equal(this LiteralString) bool {
	return s.String == this.String
}

type LiteralInt struct {
	Int int
}

func (i LiteralInt) ty() {}
func (i LiteralInt) equal(this LiteralInt) bool {
	return i.Int == this.Int
}

var symbol = map[string]string{
	"=":   ASSIGN,
	"==":  EQ,
	"!=":  NOT_EQ,
	";":   SEMICOLON,
	"(":   LPAREN,
	")":   RPAREN,
	",":   COMMA,
	"+":   PLUS,
	"{":   LBRACE,
	"}":   RBRACE,
	"!":   BANG,
	"-":   MINUS,
	"*":   ASTERISK,
	"/":   SLASH,
	"<":   LT,
	">":   GT,
	"eof": EOF,
}

var keyword = map[string]string{
	"let":    LET,
	"fn":     FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func New(word string) Ty {
	if s, ok := symbol[word]; ok {
		return Symbol{Id: s}
	} else if k, ok := keyword[word]; ok {
		return Keyword{Name: k}
	} else {
		i, err := strconv.Atoi(word)
		if err != nil {
			return Ident{Name: word}
		}
		return LiteralInt{Int: i}
	}
}

func NewSymbol(word string) Ty {
	if s, ok := symbol[word]; ok {
		return Symbol{Id: s}
	} else {
		return Symbol{Id: ILLEGAL}
	}
}

func Is(subject, target Ty) bool {
	switch subject := subject.(type) {
	case Symbol:
		switch target := target.(type) {
		case Symbol:
			return subject.equal(target)
		}
	case Keyword:
		switch target := target.(type) {
		case Keyword:
			return subject.equal(target)
		}
	case Ident:
		switch target := target.(type) {
		case Ident:
			return subject.equal(target)
		}
	case LiteralInt:
		switch target := target.(type) {
		case LiteralInt:
			return subject.equal(target)
		}
	case LiteralString:
		switch target := target.(type) {
		case LiteralString:
			return subject.equal(target)
		}
	}
	return false
}
