package instruction

import (
	"fmt"
	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/ast"
)

type Println struct {
	Expression 	interfaces.Expression
	Row 		int
	Column 		int
}

func NewPrintln(val interfaces.Expression, row int, column int) Println {
	exp := Println{val, row, column}
	return exp
}

func (p Println) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interface{} {

	var result interfaces.Value

	result = p.Expression.Compilar(env, tree, gen)

	if result.Type == interfaces.BOOLEAN {

		newLabel := gen.NewLabel()
		gen.AddLabel(result.TrueLabel)
		gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", 1))
		gen.AddGoto(newLabel)
		gen.AddLabel(result.FalseLabel)
		gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", 0))
		gen.AddLabel(newLabel)

	} else if result.Type == interfaces.INTEGER {
		gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))

	} else if result.Type == interfaces.FLOAT {
		gen.AddPrintf("f", "(double)"+fmt.Sprintf("%v", result.Value))

	}else if result.Type == interfaces.STRING {
		if !tree.IsPrimitive {
			gen.AddPrintfString()
			tree.IsPrimitive = true
		}
		temp := gen.NewTemp()
		gen.AddExpression(temp, "P", fmt.Sprintf("%v", tree.GetPos()), "+")
		gen.AddExpression(temp, temp, "1", "+")
		
		if result.IsTemp {
			gen.AddStack(temp, result.Value)
			gen.PrintfString()
		}
		temp = gen.NewTemp()
		gen.AddExpressionStack(temp,"P")
		gen.AddExpression("P", "P", fmt.Sprintf("%v", tree.GetPos()), "-")
		
	} else {
		gen.AddPrintf("c", "(char)"+fmt.Sprintf("%v", result.Value))
	}


	gen.AddPrintf("c","10")

	return result.Value
}
