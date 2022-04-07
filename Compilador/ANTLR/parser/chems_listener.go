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

	// EnterInstr_declaracion is called when entering the instr_declaracion production.
	EnterInstr_declaracion(c *Instr_declaracionContext)

	// EnterInstr_asignacion is called when entering the instr_asignacion production.
	EnterInstr_asignacion(c *Instr_asignacionContext)

	// EnterInstr_tipo is called when entering the instr_tipo production.
	EnterInstr_tipo(c *Instr_tipoContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterExpre_log is called when entering the expre_log production.
	EnterExpre_log(c *Expre_logContext)

	// EnterExpre_rel is called when entering the expre_rel production.
	EnterExpre_rel(c *Expre_relContext)

	// EnterExpre_arit is called when entering the expre_arit production.
	EnterExpre_arit(c *Expre_aritContext)

	// EnterExpre_valor is called when entering the expre_valor production.
	EnterExpre_valor(c *Expre_valorContext)

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

	// ExitInstr_declaracion is called when exiting the instr_declaracion production.
	ExitInstr_declaracion(c *Instr_declaracionContext)

	// ExitInstr_asignacion is called when exiting the instr_asignacion production.
	ExitInstr_asignacion(c *Instr_asignacionContext)

	// ExitInstr_tipo is called when exiting the instr_tipo production.
	ExitInstr_tipo(c *Instr_tipoContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitExpre_log is called when exiting the expre_log production.
	ExitExpre_log(c *Expre_logContext)

	// ExitExpre_rel is called when exiting the expre_rel production.
	ExitExpre_rel(c *Expre_relContext)

	// ExitExpre_arit is called when exiting the expre_arit production.
	ExitExpre_arit(c *Expre_aritContext)

	// ExitExpre_valor is called when exiting the expre_valor production.
	ExitExpre_valor(c *Expre_valorContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)
}
