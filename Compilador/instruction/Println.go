package instruction

import (
	"fmt"
	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/ast"
	arrayList "github.com/colegno/arraylist"
)

type Println struct {
	Expression 	*arrayList.List
	Condicion 	interfaces.Expression
	Row 		int
	Column 		int
}

func NewPrintln(val *arrayList.List, cond interfaces.Expression, row int, column int) Println {
	exp := Println{val, cond, row, column}
	return exp
}

func (p Println) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interface{} {
	
	var result interfaces.Value
	if p.Condicion == nil { 
		for _, s := range p.Expression.ToArray() {
			result = s.(ListExpre).Compilar(env, tree, gen)

			if result.Type == interfaces.EXCEPTION {
				return result.Value
			}

			if result.Type == interfaces.BOOLEAN {
				gen.AddComment("Printf Boolean")
				if !result.IsTemp {
					newLabel := gen.NewLabel()
					gen.AddBoolean(result.TrueLabel, result.FalseLabel, newLabel)
				}
				


			} else if result.Type == interfaces.INTEGER {
				gen.AddComment("Printf Integer")
				gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))

			} else if result.Type == interfaces.FLOAT {
				gen.AddComment("Printf Float")
				gen.AddPrintf("f", "(double)"+fmt.Sprintf("%v", result.Value))

			}else if result.Type == interfaces.STRING || result.Type == interfaces.CHAR {
				gen.AddComment("Printf String")
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
		}
	
	} else {

		result = p.Condicion.Compilar(env, tree, gen)

		if result.Type == interfaces.EXCEPTION {
			return result.Value
		}

		if result.Type == interfaces.STRING {
			gen.AddComment("Printf String")
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
			excep := ast.NewException("Semantico","Formato incorrecto {}, Tipo de Dato Incorrecto.", p.Row, p.Column)
			tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}
	}
	

		

	gen.AddComment("Salto de Linea \\n")
	gen.AddPrintf("c","10")

	return result.Value
}
