package transferencia

import (
	"OLC2/Compilador/interfaces"
)

type Return struct {
	Expression interfaces.Expression
	Row        int
	Column     int
}

func NewReturn(expresion interfaces.Expression, row int, column int) Return {
	instr := Return{expresion, row, column}
	return instr
}

func (p Return) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	gen.AddComment("Return")
	var result interfaces.Value
	
	result = p.Expression.Compilar(env, tree, gen)
	gen.AddStack("P", result.Value)
	return result
}
