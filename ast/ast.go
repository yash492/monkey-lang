package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Let Statement -
// let's take an example of let x = 5
// It will be parsed into let token, Name = x and with the value of 5
type LetStatement struct {
	Token token.Token // the token.LET token (let)
	Name  *Identifier // Name holds the variable name (x)
	Value Expression  // Value comes whatever comes after the assign operator (5)
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // Type - "IDENT" : Literal - variable name
	Value string      // Value - variable name
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r *ReturnStatement) statementNode() {}
func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}
