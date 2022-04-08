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
	Formato     string
	Row 		int
	Column 		int
}

func NewPrintln(val *arrayList.List, cond interfaces.Expression, formato string, row int, column int) Println {
	exp := Println{val, cond, formato, row, column}
	return exp
}

func (p Println) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interface{} {
	
	var result interfaces.Value
	if p.Condicion == nil { 
		var condBool bool = false
		// var isConca bool = false
		newPos := -1;
		contExpre := 0;

		gen.AddComment("PRINTF")


		for i := 0; i < len(p.Formato); i++ {
			
			if p.Formato[i] == 123 { //  Caracter ASCII {
				
				
				newPos = devolverPos(p.Formato, i)
				if newPos == -2 {
					excep := ast.NewException("Semantico","Formato incorrecto {}, hace falta }.", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				} 
				condBool = true
				i = newPos
				if newPos == i && condBool{

					if contExpre < p.Expression.Len() {
						auxCont := 0

						for _, s := range p.Expression.ToArray() {

							if contExpre == auxCont {

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
									
								} else if result.Type == interfaces.STRING || result.Type == interfaces.CHAR {
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
										gen.AddExpression("P", "P", fmt.Sprintf("%v", tree.GetPos()), "+")
										gen.PrintfString()
									}
									temp = gen.NewTemp()
									gen.AddExpressionStack(temp,"P")
									gen.AddExpression("P", "P", fmt.Sprintf("%v", tree.GetPos()), "-")
									
								} 
								

							}
							auxCont++;
						}
						contExpre++;
						condBool = false
					}
				}


			}
			
			if !condBool && p.Formato[i] != 125 { // ASCII }
				
					// gen.AddHeap("H",fmt.Sprintf("%v", p.Formato[i]))
					// gen.AddExpression("H","H","1","+")
				fmt.Println(string(p.Formato[i]))
				gen.AddComment("concat format {}")
				auxTemp := gen.NewTemp()
				gen.AddExpression(auxTemp,"H","0","+")

				
				gen.AddHeap("H",fmt.Sprintf("%v", p.Formato[i]))
				gen.AddExpression("H","H","1","+")
				gen.AddHeap("H","-1")
				gen.AddExpression("H","H","1","+")

				gen.AddComment("Printf {}")
				if !tree.IsPrimitive {
					gen.AddPrintfString()
					tree.IsPrimitive = true
				}
				temp := gen.NewTemp()
				gen.AddExpression(temp, "P", fmt.Sprintf("%v", tree.GetPos()), "+")
				gen.AddExpression(temp, temp, "1", "+")
				
				
				gen.AddStack(temp, auxTemp)
				gen.AddExpression("P", "P", fmt.Sprintf("%v", tree.GetPos()), "+")
				gen.PrintfString()
				
				temp = gen.NewTemp()
				gen.AddExpressionStack(temp,"P")
				gen.AddExpression("P", "P", fmt.Sprintf("%v", tree.GetPos()), "-")


				
			}

		
			
		}

		fmt.Println("------")
		fmt.Println(p.Expression.Len())
		fmt.Println(contExpre)
		fmt.Println("------")

		if p.Expression.Len() != contExpre {
			excep := ast.NewException("Semantico","Formato incorrecto {}, cantidad incorrecta de expreciones, hace falta.", p.Row, p.Column)
			tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
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
			
		} else if result.Type == interfaces.BOOLEAN {
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
			
		}else {
			excep := ast.NewException("Semantico","Formato incorrecto {}, Tipo de Dato Incorrecto.", p.Row, p.Column)
			tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}
	}
	

		

	gen.AddComment("Salto de Linea \\n")
	gen.AddPrintf("c","10")

	return result.Value
}

func devolverPos(text string, pos int) int {

	var newPos int = -2
	if text[pos] == 123 {	// ASCII  -> {
		pos++
	}
	
	for i := pos; pos < len(text); i++ { 

		if text[pos] != 32 { // ASCII  -> espace

			if text[pos] == 125 { // ASCII -> }
				return pos
			}else {
				return newPos
			}
		}
		
		pos++;
	}
	
	return newPos
}