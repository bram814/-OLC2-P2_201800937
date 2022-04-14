package variable

import (
	"OLC2/Compilador/ast"
	"OLC2/Compilador/interfaces"
	"fmt"
	// "reflect"
)

type Declaration struct {
	Id        string
	Type      interfaces.TypeExpression
	Expresion interfaces.Expression
	IsMut     bool
	Row       int
	Column    int
}

func NewDeclaration(id string, tipo interfaces.TypeExpression, val interfaces.Expression, isMut bool, row int, column int) Declaration {
	instr := Declaration{id, tipo, val, isMut, row, column}
	return instr
}

func (p Declaration) Compilar(env *interfaces.Environment, tree *ast.Arbol, gen *ast.Generator) interface{} {

	/* Buscar si el id ya existe */
	symbol := env.SearchSymbol(p.Id)

	if symbol.Type != interfaces.NULL {
		fmt.Println("No puede agregar")
		excep := ast.NewException("Semantico", "Ya Existe ese Id "+p.Id, p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	// symbol = env.GetSymbol(p.Id)

	// if symbol.Type != interfaces.NULL {
	// 	excep := ast.NewException("Semantico", "Ya Existe ese Id "+p.Id, p.Row, p.Column)
	// 	tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
	// 	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	// }

	var result interfaces.Value

	if p.Expresion != nil {
		result = p.Expresion.Compilar(env, tree, gen)
		p.Type = result.Type
	} else {

		result.Type = p.Type
	}

	if result.Type == p.Type || p.Type == interfaces.NULL {
		gen.AddComment("Declaracion")

		if env.IsAmbit() {
			// fmt.Println("Local")
			temp := gen.NewTemp()
			gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")
			gen.AddStack(temp, result.Value)
			env.AddSymbol(p.Id, result, result.Type, p.IsMut, env.Posicion)
			env.NewPos()

		} else {
			// fmt.Println("Global")
			env.AddSymbol(p.Id, result, result.Type, p.IsMut, tree.GetPos())
			gen.AddStack(fmt.Sprintf("%v", tree.GetPos()), result.Value)
			tree.NewPos()
		}

	} else {
		excep := ast.NewException("Semantico", "Tipo Incorrecto en Declaracion.", p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	return result.Value

}
