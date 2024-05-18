package ast

type ExprStmt struct {
	Expression Expr
}

func (n ExprStmt) stmt() {}

type VarDeclStmt struct {
	VariableName  string
	IsConstant    bool
	AssignedValue Expr
	ExplicitType  Type
}

func (n VarDeclStmt) stmt() {}
