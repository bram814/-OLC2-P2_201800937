#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29;

	
/************ NATIVE PRINTF STRING ************/
	
void printfString(){
	t1 = P + 1;
	t2 = stack[(int) t1];
	L1:
	t3 = heap[(int)t2];
	if(t3 == -1) goto L0;
	printf("%c", (int)t3);
	t2 = t2 + 1;
	goto L1;
	L0:
	return;
}
	
void main() {
	P = 0; H = 0;

	/************ Declaracion ************/
	stack[(int)0] = 1;
	/************ PRINTF ************/
	/************ concat format {} ************/
	t0 = H + 0;
	heap[(int)H] = 74;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t4 = P + 1;
	t4 = t4 + 1;
	stack[(int)t4] = t0;
	P = P + 1;
	printfString();
	t5 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t6 = H + 0;
	heap[(int)H] = 79;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t7 = P + 1;
	t7 = t7 + 1;
	stack[(int)t7] = t6;
	P = P + 1;
	printfString();
	t8 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t9 = H + 0;
	heap[(int)H] = 83;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t10 = P + 1;
	t10 = t10 + 1;
	stack[(int)t10] = t9;
	P = P + 1;
	printfString();
	t11 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t12 = H + 0;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t13 = P + 1;
	t13 = t13 + 1;
	stack[(int)t13] = t12;
	P = P + 1;
	printfString();
	t14 = stack[(int)P];
	P = P - 1;
	/************ PRIMITIVO STRING ************/
	t15 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t16 = P + 1;
	t16 = t16 + 1;
	stack[(int)t16] = t15;
	P = P + 1;
	printfString();
	t17 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t18 = H + 0;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t19 = P + 1;
	t19 = t19 + 1;
	stack[(int)t19] = t18;
	P = P + 1;
	printfString();
	t20 = stack[(int)P];
	P = P - 1;
	/************ Printf Integer ************/
	printf("%d",(int)1);
	/************ concat format {} ************/
	t21 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t22 = P + 1;
	t22 = t22 + 1;
	stack[(int)t22] = t21;
	P = P + 1;
	printfString();
	t23 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t24 = H + 0;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t25 = P + 1;
	t25 = t25 + 1;
	stack[(int)t25] = t24;
	P = P + 1;
	printfString();
	t26 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t27 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t28 = P + 1;
	t28 = t28 + 1;
	stack[(int)t28] = t27;
	P = P + 1;
	printfString();
	t29 = stack[(int)P];
	P = P - 1;
	
	return;
}
