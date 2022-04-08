#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34;

	
/************ NATIVE PRINTF STRING ************/
	
void printfString(){
	t4 = P + 1;
	t5 = stack[(int) t4];
	L8:
	t6 = heap[(int)t5];
	if(t6 == -1) goto L7;
	printf("%c", (int)t6);
	t5 = t5 + 1;
	goto L8;
	L7:
	return;
}
	
void main() {
	P = 0; H = 0;

	/************ PRIMITIVO BOOLEAN ************/
	goto L1;
	goto L0;
	L0:
	t1 = 1 + 0;
	goto L3;
	L1:
	t1 = 0 + 0;
	goto L3;
	L3:
	/************ Declaracion ************/
	stack[(int)0] = t0;
	/************ PRINTF ************/
	/************ Identificador ************/
	t2 = stack[(int)0];
	/************ Printf Boolean ************/
	if(t2 == 1) goto L4;
	goto L5;
	L4:
	printf("%c",(char)116);
	printf("%c",(char)114);
	printf("%c",(char)117);
	printf("%c",(char)101);
	goto L6;
	L5:
	printf("%c",(char)102);
	printf("%c",(char)97);
	printf("%c",(char)108);
	printf("%c",(char)115);
	printf("%c",(char)101);
	goto L6;
	L6:
	/************ concat format {} ************/
	t3 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t7 = P + 1;
	t7 = t7 + 1;
	stack[(int)t7] = t3;
	P = P + 1;
	printfString();
	t8 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t9 = H + 0;
	heap[(int)H] = 45;
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
	/************ PRIMITIVO BOOLEAN ************/
	goto L9;
	goto L10;
	L9:
	t12 = 1 + 0;
	goto L11;
	L10:
	t12 = 0 + 0;
	goto L11;
	L11:
	/************ Printf Boolean ************/
	if(t12 == 1) goto L12;
	goto L13;
	L12:
	printf("%c",(char)116);
	printf("%c",(char)114);
	printf("%c",(char)117);
	printf("%c",(char)101);
	goto L14;
	L13:
	printf("%c",(char)102);
	printf("%c",(char)97);
	printf("%c",(char)108);
	printf("%c",(char)115);
	printf("%c",(char)101);
	goto L14;
	L14:
	/************ concat format {} ************/
	t13 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t14 = P + 1;
	t14 = t14 + 1;
	stack[(int)t14] = t13;
	P = P + 1;
	printfString();
	t15 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t16 = H + 0;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t17 = P + 1;
	t17 = t17 + 1;
	stack[(int)t17] = t16;
	P = P + 1;
	printfString();
	t18 = stack[(int)P];
	P = P - 1;
	/************ Identificador ************/
	t19 = stack[(int)0];
	/************ Printf Boolean ************/
	if(t19 == 1) goto L15;
	goto L16;
	L15:
	printf("%c",(char)116);
	printf("%c",(char)114);
	printf("%c",(char)117);
	printf("%c",(char)101);
	goto L17;
	L16:
	printf("%c",(char)102);
	printf("%c",(char)97);
	printf("%c",(char)108);
	printf("%c",(char)115);
	printf("%c",(char)101);
	goto L17;
	L17:
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ PRINTF ************/
	/************ Identificador ************/
	t20 = stack[(int)0];
	/************ Printf Boolean ************/
	if(t20 == 1) goto L18;
	goto L19;
	L18:
	printf("%c",(char)116);
	printf("%c",(char)114);
	printf("%c",(char)117);
	printf("%c",(char)101);
	goto L20;
	L19:
	printf("%c",(char)102);
	printf("%c",(char)97);
	printf("%c",(char)108);
	printf("%c",(char)115);
	printf("%c",(char)101);
	goto L20;
	L20:
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
	/************ PRIMITIVO BOOLEAN ************/
	goto L21;
	goto L22;
	L21:
	t27 = 1 + 0;
	goto L23;
	L22:
	t27 = 0 + 0;
	goto L23;
	L23:
	/************ Printf Boolean ************/
	if(t27 == 1) goto L24;
	goto L25;
	L24:
	printf("%c",(char)116);
	printf("%c",(char)114);
	printf("%c",(char)117);
	printf("%c",(char)101);
	goto L26;
	L25:
	printf("%c",(char)102);
	printf("%c",(char)97);
	printf("%c",(char)108);
	printf("%c",(char)115);
	printf("%c",(char)101);
	goto L26;
	L26:
	/************ concat format {} ************/
	t28 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t29 = P + 1;
	t29 = t29 + 1;
	stack[(int)t29] = t28;
	P = P + 1;
	printfString();
	t30 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t31 = H + 0;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t32 = P + 1;
	t32 = t32 + 1;
	stack[(int)t32] = t31;
	P = P + 1;
	printfString();
	t33 = stack[(int)P];
	P = P - 1;
	/************ Identificador ************/
	t34 = stack[(int)0];
	/************ Printf Boolean ************/
	if(t34 == 1) goto L27;
	goto L28;
	L27:
	printf("%c",(char)116);
	printf("%c",(char)114);
	printf("%c",(char)117);
	printf("%c",(char)101);
	goto L29;
	L28:
	printf("%c",(char)102);
	printf("%c",(char)97);
	printf("%c",(char)108);
	printf("%c",(char)115);
	printf("%c",(char)101);
	goto L29;
	L29:
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return;
}
