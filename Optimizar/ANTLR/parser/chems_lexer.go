// Code generated from ChemsLexer.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 49, 310,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 6, 19, 198, 10, 19,
	13, 19, 14, 19, 199, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 6, 20, 211, 10, 20, 13, 20, 14, 20, 212, 3, 20, 3, 20, 3, 21,
	6, 21, 218, 10, 21, 13, 21, 14, 21, 219, 3, 22, 6, 22, 223, 10, 22, 13,
	22, 14, 22, 224, 3, 22, 3, 22, 6, 22, 229, 10, 22, 13, 22, 14, 22, 230,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 7, 24, 239, 10, 24, 12, 24, 14,
	24, 242, 11, 24, 3, 25, 3, 25, 7, 25, 246, 10, 25, 12, 25, 14, 25, 249,
	11, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3,
	34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38,
	3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3,
	43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48,
	6, 48, 302, 10, 48, 13, 48, 14, 48, 303, 3, 48, 3, 48, 3, 49, 3, 49, 3,
	49, 2, 2, 50, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73,
	38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91,
	47, 93, 48, 95, 49, 97, 2, 3, 2, 10, 3, 2, 49, 49, 3, 2, 12, 12, 3, 2,
	50, 59, 3, 2, 36, 36, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67,
	92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 9, 2, 34, 35, 37, 37,
	45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2, 316, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2,
	67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2,
	2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2,
	2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 3, 99, 3,
	2, 2, 2, 5, 103, 3, 2, 2, 2, 7, 110, 3, 2, 2, 2, 9, 115, 3, 2, 2, 2, 11,
	120, 3, 2, 2, 2, 13, 129, 3, 2, 2, 2, 15, 139, 3, 2, 2, 2, 17, 148, 3,
	2, 2, 2, 19, 153, 3, 2, 2, 2, 21, 159, 3, 2, 2, 2, 23, 162, 3, 2, 2, 2,
	25, 167, 3, 2, 2, 2, 27, 174, 3, 2, 2, 2, 29, 181, 3, 2, 2, 2, 31, 185,
	3, 2, 2, 2, 33, 189, 3, 2, 2, 2, 35, 191, 3, 2, 2, 2, 37, 193, 3, 2, 2,
	2, 39, 206, 3, 2, 2, 2, 41, 217, 3, 2, 2, 2, 43, 222, 3, 2, 2, 2, 45, 232,
	3, 2, 2, 2, 47, 236, 3, 2, 2, 2, 49, 243, 3, 2, 2, 2, 51, 252, 3, 2, 2,
	2, 53, 254, 3, 2, 2, 2, 55, 256, 3, 2, 2, 2, 57, 258, 3, 2, 2, 2, 59, 260,
	3, 2, 2, 2, 61, 262, 3, 2, 2, 2, 63, 264, 3, 2, 2, 2, 65, 266, 3, 2, 2,
	2, 67, 268, 3, 2, 2, 2, 69, 270, 3, 2, 2, 2, 71, 272, 3, 2, 2, 2, 73, 274,
	3, 2, 2, 2, 75, 277, 3, 2, 2, 2, 77, 280, 3, 2, 2, 2, 79, 282, 3, 2, 2,
	2, 81, 284, 3, 2, 2, 2, 83, 287, 3, 2, 2, 2, 85, 290, 3, 2, 2, 2, 87, 292,
	3, 2, 2, 2, 89, 294, 3, 2, 2, 2, 91, 296, 3, 2, 2, 2, 93, 298, 3, 2, 2,
	2, 95, 301, 3, 2, 2, 2, 97, 307, 3, 2, 2, 2, 99, 100, 7, 107, 2, 2, 100,
	101, 7, 112, 2, 2, 101, 102, 7, 118, 2, 2, 102, 4, 3, 2, 2, 2, 103, 104,
	7, 102, 2, 2, 104, 105, 7, 113, 2, 2, 105, 106, 7, 119, 2, 2, 106, 107,
	7, 100, 2, 2, 107, 108, 7, 110, 2, 2, 108, 109, 7, 103, 2, 2, 109, 6, 3,
	2, 2, 2, 110, 111, 7, 101, 2, 2, 111, 112, 7, 106, 2, 2, 112, 113, 7, 99,
	2, 2, 113, 114, 7, 116, 2, 2, 114, 8, 3, 2, 2, 2, 115, 116, 7, 120, 2,
	2, 116, 117, 7, 113, 2, 2, 117, 118, 7, 107, 2, 2, 118, 119, 7, 102, 2,
	2, 119, 10, 3, 2, 2, 2, 120, 121, 7, 37, 2, 2, 121, 122, 7, 107, 2, 2,
	122, 123, 7, 112, 2, 2, 123, 124, 7, 101, 2, 2, 124, 125, 7, 110, 2, 2,
	125, 126, 7, 119, 2, 2, 126, 127, 7, 102, 2, 2, 127, 128, 7, 103, 2, 2,
	128, 12, 3, 2, 2, 2, 129, 130, 7, 62, 2, 2, 130, 131, 7, 117, 2, 2, 131,
	132, 7, 118, 2, 2, 132, 133, 7, 102, 2, 2, 133, 134, 7, 107, 2, 2, 134,
	135, 7, 113, 2, 2, 135, 136, 7, 48, 2, 2, 136, 137, 7, 106, 2, 2, 137,
	138, 7, 64, 2, 2, 138, 14, 3, 2, 2, 2, 139, 140, 7, 62, 2, 2, 140, 141,
	7, 111, 2, 2, 141, 142, 7, 99, 2, 2, 142, 143, 7, 118, 2, 2, 143, 144,
	7, 106, 2, 2, 144, 145, 7, 48, 2, 2, 145, 146, 7, 106, 2, 2, 146, 147,
	7, 64, 2, 2, 147, 16, 3, 2, 2, 2, 148, 149, 7, 106, 2, 2, 149, 150, 7,
	103, 2, 2, 150, 151, 7, 99, 2, 2, 151, 152, 7, 114, 2, 2, 152, 18, 3, 2,
	2, 2, 153, 154, 7, 117, 2, 2, 154, 155, 7, 118, 2, 2, 155, 156, 7, 99,
	2, 2, 156, 157, 7, 101, 2, 2, 157, 158, 7, 109, 2, 2, 158, 20, 3, 2, 2,
	2, 159, 160, 7, 107, 2, 2, 160, 161, 7, 104, 2, 2, 161, 22, 3, 2, 2, 2,
	162, 163, 7, 105, 2, 2, 163, 164, 7, 113, 2, 2, 164, 165, 7, 118, 2, 2,
	165, 166, 7, 113, 2, 2, 166, 24, 3, 2, 2, 2, 167, 168, 7, 116, 2, 2, 168,
	169, 7, 103, 2, 2, 169, 170, 7, 118, 2, 2, 170, 171, 7, 119, 2, 2, 171,
	172, 7, 116, 2, 2, 172, 173, 7, 112, 2, 2, 173, 26, 3, 2, 2, 2, 174, 175,
	7, 114, 2, 2, 175, 176, 7, 116, 2, 2, 176, 177, 7, 107, 2, 2, 177, 178,
	7, 112, 2, 2, 178, 179, 7, 118, 2, 2, 179, 180, 7, 104, 2, 2, 180, 28,
	3, 2, 2, 2, 181, 182, 7, 114, 2, 2, 182, 183, 7, 113, 2, 2, 183, 184, 7,
	121, 2, 2, 184, 30, 3, 2, 2, 2, 185, 186, 7, 111, 2, 2, 186, 187, 7, 113,
	2, 2, 187, 188, 7, 102, 2, 2, 188, 32, 3, 2, 2, 2, 189, 190, 7, 82, 2,
	2, 190, 34, 3, 2, 2, 2, 191, 192, 7, 74, 2, 2, 192, 36, 3, 2, 2, 2, 193,
	194, 7, 49, 2, 2, 194, 195, 7, 44, 2, 2, 195, 197, 3, 2, 2, 2, 196, 198,
	10, 2, 2, 2, 197, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 197, 3, 2,
	2, 2, 199, 200, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 202, 7, 44, 2, 2,
	202, 203, 7, 49, 2, 2, 203, 204, 3, 2, 2, 2, 204, 205, 8, 19, 2, 2, 205,
	38, 3, 2, 2, 2, 206, 207, 7, 49, 2, 2, 207, 208, 7, 49, 2, 2, 208, 210,
	3, 2, 2, 2, 209, 211, 10, 3, 2, 2, 210, 209, 3, 2, 2, 2, 211, 212, 3, 2,
	2, 2, 212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2,
	214, 215, 8, 20, 2, 2, 215, 40, 3, 2, 2, 2, 216, 218, 9, 4, 2, 2, 217,
	216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 219, 220,
	3, 2, 2, 2, 220, 42, 3, 2, 2, 2, 221, 223, 9, 4, 2, 2, 222, 221, 3, 2,
	2, 2, 223, 224, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2,
	225, 226, 3, 2, 2, 2, 226, 228, 7, 48, 2, 2, 227, 229, 9, 4, 2, 2, 228,
	227, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 230, 231,
	3, 2, 2, 2, 231, 44, 3, 2, 2, 2, 232, 233, 7, 41, 2, 2, 233, 234, 10, 5,
	2, 2, 234, 235, 7, 41, 2, 2, 235, 46, 3, 2, 2, 2, 236, 240, 9, 6, 2, 2,
	237, 239, 9, 7, 2, 2, 238, 237, 3, 2, 2, 2, 239, 242, 3, 2, 2, 2, 240,
	238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 48, 3, 2, 2, 2, 242, 240, 3,
	2, 2, 2, 243, 247, 7, 36, 2, 2, 244, 246, 10, 5, 2, 2, 245, 244, 3, 2,
	2, 2, 246, 249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2,
	248, 250, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 251, 7, 36, 2, 2, 251,
	50, 3, 2, 2, 2, 252, 253, 7, 63, 2, 2, 253, 52, 3, 2, 2, 2, 254, 255, 7,
	61, 2, 2, 255, 54, 3, 2, 2, 2, 256, 257, 7, 46, 2, 2, 257, 56, 3, 2, 2,
	2, 258, 259, 7, 60, 2, 2, 259, 58, 3, 2, 2, 2, 260, 261, 7, 35, 2, 2, 261,
	60, 3, 2, 2, 2, 262, 263, 7, 42, 2, 2, 263, 62, 3, 2, 2, 2, 264, 265, 7,
	43, 2, 2, 265, 64, 3, 2, 2, 2, 266, 267, 7, 125, 2, 2, 267, 66, 3, 2, 2,
	2, 268, 269, 7, 127, 2, 2, 269, 68, 3, 2, 2, 2, 270, 271, 7, 93, 2, 2,
	271, 70, 3, 2, 2, 2, 272, 273, 7, 95, 2, 2, 273, 72, 3, 2, 2, 2, 274, 275,
	7, 63, 2, 2, 275, 276, 7, 63, 2, 2, 276, 74, 3, 2, 2, 2, 277, 278, 7, 35,
	2, 2, 278, 279, 7, 63, 2, 2, 279, 76, 3, 2, 2, 2, 280, 281, 7, 64, 2, 2,
	281, 78, 3, 2, 2, 2, 282, 283, 7, 62, 2, 2, 283, 80, 3, 2, 2, 2, 284, 285,
	7, 64, 2, 2, 285, 286, 7, 63, 2, 2, 286, 82, 3, 2, 2, 2, 287, 288, 7, 62,
	2, 2, 288, 289, 7, 63, 2, 2, 289, 84, 3, 2, 2, 2, 290, 291, 7, 44, 2, 2,
	291, 86, 3, 2, 2, 2, 292, 293, 7, 49, 2, 2, 293, 88, 3, 2, 2, 2, 294, 295,
	7, 39, 2, 2, 295, 90, 3, 2, 2, 2, 296, 297, 7, 45, 2, 2, 297, 92, 3, 2,
	2, 2, 298, 299, 7, 47, 2, 2, 299, 94, 3, 2, 2, 2, 300, 302, 9, 8, 2, 2,
	301, 300, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 303,
	304, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 306, 8, 48, 2, 2, 306, 96,
	3, 2, 2, 2, 307, 308, 7, 94, 2, 2, 308, 309, 9, 9, 2, 2, 309, 98, 3, 2,
	2, 2, 11, 2, 199, 212, 219, 224, 230, 240, 247, 303, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'int'", "'double'", "'char'", "'void'", "'#include'", "'<stdio.h>'",
	"'<math.h>'", "'heap'", "'stack'", "'if'", "'goto'", "'return'", "'printf'",
	"'pow'", "'mod'", "'P'", "'H'", "", "", "", "", "", "", "", "'='", "';'",
	"','", "':'", "'!'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'=='",
	"'!='", "'>'", "'<'", "'>='", "'<='", "'*'", "'/'", "'%'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "RINT", "RDOUBLE", "RCHAR", "RVOID", "RINCLUDE", "RSTDIO", "RMATH",
	"RHEAP", "RSTACK", "RIF", "RGOTO", "RRETURN", "RPRINTF", "RPOW", "RMOD",
	"RPSTACK", "RPHEAP", "MULTILINE", "INLINE", "INTEGER", "DOUBLE", "CHAR",
	"ID", "STRING", "EQUAL", "SEMICOLON", "COMMA", "COLON", "ADMIRATION", "LEFT_PAR",
	"RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET",
	"EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN",
	"LESS_EQUEALTHAN", "MUL", "DIV", "MOD", "ADD", "SUB", "WHITESPACE",
}

var lexerRuleNames = []string{
	"RINT", "RDOUBLE", "RCHAR", "RVOID", "RINCLUDE", "RSTDIO", "RMATH", "RHEAP",
	"RSTACK", "RIF", "RGOTO", "RRETURN", "RPRINTF", "RPOW", "RMOD", "RPSTACK",
	"RPHEAP", "MULTILINE", "INLINE", "INTEGER", "DOUBLE", "CHAR", "ID", "STRING",
	"EQUAL", "SEMICOLON", "COMMA", "COLON", "ADMIRATION", "LEFT_PAR", "RIGHT_PAR",
	"LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL",
	"NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN",
	"MUL", "DIV", "MOD", "ADD", "SUB", "WHITESPACE", "ESC_SEQ",
}

type ChemsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewChemsLexer(input antlr.CharStream) *ChemsLexer {

	l := new(ChemsLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ChemsLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ChemsLexer tokens.
const (
	ChemsLexerRINT              = 1
	ChemsLexerRDOUBLE           = 2
	ChemsLexerRCHAR             = 3
	ChemsLexerRVOID             = 4
	ChemsLexerRINCLUDE          = 5
	ChemsLexerRSTDIO            = 6
	ChemsLexerRMATH             = 7
	ChemsLexerRHEAP             = 8
	ChemsLexerRSTACK            = 9
	ChemsLexerRIF               = 10
	ChemsLexerRGOTO             = 11
	ChemsLexerRRETURN           = 12
	ChemsLexerRPRINTF           = 13
	ChemsLexerRPOW              = 14
	ChemsLexerRMOD              = 15
	ChemsLexerRPSTACK           = 16
	ChemsLexerRPHEAP            = 17
	ChemsLexerMULTILINE         = 18
	ChemsLexerINLINE            = 19
	ChemsLexerINTEGER           = 20
	ChemsLexerDOUBLE            = 21
	ChemsLexerCHAR              = 22
	ChemsLexerID                = 23
	ChemsLexerSTRING            = 24
	ChemsLexerEQUAL             = 25
	ChemsLexerSEMICOLON         = 26
	ChemsLexerCOMMA             = 27
	ChemsLexerCOLON             = 28
	ChemsLexerADMIRATION        = 29
	ChemsLexerLEFT_PAR          = 30
	ChemsLexerRIGHT_PAR         = 31
	ChemsLexerLEFT_KEY          = 32
	ChemsLexerRIGHT_KEY         = 33
	ChemsLexerLEFT_BRACKET      = 34
	ChemsLexerRIGHT_BRACKET     = 35
	ChemsLexerEQUEAL_EQUAL      = 36
	ChemsLexerNOTEQUAL          = 37
	ChemsLexerGREATER_THAN      = 38
	ChemsLexerLESS_THAN         = 39
	ChemsLexerGREATER_EQUALTHAN = 40
	ChemsLexerLESS_EQUEALTHAN   = 41
	ChemsLexerMUL               = 42
	ChemsLexerDIV               = 43
	ChemsLexerMOD               = 44
	ChemsLexerADD               = 45
	ChemsLexerSUB               = 46
	ChemsLexerWHITESPACE        = 47
)