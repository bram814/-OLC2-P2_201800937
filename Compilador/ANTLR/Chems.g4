parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}


@header {
    import "OLC2/Compilador/interfaces"
    import "OLC2/Compilador/expression"
    import "OLC2/Compilador/instruction"
    import "OLC2/Compilador/instruction/variable"
    import "OLC2/Compilador/instruction/control"
    import "OLC2/Compilador/instruction/ternario"
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
  | instr_if                      { $instr = $instr_if.instr }
  | instr_match                   { $instr = $instr_match.instr}

;


/******************************** [PRINTLN!] ********************************/
instr_println returns [interfaces.Instruction instr]
  : R_PRINTLN TK_PARA primitivo TK_PARC TK_PUNTOCOMA                                 { $instr = instruction.NewPrintln(nil, $primitivo.p,       "-1",         $R_PRINTLN.line, localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn()) }
  | R_PRINTLN TK_PARA STRING TK_COMA list_expression TK_PARC TK_PUNTOCOMA            { $instr = instruction.NewPrintln($list_expression.l, nil, $STRING.text[1:len($STRING.text)-1], $R_PRINTLN.line, localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn()) }
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

/******************************** [CONTROL][IF] ********************************/
instr_if returns [interfaces.Instruction instr]
  : R_IF expressions TK_LLAVEA left_instr=instrucciones TK_LLAVEC                                                       { $instr = control.NewIf($expressions.p, $left_instr.l, nil, nil,               $R_IF.line, localctx.(*Instr_ifContext).Get_R_IF().GetColumn()) }
  | R_IF expressions TK_LLAVEA left_instr=instrucciones TK_LLAVEC R_ELSE TK_LLAVEA right_intr=instrucciones TK_LLAVEC   { $instr = control.NewIf($expressions.p, $left_instr.l, $right_intr.l, nil,     $R_IF.line, localctx.(*Instr_ifContext).Get_R_IF().GetColumn()) }
  | R_IF expressions TK_LLAVEA left_instr=instrucciones TK_LLAVEC R_ELSE instr_else_if                                  { $instr = control.NewIf($expressions.p, $left_instr.l, nil, $instr_else_if.l,  $R_IF.line, localctx.(*Instr_ifContext).Get_R_IF().GetColumn()) }
;

instr_else_if returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e +=instr_if*  {
        listInt := localctx.(*Instr_else_ifContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/******************************** [TERNARIO][IF] ********************************/
instr_ternario returns [interfaces.Expression p]
  : R_IF cond=expressions TK_LLAVEA left_instr=expressions TK_LLAVEC                                                      { $p = ternario.NewIf($cond.p, $left_instr.p, nil, nil,                       $R_IF.line, localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn()) }
  | R_IF cond=expressions TK_LLAVEA left_instr=expressions TK_LLAVEC R_ELSE TK_LLAVEA right_intr=expressions TK_LLAVEC    { $p = ternario.NewIf($cond.p, $left_instr.p, $right_intr.p, nil,             $R_IF.line, localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn()) }
  | R_IF cond=expressions TK_LLAVEA left_instr=expressions TK_LLAVEC R_ELSE instr_else_if_ternario                        { $p = ternario.NewIf($cond.p, $left_instr.p, nil, $instr_else_if_ternario.l, $R_IF.line, localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn()) }
;

instr_else_if_ternario returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e +=instr_ternario*  {
        listInt := localctx.(*Instr_else_if_ternarioContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetP())
        }
    }
;

/******************************** [CONTROL][MATCH] ********************************/
instr_match returns [interfaces.Instruction instr]
  : R_MATCH expressions TK_LLAVEA list_case block_default TK_LLAVEC        { $instr = control.NewMatch($expressions.p, $list_case.l, $block_default.l, $R_MATCH.line, localctx.(*Instr_matchContext).Get_R_MATCH().GetColumn()) }
  | R_MATCH expressions TK_LLAVEA block_default TK_LLAVEC                  { $instr = control.NewMatch($expressions.p, nil, $block_default.l,          $R_MATCH.line, localctx.(*Instr_matchContext).Get_R_MATCH().GetColumn()) }
;

list_case returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instr_case+  {
        listInt := localctx.(*List_caseContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/*  CASE  */
instr_case returns [interfaces.Expression instr]
  : list_expre_case TK_IGUALMAYOR TK_LLAVEA instrucciones TK_LLAVEC     { $instr = control.NewCase(nil, $list_expre_case.l, $instrucciones.l) }
  | list_expre_case TK_IGUALMAYOR block_instr_match TK_COMA             { $instr = control.NewCase(nil, $list_expre_case.l, $block_instr_match.l) }
;

/* List Expression Case */
list_expre_case returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_case+  {
        listInt := localctx.(*List_expre_caseContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

block_case returns [interfaces.Expression instr]
  : expressions TK_BARRA                                                         { $instr =  control.NewCase($expressions.p, nil, nil)}
  | expressions                                                                  { $instr =  control.NewCase($expressions.p, nil, nil)}
;

block_instr_match returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : list += instruccion   {
        listInt := localctx.(*Block_instr_matchContext).GetList()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/*  DEFAULT  */
instr_default returns [interfaces.Instruction instr]
  : ID TK_IGUALMAYOR TK_LLAVEA instrucciones TK_LLAVEC           { $instr = control.NewDefault($instrucciones.l) }
  | ID TK_IGUALMAYOR block_instr_match TK_COMA                   { $instr = control.NewDefault($block_instr_match.l) }
;

block_default returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instr_default+  {
        listInt := localctx.(*Block_defaultContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/******************************** [TIPO] ********************************/
instr_tipo returns [interfaces.TypeExpression tipo_exp]
  : R_INT       {$tipo_exp = interfaces.INTEGER}
  | R_FLOAT     {$tipo_exp = interfaces.FLOAT}
  | R_STRING    {$tipo_exp = interfaces.STRING}
  | R_STR       {$tipo_exp = interfaces.STR}
  | R_BOOL      {$tipo_exp = interfaces.BOOLEAN}
  | R_CHAR      {$tipo_exp = interfaces.CHAR}
;

/* List Expression Case */
list_expression returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_list_expression+  {
        listInt := localctx.(*List_expressionContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetP())
        }
    }
;


block_list_expression returns [interfaces.Expression p]
  : expressions TK_COMA                    { $p =  instruction.NewListExpre($expressions.p) }
  | expressions                            { $p =  instruction.NewListExpre($expressions.p) }
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
    |ID       { $p = variable.NewIdentifier($ID.text, $ID.line, localctx.(*PrimitivoContext).Get_ID().GetColumn()) }

    | primitivo_casteo          { $p = $primitivo_casteo.p }
    | instr_ternario            { $p = $instr_ternario.p }
;


primitivo_casteo returns[interfaces.Expression p]
  : NUMBER R_AS R_INT{
              num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.INTEGER, $NUMBER.line, localctx.(*Primitivo_casteoContext).Get_NUMBER().GetColumn())
       }
    |NUMBER R_AS R_FLOAT{
              num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.FLOAT, $NUMBER.line, localctx.(*Primitivo_casteoContext).Get_NUMBER().GetColumn())
       }  
    |DOUBLE  R_AS R_INT{  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.INTEGER, $DOUBLE.line, localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetColumn())
              }
    |DOUBLE  R_AS R_FLOAT{  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.FLOAT, $DOUBLE.line, localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetColumn())
              }

    |BOOLEAN  R_AS R_INT { 
              // str:= $BOOLEAN.text[1:len($BOOLEAN.text)-1]
              exp,_ := strconv.ParseBool($BOOLEAN.text)
              $p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.INTEGER, $BOOLEAN.line, localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetColumn())
            }

    |BOOLEAN  R_AS R_FLOAT { 
              // str:= $BOOLEAN.text[1:len($BOOLEAN.text)-1]
              exp,_ := strconv.ParseBool($BOOLEAN.text)
              $p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.FLOAT, $BOOLEAN.line, localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetColumn())
            }
;