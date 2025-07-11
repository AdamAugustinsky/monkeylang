package ast

import (
	"bytes"
	"monkeylang/token"
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

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func NewProgramASTNode() *Program {
	return &Program{Statements: []Statement{}}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value *Identifier
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token token.Token
	Value Expression
}

func (ls *ExpressionStatement) statementNode() {

}

func (ls *ExpressionStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Value != nil {
		return es.Value.String()
	}
	return ""
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) statementNode() {

}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

func NewIdentifierASTNode() *Identifier {
	return &Identifier{
		Token: token.Token{},
		Value: "",
	}
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
