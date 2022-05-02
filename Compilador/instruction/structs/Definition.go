package structs

import (
	// "fmt"
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type Definition struct {
	Id        	string
	Type      	interfaces.TypeExpression
	Parameters  *arrayList.List
	Row       	int
	Column    	int
}

func NewDefinition(id string, tipo interfaces.TypeExpression, parameters *arrayList.List, row int, column int) Definition {
	instr := Definition{id, tipo, parameters, row, column}
	return instr
}

func (p Definition) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	// var result interfaces.Value

	symbol := env.GetStructs(p.Id)

	if symbol.Type != interfaces.NULL {
		excep := interfaces.NewException("Semantico", "Ya Existe ese Struct -"+p.Id + ".", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			
	}

	value := interfaces.Symbol {
		Id     : p.Id,
		Type   : p.Type,
		Value  : interfaces.SymbolStructs {
			
			Id			    : p.Id,
			Type			: p.Type,
			Parameters		: p.Parameters,
			Value           : nil,
			Size			: p.Parameters.Len(),
		},
		IsMut    : false,
		Posicion : 0,
	}

	env.AddStructs(p.Id, value, p.Type)


	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}