package expression

import (
	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/ast"
	"fmt"
	// "reflect"
)


type Primitivo struct {
	Value 	interface{}
	Type 	interfaces.TypeExpression
	Row 	int
	Column 	int
}

func NewPrimitivo(val interface{}, tipo interfaces.TypeExpression, row int, column int) Primitivo {
	exp := Primitivo{val, tipo, row, column}
	return exp
}

func (p Primitivo) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interfaces.Value {


	if p.Type == interfaces.STRING || p.Type == interfaces.CHAR {

		temp := gen.NewTemp()
		gen.AddExpression(temp,"H","0","+")

		for i := 0; i < len(p.Value.(string)); i++ {
			gen.AddHeap("H",fmt.Sprintf("%v", p.Value.(string)[i]))
			gen.AddExpression("H","H","1","+")
		}

		gen.AddHeap("H","-1")
		gen.AddExpression("H","H","1","+")

		return interfaces.Value{
			Value:      temp,
			IsTemp:     true,
			Type:       p.Type,
			TrueLabel:  "",
			FalseLabel: "",
		}



	} else if p.Type == interfaces.BOOLEAN {
		EV := gen.NewLabel()
		EF := gen.NewLabel()

		if p.Value.(bool) {
			gen.AddGoto(EV)
			gen.AddGoto(EF)
				
		}else {
			gen.AddGoto(EF)
			gen.AddGoto(EV)
			
		}

		return interfaces.Value{
			Value:      "",
			IsTemp:     false,
			Type:       p.Type,
			TrueLabel:  EV,
			FalseLabel: EF,
		}
	}


	return interfaces.Value{
		Value:      fmt.Sprintf("%v", p.Value),
		IsTemp:     false,
		Type:       p.Type,
		TrueLabel:  "",
		FalseLabel: "",
	}
}

