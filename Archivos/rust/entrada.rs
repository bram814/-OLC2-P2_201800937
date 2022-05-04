#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13;

	/************ Fucniones ************/
	/************ Main ************/
	
int main() {
	P = 0; H = 0;

	/************ Struct ************/
	t0 = P + 0;
	stack[(int)t0] = H;
	H = H + 3;
	/************ Declaracion Struct ************/
	/************ Atributo - nombre ************/
	/************ PRIMITIVO STRING ************/
	t1 = H + 0;
	heap[(int)H] = 70;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Guardando ************/
	t2 = P + 0;
	t3 = stack[(int)t2];
	t4 = t3 + 0;
	heap[(int)t4] = t1;
	/************ Atributo - edad ************/
	/************ Guardando ************/
	t5 = P + 0;
	t6 = stack[(int)t5];
	t7 = t6 + 1;
	heap[(int)t7] = 18;
	/************ Atributo - descripcion ************/
	/************ PRIMITIVO STRING ************/
	t8 = H + 0;
	heap[(int)H] = 78;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 104;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 110;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Guardando ************/
	t9 = P + 0;
	t10 = stack[(int)t9];
	t11 = t10 + 2;
	heap[(int)t11] = t8;
	/************ Asinacion - Struct ************/
	/************ SetAtributo - Struct ************/
	t12 = P + 0;
	t13 = stack[(int)t12];
	t12 = t13 + 1;
	heap[(int)t12] = 10;
	
	return 0;
}
