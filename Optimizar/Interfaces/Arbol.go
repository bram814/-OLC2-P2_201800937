package Interfaces

import (
	arrayList "github.com/colegno/arraylist"
)

type Arbol struct {
	Code         	*arrayList.List
	Temps  	     	*arrayList.List
	_RepOptimizar  	*arrayList.List
}

type RepOptimizar struct {
	Tipo 		string
	Regla       string
	Original    string
	Optimizada  string
	Row 		string
}

func NewArbol(a *arrayList.List) *Arbol {
	tree := Arbol{
		Code  			: a,
		Temps  			: nil,
		_RepOptimizar 	: arrayList.New(),
	}
	return &tree
}


/** TABLE SYMBOL*/
func NewRepOptimizar(tipo string, regla string, original string, optimizada  string, row string) *RepOptimizar {

	e := RepOptimizar{
		Tipo 		: tipo, 		
		Regla       : regla,       
		Original    : original,    
		Optimizada  : optimizada,  
		Row 		: row, 		
	}
	return &e
}

/* Add TableSymbol */
func (a *Arbol) AddRepOptimizar(e RepOptimizar) { 
	a._RepOptimizar.Add(e) 
}

/* Get TableSymbol */
func (a Arbol) GetRepOptimizar() *arrayList.List {
	return a._RepOptimizar
}


func NewTree(listCode *arrayList.List, temps *arrayList.List) Arbol {
	tree := Arbol{
		Code : listCode,
		Temps: temps,
	}
	return tree
}


/* Add text to Code */
func (a *Arbol) AddCode(input string) {
	a.Code.Add(input)
}

/* Get text to Code */
func (a Arbol) GetCode() *arrayList.List {
	return a.Code
}


/* Add text to Temps */
func (a *Arbol) AddTemps(input string) {
	a.Code.Add(input)
}

/* Get text to Temps */
func (a Arbol) GetTemps() *arrayList.List {
	return a.Code
}

