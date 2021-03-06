package ast

import "../token"

// Number - Interger literal node
type Number struct {
	Token token.Token
	Value int
}

func (num *Number) expressionNode()      {}
func (num *Number) TokenLiteral() string { return num.Token.Literal }
func (num *Number) String() string       { return num.Token.Literal }

func (num *Number) Val() int {
	return int(num.Value)
}
