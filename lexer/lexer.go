package lexer

import "interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination (only ASCII)
}

// readChar gives the next character and advances our position in the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII for "NUL"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	tok := token.New(l.ch)
	l.readChar()
	return tok
}

func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()
	return &l
}
