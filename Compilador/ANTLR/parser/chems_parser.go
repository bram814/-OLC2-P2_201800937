// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2/Compilador/interfaces"
import "OLC2/Compilador/expression"
import "OLC2/Compilador/instruction"
import "OLC2/Compilador/instruction/loops"
import "OLC2/Compilador/instruction/control"
import "OLC2/Compilador/instruction/variable"
import "OLC2/Compilador/instruction/ternario"
import "OLC2/Compilador/instruction/function"
import "OLC2/Compilador/instruction/casteo"
import "OLC2/Compilador/instruction/transferencia"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 759,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 3,
	2, 3, 2, 3, 2, 3, 3, 6, 3, 103, 10, 3, 13, 3, 14, 3, 104, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 5, 4, 112, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 157, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 175, 10, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 235, 10, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 270, 10, 10, 3, 11, 7, 11, 273, 10, 11, 12, 11, 14, 11, 276,
	11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12,
	307, 10, 12, 3, 13, 7, 13, 310, 10, 13, 12, 13, 14, 13, 313, 11, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 332, 10, 14, 3, 15, 6,
	15, 335, 10, 15, 13, 15, 14, 15, 336, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5,
	16, 354, 10, 16, 3, 17, 6, 17, 357, 10, 17, 13, 17, 14, 17, 358, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 370, 10,
	18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 388, 10, 20, 3, 21, 6,
	21, 391, 10, 21, 13, 21, 14, 21, 392, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 5, 22, 412, 10, 22, 3, 23, 6, 23, 415, 10, 23, 13, 23, 14, 23,
	416, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 434, 10, 24, 3, 25, 6, 25,
	437, 10, 25, 13, 25, 14, 25, 438, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 5, 26, 450, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 465,
	10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 494, 10, 29,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 514, 10, 32,
	3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 5, 34, 550, 10, 34, 3, 35, 7, 35, 553, 10, 35, 12, 35,
	14, 35, 556, 11, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 571, 10, 36, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 583, 10,
	37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 5, 38, 597, 10, 38, 3, 39, 6, 39, 600, 10, 39, 13, 39, 14,
	39, 601, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	5, 40, 613, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 627, 10, 41, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 637, 10, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 7, 42, 644, 10, 42, 12, 42, 14, 42, 647, 11, 42, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 7, 43, 658, 10,
	43, 12, 43, 14, 43, 661, 11, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 676, 10, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 7,
	44, 688, 10, 44, 12, 44, 14, 44, 691, 11, 44, 3, 45, 3, 45, 3, 45, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 5, 46, 720, 10, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47, 746, 10,
	47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 5, 49,
	757, 10, 49, 3, 49, 2, 5, 82, 84, 86, 50, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90,
	92, 94, 96, 2, 6, 3, 2, 57, 58, 4, 2, 38, 39, 42, 45, 3, 2, 46, 48, 3,
	2, 49, 50, 2, 788, 2, 98, 3, 2, 2, 2, 4, 102, 3, 2, 2, 2, 6, 111, 3, 2,
	2, 2, 8, 156, 3, 2, 2, 2, 10, 174, 3, 2, 2, 2, 12, 176, 3, 2, 2, 2, 14,
	234, 3, 2, 2, 2, 16, 236, 3, 2, 2, 2, 18, 269, 3, 2, 2, 2, 20, 274, 3,
	2, 2, 2, 22, 306, 3, 2, 2, 2, 24, 311, 3, 2, 2, 2, 26, 331, 3, 2, 2, 2,
	28, 334, 3, 2, 2, 2, 30, 353, 3, 2, 2, 2, 32, 356, 3, 2, 2, 2, 34, 369,
	3, 2, 2, 2, 36, 371, 3, 2, 2, 2, 38, 387, 3, 2, 2, 2, 40, 390, 3, 2, 2,
	2, 42, 411, 3, 2, 2, 2, 44, 414, 3, 2, 2, 2, 46, 433, 3, 2, 2, 2, 48, 436,
	3, 2, 2, 2, 50, 449, 3, 2, 2, 2, 52, 464, 3, 2, 2, 2, 54, 466, 3, 2, 2,
	2, 56, 493, 3, 2, 2, 2, 58, 495, 3, 2, 2, 2, 60, 501, 3, 2, 2, 2, 62, 513,
	3, 2, 2, 2, 64, 515, 3, 2, 2, 2, 66, 549, 3, 2, 2, 2, 68, 554, 3, 2, 2,
	2, 70, 570, 3, 2, 2, 2, 72, 582, 3, 2, 2, 2, 74, 596, 3, 2, 2, 2, 76, 599,
	3, 2, 2, 2, 78, 612, 3, 2, 2, 2, 80, 626, 3, 2, 2, 2, 82, 636, 3, 2, 2,
	2, 84, 648, 3, 2, 2, 2, 86, 675, 3, 2, 2, 2, 88, 692, 3, 2, 2, 2, 90, 719,
	3, 2, 2, 2, 92, 745, 3, 2, 2, 2, 94, 747, 3, 2, 2, 2, 96, 756, 3, 2, 2,
	2, 98, 99, 5, 4, 3, 2, 99, 100, 8, 2, 1, 2, 100, 3, 3, 2, 2, 2, 101, 103,
	5, 8, 5, 2, 102, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 102, 3, 2,
	2, 2, 104, 105, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 8, 3, 1, 2,
	107, 5, 3, 2, 2, 2, 108, 109, 7, 34, 2, 2, 109, 112, 8, 4, 1, 2, 110, 112,
	8, 4, 1, 2, 111, 108, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2, 112, 7, 3, 2, 2,
	2, 113, 114, 5, 10, 6, 2, 114, 115, 5, 6, 4, 2, 115, 116, 8, 5, 1, 2, 116,
	157, 3, 2, 2, 2, 117, 118, 5, 12, 7, 2, 118, 119, 8, 5, 1, 2, 119, 157,
	3, 2, 2, 2, 120, 121, 5, 14, 8, 2, 121, 122, 8, 5, 1, 2, 122, 157, 3, 2,
	2, 2, 123, 124, 5, 16, 9, 2, 124, 125, 8, 5, 1, 2, 125, 157, 3, 2, 2, 2,
	126, 127, 5, 18, 10, 2, 127, 128, 8, 5, 1, 2, 128, 157, 3, 2, 2, 2, 129,
	130, 5, 26, 14, 2, 130, 131, 8, 5, 1, 2, 131, 157, 3, 2, 2, 2, 132, 133,
	5, 54, 28, 2, 133, 134, 8, 5, 1, 2, 134, 157, 3, 2, 2, 2, 135, 136, 5,
	56, 29, 2, 136, 137, 8, 5, 1, 2, 137, 157, 3, 2, 2, 2, 138, 139, 5, 58,
	30, 2, 139, 140, 8, 5, 1, 2, 140, 157, 3, 2, 2, 2, 141, 142, 5, 62, 32,
	2, 142, 143, 5, 6, 4, 2, 143, 144, 8, 5, 1, 2, 144, 157, 3, 2, 2, 2, 145,
	146, 5, 64, 33, 2, 146, 147, 5, 6, 4, 2, 147, 148, 8, 5, 1, 2, 148, 157,
	3, 2, 2, 2, 149, 150, 5, 66, 34, 2, 150, 151, 8, 5, 1, 2, 151, 157, 3,
	2, 2, 2, 152, 153, 5, 72, 37, 2, 153, 154, 5, 6, 4, 2, 154, 155, 8, 5,
	1, 2, 155, 157, 3, 2, 2, 2, 156, 113, 3, 2, 2, 2, 156, 117, 3, 2, 2, 2,
	156, 120, 3, 2, 2, 2, 156, 123, 3, 2, 2, 2, 156, 126, 3, 2, 2, 2, 156,
	129, 3, 2, 2, 2, 156, 132, 3, 2, 2, 2, 156, 135, 3, 2, 2, 2, 156, 138,
	3, 2, 2, 2, 156, 141, 3, 2, 2, 2, 156, 145, 3, 2, 2, 2, 156, 149, 3, 2,
	2, 2, 156, 152, 3, 2, 2, 2, 157, 9, 3, 2, 2, 2, 158, 159, 7, 4, 2, 2, 159,
	160, 7, 51, 2, 2, 160, 161, 5, 90, 46, 2, 161, 162, 7, 52, 2, 2, 162, 163,
	7, 34, 2, 2, 163, 164, 8, 6, 1, 2, 164, 175, 3, 2, 2, 2, 165, 166, 7, 4,
	2, 2, 166, 167, 7, 51, 2, 2, 167, 168, 7, 29, 2, 2, 168, 169, 7, 35, 2,
	2, 169, 170, 5, 76, 39, 2, 170, 171, 7, 52, 2, 2, 171, 172, 7, 34, 2, 2,
	172, 173, 8, 6, 1, 2, 173, 175, 3, 2, 2, 2, 174, 158, 3, 2, 2, 2, 174,
	165, 3, 2, 2, 2, 175, 11, 3, 2, 2, 2, 176, 177, 7, 18, 2, 2, 177, 178,
	7, 17, 2, 2, 178, 179, 7, 51, 2, 2, 179, 180, 7, 52, 2, 2, 180, 181, 7,
	53, 2, 2, 181, 182, 5, 4, 3, 2, 182, 183, 7, 54, 2, 2, 183, 184, 8, 7,
	1, 2, 184, 13, 3, 2, 2, 2, 185, 186, 7, 5, 2, 2, 186, 187, 7, 6, 2, 2,
	187, 188, 7, 31, 2, 2, 188, 189, 7, 37, 2, 2, 189, 190, 5, 80, 41, 2, 190,
	191, 7, 34, 2, 2, 191, 192, 8, 8, 1, 2, 192, 235, 3, 2, 2, 2, 193, 194,
	7, 5, 2, 2, 194, 195, 7, 6, 2, 2, 195, 196, 7, 31, 2, 2, 196, 197, 7, 36,
	2, 2, 197, 198, 5, 74, 38, 2, 198, 199, 7, 34, 2, 2, 199, 200, 8, 8, 1,
	2, 200, 235, 3, 2, 2, 2, 201, 202, 7, 5, 2, 2, 202, 203, 7, 6, 2, 2, 203,
	204, 7, 31, 2, 2, 204, 205, 7, 36, 2, 2, 205, 206, 5, 74, 38, 2, 206, 207,
	7, 37, 2, 2, 207, 208, 5, 80, 41, 2, 208, 209, 7, 34, 2, 2, 209, 210, 8,
	8, 1, 2, 210, 235, 3, 2, 2, 2, 211, 212, 7, 5, 2, 2, 212, 213, 7, 31, 2,
	2, 213, 214, 7, 37, 2, 2, 214, 215, 5, 80, 41, 2, 215, 216, 7, 34, 2, 2,
	216, 217, 8, 8, 1, 2, 217, 235, 3, 2, 2, 2, 218, 219, 7, 5, 2, 2, 219,
	220, 7, 31, 2, 2, 220, 221, 7, 36, 2, 2, 221, 222, 5, 74, 38, 2, 222, 223,
	7, 34, 2, 2, 223, 224, 8, 8, 1, 2, 224, 235, 3, 2, 2, 2, 225, 226, 7, 5,
	2, 2, 226, 227, 7, 31, 2, 2, 227, 228, 7, 36, 2, 2, 228, 229, 5, 74, 38,
	2, 229, 230, 7, 37, 2, 2, 230, 231, 5, 80, 41, 2, 231, 232, 7, 34, 2, 2,
	232, 233, 8, 8, 1, 2, 233, 235, 3, 2, 2, 2, 234, 185, 3, 2, 2, 2, 234,
	193, 3, 2, 2, 2, 234, 201, 3, 2, 2, 2, 234, 211, 3, 2, 2, 2, 234, 218,
	3, 2, 2, 2, 234, 225, 3, 2, 2, 2, 235, 15, 3, 2, 2, 2, 236, 237, 7, 31,
	2, 2, 237, 238, 7, 37, 2, 2, 238, 239, 5, 80, 41, 2, 239, 240, 7, 34, 2,
	2, 240, 241, 8, 9, 1, 2, 241, 17, 3, 2, 2, 2, 242, 243, 7, 7, 2, 2, 243,
	244, 5, 80, 41, 2, 244, 245, 7, 53, 2, 2, 245, 246, 5, 4, 3, 2, 246, 247,
	7, 54, 2, 2, 247, 248, 8, 10, 1, 2, 248, 270, 3, 2, 2, 2, 249, 250, 7,
	7, 2, 2, 250, 251, 5, 80, 41, 2, 251, 252, 7, 53, 2, 2, 252, 253, 5, 4,
	3, 2, 253, 254, 7, 54, 2, 2, 254, 255, 7, 8, 2, 2, 255, 256, 7, 53, 2,
	2, 256, 257, 5, 4, 3, 2, 257, 258, 7, 54, 2, 2, 258, 259, 8, 10, 1, 2,
	259, 270, 3, 2, 2, 2, 260, 261, 7, 7, 2, 2, 261, 262, 5, 80, 41, 2, 262,
	263, 7, 53, 2, 2, 263, 264, 5, 4, 3, 2, 264, 265, 7, 54, 2, 2, 265, 266,
	7, 8, 2, 2, 266, 267, 5, 20, 11, 2, 267, 268, 8, 10, 1, 2, 268, 270, 3,
	2, 2, 2, 269, 242, 3, 2, 2, 2, 269, 249, 3, 2, 2, 2, 269, 260, 3, 2, 2,
	2, 270, 19, 3, 2, 2, 2, 271, 273, 5, 18, 10, 2, 272, 271, 3, 2, 2, 2, 273,
	276, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 277,
	3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 277, 278, 8, 11, 1, 2, 278, 21, 3, 2,
	2, 2, 279, 280, 7, 7, 2, 2, 280, 281, 5, 80, 41, 2, 281, 282, 7, 53, 2,
	2, 282, 283, 5, 80, 41, 2, 283, 284, 7, 54, 2, 2, 284, 285, 8, 12, 1, 2,
	285, 307, 3, 2, 2, 2, 286, 287, 7, 7, 2, 2, 287, 288, 5, 80, 41, 2, 288,
	289, 7, 53, 2, 2, 289, 290, 5, 80, 41, 2, 290, 291, 7, 54, 2, 2, 291, 292,
	7, 8, 2, 2, 292, 293, 7, 53, 2, 2, 293, 294, 5, 80, 41, 2, 294, 295, 7,
	54, 2, 2, 295, 296, 8, 12, 1, 2, 296, 307, 3, 2, 2, 2, 297, 298, 7, 7,
	2, 2, 298, 299, 5, 80, 41, 2, 299, 300, 7, 53, 2, 2, 300, 301, 5, 80, 41,
	2, 301, 302, 7, 54, 2, 2, 302, 303, 7, 8, 2, 2, 303, 304, 5, 24, 13, 2,
	304, 305, 8, 12, 1, 2, 305, 307, 3, 2, 2, 2, 306, 279, 3, 2, 2, 2, 306,
	286, 3, 2, 2, 2, 306, 297, 3, 2, 2, 2, 307, 23, 3, 2, 2, 2, 308, 310, 5,
	22, 12, 2, 309, 308, 3, 2, 2, 2, 310, 313, 3, 2, 2, 2, 311, 309, 3, 2,
	2, 2, 311, 312, 3, 2, 2, 2, 312, 314, 3, 2, 2, 2, 313, 311, 3, 2, 2, 2,
	314, 315, 8, 13, 1, 2, 315, 25, 3, 2, 2, 2, 316, 317, 7, 9, 2, 2, 317,
	318, 5, 80, 41, 2, 318, 319, 7, 53, 2, 2, 319, 320, 5, 28, 15, 2, 320,
	321, 5, 40, 21, 2, 321, 322, 7, 54, 2, 2, 322, 323, 8, 14, 1, 2, 323, 332,
	3, 2, 2, 2, 324, 325, 7, 9, 2, 2, 325, 326, 5, 80, 41, 2, 326, 327, 7,
	53, 2, 2, 327, 328, 5, 40, 21, 2, 328, 329, 7, 54, 2, 2, 329, 330, 8, 14,
	1, 2, 330, 332, 3, 2, 2, 2, 331, 316, 3, 2, 2, 2, 331, 324, 3, 2, 2, 2,
	332, 27, 3, 2, 2, 2, 333, 335, 5, 30, 16, 2, 334, 333, 3, 2, 2, 2, 335,
	336, 3, 2, 2, 2, 336, 334, 3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 338,
	3, 2, 2, 2, 338, 339, 8, 15, 1, 2, 339, 29, 3, 2, 2, 2, 340, 341, 5, 32,
	17, 2, 341, 342, 7, 40, 2, 2, 342, 343, 7, 53, 2, 2, 343, 344, 5, 4, 3,
	2, 344, 345, 7, 54, 2, 2, 345, 346, 8, 16, 1, 2, 346, 354, 3, 2, 2, 2,
	347, 348, 5, 32, 17, 2, 348, 349, 7, 40, 2, 2, 349, 350, 5, 36, 19, 2,
	350, 351, 7, 35, 2, 2, 351, 352, 8, 16, 1, 2, 352, 354, 3, 2, 2, 2, 353,
	340, 3, 2, 2, 2, 353, 347, 3, 2, 2, 2, 354, 31, 3, 2, 2, 2, 355, 357, 5,
	34, 18, 2, 356, 355, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2, 358, 356, 3, 2,
	2, 2, 358, 359, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 361, 8, 17, 1, 2,
	361, 33, 3, 2, 2, 2, 362, 363, 5, 80, 41, 2, 363, 364, 7, 59, 2, 2, 364,
	365, 8, 18, 1, 2, 365, 370, 3, 2, 2, 2, 366, 367, 5, 80, 41, 2, 367, 368,
	8, 18, 1, 2, 368, 370, 3, 2, 2, 2, 369, 362, 3, 2, 2, 2, 369, 366, 3, 2,
	2, 2, 370, 35, 3, 2, 2, 2, 371, 372, 5, 8, 5, 2, 372, 373, 8, 19, 1, 2,
	373, 37, 3, 2, 2, 2, 374, 375, 7, 31, 2, 2, 375, 376, 7, 40, 2, 2, 376,
	377, 7, 53, 2, 2, 377, 378, 5, 4, 3, 2, 378, 379, 7, 54, 2, 2, 379, 380,
	8, 20, 1, 2, 380, 388, 3, 2, 2, 2, 381, 382, 7, 31, 2, 2, 382, 383, 7,
	40, 2, 2, 383, 384, 5, 36, 19, 2, 384, 385, 7, 35, 2, 2, 385, 386, 8, 20,
	1, 2, 386, 388, 3, 2, 2, 2, 387, 374, 3, 2, 2, 2, 387, 381, 3, 2, 2, 2,
	388, 39, 3, 2, 2, 2, 389, 391, 5, 38, 20, 2, 390, 389, 3, 2, 2, 2, 391,
	392, 3, 2, 2, 2, 392, 390, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2, 393, 394,
	3, 2, 2, 2, 394, 395, 8, 21, 1, 2, 395, 41, 3, 2, 2, 2, 396, 397, 7, 9,
	2, 2, 397, 398, 5, 80, 41, 2, 398, 399, 7, 53, 2, 2, 399, 400, 5, 44, 23,
	2, 400, 401, 5, 52, 27, 2, 401, 402, 7, 54, 2, 2, 402, 403, 8, 22, 1, 2,
	403, 412, 3, 2, 2, 2, 404, 405, 7, 9, 2, 2, 405, 406, 5, 80, 41, 2, 406,
	407, 7, 53, 2, 2, 407, 408, 5, 52, 27, 2, 408, 409, 7, 54, 2, 2, 409, 410,
	8, 22, 1, 2, 410, 412, 3, 2, 2, 2, 411, 396, 3, 2, 2, 2, 411, 404, 3, 2,
	2, 2, 412, 43, 3, 2, 2, 2, 413, 415, 5, 46, 24, 2, 414, 413, 3, 2, 2, 2,
	415, 416, 3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417,
	418, 3, 2, 2, 2, 418, 419, 8, 23, 1, 2, 419, 45, 3, 2, 2, 2, 420, 421,
	5, 48, 25, 2, 421, 422, 7, 40, 2, 2, 422, 423, 5, 80, 41, 2, 423, 424,
	7, 35, 2, 2, 424, 425, 8, 24, 1, 2, 425, 434, 3, 2, 2, 2, 426, 427, 5,
	48, 25, 2, 427, 428, 7, 40, 2, 2, 428, 429, 7, 53, 2, 2, 429, 430, 5, 80,
	41, 2, 430, 431, 7, 54, 2, 2, 431, 432, 8, 24, 1, 2, 432, 434, 3, 2, 2,
	2, 433, 420, 3, 2, 2, 2, 433, 426, 3, 2, 2, 2, 434, 47, 3, 2, 2, 2, 435,
	437, 5, 50, 26, 2, 436, 435, 3, 2, 2, 2, 437, 438, 3, 2, 2, 2, 438, 436,
	3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 440, 441, 8, 25,
	1, 2, 441, 49, 3, 2, 2, 2, 442, 443, 5, 80, 41, 2, 443, 444, 7, 59, 2,
	2, 444, 445, 8, 26, 1, 2, 445, 450, 3, 2, 2, 2, 446, 447, 5, 80, 41, 2,
	447, 448, 8, 26, 1, 2, 448, 450, 3, 2, 2, 2, 449, 442, 3, 2, 2, 2, 449,
	446, 3, 2, 2, 2, 450, 51, 3, 2, 2, 2, 451, 452, 7, 31, 2, 2, 452, 453,
	7, 40, 2, 2, 453, 454, 5, 80, 41, 2, 454, 455, 7, 35, 2, 2, 455, 456, 8,
	27, 1, 2, 456, 465, 3, 2, 2, 2, 457, 458, 7, 31, 2, 2, 458, 459, 7, 40,
	2, 2, 459, 460, 7, 53, 2, 2, 460, 461, 5, 80, 41, 2, 461, 462, 7, 54, 2,
	2, 462, 463, 8, 27, 1, 2, 463, 465, 3, 2, 2, 2, 464, 451, 3, 2, 2, 2, 464,
	457, 3, 2, 2, 2, 465, 53, 3, 2, 2, 2, 466, 467, 7, 10, 2, 2, 467, 468,
	5, 80, 41, 2, 468, 469, 7, 53, 2, 2, 469, 470, 5, 4, 3, 2, 470, 471, 7,
	54, 2, 2, 471, 472, 8, 28, 1, 2, 472, 55, 3, 2, 2, 2, 473, 474, 7, 12,
	2, 2, 474, 475, 7, 31, 2, 2, 475, 476, 7, 13, 2, 2, 476, 477, 5, 80, 41,
	2, 477, 478, 7, 32, 2, 2, 478, 479, 5, 80, 41, 2, 479, 480, 7, 53, 2, 2,
	480, 481, 5, 4, 3, 2, 481, 482, 7, 54, 2, 2, 482, 483, 8, 29, 1, 2, 483,
	494, 3, 2, 2, 2, 484, 485, 7, 12, 2, 2, 485, 486, 7, 31, 2, 2, 486, 487,
	7, 13, 2, 2, 487, 488, 5, 80, 41, 2, 488, 489, 7, 53, 2, 2, 489, 490, 5,
	4, 3, 2, 490, 491, 7, 54, 2, 2, 491, 492, 8, 29, 1, 2, 492, 494, 3, 2,
	2, 2, 493, 473, 3, 2, 2, 2, 493, 484, 3, 2, 2, 2, 494, 57, 3, 2, 2, 2,
	495, 496, 7, 11, 2, 2, 496, 497, 7, 53, 2, 2, 497, 498, 5, 4, 3, 2, 498,
	499, 7, 54, 2, 2, 499, 500, 8, 30, 1, 2, 500, 59, 3, 2, 2, 2, 501, 502,
	7, 11, 2, 2, 502, 503, 7, 53, 2, 2, 503, 504, 5, 4, 3, 2, 504, 505, 7,
	54, 2, 2, 505, 506, 8, 31, 1, 2, 506, 61, 3, 2, 2, 2, 507, 508, 7, 14,
	2, 2, 508, 514, 8, 32, 1, 2, 509, 510, 7, 14, 2, 2, 510, 511, 5, 80, 41,
	2, 511, 512, 8, 32, 1, 2, 512, 514, 3, 2, 2, 2, 513, 507, 3, 2, 2, 2, 513,
	509, 3, 2, 2, 2, 514, 63, 3, 2, 2, 2, 515, 516, 7, 15, 2, 2, 516, 517,
	8, 33, 1, 2, 517, 65, 3, 2, 2, 2, 518, 519, 7, 18, 2, 2, 519, 520, 7, 31,
	2, 2, 520, 521, 7, 51, 2, 2, 521, 522, 7, 52, 2, 2, 522, 523, 7, 53, 2,
	2, 523, 524, 5, 4, 3, 2, 524, 525, 7, 54, 2, 2, 525, 526, 8, 34, 1, 2,
	526, 550, 3, 2, 2, 2, 527, 528, 7, 18, 2, 2, 528, 529, 7, 31, 2, 2, 529,
	530, 7, 51, 2, 2, 530, 531, 5, 68, 35, 2, 531, 532, 7, 52, 2, 2, 532, 533,
	7, 53, 2, 2, 533, 534, 5, 4, 3, 2, 534, 535, 7, 54, 2, 2, 535, 536, 8,
	34, 1, 2, 536, 550, 3, 2, 2, 2, 537, 538, 7, 18, 2, 2, 538, 539, 7, 31,
	2, 2, 539, 540, 7, 51, 2, 2, 540, 541, 5, 68, 35, 2, 541, 542, 7, 52, 2,
	2, 542, 543, 7, 41, 2, 2, 543, 544, 5, 74, 38, 2, 544, 545, 7, 53, 2, 2,
	545, 546, 5, 4, 3, 2, 546, 547, 7, 54, 2, 2, 547, 548, 8, 34, 1, 2, 548,
	550, 3, 2, 2, 2, 549, 518, 3, 2, 2, 2, 549, 527, 3, 2, 2, 2, 549, 537,
	3, 2, 2, 2, 550, 67, 3, 2, 2, 2, 551, 553, 5, 70, 36, 2, 552, 551, 3, 2,
	2, 2, 553, 556, 3, 2, 2, 2, 554, 552, 3, 2, 2, 2, 554, 555, 3, 2, 2, 2,
	555, 557, 3, 2, 2, 2, 556, 554, 3, 2, 2, 2, 557, 558, 8, 35, 1, 2, 558,
	69, 3, 2, 2, 2, 559, 560, 7, 31, 2, 2, 560, 561, 7, 36, 2, 2, 561, 562,
	5, 74, 38, 2, 562, 563, 7, 35, 2, 2, 563, 564, 8, 36, 1, 2, 564, 571, 3,
	2, 2, 2, 565, 566, 7, 31, 2, 2, 566, 567, 7, 36, 2, 2, 567, 568, 5, 74,
	38, 2, 568, 569, 8, 36, 1, 2, 569, 571, 3, 2, 2, 2, 570, 559, 3, 2, 2,
	2, 570, 565, 3, 2, 2, 2, 571, 71, 3, 2, 2, 2, 572, 573, 7, 31, 2, 2, 573,
	574, 7, 51, 2, 2, 574, 575, 7, 52, 2, 2, 575, 583, 8, 37, 1, 2, 576, 577,
	7, 31, 2, 2, 577, 578, 7, 51, 2, 2, 578, 579, 5, 76, 39, 2, 579, 580, 7,
	52, 2, 2, 580, 581, 8, 37, 1, 2, 581, 583, 3, 2, 2, 2, 582, 572, 3, 2,
	2, 2, 582, 576, 3, 2, 2, 2, 583, 73, 3, 2, 2, 2, 584, 585, 7, 20, 2, 2,
	585, 597, 8, 38, 1, 2, 586, 587, 7, 21, 2, 2, 587, 597, 8, 38, 1, 2, 588,
	589, 7, 22, 2, 2, 589, 597, 8, 38, 1, 2, 590, 591, 7, 25, 2, 2, 591, 597,
	8, 38, 1, 2, 592, 593, 7, 23, 2, 2, 593, 597, 8, 38, 1, 2, 594, 595, 7,
	24, 2, 2, 595, 597, 8, 38, 1, 2, 596, 584, 3, 2, 2, 2, 596, 586, 3, 2,
	2, 2, 596, 588, 3, 2, 2, 2, 596, 590, 3, 2, 2, 2, 596, 592, 3, 2, 2, 2,
	596, 594, 3, 2, 2, 2, 597, 75, 3, 2, 2, 2, 598, 600, 5, 78, 40, 2, 599,
	598, 3, 2, 2, 2, 600, 601, 3, 2, 2, 2, 601, 599, 3, 2, 2, 2, 601, 602,
	3, 2, 2, 2, 602, 603, 3, 2, 2, 2, 603, 604, 8, 39, 1, 2, 604, 77, 3, 2,
	2, 2, 605, 606, 5, 80, 41, 2, 606, 607, 7, 35, 2, 2, 607, 608, 8, 40, 1,
	2, 608, 613, 3, 2, 2, 2, 609, 610, 5, 80, 41, 2, 610, 611, 8, 40, 1, 2,
	611, 613, 3, 2, 2, 2, 612, 605, 3, 2, 2, 2, 612, 609, 3, 2, 2, 2, 613,
	79, 3, 2, 2, 2, 614, 615, 5, 82, 42, 2, 615, 616, 8, 41, 1, 2, 616, 627,
	3, 2, 2, 2, 617, 618, 5, 84, 43, 2, 618, 619, 8, 41, 1, 2, 619, 627, 3,
	2, 2, 2, 620, 621, 5, 86, 44, 2, 621, 622, 8, 41, 1, 2, 622, 627, 3, 2,
	2, 2, 623, 624, 5, 94, 48, 2, 624, 625, 8, 41, 1, 2, 625, 627, 3, 2, 2,
	2, 626, 614, 3, 2, 2, 2, 626, 617, 3, 2, 2, 2, 626, 620, 3, 2, 2, 2, 626,
	623, 3, 2, 2, 2, 627, 81, 3, 2, 2, 2, 628, 629, 8, 42, 1, 2, 629, 630,
	7, 60, 2, 2, 630, 631, 5, 82, 42, 5, 631, 632, 8, 42, 1, 2, 632, 637, 3,
	2, 2, 2, 633, 634, 5, 84, 43, 2, 634, 635, 8, 42, 1, 2, 635, 637, 3, 2,
	2, 2, 636, 628, 3, 2, 2, 2, 636, 633, 3, 2, 2, 2, 637, 645, 3, 2, 2, 2,
	638, 639, 12, 4, 2, 2, 639, 640, 9, 2, 2, 2, 640, 641, 5, 82, 42, 5, 641,
	642, 8, 42, 1, 2, 642, 644, 3, 2, 2, 2, 643, 638, 3, 2, 2, 2, 644, 647,
	3, 2, 2, 2, 645, 643, 3, 2, 2, 2, 645, 646, 3, 2, 2, 2, 646, 83, 3, 2,
	2, 2, 647, 645, 3, 2, 2, 2, 648, 649, 8, 43, 1, 2, 649, 650, 5, 86, 44,
	2, 650, 651, 8, 43, 1, 2, 651, 659, 3, 2, 2, 2, 652, 653, 12, 4, 2, 2,
	653, 654, 9, 3, 2, 2, 654, 655, 5, 84, 43, 5, 655, 656, 8, 43, 1, 2, 656,
	658, 3, 2, 2, 2, 657, 652, 3, 2, 2, 2, 658, 661, 3, 2, 2, 2, 659, 657,
	3, 2, 2, 2, 659, 660, 3, 2, 2, 2, 660, 85, 3, 2, 2, 2, 661, 659, 3, 2,
	2, 2, 662, 663, 8, 44, 1, 2, 663, 664, 7, 50, 2, 2, 664, 665, 5, 86, 44,
	7, 665, 666, 8, 44, 1, 2, 666, 676, 3, 2, 2, 2, 667, 668, 5, 88, 45, 2,
	668, 669, 8, 44, 1, 2, 669, 676, 3, 2, 2, 2, 670, 671, 7, 51, 2, 2, 671,
	672, 5, 80, 41, 2, 672, 673, 7, 52, 2, 2, 673, 674, 8, 44, 1, 2, 674, 676,
	3, 2, 2, 2, 675, 662, 3, 2, 2, 2, 675, 667, 3, 2, 2, 2, 675, 670, 3, 2,
	2, 2, 676, 689, 3, 2, 2, 2, 677, 678, 12, 6, 2, 2, 678, 679, 9, 4, 2, 2,
	679, 680, 5, 86, 44, 7, 680, 681, 8, 44, 1, 2, 681, 688, 3, 2, 2, 2, 682,
	683, 12, 5, 2, 2, 683, 684, 9, 5, 2, 2, 684, 685, 5, 86, 44, 6, 685, 686,
	8, 44, 1, 2, 686, 688, 3, 2, 2, 2, 687, 677, 3, 2, 2, 2, 687, 682, 3, 2,
	2, 2, 688, 691, 3, 2, 2, 2, 689, 687, 3, 2, 2, 2, 689, 690, 3, 2, 2, 2,
	690, 87, 3, 2, 2, 2, 691, 689, 3, 2, 2, 2, 692, 693, 5, 90, 46, 2, 693,
	694, 8, 45, 1, 2, 694, 89, 3, 2, 2, 2, 695, 696, 7, 26, 2, 2, 696, 720,
	8, 46, 1, 2, 697, 698, 7, 27, 2, 2, 698, 720, 8, 46, 1, 2, 699, 700, 7,
	29, 2, 2, 700, 720, 8, 46, 1, 2, 701, 702, 7, 30, 2, 2, 702, 720, 8, 46,
	1, 2, 703, 704, 7, 28, 2, 2, 704, 720, 8, 46, 1, 2, 705, 706, 7, 31, 2,
	2, 706, 720, 8, 46, 1, 2, 707, 708, 5, 92, 47, 2, 708, 709, 8, 46, 1, 2,
	709, 720, 3, 2, 2, 2, 710, 711, 5, 22, 12, 2, 711, 712, 8, 46, 1, 2, 712,
	720, 3, 2, 2, 2, 713, 714, 5, 42, 22, 2, 714, 715, 8, 46, 1, 2, 715, 720,
	3, 2, 2, 2, 716, 717, 5, 60, 31, 2, 717, 718, 8, 46, 1, 2, 718, 720, 3,
	2, 2, 2, 719, 695, 3, 2, 2, 2, 719, 697, 3, 2, 2, 2, 719, 699, 3, 2, 2,
	2, 719, 701, 3, 2, 2, 2, 719, 703, 3, 2, 2, 2, 719, 705, 3, 2, 2, 2, 719,
	707, 3, 2, 2, 2, 719, 710, 3, 2, 2, 2, 719, 713, 3, 2, 2, 2, 719, 716,
	3, 2, 2, 2, 720, 91, 3, 2, 2, 2, 721, 722, 7, 26, 2, 2, 722, 723, 7, 19,
	2, 2, 723, 724, 7, 20, 2, 2, 724, 746, 8, 47, 1, 2, 725, 726, 7, 26, 2,
	2, 726, 727, 7, 19, 2, 2, 727, 728, 7, 21, 2, 2, 728, 746, 8, 47, 1, 2,
	729, 730, 7, 27, 2, 2, 730, 731, 7, 19, 2, 2, 731, 732, 7, 20, 2, 2, 732,
	746, 8, 47, 1, 2, 733, 734, 7, 27, 2, 2, 734, 735, 7, 19, 2, 2, 735, 736,
	7, 21, 2, 2, 736, 746, 8, 47, 1, 2, 737, 738, 7, 30, 2, 2, 738, 739, 7,
	19, 2, 2, 739, 740, 7, 20, 2, 2, 740, 746, 8, 47, 1, 2, 741, 742, 7, 30,
	2, 2, 742, 743, 7, 19, 2, 2, 743, 744, 7, 21, 2, 2, 744, 746, 8, 47, 1,
	2, 745, 721, 3, 2, 2, 2, 745, 725, 3, 2, 2, 2, 745, 729, 3, 2, 2, 2, 745,
	733, 3, 2, 2, 2, 745, 737, 3, 2, 2, 2, 745, 741, 3, 2, 2, 2, 746, 93, 3,
	2, 2, 2, 747, 748, 5, 86, 44, 2, 748, 749, 7, 19, 2, 2, 749, 750, 5, 96,
	49, 2, 750, 751, 8, 48, 1, 2, 751, 95, 3, 2, 2, 2, 752, 753, 7, 20, 2,
	2, 753, 757, 8, 49, 1, 2, 754, 755, 7, 21, 2, 2, 755, 757, 8, 49, 1, 2,
	756, 752, 3, 2, 2, 2, 756, 754, 3, 2, 2, 2, 757, 97, 3, 2, 2, 2, 43, 104,
	111, 156, 174, 234, 269, 274, 306, 311, 331, 336, 353, 358, 369, 387, 392,
	411, 416, 433, 438, 449, 464, 493, 513, 549, 554, 570, 582, 596, 601, 612,
	626, 636, 645, 659, 675, 687, 689, 719, 745, 756,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'print!'", "'println!'", "'let'", "'mut'", "'if'", "'else'", "'match'",
	"'while'", "'loop'", "'for'", "'in'", "'break'", "'continue'", "'return'",
	"'main'", "'fn'", "'as'", "'i64'", "'f64'", "'String'", "'bool'", "'char'",
	"'&str'", "", "", "", "", "", "", "'..'", "'.'", "';'", "','", "':'", "'='",
	"'=='", "'>='", "'=>'", "'->'", "'<='", "'!='", "'>'", "'<'", "'*'", "'/'",
	"'%'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'&&'",
	"'||'", "'|'", "'!'",
}
var symbolicNames = []string{
	"", "R_PRINT", "R_PRINTLN", "R_LET", "R_MUT", "R_IF", "R_ELSE", "R_MATCH",
	"R_WHILE", "R_LOOP", "R_FOR", "R_IN", "R_BREAK", "R_CONTINUE", "R_RETURN",
	"R_MAIN", "R_FUNCTION", "R_AS", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL",
	"R_CHAR", "R_STR", "NUMBER", "DOUBLE", "CHAR", "STRING", "BOOLEAN", "ID",
	"TK_DOBLEPUNTO", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA", "TK_DOSPUNTOS",
	"TK_IGUAL", "TK_IGUALIGUAL", "TK_MAYORIGUAL", "TK_IGUALMAYOR", "TK_MENOSMAYOR",
	"TK_MENORIGUAL", "TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV",
	"TK_MODULO", "TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC",
	"TK_CORA", "TK_CORC", "TK_AND", "TK_OR", "TK_BARRA", "TK_NOT", "WHITESPACE",
	"TK_MULTI", "TK_LINE",
}

var ruleNames = []string{
	"start", "instrucciones", "end_instr", "instruccion", "instr_println",
	"instr_main", "instr_declaracion", "instr_asignacion", "instr_if", "instr_else_if",
	"instr_ternario", "instr_else_if_ternario", "instr_match", "list_case",
	"instr_case", "list_expre_case", "block_case", "block_instr_match", "instr_default",
	"block_default", "instr_match_ter", "list_case_ternario", "instr_case_ter",
	"list_expre_case_ter", "block_case_ter", "instr_default_ter", "instr_while",
	"instr_for_in", "instr_loop", "instr_loop_ternario", "instr_break", "instr_continue",
	"instr_func", "list_function_parameters", "block_parameters_fn", "instr_llamada",
	"instr_tipo", "list_expression", "block_list_expression", "expressions",
	"expre_log", "expre_rel", "expre_arit", "expre_valor", "primitivo", "primitivo_casteo",
	"expre_casteo", "type_casteo",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Chems struct {
	*antlr.BaseParser
}

func NewChems(input antlr.TokenStream) *Chems {
	this := new(Chems)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Chems.g4"

	return this
}

// Chems tokens.
const (
	ChemsEOF           = antlr.TokenEOF
	ChemsR_PRINT       = 1
	ChemsR_PRINTLN     = 2
	ChemsR_LET         = 3
	ChemsR_MUT         = 4
	ChemsR_IF          = 5
	ChemsR_ELSE        = 6
	ChemsR_MATCH       = 7
	ChemsR_WHILE       = 8
	ChemsR_LOOP        = 9
	ChemsR_FOR         = 10
	ChemsR_IN          = 11
	ChemsR_BREAK       = 12
	ChemsR_CONTINUE    = 13
	ChemsR_RETURN      = 14
	ChemsR_MAIN        = 15
	ChemsR_FUNCTION    = 16
	ChemsR_AS          = 17
	ChemsR_INT         = 18
	ChemsR_FLOAT       = 19
	ChemsR_STRING      = 20
	ChemsR_BOOL        = 21
	ChemsR_CHAR        = 22
	ChemsR_STR         = 23
	ChemsNUMBER        = 24
	ChemsDOUBLE        = 25
	ChemsCHAR          = 26
	ChemsSTRING        = 27
	ChemsBOOLEAN       = 28
	ChemsID            = 29
	ChemsTK_DOBLEPUNTO = 30
	ChemsTK_PUNTO      = 31
	ChemsTK_PUNTOCOMA  = 32
	ChemsTK_COMA       = 33
	ChemsTK_DOSPUNTOS  = 34
	ChemsTK_IGUAL      = 35
	ChemsTK_IGUALIGUAL = 36
	ChemsTK_MAYORIGUAL = 37
	ChemsTK_IGUALMAYOR = 38
	ChemsTK_MENOSMAYOR = 39
	ChemsTK_MENORIGUAL = 40
	ChemsTK_DIFIGUAL   = 41
	ChemsTK_MAYOR      = 42
	ChemsTK_MENOR      = 43
	ChemsTK_MULT       = 44
	ChemsTK_DIV        = 45
	ChemsTK_MODULO     = 46
	ChemsTK_MAS        = 47
	ChemsTK_MENOS      = 48
	ChemsTK_PARA       = 49
	ChemsTK_PARC       = 50
	ChemsTK_LLAVEA     = 51
	ChemsTK_LLAVEC     = 52
	ChemsTK_CORA       = 53
	ChemsTK_CORC       = 54
	ChemsTK_AND        = 55
	ChemsTK_OR         = 56
	ChemsTK_BARRA      = 57
	ChemsTK_NOT        = 58
	ChemsWHITESPACE    = 59
	ChemsTK_MULTI      = 60
	ChemsTK_LINE       = 61
)

// Chems rules.
const (
	ChemsRULE_start                    = 0
	ChemsRULE_instrucciones            = 1
	ChemsRULE_end_instr                = 2
	ChemsRULE_instruccion              = 3
	ChemsRULE_instr_println            = 4
	ChemsRULE_instr_main               = 5
	ChemsRULE_instr_declaracion        = 6
	ChemsRULE_instr_asignacion         = 7
	ChemsRULE_instr_if                 = 8
	ChemsRULE_instr_else_if            = 9
	ChemsRULE_instr_ternario           = 10
	ChemsRULE_instr_else_if_ternario   = 11
	ChemsRULE_instr_match              = 12
	ChemsRULE_list_case                = 13
	ChemsRULE_instr_case               = 14
	ChemsRULE_list_expre_case          = 15
	ChemsRULE_block_case               = 16
	ChemsRULE_block_instr_match        = 17
	ChemsRULE_instr_default            = 18
	ChemsRULE_block_default            = 19
	ChemsRULE_instr_match_ter          = 20
	ChemsRULE_list_case_ternario       = 21
	ChemsRULE_instr_case_ter           = 22
	ChemsRULE_list_expre_case_ter      = 23
	ChemsRULE_block_case_ter           = 24
	ChemsRULE_instr_default_ter        = 25
	ChemsRULE_instr_while              = 26
	ChemsRULE_instr_for_in             = 27
	ChemsRULE_instr_loop               = 28
	ChemsRULE_instr_loop_ternario      = 29
	ChemsRULE_instr_break              = 30
	ChemsRULE_instr_continue           = 31
	ChemsRULE_instr_func               = 32
	ChemsRULE_list_function_parameters = 33
	ChemsRULE_block_parameters_fn      = 34
	ChemsRULE_instr_llamada            = 35
	ChemsRULE_instr_tipo               = 36
	ChemsRULE_list_expression          = 37
	ChemsRULE_block_list_expression    = 38
	ChemsRULE_expressions              = 39
	ChemsRULE_expre_log                = 40
	ChemsRULE_expre_rel                = 41
	ChemsRULE_expre_arit               = 42
	ChemsRULE_expre_valor              = 43
	ChemsRULE_primitivo                = 44
	ChemsRULE_primitivo_casteo         = 45
	ChemsRULE_expre_casteo             = 46
	ChemsRULE_type_casteo              = 47
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *StartContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *StartContext) GetLista() *arrayList.List { return s.lista }

func (s *StartContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Chems) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChemsRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	localctx.(*StartContext).lista = localctx.(*StartContext).Get_instrucciones().GetL()

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetList returns the list rule context list.
	GetList() []IInstruccionContext

	// SetList sets the list rule context list.
	SetList([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruccion IInstruccionContext
	list         []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetList() []IInstruccionContext { return s.list }

func (s *InstruccionesContext) SetList(v []IInstruccionContext) { s.list = v }

func (s *InstruccionesContext) GetL() *arrayList.List { return s.l }

func (s *InstruccionesContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Chems) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChemsRULE_instrucciones)

	localctx.(*InstruccionesContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_PRINTLN)|(1<<ChemsR_LET)|(1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_WHILE)|(1<<ChemsR_LOOP)|(1<<ChemsR_FOR)|(1<<ChemsR_BREAK)|(1<<ChemsR_CONTINUE)|(1<<ChemsR_FUNCTION)|(1<<ChemsID))) != 0) {
		{
			p.SetState(99)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).list = append(localctx.(*InstruccionesContext).list, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetList()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IEnd_instrContext is an interface to support dynamic dispatch.
type IEnd_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v attribute.
	GetV() int

	// SetV sets the v attribute.
	SetV(int)

	// IsEnd_instrContext differentiates from other interfaces.
	IsEnd_instrContext()
}

type End_instrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      int
}

func NewEmptyEnd_instrContext() *End_instrContext {
	var p = new(End_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_end_instr
	return p
}

func (*End_instrContext) IsEnd_instrContext() {}

func NewEnd_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *End_instrContext {
	var p = new(End_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_end_instr

	return p
}

func (s *End_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *End_instrContext) GetV() int { return s.v }

func (s *End_instrContext) SetV(v int) { s.v = v }

func (s *End_instrContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *End_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *End_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *End_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterEnd_instr(s)
	}
}

func (s *End_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitEnd_instr(s)
	}
}

func (p *Chems) End_instr() (localctx IEnd_instrContext) {
	localctx = NewEnd_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChemsRULE_end_instr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_PUNTOCOMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*End_instrContext).v = 1

	case ChemsEOF, ChemsR_PRINTLN, ChemsR_LET, ChemsR_IF, ChemsR_MATCH, ChemsR_WHILE, ChemsR_LOOP, ChemsR_FOR, ChemsR_BREAK, ChemsR_CONTINUE, ChemsR_FUNCTION, ChemsID, ChemsTK_COMA, ChemsTK_LLAVEC:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*End_instrContext).v = 0

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_println returns the _instr_println rule contexts.
	Get_instr_println() IInstr_printlnContext

	// Get_instr_main returns the _instr_main rule contexts.
	Get_instr_main() IInstr_mainContext

	// Get_instr_declaracion returns the _instr_declaracion rule contexts.
	Get_instr_declaracion() IInstr_declaracionContext

	// Get_instr_asignacion returns the _instr_asignacion rule contexts.
	Get_instr_asignacion() IInstr_asignacionContext

	// Get_instr_if returns the _instr_if rule contexts.
	Get_instr_if() IInstr_ifContext

	// Get_instr_match returns the _instr_match rule contexts.
	Get_instr_match() IInstr_matchContext

	// Get_instr_while returns the _instr_while rule contexts.
	Get_instr_while() IInstr_whileContext

	// Get_instr_for_in returns the _instr_for_in rule contexts.
	Get_instr_for_in() IInstr_for_inContext

	// Get_instr_loop returns the _instr_loop rule contexts.
	Get_instr_loop() IInstr_loopContext

	// Get_instr_break returns the _instr_break rule contexts.
	Get_instr_break() IInstr_breakContext

	// Get_instr_continue returns the _instr_continue rule contexts.
	Get_instr_continue() IInstr_continueContext

	// Get_instr_func returns the _instr_func rule contexts.
	Get_instr_func() IInstr_funcContext

	// Get_instr_llamada returns the _instr_llamada rule contexts.
	Get_instr_llamada() IInstr_llamadaContext

	// Set_instr_println sets the _instr_println rule contexts.
	Set_instr_println(IInstr_printlnContext)

	// Set_instr_main sets the _instr_main rule contexts.
	Set_instr_main(IInstr_mainContext)

	// Set_instr_declaracion sets the _instr_declaracion rule contexts.
	Set_instr_declaracion(IInstr_declaracionContext)

	// Set_instr_asignacion sets the _instr_asignacion rule contexts.
	Set_instr_asignacion(IInstr_asignacionContext)

	// Set_instr_if sets the _instr_if rule contexts.
	Set_instr_if(IInstr_ifContext)

	// Set_instr_match sets the _instr_match rule contexts.
	Set_instr_match(IInstr_matchContext)

	// Set_instr_while sets the _instr_while rule contexts.
	Set_instr_while(IInstr_whileContext)

	// Set_instr_for_in sets the _instr_for_in rule contexts.
	Set_instr_for_in(IInstr_for_inContext)

	// Set_instr_loop sets the _instr_loop rule contexts.
	Set_instr_loop(IInstr_loopContext)

	// Set_instr_break sets the _instr_break rule contexts.
	Set_instr_break(IInstr_breakContext)

	// Set_instr_continue sets the _instr_continue rule contexts.
	Set_instr_continue(IInstr_continueContext)

	// Set_instr_func sets the _instr_func rule contexts.
	Set_instr_func(IInstr_funcContext)

	// Set_instr_llamada sets the _instr_llamada rule contexts.
	Set_instr_llamada(IInstr_llamadaContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	instr              interfaces.Instruction
	_instr_println     IInstr_printlnContext
	_instr_main        IInstr_mainContext
	_instr_declaracion IInstr_declaracionContext
	_instr_asignacion  IInstr_asignacionContext
	_instr_if          IInstr_ifContext
	_instr_match       IInstr_matchContext
	_instr_while       IInstr_whileContext
	_instr_for_in      IInstr_for_inContext
	_instr_loop        IInstr_loopContext
	_instr_break       IInstr_breakContext
	_instr_continue    IInstr_continueContext
	_instr_func        IInstr_funcContext
	_instr_llamada     IInstr_llamadaContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_instr_println() IInstr_printlnContext { return s._instr_println }

func (s *InstruccionContext) Get_instr_main() IInstr_mainContext { return s._instr_main }

func (s *InstruccionContext) Get_instr_declaracion() IInstr_declaracionContext {
	return s._instr_declaracion
}

func (s *InstruccionContext) Get_instr_asignacion() IInstr_asignacionContext {
	return s._instr_asignacion
}

func (s *InstruccionContext) Get_instr_if() IInstr_ifContext { return s._instr_if }

func (s *InstruccionContext) Get_instr_match() IInstr_matchContext { return s._instr_match }

func (s *InstruccionContext) Get_instr_while() IInstr_whileContext { return s._instr_while }

func (s *InstruccionContext) Get_instr_for_in() IInstr_for_inContext { return s._instr_for_in }

func (s *InstruccionContext) Get_instr_loop() IInstr_loopContext { return s._instr_loop }

func (s *InstruccionContext) Get_instr_break() IInstr_breakContext { return s._instr_break }

func (s *InstruccionContext) Get_instr_continue() IInstr_continueContext { return s._instr_continue }

func (s *InstruccionContext) Get_instr_func() IInstr_funcContext { return s._instr_func }

func (s *InstruccionContext) Get_instr_llamada() IInstr_llamadaContext { return s._instr_llamada }

func (s *InstruccionContext) Set_instr_println(v IInstr_printlnContext) { s._instr_println = v }

func (s *InstruccionContext) Set_instr_main(v IInstr_mainContext) { s._instr_main = v }

func (s *InstruccionContext) Set_instr_declaracion(v IInstr_declaracionContext) {
	s._instr_declaracion = v
}

func (s *InstruccionContext) Set_instr_asignacion(v IInstr_asignacionContext) { s._instr_asignacion = v }

func (s *InstruccionContext) Set_instr_if(v IInstr_ifContext) { s._instr_if = v }

func (s *InstruccionContext) Set_instr_match(v IInstr_matchContext) { s._instr_match = v }

func (s *InstruccionContext) Set_instr_while(v IInstr_whileContext) { s._instr_while = v }

func (s *InstruccionContext) Set_instr_for_in(v IInstr_for_inContext) { s._instr_for_in = v }

func (s *InstruccionContext) Set_instr_loop(v IInstr_loopContext) { s._instr_loop = v }

func (s *InstruccionContext) Set_instr_break(v IInstr_breakContext) { s._instr_break = v }

func (s *InstruccionContext) Set_instr_continue(v IInstr_continueContext) { s._instr_continue = v }

func (s *InstruccionContext) Set_instr_func(v IInstr_funcContext) { s._instr_func = v }

func (s *InstruccionContext) Set_instr_llamada(v IInstr_llamadaContext) { s._instr_llamada = v }

func (s *InstruccionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *InstruccionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *InstruccionContext) Instr_println() IInstr_printlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_printlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_printlnContext)
}

func (s *InstruccionContext) End_instr() IEnd_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnd_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnd_instrContext)
}

func (s *InstruccionContext) Instr_main() IInstr_mainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_mainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_mainContext)
}

func (s *InstruccionContext) Instr_declaracion() IInstr_declaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_declaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_declaracionContext)
}

func (s *InstruccionContext) Instr_asignacion() IInstr_asignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_asignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_asignacionContext)
}

func (s *InstruccionContext) Instr_if() IInstr_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_ifContext)
}

func (s *InstruccionContext) Instr_match() IInstr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_matchContext)
}

func (s *InstruccionContext) Instr_while() IInstr_whileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_whileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_whileContext)
}

func (s *InstruccionContext) Instr_for_in() IInstr_for_inContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_for_inContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_for_inContext)
}

func (s *InstruccionContext) Instr_loop() IInstr_loopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_loopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_loopContext)
}

func (s *InstruccionContext) Instr_break() IInstr_breakContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_breakContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_breakContext)
}

func (s *InstruccionContext) Instr_continue() IInstr_continueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_continueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_continueContext)
}

func (s *InstruccionContext) Instr_func() IInstr_funcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_funcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_funcContext)
}

func (s *InstruccionContext) Instr_llamada() IInstr_llamadaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_llamadaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_llamadaContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Chems) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChemsRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)

			var _x = p.Instr_println()

			localctx.(*InstruccionContext)._instr_println = _x
		}
		{
			p.SetState(112)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_println().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)

			var _x = p.Instr_main()

			localctx.(*InstruccionContext)._instr_main = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_main().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)

			var _x = p.Instr_declaracion()

			localctx.(*InstruccionContext)._instr_declaracion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_declaracion().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(121)

			var _x = p.Instr_asignacion()

			localctx.(*InstruccionContext)._instr_asignacion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_asignacion().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(124)

			var _x = p.Instr_if()

			localctx.(*InstruccionContext)._instr_if = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_if().GetInstr()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(127)

			var _x = p.Instr_match()

			localctx.(*InstruccionContext)._instr_match = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_match().GetInstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(130)

			var _x = p.Instr_while()

			localctx.(*InstruccionContext)._instr_while = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_while().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(133)

			var _x = p.Instr_for_in()

			localctx.(*InstruccionContext)._instr_for_in = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_for_in().GetInstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(136)

			var _x = p.Instr_loop()

			localctx.(*InstruccionContext)._instr_loop = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_loop().GetInstr()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(139)

			var _x = p.Instr_break()

			localctx.(*InstruccionContext)._instr_break = _x
		}
		{
			p.SetState(140)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_break().GetInstr()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(143)

			var _x = p.Instr_continue()

			localctx.(*InstruccionContext)._instr_continue = _x
		}
		{
			p.SetState(144)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_continue().GetInstr()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(147)

			var _x = p.Instr_func()

			localctx.(*InstruccionContext)._instr_func = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_func().GetInstr()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(150)

			var _x = p.Instr_llamada()

			localctx.(*InstruccionContext)._instr_llamada = _x
		}
		{
			p.SetState(151)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_llamada().GetInstr()

	}

	return localctx
}

// IInstr_printlnContext is an interface to support dynamic dispatch.
type IInstr_printlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_PRINTLN returns the _R_PRINTLN token.
	Get_R_PRINTLN() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_R_PRINTLN sets the _R_PRINTLN token.
	Set_R_PRINTLN(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Get_list_expression returns the _list_expression rule contexts.
	Get_list_expression() IList_expressionContext

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// Set_list_expression sets the _list_expression rule contexts.
	Set_list_expression(IList_expressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_printlnContext differentiates from other interfaces.
	IsInstr_printlnContext()
}

type Instr_printlnContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            interfaces.Instruction
	_R_PRINTLN       antlr.Token
	_primitivo       IPrimitivoContext
	_STRING          antlr.Token
	_list_expression IList_expressionContext
}

func NewEmptyInstr_printlnContext() *Instr_printlnContext {
	var p = new(Instr_printlnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_println
	return p
}

func (*Instr_printlnContext) IsInstr_printlnContext() {}

func NewInstr_printlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_printlnContext {
	var p = new(Instr_printlnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_println

	return p
}

func (s *Instr_printlnContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_printlnContext) Get_R_PRINTLN() antlr.Token { return s._R_PRINTLN }

func (s *Instr_printlnContext) Get_STRING() antlr.Token { return s._STRING }

func (s *Instr_printlnContext) Set_R_PRINTLN(v antlr.Token) { s._R_PRINTLN = v }

func (s *Instr_printlnContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *Instr_printlnContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Instr_printlnContext) Get_list_expression() IList_expressionContext {
	return s._list_expression
}

func (s *Instr_printlnContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Instr_printlnContext) Set_list_expression(v IList_expressionContext) { s._list_expression = v }

func (s *Instr_printlnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_printlnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_printlnContext) R_PRINTLN() antlr.TerminalNode {
	return s.GetToken(ChemsR_PRINTLN, 0)
}

func (s *Instr_printlnContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_printlnContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Instr_printlnContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_printlnContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_printlnContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *Instr_printlnContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_printlnContext) List_expression() IList_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Instr_printlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_printlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_printlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_println(s)
	}
}

func (s *Instr_printlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_println(s)
	}
}

func (p *Chems) Instr_println() (localctx IInstr_printlnContext) {
	localctx = NewInstr_printlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChemsRULE_instr_println)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)

			var _m = p.Match(ChemsR_PRINTLN)

			localctx.(*Instr_printlnContext)._R_PRINTLN = _m
		}
		{
			p.SetState(157)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(158)

			var _x = p.Primitivo()

			localctx.(*Instr_printlnContext)._primitivo = _x
		}
		{
			p.SetState(159)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(160)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.NewPrintln(nil, localctx.(*Instr_printlnContext).Get_primitivo().GetP(), "-1", (func() int {
			if localctx.(*Instr_printlnContext).Get_R_PRINTLN() == nil {
				return 0
			} else {
				return localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetLine()
			}
		}()), localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)

			var _m = p.Match(ChemsR_PRINTLN)

			localctx.(*Instr_printlnContext)._R_PRINTLN = _m
		}
		{
			p.SetState(164)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(165)

			var _m = p.Match(ChemsSTRING)

			localctx.(*Instr_printlnContext)._STRING = _m
		}
		{
			p.SetState(166)
			p.Match(ChemsTK_COMA)
		}
		{
			p.SetState(167)

			var _x = p.List_expression()

			localctx.(*Instr_printlnContext)._list_expression = _x
		}
		{
			p.SetState(168)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(169)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.NewPrintln(localctx.(*Instr_printlnContext).Get_list_expression().GetL(), nil, (func() string {
			if localctx.(*Instr_printlnContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*Instr_printlnContext).Get_STRING().GetText()
			}
		}())[1:len((func() string {
			if localctx.(*Instr_printlnContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*Instr_printlnContext).Get_STRING().GetText()
			}
		}()))-1], (func() int {
			if localctx.(*Instr_printlnContext).Get_R_PRINTLN() == nil {
				return 0
			} else {
				return localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetLine()
			}
		}()), localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn())

	}

	return localctx
}

// IInstr_mainContext is an interface to support dynamic dispatch.
type IInstr_mainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_MAIN returns the _R_MAIN token.
	Get_R_MAIN() antlr.Token

	// Set_R_MAIN sets the _R_MAIN token.
	Set_R_MAIN(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_mainContext differentiates from other interfaces.
	IsInstr_mainContext()
}

type Instr_mainContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_MAIN        antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_mainContext() *Instr_mainContext {
	var p = new(Instr_mainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_main
	return p
}

func (*Instr_mainContext) IsInstr_mainContext() {}

func NewInstr_mainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_mainContext {
	var p = new(Instr_mainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_main

	return p
}

func (s *Instr_mainContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_mainContext) Get_R_MAIN() antlr.Token { return s._R_MAIN }

func (s *Instr_mainContext) Set_R_MAIN(v antlr.Token) { s._R_MAIN = v }

func (s *Instr_mainContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_mainContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_mainContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_mainContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_mainContext) R_FUNCTION() antlr.TerminalNode {
	return s.GetToken(ChemsR_FUNCTION, 0)
}

func (s *Instr_mainContext) R_MAIN() antlr.TerminalNode {
	return s.GetToken(ChemsR_MAIN, 0)
}

func (s *Instr_mainContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_mainContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_mainContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_mainContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_mainContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_mainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_mainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_mainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_main(s)
	}
}

func (s *Instr_mainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_main(s)
	}
}

func (p *Chems) Instr_main() (localctx IInstr_mainContext) {
	localctx = NewInstr_mainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChemsRULE_instr_main)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(ChemsR_FUNCTION)
	}
	{
		p.SetState(175)

		var _m = p.Match(ChemsR_MAIN)

		localctx.(*Instr_mainContext)._R_MAIN = _m
	}
	{
		p.SetState(176)
		p.Match(ChemsTK_PARA)
	}
	{
		p.SetState(177)
		p.Match(ChemsTK_PARC)
	}
	{
		p.SetState(178)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(179)

		var _x = p.Instrucciones()

		localctx.(*Instr_mainContext)._instrucciones = _x
	}
	{
		p.SetState(180)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_mainContext).instr = instruction.NewMain(localctx.(*Instr_mainContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instr_mainContext).Get_R_MAIN() == nil {
			return 0
		} else {
			return localctx.(*Instr_mainContext).Get_R_MAIN().GetLine()
		}
	}()), localctx.(*Instr_mainContext).Get_R_MAIN().GetColumn())

	return localctx
}

// IInstr_declaracionContext is an interface to support dynamic dispatch.
type IInstr_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_LET returns the _R_LET token.
	Get_R_LET() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_R_LET sets the _R_LET token.
	Set_R_LET(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Get_instr_tipo returns the _instr_tipo rule contexts.
	Get_instr_tipo() IInstr_tipoContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// Set_instr_tipo sets the _instr_tipo rule contexts.
	Set_instr_tipo(IInstr_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_declaracionContext differentiates from other interfaces.
	IsInstr_declaracionContext()
}

type Instr_declaracionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruction
	_R_LET       antlr.Token
	_ID          antlr.Token
	_expressions IExpressionsContext
	_instr_tipo  IInstr_tipoContext
}

func NewEmptyInstr_declaracionContext() *Instr_declaracionContext {
	var p = new(Instr_declaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_declaracion
	return p
}

func (*Instr_declaracionContext) IsInstr_declaracionContext() {}

func NewInstr_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_declaracionContext {
	var p = new(Instr_declaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_declaracion

	return p
}

func (s *Instr_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_declaracionContext) Get_R_LET() antlr.Token { return s._R_LET }

func (s *Instr_declaracionContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_declaracionContext) Set_R_LET(v antlr.Token) { s._R_LET = v }

func (s *Instr_declaracionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_declaracionContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_declaracionContext) Get_instr_tipo() IInstr_tipoContext { return s._instr_tipo }

func (s *Instr_declaracionContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_declaracionContext) Set_instr_tipo(v IInstr_tipoContext) { s._instr_tipo = v }

func (s *Instr_declaracionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_declaracionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_declaracionContext) R_LET() antlr.TerminalNode {
	return s.GetToken(ChemsR_LET, 0)
}

func (s *Instr_declaracionContext) R_MUT() antlr.TerminalNode {
	return s.GetToken(ChemsR_MUT, 0)
}

func (s *Instr_declaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_declaracionContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUAL, 0)
}

func (s *Instr_declaracionContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_declaracionContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_declaracionContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DOSPUNTOS, 0)
}

func (s *Instr_declaracionContext) Instr_tipo() IInstr_tipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_tipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_tipoContext)
}

func (s *Instr_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_declaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_declaracion(s)
	}
}

func (s *Instr_declaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_declaracion(s)
	}
}

func (p *Chems) Instr_declaracion() (localctx IInstr_declaracionContext) {
	localctx = NewInstr_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ChemsRULE_instr_declaracion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(184)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(185)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(186)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(187)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(188)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), interfaces.NULL, localctx.(*Instr_declaracionContext).Get_expressions().GetP(), true, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(192)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(193)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(194)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(195)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(196)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), nil, true, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(199)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(200)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(201)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(202)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(203)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(204)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(205)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(206)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), localctx.(*Instr_declaracionContext).Get_expressions().GetP(), true, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(209)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(210)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(211)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(212)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(213)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), interfaces.NULL, localctx.(*Instr_declaracionContext).Get_expressions().GetP(), false, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(216)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(217)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(218)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(219)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(220)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), nil, false, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(223)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(224)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(225)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(226)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(227)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(228)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(229)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_declaracionContext).instr = variable.NewDeclaration((func() string {
			if localctx.(*Instr_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_declaracionContext).Get_instr_tipo().GetTipo_exp(), localctx.(*Instr_declaracionContext).Get_expressions().GetP(), false, (func() int {
			if localctx.(*Instr_declaracionContext).Get_R_LET() == nil {
				return 0
			} else {
				return localctx.(*Instr_declaracionContext).Get_R_LET().GetLine()
			}
		}()), localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn())

	}

	return localctx
}

// IInstr_asignacionContext is an interface to support dynamic dispatch.
type IInstr_asignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_asignacionContext differentiates from other interfaces.
	IsInstr_asignacionContext()
}

type Instr_asignacionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruction
	_ID          antlr.Token
	_expressions IExpressionsContext
}

func NewEmptyInstr_asignacionContext() *Instr_asignacionContext {
	var p = new(Instr_asignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_asignacion
	return p
}

func (*Instr_asignacionContext) IsInstr_asignacionContext() {}

func NewInstr_asignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_asignacionContext {
	var p = new(Instr_asignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_asignacion

	return p
}

func (s *Instr_asignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_asignacionContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_asignacionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_asignacionContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_asignacionContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_asignacionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_asignacionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_asignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_asignacionContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUAL, 0)
}

func (s *Instr_asignacionContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_asignacionContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_asignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_asignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_asignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_asignacion(s)
	}
}

func (s *Instr_asignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_asignacion(s)
	}
}

func (p *Chems) Instr_asignacion() (localctx IInstr_asignacionContext) {
	localctx = NewInstr_asignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ChemsRULE_instr_asignacion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)

		var _m = p.Match(ChemsID)

		localctx.(*Instr_asignacionContext)._ID = _m
	}
	{
		p.SetState(235)
		p.Match(ChemsTK_IGUAL)
	}
	{
		p.SetState(236)

		var _x = p.Expressions()

		localctx.(*Instr_asignacionContext)._expressions = _x
	}
	{
		p.SetState(237)
		p.Match(ChemsTK_PUNTOCOMA)
	}
	localctx.(*Instr_asignacionContext).instr = variable.NewAssignment((func() string {
		if localctx.(*Instr_asignacionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Instr_asignacionContext).Get_ID().GetText()
		}
	}()), localctx.(*Instr_asignacionContext).Get_expressions().GetP(), (func() int {
		if localctx.(*Instr_asignacionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Instr_asignacionContext).Get_ID().GetLine()
		}
	}()), localctx.(*Instr_asignacionContext).Get_ID().GetColumn())

	return localctx
}

// IInstr_ifContext is an interface to support dynamic dispatch.
type IInstr_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_IF returns the _R_IF token.
	Get_R_IF() antlr.Token

	// Set_R_IF sets the _R_IF token.
	Set_R_IF(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// GetLeft_instr returns the left_instr rule contexts.
	GetLeft_instr() IInstruccionesContext

	// GetRight_intr returns the right_intr rule contexts.
	GetRight_intr() IInstruccionesContext

	// Get_instr_else_if returns the _instr_else_if rule contexts.
	Get_instr_else_if() IInstr_else_ifContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// SetLeft_instr sets the left_instr rule contexts.
	SetLeft_instr(IInstruccionesContext)

	// SetRight_intr sets the right_intr rule contexts.
	SetRight_intr(IInstruccionesContext)

	// Set_instr_else_if sets the _instr_else_if rule contexts.
	Set_instr_else_if(IInstr_else_ifContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_ifContext differentiates from other interfaces.
	IsInstr_ifContext()
}

type Instr_ifContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_IF          antlr.Token
	_expressions   IExpressionsContext
	left_instr     IInstruccionesContext
	right_intr     IInstruccionesContext
	_instr_else_if IInstr_else_ifContext
}

func NewEmptyInstr_ifContext() *Instr_ifContext {
	var p = new(Instr_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_if
	return p
}

func (*Instr_ifContext) IsInstr_ifContext() {}

func NewInstr_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_ifContext {
	var p = new(Instr_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_if

	return p
}

func (s *Instr_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_ifContext) Get_R_IF() antlr.Token { return s._R_IF }

func (s *Instr_ifContext) Set_R_IF(v antlr.Token) { s._R_IF = v }

func (s *Instr_ifContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_ifContext) GetLeft_instr() IInstruccionesContext { return s.left_instr }

func (s *Instr_ifContext) GetRight_intr() IInstruccionesContext { return s.right_intr }

func (s *Instr_ifContext) Get_instr_else_if() IInstr_else_ifContext { return s._instr_else_if }

func (s *Instr_ifContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_ifContext) SetLeft_instr(v IInstruccionesContext) { s.left_instr = v }

func (s *Instr_ifContext) SetRight_intr(v IInstruccionesContext) { s.right_intr = v }

func (s *Instr_ifContext) Set_instr_else_if(v IInstr_else_ifContext) { s._instr_else_if = v }

func (s *Instr_ifContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_ifContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_ifContext) R_IF() antlr.TerminalNode {
	return s.GetToken(ChemsR_IF, 0)
}

func (s *Instr_ifContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_ifContext) AllTK_LLAVEA() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEA)
}

func (s *Instr_ifContext) TK_LLAVEA(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, i)
}

func (s *Instr_ifContext) AllTK_LLAVEC() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEC)
}

func (s *Instr_ifContext) TK_LLAVEC(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, i)
}

func (s *Instr_ifContext) AllInstrucciones() []IInstruccionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem())
	var tst = make([]IInstruccionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionesContext)
		}
	}

	return tst
}

func (s *Instr_ifContext) Instrucciones(i int) IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_ifContext) R_ELSE() antlr.TerminalNode {
	return s.GetToken(ChemsR_ELSE, 0)
}

func (s *Instr_ifContext) Instr_else_if() IInstr_else_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_else_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_else_ifContext)
}

func (s *Instr_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_if(s)
	}
}

func (s *Instr_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_if(s)
	}
}

func (p *Chems) Instr_if() (localctx IInstr_ifContext) {
	localctx = NewInstr_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ChemsRULE_instr_if)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(241)

			var _x = p.Expressions()

			localctx.(*Instr_ifContext)._expressions = _x
		}
		{
			p.SetState(242)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(243)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(244)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ifContext).instr = control.NewIf(localctx.(*Instr_ifContext).Get_expressions().GetP(), localctx.(*Instr_ifContext).GetLeft_instr().GetL(), nil, nil, (func() int {
			if localctx.(*Instr_ifContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ifContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ifContext).Get_R_IF().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(248)

			var _x = p.Expressions()

			localctx.(*Instr_ifContext)._expressions = _x
		}
		{
			p.SetState(249)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(250)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(251)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(252)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(253)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(254)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).right_intr = _x
		}
		{
			p.SetState(255)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ifContext).instr = control.NewIf(localctx.(*Instr_ifContext).Get_expressions().GetP(), localctx.(*Instr_ifContext).GetLeft_instr().GetL(), localctx.(*Instr_ifContext).GetRight_intr().GetL(), nil, (func() int {
			if localctx.(*Instr_ifContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ifContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ifContext).Get_R_IF().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(258)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ifContext)._R_IF = _m
		}
		{
			p.SetState(259)

			var _x = p.Expressions()

			localctx.(*Instr_ifContext)._expressions = _x
		}
		{
			p.SetState(260)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(261)

			var _x = p.Instrucciones()

			localctx.(*Instr_ifContext).left_instr = _x
		}
		{
			p.SetState(262)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(263)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(264)

			var _x = p.Instr_else_if()

			localctx.(*Instr_ifContext)._instr_else_if = _x
		}
		localctx.(*Instr_ifContext).instr = control.NewIf(localctx.(*Instr_ifContext).Get_expressions().GetP(), localctx.(*Instr_ifContext).GetLeft_instr().GetL(), nil, localctx.(*Instr_ifContext).Get_instr_else_if().GetL(), (func() int {
			if localctx.(*Instr_ifContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ifContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ifContext).Get_R_IF().GetColumn())

	}

	return localctx
}

// IInstr_else_ifContext is an interface to support dynamic dispatch.
type IInstr_else_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_if returns the _instr_if rule contexts.
	Get_instr_if() IInstr_ifContext

	// Set_instr_if sets the _instr_if rule contexts.
	Set_instr_if(IInstr_ifContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_ifContext

	// SetE sets the e rule context list.
	SetE([]IInstr_ifContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstr_else_ifContext differentiates from other interfaces.
	IsInstr_else_ifContext()
}

type Instr_else_ifContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	l         *arrayList.List
	_instr_if IInstr_ifContext
	e         []IInstr_ifContext
}

func NewEmptyInstr_else_ifContext() *Instr_else_ifContext {
	var p = new(Instr_else_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_else_if
	return p
}

func (*Instr_else_ifContext) IsInstr_else_ifContext() {}

func NewInstr_else_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_else_ifContext {
	var p = new(Instr_else_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_else_if

	return p
}

func (s *Instr_else_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_else_ifContext) Get_instr_if() IInstr_ifContext { return s._instr_if }

func (s *Instr_else_ifContext) Set_instr_if(v IInstr_ifContext) { s._instr_if = v }

func (s *Instr_else_ifContext) GetE() []IInstr_ifContext { return s.e }

func (s *Instr_else_ifContext) SetE(v []IInstr_ifContext) { s.e = v }

func (s *Instr_else_ifContext) GetL() *arrayList.List { return s.l }

func (s *Instr_else_ifContext) SetL(v *arrayList.List) { s.l = v }

func (s *Instr_else_ifContext) AllInstr_if() []IInstr_ifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_ifContext)(nil)).Elem())
	var tst = make([]IInstr_ifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_ifContext)
		}
	}

	return tst
}

func (s *Instr_else_ifContext) Instr_if(i int) IInstr_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_ifContext)
}

func (s *Instr_else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_else_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_else_if(s)
	}
}

func (s *Instr_else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_else_if(s)
	}
}

func (p *Chems) Instr_else_if() (localctx IInstr_else_ifContext) {
	localctx = NewInstr_else_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ChemsRULE_instr_else_if)

	localctx.(*Instr_else_ifContext).l = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(269)

				var _x = p.Instr_if()

				localctx.(*Instr_else_ifContext)._instr_if = _x
			}
			localctx.(*Instr_else_ifContext).e = append(localctx.(*Instr_else_ifContext).e, localctx.(*Instr_else_ifContext)._instr_if)

		}
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	listInt := localctx.(*Instr_else_ifContext).GetE()
	for _, e := range listInt {
		localctx.(*Instr_else_ifContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstr_ternarioContext is an interface to support dynamic dispatch.
type IInstr_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_IF returns the _R_IF token.
	Get_R_IF() antlr.Token

	// Set_R_IF sets the _R_IF token.
	Set_R_IF(antlr.Token)

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionsContext

	// GetLeft_instr returns the left_instr rule contexts.
	GetLeft_instr() IExpressionsContext

	// GetRight_intr returns the right_intr rule contexts.
	GetRight_intr() IExpressionsContext

	// Get_instr_else_if_ternario returns the _instr_else_if_ternario rule contexts.
	Get_instr_else_if_ternario() IInstr_else_if_ternarioContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionsContext)

	// SetLeft_instr sets the left_instr rule contexts.
	SetLeft_instr(IExpressionsContext)

	// SetRight_intr sets the right_intr rule contexts.
	SetRight_intr(IExpressionsContext)

	// Set_instr_else_if_ternario sets the _instr_else_if_ternario rule contexts.
	Set_instr_else_if_ternario(IInstr_else_if_ternarioContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_ternarioContext differentiates from other interfaces.
	IsInstr_ternarioContext()
}

type Instr_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser                  antlr.Parser
	p                       interfaces.Expression
	_R_IF                   antlr.Token
	cond                    IExpressionsContext
	left_instr              IExpressionsContext
	right_intr              IExpressionsContext
	_instr_else_if_ternario IInstr_else_if_ternarioContext
}

func NewEmptyInstr_ternarioContext() *Instr_ternarioContext {
	var p = new(Instr_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_ternario
	return p
}

func (*Instr_ternarioContext) IsInstr_ternarioContext() {}

func NewInstr_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_ternarioContext {
	var p = new(Instr_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_ternario

	return p
}

func (s *Instr_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_ternarioContext) Get_R_IF() antlr.Token { return s._R_IF }

func (s *Instr_ternarioContext) Set_R_IF(v antlr.Token) { s._R_IF = v }

func (s *Instr_ternarioContext) GetCond() IExpressionsContext { return s.cond }

func (s *Instr_ternarioContext) GetLeft_instr() IExpressionsContext { return s.left_instr }

func (s *Instr_ternarioContext) GetRight_intr() IExpressionsContext { return s.right_intr }

func (s *Instr_ternarioContext) Get_instr_else_if_ternario() IInstr_else_if_ternarioContext {
	return s._instr_else_if_ternario
}

func (s *Instr_ternarioContext) SetCond(v IExpressionsContext) { s.cond = v }

func (s *Instr_ternarioContext) SetLeft_instr(v IExpressionsContext) { s.left_instr = v }

func (s *Instr_ternarioContext) SetRight_intr(v IExpressionsContext) { s.right_intr = v }

func (s *Instr_ternarioContext) Set_instr_else_if_ternario(v IInstr_else_if_ternarioContext) {
	s._instr_else_if_ternario = v
}

func (s *Instr_ternarioContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_ternarioContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_ternarioContext) R_IF() antlr.TerminalNode {
	return s.GetToken(ChemsR_IF, 0)
}

func (s *Instr_ternarioContext) AllTK_LLAVEA() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEA)
}

func (s *Instr_ternarioContext) TK_LLAVEA(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, i)
}

func (s *Instr_ternarioContext) AllTK_LLAVEC() []antlr.TerminalNode {
	return s.GetTokens(ChemsTK_LLAVEC)
}

func (s *Instr_ternarioContext) TK_LLAVEC(i int) antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, i)
}

func (s *Instr_ternarioContext) AllExpressions() []IExpressionsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionsContext)(nil)).Elem())
	var tst = make([]IExpressionsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionsContext)
		}
	}

	return tst
}

func (s *Instr_ternarioContext) Expressions(i int) IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_ternarioContext) R_ELSE() antlr.TerminalNode {
	return s.GetToken(ChemsR_ELSE, 0)
}

func (s *Instr_ternarioContext) Instr_else_if_ternario() IInstr_else_if_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_else_if_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_else_if_ternarioContext)
}

func (s *Instr_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_ternario(s)
	}
}

func (s *Instr_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_ternario(s)
	}
}

func (p *Chems) Instr_ternario() (localctx IInstr_ternarioContext) {
	localctx = NewInstr_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ChemsRULE_instr_ternario)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(278)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).cond = _x
		}
		{
			p.SetState(279)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(280)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(281)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ternarioContext).p = ternario.NewIf(localctx.(*Instr_ternarioContext).GetCond().GetP(), localctx.(*Instr_ternarioContext).GetLeft_instr().GetP(), nil, nil, (func() int {
			if localctx.(*Instr_ternarioContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ternarioContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(284)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(285)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).cond = _x
		}
		{
			p.SetState(286)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(287)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(288)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(289)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(290)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(291)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).right_intr = _x
		}
		{
			p.SetState(292)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_ternarioContext).p = ternario.NewIf(localctx.(*Instr_ternarioContext).GetCond().GetP(), localctx.(*Instr_ternarioContext).GetLeft_instr().GetP(), localctx.(*Instr_ternarioContext).GetRight_intr().GetP(), nil, (func() int {
			if localctx.(*Instr_ternarioContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ternarioContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)

			var _m = p.Match(ChemsR_IF)

			localctx.(*Instr_ternarioContext)._R_IF = _m
		}
		{
			p.SetState(296)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).cond = _x
		}
		{
			p.SetState(297)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(298)

			var _x = p.Expressions()

			localctx.(*Instr_ternarioContext).left_instr = _x
		}
		{
			p.SetState(299)
			p.Match(ChemsTK_LLAVEC)
		}
		{
			p.SetState(300)
			p.Match(ChemsR_ELSE)
		}
		{
			p.SetState(301)

			var _x = p.Instr_else_if_ternario()

			localctx.(*Instr_ternarioContext)._instr_else_if_ternario = _x
		}
		localctx.(*Instr_ternarioContext).p = ternario.NewIf(localctx.(*Instr_ternarioContext).GetCond().GetP(), localctx.(*Instr_ternarioContext).GetLeft_instr().GetP(), nil, localctx.(*Instr_ternarioContext).Get_instr_else_if_ternario().GetL(), (func() int {
			if localctx.(*Instr_ternarioContext).Get_R_IF() == nil {
				return 0
			} else {
				return localctx.(*Instr_ternarioContext).Get_R_IF().GetLine()
			}
		}()), localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn())

	}

	return localctx
}

// IInstr_else_if_ternarioContext is an interface to support dynamic dispatch.
type IInstr_else_if_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_ternario returns the _instr_ternario rule contexts.
	Get_instr_ternario() IInstr_ternarioContext

	// Set_instr_ternario sets the _instr_ternario rule contexts.
	Set_instr_ternario(IInstr_ternarioContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_ternarioContext

	// SetE sets the e rule context list.
	SetE([]IInstr_ternarioContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstr_else_if_ternarioContext differentiates from other interfaces.
	IsInstr_else_if_ternarioContext()
}

type Instr_else_if_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	l               *arrayList.List
	_instr_ternario IInstr_ternarioContext
	e               []IInstr_ternarioContext
}

func NewEmptyInstr_else_if_ternarioContext() *Instr_else_if_ternarioContext {
	var p = new(Instr_else_if_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_else_if_ternario
	return p
}

func (*Instr_else_if_ternarioContext) IsInstr_else_if_ternarioContext() {}

func NewInstr_else_if_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_else_if_ternarioContext {
	var p = new(Instr_else_if_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_else_if_ternario

	return p
}

func (s *Instr_else_if_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_else_if_ternarioContext) Get_instr_ternario() IInstr_ternarioContext {
	return s._instr_ternario
}

func (s *Instr_else_if_ternarioContext) Set_instr_ternario(v IInstr_ternarioContext) {
	s._instr_ternario = v
}

func (s *Instr_else_if_ternarioContext) GetE() []IInstr_ternarioContext { return s.e }

func (s *Instr_else_if_ternarioContext) SetE(v []IInstr_ternarioContext) { s.e = v }

func (s *Instr_else_if_ternarioContext) GetL() *arrayList.List { return s.l }

func (s *Instr_else_if_ternarioContext) SetL(v *arrayList.List) { s.l = v }

func (s *Instr_else_if_ternarioContext) AllInstr_ternario() []IInstr_ternarioContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_ternarioContext)(nil)).Elem())
	var tst = make([]IInstr_ternarioContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_ternarioContext)
		}
	}

	return tst
}

func (s *Instr_else_if_ternarioContext) Instr_ternario(i int) IInstr_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ternarioContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_ternarioContext)
}

func (s *Instr_else_if_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_else_if_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_else_if_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_else_if_ternario(s)
	}
}

func (s *Instr_else_if_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_else_if_ternario(s)
	}
}

func (p *Chems) Instr_else_if_ternario() (localctx IInstr_else_if_ternarioContext) {
	localctx = NewInstr_else_if_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ChemsRULE_instr_else_if_ternario)

	localctx.(*Instr_else_if_ternarioContext).l = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(306)

				var _x = p.Instr_ternario()

				localctx.(*Instr_else_if_ternarioContext)._instr_ternario = _x
			}
			localctx.(*Instr_else_if_ternarioContext).e = append(localctx.(*Instr_else_if_ternarioContext).e, localctx.(*Instr_else_if_ternarioContext)._instr_ternario)

		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	listInt := localctx.(*Instr_else_if_ternarioContext).GetE()
	for _, e := range listInt {
		localctx.(*Instr_else_if_ternarioContext).l.Add(e.GetP())
	}

	return localctx
}

// IInstr_matchContext is an interface to support dynamic dispatch.
type IInstr_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_MATCH returns the _R_MATCH token.
	Get_R_MATCH() antlr.Token

	// Set_R_MATCH sets the _R_MATCH token.
	Set_R_MATCH(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Get_list_case returns the _list_case rule contexts.
	Get_list_case() IList_caseContext

	// Get_block_default returns the _block_default rule contexts.
	Get_block_default() IBlock_defaultContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// Set_list_case sets the _list_case rule contexts.
	Set_list_case(IList_caseContext)

	// Set_block_default sets the _block_default rule contexts.
	Set_block_default(IBlock_defaultContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_matchContext differentiates from other interfaces.
	IsInstr_matchContext()
}

type Instr_matchContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_MATCH       antlr.Token
	_expressions   IExpressionsContext
	_list_case     IList_caseContext
	_block_default IBlock_defaultContext
}

func NewEmptyInstr_matchContext() *Instr_matchContext {
	var p = new(Instr_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_match
	return p
}

func (*Instr_matchContext) IsInstr_matchContext() {}

func NewInstr_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_matchContext {
	var p = new(Instr_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_match

	return p
}

func (s *Instr_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_matchContext) Get_R_MATCH() antlr.Token { return s._R_MATCH }

func (s *Instr_matchContext) Set_R_MATCH(v antlr.Token) { s._R_MATCH = v }

func (s *Instr_matchContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_matchContext) Get_list_case() IList_caseContext { return s._list_case }

func (s *Instr_matchContext) Get_block_default() IBlock_defaultContext { return s._block_default }

func (s *Instr_matchContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_matchContext) Set_list_case(v IList_caseContext) { s._list_case = v }

func (s *Instr_matchContext) Set_block_default(v IBlock_defaultContext) { s._block_default = v }

func (s *Instr_matchContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_matchContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_matchContext) R_MATCH() antlr.TerminalNode {
	return s.GetToken(ChemsR_MATCH, 0)
}

func (s *Instr_matchContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_matchContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_matchContext) List_case() IList_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_caseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_caseContext)
}

func (s *Instr_matchContext) Block_default() IBlock_defaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_defaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_defaultContext)
}

func (s *Instr_matchContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_match(s)
	}
}

func (s *Instr_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_match(s)
	}
}

func (p *Chems) Instr_match() (localctx IInstr_matchContext) {
	localctx = NewInstr_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ChemsRULE_instr_match)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(314)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(315)

			var _x = p.Expressions()

			localctx.(*Instr_matchContext)._expressions = _x
		}
		{
			p.SetState(316)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(317)

			var _x = p.List_case()

			localctx.(*Instr_matchContext)._list_case = _x
		}
		{
			p.SetState(318)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(319)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_matchContext).instr = control.NewMatch(localctx.(*Instr_matchContext).Get_expressions().GetP(), localctx.(*Instr_matchContext).Get_list_case().GetL(), localctx.(*Instr_matchContext).Get_block_default().GetL(), (func() int {
			if localctx.(*Instr_matchContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_matchContext).Get_R_MATCH().GetLine()
			}
		}()), localctx.(*Instr_matchContext).Get_R_MATCH().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_matchContext)._R_MATCH = _m
		}
		{
			p.SetState(323)

			var _x = p.Expressions()

			localctx.(*Instr_matchContext)._expressions = _x
		}
		{
			p.SetState(324)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(325)

			var _x = p.Block_default()

			localctx.(*Instr_matchContext)._block_default = _x
		}
		{
			p.SetState(326)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_matchContext).instr = control.NewMatch(localctx.(*Instr_matchContext).Get_expressions().GetP(), nil, localctx.(*Instr_matchContext).Get_block_default().GetL(), (func() int {
			if localctx.(*Instr_matchContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_matchContext).Get_R_MATCH().GetLine()
			}
		}()), localctx.(*Instr_matchContext).Get_R_MATCH().GetColumn())

	}

	return localctx
}

// IList_caseContext is an interface to support dynamic dispatch.
type IList_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_case returns the _instr_case rule contexts.
	Get_instr_case() IInstr_caseContext

	// Set_instr_case sets the _instr_case rule contexts.
	Set_instr_case(IInstr_caseContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_caseContext

	// SetE sets the e rule context list.
	SetE([]IInstr_caseContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_caseContext differentiates from other interfaces.
	IsList_caseContext()
}

type List_caseContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	_instr_case IInstr_caseContext
	e           []IInstr_caseContext
}

func NewEmptyList_caseContext() *List_caseContext {
	var p = new(List_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_case
	return p
}

func (*List_caseContext) IsList_caseContext() {}

func NewList_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_caseContext {
	var p = new(List_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_case

	return p
}

func (s *List_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *List_caseContext) Get_instr_case() IInstr_caseContext { return s._instr_case }

func (s *List_caseContext) Set_instr_case(v IInstr_caseContext) { s._instr_case = v }

func (s *List_caseContext) GetE() []IInstr_caseContext { return s.e }

func (s *List_caseContext) SetE(v []IInstr_caseContext) { s.e = v }

func (s *List_caseContext) GetL() *arrayList.List { return s.l }

func (s *List_caseContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_caseContext) AllInstr_case() []IInstr_caseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_caseContext)(nil)).Elem())
	var tst = make([]IInstr_caseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_caseContext)
		}
	}

	return tst
}

func (s *List_caseContext) Instr_case(i int) IInstr_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_caseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_caseContext)
}

func (s *List_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_case(s)
	}
}

func (s *List_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_case(s)
	}
}

func (p *Chems) List_case() (localctx IList_caseContext) {
	localctx = NewList_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ChemsRULE_list_case)

	localctx.(*List_caseContext).l = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(331)

				var _x = p.Instr_case()

				localctx.(*List_caseContext)._instr_case = _x
			}
			localctx.(*List_caseContext).e = append(localctx.(*List_caseContext).e, localctx.(*List_caseContext)._instr_case)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	listInt := localctx.(*List_caseContext).GetE()
	for _, e := range listInt {
		localctx.(*List_caseContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstr_caseContext is an interface to support dynamic dispatch.
type IInstr_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_list_expre_case returns the _list_expre_case rule contexts.
	Get_list_expre_case() IList_expre_caseContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_block_instr_match returns the _block_instr_match rule contexts.
	Get_block_instr_match() IBlock_instr_matchContext

	// Set_list_expre_case sets the _list_expre_case rule contexts.
	Set_list_expre_case(IList_expre_caseContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_block_instr_match sets the _block_instr_match rule contexts.
	Set_block_instr_match(IBlock_instr_matchContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expression

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expression)

	// IsInstr_caseContext differentiates from other interfaces.
	IsInstr_caseContext()
}

type Instr_caseContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	instr              interfaces.Expression
	_list_expre_case   IList_expre_caseContext
	_instrucciones     IInstruccionesContext
	_block_instr_match IBlock_instr_matchContext
}

func NewEmptyInstr_caseContext() *Instr_caseContext {
	var p = new(Instr_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_case
	return p
}

func (*Instr_caseContext) IsInstr_caseContext() {}

func NewInstr_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_caseContext {
	var p = new(Instr_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_case

	return p
}

func (s *Instr_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_caseContext) Get_list_expre_case() IList_expre_caseContext { return s._list_expre_case }

func (s *Instr_caseContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_caseContext) Get_block_instr_match() IBlock_instr_matchContext {
	return s._block_instr_match
}

func (s *Instr_caseContext) Set_list_expre_case(v IList_expre_caseContext) { s._list_expre_case = v }

func (s *Instr_caseContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_caseContext) Set_block_instr_match(v IBlock_instr_matchContext) {
	s._block_instr_match = v
}

func (s *Instr_caseContext) GetInstr() interfaces.Expression { return s.instr }

func (s *Instr_caseContext) SetInstr(v interfaces.Expression) { s.instr = v }

func (s *Instr_caseContext) List_expre_case() IList_expre_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expre_caseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expre_caseContext)
}

func (s *Instr_caseContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_caseContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_caseContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_caseContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_caseContext) Block_instr_match() IBlock_instr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_instr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_instr_matchContext)
}

func (s *Instr_caseContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_case(s)
	}
}

func (s *Instr_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_case(s)
	}
}

func (p *Chems) Instr_case() (localctx IInstr_caseContext) {
	localctx = NewInstr_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ChemsRULE_instr_case)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)

			var _x = p.List_expre_case()

			localctx.(*Instr_caseContext)._list_expre_case = _x
		}
		{
			p.SetState(339)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(340)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(341)

			var _x = p.Instrucciones()

			localctx.(*Instr_caseContext)._instrucciones = _x
		}
		{
			p.SetState(342)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_caseContext).instr = control.NewCase(nil, localctx.(*Instr_caseContext).Get_list_expre_case().GetL(), localctx.(*Instr_caseContext).Get_instrucciones().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)

			var _x = p.List_expre_case()

			localctx.(*Instr_caseContext)._list_expre_case = _x
		}
		{
			p.SetState(346)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(347)

			var _x = p.Block_instr_match()

			localctx.(*Instr_caseContext)._block_instr_match = _x
		}
		{
			p.SetState(348)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Instr_caseContext).instr = control.NewCase(nil, localctx.(*Instr_caseContext).Get_list_expre_case().GetL(), localctx.(*Instr_caseContext).Get_block_instr_match().GetL())

	}

	return localctx
}

// IList_expre_caseContext is an interface to support dynamic dispatch.
type IList_expre_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_case returns the _block_case rule contexts.
	Get_block_case() IBlock_caseContext

	// Set_block_case sets the _block_case rule contexts.
	Set_block_case(IBlock_caseContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_caseContext

	// SetE sets the e rule context list.
	SetE([]IBlock_caseContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_expre_caseContext differentiates from other interfaces.
	IsList_expre_caseContext()
}

type List_expre_caseContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	_block_case IBlock_caseContext
	e           []IBlock_caseContext
}

func NewEmptyList_expre_caseContext() *List_expre_caseContext {
	var p = new(List_expre_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_expre_case
	return p
}

func (*List_expre_caseContext) IsList_expre_caseContext() {}

func NewList_expre_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expre_caseContext {
	var p = new(List_expre_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_expre_case

	return p
}

func (s *List_expre_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expre_caseContext) Get_block_case() IBlock_caseContext { return s._block_case }

func (s *List_expre_caseContext) Set_block_case(v IBlock_caseContext) { s._block_case = v }

func (s *List_expre_caseContext) GetE() []IBlock_caseContext { return s.e }

func (s *List_expre_caseContext) SetE(v []IBlock_caseContext) { s.e = v }

func (s *List_expre_caseContext) GetL() *arrayList.List { return s.l }

func (s *List_expre_caseContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_expre_caseContext) AllBlock_case() []IBlock_caseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_caseContext)(nil)).Elem())
	var tst = make([]IBlock_caseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_caseContext)
		}
	}

	return tst
}

func (s *List_expre_caseContext) Block_case(i int) IBlock_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_caseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_caseContext)
}

func (s *List_expre_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expre_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expre_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_expre_case(s)
	}
}

func (s *List_expre_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_expre_case(s)
	}
}

func (p *Chems) List_expre_case() (localctx IList_expre_caseContext) {
	localctx = NewList_expre_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ChemsRULE_list_expre_case)

	localctx.(*List_expre_caseContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsCHAR)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(ChemsTK_MENOS-48))|(1<<(ChemsTK_PARA-48))|(1<<(ChemsTK_NOT-48)))) != 0) {
		{
			p.SetState(353)

			var _x = p.Block_case()

			localctx.(*List_expre_caseContext)._block_case = _x
		}
		localctx.(*List_expre_caseContext).e = append(localctx.(*List_expre_caseContext).e, localctx.(*List_expre_caseContext)._block_case)

		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expre_caseContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expre_caseContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IBlock_caseContext is an interface to support dynamic dispatch.
type IBlock_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Expression

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Expression)

	// IsBlock_caseContext differentiates from other interfaces.
	IsBlock_caseContext()
}

type Block_caseContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Expression
	_expressions IExpressionsContext
}

func NewEmptyBlock_caseContext() *Block_caseContext {
	var p = new(Block_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_case
	return p
}

func (*Block_caseContext) IsBlock_caseContext() {}

func NewBlock_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_caseContext {
	var p = new(Block_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_case

	return p
}

func (s *Block_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_caseContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Block_caseContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Block_caseContext) GetInstr() interfaces.Expression { return s.instr }

func (s *Block_caseContext) SetInstr(v interfaces.Expression) { s.instr = v }

func (s *Block_caseContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_caseContext) TK_BARRA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_BARRA, 0)
}

func (s *Block_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_case(s)
	}
}

func (s *Block_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_case(s)
	}
}

func (p *Chems) Block_case() (localctx IBlock_caseContext) {
	localctx = NewBlock_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ChemsRULE_block_case)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(360)

			var _x = p.Expressions()

			localctx.(*Block_caseContext)._expressions = _x
		}
		{
			p.SetState(361)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_caseContext).instr = control.NewCase(localctx.(*Block_caseContext).Get_expressions().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(364)

			var _x = p.Expressions()

			localctx.(*Block_caseContext)._expressions = _x
		}
		localctx.(*Block_caseContext).instr = control.NewCase(localctx.(*Block_caseContext).Get_expressions().GetP(), nil, nil)

	}

	return localctx
}

// IBlock_instr_matchContext is an interface to support dynamic dispatch.
type IBlock_instr_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetList returns the list rule context list.
	GetList() []IInstruccionContext

	// SetList sets the list rule context list.
	SetList([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsBlock_instr_matchContext differentiates from other interfaces.
	IsBlock_instr_matchContext()
}

type Block_instr_matchContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruccion IInstruccionContext
	list         []IInstruccionContext
}

func NewEmptyBlock_instr_matchContext() *Block_instr_matchContext {
	var p = new(Block_instr_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_instr_match
	return p
}

func (*Block_instr_matchContext) IsBlock_instr_matchContext() {}

func NewBlock_instr_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_instr_matchContext {
	var p = new(Block_instr_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_instr_match

	return p
}

func (s *Block_instr_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_instr_matchContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *Block_instr_matchContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *Block_instr_matchContext) GetList() []IInstruccionContext { return s.list }

func (s *Block_instr_matchContext) SetList(v []IInstruccionContext) { s.list = v }

func (s *Block_instr_matchContext) GetL() *arrayList.List { return s.l }

func (s *Block_instr_matchContext) SetL(v *arrayList.List) { s.l = v }

func (s *Block_instr_matchContext) Instruccion() IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *Block_instr_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_instr_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_instr_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_instr_match(s)
	}
}

func (s *Block_instr_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_instr_match(s)
	}
}

func (p *Chems) Block_instr_match() (localctx IBlock_instr_matchContext) {
	localctx = NewBlock_instr_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ChemsRULE_block_instr_match)

	localctx.(*Block_instr_matchContext).l = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)

		var _x = p.Instruccion()

		localctx.(*Block_instr_matchContext)._instruccion = _x
	}
	localctx.(*Block_instr_matchContext).list = append(localctx.(*Block_instr_matchContext).list, localctx.(*Block_instr_matchContext)._instruccion)

	listInt := localctx.(*Block_instr_matchContext).GetList()
	for _, e := range listInt {
		localctx.(*Block_instr_matchContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstr_defaultContext is an interface to support dynamic dispatch.
type IInstr_defaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_block_instr_match returns the _block_instr_match rule contexts.
	Get_block_instr_match() IBlock_instr_matchContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_block_instr_match sets the _block_instr_match rule contexts.
	Set_block_instr_match(IBlock_instr_matchContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_defaultContext differentiates from other interfaces.
	IsInstr_defaultContext()
}

type Instr_defaultContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	instr              interfaces.Instruction
	_instrucciones     IInstruccionesContext
	_block_instr_match IBlock_instr_matchContext
}

func NewEmptyInstr_defaultContext() *Instr_defaultContext {
	var p = new(Instr_defaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_default
	return p
}

func (*Instr_defaultContext) IsInstr_defaultContext() {}

func NewInstr_defaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_defaultContext {
	var p = new(Instr_defaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_default

	return p
}

func (s *Instr_defaultContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_defaultContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_defaultContext) Get_block_instr_match() IBlock_instr_matchContext {
	return s._block_instr_match
}

func (s *Instr_defaultContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_defaultContext) Set_block_instr_match(v IBlock_instr_matchContext) {
	s._block_instr_match = v
}

func (s *Instr_defaultContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_defaultContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_defaultContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_defaultContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_defaultContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_defaultContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_defaultContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_defaultContext) Block_instr_match() IBlock_instr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_instr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_instr_matchContext)
}

func (s *Instr_defaultContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_defaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_defaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_defaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_default(s)
	}
}

func (s *Instr_defaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_default(s)
	}
}

func (p *Chems) Instr_default() (localctx IInstr_defaultContext) {
	localctx = NewInstr_defaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ChemsRULE_instr_default)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(372)
			p.Match(ChemsID)
		}
		{
			p.SetState(373)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(374)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(375)

			var _x = p.Instrucciones()

			localctx.(*Instr_defaultContext)._instrucciones = _x
		}
		{
			p.SetState(376)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_defaultContext).instr = control.NewDefault(localctx.(*Instr_defaultContext).Get_instrucciones().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(379)
			p.Match(ChemsID)
		}
		{
			p.SetState(380)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(381)

			var _x = p.Block_instr_match()

			localctx.(*Instr_defaultContext)._block_instr_match = _x
		}
		{
			p.SetState(382)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Instr_defaultContext).instr = control.NewDefault(localctx.(*Instr_defaultContext).Get_block_instr_match().GetL())

	}

	return localctx
}

// IBlock_defaultContext is an interface to support dynamic dispatch.
type IBlock_defaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_default returns the _instr_default rule contexts.
	Get_instr_default() IInstr_defaultContext

	// Set_instr_default sets the _instr_default rule contexts.
	Set_instr_default(IInstr_defaultContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_defaultContext

	// SetE sets the e rule context list.
	SetE([]IInstr_defaultContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsBlock_defaultContext differentiates from other interfaces.
	IsBlock_defaultContext()
}

type Block_defaultContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	l              *arrayList.List
	_instr_default IInstr_defaultContext
	e              []IInstr_defaultContext
}

func NewEmptyBlock_defaultContext() *Block_defaultContext {
	var p = new(Block_defaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_default
	return p
}

func (*Block_defaultContext) IsBlock_defaultContext() {}

func NewBlock_defaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_defaultContext {
	var p = new(Block_defaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_default

	return p
}

func (s *Block_defaultContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_defaultContext) Get_instr_default() IInstr_defaultContext { return s._instr_default }

func (s *Block_defaultContext) Set_instr_default(v IInstr_defaultContext) { s._instr_default = v }

func (s *Block_defaultContext) GetE() []IInstr_defaultContext { return s.e }

func (s *Block_defaultContext) SetE(v []IInstr_defaultContext) { s.e = v }

func (s *Block_defaultContext) GetL() *arrayList.List { return s.l }

func (s *Block_defaultContext) SetL(v *arrayList.List) { s.l = v }

func (s *Block_defaultContext) AllInstr_default() []IInstr_defaultContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_defaultContext)(nil)).Elem())
	var tst = make([]IInstr_defaultContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_defaultContext)
		}
	}

	return tst
}

func (s *Block_defaultContext) Instr_default(i int) IInstr_defaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_defaultContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_defaultContext)
}

func (s *Block_defaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_defaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_defaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_default(s)
	}
}

func (s *Block_defaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_default(s)
	}
}

func (p *Chems) Block_default() (localctx IBlock_defaultContext) {
	localctx = NewBlock_defaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ChemsRULE_block_default)

	localctx.(*Block_defaultContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ChemsID {
		{
			p.SetState(387)

			var _x = p.Instr_default()

			localctx.(*Block_defaultContext)._instr_default = _x
		}
		localctx.(*Block_defaultContext).e = append(localctx.(*Block_defaultContext).e, localctx.(*Block_defaultContext)._instr_default)

		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*Block_defaultContext).GetE()
	for _, e := range listInt {
		localctx.(*Block_defaultContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstr_match_terContext is an interface to support dynamic dispatch.
type IInstr_match_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_MATCH returns the _R_MATCH token.
	Get_R_MATCH() antlr.Token

	// Set_R_MATCH sets the _R_MATCH token.
	Set_R_MATCH(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Get_list_case_ternario returns the _list_case_ternario rule contexts.
	Get_list_case_ternario() IList_case_ternarioContext

	// Get_instr_default_ter returns the _instr_default_ter rule contexts.
	Get_instr_default_ter() IInstr_default_terContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// Set_list_case_ternario sets the _list_case_ternario rule contexts.
	Set_list_case_ternario(IList_case_ternarioContext)

	// Set_instr_default_ter sets the _instr_default_ter rule contexts.
	Set_instr_default_ter(IInstr_default_terContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_match_terContext differentiates from other interfaces.
	IsInstr_match_terContext()
}

type Instr_match_terContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	p                   interfaces.Expression
	_R_MATCH            antlr.Token
	_expressions        IExpressionsContext
	_list_case_ternario IList_case_ternarioContext
	_instr_default_ter  IInstr_default_terContext
}

func NewEmptyInstr_match_terContext() *Instr_match_terContext {
	var p = new(Instr_match_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_match_ter
	return p
}

func (*Instr_match_terContext) IsInstr_match_terContext() {}

func NewInstr_match_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_match_terContext {
	var p = new(Instr_match_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_match_ter

	return p
}

func (s *Instr_match_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_match_terContext) Get_R_MATCH() antlr.Token { return s._R_MATCH }

func (s *Instr_match_terContext) Set_R_MATCH(v antlr.Token) { s._R_MATCH = v }

func (s *Instr_match_terContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_match_terContext) Get_list_case_ternario() IList_case_ternarioContext {
	return s._list_case_ternario
}

func (s *Instr_match_terContext) Get_instr_default_ter() IInstr_default_terContext {
	return s._instr_default_ter
}

func (s *Instr_match_terContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_match_terContext) Set_list_case_ternario(v IList_case_ternarioContext) {
	s._list_case_ternario = v
}

func (s *Instr_match_terContext) Set_instr_default_ter(v IInstr_default_terContext) {
	s._instr_default_ter = v
}

func (s *Instr_match_terContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_match_terContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_match_terContext) R_MATCH() antlr.TerminalNode {
	return s.GetToken(ChemsR_MATCH, 0)
}

func (s *Instr_match_terContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_match_terContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_match_terContext) List_case_ternario() IList_case_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_case_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_case_ternarioContext)
}

func (s *Instr_match_terContext) Instr_default_ter() IInstr_default_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_default_terContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_default_terContext)
}

func (s *Instr_match_terContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_match_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_match_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_match_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_match_ter(s)
	}
}

func (s *Instr_match_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_match_ter(s)
	}
}

func (p *Chems) Instr_match_ter() (localctx IInstr_match_terContext) {
	localctx = NewInstr_match_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ChemsRULE_instr_match_ter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(394)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_match_terContext)._R_MATCH = _m
		}
		{
			p.SetState(395)

			var _x = p.Expressions()

			localctx.(*Instr_match_terContext)._expressions = _x
		}
		{
			p.SetState(396)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(397)

			var _x = p.List_case_ternario()

			localctx.(*Instr_match_terContext)._list_case_ternario = _x
		}
		{
			p.SetState(398)

			var _x = p.Instr_default_ter()

			localctx.(*Instr_match_terContext)._instr_default_ter = _x
		}
		{
			p.SetState(399)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_match_terContext).p = ternario.NewMatch(localctx.(*Instr_match_terContext).Get_expressions().GetP(), localctx.(*Instr_match_terContext).Get_list_case_ternario().GetL(), localctx.(*Instr_match_terContext).Get_instr_default_ter().GetP(), (func() int {
			if localctx.(*Instr_match_terContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_match_terContext).Get_R_MATCH().GetLine()
			}
		}()), localctx.(*Instr_match_terContext).Get_R_MATCH().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(402)

			var _m = p.Match(ChemsR_MATCH)

			localctx.(*Instr_match_terContext)._R_MATCH = _m
		}
		{
			p.SetState(403)

			var _x = p.Expressions()

			localctx.(*Instr_match_terContext)._expressions = _x
		}
		{
			p.SetState(404)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(405)

			var _x = p.Instr_default_ter()

			localctx.(*Instr_match_terContext)._instr_default_ter = _x
		}
		{
			p.SetState(406)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_match_terContext).p = ternario.NewMatch(localctx.(*Instr_match_terContext).Get_expressions().GetP(), nil, localctx.(*Instr_match_terContext).Get_instr_default_ter().GetP(), (func() int {
			if localctx.(*Instr_match_terContext).Get_R_MATCH() == nil {
				return 0
			} else {
				return localctx.(*Instr_match_terContext).Get_R_MATCH().GetLine()
			}
		}()), localctx.(*Instr_match_terContext).Get_R_MATCH().GetColumn())

	}

	return localctx
}

// IList_case_ternarioContext is an interface to support dynamic dispatch.
type IList_case_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_case_ter returns the _instr_case_ter rule contexts.
	Get_instr_case_ter() IInstr_case_terContext

	// Set_instr_case_ter sets the _instr_case_ter rule contexts.
	Set_instr_case_ter(IInstr_case_terContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_case_terContext

	// SetE sets the e rule context list.
	SetE([]IInstr_case_terContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_case_ternarioContext differentiates from other interfaces.
	IsList_case_ternarioContext()
}

type List_case_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	l               *arrayList.List
	_instr_case_ter IInstr_case_terContext
	e               []IInstr_case_terContext
}

func NewEmptyList_case_ternarioContext() *List_case_ternarioContext {
	var p = new(List_case_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_case_ternario
	return p
}

func (*List_case_ternarioContext) IsList_case_ternarioContext() {}

func NewList_case_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_case_ternarioContext {
	var p = new(List_case_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_case_ternario

	return p
}

func (s *List_case_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *List_case_ternarioContext) Get_instr_case_ter() IInstr_case_terContext {
	return s._instr_case_ter
}

func (s *List_case_ternarioContext) Set_instr_case_ter(v IInstr_case_terContext) {
	s._instr_case_ter = v
}

func (s *List_case_ternarioContext) GetE() []IInstr_case_terContext { return s.e }

func (s *List_case_ternarioContext) SetE(v []IInstr_case_terContext) { s.e = v }

func (s *List_case_ternarioContext) GetL() *arrayList.List { return s.l }

func (s *List_case_ternarioContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_case_ternarioContext) AllInstr_case_ter() []IInstr_case_terContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_case_terContext)(nil)).Elem())
	var tst = make([]IInstr_case_terContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_case_terContext)
		}
	}

	return tst
}

func (s *List_case_ternarioContext) Instr_case_ter(i int) IInstr_case_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_case_terContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_case_terContext)
}

func (s *List_case_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_case_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_case_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_case_ternario(s)
	}
}

func (s *List_case_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_case_ternario(s)
	}
}

func (p *Chems) List_case_ternario() (localctx IList_case_ternarioContext) {
	localctx = NewList_case_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ChemsRULE_list_case_ternario)

	localctx.(*List_case_ternarioContext).l = arrayList.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(411)

				var _x = p.Instr_case_ter()

				localctx.(*List_case_ternarioContext)._instr_case_ter = _x
			}
			localctx.(*List_case_ternarioContext).e = append(localctx.(*List_case_ternarioContext).e, localctx.(*List_case_ternarioContext)._instr_case_ter)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	listInt := localctx.(*List_case_ternarioContext).GetE()
	for _, e := range listInt {
		localctx.(*List_case_ternarioContext).l.Add(e.GetP())
	}

	return localctx
}

// IInstr_case_terContext is an interface to support dynamic dispatch.
type IInstr_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_list_expre_case_ter returns the _list_expre_case_ter rule contexts.
	Get_list_expre_case_ter() IList_expre_case_terContext

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_list_expre_case_ter sets the _list_expre_case_ter rule contexts.
	Set_list_expre_case_ter(IList_expre_case_terContext)

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_case_terContext differentiates from other interfaces.
	IsInstr_case_terContext()
}

type Instr_case_terContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	p                    interfaces.Expression
	_list_expre_case_ter IList_expre_case_terContext
	_expressions         IExpressionsContext
}

func NewEmptyInstr_case_terContext() *Instr_case_terContext {
	var p = new(Instr_case_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_case_ter
	return p
}

func (*Instr_case_terContext) IsInstr_case_terContext() {}

func NewInstr_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_case_terContext {
	var p = new(Instr_case_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_case_ter

	return p
}

func (s *Instr_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_case_terContext) Get_list_expre_case_ter() IList_expre_case_terContext {
	return s._list_expre_case_ter
}

func (s *Instr_case_terContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_case_terContext) Set_list_expre_case_ter(v IList_expre_case_terContext) {
	s._list_expre_case_ter = v
}

func (s *Instr_case_terContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_case_terContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_case_terContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_case_terContext) List_expre_case_ter() IList_expre_case_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expre_case_terContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expre_case_terContext)
}

func (s *Instr_case_terContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_case_terContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_case_terContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_case_terContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_case_terContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_case_ter(s)
	}
}

func (s *Instr_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_case_ter(s)
	}
}

func (p *Chems) Instr_case_ter() (localctx IInstr_case_terContext) {
	localctx = NewInstr_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ChemsRULE_instr_case_ter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(418)

			var _x = p.List_expre_case_ter()

			localctx.(*Instr_case_terContext)._list_expre_case_ter = _x
		}
		{
			p.SetState(419)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(420)

			var _x = p.Expressions()

			localctx.(*Instr_case_terContext)._expressions = _x
		}
		{
			p.SetState(421)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Instr_case_terContext).p = ternario.NewCase(nil, localctx.(*Instr_case_terContext).Get_list_expre_case_ter().GetL(), localctx.(*Instr_case_terContext).Get_expressions().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(424)

			var _x = p.List_expre_case_ter()

			localctx.(*Instr_case_terContext)._list_expre_case_ter = _x
		}
		{
			p.SetState(425)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(426)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(427)

			var _x = p.Expressions()

			localctx.(*Instr_case_terContext)._expressions = _x
		}
		{
			p.SetState(428)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_case_terContext).p = ternario.NewCase(nil, localctx.(*Instr_case_terContext).Get_list_expre_case_ter().GetL(), localctx.(*Instr_case_terContext).Get_expressions().GetP())

	}

	return localctx
}

// IList_expre_case_terContext is an interface to support dynamic dispatch.
type IList_expre_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_case_ter returns the _block_case_ter rule contexts.
	Get_block_case_ter() IBlock_case_terContext

	// Set_block_case_ter sets the _block_case_ter rule contexts.
	Set_block_case_ter(IBlock_case_terContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_case_terContext

	// SetE sets the e rule context list.
	SetE([]IBlock_case_terContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_expre_case_terContext differentiates from other interfaces.
	IsList_expre_case_terContext()
}

type List_expre_case_terContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	l               *arrayList.List
	_block_case_ter IBlock_case_terContext
	e               []IBlock_case_terContext
}

func NewEmptyList_expre_case_terContext() *List_expre_case_terContext {
	var p = new(List_expre_case_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_expre_case_ter
	return p
}

func (*List_expre_case_terContext) IsList_expre_case_terContext() {}

func NewList_expre_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expre_case_terContext {
	var p = new(List_expre_case_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_expre_case_ter

	return p
}

func (s *List_expre_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expre_case_terContext) Get_block_case_ter() IBlock_case_terContext {
	return s._block_case_ter
}

func (s *List_expre_case_terContext) Set_block_case_ter(v IBlock_case_terContext) {
	s._block_case_ter = v
}

func (s *List_expre_case_terContext) GetE() []IBlock_case_terContext { return s.e }

func (s *List_expre_case_terContext) SetE(v []IBlock_case_terContext) { s.e = v }

func (s *List_expre_case_terContext) GetL() *arrayList.List { return s.l }

func (s *List_expre_case_terContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_expre_case_terContext) AllBlock_case_ter() []IBlock_case_terContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_case_terContext)(nil)).Elem())
	var tst = make([]IBlock_case_terContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_case_terContext)
		}
	}

	return tst
}

func (s *List_expre_case_terContext) Block_case_ter(i int) IBlock_case_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_case_terContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_case_terContext)
}

func (s *List_expre_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expre_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expre_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_expre_case_ter(s)
	}
}

func (s *List_expre_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_expre_case_ter(s)
	}
}

func (p *Chems) List_expre_case_ter() (localctx IList_expre_case_terContext) {
	localctx = NewList_expre_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ChemsRULE_list_expre_case_ter)

	localctx.(*List_expre_case_terContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsCHAR)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(ChemsTK_MENOS-48))|(1<<(ChemsTK_PARA-48))|(1<<(ChemsTK_NOT-48)))) != 0) {
		{
			p.SetState(433)

			var _x = p.Block_case_ter()

			localctx.(*List_expre_case_terContext)._block_case_ter = _x
		}
		localctx.(*List_expre_case_terContext).e = append(localctx.(*List_expre_case_terContext).e, localctx.(*List_expre_case_terContext)._block_case_ter)

		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expre_case_terContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expre_case_terContext).l.Add(e.GetP())
	}

	return localctx
}

// IBlock_case_terContext is an interface to support dynamic dispatch.
type IBlock_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsBlock_case_terContext differentiates from other interfaces.
	IsBlock_case_terContext()
}

type Block_case_terContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	_expressions IExpressionsContext
}

func NewEmptyBlock_case_terContext() *Block_case_terContext {
	var p = new(Block_case_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_case_ter
	return p
}

func (*Block_case_terContext) IsBlock_case_terContext() {}

func NewBlock_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_case_terContext {
	var p = new(Block_case_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_case_ter

	return p
}

func (s *Block_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_case_terContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Block_case_terContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Block_case_terContext) GetP() interfaces.Expression { return s.p }

func (s *Block_case_terContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Block_case_terContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_case_terContext) TK_BARRA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_BARRA, 0)
}

func (s *Block_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_case_ter(s)
	}
}

func (s *Block_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_case_ter(s)
	}
}

func (p *Chems) Block_case_ter() (localctx IBlock_case_terContext) {
	localctx = NewBlock_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ChemsRULE_block_case_ter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(440)

			var _x = p.Expressions()

			localctx.(*Block_case_terContext)._expressions = _x
		}
		{
			p.SetState(441)
			p.Match(ChemsTK_BARRA)
		}
		localctx.(*Block_case_terContext).p = ternario.NewCase(localctx.(*Block_case_terContext).Get_expressions().GetP(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(444)

			var _x = p.Expressions()

			localctx.(*Block_case_terContext)._expressions = _x
		}
		localctx.(*Block_case_terContext).p = ternario.NewCase(localctx.(*Block_case_terContext).Get_expressions().GetP(), nil, nil)

	}

	return localctx
}

// IInstr_default_terContext is an interface to support dynamic dispatch.
type IInstr_default_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_default_terContext differentiates from other interfaces.
	IsInstr_default_terContext()
}

type Instr_default_terContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	_expressions IExpressionsContext
}

func NewEmptyInstr_default_terContext() *Instr_default_terContext {
	var p = new(Instr_default_terContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_default_ter
	return p
}

func (*Instr_default_terContext) IsInstr_default_terContext() {}

func NewInstr_default_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_default_terContext {
	var p = new(Instr_default_terContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_default_ter

	return p
}

func (s *Instr_default_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_default_terContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_default_terContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_default_terContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_default_terContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_default_terContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_default_terContext) TK_IGUALMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALMAYOR, 0)
}

func (s *Instr_default_terContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_default_terContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_default_terContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_default_terContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_default_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_default_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_default_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_default_ter(s)
	}
}

func (s *Instr_default_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_default_ter(s)
	}
}

func (p *Chems) Instr_default_ter() (localctx IInstr_default_terContext) {
	localctx = NewInstr_default_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ChemsRULE_instr_default_ter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(449)
			p.Match(ChemsID)
		}
		{
			p.SetState(450)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(451)

			var _x = p.Expressions()

			localctx.(*Instr_default_terContext)._expressions = _x
		}
		{
			p.SetState(452)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Instr_default_terContext).p = ternario.NewDefault(localctx.(*Instr_default_terContext).Get_expressions().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Match(ChemsID)
		}
		{
			p.SetState(456)
			p.Match(ChemsTK_IGUALMAYOR)
		}
		{
			p.SetState(457)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(458)

			var _x = p.Expressions()

			localctx.(*Instr_default_terContext)._expressions = _x
		}
		{
			p.SetState(459)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_default_terContext).p = ternario.NewDefault(localctx.(*Instr_default_terContext).Get_expressions().GetP())

	}

	return localctx
}

// IInstr_whileContext is an interface to support dynamic dispatch.
type IInstr_whileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_WHILE returns the _R_WHILE token.
	Get_R_WHILE() antlr.Token

	// Set_R_WHILE sets the _R_WHILE token.
	Set_R_WHILE(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_whileContext differentiates from other interfaces.
	IsInstr_whileContext()
}

type Instr_whileContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_WHILE       antlr.Token
	_expressions   IExpressionsContext
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_whileContext() *Instr_whileContext {
	var p = new(Instr_whileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_while
	return p
}

func (*Instr_whileContext) IsInstr_whileContext() {}

func NewInstr_whileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_whileContext {
	var p = new(Instr_whileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_while

	return p
}

func (s *Instr_whileContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_whileContext) Get_R_WHILE() antlr.Token { return s._R_WHILE }

func (s *Instr_whileContext) Set_R_WHILE(v antlr.Token) { s._R_WHILE = v }

func (s *Instr_whileContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_whileContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_whileContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_whileContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_whileContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_whileContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_whileContext) R_WHILE() antlr.TerminalNode {
	return s.GetToken(ChemsR_WHILE, 0)
}

func (s *Instr_whileContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_whileContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_whileContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_whileContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_whileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_whileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_whileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_while(s)
	}
}

func (s *Instr_whileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_while(s)
	}
}

func (p *Chems) Instr_while() (localctx IInstr_whileContext) {
	localctx = NewInstr_whileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ChemsRULE_instr_while)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)

		var _m = p.Match(ChemsR_WHILE)

		localctx.(*Instr_whileContext)._R_WHILE = _m
	}
	{
		p.SetState(465)

		var _x = p.Expressions()

		localctx.(*Instr_whileContext)._expressions = _x
	}
	{
		p.SetState(466)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(467)

		var _x = p.Instrucciones()

		localctx.(*Instr_whileContext)._instrucciones = _x
	}
	{
		p.SetState(468)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_whileContext).instr = loops.NewWhile(localctx.(*Instr_whileContext).Get_expressions().GetP(), localctx.(*Instr_whileContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instr_whileContext).Get_R_WHILE() == nil {
			return 0
		} else {
			return localctx.(*Instr_whileContext).Get_R_WHILE().GetLine()
		}
	}()), localctx.(*Instr_whileContext).Get_R_WHILE().GetColumn())

	return localctx
}

// IInstr_for_inContext is an interface to support dynamic dispatch.
type IInstr_for_inContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_FOR returns the _R_FOR token.
	Get_R_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_R_FOR sets the _R_FOR token.
	Set_R_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionsContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionsContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionsContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_for_inContext differentiates from other interfaces.
	IsInstr_for_inContext()
}

type Instr_for_inContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_FOR         antlr.Token
	_ID            antlr.Token
	left           IExpressionsContext
	right          IExpressionsContext
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_for_inContext() *Instr_for_inContext {
	var p = new(Instr_for_inContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_for_in
	return p
}

func (*Instr_for_inContext) IsInstr_for_inContext() {}

func NewInstr_for_inContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_for_inContext {
	var p = new(Instr_for_inContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_for_in

	return p
}

func (s *Instr_for_inContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_for_inContext) Get_R_FOR() antlr.Token { return s._R_FOR }

func (s *Instr_for_inContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_for_inContext) Set_R_FOR(v antlr.Token) { s._R_FOR = v }

func (s *Instr_for_inContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_for_inContext) GetLeft() IExpressionsContext { return s.left }

func (s *Instr_for_inContext) GetRight() IExpressionsContext { return s.right }

func (s *Instr_for_inContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_for_inContext) SetLeft(v IExpressionsContext) { s.left = v }

func (s *Instr_for_inContext) SetRight(v IExpressionsContext) { s.right = v }

func (s *Instr_for_inContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_for_inContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_for_inContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_for_inContext) R_FOR() antlr.TerminalNode {
	return s.GetToken(ChemsR_FOR, 0)
}

func (s *Instr_for_inContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_for_inContext) R_IN() antlr.TerminalNode {
	return s.GetToken(ChemsR_IN, 0)
}

func (s *Instr_for_inContext) TK_DOBLEPUNTO() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DOBLEPUNTO, 0)
}

func (s *Instr_for_inContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_for_inContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_for_inContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_for_inContext) AllExpressions() []IExpressionsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionsContext)(nil)).Elem())
	var tst = make([]IExpressionsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionsContext)
		}
	}

	return tst
}

func (s *Instr_for_inContext) Expressions(i int) IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_for_inContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_for_inContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_for_inContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_for_in(s)
	}
}

func (s *Instr_for_inContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_for_in(s)
	}
}

func (p *Chems) Instr_for_in() (localctx IInstr_for_inContext) {
	localctx = NewInstr_for_inContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ChemsRULE_instr_for_in)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(471)

			var _m = p.Match(ChemsR_FOR)

			localctx.(*Instr_for_inContext)._R_FOR = _m
		}
		{
			p.SetState(472)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_for_inContext)._ID = _m
		}
		{
			p.SetState(473)
			p.Match(ChemsR_IN)
		}
		{
			p.SetState(474)

			var _x = p.Expressions()

			localctx.(*Instr_for_inContext).left = _x
		}
		{
			p.SetState(475)
			p.Match(ChemsTK_DOBLEPUNTO)
		}
		{
			p.SetState(476)

			var _x = p.Expressions()

			localctx.(*Instr_for_inContext).right = _x
		}
		{
			p.SetState(477)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(478)

			var _x = p.Instrucciones()

			localctx.(*Instr_for_inContext)._instrucciones = _x
		}
		{
			p.SetState(479)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_for_inContext).instr = loops.NewFor((func() string {
			if localctx.(*Instr_for_inContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_for_inContext).Get_ID().GetText()
			}
		}()), interfaces.INTEGER, localctx.(*Instr_for_inContext).GetLeft().GetP(), localctx.(*Instr_for_inContext).GetRight().GetP(), localctx.(*Instr_for_inContext).Get_instrucciones().GetL(), (func() int {
			if localctx.(*Instr_for_inContext).Get_R_FOR() == nil {
				return 0
			} else {
				return localctx.(*Instr_for_inContext).Get_R_FOR().GetLine()
			}
		}()), localctx.(*Instr_for_inContext).Get_R_FOR().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(482)

			var _m = p.Match(ChemsR_FOR)

			localctx.(*Instr_for_inContext)._R_FOR = _m
		}
		{
			p.SetState(483)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_for_inContext)._ID = _m
		}
		{
			p.SetState(484)
			p.Match(ChemsR_IN)
		}
		{
			p.SetState(485)

			var _x = p.Expressions()

			localctx.(*Instr_for_inContext).left = _x
		}
		{
			p.SetState(486)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(487)

			var _x = p.Instrucciones()

			localctx.(*Instr_for_inContext)._instrucciones = _x
		}
		{
			p.SetState(488)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_for_inContext).instr = loops.NewFor((func() string {
			if localctx.(*Instr_for_inContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_for_inContext).Get_ID().GetText()
			}
		}()), interfaces.STRING, localctx.(*Instr_for_inContext).GetLeft().GetP(), nil, localctx.(*Instr_for_inContext).Get_instrucciones().GetL(), (func() int {
			if localctx.(*Instr_for_inContext).Get_R_FOR() == nil {
				return 0
			} else {
				return localctx.(*Instr_for_inContext).Get_R_FOR().GetLine()
			}
		}()), localctx.(*Instr_for_inContext).Get_R_FOR().GetColumn())

	}

	return localctx
}

// IInstr_loopContext is an interface to support dynamic dispatch.
type IInstr_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_LOOP returns the _R_LOOP token.
	Get_R_LOOP() antlr.Token

	// Set_R_LOOP sets the _R_LOOP token.
	Set_R_LOOP(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_loopContext differentiates from other interfaces.
	IsInstr_loopContext()
}

type Instr_loopContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_LOOP        antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_loopContext() *Instr_loopContext {
	var p = new(Instr_loopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_loop
	return p
}

func (*Instr_loopContext) IsInstr_loopContext() {}

func NewInstr_loopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_loopContext {
	var p = new(Instr_loopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_loop

	return p
}

func (s *Instr_loopContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_loopContext) Get_R_LOOP() antlr.Token { return s._R_LOOP }

func (s *Instr_loopContext) Set_R_LOOP(v antlr.Token) { s._R_LOOP = v }

func (s *Instr_loopContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_loopContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_loopContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_loopContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_loopContext) R_LOOP() antlr.TerminalNode {
	return s.GetToken(ChemsR_LOOP, 0)
}

func (s *Instr_loopContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_loopContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_loopContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_loopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_loopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_loop(s)
	}
}

func (s *Instr_loopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_loop(s)
	}
}

func (p *Chems) Instr_loop() (localctx IInstr_loopContext) {
	localctx = NewInstr_loopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ChemsRULE_instr_loop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)

		var _m = p.Match(ChemsR_LOOP)

		localctx.(*Instr_loopContext)._R_LOOP = _m
	}
	{
		p.SetState(494)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(495)

		var _x = p.Instrucciones()

		localctx.(*Instr_loopContext)._instrucciones = _x
	}
	{
		p.SetState(496)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_loopContext).instr = loops.NewLoop(localctx.(*Instr_loopContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instr_loopContext).Get_R_LOOP() == nil {
			return 0
		} else {
			return localctx.(*Instr_loopContext).Get_R_LOOP().GetLine()
		}
	}()), localctx.(*Instr_loopContext).Get_R_LOOP().GetColumn())

	return localctx
}

// IInstr_loop_ternarioContext is an interface to support dynamic dispatch.
type IInstr_loop_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_LOOP returns the _R_LOOP token.
	Get_R_LOOP() antlr.Token

	// Set_R_LOOP sets the _R_LOOP token.
	Set_R_LOOP(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_loop_ternarioContext differentiates from other interfaces.
	IsInstr_loop_ternarioContext()
}

type Instr_loop_ternarioContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	p              interfaces.Expression
	_R_LOOP        antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_loop_ternarioContext() *Instr_loop_ternarioContext {
	var p = new(Instr_loop_ternarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_loop_ternario
	return p
}

func (*Instr_loop_ternarioContext) IsInstr_loop_ternarioContext() {}

func NewInstr_loop_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_loop_ternarioContext {
	var p = new(Instr_loop_ternarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_loop_ternario

	return p
}

func (s *Instr_loop_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_loop_ternarioContext) Get_R_LOOP() antlr.Token { return s._R_LOOP }

func (s *Instr_loop_ternarioContext) Set_R_LOOP(v antlr.Token) { s._R_LOOP = v }

func (s *Instr_loop_ternarioContext) Get_instrucciones() IInstruccionesContext {
	return s._instrucciones
}

func (s *Instr_loop_ternarioContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_loop_ternarioContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_loop_ternarioContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_loop_ternarioContext) R_LOOP() antlr.TerminalNode {
	return s.GetToken(ChemsR_LOOP, 0)
}

func (s *Instr_loop_ternarioContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_loop_ternarioContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_loop_ternarioContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_loop_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_loop_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_loop_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_loop_ternario(s)
	}
}

func (s *Instr_loop_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_loop_ternario(s)
	}
}

func (p *Chems) Instr_loop_ternario() (localctx IInstr_loop_ternarioContext) {
	localctx = NewInstr_loop_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ChemsRULE_instr_loop_ternario)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)

		var _m = p.Match(ChemsR_LOOP)

		localctx.(*Instr_loop_ternarioContext)._R_LOOP = _m
	}
	{
		p.SetState(500)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(501)

		var _x = p.Instrucciones()

		localctx.(*Instr_loop_ternarioContext)._instrucciones = _x
	}
	{
		p.SetState(502)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_loop_ternarioContext).p = ternario.NewLoop(localctx.(*Instr_loop_ternarioContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instr_loop_ternarioContext).Get_R_LOOP() == nil {
			return 0
		} else {
			return localctx.(*Instr_loop_ternarioContext).Get_R_LOOP().GetLine()
		}
	}()), localctx.(*Instr_loop_ternarioContext).Get_R_LOOP().GetColumn())

	return localctx
}

// IInstr_breakContext is an interface to support dynamic dispatch.
type IInstr_breakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_BREAK returns the _R_BREAK token.
	Get_R_BREAK() antlr.Token

	// Set_R_BREAK sets the _R_BREAK token.
	Set_R_BREAK(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_breakContext differentiates from other interfaces.
	IsInstr_breakContext()
}

type Instr_breakContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruction
	_R_BREAK     antlr.Token
	_expressions IExpressionsContext
}

func NewEmptyInstr_breakContext() *Instr_breakContext {
	var p = new(Instr_breakContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_break
	return p
}

func (*Instr_breakContext) IsInstr_breakContext() {}

func NewInstr_breakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_breakContext {
	var p = new(Instr_breakContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_break

	return p
}

func (s *Instr_breakContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_breakContext) Get_R_BREAK() antlr.Token { return s._R_BREAK }

func (s *Instr_breakContext) Set_R_BREAK(v antlr.Token) { s._R_BREAK = v }

func (s *Instr_breakContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_breakContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_breakContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_breakContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_breakContext) R_BREAK() antlr.TerminalNode {
	return s.GetToken(ChemsR_BREAK, 0)
}

func (s *Instr_breakContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_breakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_breakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_breakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_break(s)
	}
}

func (s *Instr_breakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_break(s)
	}
}

func (p *Chems) Instr_break() (localctx IInstr_breakContext) {
	localctx = NewInstr_breakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ChemsRULE_instr_break)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(505)

			var _m = p.Match(ChemsR_BREAK)

			localctx.(*Instr_breakContext)._R_BREAK = _m
		}
		localctx.(*Instr_breakContext).instr = transferencia.NewBreak(nil, (func() int {
			if localctx.(*Instr_breakContext).Get_R_BREAK() == nil {
				return 0
			} else {
				return localctx.(*Instr_breakContext).Get_R_BREAK().GetLine()
			}
		}()), localctx.(*Instr_breakContext).Get_R_BREAK().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(507)

			var _m = p.Match(ChemsR_BREAK)

			localctx.(*Instr_breakContext)._R_BREAK = _m
		}
		{
			p.SetState(508)

			var _x = p.Expressions()

			localctx.(*Instr_breakContext)._expressions = _x
		}
		localctx.(*Instr_breakContext).instr = transferencia.NewBreak(localctx.(*Instr_breakContext).Get_expressions().GetP(), (func() int {
			if localctx.(*Instr_breakContext).Get_R_BREAK() == nil {
				return 0
			} else {
				return localctx.(*Instr_breakContext).Get_R_BREAK().GetLine()
			}
		}()), localctx.(*Instr_breakContext).Get_R_BREAK().GetColumn())

	}

	return localctx
}

// IInstr_continueContext is an interface to support dynamic dispatch.
type IInstr_continueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_CONTINUE returns the _R_CONTINUE token.
	Get_R_CONTINUE() antlr.Token

	// Set_R_CONTINUE sets the _R_CONTINUE token.
	Set_R_CONTINUE(antlr.Token)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_continueContext differentiates from other interfaces.
	IsInstr_continueContext()
}

type Instr_continueContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_R_CONTINUE antlr.Token
}

func NewEmptyInstr_continueContext() *Instr_continueContext {
	var p = new(Instr_continueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_continue
	return p
}

func (*Instr_continueContext) IsInstr_continueContext() {}

func NewInstr_continueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_continueContext {
	var p = new(Instr_continueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_continue

	return p
}

func (s *Instr_continueContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_continueContext) Get_R_CONTINUE() antlr.Token { return s._R_CONTINUE }

func (s *Instr_continueContext) Set_R_CONTINUE(v antlr.Token) { s._R_CONTINUE = v }

func (s *Instr_continueContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_continueContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_continueContext) R_CONTINUE() antlr.TerminalNode {
	return s.GetToken(ChemsR_CONTINUE, 0)
}

func (s *Instr_continueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_continueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_continueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_continue(s)
	}
}

func (s *Instr_continueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_continue(s)
	}
}

func (p *Chems) Instr_continue() (localctx IInstr_continueContext) {
	localctx = NewInstr_continueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ChemsRULE_instr_continue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)

		var _m = p.Match(ChemsR_CONTINUE)

		localctx.(*Instr_continueContext)._R_CONTINUE = _m
	}
	localctx.(*Instr_continueContext).instr = transferencia.NewContinue((func() int {
		if localctx.(*Instr_continueContext).Get_R_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*Instr_continueContext).Get_R_CONTINUE().GetLine()
		}
	}()), localctx.(*Instr_continueContext).Get_R_CONTINUE().GetColumn())

	return localctx
}

// IInstr_funcContext is an interface to support dynamic dispatch.
type IInstr_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_FUNCTION returns the _R_FUNCTION token.
	Get_R_FUNCTION() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_R_FUNCTION sets the _R_FUNCTION token.
	Set_R_FUNCTION(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_list_function_parameters returns the _list_function_parameters rule contexts.
	Get_list_function_parameters() IList_function_parametersContext

	// Get_instr_tipo returns the _instr_tipo rule contexts.
	Get_instr_tipo() IInstr_tipoContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_list_function_parameters sets the _list_function_parameters rule contexts.
	Set_list_function_parameters(IList_function_parametersContext)

	// Set_instr_tipo sets the _instr_tipo rule contexts.
	Set_instr_tipo(IInstr_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_funcContext differentiates from other interfaces.
	IsInstr_funcContext()
}

type Instr_funcContext struct {
	*antlr.BaseParserRuleContext
	parser                    antlr.Parser
	instr                     interfaces.Instruction
	_R_FUNCTION               antlr.Token
	_ID                       antlr.Token
	_instrucciones            IInstruccionesContext
	_list_function_parameters IList_function_parametersContext
	_instr_tipo               IInstr_tipoContext
}

func NewEmptyInstr_funcContext() *Instr_funcContext {
	var p = new(Instr_funcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_func
	return p
}

func (*Instr_funcContext) IsInstr_funcContext() {}

func NewInstr_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_funcContext {
	var p = new(Instr_funcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_func

	return p
}

func (s *Instr_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_funcContext) Get_R_FUNCTION() antlr.Token { return s._R_FUNCTION }

func (s *Instr_funcContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_funcContext) Set_R_FUNCTION(v antlr.Token) { s._R_FUNCTION = v }

func (s *Instr_funcContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_funcContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_funcContext) Get_list_function_parameters() IList_function_parametersContext {
	return s._list_function_parameters
}

func (s *Instr_funcContext) Get_instr_tipo() IInstr_tipoContext { return s._instr_tipo }

func (s *Instr_funcContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_funcContext) Set_list_function_parameters(v IList_function_parametersContext) {
	s._list_function_parameters = v
}

func (s *Instr_funcContext) Set_instr_tipo(v IInstr_tipoContext) { s._instr_tipo = v }

func (s *Instr_funcContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_funcContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_funcContext) R_FUNCTION() antlr.TerminalNode {
	return s.GetToken(ChemsR_FUNCTION, 0)
}

func (s *Instr_funcContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_funcContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_funcContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_funcContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_funcContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_funcContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_funcContext) List_function_parameters() IList_function_parametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_function_parametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_function_parametersContext)
}

func (s *Instr_funcContext) TK_MENOSMAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOSMAYOR, 0)
}

func (s *Instr_funcContext) Instr_tipo() IInstr_tipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_tipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_tipoContext)
}

func (s *Instr_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_funcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_func(s)
	}
}

func (s *Instr_funcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_func(s)
	}
}

func (p *Chems) Instr_func() (localctx IInstr_funcContext) {
	localctx = NewInstr_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ChemsRULE_instr_func)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(516)

			var _m = p.Match(ChemsR_FUNCTION)

			localctx.(*Instr_funcContext)._R_FUNCTION = _m
		}
		{
			p.SetState(517)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_funcContext)._ID = _m
		}
		{
			p.SetState(518)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(519)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(520)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(521)

			var _x = p.Instrucciones()

			localctx.(*Instr_funcContext)._instrucciones = _x
		}
		{
			p.SetState(522)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_funcContext).instr = function.NewFunction((func() string {
			if localctx.(*Instr_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_funcContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*Instr_funcContext).Get_instrucciones().GetL(), interfaces.NULL, (func() int {
			if localctx.(*Instr_funcContext).Get_R_FUNCTION() == nil {
				return 0
			} else {
				return localctx.(*Instr_funcContext).Get_R_FUNCTION().GetLine()
			}
		}()), localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(525)

			var _m = p.Match(ChemsR_FUNCTION)

			localctx.(*Instr_funcContext)._R_FUNCTION = _m
		}
		{
			p.SetState(526)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_funcContext)._ID = _m
		}
		{
			p.SetState(527)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(528)

			var _x = p.List_function_parameters()

			localctx.(*Instr_funcContext)._list_function_parameters = _x
		}
		{
			p.SetState(529)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(530)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(531)

			var _x = p.Instrucciones()

			localctx.(*Instr_funcContext)._instrucciones = _x
		}
		{
			p.SetState(532)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_funcContext).instr = function.NewFunction((func() string {
			if localctx.(*Instr_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_funcContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_funcContext).Get_list_function_parameters().GetL(), localctx.(*Instr_funcContext).Get_instrucciones().GetL(), interfaces.NULL, (func() int {
			if localctx.(*Instr_funcContext).Get_R_FUNCTION() == nil {
				return 0
			} else {
				return localctx.(*Instr_funcContext).Get_R_FUNCTION().GetLine()
			}
		}()), localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(535)

			var _m = p.Match(ChemsR_FUNCTION)

			localctx.(*Instr_funcContext)._R_FUNCTION = _m
		}
		{
			p.SetState(536)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_funcContext)._ID = _m
		}
		{
			p.SetState(537)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(538)

			var _x = p.List_function_parameters()

			localctx.(*Instr_funcContext)._list_function_parameters = _x
		}
		{
			p.SetState(539)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(540)
			p.Match(ChemsTK_MENOSMAYOR)
		}
		{
			p.SetState(541)

			var _x = p.Instr_tipo()

			localctx.(*Instr_funcContext)._instr_tipo = _x
		}
		{
			p.SetState(542)
			p.Match(ChemsTK_LLAVEA)
		}
		{
			p.SetState(543)

			var _x = p.Instrucciones()

			localctx.(*Instr_funcContext)._instrucciones = _x
		}
		{
			p.SetState(544)
			p.Match(ChemsTK_LLAVEC)
		}
		localctx.(*Instr_funcContext).instr = function.NewFunction((func() string {
			if localctx.(*Instr_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_funcContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_funcContext).Get_list_function_parameters().GetL(), localctx.(*Instr_funcContext).Get_instrucciones().GetL(), localctx.(*Instr_funcContext).Get_instr_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Instr_funcContext).Get_R_FUNCTION() == nil {
				return 0
			} else {
				return localctx.(*Instr_funcContext).Get_R_FUNCTION().GetLine()
			}
		}()), localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn())

	}

	return localctx
}

// IList_function_parametersContext is an interface to support dynamic dispatch.
type IList_function_parametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_parameters_fn returns the _block_parameters_fn rule contexts.
	Get_block_parameters_fn() IBlock_parameters_fnContext

	// Set_block_parameters_fn sets the _block_parameters_fn rule contexts.
	Set_block_parameters_fn(IBlock_parameters_fnContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_parameters_fnContext

	// SetE sets the e rule context list.
	SetE([]IBlock_parameters_fnContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_function_parametersContext differentiates from other interfaces.
	IsList_function_parametersContext()
}

type List_function_parametersContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	l                    *arrayList.List
	_block_parameters_fn IBlock_parameters_fnContext
	e                    []IBlock_parameters_fnContext
}

func NewEmptyList_function_parametersContext() *List_function_parametersContext {
	var p = new(List_function_parametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_function_parameters
	return p
}

func (*List_function_parametersContext) IsList_function_parametersContext() {}

func NewList_function_parametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_function_parametersContext {
	var p = new(List_function_parametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_function_parameters

	return p
}

func (s *List_function_parametersContext) GetParser() antlr.Parser { return s.parser }

func (s *List_function_parametersContext) Get_block_parameters_fn() IBlock_parameters_fnContext {
	return s._block_parameters_fn
}

func (s *List_function_parametersContext) Set_block_parameters_fn(v IBlock_parameters_fnContext) {
	s._block_parameters_fn = v
}

func (s *List_function_parametersContext) GetE() []IBlock_parameters_fnContext { return s.e }

func (s *List_function_parametersContext) SetE(v []IBlock_parameters_fnContext) { s.e = v }

func (s *List_function_parametersContext) GetL() *arrayList.List { return s.l }

func (s *List_function_parametersContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_function_parametersContext) AllBlock_parameters_fn() []IBlock_parameters_fnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_parameters_fnContext)(nil)).Elem())
	var tst = make([]IBlock_parameters_fnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_parameters_fnContext)
		}
	}

	return tst
}

func (s *List_function_parametersContext) Block_parameters_fn(i int) IBlock_parameters_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_parameters_fnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_parameters_fnContext)
}

func (s *List_function_parametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_function_parametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_function_parametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_function_parameters(s)
	}
}

func (s *List_function_parametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_function_parameters(s)
	}
}

func (p *Chems) List_function_parameters() (localctx IList_function_parametersContext) {
	localctx = NewList_function_parametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ChemsRULE_list_function_parameters)

	localctx.(*List_function_parametersContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ChemsID {
		{
			p.SetState(549)

			var _x = p.Block_parameters_fn()

			localctx.(*List_function_parametersContext)._block_parameters_fn = _x
		}
		localctx.(*List_function_parametersContext).e = append(localctx.(*List_function_parametersContext).e, localctx.(*List_function_parametersContext)._block_parameters_fn)

		p.SetState(554)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_function_parametersContext).GetE()
	for _, e := range listInt {
		localctx.(*List_function_parametersContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IBlock_parameters_fnContext is an interface to support dynamic dispatch.
type IBlock_parameters_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_instr_tipo returns the _instr_tipo rule contexts.
	Get_instr_tipo() IInstr_tipoContext

	// Set_instr_tipo sets the _instr_tipo rule contexts.
	Set_instr_tipo(IInstr_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsBlock_parameters_fnContext differentiates from other interfaces.
	IsBlock_parameters_fnContext()
}

type Block_parameters_fnContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfaces.Instruction
	_ID         antlr.Token
	_instr_tipo IInstr_tipoContext
}

func NewEmptyBlock_parameters_fnContext() *Block_parameters_fnContext {
	var p = new(Block_parameters_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_parameters_fn
	return p
}

func (*Block_parameters_fnContext) IsBlock_parameters_fnContext() {}

func NewBlock_parameters_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_parameters_fnContext {
	var p = new(Block_parameters_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_parameters_fn

	return p
}

func (s *Block_parameters_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_parameters_fnContext) Get_ID() antlr.Token { return s._ID }

func (s *Block_parameters_fnContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Block_parameters_fnContext) Get_instr_tipo() IInstr_tipoContext { return s._instr_tipo }

func (s *Block_parameters_fnContext) Set_instr_tipo(v IInstr_tipoContext) { s._instr_tipo = v }

func (s *Block_parameters_fnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Block_parameters_fnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Block_parameters_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Block_parameters_fnContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DOSPUNTOS, 0)
}

func (s *Block_parameters_fnContext) Instr_tipo() IInstr_tipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_tipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_tipoContext)
}

func (s *Block_parameters_fnContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Block_parameters_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_parameters_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_parameters_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_parameters_fn(s)
	}
}

func (s *Block_parameters_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_parameters_fn(s)
	}
}

func (p *Chems) Block_parameters_fn() (localctx IBlock_parameters_fnContext) {
	localctx = NewBlock_parameters_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ChemsRULE_block_parameters_fn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(568)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(557)

			var _m = p.Match(ChemsID)

			localctx.(*Block_parameters_fnContext)._ID = _m
		}
		{
			p.SetState(558)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(559)

			var _x = p.Instr_tipo()

			localctx.(*Block_parameters_fnContext)._instr_tipo = _x
		}
		{
			p.SetState(560)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Block_parameters_fnContext).instr = function.NewListExpre((func() string {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetText()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_instr_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetLine()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_ID().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(563)

			var _m = p.Match(ChemsID)

			localctx.(*Block_parameters_fnContext)._ID = _m
		}
		{
			p.SetState(564)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(565)

			var _x = p.Instr_tipo()

			localctx.(*Block_parameters_fnContext)._instr_tipo = _x
		}
		localctx.(*Block_parameters_fnContext).instr = function.NewListExpre((func() string {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetText()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_instr_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetLine()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_ID().GetColumn())

	}

	return localctx
}

// IInstr_llamadaContext is an interface to support dynamic dispatch.
type IInstr_llamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_list_expression returns the _list_expression rule contexts.
	Get_list_expression() IList_expressionContext

	// Set_list_expression sets the _list_expression rule contexts.
	Set_list_expression(IList_expressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_llamadaContext differentiates from other interfaces.
	IsInstr_llamadaContext()
}

type Instr_llamadaContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            interfaces.Instruction
	_ID              antlr.Token
	_list_expression IList_expressionContext
}

func NewEmptyInstr_llamadaContext() *Instr_llamadaContext {
	var p = new(Instr_llamadaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_llamada
	return p
}

func (*Instr_llamadaContext) IsInstr_llamadaContext() {}

func NewInstr_llamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_llamadaContext {
	var p = new(Instr_llamadaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_llamada

	return p
}

func (s *Instr_llamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_llamadaContext) Get_ID() antlr.Token { return s._ID }

func (s *Instr_llamadaContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instr_llamadaContext) Get_list_expression() IList_expressionContext {
	return s._list_expression
}

func (s *Instr_llamadaContext) Set_list_expression(v IList_expressionContext) { s._list_expression = v }

func (s *Instr_llamadaContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_llamadaContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_llamadaContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Instr_llamadaContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_llamadaContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_llamadaContext) List_expression() IList_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Instr_llamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_llamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_llamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_llamada(s)
	}
}

func (s *Instr_llamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_llamada(s)
	}
}

func (p *Chems) Instr_llamada() (localctx IInstr_llamadaContext) {
	localctx = NewInstr_llamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ChemsRULE_instr_llamada)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(570)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_llamadaContext)._ID = _m
		}
		{
			p.SetState(571)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(572)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Instr_llamadaContext).instr = function.NewLlamada((func() string {
			if localctx.(*Instr_llamadaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_llamadaContext).Get_ID().GetText()
			}
		}()), nil, (func() int {
			if localctx.(*Instr_llamadaContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Instr_llamadaContext).Get_ID().GetLine()
			}
		}()), localctx.(*Instr_llamadaContext).Get_ID().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(574)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_llamadaContext)._ID = _m
		}
		{
			p.SetState(575)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(576)

			var _x = p.List_expression()

			localctx.(*Instr_llamadaContext)._list_expression = _x
		}
		{
			p.SetState(577)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Instr_llamadaContext).instr = function.NewLlamada((func() string {
			if localctx.(*Instr_llamadaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instr_llamadaContext).Get_ID().GetText()
			}
		}()), localctx.(*Instr_llamadaContext).Get_list_expression().GetL(), (func() int {
			if localctx.(*Instr_llamadaContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Instr_llamadaContext).Get_ID().GetLine()
			}
		}()), localctx.(*Instr_llamadaContext).Get_ID().GetColumn())

	}

	return localctx
}

// IInstr_tipoContext is an interface to support dynamic dispatch.
type IInstr_tipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipo_exp returns the tipo_exp attribute.
	GetTipo_exp() interfaces.TypeExpression

	// SetTipo_exp sets the tipo_exp attribute.
	SetTipo_exp(interfaces.TypeExpression)

	// IsInstr_tipoContext differentiates from other interfaces.
	IsInstr_tipoContext()
}

type Instr_tipoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	tipo_exp interfaces.TypeExpression
}

func NewEmptyInstr_tipoContext() *Instr_tipoContext {
	var p = new(Instr_tipoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_tipo
	return p
}

func (*Instr_tipoContext) IsInstr_tipoContext() {}

func NewInstr_tipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_tipoContext {
	var p = new(Instr_tipoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_tipo

	return p
}

func (s *Instr_tipoContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_tipoContext) GetTipo_exp() interfaces.TypeExpression { return s.tipo_exp }

func (s *Instr_tipoContext) SetTipo_exp(v interfaces.TypeExpression) { s.tipo_exp = v }

func (s *Instr_tipoContext) R_INT() antlr.TerminalNode {
	return s.GetToken(ChemsR_INT, 0)
}

func (s *Instr_tipoContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsR_FLOAT, 0)
}

func (s *Instr_tipoContext) R_STRING() antlr.TerminalNode {
	return s.GetToken(ChemsR_STRING, 0)
}

func (s *Instr_tipoContext) R_STR() antlr.TerminalNode {
	return s.GetToken(ChemsR_STR, 0)
}

func (s *Instr_tipoContext) R_BOOL() antlr.TerminalNode {
	return s.GetToken(ChemsR_BOOL, 0)
}

func (s *Instr_tipoContext) R_CHAR() antlr.TerminalNode {
	return s.GetToken(ChemsR_CHAR, 0)
}

func (s *Instr_tipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_tipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_tipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_tipo(s)
	}
}

func (s *Instr_tipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_tipo(s)
	}
}

func (p *Chems) Instr_tipo() (localctx IInstr_tipoContext) {
	localctx = NewInstr_tipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ChemsRULE_instr_tipo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(594)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(582)
			p.Match(ChemsR_INT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.INTEGER

	case ChemsR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(584)
			p.Match(ChemsR_FLOAT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.FLOAT

	case ChemsR_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(586)
			p.Match(ChemsR_STRING)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_STR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(588)
			p.Match(ChemsR_STR)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STR

	case ChemsR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(590)
			p.Match(ChemsR_BOOL)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.BOOLEAN

	case ChemsR_CHAR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(592)
			p.Match(ChemsR_CHAR)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.CHAR

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IList_expressionContext is an interface to support dynamic dispatch.
type IList_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_list_expression returns the _block_list_expression rule contexts.
	Get_block_list_expression() IBlock_list_expressionContext

	// Set_block_list_expression sets the _block_list_expression rule contexts.
	Set_block_list_expression(IBlock_list_expressionContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_list_expressionContext

	// SetE sets the e rule context list.
	SetE([]IBlock_list_expressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_expressionContext differentiates from other interfaces.
	IsList_expressionContext()
}

type List_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser                 antlr.Parser
	l                      *arrayList.List
	_block_list_expression IBlock_list_expressionContext
	e                      []IBlock_list_expressionContext
}

func NewEmptyList_expressionContext() *List_expressionContext {
	var p = new(List_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_list_expression
	return p
}

func (*List_expressionContext) IsList_expressionContext() {}

func NewList_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expressionContext {
	var p = new(List_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_list_expression

	return p
}

func (s *List_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expressionContext) Get_block_list_expression() IBlock_list_expressionContext {
	return s._block_list_expression
}

func (s *List_expressionContext) Set_block_list_expression(v IBlock_list_expressionContext) {
	s._block_list_expression = v
}

func (s *List_expressionContext) GetE() []IBlock_list_expressionContext { return s.e }

func (s *List_expressionContext) SetE(v []IBlock_list_expressionContext) { s.e = v }

func (s *List_expressionContext) GetL() *arrayList.List { return s.l }

func (s *List_expressionContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_expressionContext) AllBlock_list_expression() []IBlock_list_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlock_list_expressionContext)(nil)).Elem())
	var tst = make([]IBlock_list_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlock_list_expressionContext)
		}
	}

	return tst
}

func (s *List_expressionContext) Block_list_expression(i int) IBlock_list_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_list_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlock_list_expressionContext)
}

func (s *List_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterList_expression(s)
	}
}

func (s *List_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitList_expression(s)
	}
}

func (p *Chems) List_expression() (localctx IList_expressionContext) {
	localctx = NewList_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ChemsRULE_list_expression)

	localctx.(*List_expressionContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(597)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_IF)|(1<<ChemsR_MATCH)|(1<<ChemsR_LOOP)|(1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsCHAR)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(ChemsTK_MENOS-48))|(1<<(ChemsTK_PARA-48))|(1<<(ChemsTK_NOT-48)))) != 0) {
		{
			p.SetState(596)

			var _x = p.Block_list_expression()

			localctx.(*List_expressionContext)._block_list_expression = _x
		}
		localctx.(*List_expressionContext).e = append(localctx.(*List_expressionContext).e, localctx.(*List_expressionContext)._block_list_expression)

		p.SetState(599)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expressionContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expressionContext).l.Add(e.GetP())
	}

	return localctx
}

// IBlock_list_expressionContext is an interface to support dynamic dispatch.
type IBlock_list_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsBlock_list_expressionContext differentiates from other interfaces.
	IsBlock_list_expressionContext()
}

type Block_list_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	_expressions IExpressionsContext
}

func NewEmptyBlock_list_expressionContext() *Block_list_expressionContext {
	var p = new(Block_list_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_block_list_expression
	return p
}

func (*Block_list_expressionContext) IsBlock_list_expressionContext() {}

func NewBlock_list_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_list_expressionContext {
	var p = new(Block_list_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_block_list_expression

	return p
}

func (s *Block_list_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_list_expressionContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Block_list_expressionContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Block_list_expressionContext) GetP() interfaces.Expression { return s.p }

func (s *Block_list_expressionContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Block_list_expressionContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_list_expressionContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Block_list_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_list_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_list_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBlock_list_expression(s)
	}
}

func (s *Block_list_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBlock_list_expression(s)
	}
}

func (p *Chems) Block_list_expression() (localctx IBlock_list_expressionContext) {
	localctx = NewBlock_list_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ChemsRULE_block_list_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(603)

			var _x = p.Expressions()

			localctx.(*Block_list_expressionContext)._expressions = _x
		}
		{
			p.SetState(604)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Block_list_expressionContext).p = instruction.NewListExpre(localctx.(*Block_list_expressionContext).Get_expressions().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(607)

			var _x = p.Expressions()

			localctx.(*Block_list_expressionContext)._expressions = _x
		}
		localctx.(*Block_list_expressionContext).p = instruction.NewListExpre(localctx.(*Block_list_expressionContext).Get_expressions().GetP())

	}

	return localctx
}

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expre_log returns the _expre_log rule contexts.
	Get_expre_log() IExpre_logContext

	// Get_expre_rel returns the _expre_rel rule contexts.
	Get_expre_rel() IExpre_relContext

	// Get_expre_arit returns the _expre_arit rule contexts.
	Get_expre_arit() IExpre_aritContext

	// Get_expre_casteo returns the _expre_casteo rule contexts.
	Get_expre_casteo() IExpre_casteoContext

	// Set_expre_log sets the _expre_log rule contexts.
	Set_expre_log(IExpre_logContext)

	// Set_expre_rel sets the _expre_rel rule contexts.
	Set_expre_rel(IExpre_relContext)

	// Set_expre_arit sets the _expre_arit rule contexts.
	Set_expre_arit(IExpre_aritContext)

	// Set_expre_casteo sets the _expre_casteo rule contexts.
	Set_expre_casteo(IExpre_casteoContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	p             interfaces.Expression
	_expre_log    IExpre_logContext
	_expre_rel    IExpre_relContext
	_expre_arit   IExpre_aritContext
	_expre_casteo IExpre_casteoContext
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) Get_expre_log() IExpre_logContext { return s._expre_log }

func (s *ExpressionsContext) Get_expre_rel() IExpre_relContext { return s._expre_rel }

func (s *ExpressionsContext) Get_expre_arit() IExpre_aritContext { return s._expre_arit }

func (s *ExpressionsContext) Get_expre_casteo() IExpre_casteoContext { return s._expre_casteo }

func (s *ExpressionsContext) Set_expre_log(v IExpre_logContext) { s._expre_log = v }

func (s *ExpressionsContext) Set_expre_rel(v IExpre_relContext) { s._expre_rel = v }

func (s *ExpressionsContext) Set_expre_arit(v IExpre_aritContext) { s._expre_arit = v }

func (s *ExpressionsContext) Set_expre_casteo(v IExpre_casteoContext) { s._expre_casteo = v }

func (s *ExpressionsContext) GetP() interfaces.Expression { return s.p }

func (s *ExpressionsContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ExpressionsContext) Expre_log() IExpre_logContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_logContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_logContext)
}

func (s *ExpressionsContext) Expre_rel() IExpre_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_relContext)
}

func (s *ExpressionsContext) Expre_arit() IExpre_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *ExpressionsContext) Expre_casteo() IExpre_casteoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_casteoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_casteoContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *Chems) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ChemsRULE_expressions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(624)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(612)

			var _x = p.expre_log(0)

			localctx.(*ExpressionsContext)._expre_log = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_log().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(615)

			var _x = p.expre_rel(0)

			localctx.(*ExpressionsContext)._expre_rel = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_rel().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(618)

			var _x = p.expre_arit(0)

			localctx.(*ExpressionsContext)._expre_arit = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_arit().GetP()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(621)

			var _x = p.Expre_casteo()

			localctx.(*ExpressionsContext)._expre_casteo = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_casteo().GetP()

	}

	return localctx
}

// IExpre_logContext is an interface to support dynamic dispatch.
type IExpre_logContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpre_logContext

	// Get_expre_rel returns the _expre_rel rule contexts.
	Get_expre_rel() IExpre_relContext

	// GetRight returns the right rule contexts.
	GetRight() IExpre_logContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpre_logContext)

	// Set_expre_rel sets the _expre_rel rule contexts.
	Set_expre_rel(IExpre_relContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpre_logContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpre_logContext differentiates from other interfaces.
	IsExpre_logContext()
}

type Expre_logContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          interfaces.Expression
	left       IExpre_logContext
	op         antlr.Token
	_expre_rel IExpre_relContext
	right      IExpre_logContext
}

func NewEmptyExpre_logContext() *Expre_logContext {
	var p = new(Expre_logContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expre_log
	return p
}

func (*Expre_logContext) IsExpre_logContext() {}

func NewExpre_logContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_logContext {
	var p = new(Expre_logContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expre_log

	return p
}

func (s *Expre_logContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_logContext) GetOp() antlr.Token { return s.op }

func (s *Expre_logContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expre_logContext) GetLeft() IExpre_logContext { return s.left }

func (s *Expre_logContext) Get_expre_rel() IExpre_relContext { return s._expre_rel }

func (s *Expre_logContext) GetRight() IExpre_logContext { return s.right }

func (s *Expre_logContext) SetLeft(v IExpre_logContext) { s.left = v }

func (s *Expre_logContext) Set_expre_rel(v IExpre_relContext) { s._expre_rel = v }

func (s *Expre_logContext) SetRight(v IExpre_logContext) { s.right = v }

func (s *Expre_logContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_logContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_logContext) TK_NOT() antlr.TerminalNode {
	return s.GetToken(ChemsTK_NOT, 0)
}

func (s *Expre_logContext) AllExpre_log() []IExpre_logContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpre_logContext)(nil)).Elem())
	var tst = make([]IExpre_logContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpre_logContext)
		}
	}

	return tst
}

func (s *Expre_logContext) Expre_log(i int) IExpre_logContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_logContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpre_logContext)
}

func (s *Expre_logContext) Expre_rel() IExpre_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_relContext)
}

func (s *Expre_logContext) TK_AND() antlr.TerminalNode {
	return s.GetToken(ChemsTK_AND, 0)
}

func (s *Expre_logContext) TK_OR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_OR, 0)
}

func (s *Expre_logContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_logContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_logContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpre_log(s)
	}
}

func (s *Expre_logContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpre_log(s)
	}
}

func (p *Chems) Expre_log() (localctx IExpre_logContext) {
	return p.expre_log(0)
}

func (p *Chems) expre_log(_p int) (localctx IExpre_logContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpre_logContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_logContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 80
	p.EnterRecursionRule(localctx, 80, ChemsRULE_expre_log, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(634)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_NOT:
		{
			p.SetState(627)

			var _m = p.Match(ChemsTK_NOT)

			localctx.(*Expre_logContext).op = _m
		}
		{
			p.SetState(628)

			var _x = p.expre_log(3)

			localctx.(*Expre_logContext).left = _x
		}
		localctx.(*Expre_logContext).p = expression.NewOperacion(localctx.(*Expre_logContext).GetLeft().GetP(), (func() string {
			if localctx.(*Expre_logContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Expre_logContext).GetOp().GetText()
			}
		}()), nil, true, (func() int {
			if localctx.(*Expre_logContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Expre_logContext).GetOp().GetLine()
			}
		}()), localctx.(*Expre_logContext).GetOp().GetColumn())

	case ChemsR_IF, ChemsR_MATCH, ChemsR_LOOP, ChemsNUMBER, ChemsDOUBLE, ChemsCHAR, ChemsSTRING, ChemsBOOLEAN, ChemsID, ChemsTK_MENOS, ChemsTK_PARA:
		{
			p.SetState(631)

			var _x = p.expre_rel(0)

			localctx.(*Expre_logContext)._expre_rel = _x
		}
		localctx.(*Expre_logContext).p = localctx.(*Expre_logContext).Get_expre_rel().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(643)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpre_logContext(p, _parentctx, _parentState)
			localctx.(*Expre_logContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_log)
			p.SetState(636)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(637)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expre_logContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ChemsTK_AND || _la == ChemsTK_OR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expre_logContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(638)

				var _x = p.expre_log(3)

				localctx.(*Expre_logContext).right = _x
			}
			localctx.(*Expre_logContext).p = expression.NewOperacion(localctx.(*Expre_logContext).GetLeft().GetP(), (func() string {
				if localctx.(*Expre_logContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_logContext).GetOp().GetText()
				}
			}()), localctx.(*Expre_logContext).GetRight().GetP(), false, (func() int {
				if localctx.(*Expre_logContext).GetOp() == nil {
					return 0
				} else {
					return localctx.(*Expre_logContext).GetOp().GetLine()
				}
			}()), localctx.(*Expre_logContext).GetOp().GetColumn())

		}
		p.SetState(645)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}

	return localctx
}

// IExpre_relContext is an interface to support dynamic dispatch.
type IExpre_relContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpre_relContext

	// Get_expre_arit returns the _expre_arit rule contexts.
	Get_expre_arit() IExpre_aritContext

	// GetRight returns the right rule contexts.
	GetRight() IExpre_relContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpre_relContext)

	// Set_expre_arit sets the _expre_arit rule contexts.
	Set_expre_arit(IExpre_aritContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpre_relContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpre_relContext differentiates from other interfaces.
	IsExpre_relContext()
}

type Expre_relContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expression
	left        IExpre_relContext
	_expre_arit IExpre_aritContext
	op          antlr.Token
	right       IExpre_relContext
}

func NewEmptyExpre_relContext() *Expre_relContext {
	var p = new(Expre_relContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expre_rel
	return p
}

func (*Expre_relContext) IsExpre_relContext() {}

func NewExpre_relContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_relContext {
	var p = new(Expre_relContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expre_rel

	return p
}

func (s *Expre_relContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_relContext) GetOp() antlr.Token { return s.op }

func (s *Expre_relContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expre_relContext) GetLeft() IExpre_relContext { return s.left }

func (s *Expre_relContext) Get_expre_arit() IExpre_aritContext { return s._expre_arit }

func (s *Expre_relContext) GetRight() IExpre_relContext { return s.right }

func (s *Expre_relContext) SetLeft(v IExpre_relContext) { s.left = v }

func (s *Expre_relContext) Set_expre_arit(v IExpre_aritContext) { s._expre_arit = v }

func (s *Expre_relContext) SetRight(v IExpre_relContext) { s.right = v }

func (s *Expre_relContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_relContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_relContext) Expre_arit() IExpre_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *Expre_relContext) AllExpre_rel() []IExpre_relContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpre_relContext)(nil)).Elem())
	var tst = make([]IExpre_relContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpre_relContext)
		}
	}

	return tst
}

func (s *Expre_relContext) Expre_rel(i int) IExpre_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_relContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpre_relContext)
}

func (s *Expre_relContext) TK_MENOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOR, 0)
}

func (s *Expre_relContext) TK_MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENORIGUAL, 0)
}

func (s *Expre_relContext) TK_MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAYORIGUAL, 0)
}

func (s *Expre_relContext) TK_MAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAYOR, 0)
}

func (s *Expre_relContext) TK_DIFIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DIFIGUAL, 0)
}

func (s *Expre_relContext) TK_IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALIGUAL, 0)
}

func (s *Expre_relContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_relContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_relContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpre_rel(s)
	}
}

func (s *Expre_relContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpre_rel(s)
	}
}

func (p *Chems) Expre_rel() (localctx IExpre_relContext) {
	return p.expre_rel(0)
}

func (p *Chems) expre_rel(_p int) (localctx IExpre_relContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpre_relContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_relContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 82
	p.EnterRecursionRule(localctx, 82, ChemsRULE_expre_rel, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(647)

		var _x = p.expre_arit(0)

		localctx.(*Expre_relContext)._expre_arit = _x
	}
	localctx.(*Expre_relContext).p = localctx.(*Expre_relContext).Get_expre_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(657)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpre_relContext(p, _parentctx, _parentState)
			localctx.(*Expre_relContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_rel)
			p.SetState(650)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(651)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expre_relContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ChemsTK_IGUALIGUAL-36))|(1<<(ChemsTK_MAYORIGUAL-36))|(1<<(ChemsTK_MENORIGUAL-36))|(1<<(ChemsTK_DIFIGUAL-36))|(1<<(ChemsTK_MAYOR-36))|(1<<(ChemsTK_MENOR-36)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expre_relContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(652)

				var _x = p.expre_rel(3)

				localctx.(*Expre_relContext).right = _x
			}
			localctx.(*Expre_relContext).p = expression.NewOperacion(localctx.(*Expre_relContext).GetLeft().GetP(), (func() string {
				if localctx.(*Expre_relContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_relContext).GetOp().GetText()
				}
			}()), localctx.(*Expre_relContext).GetRight().GetP(), false, (func() int {
				if localctx.(*Expre_relContext).GetOp() == nil {
					return 0
				} else {
					return localctx.(*Expre_relContext).GetOp().GetLine()
				}
			}()), localctx.(*Expre_relContext).GetOp().GetColumn())

		}
		p.SetState(659)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}

	return localctx
}

// IExpre_aritContext is an interface to support dynamic dispatch.
type IExpre_aritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpre_aritContext

	// Get_expre_valor returns the _expre_valor rule contexts.
	Get_expre_valor() IExpre_valorContext

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// GetRight returns the right rule contexts.
	GetRight() IExpre_aritContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpre_aritContext)

	// Set_expre_valor sets the _expre_valor rule contexts.
	Set_expre_valor(IExpre_valorContext)

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpre_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpre_aritContext differentiates from other interfaces.
	IsExpre_aritContext()
}

type Expre_aritContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	left         IExpre_aritContext
	op           antlr.Token
	_expre_valor IExpre_valorContext
	_expressions IExpressionsContext
	right        IExpre_aritContext
}

func NewEmptyExpre_aritContext() *Expre_aritContext {
	var p = new(Expre_aritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expre_arit
	return p
}

func (*Expre_aritContext) IsExpre_aritContext() {}

func NewExpre_aritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_aritContext {
	var p = new(Expre_aritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expre_arit

	return p
}

func (s *Expre_aritContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expre_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expre_aritContext) GetLeft() IExpre_aritContext { return s.left }

func (s *Expre_aritContext) Get_expre_valor() IExpre_valorContext { return s._expre_valor }

func (s *Expre_aritContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Expre_aritContext) GetRight() IExpre_aritContext { return s.right }

func (s *Expre_aritContext) SetLeft(v IExpre_aritContext) { s.left = v }

func (s *Expre_aritContext) Set_expre_valor(v IExpre_valorContext) { s._expre_valor = v }

func (s *Expre_aritContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Expre_aritContext) SetRight(v IExpre_aritContext) { s.right = v }

func (s *Expre_aritContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_aritContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_aritContext) TK_MENOS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOS, 0)
}

func (s *Expre_aritContext) AllExpre_arit() []IExpre_aritContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpre_aritContext)(nil)).Elem())
	var tst = make([]IExpre_aritContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpre_aritContext)
		}
	}

	return tst
}

func (s *Expre_aritContext) Expre_arit(i int) IExpre_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *Expre_aritContext) Expre_valor() IExpre_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_valorContext)
}

func (s *Expre_aritContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Expre_aritContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Expre_aritContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Expre_aritContext) TK_MULT() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MULT, 0)
}

func (s *Expre_aritContext) TK_DIV() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DIV, 0)
}

func (s *Expre_aritContext) TK_MODULO() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MODULO, 0)
}

func (s *Expre_aritContext) TK_MAS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAS, 0)
}

func (s *Expre_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_aritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpre_arit(s)
	}
}

func (s *Expre_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpre_arit(s)
	}
}

func (p *Chems) Expre_arit() (localctx IExpre_aritContext) {
	return p.expre_arit(0)
}

func (p *Chems) expre_arit(_p int) (localctx IExpre_aritContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpre_aritContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_aritContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 84
	p.EnterRecursionRule(localctx, 84, ChemsRULE_expre_arit, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(673)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_MENOS:
		{
			p.SetState(661)

			var _m = p.Match(ChemsTK_MENOS)

			localctx.(*Expre_aritContext).op = _m
		}
		{
			p.SetState(662)

			var _x = p.expre_arit(5)

			localctx.(*Expre_aritContext).left = _x
		}
		localctx.(*Expre_aritContext).p = expression.NewOperacion(localctx.(*Expre_aritContext).GetLeft().GetP(), (func() string {
			if localctx.(*Expre_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Expre_aritContext).GetOp().GetText()
			}
		}()), nil, true, (func() int {
			if localctx.(*Expre_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Expre_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Expre_aritContext).GetOp().GetColumn())

	case ChemsR_IF, ChemsR_MATCH, ChemsR_LOOP, ChemsNUMBER, ChemsDOUBLE, ChemsCHAR, ChemsSTRING, ChemsBOOLEAN, ChemsID:
		{
			p.SetState(665)

			var _x = p.Expre_valor()

			localctx.(*Expre_aritContext)._expre_valor = _x
		}
		localctx.(*Expre_aritContext).p = localctx.(*Expre_aritContext).Get_expre_valor().GetP()

	case ChemsTK_PARA:
		{
			p.SetState(668)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(669)

			var _x = p.Expressions()

			localctx.(*Expre_aritContext)._expressions = _x
		}
		{
			p.SetState(670)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Expre_aritContext).p = localctx.(*Expre_aritContext).Get_expressions().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(687)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(685)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpre_aritContext(p, _parentctx, _parentState)
				localctx.(*Expre_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_arit)
				p.SetState(675)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(676)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expre_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(ChemsTK_MULT-44))|(1<<(ChemsTK_DIV-44))|(1<<(ChemsTK_MODULO-44)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expre_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(677)

					var _x = p.expre_arit(5)

					localctx.(*Expre_aritContext).right = _x
				}
				localctx.(*Expre_aritContext).p = expression.NewOperacion(localctx.(*Expre_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expre_aritContext).GetRight().GetP(), false, (func() int {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Expre_aritContext).GetOp().GetColumn())

			case 2:
				localctx = NewExpre_aritContext(p, _parentctx, _parentState)
				localctx.(*Expre_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_arit)
				p.SetState(680)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(681)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expre_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsTK_MAS || _la == ChemsTK_MENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expre_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(682)

					var _x = p.expre_arit(4)

					localctx.(*Expre_aritContext).right = _x
				}
				localctx.(*Expre_aritContext).p = expression.NewOperacion(localctx.(*Expre_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expre_aritContext).GetRight().GetP(), false, (func() int {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Expre_aritContext).GetOp().GetColumn())

			}

		}
		p.SetState(689)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// IExpre_valorContext is an interface to support dynamic dispatch.
type IExpre_valorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpre_valorContext differentiates from other interfaces.
	IsExpre_valorContext()
}

type Expre_valorContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          interfaces.Expression
	_primitivo IPrimitivoContext
}

func NewEmptyExpre_valorContext() *Expre_valorContext {
	var p = new(Expre_valorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expre_valor
	return p
}

func (*Expre_valorContext) IsExpre_valorContext() {}

func NewExpre_valorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_valorContext {
	var p = new(Expre_valorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expre_valor

	return p
}

func (s *Expre_valorContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_valorContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Expre_valorContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Expre_valorContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_valorContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_valorContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Expre_valorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_valorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_valorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpre_valor(s)
	}
}

func (s *Expre_valorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpre_valor(s)
	}
}

func (p *Chems) Expre_valor() (localctx IExpre_valorContext) {
	localctx = NewExpre_valorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ChemsRULE_expre_valor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(690)

		var _x = p.Primitivo()

		localctx.(*Expre_valorContext)._primitivo = _x
	}
	localctx.(*Expre_valorContext).p = localctx.(*Expre_valorContext).Get_primitivo().GetP()

	return localctx
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_DOUBLE returns the _DOUBLE token.
	Get_DOUBLE() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_BOOLEAN returns the _BOOLEAN token.
	Get_BOOLEAN() antlr.Token

	// Get_CHAR returns the _CHAR token.
	Get_CHAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_DOUBLE sets the _DOUBLE token.
	Set_DOUBLE(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_BOOLEAN sets the _BOOLEAN token.
	Set_BOOLEAN(antlr.Token)

	// Set_CHAR sets the _CHAR token.
	Set_CHAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_primitivo_casteo returns the _primitivo_casteo rule contexts.
	Get_primitivo_casteo() IPrimitivo_casteoContext

	// Get_instr_ternario returns the _instr_ternario rule contexts.
	Get_instr_ternario() IInstr_ternarioContext

	// Get_instr_match_ter returns the _instr_match_ter rule contexts.
	Get_instr_match_ter() IInstr_match_terContext

	// Get_instr_loop_ternario returns the _instr_loop_ternario rule contexts.
	Get_instr_loop_ternario() IInstr_loop_ternarioContext

	// Set_primitivo_casteo sets the _primitivo_casteo rule contexts.
	Set_primitivo_casteo(IPrimitivo_casteoContext)

	// Set_instr_ternario sets the _instr_ternario rule contexts.
	Set_instr_ternario(IInstr_ternarioContext)

	// Set_instr_match_ter sets the _instr_match_ter rule contexts.
	Set_instr_match_ter(IInstr_match_terContext)

	// Set_instr_loop_ternario sets the _instr_loop_ternario rule contexts.
	Set_instr_loop_ternario(IInstr_loop_ternarioContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	p                    interfaces.Expression
	_NUMBER              antlr.Token
	_DOUBLE              antlr.Token
	_STRING              antlr.Token
	_BOOLEAN             antlr.Token
	_CHAR                antlr.Token
	_ID                  antlr.Token
	_primitivo_casteo    IPrimitivo_casteoContext
	_instr_ternario      IInstr_ternarioContext
	_instr_match_ter     IInstr_match_terContext
	_instr_loop_ternario IInstr_loop_ternarioContext
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_primitivo
	return p
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *PrimitivoContext) Get_DOUBLE() antlr.Token { return s._DOUBLE }

func (s *PrimitivoContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitivoContext) Get_BOOLEAN() antlr.Token { return s._BOOLEAN }

func (s *PrimitivoContext) Get_CHAR() antlr.Token { return s._CHAR }

func (s *PrimitivoContext) Get_ID() antlr.Token { return s._ID }

func (s *PrimitivoContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitivoContext) Set_DOUBLE(v antlr.Token) { s._DOUBLE = v }

func (s *PrimitivoContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitivoContext) Set_BOOLEAN(v antlr.Token) { s._BOOLEAN = v }

func (s *PrimitivoContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *PrimitivoContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *PrimitivoContext) Get_primitivo_casteo() IPrimitivo_casteoContext { return s._primitivo_casteo }

func (s *PrimitivoContext) Get_instr_ternario() IInstr_ternarioContext { return s._instr_ternario }

func (s *PrimitivoContext) Get_instr_match_ter() IInstr_match_terContext { return s._instr_match_ter }

func (s *PrimitivoContext) Get_instr_loop_ternario() IInstr_loop_ternarioContext {
	return s._instr_loop_ternario
}

func (s *PrimitivoContext) Set_primitivo_casteo(v IPrimitivo_casteoContext) { s._primitivo_casteo = v }

func (s *PrimitivoContext) Set_instr_ternario(v IInstr_ternarioContext) { s._instr_ternario = v }

func (s *PrimitivoContext) Set_instr_match_ter(v IInstr_match_terContext) { s._instr_match_ter = v }

func (s *PrimitivoContext) Set_instr_loop_ternario(v IInstr_loop_ternarioContext) {
	s._instr_loop_ternario = v
}

func (s *PrimitivoContext) GetP() interfaces.Expression { return s.p }

func (s *PrimitivoContext) SetP(v interfaces.Expression) { s.p = v }

func (s *PrimitivoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChemsNUMBER, 0)
}

func (s *PrimitivoContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ChemsDOUBLE, 0)
}

func (s *PrimitivoContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *PrimitivoContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChemsBOOLEAN, 0)
}

func (s *PrimitivoContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ChemsCHAR, 0)
}

func (s *PrimitivoContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *PrimitivoContext) Primitivo_casteo() IPrimitivo_casteoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivo_casteoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivo_casteoContext)
}

func (s *PrimitivoContext) Instr_ternario() IInstr_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_ternarioContext)
}

func (s *PrimitivoContext) Instr_match_ter() IInstr_match_terContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_match_terContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_match_terContext)
}

func (s *PrimitivoContext) Instr_loop_ternario() IInstr_loop_ternarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_loop_ternarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_loop_ternarioContext)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *Chems) Primitivo() (localctx IPrimitivoContext) {
	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ChemsRULE_primitivo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(717)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(693)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*PrimitivoContext)._NUMBER = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}

		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(695)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*PrimitivoContext)._DOUBLE = _m
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(697)

			var _m = p.Match(ChemsSTRING)

			localctx.(*PrimitivoContext)._STRING = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(str, interfaces.STRING, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_STRING().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(699)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*PrimitivoContext)._BOOLEAN = _m
		}

		// str:= (func() string { if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText() }}())[1:len((func() string { if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText() }}()))-1]
		exp, _ := strconv.ParseBool((func() string {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText()
			}
		}()))
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(701)

			var _m = p.Match(ChemsCHAR)

			localctx.(*PrimitivoContext)._CHAR = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_CHAR() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CHAR().GetText()
			}
		}())[1]
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(string(str), interfaces.CHAR, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_CHAR() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_CHAR().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_CHAR().GetColumn())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(703)

			var _m = p.Match(ChemsID)

			localctx.(*PrimitivoContext)._ID = _m
		}
		localctx.(*PrimitivoContext).p = variable.NewIdentifier((func() string {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetText()
			}
		}()), (func() int {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_ID().GetColumn())

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(705)

			var _x = p.Primitivo_casteo()

			localctx.(*PrimitivoContext)._primitivo_casteo = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_primitivo_casteo().GetP()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(708)

			var _x = p.Instr_ternario()

			localctx.(*PrimitivoContext)._instr_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_ternario().GetP()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(711)

			var _x = p.Instr_match_ter()

			localctx.(*PrimitivoContext)._instr_match_ter = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_match_ter().GetP()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(714)

			var _x = p.Instr_loop_ternario()

			localctx.(*PrimitivoContext)._instr_loop_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instr_loop_ternario().GetP()

	}

	return localctx
}

// IPrimitivo_casteoContext is an interface to support dynamic dispatch.
type IPrimitivo_casteoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_DOUBLE returns the _DOUBLE token.
	Get_DOUBLE() antlr.Token

	// Get_BOOLEAN returns the _BOOLEAN token.
	Get_BOOLEAN() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_DOUBLE sets the _DOUBLE token.
	Set_DOUBLE(antlr.Token)

	// Set_BOOLEAN sets the _BOOLEAN token.
	Set_BOOLEAN(antlr.Token)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsPrimitivo_casteoContext differentiates from other interfaces.
	IsPrimitivo_casteoContext()
}

type Primitivo_casteoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	p        interfaces.Expression
	_NUMBER  antlr.Token
	_DOUBLE  antlr.Token
	_BOOLEAN antlr.Token
}

func NewEmptyPrimitivo_casteoContext() *Primitivo_casteoContext {
	var p = new(Primitivo_casteoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_primitivo_casteo
	return p
}

func (*Primitivo_casteoContext) IsPrimitivo_casteoContext() {}

func NewPrimitivo_casteoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primitivo_casteoContext {
	var p = new(Primitivo_casteoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_primitivo_casteo

	return p
}

func (s *Primitivo_casteoContext) GetParser() antlr.Parser { return s.parser }

func (s *Primitivo_casteoContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *Primitivo_casteoContext) Get_DOUBLE() antlr.Token { return s._DOUBLE }

func (s *Primitivo_casteoContext) Get_BOOLEAN() antlr.Token { return s._BOOLEAN }

func (s *Primitivo_casteoContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *Primitivo_casteoContext) Set_DOUBLE(v antlr.Token) { s._DOUBLE = v }

func (s *Primitivo_casteoContext) Set_BOOLEAN(v antlr.Token) { s._BOOLEAN = v }

func (s *Primitivo_casteoContext) GetP() interfaces.Expression { return s.p }

func (s *Primitivo_casteoContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Primitivo_casteoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChemsNUMBER, 0)
}

func (s *Primitivo_casteoContext) R_AS() antlr.TerminalNode {
	return s.GetToken(ChemsR_AS, 0)
}

func (s *Primitivo_casteoContext) R_INT() antlr.TerminalNode {
	return s.GetToken(ChemsR_INT, 0)
}

func (s *Primitivo_casteoContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsR_FLOAT, 0)
}

func (s *Primitivo_casteoContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ChemsDOUBLE, 0)
}

func (s *Primitivo_casteoContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChemsBOOLEAN, 0)
}

func (s *Primitivo_casteoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_casteoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primitivo_casteoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrimitivo_casteo(s)
	}
}

func (s *Primitivo_casteoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrimitivo_casteo(s)
	}
}

func (p *Chems) Primitivo_casteo() (localctx IPrimitivo_casteoContext) {
	localctx = NewPrimitivo_casteoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ChemsRULE_primitivo_casteo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(743)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(719)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*Primitivo_casteoContext)._NUMBER = _m
		}
		{
			p.SetState(720)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(721)
			p.Match(ChemsR_INT)
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*Primitivo_casteoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}

		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.INTEGER, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_NUMBER().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(723)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*Primitivo_casteoContext)._NUMBER = _m
		}
		{
			p.SetState(724)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(725)
			p.Match(ChemsR_FLOAT)
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*Primitivo_casteoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}

		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.FLOAT, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_NUMBER().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(727)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*Primitivo_casteoContext)._DOUBLE = _m
		}
		{
			p.SetState(728)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(729)
			p.Match(ChemsR_INT)
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*Primitivo_casteoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.INTEGER, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(731)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*Primitivo_casteoContext)._DOUBLE = _m
		}
		{
			p.SetState(732)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(733)
			p.Match(ChemsR_FLOAT)
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*Primitivo_casteoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.FLOAT, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(735)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*Primitivo_casteoContext)._BOOLEAN = _m
		}
		{
			p.SetState(736)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(737)
			p.Match(ChemsR_INT)
		}

		// str:= (func() string { if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText() }}())[1:len((func() string { if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText() }}()))-1]
		exp, _ := strconv.ParseBool((func() string {
			if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText()
			}
		}()))
		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.INTEGER, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetColumn())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(739)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*Primitivo_casteoContext)._BOOLEAN = _m
		}
		{
			p.SetState(740)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(741)
			p.Match(ChemsR_FLOAT)
		}

		// str:= (func() string { if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText() }}())[1:len((func() string { if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText() }}()))-1]
		exp, _ := strconv.ParseBool((func() string {
			if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetText()
			}
		}()))
		localctx.(*Primitivo_casteoContext).p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.FLOAT, (func() int {
			if localctx.(*Primitivo_casteoContext).Get_BOOLEAN() == nil {
				return 0
			} else {
				return localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetLine()
			}
		}()), localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetColumn())

	}

	return localctx
}

// IExpre_casteoContext is an interface to support dynamic dispatch.
type IExpre_casteoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_AS returns the _R_AS token.
	Get_R_AS() antlr.Token

	// Set_R_AS sets the _R_AS token.
	Set_R_AS(antlr.Token)

	// Get_expre_arit returns the _expre_arit rule contexts.
	Get_expre_arit() IExpre_aritContext

	// Get_type_casteo returns the _type_casteo rule contexts.
	Get_type_casteo() IType_casteoContext

	// Set_expre_arit sets the _expre_arit rule contexts.
	Set_expre_arit(IExpre_aritContext)

	// Set_type_casteo sets the _type_casteo rule contexts.
	Set_type_casteo(IType_casteoContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpre_casteoContext differentiates from other interfaces.
	IsExpre_casteoContext()
}

type Expre_casteoContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	_expre_arit  IExpre_aritContext
	_R_AS        antlr.Token
	_type_casteo IType_casteoContext
}

func NewEmptyExpre_casteoContext() *Expre_casteoContext {
	var p = new(Expre_casteoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expre_casteo
	return p
}

func (*Expre_casteoContext) IsExpre_casteoContext() {}

func NewExpre_casteoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_casteoContext {
	var p = new(Expre_casteoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expre_casteo

	return p
}

func (s *Expre_casteoContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_casteoContext) Get_R_AS() antlr.Token { return s._R_AS }

func (s *Expre_casteoContext) Set_R_AS(v antlr.Token) { s._R_AS = v }

func (s *Expre_casteoContext) Get_expre_arit() IExpre_aritContext { return s._expre_arit }

func (s *Expre_casteoContext) Get_type_casteo() IType_casteoContext { return s._type_casteo }

func (s *Expre_casteoContext) Set_expre_arit(v IExpre_aritContext) { s._expre_arit = v }

func (s *Expre_casteoContext) Set_type_casteo(v IType_casteoContext) { s._type_casteo = v }

func (s *Expre_casteoContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_casteoContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_casteoContext) Expre_arit() IExpre_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *Expre_casteoContext) R_AS() antlr.TerminalNode {
	return s.GetToken(ChemsR_AS, 0)
}

func (s *Expre_casteoContext) Type_casteo() IType_casteoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_casteoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_casteoContext)
}

func (s *Expre_casteoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_casteoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_casteoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpre_casteo(s)
	}
}

func (s *Expre_casteoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpre_casteo(s)
	}
}

func (p *Chems) Expre_casteo() (localctx IExpre_casteoContext) {
	localctx = NewExpre_casteoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ChemsRULE_expre_casteo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(745)

		var _x = p.expre_arit(0)

		localctx.(*Expre_casteoContext)._expre_arit = _x
	}
	{
		p.SetState(746)

		var _m = p.Match(ChemsR_AS)

		localctx.(*Expre_casteoContext)._R_AS = _m
	}
	{
		p.SetState(747)

		var _x = p.Type_casteo()

		localctx.(*Expre_casteoContext)._type_casteo = _x
	}
	localctx.(*Expre_casteoContext).p = casteo.NewCasteo(localctx.(*Expre_casteoContext).Get_expre_arit().GetP(), localctx.(*Expre_casteoContext).Get_type_casteo().GetTipo_exp(), (func() int {
		if localctx.(*Expre_casteoContext).Get_R_AS() == nil {
			return 0
		} else {
			return localctx.(*Expre_casteoContext).Get_R_AS().GetLine()
		}
	}()), localctx.(*Expre_casteoContext).Get_R_AS().GetColumn())

	return localctx
}

// IType_casteoContext is an interface to support dynamic dispatch.
type IType_casteoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipo_exp returns the tipo_exp attribute.
	GetTipo_exp() interfaces.TypeExpression

	// SetTipo_exp sets the tipo_exp attribute.
	SetTipo_exp(interfaces.TypeExpression)

	// IsType_casteoContext differentiates from other interfaces.
	IsType_casteoContext()
}

type Type_casteoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	tipo_exp interfaces.TypeExpression
}

func NewEmptyType_casteoContext() *Type_casteoContext {
	var p = new(Type_casteoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_type_casteo
	return p
}

func (*Type_casteoContext) IsType_casteoContext() {}

func NewType_casteoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_casteoContext {
	var p = new(Type_casteoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_type_casteo

	return p
}

func (s *Type_casteoContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_casteoContext) GetTipo_exp() interfaces.TypeExpression { return s.tipo_exp }

func (s *Type_casteoContext) SetTipo_exp(v interfaces.TypeExpression) { s.tipo_exp = v }

func (s *Type_casteoContext) R_INT() antlr.TerminalNode {
	return s.GetToken(ChemsR_INT, 0)
}

func (s *Type_casteoContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsR_FLOAT, 0)
}

func (s *Type_casteoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_casteoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_casteoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterType_casteo(s)
	}
}

func (s *Type_casteoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitType_casteo(s)
	}
}

func (p *Chems) Type_casteo() (localctx IType_casteoContext) {
	localctx = NewType_casteoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ChemsRULE_type_casteo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(754)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(750)
			p.Match(ChemsR_INT)
		}
		localctx.(*Type_casteoContext).tipo_exp = interfaces.INTEGER

	case ChemsR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(752)
			p.Match(ChemsR_FLOAT)
		}
		localctx.(*Type_casteoContext).tipo_exp = interfaces.FLOAT

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 40:
		var t *Expre_logContext = nil
		if localctx != nil {
			t = localctx.(*Expre_logContext)
		}
		return p.Expre_log_Sempred(t, predIndex)

	case 41:
		var t *Expre_relContext = nil
		if localctx != nil {
			t = localctx.(*Expre_relContext)
		}
		return p.Expre_rel_Sempred(t, predIndex)

	case 42:
		var t *Expre_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expre_aritContext)
		}
		return p.Expre_arit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Chems) Expre_log_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Expre_rel_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Expre_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
