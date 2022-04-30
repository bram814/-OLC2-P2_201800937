package transferencia

import (
	"fmt"
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


	

	if tree.IsReturn {

		pos := fmt.Sprintf("%v", tree.PosReturn)
		display := tree.GetDisplay(pos)
		if display.Type == result.Type {
			gen.AddStack("P", result.Value)
			gen.AddGoto(display.LFinal)
		} else {
			excep := interfaces.NewException("Semantico", "Error en return, tipos de datos incorrectos a Retornar ", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}


	} else {

		excep := interfaces.NewException("Semantico", "Error en return, tipos de datos incorrectos a Retornar ", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

	}

	
	return result
}
