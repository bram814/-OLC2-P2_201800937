package main

import (
	"fmt"
	// "reflect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	/* COMPILADOR */
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"OLC2/Compilador/ANTLR/parser"

	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/environment"
	"OLC2/Compilador/ast"

)

var CODE_OUT_ string = ""
var tablaSimboloP []interface{}

func main() {

	

	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"CODE_INPUT": "",
			"CODE_OUT": "",
			"Tabla_Error" : nil,
		})
	})

	app.Post("/compilar", Execute)

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
		"CODE_INPUT"  : data.Input,
		"CODE_OUT"	  : CODE_OUT_,
		"Tabla_Error" : tablaSimboloP,
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

	/* SALIDA */
	var _salida string
	_salida = ""
	CODE_OUT_ ="";
	/* TREE */
	var tree *ast.Arbol
	tree = ast.NewArbol()
	/* ENVIRONMENT */
	var globalEnv environment.Environment
	globalEnv = environment.NewEnvironment(nil)

	/* GENERATOR */
	var gen *ast.Generator
	gen = ast.NewGenerator()

	
	// var contMain int = 0


	for _, s := range result.ToArray() {
		s.(interfaces.Instruction).Compilar(globalEnv, tree, gen)
	}

	_salida += "#include <stdio.h>\n"
	_salida += "#include <math.h>\n"
	_salida += "double HEAP[82000];\n"
	_salida += "double STACK[82000];\n"
	_salida += "double P;\n"
	_salida += "double H;\n"
	_salida += "double "

	_salida += fmt.Sprintf("%v", gen.GetTemp().GetValue(0))
	gen.GetTemp().RemoveAtIndex(0)

	for _, s := range gen.GetTemp().ToArray() {
		_salida += ", "
		_salida += fmt.Sprintf("%v", s)
	}

	_salida += ";\n\n"
	_salida += "\nvoid main(){\n"


	for _, s := range gen.GetCode().ToArray() {
		_salida += fmt.Sprintf("%v", s)
		_salida += "\n"
	}

	_salida += "\nreturn;\n}\n"



	CODE_OUT_ = _salida



}

