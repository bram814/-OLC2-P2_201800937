package Grammar

import (
	"fmt"
	"reflect"
	/* Optimizar */
	"OLC2/Optimizar/Interfaces"
   	"OLC2/Optimizar/ANTLR/parser"
   	"OLC2/Optimizar/Function"
	"github.com/antlr/antlr4/runtime/Go/antlr"


)
var CODE_OUT string
var CODE_HEAD string
var TSGlobal []interface{}

func Optimizar(datos string, head string) (string,[]interface {}) {

	CODE_HEAD = head
	
	is := antlr.NewInputStream(datos)

	lexer := parser.NewChemsLexer(is) // aqui van tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewChems(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	result := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(result, tree)
	//fmt.Println(CODE_OUT)
	return CODE_OUT, TSGlobal
}

/* ANTLR - OPTIMIZAR */

type TreeShapeListener struct {
	*parser.BaseChemsListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}



func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	ast := ctx.GetAST()
	TSGlobal = nil
	salida_ := ""

	if CODE_HEAD == "" {
		salida_+= "#include <stdio.h>\n"
		salida_ += "#include <math.h>\n"
		salida_ += "double heap[23111998];\n"
		salida_ += "double stack[23111998];\n"
		salida_ += "double P;\n"
		salida_ += "double H;\n"

		salida_ += "double "
		for i := 0; i < ast.Temps.Len(); i++ {
			if i < ast.Temps.Len()-1 {
				salida_ += ast.Temps.GetValue(i).(string) + ", "
			} else {
				salida_ += ast.Temps.GetValue(i).(string)
			}
		}
	salida_ += ";\n\n"
	} else {
		salida_ = CODE_HEAD + "\n"
	}
	
	


	

	salida_ += "void main() {\n"
	if ast.Code != nil {
		for i := 0; i < ast.Code.Len(); i++ {
			r := ast.Code.GetValue(i)

			if r != nil {
				salida_ += r.(Interfaces.Instruction).Execute().(string)
			}

			if reflect.TypeOf(ast.Code.GetValue(i)).String() == "Function.Function" {



				for _, s := range ast.Code.GetValue(i).(Function.Function).Tree.GetRepOptimizar().ToArray() {

					m := make(map[string]string)

					m["Tipo"] = fmt.Sprintf("%v", s.(Interfaces.RepOptimizar).Tipo)
					m["Regla"] = fmt.Sprintf("%v", s.(Interfaces.RepOptimizar).Regla)
					m["Original"] = fmt.Sprintf("%v", s.(Interfaces.RepOptimizar).Original)
					m["Optimizada"] = fmt.Sprintf("%v", s.(Interfaces.RepOptimizar).Optimizada)
					m["Row"] = fmt.Sprintf("%v", s.(Interfaces.RepOptimizar).Row)

					TSGlobal = append(TSGlobal, m)
				}
			}
		}
	}
	salida_ += "}"
	CODE_OUT = salida_

}
