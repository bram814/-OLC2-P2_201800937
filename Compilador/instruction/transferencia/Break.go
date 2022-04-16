package transferencia

import (
	"OLC2/Compilador/ast"
	"OLC2/Compilador/interfaces"
	"fmt"
)


type Break struct {
	Expression 	interfaces.Expression
	Row 		int
	Column 		int
}

func NewBreak(expresion interfaces.Expression, row int, column int) Break {
	instr := Break{expresion, row, column}
	return instr
}


func (p Break) Compilar(env *interfaces.Environment, tree *ast.Arbol, gen *ast.Generator) interface{} {
	
	gen.AddComment("Break")
	var result interfaces.Value 
	
	pos := fmt.Sprintf("%v", tree.PosDisplay-1)
	display := tree.GetDisplay(pos)

	if p.Expression != nil {
		result = p.Expression.Compilar(env, tree, gen)

	}


	gen.AddGoto(display.LFinal)

	return result
}