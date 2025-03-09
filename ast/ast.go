package ast

import "interpreter/token"

type Node interface {
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

type LetStatement struct {
	Token token.Ty
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

type Identifier struct {
	Token token.Ty
	Value string
}

func (i *Identifier) expressionNode() {}
