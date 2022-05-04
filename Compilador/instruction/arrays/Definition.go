package arrays

import (
	"fmt"
	"reflect"
	"strconv"
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type Definition struct {
	Id       	 string
	IsMut  		 bool
	Declaration  *arrayList.List
	Datos        *arrayList.List
	Row  		 int
	Column   	 int
}

func NewDefinition(Id string, IsMut bool, Decla *arrayList.List, Datos *arrayList.List, Row int, Column int) Definition {
	instr := Definition{Id, IsMut, Decla, Datos, Row, Column}
	return instr
}

func (p Definition) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {
	

	symbol := env.SearchSymbol(p.Id)

	if symbol.Type != interfaces.NULL {
		// fmt.Println("No puede agregar")
		excep := interfaces.NewException("Semantico", "Error en Declaración de Array, Ya Existe ese Id "+p.Id, p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	var saveTemps *arrayList.List = arrayList.New();
	var dimUno interfaces.Value 
	var size int
	var typeArray interfaces.TypeExpression
	for _, s := range p.Declaration.ToArray() { 
		if reflect.TypeOf(s).String() == "arrays.DimUno" {
			fmt.Println(reflect.TypeOf(s))
			fmt.Println(s.(DimUno).Type)
			typeArray = s.(DimUno).Type
			dimUno = s.(DimUno).Expression.Compilar(env, tree, gen)
			
			if dimUno.Type == interfaces.INTEGER && !dimUno.IsTemp{
				fmt.Println(dimUno.Value)
				aux,_ := strconv.Atoi(dimUno.Value)
				size = aux
				saveTemps.Add(size)
			} else {
				fmt.Println("error semantico, debe de ser un entero")
			}
		}
		
	}

	var symbolArray interfaces.Symbol
	var valueArray interfaces.SymbolArrays
	valueArray.Id = p.Id 
	valueArray.Type = typeArray
	valueArray.Tam = saveTemps
	valueArray.Size = size
	
	fmt.Println("---")
	fmt.Println(typeArray)
	fmt.Println("---")
	gen.AddComment("Array")

	temp  := gen.NewTemp()
	temp1 := gen.NewTemp()
	gen.AddExpression(temp,"H","0","+")
	gen.AddExpression(temp1,"P", fmt.Sprintf("%v",env.Posicion),"+")
	symbolArray.Posicion = env.Posicion
	env.NewPos()
	gen.AddStack(temp1,temp)
	gen.AddHeap(temp, fmt.Sprintf("%v", size))
	gen.AddExpression("H", "H", fmt.Sprintf("%v", size),"+")

	cont := 0;
	for _, s := range p.Datos.ToArray() { 

		fmt.Println(reflect.TypeOf(s))

		if reflect.TypeOf(s).String() == "arrays.DimUnoDatos" {
			
			for _, z := range s.(DimUnoDatos).Datos.ToArray() { 
				// fmt.Println(reflect.TypeOf(z))
				result := z.(ListDimUno).Expression.Compilar(env, tree, gen)
				if result.Type == typeArray {

					// fmt.Println(result.Value)

					gen.AddComment("Guardando")

					temp   = gen.NewTemp()
					temp1  = gen.NewTemp()
					gen.AddExpression(temp,"P",fmt.Sprintf("%v", symbolArray.Posicion ), "+")
					gen.AddExpressionStack(temp1, temp)
					temp = gen.NewTemp()
					gen.AddExpression(temp, temp1, fmt.Sprintf("%v", cont), "+")
					gen.AddHeap(temp,result.Value)

					cont ++;
				} else {
					fmt.Println("Error Semantico, tipo de dato diferente en Array (Dimension Uno).")
				}
				
			}
		}

	}

	
	symbolArray.Value = valueArray
	
	env.AddSymbolArrays(p.Id, symbolArray, interfaces.ARRAY, p.IsMut, symbolArray.Posicion, env)


	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}