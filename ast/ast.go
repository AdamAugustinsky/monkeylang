package ast

import (
	"monkeylang/token"
)

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

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func NewProgramASTNode() *Program {
	return &Program{Statements: []Statement{}}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
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

// type ExpressionASTNode struct {
// 	Left     ExpressionASTNode
// 	Operator token.Token
// 	Right    ExpressionASTNode
// }
