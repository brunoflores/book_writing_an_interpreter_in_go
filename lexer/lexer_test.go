package lexer

import (
	"strings"
	"testing"

	"interpreter/token"

	"github.com/stretchr/testify/assert"
)

func assertToken(want token.Ty) func(*testing.T, int, token.Ty) {
	return func(t *testing.T, i int, got token.Ty) {
		assert.True(t, token.Is(got, want),
			"test %d: expected %v, got %v", i, want, got)
	}
}

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

if (5 < 10) {
  return true;
} else {
  return false;
}

10 == 10;
10 != 9;
`

	// One test per token expected
	tests := []func(*testing.T, int, token.Ty){
		assertToken(token.Keyword{Name: token.LET}),
		assertToken(token.Ident{Name: "answer"}),
		assertToken(token.Symbol{Id: token.ASSIGN}),
		assertToken(token.LiteralInt{Int: 42}),
	}

	// {Keyword: token.LET},
	// {Ident: "answer"},
	// {Symbol: token.ASSIGN},
	// {Literal: 42},
	// {Symbol: token.SEMICOLON},
	// {Keyword: token.LET},
	// {Ident: "ten"},
	// {Symbol: token.ASSIGN},
	// {Literal: 10},
	// {Symbol: token.SEMICOLON},
	// {Keyword: token.LET},
	// {Ident: "add"},
	// {Symbol: token.ASSIGN},
	// {Keyword: token.FUNCTION},
	// {Symbol: token.LPAREN},
	// {Ident: "x"},
	// {Symbol: token.COMMA},
	// {Ident: "y"},
	// {Symbol: token.RPAREN},
	// {Symbol: token.LBRACE},
	// {Ident: "x"},
	// {Symbol: token.PLUS},
	// {Ident: "y"},
	// {Symbol: token.SEMICOLON},
	// {Symbol: token.RBRACE},
	// {Symbol: token.SEMICOLON},
	// {Keyword: token.LET},
	// {Ident: "result"},
	// {Symbol: token.ASSIGN},
	// {Ident: "add"},
	// {Symbol: token.LPAREN},
	// {Ident: "five"},
	// {Symbol: token.COMMA},
	// {Ident: "ten"},
	// {Symbol: token.RPAREN},
	// {Symbol: token.SEMICOLON},
	// {Symbol: token.BANG},
	// {Symbol: token.MINUS},
	// {Symbol: token.SLASH},
	// {Literal: 5},
	// {Symbol: token.SEMICOLON},
	// {Literal: 5},
	// {Symbol: token.LT},
	// {Literal: 10},
	// {Symbol: token.GT},
	// {Literal: 5},
	// {Symbol: token.SEMICOLON},
	// {Keyword: token.IF},
	// {Symbol: token.LPAREN},
	// {Literal: 5},
	// {Symbol: token.LT},
	// {Literal: 10},
	// {Symbol: token.RPAREN},
	// {Symbol: token.LBRACE},
	// {Keyword: token.RETURN},
	// {Keyword: token.TRUE},
	// {Symbol: token.SEMICOLON},
	// {Symbol: token.RBRACE},
	// {Keyword: token.ELSE},
	// {Symbol: token.LBRACE},
	// {Keyword: token.RETURN},
	// {Keyword: token.FALSE},
	// {Symbol: token.SEMICOLON},
	// {Symbol: token.RBRACE},
	// {Literal: 10},
	// {Symbol: token.EQ},
	// {Literal: 10},
	// {Symbol: token.SEMICOLON},
	// {Literal: 10},
	// {Symbol: token.NOT_EQ},
	// {Literal: 9},
	// {Symbol: token.SEMICOLON},
	// {Symbol: token.EOF},

	reader := strings.NewReader(input)
	l := New(reader, "-")
	for i, tt := range tests {
		tok := l.NextToken()
		tt(t, i, tok)
	}
}
