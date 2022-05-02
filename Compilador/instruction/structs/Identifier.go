package structs

import (
	"fmt"
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type Identifier struct {
	Id     		string
	Atributo    *arrayList.List
	Row    		int
	Column 		int
}

func NewIdentifier(id string, atr *arrayList.List, row int, column int) Identifier {
	instr := Identifier{id, atr, row, column}
	return instr
}

func (p Identifier) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	/* Buscar si el id ya existe */
	symbol := env.SearchSymbol(p.Id)

	if symbol.Type == interfaces.NULL {
		symbol = env.GetSymbol(p.Id)

		if symbol.Type == interfaces.NULL {
			excep := interfaces.NewException("Semantico", "No Existe ese Id "+p.Id + "(identificador).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

		}
	}

	

	if symbol.Type != interfaces.STRUCT {
		excep := interfaces.NewException("Semantico", "No Existe ese Id "+p.Id + "(identificador).", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

	} 

	gen.AddComment("Identificador - Struct")

	temp0 := gen.NewTemp()
	temp1 := gen.NewTemp()
	tempf := gen.NewTemp()
	gen.AddExpression(temp0, "P", fmt.Sprintf("%v", symbol.Posicion), "+")
	gen.AddExpressionStack(temp1, temp0)
	

	var result interfaces.Value

	for s := 0; s < p.Atributo.Len(); s++ {

		value := SearchAttributo(symbol.Value.(interfaces.Symbol).Value.([]interface {}), p.Atributo.GetValue(s).(ListIdentifier).Id)
		if value.Type == interfaces.EXCEPTION {
			excep := interfaces.NewException("Semantico", "No Existe ese Atributo " + value.Id + " en Struct (identificador).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}

		if value.Type != interfaces.STRUCT {
			gen.AddExpression(temp0,temp1,fmt.Sprintf("%v", value.Posicion),"+")
			gen.AddExpressionHeap(tempf, temp0)
			result.Type = value.Type
			result.Value = tempf

		}
	}
	
	return interfaces.Value{Value: result.Value, IsTemp: true, Type: result.Type, TrueLabel: "", FalseLabel: ""}
}


func SearchAttributo(auxVec []interface {}, atributo string) interfaces.ValueStructs {

	for _, s := range auxVec {
		if s.(interfaces.ValueStructs).Id == atributo {
			return s.(interfaces.ValueStructs)
		}
	}

	return interfaces.ValueStructs{Id: atributo, Type: interfaces.EXCEPTION, Posicion:0}
}