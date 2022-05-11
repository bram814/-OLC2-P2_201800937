package nativa

import (
	"fmt"
	"reflect"
	"OLC2/Compilador/interfaces"
)

type Potencia struct {

	Base   			interfaces.Expression
	Exponente 	    interfaces.Expression
	Type 			interfaces.TypeExpression
	Row				int
	Column			int

}

func NewPotencia(Base interfaces.Expression, Exponente interfaces.Expression, Type interfaces.TypeExpression, Row int, Column int) Potencia {
	instr := Potencia{Base, Exponente, Type, Row, Column}
	return instr
}

func (p Potencia) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	
	gen.AddComment("Potencia")


	var base interfaces.Value
	var exponente interfaces.Value

	if reflect.TypeOf(p.Exponente).String() == "function.LlamadaExpresion" {

		base = p.Base.Compilar(env, tree, gen)

		if base.Type != p.Type {
			excep := interfaces.NewException("Semantico", "Error en Potencia, tipos de datos Incorrectos (Base).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}

		temp := gen.NewTemp()
		gen.AddComment("Guardando Temporal")
		gen.AddExpression(temp,"P", fmt.Sprintf("%v", env.Posicion), "+")
		aux := env.Posicion
		gen.AddStack(temp, base.Value)
		tree.AddTableSymbol(*interfaces.NewTableSymbol(base.Value,"Temporal","Local", p.Row, p.Column, "--", fmt.Sprintf("%v", env.Posicion)))
		env.NewPos()

		exponente = p.Exponente.Compilar(env, tree, gen)

		if exponente.Type != p.Type {
			excep := interfaces.NewException("Semantico", "Error en Potencia, tipos de datos Incorrectos (Exponente).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}

		temp = gen.NewTemp()
		gen.AddExpression(temp, "P", fmt.Sprintf("%v", aux), "+")
		gen.AddExpressionStack(base.Value,temp)


	} else {

		base = p.Base.Compilar(env, tree, gen)

		if base.Type != p.Type  {
			excep := interfaces.NewException("Semantico", "Error en Potencia, tipos de datos Incorrectos (Base).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		 
		}
		exponente = p.Exponente.Compilar(env, tree, gen)

		if exponente.Type != p.Type {
			excep := interfaces.NewException("Semantico", "Error en Potencia, tipos de datos Incorrectos (Exponente).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}
	}

	temp := gen.NewTemp()
	temp1 := gen.NewTemp()
	

	Linicio := gen.NewLabel()
	Lfinal  := gen.NewLabel()

	gen.AddExpression1(temp, "1",)
	gen.AddExpression1(temp1, base.Value)

	gen.AddLabel(Linicio)

	gen.AddIf(temp, exponente.Value, "==", Lfinal)
	gen.AddExpression(temp1, temp1, base.Value, "*")
	gen.AddExpression(temp, temp, "1", "+")

	gen.AddGoto(Linicio)
	gen.AddLabel(Lfinal)

	

	return interfaces.Value{Value: temp1, IsTemp: true, Type: p.Type, TrueLabel: "", FalseLabel: ""}
}