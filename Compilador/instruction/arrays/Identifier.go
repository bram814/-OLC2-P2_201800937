package arrays

import (
	"fmt"
	"reflect"
	"strconv"
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
	// var result interfaces.Value
	gen.AddComment("Acceso Array")

	
	symbol := env.SearchSymbol(p.Id)
	if symbol.Type == interfaces.NULL {
		symbol = env.GetSymbol(p.Id)

		if symbol.Type == interfaces.NULL {
			excep := interfaces.NewException("Semantico", "No Existe ese Id "+p.Id + "(identificador).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

		}
	}
	fmt.Println(reflect.TypeOf(symbol.Value.(interfaces.Symbol).Value.(interfaces.SymbolArrays)))

	var saveTemps []interface{}
	for _, s := range p.Atributo.ToArray() {
		result := s.(ListIdentifier).Size.Compilar(env, tree, gen)

		fmt.Println(reflect.TypeOf(result))
		if result.Type != interfaces.INTEGER {
			excep := interfaces.NewException("Semantico", "Error en Array (indetifier), Las posiciones deben de ser Enteras.", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		} 
		saveTemps = append(saveTemps, result.Value)
		

	}

	// fmt.Println(len(saveTemps))
	var size int
	for _, s := range saveTemps {
		fmt.Println(reflect.TypeOf(s))
		aux,_ := strconv.Atoi(s.(string))
		size = aux
	}



	temp := gen.NewTemp()
	temp1 := gen.NewTemp()
	gen.AddExpression(temp, "P" , fmt.Sprintf("%v", symbol.Posicion), "+")
	gen.AddExpressionStack(temp1, temp)
	gen.AddExpression(temp,temp1,fmt.Sprintf("%v",size), "+")
	temp1 = gen.NewTemp()
	gen.AddExpressionHeap(temp1,temp)

	
	return interfaces.Value{Value: temp1, IsTemp: true, Type: symbol.Value.(interfaces.Symbol).Value.(interfaces.SymbolArrays).Type, TrueLabel: "", FalseLabel: ""}
}



