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
import "OLC2/Compilador/instruction/variable"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 64, 304,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 43, 10, 3, 13, 3, 14, 3,
	44, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 52, 10, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 67, 10,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 85, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 145, 10, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 165, 10, 10, 3, 11, 6, 11, 168, 10,
	11, 13, 11, 14, 11, 169, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 5, 12, 181, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 192, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 202, 10, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 7, 14, 209, 10, 14, 12, 14, 14, 14, 212, 11, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 223, 10, 15, 12,
	15, 14, 15, 226, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 241, 10, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 253,
	10, 16, 12, 16, 14, 16, 256, 11, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 5, 18, 276, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 302,
	10, 19, 3, 19, 2, 5, 26, 28, 30, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 2, 6, 3, 2, 58, 59, 4, 2, 39, 40, 43, 46,
	3, 2, 47, 49, 3, 2, 50, 51, 2, 323, 2, 38, 3, 2, 2, 2, 4, 42, 3, 2, 2,
	2, 6, 51, 3, 2, 2, 2, 8, 66, 3, 2, 2, 2, 10, 84, 3, 2, 2, 2, 12, 86, 3,
	2, 2, 2, 14, 144, 3, 2, 2, 2, 16, 146, 3, 2, 2, 2, 18, 164, 3, 2, 2, 2,
	20, 167, 3, 2, 2, 2, 22, 180, 3, 2, 2, 2, 24, 191, 3, 2, 2, 2, 26, 201,
	3, 2, 2, 2, 28, 213, 3, 2, 2, 2, 30, 240, 3, 2, 2, 2, 32, 257, 3, 2, 2,
	2, 34, 275, 3, 2, 2, 2, 36, 301, 3, 2, 2, 2, 38, 39, 5, 4, 3, 2, 39, 40,
	8, 2, 1, 2, 40, 3, 3, 2, 2, 2, 41, 43, 5, 8, 5, 2, 42, 41, 3, 2, 2, 2,
	43, 44, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 3,
	2, 2, 2, 46, 47, 8, 3, 1, 2, 47, 5, 3, 2, 2, 2, 48, 49, 7, 35, 2, 2, 49,
	52, 8, 4, 1, 2, 50, 52, 8, 4, 1, 2, 51, 48, 3, 2, 2, 2, 51, 50, 3, 2, 2,
	2, 52, 7, 3, 2, 2, 2, 53, 54, 5, 10, 6, 2, 54, 55, 5, 6, 4, 2, 55, 56,
	8, 5, 1, 2, 56, 67, 3, 2, 2, 2, 57, 58, 5, 12, 7, 2, 58, 59, 8, 5, 1, 2,
	59, 67, 3, 2, 2, 2, 60, 61, 5, 14, 8, 2, 61, 62, 8, 5, 1, 2, 62, 67, 3,
	2, 2, 2, 63, 64, 5, 16, 9, 2, 64, 65, 8, 5, 1, 2, 65, 67, 3, 2, 2, 2, 66,
	53, 3, 2, 2, 2, 66, 57, 3, 2, 2, 2, 66, 60, 3, 2, 2, 2, 66, 63, 3, 2, 2,
	2, 67, 9, 3, 2, 2, 2, 68, 69, 7, 4, 2, 2, 69, 70, 7, 52, 2, 2, 70, 71,
	5, 34, 18, 2, 71, 72, 7, 53, 2, 2, 72, 73, 7, 35, 2, 2, 73, 74, 8, 6, 1,
	2, 74, 85, 3, 2, 2, 2, 75, 76, 7, 4, 2, 2, 76, 77, 7, 52, 2, 2, 77, 78,
	7, 30, 2, 2, 78, 79, 7, 36, 2, 2, 79, 80, 5, 20, 11, 2, 80, 81, 7, 53,
	2, 2, 81, 82, 7, 35, 2, 2, 82, 83, 8, 6, 1, 2, 83, 85, 3, 2, 2, 2, 84,
	68, 3, 2, 2, 2, 84, 75, 3, 2, 2, 2, 85, 11, 3, 2, 2, 2, 86, 87, 7, 19,
	2, 2, 87, 88, 7, 18, 2, 2, 88, 89, 7, 52, 2, 2, 89, 90, 7, 53, 2, 2, 90,
	91, 7, 54, 2, 2, 91, 92, 5, 4, 3, 2, 92, 93, 7, 55, 2, 2, 93, 94, 8, 7,
	1, 2, 94, 13, 3, 2, 2, 2, 95, 96, 7, 6, 2, 2, 96, 97, 7, 7, 2, 2, 97, 98,
	7, 32, 2, 2, 98, 99, 7, 38, 2, 2, 99, 100, 5, 24, 13, 2, 100, 101, 7, 35,
	2, 2, 101, 102, 8, 8, 1, 2, 102, 145, 3, 2, 2, 2, 103, 104, 7, 6, 2, 2,
	104, 105, 7, 7, 2, 2, 105, 106, 7, 32, 2, 2, 106, 107, 7, 37, 2, 2, 107,
	108, 5, 18, 10, 2, 108, 109, 7, 35, 2, 2, 109, 110, 8, 8, 1, 2, 110, 145,
	3, 2, 2, 2, 111, 112, 7, 6, 2, 2, 112, 113, 7, 7, 2, 2, 113, 114, 7, 32,
	2, 2, 114, 115, 7, 37, 2, 2, 115, 116, 5, 18, 10, 2, 116, 117, 7, 38, 2,
	2, 117, 118, 5, 24, 13, 2, 118, 119, 7, 35, 2, 2, 119, 120, 8, 8, 1, 2,
	120, 145, 3, 2, 2, 2, 121, 122, 7, 6, 2, 2, 122, 123, 7, 32, 2, 2, 123,
	124, 7, 38, 2, 2, 124, 125, 5, 24, 13, 2, 125, 126, 7, 35, 2, 2, 126, 127,
	8, 8, 1, 2, 127, 145, 3, 2, 2, 2, 128, 129, 7, 6, 2, 2, 129, 130, 7, 32,
	2, 2, 130, 131, 7, 37, 2, 2, 131, 132, 5, 18, 10, 2, 132, 133, 7, 35, 2,
	2, 133, 134, 8, 8, 1, 2, 134, 145, 3, 2, 2, 2, 135, 136, 7, 6, 2, 2, 136,
	137, 7, 32, 2, 2, 137, 138, 7, 37, 2, 2, 138, 139, 5, 18, 10, 2, 139, 140,
	7, 38, 2, 2, 140, 141, 5, 24, 13, 2, 141, 142, 7, 35, 2, 2, 142, 143, 8,
	8, 1, 2, 143, 145, 3, 2, 2, 2, 144, 95, 3, 2, 2, 2, 144, 103, 3, 2, 2,
	2, 144, 111, 3, 2, 2, 2, 144, 121, 3, 2, 2, 2, 144, 128, 3, 2, 2, 2, 144,
	135, 3, 2, 2, 2, 145, 15, 3, 2, 2, 2, 146, 147, 7, 32, 2, 2, 147, 148,
	7, 38, 2, 2, 148, 149, 5, 24, 13, 2, 149, 150, 7, 35, 2, 2, 150, 151, 8,
	9, 1, 2, 151, 17, 3, 2, 2, 2, 152, 153, 7, 21, 2, 2, 153, 165, 8, 10, 1,
	2, 154, 155, 7, 22, 2, 2, 155, 165, 8, 10, 1, 2, 156, 157, 7, 23, 2, 2,
	157, 165, 8, 10, 1, 2, 158, 159, 7, 26, 2, 2, 159, 165, 8, 10, 1, 2, 160,
	161, 7, 24, 2, 2, 161, 165, 8, 10, 1, 2, 162, 163, 7, 25, 2, 2, 163, 165,
	8, 10, 1, 2, 164, 152, 3, 2, 2, 2, 164, 154, 3, 2, 2, 2, 164, 156, 3, 2,
	2, 2, 164, 158, 3, 2, 2, 2, 164, 160, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2,
	165, 19, 3, 2, 2, 2, 166, 168, 5, 22, 12, 2, 167, 166, 3, 2, 2, 2, 168,
	169, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 171,
	3, 2, 2, 2, 171, 172, 8, 11, 1, 2, 172, 21, 3, 2, 2, 2, 173, 174, 5, 24,
	13, 2, 174, 175, 7, 36, 2, 2, 175, 176, 8, 12, 1, 2, 176, 181, 3, 2, 2,
	2, 177, 178, 5, 24, 13, 2, 178, 179, 8, 12, 1, 2, 179, 181, 3, 2, 2, 2,
	180, 173, 3, 2, 2, 2, 180, 177, 3, 2, 2, 2, 181, 23, 3, 2, 2, 2, 182, 183,
	5, 26, 14, 2, 183, 184, 8, 13, 1, 2, 184, 192, 3, 2, 2, 2, 185, 186, 5,
	28, 15, 2, 186, 187, 8, 13, 1, 2, 187, 192, 3, 2, 2, 2, 188, 189, 5, 30,
	16, 2, 189, 190, 8, 13, 1, 2, 190, 192, 3, 2, 2, 2, 191, 182, 3, 2, 2,
	2, 191, 185, 3, 2, 2, 2, 191, 188, 3, 2, 2, 2, 192, 25, 3, 2, 2, 2, 193,
	194, 8, 14, 1, 2, 194, 195, 7, 61, 2, 2, 195, 196, 5, 26, 14, 5, 196, 197,
	8, 14, 1, 2, 197, 202, 3, 2, 2, 2, 198, 199, 5, 28, 15, 2, 199, 200, 8,
	14, 1, 2, 200, 202, 3, 2, 2, 2, 201, 193, 3, 2, 2, 2, 201, 198, 3, 2, 2,
	2, 202, 210, 3, 2, 2, 2, 203, 204, 12, 4, 2, 2, 204, 205, 9, 2, 2, 2, 205,
	206, 5, 26, 14, 5, 206, 207, 8, 14, 1, 2, 207, 209, 3, 2, 2, 2, 208, 203,
	3, 2, 2, 2, 209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2,
	2, 2, 211, 27, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 214, 8, 15, 1, 2,
	214, 215, 5, 30, 16, 2, 215, 216, 8, 15, 1, 2, 216, 224, 3, 2, 2, 2, 217,
	218, 12, 4, 2, 2, 218, 219, 9, 3, 2, 2, 219, 220, 5, 28, 15, 5, 220, 221,
	8, 15, 1, 2, 221, 223, 3, 2, 2, 2, 222, 217, 3, 2, 2, 2, 223, 226, 3, 2,
	2, 2, 224, 222, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 29, 3, 2, 2, 2,
	226, 224, 3, 2, 2, 2, 227, 228, 8, 16, 1, 2, 228, 229, 7, 51, 2, 2, 229,
	230, 5, 30, 16, 7, 230, 231, 8, 16, 1, 2, 231, 241, 3, 2, 2, 2, 232, 233,
	5, 32, 17, 2, 233, 234, 8, 16, 1, 2, 234, 241, 3, 2, 2, 2, 235, 236, 7,
	52, 2, 2, 236, 237, 5, 24, 13, 2, 237, 238, 7, 53, 2, 2, 238, 239, 8, 16,
	1, 2, 239, 241, 3, 2, 2, 2, 240, 227, 3, 2, 2, 2, 240, 232, 3, 2, 2, 2,
	240, 235, 3, 2, 2, 2, 241, 254, 3, 2, 2, 2, 242, 243, 12, 6, 2, 2, 243,
	244, 9, 4, 2, 2, 244, 245, 5, 30, 16, 7, 245, 246, 8, 16, 1, 2, 246, 253,
	3, 2, 2, 2, 247, 248, 12, 5, 2, 2, 248, 249, 9, 5, 2, 2, 249, 250, 5, 30,
	16, 6, 250, 251, 8, 16, 1, 2, 251, 253, 3, 2, 2, 2, 252, 242, 3, 2, 2,
	2, 252, 247, 3, 2, 2, 2, 253, 256, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 254,
	255, 3, 2, 2, 2, 255, 31, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 257, 258, 5,
	34, 18, 2, 258, 259, 8, 17, 1, 2, 259, 33, 3, 2, 2, 2, 260, 261, 7, 27,
	2, 2, 261, 276, 8, 18, 1, 2, 262, 263, 7, 28, 2, 2, 263, 276, 8, 18, 1,
	2, 264, 265, 7, 30, 2, 2, 265, 276, 8, 18, 1, 2, 266, 267, 7, 31, 2, 2,
	267, 276, 8, 18, 1, 2, 268, 269, 7, 29, 2, 2, 269, 276, 8, 18, 1, 2, 270,
	271, 7, 32, 2, 2, 271, 276, 8, 18, 1, 2, 272, 273, 5, 36, 19, 2, 273, 274,
	8, 18, 1, 2, 274, 276, 3, 2, 2, 2, 275, 260, 3, 2, 2, 2, 275, 262, 3, 2,
	2, 2, 275, 264, 3, 2, 2, 2, 275, 266, 3, 2, 2, 2, 275, 268, 3, 2, 2, 2,
	275, 270, 3, 2, 2, 2, 275, 272, 3, 2, 2, 2, 276, 35, 3, 2, 2, 2, 277, 278,
	7, 27, 2, 2, 278, 279, 7, 20, 2, 2, 279, 280, 7, 21, 2, 2, 280, 302, 8,
	19, 1, 2, 281, 282, 7, 27, 2, 2, 282, 283, 7, 20, 2, 2, 283, 284, 7, 22,
	2, 2, 284, 302, 8, 19, 1, 2, 285, 286, 7, 28, 2, 2, 286, 287, 7, 20, 2,
	2, 287, 288, 7, 21, 2, 2, 288, 302, 8, 19, 1, 2, 289, 290, 7, 28, 2, 2,
	290, 291, 7, 20, 2, 2, 291, 292, 7, 22, 2, 2, 292, 302, 8, 19, 1, 2, 293,
	294, 7, 31, 2, 2, 294, 295, 7, 20, 2, 2, 295, 296, 7, 21, 2, 2, 296, 302,
	8, 19, 1, 2, 297, 298, 7, 31, 2, 2, 298, 299, 7, 20, 2, 2, 299, 300, 7,
	22, 2, 2, 300, 302, 8, 19, 1, 2, 301, 277, 3, 2, 2, 2, 301, 281, 3, 2,
	2, 2, 301, 285, 3, 2, 2, 2, 301, 289, 3, 2, 2, 2, 301, 293, 3, 2, 2, 2,
	301, 297, 3, 2, 2, 2, 302, 37, 3, 2, 2, 2, 19, 44, 51, 66, 84, 144, 164,
	169, 180, 191, 201, 210, 224, 240, 252, 254, 275, 301,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'print!'", "'println!'", "'number'", "'let'", "'mut'", "'if'", "'else'",
	"'match'", "'while'", "'break'", "'continue'", "'return'", "'loop'", "'for'",
	"'in'", "'main'", "'fn'", "'as'", "'i64'", "'f64'", "'String'", "'bool'",
	"'char'", "'&str'", "", "", "", "", "", "", "'..'", "'.'", "';'", "','",
	"':'", "'='", "'=='", "'>='", "'=>'", "'->'", "'<='", "'!='", "'>'", "'<'",
	"'*'", "'/'", "'%'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'",
	"'&&'", "'||'", "'|'", "'!'",
}
var symbolicNames = []string{
	"", "R_PRINT", "R_PRINTLN", "P_NUMBER", "R_LET", "R_MUT", "R_IF", "R_ELSE",
	"R_MATCH", "R_WHILE", "R_BREAK", "R_CONTINUE", "R_RETURN", "R_LOOP", "R_FOR",
	"R_IN", "R_MAIN", "R_FUNCTION", "R_AS", "R_INT", "R_FLOAT", "R_STRING",
	"R_BOOL", "R_CHAR", "R_STR", "NUMBER", "DOUBLE", "CHAR", "STRING", "BOOLEAN",
	"ID", "TK_DOBLEPUNTO", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA", "TK_DOSPUNTOS",
	"TK_IGUAL", "TK_IGUALIGUAL", "TK_MAYORIGUAL", "TK_IGUALMAYOR", "TK_MENOSMAYOR",
	"TK_MENORIGUAL", "TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV",
	"TK_MODULO", "TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC",
	"TK_CORA", "TK_CORC", "TK_AND", "TK_OR", "TK_BARRA", "TK_NOT", "WHITESPACE",
	"TK_MULTI", "TK_LINE",
}

var ruleNames = []string{
	"start", "instrucciones", "end_instr", "instruccion", "instr_println",
	"instr_main", "instr_declaracion", "instr_asignacion", "instr_tipo", "list_expression",
	"block_list_expression", "expressions", "expre_log", "expre_rel", "expre_arit",
	"expre_valor", "primitivo", "primitivo_casteo",
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
	ChemsP_NUMBER      = 3
	ChemsR_LET         = 4
	ChemsR_MUT         = 5
	ChemsR_IF          = 6
	ChemsR_ELSE        = 7
	ChemsR_MATCH       = 8
	ChemsR_WHILE       = 9
	ChemsR_BREAK       = 10
	ChemsR_CONTINUE    = 11
	ChemsR_RETURN      = 12
	ChemsR_LOOP        = 13
	ChemsR_FOR         = 14
	ChemsR_IN          = 15
	ChemsR_MAIN        = 16
	ChemsR_FUNCTION    = 17
	ChemsR_AS          = 18
	ChemsR_INT         = 19
	ChemsR_FLOAT       = 20
	ChemsR_STRING      = 21
	ChemsR_BOOL        = 22
	ChemsR_CHAR        = 23
	ChemsR_STR         = 24
	ChemsNUMBER        = 25
	ChemsDOUBLE        = 26
	ChemsCHAR          = 27
	ChemsSTRING        = 28
	ChemsBOOLEAN       = 29
	ChemsID            = 30
	ChemsTK_DOBLEPUNTO = 31
	ChemsTK_PUNTO      = 32
	ChemsTK_PUNTOCOMA  = 33
	ChemsTK_COMA       = 34
	ChemsTK_DOSPUNTOS  = 35
	ChemsTK_IGUAL      = 36
	ChemsTK_IGUALIGUAL = 37
	ChemsTK_MAYORIGUAL = 38
	ChemsTK_IGUALMAYOR = 39
	ChemsTK_MENOSMAYOR = 40
	ChemsTK_MENORIGUAL = 41
	ChemsTK_DIFIGUAL   = 42
	ChemsTK_MAYOR      = 43
	ChemsTK_MENOR      = 44
	ChemsTK_MULT       = 45
	ChemsTK_DIV        = 46
	ChemsTK_MODULO     = 47
	ChemsTK_MAS        = 48
	ChemsTK_MENOS      = 49
	ChemsTK_PARA       = 50
	ChemsTK_PARC       = 51
	ChemsTK_LLAVEA     = 52
	ChemsTK_LLAVEC     = 53
	ChemsTK_CORA       = 54
	ChemsTK_CORC       = 55
	ChemsTK_AND        = 56
	ChemsTK_OR         = 57
	ChemsTK_BARRA      = 58
	ChemsTK_NOT        = 59
	ChemsWHITESPACE    = 60
	ChemsTK_MULTI      = 61
	ChemsTK_LINE       = 62
)

// Chems rules.
const (
	ChemsRULE_start                 = 0
	ChemsRULE_instrucciones         = 1
	ChemsRULE_end_instr             = 2
	ChemsRULE_instruccion           = 3
	ChemsRULE_instr_println         = 4
	ChemsRULE_instr_main            = 5
	ChemsRULE_instr_declaracion     = 6
	ChemsRULE_instr_asignacion      = 7
	ChemsRULE_instr_tipo            = 8
	ChemsRULE_list_expression       = 9
	ChemsRULE_block_list_expression = 10
	ChemsRULE_expressions           = 11
	ChemsRULE_expre_log             = 12
	ChemsRULE_expre_rel             = 13
	ChemsRULE_expre_arit            = 14
	ChemsRULE_expre_valor           = 15
	ChemsRULE_primitivo             = 16
	ChemsRULE_primitivo_casteo      = 17
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
		p.SetState(36)

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
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsR_PRINTLN)|(1<<ChemsR_LET)|(1<<ChemsR_FUNCTION)|(1<<ChemsID))) != 0) {
		{
			p.SetState(39)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).list = append(localctx.(*InstruccionesContext).list, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(42)
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

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_PUNTOCOMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*End_instrContext).v = 1

	case ChemsEOF, ChemsR_PRINTLN, ChemsR_LET, ChemsR_FUNCTION, ChemsID, ChemsTK_LLAVEC:
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

	// Set_instr_println sets the _instr_println rule contexts.
	Set_instr_println(IInstr_printlnContext)

	// Set_instr_main sets the _instr_main rule contexts.
	Set_instr_main(IInstr_mainContext)

	// Set_instr_declaracion sets the _instr_declaracion rule contexts.
	Set_instr_declaracion(IInstr_declaracionContext)

	// Set_instr_asignacion sets the _instr_asignacion rule contexts.
	Set_instr_asignacion(IInstr_asignacionContext)

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

func (s *InstruccionContext) Set_instr_println(v IInstr_printlnContext) { s._instr_println = v }

func (s *InstruccionContext) Set_instr_main(v IInstr_mainContext) { s._instr_main = v }

func (s *InstruccionContext) Set_instr_declaracion(v IInstr_declaracionContext) {
	s._instr_declaracion = v
}

func (s *InstruccionContext) Set_instr_asignacion(v IInstr_asignacionContext) { s._instr_asignacion = v }

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

	p.SetState(64)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_PRINTLN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)

			var _x = p.Instr_println()

			localctx.(*InstruccionContext)._instr_println = _x
		}
		{
			p.SetState(52)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_println().GetInstr()

	case ChemsR_FUNCTION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)

			var _x = p.Instr_main()

			localctx.(*InstruccionContext)._instr_main = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_main().GetInstr()

	case ChemsR_LET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)

			var _x = p.Instr_declaracion()

			localctx.(*InstruccionContext)._instr_declaracion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_declaracion().GetInstr()

	case ChemsID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)

			var _x = p.Instr_asignacion()

			localctx.(*InstruccionContext)._instr_asignacion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_asignacion().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)

			var _m = p.Match(ChemsR_PRINTLN)

			localctx.(*Instr_printlnContext)._R_PRINTLN = _m
		}
		{
			p.SetState(67)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(68)

			var _x = p.Primitivo()

			localctx.(*Instr_printlnContext)._primitivo = _x
		}
		{
			p.SetState(69)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(70)
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
			p.SetState(73)

			var _m = p.Match(ChemsR_PRINTLN)

			localctx.(*Instr_printlnContext)._R_PRINTLN = _m
		}
		{
			p.SetState(74)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(75)

			var _m = p.Match(ChemsSTRING)

			localctx.(*Instr_printlnContext)._STRING = _m
		}
		{
			p.SetState(76)
			p.Match(ChemsTK_COMA)
		}
		{
			p.SetState(77)

			var _x = p.List_expression()

			localctx.(*Instr_printlnContext)._list_expression = _x
		}
		{
			p.SetState(78)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(79)
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
		p.SetState(84)
		p.Match(ChemsR_FUNCTION)
	}
	{
		p.SetState(85)

		var _m = p.Match(ChemsR_MAIN)

		localctx.(*Instr_mainContext)._R_MAIN = _m
	}
	{
		p.SetState(86)
		p.Match(ChemsTK_PARA)
	}
	{
		p.SetState(87)
		p.Match(ChemsTK_PARC)
	}
	{
		p.SetState(88)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(89)

		var _x = p.Instrucciones()

		localctx.(*Instr_mainContext)._instrucciones = _x
	}
	{
		p.SetState(90)
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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(94)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(95)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(96)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(97)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(98)
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
			p.SetState(101)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(102)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(103)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(104)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(105)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(106)
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
			p.SetState(109)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(110)
			p.Match(ChemsR_MUT)
		}
		{
			p.SetState(111)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(112)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(113)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(114)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(115)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(116)
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
			p.SetState(119)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(120)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(121)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(122)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(123)
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
			p.SetState(126)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(127)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(128)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(129)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(130)
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
			p.SetState(133)

			var _m = p.Match(ChemsR_LET)

			localctx.(*Instr_declaracionContext)._R_LET = _m
		}
		{
			p.SetState(134)

			var _m = p.Match(ChemsID)

			localctx.(*Instr_declaracionContext)._ID = _m
		}
		{
			p.SetState(135)
			p.Match(ChemsTK_DOSPUNTOS)
		}
		{
			p.SetState(136)

			var _x = p.Instr_tipo()

			localctx.(*Instr_declaracionContext)._instr_tipo = _x
		}
		{
			p.SetState(137)
			p.Match(ChemsTK_IGUAL)
		}
		{
			p.SetState(138)

			var _x = p.Expressions()

			localctx.(*Instr_declaracionContext)._expressions = _x
		}
		{
			p.SetState(139)
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
		p.SetState(144)

		var _m = p.Match(ChemsID)

		localctx.(*Instr_asignacionContext)._ID = _m
	}
	{
		p.SetState(145)
		p.Match(ChemsTK_IGUAL)
	}
	{
		p.SetState(146)

		var _x = p.Expressions()

		localctx.(*Instr_asignacionContext)._expressions = _x
	}
	{
		p.SetState(147)
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
	p.EnterRule(localctx, 16, ChemsRULE_instr_tipo)

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

	p.SetState(162)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Match(ChemsR_INT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.INTEGER

	case ChemsR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Match(ChemsR_FLOAT)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.FLOAT

	case ChemsR_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.Match(ChemsR_STRING)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STRING

	case ChemsR_STR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(156)
			p.Match(ChemsR_STR)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.STR

	case ChemsR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(158)
			p.Match(ChemsR_BOOL)
		}
		localctx.(*Instr_tipoContext).tipo_exp = interfaces.BOOLEAN

	case ChemsR_CHAR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(160)
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
	p.EnterRule(localctx, 18, ChemsRULE_list_expression)

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
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsNUMBER)|(1<<ChemsDOUBLE)|(1<<ChemsCHAR)|(1<<ChemsSTRING)|(1<<ChemsBOOLEAN)|(1<<ChemsID))) != 0) || (((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(ChemsTK_MENOS-49))|(1<<(ChemsTK_PARA-49))|(1<<(ChemsTK_NOT-49)))) != 0) {
		{
			p.SetState(164)

			var _x = p.Block_list_expression()

			localctx.(*List_expressionContext)._block_list_expression = _x
		}
		localctx.(*List_expressionContext).e = append(localctx.(*List_expressionContext).e, localctx.(*List_expressionContext)._block_list_expression)

		p.SetState(167)
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
	p.EnterRule(localctx, 20, ChemsRULE_block_list_expression)

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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)

			var _x = p.Expressions()

			localctx.(*Block_list_expressionContext)._expressions = _x
		}
		{
			p.SetState(172)
			p.Match(ChemsTK_COMA)
		}
		localctx.(*Block_list_expressionContext).p = instruction.NewListExpre(localctx.(*Block_list_expressionContext).Get_expressions().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)

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

	// Set_expre_log sets the _expre_log rule contexts.
	Set_expre_log(IExpre_logContext)

	// Set_expre_rel sets the _expre_rel rule contexts.
	Set_expre_rel(IExpre_relContext)

	// Set_expre_arit sets the _expre_arit rule contexts.
	Set_expre_arit(IExpre_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expression
	_expre_log  IExpre_logContext
	_expre_rel  IExpre_relContext
	_expre_arit IExpre_aritContext
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

func (s *ExpressionsContext) Set_expre_log(v IExpre_logContext) { s._expre_log = v }

func (s *ExpressionsContext) Set_expre_rel(v IExpre_relContext) { s._expre_rel = v }

func (s *ExpressionsContext) Set_expre_arit(v IExpre_aritContext) { s._expre_arit = v }

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
	p.EnterRule(localctx, 22, ChemsRULE_expressions)

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

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)

			var _x = p.expre_log(0)

			localctx.(*ExpressionsContext)._expre_log = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_log().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)

			var _x = p.expre_rel(0)

			localctx.(*ExpressionsContext)._expre_rel = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_rel().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(186)

			var _x = p.expre_arit(0)

			localctx.(*ExpressionsContext)._expre_arit = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_arit().GetP()

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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ChemsRULE_expre_log, _p)
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
	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_NOT:
		{
			p.SetState(192)

			var _m = p.Match(ChemsTK_NOT)

			localctx.(*Expre_logContext).op = _m
		}
		{
			p.SetState(193)

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

	case ChemsNUMBER, ChemsDOUBLE, ChemsCHAR, ChemsSTRING, ChemsBOOLEAN, ChemsID, ChemsTK_MENOS, ChemsTK_PARA:
		{
			p.SetState(196)

			var _x = p.expre_rel(0)

			localctx.(*Expre_logContext)._expre_rel = _x
		}
		localctx.(*Expre_logContext).p = localctx.(*Expre_logContext).Get_expre_rel().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpre_logContext(p, _parentctx, _parentState)
			localctx.(*Expre_logContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_log)
			p.SetState(201)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(202)

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
				p.SetState(203)

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
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, ChemsRULE_expre_rel, _p)
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
		p.SetState(212)

		var _x = p.expre_arit(0)

		localctx.(*Expre_relContext)._expre_arit = _x
	}
	localctx.(*Expre_relContext).p = localctx.(*Expre_relContext).Get_expre_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpre_relContext(p, _parentctx, _parentState)
			localctx.(*Expre_relContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_rel)
			p.SetState(215)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(216)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expre_relContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(ChemsTK_IGUALIGUAL-37))|(1<<(ChemsTK_MAYORIGUAL-37))|(1<<(ChemsTK_MENORIGUAL-37))|(1<<(ChemsTK_DIFIGUAL-37))|(1<<(ChemsTK_MAYOR-37))|(1<<(ChemsTK_MENOR-37)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expre_relContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(217)

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
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	_startState := 28
	p.EnterRecursionRule(localctx, 28, ChemsRULE_expre_arit, _p)
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
	p.SetState(238)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_MENOS:
		{
			p.SetState(226)

			var _m = p.Match(ChemsTK_MENOS)

			localctx.(*Expre_aritContext).op = _m
		}
		{
			p.SetState(227)

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

	case ChemsNUMBER, ChemsDOUBLE, ChemsCHAR, ChemsSTRING, ChemsBOOLEAN, ChemsID:
		{
			p.SetState(230)

			var _x = p.Expre_valor()

			localctx.(*Expre_aritContext)._expre_valor = _x
		}
		localctx.(*Expre_aritContext).p = localctx.(*Expre_aritContext).Get_expre_valor().GetP()

	case ChemsTK_PARA:
		{
			p.SetState(233)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(234)

			var _x = p.Expressions()

			localctx.(*Expre_aritContext)._expressions = _x
		}
		{
			p.SetState(235)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Expre_aritContext).p = localctx.(*Expre_aritContext).Get_expressions().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpre_aritContext(p, _parentctx, _parentState)
				localctx.(*Expre_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expre_arit)
				p.SetState(240)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(241)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expre_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(ChemsTK_MULT-45))|(1<<(ChemsTK_DIV-45))|(1<<(ChemsTK_MODULO-45)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expre_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(242)

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
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(246)

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
					p.SetState(247)

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
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 30, ChemsRULE_expre_valor)

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
		p.SetState(255)

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

	// Set_primitivo_casteo sets the _primitivo_casteo rule contexts.
	Set_primitivo_casteo(IPrimitivo_casteoContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	p                 interfaces.Expression
	_NUMBER           antlr.Token
	_DOUBLE           antlr.Token
	_STRING           antlr.Token
	_BOOLEAN          antlr.Token
	_CHAR             antlr.Token
	_ID               antlr.Token
	_primitivo_casteo IPrimitivo_casteoContext
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

func (s *PrimitivoContext) Set_primitivo_casteo(v IPrimitivo_casteoContext) { s._primitivo_casteo = v }

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
	p.EnterRule(localctx, 32, ChemsRULE_primitivo)

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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)

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
			p.SetState(260)

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
			p.SetState(262)

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
			p.SetState(264)

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
			p.SetState(266)

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
			p.SetState(268)

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
			p.SetState(270)

			var _x = p.Primitivo_casteo()

			localctx.(*PrimitivoContext)._primitivo_casteo = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_primitivo_casteo().GetP()

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
	p.EnterRule(localctx, 34, ChemsRULE_primitivo_casteo)

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

	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*Primitivo_casteoContext)._NUMBER = _m
		}
		{
			p.SetState(276)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(277)
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
			p.SetState(279)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*Primitivo_casteoContext)._NUMBER = _m
		}
		{
			p.SetState(280)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(281)
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
			p.SetState(283)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*Primitivo_casteoContext)._DOUBLE = _m
		}
		{
			p.SetState(284)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(285)
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
			p.SetState(287)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*Primitivo_casteoContext)._DOUBLE = _m
		}
		{
			p.SetState(288)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(289)
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
			p.SetState(291)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*Primitivo_casteoContext)._BOOLEAN = _m
		}
		{
			p.SetState(292)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(293)
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
			p.SetState(295)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*Primitivo_casteoContext)._BOOLEAN = _m
		}
		{
			p.SetState(296)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(297)
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

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *Expre_logContext = nil
		if localctx != nil {
			t = localctx.(*Expre_logContext)
		}
		return p.Expre_log_Sempred(t, predIndex)

	case 13:
		var t *Expre_relContext = nil
		if localctx != nil {
			t = localctx.(*Expre_relContext)
		}
		return p.Expre_rel_Sempred(t, predIndex)

	case 14:
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
