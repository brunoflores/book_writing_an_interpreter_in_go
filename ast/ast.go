package ast

import "interpreter/token"

type Node interface {
	TokenLiteral() int
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() int {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return 0
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()    {}
func (ls *LetStatement) TokenLiteral() int { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()   {}
func (i *Identifier) TokenLiteral() int { return i.Token.Literal }
