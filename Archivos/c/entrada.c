#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16;

	/************ Fucniones ************/
	/************ Main ************/
	
int main() {
	P = 0; H = 0;

	/************ Array ************/
	t0 = H + 0;
	t1 = P + 0;
	stack[(int)t1] = t0;
	heap[(int)t0] = 4;
	H = H + 4;
	/************ Guardando ************/
	t2 = P + 0;
	t3 = stack[(int)t2];
	t4 = t3 + 0;
	heap[(int)t4] = 8;
	/************ Guardando ************/
	t5 = P + 0;
	t6 = stack[(int)t5];
	t7 = t6 + 1;
	heap[(int)t7] = 4;
	/************ Guardando ************/
	t8 = P + 0;
	t9 = stack[(int)t8];
	t10 = t9 + 2;
	heap[(int)t10] = 6;
	/************ Guardando ************/
	t11 = P + 0;
	t12 = stack[(int)t11];
	t13 = t12 + 3;
	heap[(int)t13] = 2;
	/************ Acceso Array ************/
	t14 = P + 0;
	t15 = stack[(int)t14];
	t14 = t15 + 1;
	t16 = heap[(int)t14];
	/************ Printf Integer ************/
	printf("%d",(int)t16);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return 0;
}
