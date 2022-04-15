package instruction

import (
	"OLC2/Compilador/ast"
	"OLC2/Compilador/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Println struct {
	Expression *arrayList.List
	Condicion  interfaces.Expression
	Formato    string
	Row        int
	Column     int
}

func NewPrintln(val *arrayList.List, cond interfaces.Expression, formato string, row int, column int) Println {
	exp := Println{val, cond, formato, row, column}
	return exp
}

func (p Println) Compilar(env *interfaces.Environment, tree *ast.Arbol, gen *ast.Generator) interface{} {

	// var conca string = "\n"
	var result interfaces.Value
	if p.Condicion == nil {
		var condBool bool = false

		newPos := -1
		contExpre := 0

		for i := 0; i < len(p.Formato); i++ {

			if p.Formato[i] == 123 { //  Caracter ASCII {

				newPos = devolverPos(p.Formato, i)
				if newPos == -2 {
					excep := ast.NewException("Semantico", "Formato incorrecto {}, hace falta }.", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				}
				condBool = true
				i = newPos
				if newPos == i && condBool {

					if contExpre < p.Expression.Len() {
						// auxCont := 0

						// for _, s := range p.Expression.ToArray() {

						// if contExpre == auxCont {
						result = p.Expression.GetValue(contExpre).(ListExpre).Compilar(env, tree, gen)
						// result = s.(ListExpre).Compilar(env, tree, gen)

						if result.Type == interfaces.EXCEPTION {
							return result.Value
						}

						if result.Type == interfaces.BOOLEAN {
							gen.AddComment("Printf Boolean")
							// conca += "\t" + "/************ Printf Boolean ************/\n"
							if result.IsTemp {

								newLabelEV := gen.NewLabel()
								newLabelEF := gen.NewLabel()

								gen.AddIf(result.Value, "1", "==", newLabelEV)
								// conca += "\t" + "if(" + result.Value + " " + "==" + " " + "1" + ") goto " + newLabelEV + ";\n"
								gen.AddGoto(newLabelEF)
								// conca += "\t" + "goto " + newLabelEF + ";\n"
								newLabel := gen.NewLabel()
								gen.AddBoolean(newLabelEV, newLabelEF, newLabel)

								// conca += "\t" + newLabelEV + ":\n"
								// conca += "\t" + "printf(\"%" + "c" + "\"," + "(char)116" + ");\n"
								// conca += "\t" + "printf(\"%" + "c" + "\"," + "(char)114" + ");\n"
								// conca += "\t" + "printf(\"%" + "c" + "\"," + "(char)117" + ");\n"
								// conca += "\t" + "printf(\"%" + "c" + "\"," + "(char)101" + ");\n"
								// conca += "\t" + "goto " + newLabel + ";\n"
								// conca += "\t" + newLabelEF + ":\n"
								// conca += "\t" + "printf(\"%" + "c" + "\"," + "(char)102" + ");\n"
								// conca += "\t" + "printf(\"%" + "c" + "\"," + "(char)97" + ");\n"
								// conca += "\t" + "printf(\"%" + "c" + "\"," + "(char)108" + ");\n"
								// conca += "\t" + "printf(\"%" + "c" + "\"," + "(char)115" + ");\n"
								// conca += "\t" + "printf(\"%" + "c" + "\"," + "(char)101" + ");\n"
								// conca += "\t" + "goto " + newLabel + ";\n"
								// conca += "\t" + newLabel + ":\n"

							}

						} else if result.Type == interfaces.INTEGER {
							gen.AddComment("Printf Integer")
							// conca += "\t" + "/************ Printf Integer ************/\n"
							gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
							// conca += "\t" + "printf(\"%" + "d" + "\"," + "(int)"+fmt.Sprintf("%v", result.Value) + ");\n"

						} else if result.Type == interfaces.FLOAT {
							gen.AddComment("Printf Float")
							// conca += "\t" + "/************ Printf Float ************/\n"
							gen.AddPrintf("f", "(double)"+fmt.Sprintf("%v", result.Value))
							// conca += "\t" + "printf(\"%" + "f" + "\"," + "(int)"+fmt.Sprintf("%v", result.Value) + ");\n"

						} else if result.Type == interfaces.STRING {
							gen.AddComment("Printf String")
							// conca += "\t" + "/************ Printf String ************/\n"
							if !tree.IsPrimitive {
								gen.AddPrintfString()
								tree.IsPrimitive = true
							}
							temp := gen.NewTemp()
							gen.AddExpression(temp, "P", fmt.Sprintf("%v", tree.GetPos()), "+")
							// conca += "\t" + temp + " = P + " + fmt.Sprintf("%v", tree.GetPos()) + ";\n"
							gen.AddExpression(temp, temp, "1", "+")
							// conca += "\t" + temp + " = " + temp + " + " + fmt.Sprintf("%v", tree.GetPos()) + ";\n"

							if result.IsTemp {
								gen.AddStack(temp, result.Value)
								// conca += "\t" + "stack[(int)" + temp + "] = "+ result.Value +";\n"
								gen.AddExpression("P", "P", fmt.Sprintf("%v", tree.GetPos()), "+")
								// conca += "\t" + "P = P + " + fmt.Sprintf("%v", tree.GetPos()) + ";\n"
								gen.PrintfString()
								// conca += "\t" + "printfString();\n"
							}
							temp = gen.NewTemp()
							gen.AddExpressionStack(temp, "P")
							// conca += "\t" + temp + " = stack[(int)P];\n"
							gen.AddExpression("P", "P", fmt.Sprintf("%v", tree.GetPos()), "-")
							// conca += "\t" + "P = P - " + fmt.Sprintf("%v", tree.GetPos()) + ";\n"

						} else if result.Type == interfaces.CHAR  {
							gen.AddComment("Printf Char")
							temp := gen.NewTemp()
							gen.AddExpressionHeap(temp,result.Value)
							gen.AddPrintf("c", "(char)"+temp)
						}

						contExpre++
						condBool = false
					}
				}

			}

			if !condBool && p.Formato[i] != 125 { // ASCII }

				gen.AddComment("Printf format {}")
				// conca += ("\t" + "/************ Printf format {} ************/\n")
				auxTemp := gen.NewTemp()
				gen.AddExpression(auxTemp, "H", "0", "+")
				// conca += "\t" + auxTemp + " = H + 0;\n"

				gen.AddHeap("H", fmt.Sprintf("%v", p.Formato[i]))
				// conca += "\theap[(int)H] = "+ fmt.Sprintf("%v", p.Formato[i]) +";\n"
				gen.AddExpression("H", "H", "1", "+")
				// conca += "\tH = H + 1;\n"
				gen.AddHeap("H", "-1")
				// conca += "\theap[(int)H] = -1;\n"
				gen.AddExpression("H", "H", "1", "+")
				// conca += "\tH = H + 1;\n"

				gen.AddComment("Printf {}")
				// conca += "\t/************ Printf {} ************/\n"
				if !tree.IsPrimitive {
					gen.AddPrintfString()
					tree.IsPrimitive = true
				}
				temp := gen.NewTemp()
				gen.AddExpression(temp, "P", fmt.Sprintf("%v", tree.GetPos()), "+")
				// conca += "\t" + temp + "= P + " + fmt.Sprintf("%v", tree.GetPos()) + ";\n"
				gen.AddExpression(temp, temp, "1", "+")
				// conca += "\t" + temp + " = " + temp + " + 1;\n"

				gen.AddStack(temp, auxTemp)
				// conca += "\t" + "stack[(int)" + temp + "] = "+ auxTemp +";\n"
				gen.AddExpression("P", "P", fmt.Sprintf("%v", tree.GetPos()), "+")
				// conca += "\t" + "P = P + " + fmt.Sprintf("%v", tree.GetPos()) + ";\n"
				gen.PrintfString()
				// conca += "\t" + "printfString();\n"

				temp = gen.NewTemp()
				gen.AddExpressionStack(temp, "P")
				// conca += "\t" + temp + " = stack[(int)P];\n"
				gen.AddExpression("P", "P", fmt.Sprintf("%v", tree.GetPos()), "-")
				// conca += "\t" + "P = P - " + fmt.Sprintf("%v", tree.GetPos()) + ";\n"

			}

		}

		// gen.AddConcaPrintf(conca)

		if p.Expression.Len() != contExpre {
			excep := ast.NewException("Semantico", "Formato incorrecto {}, cantidad incorrecta de expreciones, hace falta.", p.Row, p.Column)
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
			gen.AddExpressionStack(temp, "P")
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

		} else if result.Type == interfaces.CHAR  {
			gen.AddComment("Printf Char")
			temp := gen.NewTemp()
			gen.AddExpressionHeap(temp,result.Value)
			gen.AddPrintf("c", "(char)"+temp)
		} else {
			excep := ast.NewException("Semantico", "Formato incorrecto {}, Tipo de Dato Incorrecto.", p.Row, p.Column)
			tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}
	}

	gen.AddComment("Salto de Linea \\n")
	gen.AddPrintf("c", "10")

	return result.Value
}

func devolverPos(text string, pos int) int {

	var newPos int = -2
	if text[pos] == 123 { // ASCII  -> {
		pos++
	}

	for i := pos; pos < len(text); i++ {

		if text[pos] != 32 { // ASCII  -> espace

			if text[pos] == 125 { // ASCII -> }
				return pos
			} else {
				return newPos
			}
		}

		pos++
	}

	return newPos
}
