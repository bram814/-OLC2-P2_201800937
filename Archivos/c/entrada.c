#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58, t59, t60, t61, t62, t63, t64, t65, t66, t67, t68, t69, t70, t71, t72, t73, t74, t75, t76, t77, t78, t79, t80, t81, t82, t83, t84, t85, t86, t87, t88, t89, t90, t91, t92, t93, t94, t95;

	
/************ NATIVE PRINTF STRING ************/
	
void printfString(){
	t38 = P + 1;
	t39 = stack[(int) t38];
	L15:
	t40 = heap[(int)t39];
	if(t40 == -1) goto L14;
	printf("%c", (int)t40);
	t39 = t39 + 1;
	goto L15;
	L14:
	return;
}
	/************ Fucniones ************/
	
void valor() {
	/************ Identificador ************/
	t0 = P + 1;
	t1 = stack[(int)t0];
	/************ Return ************/
	/************ Identificador ************/
	t2 = P + 1;
	t3 = stack[(int)t2];
	stack[(int)P] = t3;
	goto L0;
	L0:
	
	return;
}
	
void ack() {
	/************ Identificador ************/
	t4 = P + 1;
	t5 = stack[(int)t4];
	/************ Identificador ************/
	t6 = P + 2;
	t7 = stack[(int)t6];
	/************ Identificador ************/
	t8 = P + 1;
	t9 = stack[(int)t8];
	/************ Relacional == ************/
	if(t9 == 0) goto L2;
	goto L3;
	L2:
	t10 = 1 + 0;
	goto L4;
	L3:
	t10 = 0 + 0;
	goto L4;
	L4:
	/************ Control - If ************/
	if(t10 == 1) goto L5;
	goto L6;
	L5:
	/************ Return ************/
	/************ Identificador ************/
	t11 = P + 2;
	t12 = stack[(int)t11];
	/************ Aritmetica + ************/
	t13 = t12 + 1;
	stack[(int)P] = t13;
	goto L1;
	goto L7;
	L6:
	/************ Control - If (else if) ************/
	/************ Identificador ************/
	t14 = P + 2;
	t15 = stack[(int)t14];
	/************ Relacional == ************/
	if(t15 == 0) goto L8;
	goto L9;
	L8:
	t16 = 1 + 0;
	goto L10;
	L9:
	t16 = 0 + 0;
	goto L10;
	L10:
	/************ Control - If ************/
	if(t16 == 1) goto L11;
	goto L12;
	L11:
	/************ Return ************/
	/************ Llamada de Funcion ************/
	/************ Identificador ************/
	t19 = P + 1;
	t20 = stack[(int)t19];
	/************ Aritmetica - ************/
	t18 = t20 - 1;
	t17 = P + 4;
	stack[(int)t17] = t18;
	t17 = t17 + 1;
	stack[(int)t17] = 1;
	P = P + 3;
	ack();
	t17 = stack[(int)P];
	P = P - 3;
	stack[(int)P] = t17;
	goto L1;
	goto L13;
	L12:
	/************ Control - If (else) ************/
	/************ Return ************/
	/************ Llamada de Funcion ************/
	/************ Identificador ************/
	t23 = P + 1;
	t24 = stack[(int)t23];
	/************ Aritmetica - ************/
	t22 = t24 - 1;
	/************ Guardando Temporal ************/
	t25 = P + 3;
	stack[(int)t25] = t22;
	/************ Llamada de Funcion ************/
	/************ Identificador ************/
	t27 = P + 1;
	t28 = stack[(int)t27];
	t26 = P + 5;
	stack[(int)t26] = t28;
	/************ Identificador ************/
	t30 = P + 2;
	t31 = stack[(int)t30];
	/************ Aritmetica - ************/
	t29 = t31 - 1;
	t26 = t26 + 1;
	stack[(int)t26] = t29;
	P = P + 4;
	ack();
	t26 = stack[(int)P];
	P = P - 4;
	/************ Guardando Temporal ************/
	t32 = P + 4;
	stack[(int)t32] = t26;
	/************ Llamado de Temporales ************/
	t21 = P + 6;
	t33 = P + 3;
	t34 = stack[(int)t33];
	stack[(int)t21] = t34;
	/************ Llamado de Temporales ************/
	t21 = t21 + 1;
	t35 = P + 4;
	t36 = stack[(int)t35];
	stack[(int)t21] = t36;
	P = P + 5;
	ack();
	t21 = stack[(int)P];
	P = P - 5;
	stack[(int)P] = t21;
	goto L1;
	goto L13;
	L13:
	goto L7;
	L7:
	L1:
	
	return;
}
	/************ Main ************/
	
int main() {
	P = 0; H = 0;

	/************ Printf format {} ************/
	t37 = H + 0;
	heap[(int)H] = 65;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t41 = P + 0;
	t41 = t41 + 1;
	stack[(int)t41] = t37;
	P = P + 0;
	printfString();
	t42 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t43 = H + 0;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t44 = P + 0;
	t44 = t44 + 1;
	stack[(int)t44] = t43;
	P = P + 0;
	printfString();
	t45 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t46 = H + 0;
	heap[(int)H] = 107;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t47 = P + 0;
	t47 = t47 + 1;
	stack[(int)t47] = t46;
	P = P + 0;
	printfString();
	t48 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t49 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t50 = P + 0;
	t50 = t50 + 1;
	stack[(int)t50] = t49;
	P = P + 0;
	printfString();
	t51 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t52 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t53 = P + 0;
	t53 = t53 + 1;
	stack[(int)t53] = t52;
	P = P + 0;
	printfString();
	t54 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t55 = H + 0;
	heap[(int)H] = 109;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t56 = P + 0;
	t56 = t56 + 1;
	stack[(int)t56] = t55;
	P = P + 0;
	printfString();
	t57 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t58 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t59 = P + 0;
	t59 = t59 + 1;
	stack[(int)t59] = t58;
	P = P + 0;
	printfString();
	t60 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t61 = H + 0;
	heap[(int)H] = 110;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t62 = P + 0;
	t62 = t62 + 1;
	stack[(int)t62] = t61;
	P = P + 0;
	printfString();
	t63 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t64 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t65 = P + 0;
	t65 = t65 + 1;
	stack[(int)t65] = t64;
	P = P + 0;
	printfString();
	t66 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t67 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t68 = P + 0;
	t68 = t68 + 1;
	stack[(int)t68] = t67;
	P = P + 0;
	printfString();
	t69 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t70 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t71 = P + 0;
	t71 = t71 + 1;
	stack[(int)t71] = t70;
	P = P + 0;
	printfString();
	t72 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t73 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t74 = P + 0;
	t74 = t74 + 1;
	stack[(int)t74] = t73;
	P = P + 0;
	printfString();
	t75 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t76 = H + 0;
	heap[(int)H] = 51;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t77 = P + 0;
	t77 = t77 + 1;
	stack[(int)t77] = t76;
	P = P + 0;
	printfString();
	t78 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t79 = H + 0;
	heap[(int)H] = 44;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t80 = P + 0;
	t80 = t80 + 1;
	stack[(int)t80] = t79;
	P = P + 0;
	printfString();
	t81 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t82 = H + 0;
	heap[(int)H] = 48;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t83 = P + 0;
	t83 = t83 + 1;
	stack[(int)t83] = t82;
	P = P + 0;
	printfString();
	t84 = stack[(int)P];
	P = P - 0;
	/************ Printf format {} ************/
	t85 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t86 = P + 0;
	t86 = t86 + 1;
	stack[(int)t86] = t85;
	P = P + 0;
	printfString();
	t87 = stack[(int)P];
	P = P - 0;
	/************ Llamada de Funcion ************/
	/************ Guardando Temporal ************/
	t89 = P + 0;
	stack[(int)t89] = 2;
	/************ Llamada de Funcion ************/
	t90 = P + 2;
	stack[(int)t90] = 8;
	P = P + 1;
	valor();
	t90 = stack[(int)P];
	P = P - 1;
	/************ Guardando Temporal ************/
	t91 = P + 1;
	stack[(int)t91] = t90;
	/************ Llamado de Temporales ************/
	t88 = P + 3;
	t92 = P + 0;
	t93 = stack[(int)t92];
	stack[(int)t88] = t93;
	/************ Llamado de Temporales ************/
	t88 = t88 + 1;
	t94 = P + 1;
	t95 = stack[(int)t94];
	stack[(int)t88] = t95;
	P = P + 2;
	ack();
	t88 = stack[(int)P];
	P = P - 2;
	/************ Printf Integer ************/
	printf("%d",(int)t88);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return 0;
}
