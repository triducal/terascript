package ast

type File struct {
	Statements []Stmt
}

type Stmt interface {
	stmt()
}

type Expr interface {
	expr()
}

type Type interface {
	_type()
}
