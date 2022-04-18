package function

import (

	"OLC2/Compilador/interfaces"
)

type ListExpre struct {
	Id 	 	string
	Type 	interfaces.TypeExpression
	Row  	int
	Column  int

}

func NewListExpre(val string, tipo interfaces.TypeExpression, row int, column int) ListExpre {
	exp := ListExpre{val, tipo, row, column}
	return exp
}

func (p ListExpre) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	return p
}