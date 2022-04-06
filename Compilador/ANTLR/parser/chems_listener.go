// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChemsListener is a complete listener for a parse tree produced by Chems.
type ChemsListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterEnd_instr is called when entering the end_instr production.
	EnterEnd_instr(c *End_instrContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterInstr_println is called when entering the instr_println production.
	EnterInstr_println(c *Instr_printlnContext)

	// EnterInstr_main is called when entering the instr_main production.
	EnterInstr_main(c *Instr_mainContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterInstr_expre is called when entering the instr_expre production.
	EnterInstr_expre(c *Instr_expreContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitEnd_instr is called when exiting the end_instr production.
	ExitEnd_instr(c *End_instrContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitInstr_println is called when exiting the instr_println production.
	ExitInstr_println(c *Instr_printlnContext)

	// ExitInstr_main is called when exiting the instr_main production.
	ExitInstr_main(c *Instr_mainContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitInstr_expre is called when exiting the instr_expre production.
	ExitInstr_expre(c *Instr_expreContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)
}
