package arrays

import (

	"OLC2/Compilador/interfaces"
)

type ListIdentifier struct {
	Size   	interfaces.Expression
}

func NewListIdentifier(Size interfaces.Expression) ListIdentifier {
	exp := ListIdentifier{Size}
	return exp
}

func (p ListIdentifier) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	return p
}