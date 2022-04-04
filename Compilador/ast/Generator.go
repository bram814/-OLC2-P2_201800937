package ast

import (
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

type Generator struct {
	temporal int
	label    int
	code     *arrayList.List
	tempList *arrayList.List
}

func NewGenerator() *Generator {
	generator := Generator{temporal: 0, label: 0, code: arrayList.New(), tempList: arrayList.New()}
	return &generator
}

/************************************************* [CODE] *************************************************/   
func (g Generator) GetCode() *arrayList.List {
	return g.code
}

/************************************************* [TEMPORALES] *************************************************/                        
func (g Generator) GetTemp() *arrayList.List {
	return g.tempList
}

//Generar un nuevo temporal
func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.temporal)
	g.temporal = g.temporal + 1
	g.tempList.Add(temp)
	return temp
}

/************************************************* [LABEL] *************************************************/          
func (g *Generator) NewLabel() string {
	temp := g.label
	g.label = g.label + 1
	return "L" + fmt.Sprintf("%v", temp)
}

//Añade label al codigo
func (g *Generator) AddLabel(label string) {
	g.code.Add(label + ":")
}

/************************************************* [IF] *************************************************/   
func (g *Generator) AddIf(left string, right string, operator string, label string) {
	g.code.Add("if(" + left + " " + operator + " " + right + ") goto " + label + ";")
}

/************************************************* [GOTO] *************************************************/   
func (g *Generator) AddGoto(label string) {
	g.code.Add("goto " + label + ";")
}

/************************************************* [EXPRESSION] *************************************************/   
func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	g.code.Add(target + " = " + left + " " + operator + " " + right + ";")
}

/************************************************* [PRINTF] *************************************************/   
func (g *Generator) AddPrintf(typePrint string, value string) {
	g.code.Add("printf(\"%" + typePrint + "\"," + value + ");")
}
