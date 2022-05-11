package db

import (

	"OLC2/Compilador/interfaces"
)

type Definition struct {
	Id       	 string
	Row  		 int
	Column   	 int
}

func NewDefinition(Id string, Row int, Column int) Definition {
	instr := Definition{Id, Row, Column}
	return instr
}

func (p Definition) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	tree.AddTableDB(*interfaces.NewTableDB(p.Id, p.Row, p.Column ))
		
	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}

}
	
