package ast

type SymbolType struct {
	Name string
}

func (t SymbolType) _type() {}

type ArrayType struct {
	Underlying Type
}

func (t ArrayType) _type() {}
