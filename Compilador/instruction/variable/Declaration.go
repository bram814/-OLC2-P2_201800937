package variable



import (

	"OLC2/Interprete/interfaces"
	"OLC2/Interprete/environment"
	"OLC2/Interprete/ast"
	"fmt"
)

type Declaration struct {
	Id 			string
	Tipo 		interfaces.TipoExpresion
	Expresion 	interfaces.Expresion
	IsMut		bool
	Row			int
	Column		int
}

func NewDeclaration(id string, tipo interfaces.TipoExpresion, val interfaces.Expresion, isMut bool, row int, column int) Declaration {
	instr := Declaration{id, tipo, val, isMut, row, column}
	return instr
}


func (p Declaration) Interpretar(env interface{}, tree *ast.Arbol, gen *ast.Generator)  interface{} {


	var value string = ""

	/* Buscar si el id ya existe */
	symbol := env.(environment.Environment).GetSymbol(p.Id)

	if symbol.Tipo != interfaces.NULL {
		excep := ast.NewException("Semantico","Ya Existe ese Id "+p.Id, p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Ya Existe ese Id "+ p.Id, Row: p.Row, Column: p.Column})
		// return interfaces.Symbol{Id: "", Tipo: interfaces.EXCEPTION, Valor: excep}
		return nil
	}

	var result interfaces.Value

	if p.Expresion != nil {
		result = p.Expresion.Compilar(env, tree, gen)
		p.Tipo = result.Tipo
	}else{
		fmt.Println("oye es declaracion vacia")
		result.Tipo = p.Tipo
	}

	
	if (result.Tipo == p.Tipo || p.Tipo == interfaces.NULL) {
		env.(environment.Environment).AddSymbol(p.Id, result, result.Tipo, p.IsMut)
	

	}else {
		excep := ast.NewException("Semantico","Tipo incorrecto en Declaracion.", p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo:"Semantico", Descripcion: "Tipo incorrecto en Declaracion.", Row: p.Row, Column: p.Column})
	}


	return result.Value

}