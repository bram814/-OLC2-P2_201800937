package casteo

import (
	"OLC2/Compilador/interfaces"
	// "fmt"
)

type Casteo struct {
	Expression  	interfaces.Expression
	Type  		interfaces.TypeExpression
	Row         int
	Column      int
}

func NewCasteo(expre interfaces.Expression, tipo interfaces.TypeExpression, row int, column int) Casteo {
	instr := Casteo{expre, tipo, row, column}
	return instr
}

func (p Casteo) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	var result interfaces.Value

	result = p.Expression.Compilar(env, tree, gen)
	if result.Type == interfaces.EXCEPTION {
		return result
	}
	

	result.Type = p.Type


	return interfaces.Value{Value: result.Value, IsTemp: result.IsTemp, Type: p.Type, TrueLabel: result.TrueLabel, FalseLabel: result.FalseLabel}

}
