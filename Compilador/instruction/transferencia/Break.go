package transferencia

import (
	"OLC2/Compilador/interfaces"
	"fmt"
)

type Break struct {
	Expression interfaces.Expression
	Row        int
	Column     int
}

func NewBreak(expresion interfaces.Expression, row int, column int) Break {
	instr := Break{expresion, row, column}
	return instr
}

func (p Break) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	gen.AddComment("Break")
	var result interfaces.Value

	pos := fmt.Sprintf("%v", tree.PosDisplay-1)
	display := tree.GetDisplay(pos)

	if p.Expression != nil {
		result = p.Expression.Compilar(env, tree, gen)
		display.IsTemp = true
		temp := gen.NewTemp()
		gen.AddExpression1(temp, result.Value)

		tree.SetDisplayTemp(pos, temp, true, result.Type)
		display = tree.GetDisplay(pos)
		fmt.Println(display)

	}else {
		tree.SetDisplayTemp(pos, "-1", false, interfaces.BREAK)
		display = tree.GetDisplay(pos)
	}

	gen.AddGoto(display.LFinal)

	return result
}
