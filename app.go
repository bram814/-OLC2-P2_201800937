package main

import (

	"fmt"
	"reflect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	/* COMPILADOR */
	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/instruction"
	"OLC2/Compilador/instruction/function"

	/* ANTLR */
	"OLC2/Compilador/ANTLR/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"OLC2/Optimizar/Grammar"

)

var CODE_OUT_ string = ""
var CODE_HEAD string = ""
var TSGlobalError []interface{}
var TSGlobalSymbol []interface{}
var TSGlobalDB []interface{}

func main() {

	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	text := ""
	text += "fn main() {\n"
    text +=	"\tprintln!(\"hello World!\");\n}"

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"CODE_INPUT":  text,
			"CODE_OUT":    "",
			"Tabla_Error": nil,
		})
	})

	app.Get("/optimizar", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("Optimizar", fiber.Map{
			"CODE_INPUT":  CODE_OUT_,
			"CODE_OUT":    "",
		})
	})

	app.Post("/compilar", Execute)
	app.Post("/ExecuteOptimizar", Optimizar)


	/*
		Se ejecuta el servidor y en caso de fallar, muetra log.Fatal con el error.
	*/
	_ = app.Listen(":3000")
}

type getInput struct {
	Input string `form:"editor"`
}

func Execute(c *fiber.Ctx) error {
	data := new(getInput)
	fmt.Println(data)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	is := antlr.NewInputStream(data.Input)

	lexer := parser.NewChemsLexer(is) // aqui van tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewChems(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	result := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(result, tree)

	return c.Render("index", fiber.Map{
		"CODE_INPUT":  data.Input,
		"CODE_OUT":    CODE_OUT_,
		"Tabla_Error": TSGlobalError,
		"Tabla_Symbol": TSGlobalSymbol,
		"Tabla_DB": TSGlobalDB,
	})
}

func Optimizar(c *fiber.Ctx) error {
	data := new(getInput)
	fmt.Println(data)

	if err := c.BodyParser(data); err != nil {
		return err
	}
	// fmt.Println(reflect.TypeOf(data.Input))
	var out, ts = Grammar.Optimizar(data.Input, CODE_HEAD)


	return c.Render("Optimizar", fiber.Map{
		"CODE_INPUT":  data.Input,
		"CODE_OUT":    out,
		"Tabla_Optimizer": ts,
		})
	
}


/* ANTLR*/

type TreeShapeListener struct {
	*parser.BaseChemsListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	result := ctx.GetLista()

	TSGlobalError = nil
	TSGlobalSymbol = nil
	TSGlobalDB = nil
	/* SALIDA */
	var _salida string
	_salida = ""
	CODE_OUT_ = ""
	/* TREE */
	var tree *interfaces.Arbol
	tree = interfaces.NewArbol()
	/* ENVIRONMENT */
	var globalEnv interfaces.Environment
	globalEnv = interfaces.NewEnvironment(nil)

	/* GENERATOR */
	var gen *interfaces.Generator
	gen = interfaces.NewGenerator()

	var contMain int = 0
	gen.AddComment("Fucniones")
	for _ , s := range result.ToArray() {
		newInstr := s.(interfaces.Instruction)
		if reflect.TypeOf(newInstr).String() != "instruction.Main" && reflect.TypeOf(newInstr).String() != "function.Function" && reflect.TypeOf(newInstr).String() != "structs.Definition" && reflect.TypeOf(newInstr).String() != "db.Definition"  { 
			excep := interfaces.NewException("Semantico","Solo puede ir Main, Func, Array y Mod.", -1, -1)
			tree.AddException(interfaces.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
			break
		}

		if reflect.TypeOf(newInstr).String() == "function.Function" {

			gen.AddFunction("void", newInstr.(function.Function).Id + "()")
			value := interfaces.Symbol {
				Id     : newInstr.(function.Function).Id,
				Type   : newInstr.(function.Function).Type,
				Value  : interfaces.SymbolFunction {
					
					Id			    : newInstr.(function.Function).Id,
					Type			: newInstr.(function.Function).Type,
					Instrucciones	: newInstr.(function.Function).Instrucciones,
					Parametro		: newInstr.(function.Function).Parametro,
					IsMut			: true,
				},
				IsMut  : true,
				Posicion		: 0,
			}
			tree.AddTableSymbol(*interfaces.NewTableSymbol(value.Id,"Function","Global", newInstr.(function.Function).Row, newInstr.(function.Function).Column, "--", "--"))
						
			globalEnv.AddFunction(newInstr.(function.Function).Id, value, newInstr.(function.Function).Type)
			s.(interfaces.Instruction).Compilar(&globalEnv, tree, gen)
			gen.AddFunctionEnd(false)
			globalEnv.UpdatePos(0, 0, true, &globalEnv)

		} else if reflect.TypeOf(newInstr).String() == "structs.Definition" || reflect.TypeOf(newInstr).String() == "db.Definition" {
			s.(interfaces.Instruction).Compilar(&globalEnv, tree, gen)
		}
	}


	globalEnv.UpdatePos(0, 0, true, &globalEnv)
	gen.AddComment("Main")
	gen.AddFunction("int", "main()")
	gen.StachHeap()
	for _, s := range result.ToArray() {
		newInstr := s.(interfaces.Instruction)

		if reflect.TypeOf(newInstr).String() == "instruction.Main" {
			if contMain > 0 {
				excep := interfaces.NewException("Semantico","Existen dos funciones Main.", newInstr.(instruction.Main).Row, newInstr.(instruction.Main).Column)
				tree.AddException(interfaces.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				break
			}

			s.(interfaces.Instruction).Compilar(&globalEnv, tree, gen)
			contMain++
		}
		
		if reflect.TypeOf(s).String() == "transferencia.Break" 	 { 
			excep := interfaces.NewException("Semantico","Sentencia Break fuera de Ciclo.", -1, -1)
			tree.AddException(interfaces.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
			break
		}
		if reflect.TypeOf(s).String() == "transferencia.Continue" { 
			excep := interfaces.NewException("Semantico","Sentencia Continue fuera de Ciclo.", -1, -1)
			tree.AddException(interfaces.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
			break
		}
		if reflect.TypeOf(s).String() == "transferencia.Return"   { 
			excep := interfaces.NewException("Semantico","Sentencia Return fuera de Ciclo.", -1, -1)
			tree.AddException(interfaces.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
			break
		}

		


	}

	gen.AddFunctionEnd(true)

	_salida += "#include <stdio.h>\n"
	_salida += "#include <math.h>\n"
	_salida += "double heap[23111998];\n"
	_salida += "double stack[23111998];\n"
	_salida += "double P;\n"
	_salida += "double H;\n"

	if gen.GetTemp().Len() != 0 {
		_salida += "double "

		_salida += fmt.Sprintf("%v", gen.GetTemp().GetValue(0))
		gen.GetTemp().RemoveAtIndex(0)

		for _, s := range gen.GetTemp().ToArray() {
			_salida += ", "
			_salida += fmt.Sprintf("%v", s)
		}

		_salida += ";\n\n"

	}

	for _, s := range gen.GetNative().ToArray() {
		_salida += "\t" + fmt.Sprintf("%v", s)
		_salida += "\n"
	}
	CODE_HEAD = _salida 
	for _, s := range gen.GetCode().ToArray() {
		_salida += "\t" + fmt.Sprintf("%v", s)
		_salida += "\n"
	}


	fmt.Println("----------")
	var OutException string
	OutException = ""

	for _, s := range tree.GetException().ToArray() {
		OutException += fmt.Sprintf("%v", s)
		m := make(map[string]string)
		m["Id"] = fmt.Sprintf("%v", s.(interfaces.Exception).Tipo)
		m["Descripcion"] = fmt.Sprintf("%v", s.(interfaces.Exception).Descripcion)
		m["Row"] = fmt.Sprintf("%v", s.(interfaces.Exception).Row)
		m["Column"] = fmt.Sprintf("%v", s.(interfaces.Exception).Column)
		m["Time"] = fmt.Sprintf("%v", s.(interfaces.Exception).Time)

		TSGlobalError = append(TSGlobalError, m)
	}

	fmt.Println(OutException)
	fmt.Println("----------")
	OutException = ""

	for _, s := range tree.GetTableSymbol().ToArray() {
		OutException += fmt.Sprintf("%v", s)
		m := make(map[string]string)
		m["Name"] 		= fmt.Sprintf("%v", s.(interfaces.TableSymbol).Name)
		m["Tipo"] 		= fmt.Sprintf("%v", s.(interfaces.TableSymbol).Tipo)
		m["Ambito"] 	= fmt.Sprintf("%v", s.(interfaces.TableSymbol).Ambito)
		m["Row"] 		= fmt.Sprintf("%v", s.(interfaces.TableSymbol).Row)
		m["Column"] 	= fmt.Sprintf("%v", s.(interfaces.TableSymbol).Column)
		m["Size"] 		= fmt.Sprintf("%v", s.(interfaces.TableSymbol).Size)
		m["Posicion"] 	= fmt.Sprintf("%v", s.(interfaces.TableSymbol).Posicion)

		TSGlobalSymbol = append(TSGlobalSymbol, m)
	}

	fmt.Println(OutException)

	fmt.Println("----------")

	auxCont := 1
	for _, s := range tree.GetTableDB().ToArray() {

		m := make(map[string]string)
		m["Id"]			= fmt.Sprintf("%v", auxCont)
		m["Name"] 		= fmt.Sprintf("%v", s.(interfaces.TableDB).Name)
		m["Row"] 		= fmt.Sprintf("%v", s.(interfaces.TableDB).Row)
		m["Column"] 	= fmt.Sprintf("%v", s.(interfaces.TableDB).Column)

		TSGlobalDB = append(TSGlobalDB, m)
		auxCont++
	}

	CODE_OUT_ = _salida

}
