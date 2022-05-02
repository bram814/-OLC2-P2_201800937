package structs

import (
	"fmt"
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type Assignment struct {
	Id        	string
	Atributo    *arrayList.List
	Expresion 	interfaces.Expression
	Row       	int
	Column    	int
}

func NewAssignment(id string, atri *arrayList.List, expre interfaces.Expression, row int, column int) Assignment {
	instr := Assignment{id, atri, expre, row, column}
	return instr
}

func (p Assignment) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {


	/* Buscar si el id ya existe */
	symbol := env.GetSymbol(p.Id)

	if symbol.Type == interfaces.NULL {
		excep := interfaces.NewException("Semantico", "No Existe ese Id " + p.Id + " (Struct - Assignment).", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	var result interfaces.Value
	gen.AddComment("Asinacion - Struct")
	result = p.Expresion.Compilar(env, tree, gen)
	fmt.Println()

	for s := 0; s < p.Atributo.Len(); s++ {
		
		value := SearchAttributo(symbol.Value.(interfaces.Symbol).Value.([]interface {}),p.Atributo.GetValue(s).(ListIdentifier).Id)
		if value.Type == interfaces.EXCEPTION {
			excep := interfaces.NewException("Semantico", "No Existe ese Atributo " + value.Id + " en Struct (Assingment).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}
		if symbol.IsMut {
			if value.Type != interfaces.STRUCT  {

				if value.Type == result.Type {
					temp0 := gen.NewTemp()
					temp1 := gen.NewTemp()
					gen.AddComment("SetAtributo - Struct")
					gen.AddExpression(temp0, "P", fmt.Sprintf("%v", symbol.Posicion), "+")
					gen.AddExpressionStack(temp1, temp0)
					gen.AddExpression(temp0,temp1,fmt.Sprintf("%v", value.Posicion),"+")
					gen.AddHeap(temp0, result.Value)
				} else {
					excep := interfaces.NewException("Semantico", "No se puede asignar a "+p.Id+", diferente tipo de dato en Struct (Assingment).", p.Row, p.Column)
					tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				}

			} 

		} else {
			excep := interfaces.NewException("Semantico", "No se puede asignar a "+p.Id+", no es mutable en Struct (Assingment).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}
		
	}

	// if symbol.IsMut {
	// 	gen.AddComment("Asignacion - Struct")

	// 	if symbol.Type == result.Type || symbol.Type == interfaces.NULL{
	// 		symbol.IsMut = true
	// 		env.SetSymbol(p.Id, result, true, symbol.Posicion)
	// 		temp := gen.NewTemp()
	// 		gen.AddExpression(temp, "P", fmt.Sprintf("%v", symbol.Posicion), "+")
	// 		gen.AddStack(temp, result.Value)

	// 	} else {
	// 		excep := interfaces.NewException("Semantico", "No se puede asignar, tipo de datos diferentes (Struct).", p.Row, p.Column)
	// 		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
	// 		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	
	// 	}

	//

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}

}
