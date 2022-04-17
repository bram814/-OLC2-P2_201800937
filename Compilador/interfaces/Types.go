package interfaces

type TypeExpression int

const (
	INTEGER   TypeExpression = iota // 0
	FLOAT                           // 1
	STRING                          // 2
	CHAR                            // 3
	BOOLEAN                         // 4
	ARRAY                           // 5
	NULL                            // 6
	EXCEPTION                       // 7
	STR                             // 8
	BREAK  							// 9
	CONTINUE  						// 10
	RETURN  						// 11 
)
