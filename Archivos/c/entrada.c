#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9;

	
/************ NATIVE COMPARE STRING ************/
	
void compareString(){
	t2 = P + 1;
	t3 = stack[(int)t2];
	t2 = t2 + 1;
	t4 = stack[(int)t2];
	L3:
	t5 = heap[(int)t3];
	t6 = heap[(int)t4];
	if (t5 != t6) goto L5;
	if (t5 == -1) goto L4;
	t3 = t3 + 1;
	t4 = t4 + 1;
	goto L3;
	L4:
	stack[(int)P] = 1;
	goto L2;
	L5:
	stack[(int)P] = 0;
	L2:
	return;
}
	
void main() {
	P = 0; H = 0;

	/************ PRIMITIVO STRING ************/
	t0 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ PRIMITIVO STRING ************/
	t1 = H + 0;
	heap[(int)H] = 98;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Relacional != ************/
	t7 = P + 0;
	t7 = t7 + 1;
	stack[(int)t7] = t0;
	t7 = t7 + 1;
	stack[(int)t7] = t1;
	P = P + 0;
	compareString();
	t8 = stack[(int)P];
	P = P - 0;
	if(t8 == 0) goto L0;
	goto L1;
	L0:
	t9 = 1 + 0;
	goto L6;
	L1:
	t9 = 0 + 0;
	goto L6;
	L6:
	/************ Printf Boolean ************/
	if(t9 == 1) goto L7;
	goto L8;
	L7:
	printf("%c",(char)116);
	printf("%c",(char)114);
	printf("%c",(char)117);
	printf("%c",(char)101);
	goto L9;
	L8:
	printf("%c",(char)102);
	printf("%c",(char)97);
	printf("%c",(char)108);
	printf("%c",(char)115);
	printf("%c",(char)101);
	goto L9;
	L9:
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return;
}
