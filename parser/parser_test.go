package parser

import (
	"interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len := len(program.Statements); len != 2 {
		t.Fatalf("Statements does not contain 2 statements. got=%d", len)
	}
}
