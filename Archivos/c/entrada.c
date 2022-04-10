#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21;

	
/************ NATIVE PRINTF STRING ************/
	
void printfString(){
	t2 = P + 1;
	t3 = stack[(int) t2];
	L7:
	t4 = heap[(int)t3];
	if(t4 == -1) goto L6;
	printf("%c", (int)t4);
	t3 = t3 + 1;
	goto L7;
	L6:
	return;
}
	
void main() {
	P = 0; H = 0;

	/************ Relacional == ************/
	if(1 == 1) goto L0;
	goto L1;
	L0:
	t0 = 1 + 0;
	goto L2;
	L1:
	t0 = 0 + 0;
	goto L2;
	L2:
	/************ Control - If ************/
	if(t0 == 1) goto L3;
	goto L4;
	L3:
	/************ PRIMITIVO STRING ************/
	t1 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t5 = P + 0;
	t5 = t5 + 1;
	stack[(int)t5] = t1;
	printfString();
	t6 = stack[(int)P];
	P = P - 0;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L5;
	L4:
	/************ Control - If (else if) ************/
	/************ Relacional != ************/
	if(2 != 2) goto L8;
	goto L9;
	L8:
	t7 = 1 + 0;
	goto L10;
	L9:
	t7 = 0 + 0;
	goto L10;
	L10:
	/************ Control - If ************/
	if(t7 == 1) goto L11;
	goto L12;
	L11:
	/************ PRIMITIVO STRING ************/
	t8 = H + 0;
	heap[(int)H] = 98;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t9 = P + 0;
	t9 = t9 + 1;
	stack[(int)t9] = t8;
	printfString();
	t10 = stack[(int)P];
	P = P - 0;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L13;
	L12:
	/************ Control - If (else if) ************/
	/************ Relacional != ************/
	if(3 != 3) goto L14;
	goto L15;
	L14:
	t11 = 1 + 0;
	goto L16;
	L15:
	t11 = 0 + 0;
	goto L16;
	L16:
	/************ Control - If ************/
	if(t11 == 1) goto L17;
	goto L18;
	L17:
	/************ PRIMITIVO STRING ************/
	t12 = H + 0;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t13 = P + 0;
	t13 = t13 + 1;
	stack[(int)t13] = t12;
	printfString();
	t14 = stack[(int)P];
	P = P - 0;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L19;
	L18:
	/************ Control - If (else if) ************/
	/************ Relacional != ************/
	if(4 != 4) goto L20;
	goto L21;
	L20:
	t15 = 1 + 0;
	goto L22;
	L21:
	t15 = 0 + 0;
	goto L22;
	L22:
	/************ Control - If ************/
	if(t15 == 1) goto L23;
	goto L24;
	L23:
	/************ PRIMITIVO STRING ************/
	t16 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t17 = P + 0;
	t17 = t17 + 1;
	stack[(int)t17] = t16;
	printfString();
	t18 = stack[(int)P];
	P = P - 0;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L25;
	L24:
	/************ Control - If (else) ************/
	/************ PRIMITIVO STRING ************/
	t19 = H + 0;
	heap[(int)H] = 102;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t20 = P + 0;
	t20 = t20 + 1;
	stack[(int)t20] = t19;
	printfString();
	t21 = stack[(int)P];
	P = P - 0;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L25;
	L25:
	goto L19;
	L19:
	goto L13;
	L13:
	goto L5;
	L5:
	
	return;
}
