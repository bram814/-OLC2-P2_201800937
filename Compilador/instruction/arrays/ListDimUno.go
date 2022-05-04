package arrays

import (
	// "fmt"
	"OLC2/Compilador/interfaces"
)

type ListDimUno struct {
	Expression    interfaces.Expression
}

func NewListDimUno(Expression interfaces.Expression) ListDimUno {
	instr := ListDimUno{Expression}
	return instr
}

func (p ListDimUno) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {
	return p
}