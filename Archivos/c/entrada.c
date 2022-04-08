#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58, t59, t60, t61, t62, t63, t64, t65, t66, t67, t68, t69, t70, t71, t72, t73, t74, t75, t76, t77, t78, t79, t80, t81, t82, t83, t84, t85, t86, t87, t88, t89, t90, t91, t92, t93, t94, t95, t96, t97, t98, t99, t100, t101, t102, t103, t104, t105, t106, t107, t108, t109, t110, t111, t112, t113, t114, t115, t116, t117, t118, t119, t120, t121, t122, t123, t124, t125, t126, t127, t128, t129, t130, t131, t132, t133, t134, t135, t136, t137, t138, t139, t140, t141, t142, t143, t144, t145, t146, t147, t148, t149, t150, t151, t152, t153, t154, t155, t156, t157, t158, t159, t160, t161, t162, t163, t164, t165, t166, t167, t168, t169, t170, t171, t172, t173, t174, t175, t176, t177, t178, t179, t180, t181, t182, t183, t184, t185, t186, t187, t188, t189, t190, t191, t192, t193, t194, t195, t196, t197, t198, t199, t200, t201, t202, t203, t204, t205, t206, t207, t208, t209, t210, t211, t212, t213, t214, t215, t216, t217, t218, t219, t220, t221, t222, t223, t224, t225, t226, t227, t228, t229, t230, t231, t232, t233, t234, t235, t236, t237, t238, t239, t240, t241, t242, t243, t244, t245, t246, t247, t248, t249, t250, t251, t252, t253, t254, t255, t256, t257, t258, t259, t260, t261, t262, t263, t264, t265, t266, t267, t268, t269, t270, t271, t272, t273, t274, t275, t276, t277, t278, t279, t280, t281, t282, t283, t284, t285, t286, t287, t288, t289, t290, t291, t292, t293, t294, t295, t296, t297, t298, t299, t300, t301, t302, t303, t304, t305, t306, t307, t308, t309, t310, t311, t312, t313, t314, t315, t316, t317, t318, t319, t320, t321, t322, t323, t324, t325, t326, t327, t328, t329, t330, t331, t332, t333, t334, t335, t336, t337, t338, t339, t340, t341, t342, t343, t344, t345, t346, t347, t348, t349, t350, t351, t352, t353, t354, t355, t356, t357, t358, t359, t360, t361, t362, t363, t364, t365, t366, t367, t368, t369, t370, t371, t372, t373, t374, t375, t376, t377, t378, t379, t380, t381, t382, t383, t384, t385, t386, t387, t388, t389, t390, t391, t392, t393, t394, t395, t396, t397, t398, t399, t400, t401, t402, t403, t404, t405, t406, t407, t408, t409, t410, t411, t412, t413, t414, t415, t416, t417, t418, t419, t420, t421, t422, t423, t424, t425, t426, t427, t428, t429, t430, t431, t432, t433, t434, t435, t436, t437, t438, t439, t440, t441, t442, t443, t444, t445, t446, t447, t448, t449, t450, t451, t452, t453, t454, t455, t456, t457, t458, t459, t460, t461, t462, t463, t464, t465, t466, t467, t468, t469, t470, t471, t472, t473, t474, t475, t476, t477, t478, t479, t480, t481, t482, t483, t484, t485, t486, t487, t488, t489, t490, t491, t492, t493, t494, t495, t496, t497, t498, t499, t500, t501, t502, t503, t504, t505, t506, t507, t508, t509, t510, t511, t512, t513, t514, t515, t516, t517, t518, t519, t520, t521, t522, t523, t524, t525, t526, t527, t528, t529, t530, t531, t532, t533, t534, t535, t536, t537, t538, t539, t540, t541, t542, t543, t544, t545, t546, t547, t548, t549, t550, t551, t552, t553, t554, t555, t556, t557, t558, t559, t560, t561, t562, t563, t564, t565, t566, t567, t568, t569, t570, t571, t572, t573, t574, t575, t576, t577, t578, t579, t580, t581, t582, t583, t584, t585, t586, t587, t588, t589, t590, t591, t592, t593, t594, t595, t596, t597, t598, t599, t600, t601, t602, t603, t604, t605, t606, t607, t608, t609, t610, t611, t612, t613, t614, t615, t616, t617, t618, t619, t620, t621, t622, t623, t624, t625, t626, t627, t628, t629, t630, t631, t632, t633, t634, t635, t636, t637, t638, t639, t640, t641, t642, t643, t644, t645, t646, t647, t648, t649, t650, t651, t652, t653, t654, t655, t656, t657, t658, t659, t660, t661, t662, t663, t664, t665, t666, t667, t668, t669, t670, t671, t672, t673, t674, t675, t676, t677, t678, t679, t680, t681, t682, t683, t684, t685, t686, t687, t688, t689, t690, t691, t692, t693, t694, t695, t696, t697, t698, t699, t700, t701, t702, t703, t704, t705, t706, t707, t708, t709, t710, t711, t712, t713, t714, t715, t716, t717, t718, t719, t720, t721, t722, t723, t724, t725, t726, t727, t728, t729, t730, t731, t732, t733, t734, t735, t736, t737, t738, t739, t740, t741, t742, t743, t744, t745, t746, t747, t748, t749, t750, t751, t752, t753, t754, t755, t756, t757, t758, t759, t760, t761, t762, t763, t764, t765, t766, t767, t768, t769, t770, t771, t772, t773, t774, t775, t776, t777, t778, t779, t780, t781, t782, t783, t784, t785, t786, t787, t788, t789, t790, t791, t792, t793, t794, t795, t796, t797, t798, t799, t800, t801, t802, t803, t804, t805, t806, t807, t808, t809, t810, t811, t812, t813, t814, t815, t816, t817, t818, t819, t820, t821, t822, t823, t824, t825, t826, t827, t828, t829, t830, t831, t832, t833, t834, t835, t836, t837, t838, t839, t840, t841, t842, t843, t844, t845, t846, t847, t848, t849, t850, t851, t852, t853, t854, t855, t856, t857, t858, t859, t860, t861, t862, t863, t864, t865, t866, t867, t868, t869, t870, t871, t872, t873, t874, t875, t876, t877, t878, t879, t880, t881, t882, t883, t884, t885, t886, t887, t888, t889, t890, t891, t892, t893, t894, t895, t896, t897, t898, t899, t900, t901, t902, t903, t904, t905, t906, t907, t908, t909, t910, t911, t912, t913, t914, t915, t916, t917, t918, t919, t920, t921, t922, t923, t924, t925, t926, t927, t928, t929, t930, t931, t932, t933, t934, t935, t936, t937, t938, t939, t940, t941, t942, t943, t944, t945, t946, t947, t948, t949, t950, t951, t952, t953, t954, t955, t956, t957, t958, t959, t960, t961, t962, t963, t964, t965, t966, t967, t968, t969, t970, t971, t972, t973, t974, t975, t976, t977, t978, t979, t980, t981, t982, t983, t984, t985, t986, t987, t988, t989, t990, t991, t992, t993, t994, t995, t996, t997, t998, t999, t1000, t1001, t1002, t1003, t1004, t1005, t1006, t1007, t1008, t1009, t1010, t1011, t1012, t1013, t1014, t1015, t1016;

	
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

	/************ PRIMITIVO STRING ************/
	t0 = H + 0;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t4 = P + 0;
	t4 = t4 + 1;
	stack[(int)t4] = t0;
	printfString();
	t5 = stack[(int)P];
	P = P - 0;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ PRIMITIVO STRING ************/
	t6 = H + 0;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 65;
	H = H + 1;
	heap[(int)H] = 82;
	H = H + 1;
	heap[(int)H] = 67;
	H = H + 1;
	heap[(int)H] = 72;
	H = H + 1;
	heap[(int)H] = 73;
	H = H + 1;
	heap[(int)H] = 86;
	H = H + 1;
	heap[(int)H] = 79;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 66;
	H = H + 1;
	heap[(int)H] = 65;
	H = H + 1;
	heap[(int)H] = 83;
	H = H + 1;
	heap[(int)H] = 73;
	H = H + 1;
	heap[(int)H] = 67;
	H = H + 1;
	heap[(int)H] = 79;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t7 = P + 0;
	t7 = t7 + 1;
	stack[(int)t7] = t6;
	printfString();
	t8 = stack[(int)P];
	P = P - 0;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ PRIMITIVO STRING ************/
	t9 = H + 0;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t10 = P + 0;
	t10 = t10 + 1;
	stack[(int)t10] = t9;
	printfString();
	t11 = stack[(int)P];
	P = P - 0;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ PRIMITIVO BOOLEAN ************/
	goto L3;
	goto L2;
	L2:
	t13 = 1 + 0;
	goto L5;
	L3:
	t13 = 0 + 0;
	goto L5;
	L5:
	/************ Declaracion ************/
	stack[(int)0] = t12;
	/************ Identificador ************/
	t14 = stack[(int)0];
	/************ Logico ! ************/
	if(t14 == 1) goto L6;
	goto L7;
	L6:
	t15 = 0 + 0;
	goto L8;
	L7:
	t15 = 1 + 0;
	goto L8;
	L8:
	/************ Declaracion ************/
	stack[(int)1] = t15;
	/************ PRIMITIVO STRING ************/
	t16 = H + 0;
	heap[(int)H] = 105;
	H = H + 1;
	heap[(int)H] = 109;
	H = H + 1;
	heap[(int)H] = 112;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 105;
	H = H + 1;
	heap[(int)H] = 109;
	H = H + 1;
	heap[(int)H] = 105;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Declaracion ************/
	stack[(int)2] = t16;
	/************ PRIMITIVO STRING ************/
	t17 = H + 0;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 110;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 105;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Declaracion ************/
	stack[(int)3] = t17;
	/************ PRIMITIVO STRING ************/
	t18 = H + 0;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Declaracion ************/
	stack[(int)4] = t18;
	/************ Aritmetica * ************/
	t21 = 2 * 3;
	/************ Aritmetica + ************/
	t22 = 5 + t21;
	/************ Aritmetica * ************/
	t23 = 4 * t22;
	/************ Aritmetica + ************/
	t24 = 2 + t23;
	/************ Aritmetica * ************/
	t25 = 10 * t24;
	/************ Aritmetica + ************/
	t26 = 5 + t25;
	/************ Aritmetica * ************/
	t27 = 8 * 3;
	/************ Aritmetica * ************/
	t28 = t27 * 3;
	/************ Aritmetica - ************/
	t20 = t26 - t28;
	/************ Aritmetica - ************/
	t19 = 7 - t20;
	/************ Aritmetica * ************/
	t29 = 6 * 2;
	/************ Aritmetica * ************/
	t30 = 50 * t29;
	/************ Aritmetica + ************/
	t31 = t19 + t30;
	/************ Declaracion ************/
	stack[(int)5] = t31;
	/************ Aritmetica * ************/
	t35 = 2 * 2;
	/************ Aritmetica * ************/
	t36 = t35 * 2;
	/************ Aritmetica * ************/
	t37 = t36 * 2;
	/************ Aritmetica - ************/
	t34 = t37 - 9;
	/************ Aritmetica - ************/
	t39 = 8 - 6;
	/************ Aritmetica * ************/
	t43 = 3 * 3;
	/************ Aritmetica * ************/
	t44 = 6 * 5;
	/************ Aritmetica - ************/
	t42 = t43 - t44;
	/************ Aritmetica - ************/
	t41 = t42 - 7;
	/************ Aritmetica * ************/
	t45 = 7 * 7;
	/************ Aritmetica * ************/
	t46 = t45 * 7;
	/************ Aritmetica + ************/
	t47 = 9 + t46;
	/************ Aritmetica - ************/
	t40 = t41 - t47;
	/************ Aritmetica + ************/
	t48 = t40 + 10;
	/************ Aritmetica + ************/
	t49 = t39 + t48;
	/************ Aritmetica - ************/
	t38 = t49 - 5;
	/************ Aritmetica - ************/
	t33 = t34 - t38;
	/************ Aritmetica + ************/
	t50 = t33 + 8;
	/************ Aritmetica * ************/
	t52 = 2 * 3;
	/************ Aritmetica * ************/
	t53 = 5 * t52;
	/************ Aritmetica - ************/
	t51 = 6 - t53;
	/************ Aritmetica - ************/
	t32 = t50 - t51;
	/************ Declaracion ************/
	stack[(int)6] = t32;
	/************ Identificador ************/
	t55 = stack[(int)5];
	/************ Identificador ************/
	t57 = stack[(int)6];
	/************ Aritmetica * ************/
	t58 = t57 * 3;
	/************ Aritmetica + ************/
	t59 = 2 + t58;
	/************ Aritmetica + ************/
	t60 = t59 + 1;
	/************ Aritmetica * ************/
	t62 = 2 * 2;
	/************ Aritmetica * ************/
	t63 = t62 * 2;
	/************ Aritmetica - ************/
	t61 = t63 - 2;
	/************ Aritmetica * ************/
	t64 = t61 * 2;
	/************ Aritmetica - ************/
	t56 = t60 - t64;
	/************ Aritmetica + ************/
	t65 = t55 + t56;
	/************ Aritmetica - ************/
	t54 = t65 - 2;
	/************ Declaracion ************/
	stack[(int)7] = t54;
	/************ Printf format {} ************/
	t66 = H + 0;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t67 = P + 8;
	t67 = t67 + 1;
	stack[(int)t67] = t66;
	P = P + 8;
	printfString();
	t68 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t69 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t70 = P + 8;
	t70 = t70 + 1;
	stack[(int)t70] = t69;
	P = P + 8;
	printfString();
	t71 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t72 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t73 = P + 8;
	t73 = t73 + 1;
	stack[(int)t73] = t72;
	P = P + 8;
	printfString();
	t74 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t75 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t76 = P + 8;
	t76 = t76 + 1;
	stack[(int)t76] = t75;
	P = P + 8;
	printfString();
	t77 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t78 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t79 = P + 8;
	t79 = t79 + 1;
	stack[(int)t79] = t78;
	P = P + 8;
	printfString();
	t80 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t81 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t82 = P + 8;
	t82 = t82 + 1;
	stack[(int)t82] = t81;
	P = P + 8;
	printfString();
	t83 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t84 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t85 = P + 8;
	t85 = t85 + 1;
	stack[(int)t85] = t84;
	P = P + 8;
	printfString();
	t86 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t87 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t88 = P + 8;
	t88 = t88 + 1;
	stack[(int)t88] = t87;
	P = P + 8;
	printfString();
	t89 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t90 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t91 = P + 8;
	t91 = t91 + 1;
	stack[(int)t91] = t90;
	P = P + 8;
	printfString();
	t92 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t93 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t94 = P + 8;
	t94 = t94 + 1;
	stack[(int)t94] = t93;
	P = P + 8;
	printfString();
	t95 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t96 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t97 = P + 8;
	t97 = t97 + 1;
	stack[(int)t97] = t96;
	P = P + 8;
	printfString();
	t98 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t99 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t100 = P + 8;
	t100 = t100 + 1;
	stack[(int)t100] = t99;
	P = P + 8;
	printfString();
	t101 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t102 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t103 = P + 8;
	t103 = t103 + 1;
	stack[(int)t103] = t102;
	P = P + 8;
	printfString();
	t104 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t105 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t106 = P + 8;
	t106 = t106 + 1;
	stack[(int)t106] = t105;
	P = P + 8;
	printfString();
	t107 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t108 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t109 = P + 8;
	t109 = t109 + 1;
	stack[(int)t109] = t108;
	P = P + 8;
	printfString();
	t110 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t111 = H + 0;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t112 = P + 8;
	t112 = t112 + 1;
	stack[(int)t112] = t111;
	P = P + 8;
	printfString();
	t113 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t114 = H + 0;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t115 = P + 8;
	t115 = t115 + 1;
	stack[(int)t115] = t114;
	P = P + 8;
	printfString();
	t116 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t117 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t118 = P + 8;
	t118 = t118 + 1;
	stack[(int)t118] = t117;
	P = P + 8;
	printfString();
	t119 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t120 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t121 = P + 8;
	t121 = t121 + 1;
	stack[(int)t121] = t120;
	P = P + 8;
	printfString();
	t122 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t123 = H + 0;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t124 = P + 8;
	t124 = t124 + 1;
	stack[(int)t124] = t123;
	P = P + 8;
	printfString();
	t125 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t126 = H + 0;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t127 = P + 8;
	t127 = t127 + 1;
	stack[(int)t127] = t126;
	P = P + 8;
	printfString();
	t128 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t129 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t130 = P + 8;
	t130 = t130 + 1;
	stack[(int)t130] = t129;
	P = P + 8;
	printfString();
	t131 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t132 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t133 = P + 8;
	t133 = t133 + 1;
	stack[(int)t133] = t132;
	P = P + 8;
	printfString();
	t134 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t135 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t136 = P + 8;
	t136 = t136 + 1;
	stack[(int)t136] = t135;
	P = P + 8;
	printfString();
	t137 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t138 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t139 = P + 8;
	t139 = t139 + 1;
	stack[(int)t139] = t138;
	P = P + 8;
	printfString();
	t140 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t141 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t142 = P + 8;
	t142 = t142 + 1;
	stack[(int)t142] = t141;
	P = P + 8;
	printfString();
	t143 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t144 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t145 = P + 8;
	t145 = t145 + 1;
	stack[(int)t145] = t144;
	P = P + 8;
	printfString();
	t146 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t147 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t148 = P + 8;
	t148 = t148 + 1;
	stack[(int)t148] = t147;
	P = P + 8;
	printfString();
	t149 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t150 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t151 = P + 8;
	t151 = t151 + 1;
	stack[(int)t151] = t150;
	P = P + 8;
	printfString();
	t152 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t153 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t154 = P + 8;
	t154 = t154 + 1;
	stack[(int)t154] = t153;
	P = P + 8;
	printfString();
	t155 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t156 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t157 = P + 8;
	t157 = t157 + 1;
	stack[(int)t157] = t156;
	P = P + 8;
	printfString();
	t158 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t159 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t160 = P + 8;
	t160 = t160 + 1;
	stack[(int)t160] = t159;
	P = P + 8;
	printfString();
	t161 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t162 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t163 = P + 8;
	t163 = t163 + 1;
	stack[(int)t163] = t162;
	P = P + 8;
	printfString();
	t164 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t165 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t166 = P + 8;
	t166 = t166 + 1;
	stack[(int)t166] = t165;
	P = P + 8;
	printfString();
	t167 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t168 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t169 = P + 8;
	t169 = t169 + 1;
	stack[(int)t169] = t168;
	P = P + 8;
	printfString();
	t170 = stack[(int)P];
	P = P - 8;
	/************ Identificador ************/
	t171 = stack[(int)5];
	/************ Printf Integer ************/
	printf("%d",(int)t171);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Printf format {} ************/
	t172 = H + 0;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t173 = P + 8;
	t173 = t173 + 1;
	stack[(int)t173] = t172;
	P = P + 8;
	printfString();
	t174 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t175 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t176 = P + 8;
	t176 = t176 + 1;
	stack[(int)t176] = t175;
	P = P + 8;
	printfString();
	t177 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t178 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t179 = P + 8;
	t179 = t179 + 1;
	stack[(int)t179] = t178;
	P = P + 8;
	printfString();
	t180 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t181 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t182 = P + 8;
	t182 = t182 + 1;
	stack[(int)t182] = t181;
	P = P + 8;
	printfString();
	t183 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t184 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t185 = P + 8;
	t185 = t185 + 1;
	stack[(int)t185] = t184;
	P = P + 8;
	printfString();
	t186 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t187 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t188 = P + 8;
	t188 = t188 + 1;
	stack[(int)t188] = t187;
	P = P + 8;
	printfString();
	t189 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t190 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t191 = P + 8;
	t191 = t191 + 1;
	stack[(int)t191] = t190;
	P = P + 8;
	printfString();
	t192 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t193 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t194 = P + 8;
	t194 = t194 + 1;
	stack[(int)t194] = t193;
	P = P + 8;
	printfString();
	t195 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t196 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t197 = P + 8;
	t197 = t197 + 1;
	stack[(int)t197] = t196;
	P = P + 8;
	printfString();
	t198 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t199 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t200 = P + 8;
	t200 = t200 + 1;
	stack[(int)t200] = t199;
	P = P + 8;
	printfString();
	t201 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t202 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t203 = P + 8;
	t203 = t203 + 1;
	stack[(int)t203] = t202;
	P = P + 8;
	printfString();
	t204 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t205 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t206 = P + 8;
	t206 = t206 + 1;
	stack[(int)t206] = t205;
	P = P + 8;
	printfString();
	t207 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t208 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t209 = P + 8;
	t209 = t209 + 1;
	stack[(int)t209] = t208;
	P = P + 8;
	printfString();
	t210 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t211 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t212 = P + 8;
	t212 = t212 + 1;
	stack[(int)t212] = t211;
	P = P + 8;
	printfString();
	t213 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t214 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t215 = P + 8;
	t215 = t215 + 1;
	stack[(int)t215] = t214;
	P = P + 8;
	printfString();
	t216 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t217 = H + 0;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t218 = P + 8;
	t218 = t218 + 1;
	stack[(int)t218] = t217;
	P = P + 8;
	printfString();
	t219 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t220 = H + 0;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t221 = P + 8;
	t221 = t221 + 1;
	stack[(int)t221] = t220;
	P = P + 8;
	printfString();
	t222 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t223 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t224 = P + 8;
	t224 = t224 + 1;
	stack[(int)t224] = t223;
	P = P + 8;
	printfString();
	t225 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t226 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t227 = P + 8;
	t227 = t227 + 1;
	stack[(int)t227] = t226;
	P = P + 8;
	printfString();
	t228 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t229 = H + 0;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t230 = P + 8;
	t230 = t230 + 1;
	stack[(int)t230] = t229;
	P = P + 8;
	printfString();
	t231 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t232 = H + 0;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t233 = P + 8;
	t233 = t233 + 1;
	stack[(int)t233] = t232;
	P = P + 8;
	printfString();
	t234 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t235 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t236 = P + 8;
	t236 = t236 + 1;
	stack[(int)t236] = t235;
	P = P + 8;
	printfString();
	t237 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t238 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t239 = P + 8;
	t239 = t239 + 1;
	stack[(int)t239] = t238;
	P = P + 8;
	printfString();
	t240 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t241 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t242 = P + 8;
	t242 = t242 + 1;
	stack[(int)t242] = t241;
	P = P + 8;
	printfString();
	t243 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t244 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t245 = P + 8;
	t245 = t245 + 1;
	stack[(int)t245] = t244;
	P = P + 8;
	printfString();
	t246 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t247 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t248 = P + 8;
	t248 = t248 + 1;
	stack[(int)t248] = t247;
	P = P + 8;
	printfString();
	t249 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t250 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t251 = P + 8;
	t251 = t251 + 1;
	stack[(int)t251] = t250;
	P = P + 8;
	printfString();
	t252 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t253 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t254 = P + 8;
	t254 = t254 + 1;
	stack[(int)t254] = t253;
	P = P + 8;
	printfString();
	t255 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t256 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t257 = P + 8;
	t257 = t257 + 1;
	stack[(int)t257] = t256;
	P = P + 8;
	printfString();
	t258 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t259 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t260 = P + 8;
	t260 = t260 + 1;
	stack[(int)t260] = t259;
	P = P + 8;
	printfString();
	t261 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t262 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t263 = P + 8;
	t263 = t263 + 1;
	stack[(int)t263] = t262;
	P = P + 8;
	printfString();
	t264 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t265 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t266 = P + 8;
	t266 = t266 + 1;
	stack[(int)t266] = t265;
	P = P + 8;
	printfString();
	t267 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t268 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t269 = P + 8;
	t269 = t269 + 1;
	stack[(int)t269] = t268;
	P = P + 8;
	printfString();
	t270 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t271 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t272 = P + 8;
	t272 = t272 + 1;
	stack[(int)t272] = t271;
	P = P + 8;
	printfString();
	t273 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t274 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t275 = P + 8;
	t275 = t275 + 1;
	stack[(int)t275] = t274;
	P = P + 8;
	printfString();
	t276 = stack[(int)P];
	P = P - 8;
	/************ Identificador ************/
	t277 = stack[(int)6];
	/************ Printf Integer ************/
	printf("%d",(int)t277);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Printf format {} ************/
	t278 = H + 0;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t279 = P + 8;
	t279 = t279 + 1;
	stack[(int)t279] = t278;
	P = P + 8;
	printfString();
	t280 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t281 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t282 = P + 8;
	t282 = t282 + 1;
	stack[(int)t282] = t281;
	P = P + 8;
	printfString();
	t283 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t284 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t285 = P + 8;
	t285 = t285 + 1;
	stack[(int)t285] = t284;
	P = P + 8;
	printfString();
	t286 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t287 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t288 = P + 8;
	t288 = t288 + 1;
	stack[(int)t288] = t287;
	P = P + 8;
	printfString();
	t289 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t290 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t291 = P + 8;
	t291 = t291 + 1;
	stack[(int)t291] = t290;
	P = P + 8;
	printfString();
	t292 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t293 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t294 = P + 8;
	t294 = t294 + 1;
	stack[(int)t294] = t293;
	P = P + 8;
	printfString();
	t295 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t296 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t297 = P + 8;
	t297 = t297 + 1;
	stack[(int)t297] = t296;
	P = P + 8;
	printfString();
	t298 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t299 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t300 = P + 8;
	t300 = t300 + 1;
	stack[(int)t300] = t299;
	P = P + 8;
	printfString();
	t301 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t302 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t303 = P + 8;
	t303 = t303 + 1;
	stack[(int)t303] = t302;
	P = P + 8;
	printfString();
	t304 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t305 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t306 = P + 8;
	t306 = t306 + 1;
	stack[(int)t306] = t305;
	P = P + 8;
	printfString();
	t307 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t308 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t309 = P + 8;
	t309 = t309 + 1;
	stack[(int)t309] = t308;
	P = P + 8;
	printfString();
	t310 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t311 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t312 = P + 8;
	t312 = t312 + 1;
	stack[(int)t312] = t311;
	P = P + 8;
	printfString();
	t313 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t314 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t315 = P + 8;
	t315 = t315 + 1;
	stack[(int)t315] = t314;
	P = P + 8;
	printfString();
	t316 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t317 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t318 = P + 8;
	t318 = t318 + 1;
	stack[(int)t318] = t317;
	P = P + 8;
	printfString();
	t319 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t320 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t321 = P + 8;
	t321 = t321 + 1;
	stack[(int)t321] = t320;
	P = P + 8;
	printfString();
	t322 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t323 = H + 0;
	heap[(int)H] = 51;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t324 = P + 8;
	t324 = t324 + 1;
	stack[(int)t324] = t323;
	P = P + 8;
	printfString();
	t325 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t326 = H + 0;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t327 = P + 8;
	t327 = t327 + 1;
	stack[(int)t327] = t326;
	P = P + 8;
	printfString();
	t328 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t329 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t330 = P + 8;
	t330 = t330 + 1;
	stack[(int)t330] = t329;
	P = P + 8;
	printfString();
	t331 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t332 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t333 = P + 8;
	t333 = t333 + 1;
	stack[(int)t333] = t332;
	P = P + 8;
	printfString();
	t334 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t335 = H + 0;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t336 = P + 8;
	t336 = t336 + 1;
	stack[(int)t336] = t335;
	P = P + 8;
	printfString();
	t337 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t338 = H + 0;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t339 = P + 8;
	t339 = t339 + 1;
	stack[(int)t339] = t338;
	P = P + 8;
	printfString();
	t340 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t341 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t342 = P + 8;
	t342 = t342 + 1;
	stack[(int)t342] = t341;
	P = P + 8;
	printfString();
	t343 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t344 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t345 = P + 8;
	t345 = t345 + 1;
	stack[(int)t345] = t344;
	P = P + 8;
	printfString();
	t346 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t347 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t348 = P + 8;
	t348 = t348 + 1;
	stack[(int)t348] = t347;
	P = P + 8;
	printfString();
	t349 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t350 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t351 = P + 8;
	t351 = t351 + 1;
	stack[(int)t351] = t350;
	P = P + 8;
	printfString();
	t352 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t353 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t354 = P + 8;
	t354 = t354 + 1;
	stack[(int)t354] = t353;
	P = P + 8;
	printfString();
	t355 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t356 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t357 = P + 8;
	t357 = t357 + 1;
	stack[(int)t357] = t356;
	P = P + 8;
	printfString();
	t358 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t359 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t360 = P + 8;
	t360 = t360 + 1;
	stack[(int)t360] = t359;
	P = P + 8;
	printfString();
	t361 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t362 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t363 = P + 8;
	t363 = t363 + 1;
	stack[(int)t363] = t362;
	P = P + 8;
	printfString();
	t364 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t365 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t366 = P + 8;
	t366 = t366 + 1;
	stack[(int)t366] = t365;
	P = P + 8;
	printfString();
	t367 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t368 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t369 = P + 8;
	t369 = t369 + 1;
	stack[(int)t369] = t368;
	P = P + 8;
	printfString();
	t370 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t371 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t372 = P + 8;
	t372 = t372 + 1;
	stack[(int)t372] = t371;
	P = P + 8;
	printfString();
	t373 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t374 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t375 = P + 8;
	t375 = t375 + 1;
	stack[(int)t375] = t374;
	P = P + 8;
	printfString();
	t376 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t377 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t378 = P + 8;
	t378 = t378 + 1;
	stack[(int)t378] = t377;
	P = P + 8;
	printfString();
	t379 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t380 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t381 = P + 8;
	t381 = t381 + 1;
	stack[(int)t381] = t380;
	P = P + 8;
	printfString();
	t382 = stack[(int)P];
	P = P - 8;
	/************ Identificador ************/
	t383 = stack[(int)7];
	/************ Printf Integer ************/
	printf("%d",(int)t383);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Printf format {} ************/
	t384 = H + 0;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t385 = P + 8;
	t385 = t385 + 1;
	stack[(int)t385] = t384;
	P = P + 8;
	printfString();
	t386 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t387 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t388 = P + 8;
	t388 = t388 + 1;
	stack[(int)t388] = t387;
	P = P + 8;
	printfString();
	t389 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t390 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t391 = P + 8;
	t391 = t391 + 1;
	stack[(int)t391] = t390;
	P = P + 8;
	printfString();
	t392 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t393 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t394 = P + 8;
	t394 = t394 + 1;
	stack[(int)t394] = t393;
	P = P + 8;
	printfString();
	t395 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t396 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t397 = P + 8;
	t397 = t397 + 1;
	stack[(int)t397] = t396;
	P = P + 8;
	printfString();
	t398 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t399 = H + 0;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t400 = P + 8;
	t400 = t400 + 1;
	stack[(int)t400] = t399;
	P = P + 8;
	printfString();
	t401 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t402 = H + 0;
	heap[(int)H] = 117;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t403 = P + 8;
	t403 = t403 + 1;
	stack[(int)t403] = t402;
	P = P + 8;
	printfString();
	t404 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t405 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t406 = P + 8;
	t406 = t406 + 1;
	stack[(int)t406] = t405;
	P = P + 8;
	printfString();
	t407 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t408 = H + 0;
	heap[(int)H] = 116;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t409 = P + 8;
	t409 = t409 + 1;
	stack[(int)t409] = t408;
	P = P + 8;
	printfString();
	t410 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t411 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t412 = P + 8;
	t412 = t412 + 1;
	stack[(int)t412] = t411;
	P = P + 8;
	printfString();
	t413 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t414 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t415 = P + 8;
	t415 = t415 + 1;
	stack[(int)t415] = t414;
	P = P + 8;
	printfString();
	t416 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t417 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t418 = P + 8;
	t418 = t418 + 1;
	stack[(int)t418] = t417;
	P = P + 8;
	printfString();
	t419 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t420 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t421 = P + 8;
	t421 = t421 + 1;
	stack[(int)t421] = t420;
	P = P + 8;
	printfString();
	t422 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t423 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t424 = P + 8;
	t424 = t424 + 1;
	stack[(int)t424] = t423;
	P = P + 8;
	printfString();
	t425 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t426 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t427 = P + 8;
	t427 = t427 + 1;
	stack[(int)t427] = t426;
	P = P + 8;
	printfString();
	t428 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t429 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t430 = P + 8;
	t430 = t430 + 1;
	stack[(int)t430] = t429;
	P = P + 8;
	printfString();
	t431 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t432 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t433 = P + 8;
	t433 = t433 + 1;
	stack[(int)t433] = t432;
	P = P + 8;
	printfString();
	t434 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t435 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t436 = P + 8;
	t436 = t436 + 1;
	stack[(int)t436] = t435;
	P = P + 8;
	printfString();
	t437 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t438 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t439 = P + 8;
	t439 = t439 + 1;
	stack[(int)t439] = t438;
	P = P + 8;
	printfString();
	t440 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t441 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t442 = P + 8;
	t442 = t442 + 1;
	stack[(int)t442] = t441;
	P = P + 8;
	printfString();
	t443 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t444 = H + 0;
	heap[(int)H] = 112;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t445 = P + 8;
	t445 = t445 + 1;
	stack[(int)t445] = t444;
	P = P + 8;
	printfString();
	t446 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t447 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t448 = P + 8;
	t448 = t448 + 1;
	stack[(int)t448] = t447;
	P = P + 8;
	printfString();
	t449 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t450 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t451 = P + 8;
	t451 = t451 + 1;
	stack[(int)t451] = t450;
	P = P + 8;
	printfString();
	t452 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t453 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t454 = P + 8;
	t454 = t454 + 1;
	stack[(int)t454] = t453;
	P = P + 8;
	printfString();
	t455 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t456 = H + 0;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t457 = P + 8;
	t457 = t457 + 1;
	stack[(int)t457] = t456;
	P = P + 8;
	printfString();
	t458 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t459 = H + 0;
	heap[(int)H] = 105;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t460 = P + 8;
	t460 = t460 + 1;
	stack[(int)t460] = t459;
	P = P + 8;
	printfString();
	t461 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t462 = H + 0;
	heap[(int)H] = 195;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t463 = P + 8;
	t463 = t463 + 1;
	stack[(int)t463] = t462;
	P = P + 8;
	printfString();
	t464 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t465 = H + 0;
	heap[(int)H] = 179;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t466 = P + 8;
	t466 = t466 + 1;
	stack[(int)t466] = t465;
	P = P + 8;
	printfString();
	t467 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t468 = H + 0;
	heap[(int)H] = 110;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t469 = P + 8;
	t469 = t469 + 1;
	stack[(int)t469] = t468;
	P = P + 8;
	printfString();
	t470 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t471 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t472 = P + 8;
	t472 = t472 + 1;
	stack[(int)t472] = t471;
	P = P + 8;
	printfString();
	t473 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t474 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t475 = P + 8;
	t475 = t475 + 1;
	stack[(int)t475] = t474;
	P = P + 8;
	printfString();
	t476 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t477 = H + 0;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t478 = P + 8;
	t478 = t478 + 1;
	stack[(int)t478] = t477;
	P = P + 8;
	printfString();
	t479 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t480 = H + 0;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t481 = P + 8;
	t481 = t481 + 1;
	stack[(int)t481] = t480;
	P = P + 8;
	printfString();
	t482 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t483 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t484 = P + 8;
	t484 = t484 + 1;
	stack[(int)t484] = t483;
	P = P + 8;
	printfString();
	t485 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t486 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t487 = P + 8;
	t487 = t487 + 1;
	stack[(int)t487] = t486;
	P = P + 8;
	printfString();
	t488 = stack[(int)P];
	P = P - 8;
	/************ Identificador ************/
	t489 = stack[(int)7];
	/************ Printf Integer ************/
	printf("%d",(int)t489);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Printf format {} ************/
	t490 = H + 0;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t491 = P + 8;
	t491 = t491 + 1;
	stack[(int)t491] = t490;
	P = P + 8;
	printfString();
	t492 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t493 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t494 = P + 8;
	t494 = t494 + 1;
	stack[(int)t494] = t493;
	P = P + 8;
	printfString();
	t495 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t496 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t497 = P + 8;
	t497 = t497 + 1;
	stack[(int)t497] = t496;
	P = P + 8;
	printfString();
	t498 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t499 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t500 = P + 8;
	t500 = t500 + 1;
	stack[(int)t500] = t499;
	P = P + 8;
	printfString();
	t501 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t502 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t503 = P + 8;
	t503 = t503 + 1;
	stack[(int)t503] = t502;
	P = P + 8;
	printfString();
	t504 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t505 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t506 = P + 8;
	t506 = t506 + 1;
	stack[(int)t506] = t505;
	P = P + 8;
	printfString();
	t507 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t508 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t509 = P + 8;
	t509 = t509 + 1;
	stack[(int)t509] = t508;
	P = P + 8;
	printfString();
	t510 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t511 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t512 = P + 8;
	t512 = t512 + 1;
	stack[(int)t512] = t511;
	P = P + 8;
	printfString();
	t513 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t514 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t515 = P + 8;
	t515 = t515 + 1;
	stack[(int)t515] = t514;
	P = P + 8;
	printfString();
	t516 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t517 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t518 = P + 8;
	t518 = t518 + 1;
	stack[(int)t518] = t517;
	P = P + 8;
	printfString();
	t519 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t520 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t521 = P + 8;
	t521 = t521 + 1;
	stack[(int)t521] = t520;
	P = P + 8;
	printfString();
	t522 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t523 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t524 = P + 8;
	t524 = t524 + 1;
	stack[(int)t524] = t523;
	P = P + 8;
	printfString();
	t525 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t526 = H + 0;
	heap[(int)H] = 98;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t527 = P + 8;
	t527 = t527 + 1;
	stack[(int)t527] = t526;
	P = P + 8;
	printfString();
	t528 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t529 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t530 = P + 8;
	t530 = t530 + 1;
	stack[(int)t530] = t529;
	P = P + 8;
	printfString();
	t531 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t532 = H + 0;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t533 = P + 8;
	t533 = t533 + 1;
	stack[(int)t533] = t532;
	P = P + 8;
	printfString();
	t534 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t535 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t536 = P + 8;
	t536 = t536 + 1;
	stack[(int)t536] = t535;
	P = P + 8;
	printfString();
	t537 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t538 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t539 = P + 8;
	t539 = t539 + 1;
	stack[(int)t539] = t538;
	P = P + 8;
	printfString();
	t540 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t541 = H + 0;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t542 = P + 8;
	t542 = t542 + 1;
	stack[(int)t542] = t541;
	P = P + 8;
	printfString();
	t543 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t544 = H + 0;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t545 = P + 8;
	t545 = t545 + 1;
	stack[(int)t545] = t544;
	P = P + 8;
	printfString();
	t546 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t547 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t548 = P + 8;
	t548 = t548 + 1;
	stack[(int)t548] = t547;
	P = P + 8;
	printfString();
	t549 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t550 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t551 = P + 8;
	t551 = t551 + 1;
	stack[(int)t551] = t550;
	P = P + 8;
	printfString();
	t552 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t553 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t554 = P + 8;
	t554 = t554 + 1;
	stack[(int)t554] = t553;
	P = P + 8;
	printfString();
	t555 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t556 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t557 = P + 8;
	t557 = t557 + 1;
	stack[(int)t557] = t556;
	P = P + 8;
	printfString();
	t558 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t559 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t560 = P + 8;
	t560 = t560 + 1;
	stack[(int)t560] = t559;
	P = P + 8;
	printfString();
	t561 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t562 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t563 = P + 8;
	t563 = t563 + 1;
	stack[(int)t563] = t562;
	P = P + 8;
	printfString();
	t564 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t565 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t566 = P + 8;
	t566 = t566 + 1;
	stack[(int)t566] = t565;
	P = P + 8;
	printfString();
	t567 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t568 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t569 = P + 8;
	t569 = t569 + 1;
	stack[(int)t569] = t568;
	P = P + 8;
	printfString();
	t570 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t571 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t572 = P + 8;
	t572 = t572 + 1;
	stack[(int)t572] = t571;
	P = P + 8;
	printfString();
	t573 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t574 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t575 = P + 8;
	t575 = t575 + 1;
	stack[(int)t575] = t574;
	P = P + 8;
	printfString();
	t576 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t577 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t578 = P + 8;
	t578 = t578 + 1;
	stack[(int)t578] = t577;
	P = P + 8;
	printfString();
	t579 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t580 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t581 = P + 8;
	t581 = t581 + 1;
	stack[(int)t581] = t580;
	P = P + 8;
	printfString();
	t582 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t583 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t584 = P + 8;
	t584 = t584 + 1;
	stack[(int)t584] = t583;
	P = P + 8;
	printfString();
	t585 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t586 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t587 = P + 8;
	t587 = t587 + 1;
	stack[(int)t587] = t586;
	P = P + 8;
	printfString();
	t588 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t589 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t590 = P + 8;
	t590 = t590 + 1;
	stack[(int)t590] = t589;
	P = P + 8;
	printfString();
	t591 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t592 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t593 = P + 8;
	t593 = t593 + 1;
	stack[(int)t593] = t592;
	P = P + 8;
	printfString();
	t594 = stack[(int)P];
	P = P - 8;
	/************ Identificador ************/
	t595 = stack[(int)0];
	/************ Printf Boolean ************/
	if(t595 == 1) goto L9;
	goto L10;
	L9:
	printf("%c",(char)116);
	printf("%c",(char)114);
	printf("%c",(char)117);
	printf("%c",(char)101);
	goto L11;
	L10:
	printf("%c",(char)102);
	printf("%c",(char)97);
	printf("%c",(char)108);
	printf("%c",(char)115);
	printf("%c",(char)101);
	goto L11;
	L11:
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Printf format {} ************/
	t596 = H + 0;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t597 = P + 8;
	t597 = t597 + 1;
	stack[(int)t597] = t596;
	P = P + 8;
	printfString();
	t598 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t599 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t600 = P + 8;
	t600 = t600 + 1;
	stack[(int)t600] = t599;
	P = P + 8;
	printfString();
	t601 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t602 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t603 = P + 8;
	t603 = t603 + 1;
	stack[(int)t603] = t602;
	P = P + 8;
	printfString();
	t604 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t605 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t606 = P + 8;
	t606 = t606 + 1;
	stack[(int)t606] = t605;
	P = P + 8;
	printfString();
	t607 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t608 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t609 = P + 8;
	t609 = t609 + 1;
	stack[(int)t609] = t608;
	P = P + 8;
	printfString();
	t610 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t611 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t612 = P + 8;
	t612 = t612 + 1;
	stack[(int)t612] = t611;
	P = P + 8;
	printfString();
	t613 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t614 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t615 = P + 8;
	t615 = t615 + 1;
	stack[(int)t615] = t614;
	P = P + 8;
	printfString();
	t616 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t617 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t618 = P + 8;
	t618 = t618 + 1;
	stack[(int)t618] = t617;
	P = P + 8;
	printfString();
	t619 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t620 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t621 = P + 8;
	t621 = t621 + 1;
	stack[(int)t621] = t620;
	P = P + 8;
	printfString();
	t622 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t623 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t624 = P + 8;
	t624 = t624 + 1;
	stack[(int)t624] = t623;
	P = P + 8;
	printfString();
	t625 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t626 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t627 = P + 8;
	t627 = t627 + 1;
	stack[(int)t627] = t626;
	P = P + 8;
	printfString();
	t628 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t629 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t630 = P + 8;
	t630 = t630 + 1;
	stack[(int)t630] = t629;
	P = P + 8;
	printfString();
	t631 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t632 = H + 0;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t633 = P + 8;
	t633 = t633 + 1;
	stack[(int)t633] = t632;
	P = P + 8;
	printfString();
	t634 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t635 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t636 = P + 8;
	t636 = t636 + 1;
	stack[(int)t636] = t635;
	P = P + 8;
	printfString();
	t637 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t638 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t639 = P + 8;
	t639 = t639 + 1;
	stack[(int)t639] = t638;
	P = P + 8;
	printfString();
	t640 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t641 = H + 0;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t642 = P + 8;
	t642 = t642 + 1;
	stack[(int)t642] = t641;
	P = P + 8;
	printfString();
	t643 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t644 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t645 = P + 8;
	t645 = t645 + 1;
	stack[(int)t645] = t644;
	P = P + 8;
	printfString();
	t646 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t647 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t648 = P + 8;
	t648 = t648 + 1;
	stack[(int)t648] = t647;
	P = P + 8;
	printfString();
	t649 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t650 = H + 0;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t651 = P + 8;
	t651 = t651 + 1;
	stack[(int)t651] = t650;
	P = P + 8;
	printfString();
	t652 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t653 = H + 0;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t654 = P + 8;
	t654 = t654 + 1;
	stack[(int)t654] = t653;
	P = P + 8;
	printfString();
	t655 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t656 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t657 = P + 8;
	t657 = t657 + 1;
	stack[(int)t657] = t656;
	P = P + 8;
	printfString();
	t658 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t659 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t660 = P + 8;
	t660 = t660 + 1;
	stack[(int)t660] = t659;
	P = P + 8;
	printfString();
	t661 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t662 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t663 = P + 8;
	t663 = t663 + 1;
	stack[(int)t663] = t662;
	P = P + 8;
	printfString();
	t664 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t665 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t666 = P + 8;
	t666 = t666 + 1;
	stack[(int)t666] = t665;
	P = P + 8;
	printfString();
	t667 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t668 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t669 = P + 8;
	t669 = t669 + 1;
	stack[(int)t669] = t668;
	P = P + 8;
	printfString();
	t670 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t671 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t672 = P + 8;
	t672 = t672 + 1;
	stack[(int)t672] = t671;
	P = P + 8;
	printfString();
	t673 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t674 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t675 = P + 8;
	t675 = t675 + 1;
	stack[(int)t675] = t674;
	P = P + 8;
	printfString();
	t676 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t677 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t678 = P + 8;
	t678 = t678 + 1;
	stack[(int)t678] = t677;
	P = P + 8;
	printfString();
	t679 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t680 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t681 = P + 8;
	t681 = t681 + 1;
	stack[(int)t681] = t680;
	P = P + 8;
	printfString();
	t682 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t683 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t684 = P + 8;
	t684 = t684 + 1;
	stack[(int)t684] = t683;
	P = P + 8;
	printfString();
	t685 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t686 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t687 = P + 8;
	t687 = t687 + 1;
	stack[(int)t687] = t686;
	P = P + 8;
	printfString();
	t688 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t689 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t690 = P + 8;
	t690 = t690 + 1;
	stack[(int)t690] = t689;
	P = P + 8;
	printfString();
	t691 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t692 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t693 = P + 8;
	t693 = t693 + 1;
	stack[(int)t693] = t692;
	P = P + 8;
	printfString();
	t694 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t695 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t696 = P + 8;
	t696 = t696 + 1;
	stack[(int)t696] = t695;
	P = P + 8;
	printfString();
	t697 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t698 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t699 = P + 8;
	t699 = t699 + 1;
	stack[(int)t699] = t698;
	P = P + 8;
	printfString();
	t700 = stack[(int)P];
	P = P - 8;
	/************ Identificador ************/
	t701 = stack[(int)2];
	/************ Printf String ************/
	t702 = P + 8;
	t702 = t702 + 1;
	stack[(int)t702] = t701;
	P = P + 8;
	printfString();
	t703 = stack[(int)P];
	P = P - 8;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Printf format {} ************/
	t704 = H + 0;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t705 = P + 8;
	t705 = t705 + 1;
	stack[(int)t705] = t704;
	P = P + 8;
	printfString();
	t706 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t707 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t708 = P + 8;
	t708 = t708 + 1;
	stack[(int)t708] = t707;
	P = P + 8;
	printfString();
	t709 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t710 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t711 = P + 8;
	t711 = t711 + 1;
	stack[(int)t711] = t710;
	P = P + 8;
	printfString();
	t712 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t713 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t714 = P + 8;
	t714 = t714 + 1;
	stack[(int)t714] = t713;
	P = P + 8;
	printfString();
	t715 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t716 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t717 = P + 8;
	t717 = t717 + 1;
	stack[(int)t717] = t716;
	P = P + 8;
	printfString();
	t718 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t719 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t720 = P + 8;
	t720 = t720 + 1;
	stack[(int)t720] = t719;
	P = P + 8;
	printfString();
	t721 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t722 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t723 = P + 8;
	t723 = t723 + 1;
	stack[(int)t723] = t722;
	P = P + 8;
	printfString();
	t724 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t725 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t726 = P + 8;
	t726 = t726 + 1;
	stack[(int)t726] = t725;
	P = P + 8;
	printfString();
	t727 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t728 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t729 = P + 8;
	t729 = t729 + 1;
	stack[(int)t729] = t728;
	P = P + 8;
	printfString();
	t730 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t731 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t732 = P + 8;
	t732 = t732 + 1;
	stack[(int)t732] = t731;
	P = P + 8;
	printfString();
	t733 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t734 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t735 = P + 8;
	t735 = t735 + 1;
	stack[(int)t735] = t734;
	P = P + 8;
	printfString();
	t736 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t737 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t738 = P + 8;
	t738 = t738 + 1;
	stack[(int)t738] = t737;
	P = P + 8;
	printfString();
	t739 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t740 = H + 0;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t741 = P + 8;
	t741 = t741 + 1;
	stack[(int)t741] = t740;
	P = P + 8;
	printfString();
	t742 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t743 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t744 = P + 8;
	t744 = t744 + 1;
	stack[(int)t744] = t743;
	P = P + 8;
	printfString();
	t745 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t746 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t747 = P + 8;
	t747 = t747 + 1;
	stack[(int)t747] = t746;
	P = P + 8;
	printfString();
	t748 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t749 = H + 0;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t750 = P + 8;
	t750 = t750 + 1;
	stack[(int)t750] = t749;
	P = P + 8;
	printfString();
	t751 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t752 = H + 0;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t753 = P + 8;
	t753 = t753 + 1;
	stack[(int)t753] = t752;
	P = P + 8;
	printfString();
	t754 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t755 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t756 = P + 8;
	t756 = t756 + 1;
	stack[(int)t756] = t755;
	P = P + 8;
	printfString();
	t757 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t758 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t759 = P + 8;
	t759 = t759 + 1;
	stack[(int)t759] = t758;
	P = P + 8;
	printfString();
	t760 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t761 = H + 0;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t762 = P + 8;
	t762 = t762 + 1;
	stack[(int)t762] = t761;
	P = P + 8;
	printfString();
	t763 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t764 = H + 0;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t765 = P + 8;
	t765 = t765 + 1;
	stack[(int)t765] = t764;
	P = P + 8;
	printfString();
	t766 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t767 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t768 = P + 8;
	t768 = t768 + 1;
	stack[(int)t768] = t767;
	P = P + 8;
	printfString();
	t769 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t770 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t771 = P + 8;
	t771 = t771 + 1;
	stack[(int)t771] = t770;
	P = P + 8;
	printfString();
	t772 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t773 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t774 = P + 8;
	t774 = t774 + 1;
	stack[(int)t774] = t773;
	P = P + 8;
	printfString();
	t775 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t776 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t777 = P + 8;
	t777 = t777 + 1;
	stack[(int)t777] = t776;
	P = P + 8;
	printfString();
	t778 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t779 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t780 = P + 8;
	t780 = t780 + 1;
	stack[(int)t780] = t779;
	P = P + 8;
	printfString();
	t781 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t782 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t783 = P + 8;
	t783 = t783 + 1;
	stack[(int)t783] = t782;
	P = P + 8;
	printfString();
	t784 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t785 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t786 = P + 8;
	t786 = t786 + 1;
	stack[(int)t786] = t785;
	P = P + 8;
	printfString();
	t787 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t788 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t789 = P + 8;
	t789 = t789 + 1;
	stack[(int)t789] = t788;
	P = P + 8;
	printfString();
	t790 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t791 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t792 = P + 8;
	t792 = t792 + 1;
	stack[(int)t792] = t791;
	P = P + 8;
	printfString();
	t793 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t794 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t795 = P + 8;
	t795 = t795 + 1;
	stack[(int)t795] = t794;
	P = P + 8;
	printfString();
	t796 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t797 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t798 = P + 8;
	t798 = t798 + 1;
	stack[(int)t798] = t797;
	P = P + 8;
	printfString();
	t799 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t800 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t801 = P + 8;
	t801 = t801 + 1;
	stack[(int)t801] = t800;
	P = P + 8;
	printfString();
	t802 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t803 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t804 = P + 8;
	t804 = t804 + 1;
	stack[(int)t804] = t803;
	P = P + 8;
	printfString();
	t805 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t806 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t807 = P + 8;
	t807 = t807 + 1;
	stack[(int)t807] = t806;
	P = P + 8;
	printfString();
	t808 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t809 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t810 = P + 8;
	t810 = t810 + 1;
	stack[(int)t810] = t809;
	P = P + 8;
	printfString();
	t811 = stack[(int)P];
	P = P - 8;
	/************ Identificador ************/
	t812 = stack[(int)3];
	/************ Printf String ************/
	t813 = P + 8;
	t813 = t813 + 1;
	stack[(int)t813] = t812;
	P = P + 8;
	printfString();
	t814 = stack[(int)P];
	P = P - 8;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Printf format {} ************/
	t815 = H + 0;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t816 = P + 8;
	t816 = t816 + 1;
	stack[(int)t816] = t815;
	P = P + 8;
	printfString();
	t817 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t818 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t819 = P + 8;
	t819 = t819 + 1;
	stack[(int)t819] = t818;
	P = P + 8;
	printfString();
	t820 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t821 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t822 = P + 8;
	t822 = t822 + 1;
	stack[(int)t822] = t821;
	P = P + 8;
	printfString();
	t823 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t824 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t825 = P + 8;
	t825 = t825 + 1;
	stack[(int)t825] = t824;
	P = P + 8;
	printfString();
	t826 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t827 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t828 = P + 8;
	t828 = t828 + 1;
	stack[(int)t828] = t827;
	P = P + 8;
	printfString();
	t829 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t830 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t831 = P + 8;
	t831 = t831 + 1;
	stack[(int)t831] = t830;
	P = P + 8;
	printfString();
	t832 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t833 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t834 = P + 8;
	t834 = t834 + 1;
	stack[(int)t834] = t833;
	P = P + 8;
	printfString();
	t835 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t836 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t837 = P + 8;
	t837 = t837 + 1;
	stack[(int)t837] = t836;
	P = P + 8;
	printfString();
	t838 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t839 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t840 = P + 8;
	t840 = t840 + 1;
	stack[(int)t840] = t839;
	P = P + 8;
	printfString();
	t841 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t842 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t843 = P + 8;
	t843 = t843 + 1;
	stack[(int)t843] = t842;
	P = P + 8;
	printfString();
	t844 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t845 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t846 = P + 8;
	t846 = t846 + 1;
	stack[(int)t846] = t845;
	P = P + 8;
	printfString();
	t847 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t848 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t849 = P + 8;
	t849 = t849 + 1;
	stack[(int)t849] = t848;
	P = P + 8;
	printfString();
	t850 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t851 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t852 = P + 8;
	t852 = t852 + 1;
	stack[(int)t852] = t851;
	P = P + 8;
	printfString();
	t853 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t854 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t855 = P + 8;
	t855 = t855 + 1;
	stack[(int)t855] = t854;
	P = P + 8;
	printfString();
	t856 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t857 = H + 0;
	heap[(int)H] = 116;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t858 = P + 8;
	t858 = t858 + 1;
	stack[(int)t858] = t857;
	P = P + 8;
	printfString();
	t859 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t860 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t861 = P + 8;
	t861 = t861 + 1;
	stack[(int)t861] = t860;
	P = P + 8;
	printfString();
	t862 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t863 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t864 = P + 8;
	t864 = t864 + 1;
	stack[(int)t864] = t863;
	P = P + 8;
	printfString();
	t865 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t866 = H + 0;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t867 = P + 8;
	t867 = t867 + 1;
	stack[(int)t867] = t866;
	P = P + 8;
	printfString();
	t868 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t869 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t870 = P + 8;
	t870 = t870 + 1;
	stack[(int)t870] = t869;
	P = P + 8;
	printfString();
	t871 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t872 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t873 = P + 8;
	t873 = t873 + 1;
	stack[(int)t873] = t872;
	P = P + 8;
	printfString();
	t874 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t875 = H + 0;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t876 = P + 8;
	t876 = t876 + 1;
	stack[(int)t876] = t875;
	P = P + 8;
	printfString();
	t877 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t878 = H + 0;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t879 = P + 8;
	t879 = t879 + 1;
	stack[(int)t879] = t878;
	P = P + 8;
	printfString();
	t880 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t881 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t882 = P + 8;
	t882 = t882 + 1;
	stack[(int)t882] = t881;
	P = P + 8;
	printfString();
	t883 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t884 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t885 = P + 8;
	t885 = t885 + 1;
	stack[(int)t885] = t884;
	P = P + 8;
	printfString();
	t886 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t887 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t888 = P + 8;
	t888 = t888 + 1;
	stack[(int)t888] = t887;
	P = P + 8;
	printfString();
	t889 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t890 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t891 = P + 8;
	t891 = t891 + 1;
	stack[(int)t891] = t890;
	P = P + 8;
	printfString();
	t892 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t893 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t894 = P + 8;
	t894 = t894 + 1;
	stack[(int)t894] = t893;
	P = P + 8;
	printfString();
	t895 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t896 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t897 = P + 8;
	t897 = t897 + 1;
	stack[(int)t897] = t896;
	P = P + 8;
	printfString();
	t898 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t899 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t900 = P + 8;
	t900 = t900 + 1;
	stack[(int)t900] = t899;
	P = P + 8;
	printfString();
	t901 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t902 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t903 = P + 8;
	t903 = t903 + 1;
	stack[(int)t903] = t902;
	P = P + 8;
	printfString();
	t904 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t905 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t906 = P + 8;
	t906 = t906 + 1;
	stack[(int)t906] = t905;
	P = P + 8;
	printfString();
	t907 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t908 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t909 = P + 8;
	t909 = t909 + 1;
	stack[(int)t909] = t908;
	P = P + 8;
	printfString();
	t910 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t911 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t912 = P + 8;
	t912 = t912 + 1;
	stack[(int)t912] = t911;
	P = P + 8;
	printfString();
	t913 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t914 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t915 = P + 8;
	t915 = t915 + 1;
	stack[(int)t915] = t914;
	P = P + 8;
	printfString();
	t916 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t917 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t918 = P + 8;
	t918 = t918 + 1;
	stack[(int)t918] = t917;
	P = P + 8;
	printfString();
	t919 = stack[(int)P];
	P = P - 8;
	/************ Identificador ************/
	t920 = stack[(int)4];
	/************ Printf String ************/
	t921 = P + 8;
	t921 = t921 + 1;
	stack[(int)t921] = t920;
	P = P + 8;
	printfString();
	t922 = stack[(int)P];
	P = P - 8;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Printf format {} ************/
	t923 = H + 0;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t924 = P + 8;
	t924 = t924 + 1;
	stack[(int)t924] = t923;
	P = P + 8;
	printfString();
	t925 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t926 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t927 = P + 8;
	t927 = t927 + 1;
	stack[(int)t927] = t926;
	P = P + 8;
	printfString();
	t928 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t929 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t930 = P + 8;
	t930 = t930 + 1;
	stack[(int)t930] = t929;
	P = P + 8;
	printfString();
	t931 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t932 = H + 0;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t933 = P + 8;
	t933 = t933 + 1;
	stack[(int)t933] = t932;
	P = P + 8;
	printfString();
	t934 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t935 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t936 = P + 8;
	t936 = t936 + 1;
	stack[(int)t936] = t935;
	P = P + 8;
	printfString();
	t937 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t938 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t939 = P + 8;
	t939 = t939 + 1;
	stack[(int)t939] = t938;
	P = P + 8;
	printfString();
	t940 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t941 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t942 = P + 8;
	t942 = t942 + 1;
	stack[(int)t942] = t941;
	P = P + 8;
	printfString();
	t943 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t944 = H + 0;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t945 = P + 8;
	t945 = t945 + 1;
	stack[(int)t945] = t944;
	P = P + 8;
	printfString();
	t946 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t947 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t948 = P + 8;
	t948 = t948 + 1;
	stack[(int)t948] = t947;
	P = P + 8;
	printfString();
	t949 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t950 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t951 = P + 8;
	t951 = t951 + 1;
	stack[(int)t951] = t950;
	P = P + 8;
	printfString();
	t952 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t953 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t954 = P + 8;
	t954 = t954 + 1;
	stack[(int)t954] = t953;
	P = P + 8;
	printfString();
	t955 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t956 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t957 = P + 8;
	t957 = t957 + 1;
	stack[(int)t957] = t956;
	P = P + 8;
	printfString();
	t958 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t959 = H + 0;
	heap[(int)H] = 98;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t960 = P + 8;
	t960 = t960 + 1;
	stack[(int)t960] = t959;
	P = P + 8;
	printfString();
	t961 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t962 = H + 0;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t963 = P + 8;
	t963 = t963 + 1;
	stack[(int)t963] = t962;
	P = P + 8;
	printfString();
	t964 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t965 = H + 0;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t966 = P + 8;
	t966 = t966 + 1;
	stack[(int)t966] = t965;
	P = P + 8;
	printfString();
	t967 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t968 = H + 0;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t969 = P + 8;
	t969 = t969 + 1;
	stack[(int)t969] = t968;
	P = P + 8;
	printfString();
	t970 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t971 = H + 0;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t972 = P + 8;
	t972 = t972 + 1;
	stack[(int)t972] = t971;
	P = P + 8;
	printfString();
	t973 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t974 = H + 0;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t975 = P + 8;
	t975 = t975 + 1;
	stack[(int)t975] = t974;
	P = P + 8;
	printfString();
	t976 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t977 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t978 = P + 8;
	t978 = t978 + 1;
	stack[(int)t978] = t977;
	P = P + 8;
	printfString();
	t979 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t980 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t981 = P + 8;
	t981 = t981 + 1;
	stack[(int)t981] = t980;
	P = P + 8;
	printfString();
	t982 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t983 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t984 = P + 8;
	t984 = t984 + 1;
	stack[(int)t984] = t983;
	P = P + 8;
	printfString();
	t985 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t986 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t987 = P + 8;
	t987 = t987 + 1;
	stack[(int)t987] = t986;
	P = P + 8;
	printfString();
	t988 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t989 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t990 = P + 8;
	t990 = t990 + 1;
	stack[(int)t990] = t989;
	P = P + 8;
	printfString();
	t991 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t992 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t993 = P + 8;
	t993 = t993 + 1;
	stack[(int)t993] = t992;
	P = P + 8;
	printfString();
	t994 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t995 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t996 = P + 8;
	t996 = t996 + 1;
	stack[(int)t996] = t995;
	P = P + 8;
	printfString();
	t997 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t998 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t999 = P + 8;
	t999 = t999 + 1;
	stack[(int)t999] = t998;
	P = P + 8;
	printfString();
	t1000 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t1001 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t1002 = P + 8;
	t1002 = t1002 + 1;
	stack[(int)t1002] = t1001;
	P = P + 8;
	printfString();
	t1003 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t1004 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t1005 = P + 8;
	t1005 = t1005 + 1;
	stack[(int)t1005] = t1004;
	P = P + 8;
	printfString();
	t1006 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t1007 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t1008 = P + 8;
	t1008 = t1008 + 1;
	stack[(int)t1008] = t1007;
	P = P + 8;
	printfString();
	t1009 = stack[(int)P];
	P = P - 8;
	/************ Printf format {} ************/
	t1010 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t1011 = P + 8;
	t1011 = t1011 + 1;
	stack[(int)t1011] = t1010;
	P = P + 8;
	printfString();
	t1012 = stack[(int)P];
	P = P - 8;
	/************ Identificador ************/
	t1013 = stack[(int)1];
	/************ Printf Boolean ************/
	if(t1013 == 1) goto L12;
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
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ PRIMITIVO STRING ************/
	t1014 = H + 0;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t1015 = P + 8;
	t1015 = t1015 + 1;
	stack[(int)t1015] = t1014;
	printfString();
	t1016 = stack[(int)P];
	P = P - 8;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return;
}
