package ast

type Project struct {
	Sprites map[string]Sprite
}

type Sprite struct {
	Name     string
	Vars     []string
	Lists    []string
	Costumes []string
	Body     []Stmt
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
