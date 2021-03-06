package nativa

import (
	"OLC2/Compilador/interfaces"
)

type Abs struct {

	Expression 	interfaces.Expression
	Row			int
	Column		int

}

func NewAbs(Expression 	interfaces.Expression, Row int, Column int) Abs {
	instr := Abs{Expression, Row, Column}
	return instr
}

func (p Abs) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	gen.AddComment("Abs")

	result := p.Expression.Compilar(env, tree, gen)
	if result.Type != interfaces.INTEGER && result.Type != interfaces.FLOAT {
		excep := interfaces.NewException("Semantico", "Error en Abs, tipos de datos Incorrectos.", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		
	}
	EV := gen.NewLabel()
	EF := gen.NewLabel()

	gen.AddIf(result.Value, "0", "<", EV)
	gen.AddGoto(EF)


	gen.AddLabel(EV)

	temp := gen.NewTemp()	
	gen.AddExpression(temp, "0", "1", "-")
	gen.AddExpression(result.Value, result.Value, temp, "*")
	gen.AddLabel(EF)


	return interfaces.Value{Value: result.Value, IsTemp: true, Type: result.Type, TrueLabel: "", FalseLabel: ""}
}