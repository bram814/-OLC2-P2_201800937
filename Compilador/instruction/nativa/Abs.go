package nativa

import (
	"fmt"
	"OLC2/Compilador/interfaces"
)

type Abs struct {

	Id 		string
	Row		int
	Column	int

}

func NewAbs(Id string, Row int, Column int) Abs {
	instr := Abs{Id, Row, Column}
	return instr
}

func (p Abs) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	symbol := env.SearchSymbol(p.Id)

	if symbol.Type == interfaces.NULL {
		symbol = env.GetSymbol(p.Id)

		if symbol.Type == interfaces.NULL {
			excep := interfaces.NewException("Semantico", "No Existe ese Id "+p.Id + "(identificador).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

		}
	}

	gen.AddComment("Identificador")
	temp0 := gen.NewTemp()
	temp1 := gen.NewTemp()
	gen.AddExpression(temp0, "P", fmt.Sprintf("%v", symbol.Posicion), "+")
	gen.AddExpressionStack(temp1, temp0)

	gen.AddComment("Abs")

	EV := gen.NewLabel()
	EF := gen.NewLabel()

	gen.AddIf(temp1, "0", "<", EV)
	gen.AddGoto(EF)


	gen.AddLabel(EV)

	temp := gen.NewTemp()	
	gen.AddExpression(temp, "0", "1", "-")
	gen.AddExpression(temp1, temp1, temp, "*")
	gen.AddLabel(EF)


	return interfaces.Value{Value: temp1, IsTemp: true, Type: symbol.Value.(interfaces.Value).Type, TrueLabel: "", FalseLabel: ""}
}