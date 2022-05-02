package structs

import (
	"fmt"
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type Declaration struct {
	Id        	string
	Parameters  *arrayList.List
	IsMut  		bool
	IdStruct    string
	Row       	int
	Column    	int
}

func NewDeclaration(id string, parameters *arrayList.List, isMut bool, idStruct string, row int, column int) Declaration {
	instr := Declaration{id, parameters, isMut, idStruct, row, column}
	return instr
}

func (p Declaration) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	/* Buscar si el id ya existe */
	symbol := env.SearchSymbol(p.Id)

	if symbol.Type != interfaces.NULL {

		excep := interfaces.NewException("Semantico", "Ya Existe ese Id "+p.Id + " (Struct).", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}



	syStruct := env.GetStructs(p.IdStruct)

	if syStruct.Type == interfaces.NULL {

		excep := interfaces.NewException("Semantico", "No Existe ese Struct - "+p.Id + " (Struct- Declaration).", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

	}

	var result interfaces.Value
	var value interfaces.Symbol
	STRUCT := syStruct.Value.(interfaces.Symbol).Value.(interfaces.SymbolStructs)


	if STRUCT.Parameters.Len() == p.Parameters.Len() {

		for s := 0; s < p.Parameters.Len(); s++ {

			if p.Parameters.GetValue(s).(ListExpreDecla).Id != STRUCT.Parameters.GetValue(s).(ListExpre).Id {
				excep := interfaces.NewException("Semantico", "Error en Declaracion Struct, Id incorrecto " + p.Parameters.GetValue(s).(ListExpreDecla).Id + ".", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}
		}

		value.IsMut = p.IsMut
		value.Type = interfaces.STRUCT
		value.Id = p.Id

		gen.AddComment("Struct")
		
		temp := gen.NewTemp()
		gen.AddExpression(temp,"P",fmt.Sprintf("%v", env.Posicion),"+")
		value.Posicion = env.Posicion
		env.NewPos()
		gen.AddStack(temp, "H")
		gen.AddExpression("H","H",fmt.Sprintf("%v", p.Parameters.Len()),"+")


		gen.AddComment("Declaracion Struct")
		var saveAtributos []interface{}

		for s := 0; s <  p.Parameters.Len(); s++ {

			gen.AddComment("Atributo - "+ p.Parameters.GetValue(s).(ListExpreDecla).Id)
			
			result = p.Parameters.GetValue(s).(ListExpreDecla).Expression.Compilar(env, tree, gen)
			if result.Type == interfaces.EXCEPTION {
				return result
			}

			gen.AddComment("Guardando")
			temp   = gen.NewTemp()
			temp1 := gen.NewTemp()
			gen.AddExpression(temp,"P",fmt.Sprintf("%v", value.Posicion), "+")
			gen.AddExpressionStack(temp1, temp)
			temp = gen.NewTemp()
			gen.AddExpression(temp, temp1, fmt.Sprintf("%v", s), "+")
			gen.AddHeap(temp,result.Value)

			valueStructs := interfaces.ValueStructs {
				Id  	 :	p.Parameters.GetValue(s).(ListExpreDecla).Id,
				Type 	 :	result.Type,
				Posicion : 	s,
			}

			saveAtributos = append(saveAtributos, valueStructs)
			
		}

		// fmt.Println("----")
		value.Value = saveAtributos
		// fmt.Println(value)
		// fmt.Println(saveAtributos)

		env.AddSymbolStruct(value.Id, value, value.Type, value.IsMut, value.Posicion, env)



	} else {
		excep := interfaces.NewException("Semantico", "Error en Declaracion de Struct, cantidad de parametros incorrectos..", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	
	}


	
	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}