package structs

import (

	"OLC2/Compilador/interfaces"
)

type ListIdentifier struct {
	Id 	 	string
}

func NewListIdentifier(val string) ListIdentifier {
	exp := ListIdentifier{val}
	return exp
}

func (p ListIdentifier) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	return p
}