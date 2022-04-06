package ast

import (
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

type Generator struct {
	temporal int
	label    int
	heap 	 int
	stack  	 int
	code     *arrayList.List
	native   *arrayList.List
	tempList *arrayList.List
}

func NewGenerator() *Generator {
	generator := Generator{
		temporal : 	0, 
		label 	 : 	0, 
		heap 	 : 	0, 
		stack 	 : 	0, 
		code 	 : 	arrayList.New(), 
		native 	 :  arrayList.New(), 
		tempList : 	arrayList.New()}
	return &generator
}

/************************************************* [CODE] *************************************************/   
func (g Generator) GetCode() *arrayList.List {
	return g.code
}

/************************************************* [NATIVE] *************************************************/   
func (g Generator) GetNative() *arrayList.List {
	return g.native
}

/************************************************* [TEMPORALES] *************************************************/                        
func (g Generator) GetTemp() *arrayList.List {
	return g.tempList
}

/************************************************* [TEMP] *************************************************/  
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


/************************************************* [ADD][LABEL] *************************************************/  
func (g *Generator) AddLabel(label string) {
	g.code.Add(label + ":")
}

/************************************************* [ADD][IF] *************************************************/   
func (g *Generator) AddIf(left string, right string, operator string, label string) {
	g.code.Add("if(" + left + " " + operator + " " + right + ") goto " + label + ";")
}

/************************************************* [ADD][GOTO] *************************************************/   
func (g *Generator) AddGoto(label string) {
	g.code.Add("goto " + label + ";")
}

/************************************************* [ADD][EXPRESSION][TEMP] *************************************************/   
func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	g.code.Add(target + " = " + left + " " + operator + " " + right + ";")
}

/************************************************* [ADD][HEAP] *************************************************/   
func (g *Generator) AddStack(temp string, result string) {
	g.code.Add("stack[(int)" + temp + "] = "+ result +";")
}
/************************************************* [ADD][EXPRESSION][STACK] *************************************************/   
func (g *Generator) AddExpressionStack(target string, pos string) {
	g.code.Add(target + " = stack[(int)" + pos + "];")
}

/************************************************* [ADD][HEAP] *************************************************/   
func (g *Generator) AddHeap(temp string, ascii string) {
	g.code.Add("heap[(int)" + temp + "] = "+ ascii +";")
}
/************************************************* [ADD][HEAP][EXPRESSION] *************************************************/   
func (g *Generator) AddExpressionHeap(target string, right string) {
	g.code.Add(target + " = heap[(int)" + right + "];")
}


/************************************************* [ADD][PRINTF] *************************************************/   
func (g *Generator) AddPrintf(typePrint string, value string) {
	g.code.Add("printf(\"%" + typePrint + "\"," + value + ");")
}


/************************************************* [NATIVE][STRING] *************************************************/
 func (g *Generator) AddPrintfString() {
 	
 	temp 	:= g.NewTemp()
 	auxTemp := g.NewTemp()
 	EV := g.NewLabel()
 	EF :=  g.NewLabel()

 	g.native.Add("\rvoid printfString(){")
 	g.native.Add(temp + " = P + 1;")
 	g.native.Add(auxTemp + " = stack[(int) " + temp + "];")
 	g.native.Add(EF + ":")

 	temp = g.NewTemp()

 	g.native.Add(temp + " = heap[(int)" + auxTemp + "];")
 	g.native.Add("if(" + temp + " == -1) goto " + EV + ";")
 	g.native.Add("printf(\"%c\"," + " (int)"+fmt.Sprintf("%v", temp) + ");")

 	g.native.Add(auxTemp + " = " + auxTemp + " + 1;")
 	g.native.Add("goto " + EF + ";")
 	g.native.Add(EV + ":")
 	g.native.Add("return;\n}")

}

func (g *Generator) PrintfString() {
	g.code.Add("printfString();")
}


/************************************************* [BOOLEAN] *************************************************/
 func (g *Generator) AddBoolean(EV string, EF string, newLabel string) {

 	g.AddLabel(EV)
 	g.AddPrintf("c", "(char)116")
 	g.AddPrintf("c", "(char)114")
 	g.AddPrintf("c", "(char)117")
 	g.AddPrintf("c", "(char)101")
 	g.AddGoto(newLabel)
 	g.AddLabel(EF)
 	g.AddPrintf("c", "(char)102")
 	g.AddPrintf("c", "(char)97")
 	g.AddPrintf("c", "(char)108")
 	g.AddPrintf("c", "(char)115")
 	g.AddPrintf("c", "(char)101")
 	g.AddGoto(newLabel)
 	g.AddLabel(newLabel)

 }

  func (g *Generator) Boolean(EV string, EF string, newLabel string, temp string) {

 	g.AddLabel(EV)
 	g.AddExpression(temp,"1","0","+")
 	g.AddGoto(newLabel)
 	g.AddLabel(EF)
 	g.AddExpression(temp,"0","0","+")
 	g.AddGoto(newLabel)
 	g.AddLabel(newLabel)

 }


/************************************************* [FUNCIONES] *************************************************/
func (g *Generator) StachHeap() {
 	g.code.Add("P = 0; H = 0;\n")

 }
 func (g *Generator) AddFunction(tipo string, nombre string) {
 	g.code.Add("\r" + tipo + " " + nombre + " {")

 }

func (g *Generator) AddFunctionEnd() {
 	g.code.Add("\n\treturn;\r}")

 }



