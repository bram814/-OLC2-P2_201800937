package structs

import (

	"OLC2/Compilador/interfaces"
)

type ListExpreDecla struct {
	Id 	 		string
	Expression 	interfaces.Expression
	Row  		int
	Column  	int

}

func NewListExpreDecla(val string, expre interfaces.Expression, row int, column int) ListExpreDecla {
	exp := ListExpreDecla{val, expre, row, column}
	return exp
}

func (p ListExpreDecla) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	return p
}