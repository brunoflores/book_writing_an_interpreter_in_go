package lexer

import (
	"testing"

	"interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `
let answer =     42;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/5;
5 < 10 > 5;
`

	tests := []token.Token{
		{Keyword: token.LET},
		{Ident: "answer"},
		{Symbol: token.ASSIGN},
		{Literal: 42},
		{Symbol: token.SEMICOLON},
		{Keyword: token.LET},
		{Ident: "ten"},
		{Symbol: token.ASSIGN},
		{Literal: 10},
		{Symbol: token.SEMICOLON},
		{Keyword: token.LET},
		{Ident: "add"},
		{Symbol: token.ASSIGN},
		{Keyword: token.FUNCTION},
		{Symbol: token.LPAREN},
		{Ident: "x"},
		{Symbol: token.COMMA},
		{Ident: "y"},
		{Symbol: token.RPAREN},
		{Symbol: token.LBRACE},
		{Ident: "x"},
		{Symbol: token.PLUS},
		{Ident: "y"},
		{Symbol: token.SEMICOLON},
		{Symbol: token.RBRACE},
		{Symbol: token.SEMICOLON},
		{Keyword: token.LET},
		{Ident: "result"},
		{Symbol: token.ASSIGN},
		{Ident: "add"},
		{Symbol: token.LPAREN},
		{Ident: "five"},
		{Symbol: token.COMMA},
		{Ident: "ten"},
		{Symbol: token.RPAREN},
		{Symbol: token.SEMICOLON},
		{Symbol: token.BANG},
		{Symbol: token.MINUS},
		{Symbol: token.SLASH},
		{Literal: 5},
		{Symbol: token.SEMICOLON},
		{Literal: 5},
		{Symbol: token.LT},
		{Literal: 10},
		{Symbol: token.GT},
		{Literal: 5},
		{Symbol: token.SEMICOLON},
		{Symbol: token.EOF},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok != tt {
			t.Fatalf("tests[%d] - expected=%v, got=%v", i, tt, tok)
		}
	}
}
