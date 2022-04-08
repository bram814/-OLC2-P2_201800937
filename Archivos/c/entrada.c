#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3;

	
void main() {
	P = 0; H = 0;

	/************ PRINTF ************/
	/************ Relacional == ************/
	if(11 == 11) goto L0;
	goto L1;
	L0:
	t0 = 1 + 0;
	goto L2;
	L1:
	t0 = 0 + 0;
	goto L2;
	L2:
	/************ PRIMITIVO BOOLEAN ************/
	goto L3;
	goto L4;
	L3:
	t1 = 1 + 0;
	goto L5;
	L4:
	t1 = 0 + 0;
	goto L5;
	L5:
	/************ Logico || ************/
	if(t0 == 1) goto L6;
	goto L7;
	L7:
	if(t1 == 1) goto L8;
	goto L9;
	L6:
	t2 = 1 + 0;
	goto L10;
	L8:
	t2 = 1 + 0;
	goto L10;
	L9:
	t2 = 0 + 0;
	goto L10;
	L10:
	if(t2 == 1) goto L11;
	goto L12;
	L11:
	t3 = 0 + 0;
	goto L13;
	L12:
	t3 = 1 + 0;
	goto L13;
	L13:
	/************ Printf Boolean ************/
	if(t3 == 1) goto L14;
	goto L15;
	L14:
	printf("%c",(char)116);
	printf("%c",(char)114);
	printf("%c",(char)117);
	printf("%c",(char)101);
	goto L16;
	L15:
	printf("%c",(char)102);
	printf("%c",(char)97);
	printf("%c",(char)108);
	printf("%c",(char)115);
	printf("%c",(char)101);
	goto L16;
	L16:
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return;
}
