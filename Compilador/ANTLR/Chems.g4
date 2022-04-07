parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}


@header {
    import "OLC2/Compilador/interfaces"
    import "OLC2/Compilador/expression"
    import "OLC2/Compilador/instruction"
    import "OLC2/Compilador/instruction/variable"
    import arrayList "github.com/colegno/arraylist"
}


start returns [*arrayList.List lista]
  : instrucciones {$lista = $instrucciones.l}
;

instrucciones returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : list +=instruccion+   {
        listInt := localctx.(*InstruccionesContext).GetList()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

end_instr returns [int v]
  : TK_PUNTOCOMA            { $v = 1}
  |                         { $v = 0}
;


instruccion returns [interfaces.Instruction instr]
  : instr_println end_instr       { $instr = $instr_println.instr }
  | instr_main                    { $instr = $instr_main.instr }
  | instr_declaracion             { $instr = $instr_declaracion.instr }
  | instr_asignacion              { $instr = $instr_asignacion.instr }

;


/******************************** [PRINTLN!] ********************************/
instr_println returns [interfaces.Instruction instr]
  : R_PRINTLN TK_PARA expressions TK_PARC TK_PUNTOCOMA                           { $instr = instruction.NewPrintln($expressions.p, $R_PRINTLN.line, localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn()) }
  | R_PRINTLN TK_PARA STRING TK_COMA expressions TK_PARC TK_PUNTOCOMA            { $instr = instruction.NewPrintln($expressions.p, $R_PRINTLN.line, localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn()) }
;  

/******************************** [MAIN] ********************************/
instr_main returns [interfaces.Instruction instr]
  : R_FUNCTION R_MAIN TK_PARA TK_PARC TK_LLAVEA instrucciones TK_LLAVEC          { $instr = instruction.NewMain($instrucciones.l, $R_MAIN.line, localctx.(*Instr_mainContext).Get_R_MAIN().GetColumn()) }
;


/******************************** [DECLARACION][VARIABLE] ********************************/
instr_declaracion returns [interfaces.Instruction instr]
  : R_LET R_MUT ID TK_IGUAL expressions TK_PUNTOCOMA                           { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $expressions.p, true,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET R_MUT ID TK_DOSPUNTOS instr_tipo TK_PUNTOCOMA                        { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp,  nil,           true,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET R_MUT ID TK_DOSPUNTOS instr_tipo TK_IGUAL expressions TK_PUNTOCOMA   { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $expressions.p, true,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET ID TK_IGUAL expressions TK_PUNTOCOMA                                 { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $expressions.p, false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET ID TK_DOSPUNTOS instr_tipo TK_PUNTOCOMA                              { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, nil,            false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET ID TK_DOSPUNTOS instr_tipo TK_IGUAL expressions TK_PUNTOCOMA         { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $expressions.p, false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  
;

/******************************** [ASIGNACION][VARIABLE] ********************************/
instr_asignacion returns [interfaces.Instruction instr]
  : ID TK_IGUAL expressions TK_PUNTOCOMA                                       { $instr = variable.NewAssignment($ID.text, $expressions.p, $ID.line, localctx.(*Instr_asignacionContext).Get_ID().GetColumn()) }
;

/******************************** [TIPO] ********************************/
instr_tipo returns [interfaces.TypeExpression tipo_exp]
  : R_INT       {$tipo_exp = interfaces.INTEGER}
  | R_FLOAT     {$tipo_exp = interfaces.FLOAT}
  | R_STRING    {$tipo_exp = interfaces.STRING}
  | R_STR       {$tipo_exp = interfaces.STR}
  | R_BOOL      {$tipo_exp = interfaces.BOOLEAN}
;

/******************************** [EXPRESIONES] ********************************/
expressions returns [interfaces.Expression p]
  : expre_log                                   { $p = $expre_log.p }
  | expre_rel                                   { $p = $expre_rel.p } 
  | expre_arit                                  { $p = $expre_arit.p }

;

expre_log returns [interfaces.Expression p]
  : op='!' left=expre_log                                               { $p = expression.NewOperacion($left.p, $op.text, nil,      true,  $op.line, localctx.(*Expre_logContext).GetOp().GetColumn()) }
  | left=expre_log op=('&&'|'||') right=expre_log                       { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_logContext).GetOp().GetColumn()) }
  | expre_rel                                                           { $p = $expre_rel.p } 
;

expre_rel returns [interfaces.Expression p]
  : left=expre_rel op=('<'|'<='|'>='|'>'|'!='|'==') right=expre_rel     { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_relContext).GetOp().GetColumn()) }
  | expre_arit                                                          { $p = $expre_arit.p }
;

expre_arit returns [interfaces.Expression p]
  : op='-' left=expre_arit                                              { $p = expression.NewOperacion($left.p, $op.text, nil,      true,  $op.line, localctx.(*Expre_aritContext).GetOp().GetColumn()) }
  | left=expre_arit op=('*'|'/'|'%') right=expre_arit                   { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_aritContext).GetOp().GetColumn()) }
  | left=expre_arit op=('+'|'-') right=expre_arit                       { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_aritContext).GetOp().GetColumn()) }
  | expre_valor                                                         { $p = $expre_valor.p }
  | TK_PARA expressions TK_PARC                                         { $p = $expressions.p }
;


expre_valor returns [interfaces.Expression p]
  : primitivo                                                             { $p = $primitivo.p }
  
;

primitivo returns[interfaces.Expression p]
    :NUMBER {
              num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.NULL, $NUMBER.line, localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())
       } 
    |DOUBLE  {  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.NULL, $DOUBLE.line, localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())
              }
    |STRING { 
              str:= $STRING.text[1:len($STRING.text)-1]
              $p = expression.NewPrimitivo(str, interfaces.STRING, interfaces.NULL, $STRING.line, localctx.(*PrimitivoContext).Get_STRING().GetColumn())
            
            }
            
    |BOOLEAN { 
              // str:= $BOOLEAN.text[1:len($BOOLEAN.text)-1]
              exp,_ := strconv.ParseBool($BOOLEAN.text)
              $p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.NULL, $BOOLEAN.line, localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())
            }

    |CHAR {

            str:= $CHAR.text[1]
            $p = expression.NewPrimitivo(string(str), interfaces.CHAR, interfaces.NULL, $CHAR.line, localctx.(*PrimitivoContext).Get_CHAR().GetColumn())
          
          }
    |NUMBER R_AS R_INT{
              num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.INTEGER, $NUMBER.line, localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())
       }
    |NUMBER R_AS R_FLOAT{
              num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.FLOAT, $NUMBER.line, localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())
       }  
    |DOUBLE  R_AS R_INT{  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.INTEGER, $DOUBLE.line, localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())
              }
    |DOUBLE  R_AS R_FLOAT{  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.FLOAT, $DOUBLE.line, localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())
              }
    |ID       { $p = variable.NewIdentifier($ID.text, $ID.line, localctx.(*PrimitivoContext).Get_ID().GetColumn()) }
;