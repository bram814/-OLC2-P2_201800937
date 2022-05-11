// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChemsListener is a complete listener for a parse tree produced by Chems.
type BaseChemsListener struct{}

var _ ChemsListener = &BaseChemsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChemsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChemsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChemsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChemsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseChemsListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseChemsListener) ExitStart(ctx *StartContext) {}

// EnterHeaders is called when production headers is entered.
func (s *BaseChemsListener) EnterHeaders(ctx *HeadersContext) {}

// ExitHeaders is called when production headers is exited.
func (s *BaseChemsListener) ExitHeaders(ctx *HeadersContext) {}

// EnterList_temps is called when production list_temps is entered.
func (s *BaseChemsListener) EnterList_temps(ctx *List_tempsContext) {}

// ExitList_temps is called when production list_temps is exited.
func (s *BaseChemsListener) ExitList_temps(ctx *List_tempsContext) {}

// EnterTemps is called when production temps is entered.
func (s *BaseChemsListener) EnterTemps(ctx *TempsContext) {}

// ExitTemps is called when production temps is exited.
func (s *BaseChemsListener) ExitTemps(ctx *TempsContext) {}

// EnterListFuncs is called when production listFuncs is entered.
func (s *BaseChemsListener) EnterListFuncs(ctx *ListFuncsContext) {}

// ExitListFuncs is called when production listFuncs is exited.
func (s *BaseChemsListener) ExitListFuncs(ctx *ListFuncsContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseChemsListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseChemsListener) ExitFunction(ctx *FunctionContext) {}

// EnterInstructions is called when production instructions is entered.
func (s *BaseChemsListener) EnterInstructions(ctx *InstructionsContext) {}

// ExitInstructions is called when production instructions is exited.
func (s *BaseChemsListener) ExitInstructions(ctx *InstructionsContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseChemsListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseChemsListener) ExitInstruction(ctx *InstructionContext) {}

// EnterAssign_stack is called when production assign_stack is entered.
func (s *BaseChemsListener) EnterAssign_stack(ctx *Assign_stackContext) {}

// ExitAssign_stack is called when production assign_stack is exited.
func (s *BaseChemsListener) ExitAssign_stack(ctx *Assign_stackContext) {}

// EnterAssign_heap is called when production assign_heap is entered.
func (s *BaseChemsListener) EnterAssign_heap(ctx *Assign_heapContext) {}

// ExitAssign_heap is called when production assign_heap is exited.
func (s *BaseChemsListener) ExitAssign_heap(ctx *Assign_heapContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseChemsListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseChemsListener) ExitAssign(ctx *AssignContext) {}

// EnterIf_instr is called when production if_instr is entered.
func (s *BaseChemsListener) EnterIf_instr(ctx *If_instrContext) {}

// ExitIf_instr is called when production if_instr is exited.
func (s *BaseChemsListener) ExitIf_instr(ctx *If_instrContext) {}

// EnterGoto_instr is called when production goto_instr is entered.
func (s *BaseChemsListener) EnterGoto_instr(ctx *Goto_instrContext) {}

// ExitGoto_instr is called when production goto_instr is exited.
func (s *BaseChemsListener) ExitGoto_instr(ctx *Goto_instrContext) {}

// EnterLabel_instr is called when production label_instr is entered.
func (s *BaseChemsListener) EnterLabel_instr(ctx *Label_instrContext) {}

// ExitLabel_instr is called when production label_instr is exited.
func (s *BaseChemsListener) ExitLabel_instr(ctx *Label_instrContext) {}

// EnterPrintf_instr is called when production printf_instr is entered.
func (s *BaseChemsListener) EnterPrintf_instr(ctx *Printf_instrContext) {}

// ExitPrintf_instr is called when production printf_instr is exited.
func (s *BaseChemsListener) ExitPrintf_instr(ctx *Printf_instrContext) {}

// EnterReturn_instr is called when production return_instr is entered.
func (s *BaseChemsListener) EnterReturn_instr(ctx *Return_instrContext) {}

// ExitReturn_instr is called when production return_instr is exited.
func (s *BaseChemsListener) ExitReturn_instr(ctx *Return_instrContext) {}

// EnterGetHeap_instr is called when production getHeap_instr is entered.
func (s *BaseChemsListener) EnterGetHeap_instr(ctx *GetHeap_instrContext) {}

// ExitGetHeap_instr is called when production getHeap_instr is exited.
func (s *BaseChemsListener) ExitGetHeap_instr(ctx *GetHeap_instrContext) {}

// EnterGetStack_instr is called when production getStack_instr is entered.
func (s *BaseChemsListener) EnterGetStack_instr(ctx *GetStack_instrContext) {}

// ExitGetStack_instr is called when production getStack_instr is exited.
func (s *BaseChemsListener) ExitGetStack_instr(ctx *GetStack_instrContext) {}

// EnterCall_instr is called when production call_instr is entered.
func (s *BaseChemsListener) EnterCall_instr(ctx *Call_instrContext) {}

// ExitCall_instr is called when production call_instr is exited.
func (s *BaseChemsListener) ExitCall_instr(ctx *Call_instrContext) {}

// EnterExpr_print is called when production expr_print is entered.
func (s *BaseChemsListener) EnterExpr_print(ctx *Expr_printContext) {}

// ExitExpr_print is called when production expr_print is exited.
func (s *BaseChemsListener) ExitExpr_print(ctx *Expr_printContext) {}

// EnterConvert is called when production convert is entered.
func (s *BaseChemsListener) EnterConvert(ctx *ConvertContext) {}

// ExitConvert is called when production convert is exited.
func (s *BaseChemsListener) ExitConvert(ctx *ConvertContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseChemsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseChemsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpr_rel is called when production expr_rel is entered.
func (s *BaseChemsListener) EnterExpr_rel(ctx *Expr_relContext) {}

// ExitExpr_rel is called when production expr_rel is exited.
func (s *BaseChemsListener) ExitExpr_rel(ctx *Expr_relContext) {}

// EnterExpr_arit is called when production expr_arit is entered.
func (s *BaseChemsListener) EnterExpr_arit(ctx *Expr_aritContext) {}

// ExitExpr_arit is called when production expr_arit is exited.
func (s *BaseChemsListener) ExitExpr_arit(ctx *Expr_aritContext) {}

// EnterData_type is called when production data_type is entered.
func (s *BaseChemsListener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BaseChemsListener) ExitData_type(ctx *Data_typeContext) {}

// EnterExpr_valor is called when production expr_valor is entered.
func (s *BaseChemsListener) EnterExpr_valor(ctx *Expr_valorContext) {}

// ExitExpr_valor is called when production expr_valor is exited.
func (s *BaseChemsListener) ExitExpr_valor(ctx *Expr_valorContext) {}
