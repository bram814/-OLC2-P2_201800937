package transferencia

import (
	"OLC2/Compilador/ast"
	"OLC2/Compilador/interfaces"
	"fmt"
)

type Continue struct {
	Row 	int
	Column 	int
}

func NewContinue(row int, column int) Continue {
	instr := Continue{row, column}
	return instr
}


func (p Continue) Compilar(env *interfaces.Environment, tree *ast.Arbol, gen *ast.Generator) interface{} {

	gen.AddComment("Continue")

	pos := fmt.Sprintf("%v", tree.PosDisplay-1)
	display := tree.GetDisplay(pos)

	gen.AddGoto(display.LInicio)
	return p
}