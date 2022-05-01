package expression

import (
	"OLC2/Compilador/interfaces"
	"fmt"
	// "reflect"
)

type Primitivo struct {
	Value  interface{}
	Type   interfaces.TypeExpression
	Casteo interfaces.TypeExpression
	Row    int
	Column int
}

func NewPrimitivo(val interface{}, tipo interfaces.TypeExpression, casteo interfaces.TypeExpression, row int, column int) Primitivo {
	exp := Primitivo{val, tipo, casteo, row, column}
	return exp
}

func (p Primitivo) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	if p.Type == interfaces.STRING || p.Type == interfaces.CHAR || p.Type == interfaces.STR {
		gen.AddComment("PRIMITIVO STRING")
		temp := gen.NewTemp()
		gen.AddExpression(temp, "H", "0", "+")

		for i := 0; i < len(p.Value.(string)); i++ {
			gen.AddHeap("H", fmt.Sprintf("%v", p.Value.(string)[i]))
			gen.AddExpression("H", "H", "1", "+")
		}

		gen.AddHeap("H", "-1")
		gen.AddExpression("H", "H", "1", "+")

		return interfaces.Value{
			Value:      temp,
			IsTemp:     true,
			Type:       p.Type,
			TrueLabel:  "",
			FalseLabel: "",
		}

	} else if p.Type == interfaces.BOOLEAN {
		gen.AddComment("PRIMITIVO BOOLEAN")

		if p.Casteo != interfaces.NULL {

			if p.Value.(bool) {

				if p.Casteo == interfaces.FLOAT {
					return interfaces.Value{
						Value:      fmt.Sprintf("%v", 1),
						IsTemp:     false,
						Type:       p.Casteo,
						TrueLabel:  "",
						FalseLabel: "",
					}
				} else if p.Casteo == interfaces.INTEGER {
					return interfaces.Value{
						Value:      fmt.Sprintf("%v", 1),
						IsTemp:     false,
						Type:       p.Casteo,
						TrueLabel:  "",
						FalseLabel: "",
					}
				}

			} else {

				if p.Casteo == interfaces.FLOAT {
					return interfaces.Value{
						Value:      fmt.Sprintf("%v", 0),
						IsTemp:     false,
						Type:       p.Casteo,
						TrueLabel:  "",
						FalseLabel: "",
					}
				} else if p.Casteo == interfaces.INTEGER {
					return interfaces.Value{
						Value:      fmt.Sprintf("%v", 0),
						IsTemp:     false,
						Type:       p.Casteo,
						TrueLabel:  "",
						FalseLabel: "",
					}
				}

			}

		} else {
			EV := gen.NewLabel()
			EF := gen.NewLabel()
			newLabel := gen.NewLabel()
			newTemp := gen.NewTemp()

			if p.Value.(bool) {
				gen.AddGoto(EV)
				gen.AddGoto(EF)

				gen.Boolean(EV, EF, newLabel, newTemp)

			} else {
				gen.AddGoto(EF)
				gen.AddGoto(EV)
				newLabel := gen.NewLabel()
				newTemp := gen.NewTemp()
				gen.Boolean(EV, EF, newLabel, newTemp)

			}

			return interfaces.Value{
				Value:      newTemp,
				IsTemp:     true,
				Type:       p.Type,
				TrueLabel:  "",
				FalseLabel: "",
			}

		}

		return interfaces.Value{
			Value:      "",
			IsTemp:     false,
			Type:       p.Type,
			TrueLabel:  "",
			FalseLabel: "",
		}
	} else if p.Type == interfaces.INTEGER || p.Type == interfaces.FLOAT {

		if p.Casteo != interfaces.NULL {

			if p.Type == interfaces.INTEGER && p.Casteo == interfaces.FLOAT {
				return interfaces.Value{
					Value:      fmt.Sprintf("%v", float64(p.Value.(int))),
					IsTemp:     false,
					Type:       p.Casteo,
					TrueLabel:  "",
					FalseLabel: "",
				}
			} else if p.Type == interfaces.FLOAT && p.Casteo == interfaces.INTEGER {
				return interfaces.Value{
					Value:      fmt.Sprintf("%v", int(p.Value.(float64))),
					IsTemp:     false,
					Type:       p.Casteo,
					TrueLabel:  "",
					FalseLabel: "",
				}
			}
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
