package variable

import (
	"OLC2/Compilador/ast"
	"OLC2/Compilador/interfaces"
	"fmt"
)

type Identifier struct {
	Id     string
	Row    int
	Column int
}

func NewIdentifier(id string, row int, column int) Identifier {
	instr := Identifier{id, row, column}
	return instr
}

func (p Identifier) Compilar(env *interfaces.Environment, tree *ast.Arbol, gen *ast.Generator) interfaces.Value {

	symbol := env.GetSymbol(p.Id)

	if symbol.Type == interfaces.NULL {
		excep := ast.NewException("Semantico", "No Existe ese Id "+p.Id, p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

	}

	gen.AddComment("Identificador")
	fmt.Println(symbol.Value.(interfaces.Value).Type)
	temp := gen.NewTemp()
	gen.AddExpressionStack(temp, fmt.Sprintf("%v", symbol.Posicion))

	return interfaces.Value{
		Value:      temp,
		IsTemp:     true,
		Type:       symbol.Value.(interfaces.Value).Type,
		TrueLabel:  "",
		FalseLabel: "",
	}

}
