package lexer

import (
	"interpreter/token"
	"io"
)

type Lexer struct {
	input       io.Reader
	bufsize     int
	read        int
	cbuf        int
	buf         []byte
	fname       string
	cch         int  // current column number
	lch         int  // current line number
	ch          byte // current char under examination (only ASCII)
	cread       int  // current reading position (after current char)
	eof_pending bool
	eof         bool
}

func (l *Lexer) ensureBuffer() {
	if !l.eof_pending && (l.cbuf >= l.read) {
		read, err := l.input.Read(l.buf)
		l.read = read
		l.cbuf = 0
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			l.eof_pending = true
		}
	}
}

func (l *Lexer) readChar() {
	l.ensureBuffer()

	if l.eof_pending && l.cbuf >= l.read {
		l.eof = true
	} else {
		l.ch = l.buf[l.cbuf]
		l.cbuf += 1

		if l.ch == '\n' {
			l.lch += 1
			l.cch = 0
			l.cread = 0
		} else {
			l.cch = l.cread
			l.cread += 1
		}
	}
}

func (l *Lexer) peekChar() byte {
	if (l.cch + 1) >= len(l.buf) {
		return 0
	} else {
		return l.buf[l.cread]
	}
}

func (l *Lexer) backtrack() {
	l.cch -= 1
	l.eof = false
	if l.cch > 0 {
		l.ch = l.buf[l.cch]
	}
}

func (l *Lexer) readWord() string {
	var word []byte
	for isLetter(l.ch) && !l.eof {
		word = append(word, l.ch)
		l.readChar()
	}
	return string(word)
}

func (l *Lexer) readNumber() string {
	var number []byte
	for isDigit(l.ch) && !l.eof {
		number = append(number, l.ch)
		l.readChar()
	}
	return string(number)
}

func (l *Lexer) readSymbol() string {
	var symbol []byte
	for !isLetter(l.ch) && !isDigit(l.ch) && !isWhitespace(l.ch) && !l.eof {
		symbol = append(symbol, l.ch)
		l.readChar()
	}
	return string(symbol)
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) && !l.eof {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Ty {
	l.skipWhitespace()

	pos := token.NewPosition(l.cch, l.lch, l.fname)
	if l.eof {
		return token.NewSymbol("eof", pos)
	}

	if isLetter(l.ch) {
		return token.New(l.readWord(), pos)
	} else if isDigit(l.ch) {
		return token.New(l.readNumber(), pos)
	} else {
		word := l.readSymbol()
		step := len(word)
		var s token.Ty
		for step > 0 {
			s = token.NewSymbol(word[0:step], pos)
			if !token.Is(s, token.Symbol{Id: token.ILLEGAL}) {
				return s
			}
			step -= 1
			l.backtrack()
		}
		return s
	}
}

func New(reader io.Reader, fname string) *Lexer {
	const bufsize = 1024
	l := Lexer{input: reader, fname: fname, bufsize: bufsize, buf: make([]byte, bufsize)}
	l.readChar()
	return &l
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
