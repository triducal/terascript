package ast

import "terascript/src/lexer"

//LITERALS

type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {}

//COMPLEX

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (n BinaryExpr) expr() {}

type PrefixExpr struct {
	Operator lexer.Token
	Right    Expr
}

func (n PrefixExpr) expr() {}

type AssignmentExpr struct {
	Assigne  Expr
	Operator lexer.Token
	Right    Expr
}

func (n AssignmentExpr) expr() {}
