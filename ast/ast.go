package ast

import (
	"monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
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
	Statements []Statement `json:"statements"`
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out strings.Builder
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Let Statement -
// let's take an example of let x = 5
// It will be parsed into let token, Name = x and with the value of 5
type LetStatement struct {
	Token token.Token `json:"token"` // the token.LET token (let)
	Name  *Identifier `json:"name"`  // Name holds the variable name (x)
	Value Expression  `json:"value"` // Value comes whatever comes after the assign operator (5)
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) String() string {
	var out strings.Builder
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.TokenLiteral())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token `json:"token"` // Type - "IDENT" : Literal - variable name
	Value string      `json:"value"` // Value - variable name
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

type AssignStatement struct {
	Token token.Token `json:"token"` // Token of an identifier
	Name  *Identifier `json:"name"`  // Name holds the variable name (x)
	Value Expression  `json:"value"` // Value comes whatever comes after the assign operator (5)
}

func (ls *AssignStatement) statementNode() {}
func (ls *AssignStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *AssignStatement) String() string {
	var out strings.Builder
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type ReturnStatement struct {
	Token token.Token `json:"token"`
	Value Expression  `json:"value"`
}

func (r *ReturnStatement) statementNode() {}
func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) String() string {
	var out strings.Builder
	out.WriteString(r.TokenLiteral())
	if r.Value != nil {

		out.WriteString(" " + r.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type WhileStatement struct {
	Token       token.Token     `json:"token"` // The 'while' token
	Condition   Expression      `json:"condition"`
	Consequence *BlockStatement `json:"consequence"`
}

func (ws *WhileStatement) statementNode() {}
func (ws *WhileStatement) TokenLiteral() string {
	return ws.Token.Literal
}
func (ws *WhileStatement) String() string {
	out := strings.Builder{}
	out.WriteString("if")
	out.WriteString(ws.Condition.String())

	if ws.Consequence != nil {
		out.WriteString(" ")
		out.WriteString(ws.Consequence.String())
	}
	return out.String()

}

type ExpressionStatement struct {
	Token      token.Token `json:"token"` // the first token of the expression
	Expression Expression  `json:"expression"`
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token `json:"token"`
	Value int64       `json:"value"`
}

func (i *IntegerLiteral) expressionNode() {}
func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}
func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token `json:"token"` // The prefix token eg: !, -
	Operator string      `json:"operator"`
	Right    Expression  `json:"right"`
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	sb := strings.Builder{}
	sb.WriteString("(")
	sb.WriteString(pe.Operator)
	sb.WriteString(pe.Right.String())
	sb.WriteString(")")
	return sb.String()
}

type InfixExpression struct {
	Token    token.Token `json:"token"` // operator token
	Left     Expression  `json:"left"`
	Operator string      `json:"operator"`
	Right    Expression  `json:"right"`
}

func (oe *InfixExpression) expressionNode() {}
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

func (oe *InfixExpression) String() string {
	sb := strings.Builder{}
	sb.WriteString("(")
	sb.WriteString(oe.Left.String())
	sb.WriteString(" " + oe.Operator + " ")
	sb.WriteString(oe.Right.String())
	sb.WriteString(")")
	return sb.String()
}

type Boolean struct {
	Token token.Token `json:"token"`
	Value bool        `json:"value"`
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

type IfExpression struct {
	Token       token.Token     `json:"token"` // The 'if' token
	Condition   Expression      `json:"condition"`
	Consequence *BlockStatement `json:"consequence"`
	Alternative *BlockStatement `json:"alternative"`
}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) String() string {
	out := strings.Builder{}
	out.WriteString("if")
	out.WriteString(ie.Condition.String())

	if ie.Consequence != nil {
		out.WriteString(" ")
		out.WriteString(ie.Consequence.String())
	}

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token `json:"token"` // the '{' token
	Statements []Statement `json:"statements"`
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	out := strings.Builder{}

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token     `json:"token"` // The fn token
	Parameters []*Identifier   `json:"parameters"`
	Body       *BlockStatement `json:"body"`
}

func (fl *FunctionLiteral) String() string {
	var out strings.Builder
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

type CallExpression struct {
	Token     token.Token  `json:"token"` // T`he '(' token
	Function  Expression   `json:"function"`
	Arguments []Expression `json:"arguments"`
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	out := strings.Builder{}

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ","))
	out.WriteString(")")

	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}


func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) TokenLiteral() string {return sl.Token.Literal}
func (sl *StringLiteral) String() string {return sl.Token.Literal}
