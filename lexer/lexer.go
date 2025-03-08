package lexer

import (
	"interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination (only ASCII)
	eof          bool
}

// readChar gives the next character and advances our position in the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.eof = true
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readWord() string {
	position := l.position
	for isLetter(l.ch) && !l.eof {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) && !l.eof {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for (l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r') && !l.eof {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	if l.eof {
		return token.Token{Symbol: token.EOF}
	}

	b := l.ch
	if isLetter(b) {
		return token.New(l.readWord())
	} else if isDigit(b) {
		return token.New(l.readNumber())
	} else {
		defer l.readChar()
		return token.New(string(b))
	}
}

func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()
	return &l
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
