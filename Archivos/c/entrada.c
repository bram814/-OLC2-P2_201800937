#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34;

	/************ Fucniones ************/
	/************ Main ************/
	
int main() {
	P = 0; H = 0;

	/************ Array ************/
	t0 = H + 0;
	t1 = P + 0;
	stack[(int)t1] = t0;
	heap[(int)t0] = 10;
	H = H + 11;
	/************ Guardando ************/
	t2 = P + 0;
	t3 = stack[(int)t2];
	t4 = t3 + 0;
	heap[(int)t4] = 5;
	/************ Guardando ************/
	t5 = P + 0;
	t6 = stack[(int)t5];
	t7 = t6 + 1;
	heap[(int)t7] = 20;
	/************ Guardando ************/
	t8 = P + 0;
	t9 = stack[(int)t8];
	t10 = t9 + 2;
	heap[(int)t10] = 8;
	/************ Guardando ************/
	t11 = P + 0;
	t12 = stack[(int)t11];
	t13 = t12 + 3;
	heap[(int)t13] = 17;
	/************ Guardando ************/
	t14 = P + 0;
	t15 = stack[(int)t14];
	t16 = t15 + 4;
	heap[(int)t16] = 65;
	/************ Guardando ************/
	t17 = P + 0;
	t18 = stack[(int)t17];
	t19 = t18 + 5;
	heap[(int)t19] = 2;
	/************ Guardando ************/
	t20 = P + 0;
	t21 = stack[(int)t20];
	t22 = t21 + 6;
	heap[(int)t22] = 40;
	/************ Guardando ************/
	t23 = P + 0;
	t24 = stack[(int)t23];
	t25 = t24 + 7;
	heap[(int)t25] = 4;
	/************ Guardando ************/
	t26 = P + 0;
	t27 = stack[(int)t26];
	t28 = t27 + 8;
	heap[(int)t28] = 35;
	/************ Guardando ************/
	t29 = P + 0;
	t30 = stack[(int)t29];
	t31 = t30 + 9;
	heap[(int)t31] = 90;
	/************ Acceso Array ************/
	t32 = P + 0;
	t33 = stack[(int)t32];
	t32 = t33 + 5;
	t34 = heap[(int)t32];
	/************ Printf Integer ************/
	printf("%d",(int)t34);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return 0;
}
