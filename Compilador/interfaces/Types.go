package interfaces

type TypeExpression int

const (
	INTEGER TypeExpression = iota 	// 0
	FLOAT							// 1
	STRING							// 2
	BOOLEAN							// 3
	ARRAY							// 4
	NULL							// 5
	EXCEPTION 						// 6
)
