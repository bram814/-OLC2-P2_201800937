package variable

import (
	"OLC2/Compilador/interfaces"
	"fmt"
)

type Assignment struct {
	Id        string
	Expresion interfaces.Expression
	Row       int
	Column    int
}

func NewAssignment(id string, val interfaces.Expression, row int, column int) Assignment {
	instr := Assignment{id, val, row, column}
	return instr
}

func (p Assignment) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	/* Buscar si el id ya existe */
	symbol := env.GetSymbol(p.Id)

	if symbol.Type == interfaces.NULL {
		excep := interfaces.NewException("Semantico", "No Existe ese Id "+p.Id, p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	var result interfaces.Value
	result = p.Expresion.Compilar(env, tree, gen)

	if symbol.IsMut {
		gen.AddComment("Asignacion")
		// fmt.Println(symbol.Type)
		// fmt.Println(symbol.Value.(interfaces.Value).Type )
		if symbol.Type == result.Type || symbol.Type == interfaces.NULL{
			symbol.IsMut = true
			env.SetSymbol(p.Id, result, true, symbol.Posicion)
			gen.AddStack(fmt.Sprintf("%v", symbol.Posicion), result.Value)

		} else {
			excep := interfaces.NewException("Semantico", "No se puede asignar, tipo de datos diferentes.", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	
		}

	} else {
		excep := interfaces.NewException("Semantico", "No se puede asignar a "+p.Id+", no es mutable.", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	return result.Value

}
