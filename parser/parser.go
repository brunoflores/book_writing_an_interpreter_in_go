package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Ty
	peekToken token.Ty

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := Parser{l: l}

	// Read two tokens so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return &p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.Program{}

	for !token.Is(p.curToken, token.Symbol{Id: token.EOF}) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return &program
}

func (p *Parser) parseStatement() ast.Statement {
	if token.Is(p.curToken, token.Keyword{Name: token.LET}) {
		return p.parseLetStatement()
	}
	return nil
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.Keyword{Name: token.ASSIGN}) {
		return nil
	}

	// stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal()}
	switch ident := p.curToken.(type) {
	case token.Ident:
		stmt.Name = &ast.Identifier{Token: ident, Value: ident.Name}
	default:
		return nil
	}

	if !token.Is(p.curToken, token.Keyword{Name: token.ASSIGN}) {
		return nil
	}

	// TODO
	for !token.Is(p.curToken, token.Symbol{Id: token.SEMICOLON}) {
		p.nextToken()
	}

	return &stmt
}

// func (p *Parser) curTokenIs(ty token.Type) bool {
// 	return p.curToken.Type() == ty
// }

// func (p *Parser) peekTokenIs(ty token.Type) bool {
// 	return p.peekToken.Type() == ty
// }

func (p *Parser) expectPeek(tok token.Ty) bool {
	if token.Is(p.peekToken, tok) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
