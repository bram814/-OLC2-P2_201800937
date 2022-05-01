package nativa

import (
	// "fmt"
	// "reflect"
	"OLC2/Compilador/interfaces"
)

type ToString struct {

	Expression		interfaces.Expression
	Row				int
	Column			int

}

func NewToString(Expression	interfaces.Expression, Row int, Column int) ToString {
	instr := ToString{Expression, Row, Column}
	return instr
}

func (p ToString) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {


	result := p.Expression.Compilar(env, tree, gen)

	if result.Type != interfaces.STRING && result.Type != interfaces.STR {
		excep := interfaces.NewException("Semantico", "Error en ToString, tipo incorrecto, no es un String.", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	return interfaces.Value{Value: result.Value, IsTemp: result.IsTemp, Type: interfaces.STRING, TrueLabel: "", FalseLabel: ""}
}