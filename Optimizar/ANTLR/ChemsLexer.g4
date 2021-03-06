// CalcLexer.g4
lexer grammar ChemsLexer;

// Reserve words
RINT:       'int';
RDOUBLE:    'double';
RCHAR:      'char';
RVOID:      'void';
RINCLUDE:   '#include';
RSTDIO:     '<stdio.h>';
RMATH:      '<math.h>';
RHEAP:      'heap';
RSTACK:     'stack';
RIF:        'if';
RGOTO:      'goto';
RRETURN:    'return';
RPRINTF:    'printf';
RPOW:       'pow';
RMOD:       'mod';
RPSTACK:    'P';
RPHEAP:     'H';

// Comments
MULTILINE:      '/*' (~[/])+ '*/' -> skip;
INLINE:         '//'(~[\n])+ -> skip;

// Regular expressions
INTEGER:                [0-9]+;
DOUBLE:                 [0-9]+'.'[0-9]+;
CHAR:                   '\''~["]'\'';
ID:                     ([a-zA-Z_])[a-zA-Z0-9_]*;
STRING:                 '"'~["]*'"';

// General Operators
EQUAL:                  '=';
SEMICOLON:              ';';
COMMA:                  ',';
COLON:                  ':';
ADMIRATION:             '!';

// Grouping Operators
LEFT_PAR:               '(';
RIGHT_PAR:              ')';
LEFT_KEY:               '{';
RIGHT_KEY:              '}';
LEFT_BRACKET:           '[';
RIGHT_BRACKET:          ']';

// Relational Operators
EQUEAL_EQUAL:           '==';
NOTEQUAL:               '!=';
GREATER_THAN:           '>';
LESS_THAN:              '<';
GREATER_EQUALTHAN:      '>=';
LESS_EQUEALTHAN:        '<=';

// Aritmethic Operators
MUL:                    '*';
DIV:                    '/';
MOD:                    '%';
ADD:                    '+';
SUB:                    '-';

// White spaces
WHITESPACE: [ \r\n\t]+ -> skip;

// Escape characters
fragment
ESC_SEQ
    : '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ');