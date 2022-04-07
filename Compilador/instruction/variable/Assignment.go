package variable

import (
	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/environment"
	"OLC2/Compilador/ast"
	"fmt"
)

type Assignment struct {
	Id 			string
	Expresion	interfaces.Expression
	Row			int
	Column		int
}


func NewAssignment(id string, val interfaces.Expression, row int, column int) Assignment {
	instr := Assignment{id, val, row, column}
	return instr
}


func (p Assignment) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interface{} {

	/* Buscar si el id ya existe */
	symbol := env.(environment.Environment).GetSymbol(p.Id)

	if symbol.Type == interfaces.NULL {
		excep := ast.NewException("Semantico","No Existe ese Id " + p.Id, p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}
	
	var result interfaces.Value
	result = p.Expresion.Compilar(env, tree, gen)

	if symbol.IsMut {
		gen.AddComment("Asignacion")
		symbol.IsMut = true
		env.(environment.Environment).SetSymbol(p.Id, result, true, symbol.Posicion)
		gen.AddStack(fmt.Sprintf("%v", symbol.Posicion),result.Value)

	}else {
		excep := ast.NewException("Semantico","No se puede asignar a " + p.Id + ", no es mutable.", p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}
	
	
	return result.Value

}