package arrays

import (
	// "fmt"
	"OLC2/Compilador/interfaces"
)

type DimUno struct {
	Type      	interfaces.TypeExpression
	Expression  interfaces.Expression
}

func NewDimUno(tipo interfaces.TypeExpression, Expression  interfaces.Expression) DimUno {
	instr := DimUno{tipo, Expression}
	return instr
}

func (p DimUno) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {
	return p
}