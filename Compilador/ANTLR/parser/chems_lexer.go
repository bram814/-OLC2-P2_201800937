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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 66, 412,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 28, 6, 28, 267, 10, 28, 13, 28, 14, 28, 268,
	3, 29, 6, 29, 272, 10, 29, 13, 29, 14, 29, 273, 3, 29, 3, 29, 6, 29, 278,
	10, 29, 13, 29, 14, 29, 279, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31,
	7, 31, 288, 10, 31, 12, 31, 14, 31, 291, 11, 31, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 304, 10, 32,
	3, 33, 3, 33, 7, 33, 308, 10, 33, 12, 33, 14, 33, 311, 11, 33, 3, 34, 3,
	34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39,
	3, 39, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3,
	43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46,
	3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3,
	52, 3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57,
	3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 61, 3,
	61, 3, 62, 3, 62, 3, 63, 6, 63, 381, 10, 63, 13, 63, 14, 63, 382, 3, 63,
	3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 6, 64, 391, 10, 64, 13, 64, 14, 64,
	392, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65, 6,
	65, 404, 10, 65, 13, 65, 14, 65, 405, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66,
	2, 2, 67, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29,
	57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38,
	75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47,
	93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109,
	56, 111, 57, 113, 58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125,
	64, 127, 65, 129, 66, 131, 2, 3, 2, 10, 3, 2, 50, 59, 3, 2, 36, 36, 5,
	2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2,
	11, 12, 15, 15, 34, 34, 3, 2, 49, 49, 3, 2, 12, 12, 9, 2, 34, 35, 37, 37,
	45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2, 419, 2, 3, 3, 2, 2, 2, 2, 5,
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
	2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3,
	2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2,
	105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2,
	2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119,
	3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2,
	2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 3, 133, 3, 2, 2, 2, 5, 140, 3,
	2, 2, 2, 7, 149, 3, 2, 2, 2, 9, 153, 3, 2, 2, 2, 11, 157, 3, 2, 2, 2, 13,
	160, 3, 2, 2, 2, 15, 165, 3, 2, 2, 2, 17, 171, 3, 2, 2, 2, 19, 177, 3,
	2, 2, 2, 21, 182, 3, 2, 2, 2, 23, 186, 3, 2, 2, 2, 25, 189, 3, 2, 2, 2,
	27, 195, 3, 2, 2, 2, 29, 204, 3, 2, 2, 2, 31, 211, 3, 2, 2, 2, 33, 216,
	3, 2, 2, 2, 35, 219, 3, 2, 2, 2, 37, 223, 3, 2, 2, 2, 39, 228, 3, 2, 2,
	2, 41, 232, 3, 2, 2, 2, 43, 235, 3, 2, 2, 2, 45, 239, 3, 2, 2, 2, 47, 243,
	3, 2, 2, 2, 49, 250, 3, 2, 2, 2, 51, 255, 3, 2, 2, 2, 53, 260, 3, 2, 2,
	2, 55, 266, 3, 2, 2, 2, 57, 271, 3, 2, 2, 2, 59, 281, 3, 2, 2, 2, 61, 285,
	3, 2, 2, 2, 63, 303, 3, 2, 2, 2, 65, 305, 3, 2, 2, 2, 67, 312, 3, 2, 2,
	2, 69, 315, 3, 2, 2, 2, 71, 317, 3, 2, 2, 2, 73, 319, 3, 2, 2, 2, 75, 321,
	3, 2, 2, 2, 77, 323, 3, 2, 2, 2, 79, 325, 3, 2, 2, 2, 81, 328, 3, 2, 2,
	2, 83, 331, 3, 2, 2, 2, 85, 334, 3, 2, 2, 2, 87, 337, 3, 2, 2, 2, 89, 340,
	3, 2, 2, 2, 91, 343, 3, 2, 2, 2, 93, 345, 3, 2, 2, 2, 95, 347, 3, 2, 2,
	2, 97, 349, 3, 2, 2, 2, 99, 351, 3, 2, 2, 2, 101, 353, 3, 2, 2, 2, 103,
	355, 3, 2, 2, 2, 105, 357, 3, 2, 2, 2, 107, 359, 3, 2, 2, 2, 109, 361,
	3, 2, 2, 2, 111, 363, 3, 2, 2, 2, 113, 365, 3, 2, 2, 2, 115, 367, 3, 2,
	2, 2, 117, 369, 3, 2, 2, 2, 119, 372, 3, 2, 2, 2, 121, 375, 3, 2, 2, 2,
	123, 377, 3, 2, 2, 2, 125, 380, 3, 2, 2, 2, 127, 386, 3, 2, 2, 2, 129,
	399, 3, 2, 2, 2, 131, 409, 3, 2, 2, 2, 133, 134, 7, 114, 2, 2, 134, 135,
	7, 116, 2, 2, 135, 136, 7, 107, 2, 2, 136, 137, 7, 112, 2, 2, 137, 138,
	7, 118, 2, 2, 138, 139, 7, 35, 2, 2, 139, 4, 3, 2, 2, 2, 140, 141, 7, 114,
	2, 2, 141, 142, 7, 116, 2, 2, 142, 143, 7, 107, 2, 2, 143, 144, 7, 112,
	2, 2, 144, 145, 7, 118, 2, 2, 145, 146, 7, 110, 2, 2, 146, 147, 7, 112,
	2, 2, 147, 148, 7, 35, 2, 2, 148, 6, 3, 2, 2, 2, 149, 150, 7, 110, 2, 2,
	150, 151, 7, 103, 2, 2, 151, 152, 7, 118, 2, 2, 152, 8, 3, 2, 2, 2, 153,
	154, 7, 111, 2, 2, 154, 155, 7, 119, 2, 2, 155, 156, 7, 118, 2, 2, 156,
	10, 3, 2, 2, 2, 157, 158, 7, 107, 2, 2, 158, 159, 7, 104, 2, 2, 159, 12,
	3, 2, 2, 2, 160, 161, 7, 103, 2, 2, 161, 162, 7, 110, 2, 2, 162, 163, 7,
	117, 2, 2, 163, 164, 7, 103, 2, 2, 164, 14, 3, 2, 2, 2, 165, 166, 7, 111,
	2, 2, 166, 167, 7, 99, 2, 2, 167, 168, 7, 118, 2, 2, 168, 169, 7, 101,
	2, 2, 169, 170, 7, 106, 2, 2, 170, 16, 3, 2, 2, 2, 171, 172, 7, 121, 2,
	2, 172, 173, 7, 106, 2, 2, 173, 174, 7, 107, 2, 2, 174, 175, 7, 110, 2,
	2, 175, 176, 7, 103, 2, 2, 176, 18, 3, 2, 2, 2, 177, 178, 7, 110, 2, 2,
	178, 179, 7, 113, 2, 2, 179, 180, 7, 113, 2, 2, 180, 181, 7, 114, 2, 2,
	181, 20, 3, 2, 2, 2, 182, 183, 7, 104, 2, 2, 183, 184, 7, 113, 2, 2, 184,
	185, 7, 116, 2, 2, 185, 22, 3, 2, 2, 2, 186, 187, 7, 107, 2, 2, 187, 188,
	7, 112, 2, 2, 188, 24, 3, 2, 2, 2, 189, 190, 7, 100, 2, 2, 190, 191, 7,
	116, 2, 2, 191, 192, 7, 103, 2, 2, 192, 193, 7, 99, 2, 2, 193, 194, 7,
	109, 2, 2, 194, 26, 3, 2, 2, 2, 195, 196, 7, 101, 2, 2, 196, 197, 7, 113,
	2, 2, 197, 198, 7, 112, 2, 2, 198, 199, 7, 118, 2, 2, 199, 200, 7, 107,
	2, 2, 200, 201, 7, 112, 2, 2, 201, 202, 7, 119, 2, 2, 202, 203, 7, 103,
	2, 2, 203, 28, 3, 2, 2, 2, 204, 205, 7, 116, 2, 2, 205, 206, 7, 103, 2,
	2, 206, 207, 7, 118, 2, 2, 207, 208, 7, 119, 2, 2, 208, 209, 7, 116, 2,
	2, 209, 210, 7, 112, 2, 2, 210, 30, 3, 2, 2, 2, 211, 212, 7, 111, 2, 2,
	212, 213, 7, 99, 2, 2, 213, 214, 7, 107, 2, 2, 214, 215, 7, 112, 2, 2,
	215, 32, 3, 2, 2, 2, 216, 217, 7, 104, 2, 2, 217, 218, 7, 112, 2, 2, 218,
	34, 3, 2, 2, 2, 219, 220, 7, 114, 2, 2, 220, 221, 7, 113, 2, 2, 221, 222,
	7, 121, 2, 2, 222, 36, 3, 2, 2, 2, 223, 224, 7, 114, 2, 2, 224, 225, 7,
	113, 2, 2, 225, 226, 7, 121, 2, 2, 226, 227, 7, 104, 2, 2, 227, 38, 3,
	2, 2, 2, 228, 229, 7, 99, 2, 2, 229, 230, 7, 100, 2, 2, 230, 231, 7, 117,
	2, 2, 231, 40, 3, 2, 2, 2, 232, 233, 7, 99, 2, 2, 233, 234, 7, 117, 2,
	2, 234, 42, 3, 2, 2, 2, 235, 236, 7, 107, 2, 2, 236, 237, 7, 56, 2, 2,
	237, 238, 7, 54, 2, 2, 238, 44, 3, 2, 2, 2, 239, 240, 7, 104, 2, 2, 240,
	241, 7, 56, 2, 2, 241, 242, 7, 54, 2, 2, 242, 46, 3, 2, 2, 2, 243, 244,
	7, 85, 2, 2, 244, 245, 7, 118, 2, 2, 245, 246, 7, 116, 2, 2, 246, 247,
	7, 107, 2, 2, 247, 248, 7, 112, 2, 2, 248, 249, 7, 105, 2, 2, 249, 48,
	3, 2, 2, 2, 250, 251, 7, 100, 2, 2, 251, 252, 7, 113, 2, 2, 252, 253, 7,
	113, 2, 2, 253, 254, 7, 110, 2, 2, 254, 50, 3, 2, 2, 2, 255, 256, 7, 101,
	2, 2, 256, 257, 7, 106, 2, 2, 257, 258, 7, 99, 2, 2, 258, 259, 7, 116,
	2, 2, 259, 52, 3, 2, 2, 2, 260, 261, 7, 40, 2, 2, 261, 262, 7, 117, 2,
	2, 262, 263, 7, 118, 2, 2, 263, 264, 7, 116, 2, 2, 264, 54, 3, 2, 2, 2,
	265, 267, 9, 2, 2, 2, 266, 265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268,
	266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 56, 3, 2, 2, 2, 270, 272, 9,
	2, 2, 2, 271, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 271, 3, 2, 2,
	2, 273, 274, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 277, 7, 48, 2, 2, 276,
	278, 9, 2, 2, 2, 277, 276, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 277,
	3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 58, 3, 2, 2, 2, 281, 282, 7, 41,
	2, 2, 282, 283, 10, 3, 2, 2, 283, 284, 7, 41, 2, 2, 284, 60, 3, 2, 2, 2,
	285, 289, 7, 36, 2, 2, 286, 288, 10, 3, 2, 2, 287, 286, 3, 2, 2, 2, 288,
	291, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 292,
	3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 292, 293, 7, 36, 2, 2, 293, 62, 3, 2,
	2, 2, 294, 295, 7, 118, 2, 2, 295, 296, 7, 116, 2, 2, 296, 297, 7, 119,
	2, 2, 297, 304, 7, 103, 2, 2, 298, 299, 7, 104, 2, 2, 299, 300, 7, 99,
	2, 2, 300, 301, 7, 110, 2, 2, 301, 302, 7, 117, 2, 2, 302, 304, 7, 103,
	2, 2, 303, 294, 3, 2, 2, 2, 303, 298, 3, 2, 2, 2, 304, 64, 3, 2, 2, 2,
	305, 309, 9, 4, 2, 2, 306, 308, 9, 5, 2, 2, 307, 306, 3, 2, 2, 2, 308,
	311, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 66, 3,
	2, 2, 2, 311, 309, 3, 2, 2, 2, 312, 313, 7, 48, 2, 2, 313, 314, 7, 48,
	2, 2, 314, 68, 3, 2, 2, 2, 315, 316, 7, 48, 2, 2, 316, 70, 3, 2, 2, 2,
	317, 318, 7, 61, 2, 2, 318, 72, 3, 2, 2, 2, 319, 320, 7, 46, 2, 2, 320,
	74, 3, 2, 2, 2, 321, 322, 7, 60, 2, 2, 322, 76, 3, 2, 2, 2, 323, 324, 7,
	63, 2, 2, 324, 78, 3, 2, 2, 2, 325, 326, 7, 63, 2, 2, 326, 327, 7, 63,
	2, 2, 327, 80, 3, 2, 2, 2, 328, 329, 7, 64, 2, 2, 329, 330, 7, 63, 2, 2,
	330, 82, 3, 2, 2, 2, 331, 332, 7, 63, 2, 2, 332, 333, 7, 64, 2, 2, 333,
	84, 3, 2, 2, 2, 334, 335, 7, 47, 2, 2, 335, 336, 7, 64, 2, 2, 336, 86,
	3, 2, 2, 2, 337, 338, 7, 62, 2, 2, 338, 339, 7, 63, 2, 2, 339, 88, 3, 2,
	2, 2, 340, 341, 7, 35, 2, 2, 341, 342, 7, 63, 2, 2, 342, 90, 3, 2, 2, 2,
	343, 344, 7, 64, 2, 2, 344, 92, 3, 2, 2, 2, 345, 346, 7, 62, 2, 2, 346,
	94, 3, 2, 2, 2, 347, 348, 7, 44, 2, 2, 348, 96, 3, 2, 2, 2, 349, 350, 7,
	49, 2, 2, 350, 98, 3, 2, 2, 2, 351, 352, 7, 39, 2, 2, 352, 100, 3, 2, 2,
	2, 353, 354, 7, 45, 2, 2, 354, 102, 3, 2, 2, 2, 355, 356, 7, 47, 2, 2,
	356, 104, 3, 2, 2, 2, 357, 358, 7, 42, 2, 2, 358, 106, 3, 2, 2, 2, 359,
	360, 7, 43, 2, 2, 360, 108, 3, 2, 2, 2, 361, 362, 7, 125, 2, 2, 362, 110,
	3, 2, 2, 2, 363, 364, 7, 127, 2, 2, 364, 112, 3, 2, 2, 2, 365, 366, 7,
	93, 2, 2, 366, 114, 3, 2, 2, 2, 367, 368, 7, 95, 2, 2, 368, 116, 3, 2,
	2, 2, 369, 370, 7, 40, 2, 2, 370, 371, 7, 40, 2, 2, 371, 118, 3, 2, 2,
	2, 372, 373, 7, 126, 2, 2, 373, 374, 7, 126, 2, 2, 374, 120, 3, 2, 2, 2,
	375, 376, 7, 126, 2, 2, 376, 122, 3, 2, 2, 2, 377, 378, 7, 35, 2, 2, 378,
	124, 3, 2, 2, 2, 379, 381, 9, 6, 2, 2, 380, 379, 3, 2, 2, 2, 381, 382,
	3, 2, 2, 2, 382, 380, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 384, 3, 2,
	2, 2, 384, 385, 8, 63, 2, 2, 385, 126, 3, 2, 2, 2, 386, 387, 7, 49, 2,
	2, 387, 388, 7, 44, 2, 2, 388, 390, 3, 2, 2, 2, 389, 391, 10, 7, 2, 2,
	390, 389, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 390, 3, 2, 2, 2, 392,
	393, 3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 395, 7, 44, 2, 2, 395, 396,
	7, 49, 2, 2, 396, 397, 3, 2, 2, 2, 397, 398, 8, 64, 2, 2, 398, 128, 3,
	2, 2, 2, 399, 400, 7, 49, 2, 2, 400, 401, 7, 49, 2, 2, 401, 403, 3, 2,
	2, 2, 402, 404, 10, 8, 2, 2, 403, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2,
	405, 403, 3, 2, 2, 2, 405, 406, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407,
	408, 8, 65, 2, 2, 408, 130, 3, 2, 2, 2, 409, 410, 7, 94, 2, 2, 410, 411,
	9, 9, 2, 2, 411, 132, 3, 2, 2, 2, 12, 2, 268, 273, 279, 289, 303, 309,
	382, 392, 405, 3, 8, 2, 2,
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
	"", "'print!'", "'println!'", "'let'", "'mut'", "'if'", "'else'", "'match'",
	"'while'", "'loop'", "'for'", "'in'", "'break'", "'continue'", "'return'",
	"'main'", "'fn'", "'pow'", "'powf'", "'abs'", "'as'", "'i64'", "'f64'",
	"'String'", "'bool'", "'char'", "'&str'", "", "", "", "", "", "", "'..'",
	"'.'", "';'", "','", "':'", "'='", "'=='", "'>='", "'=>'", "'->'", "'<='",
	"'!='", "'>'", "'<'", "'*'", "'/'", "'%'", "'+'", "'-'", "'('", "')'",
	"'{'", "'}'", "'['", "']'", "'&&'", "'||'", "'|'", "'!'",
}

var lexerSymbolicNames = []string{
	"", "R_PRINT", "R_PRINTLN", "R_LET", "R_MUT", "R_IF", "R_ELSE", "R_MATCH",
	"R_WHILE", "R_LOOP", "R_FOR", "R_IN", "R_BREAK", "R_CONTINUE", "R_RETURN",
	"R_MAIN", "R_FUNCTION", "R_POW", "R_POWF", "R_ABS", "R_AS", "R_INT", "R_FLOAT",
	"R_STRING", "R_BOOL", "R_CHAR", "R_STR", "NUMBER", "DOUBLE", "CHAR", "STRING",
	"BOOLEAN", "ID", "TK_DOBLEPUNTO", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA",
	"TK_DOSPUNTOS", "TK_IGUAL", "TK_IGUALIGUAL", "TK_MAYORIGUAL", "TK_IGUALMAYOR",
	"TK_MENOSMAYOR", "TK_MENORIGUAL", "TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR",
	"TK_MULT", "TK_DIV", "TK_MODULO", "TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC",
	"TK_LLAVEA", "TK_LLAVEC", "TK_CORA", "TK_CORC", "TK_AND", "TK_OR", "TK_BARRA",
	"TK_NOT", "WHITESPACE", "TK_MULTI", "TK_LINE",
}

var lexerRuleNames = []string{
	"R_PRINT", "R_PRINTLN", "R_LET", "R_MUT", "R_IF", "R_ELSE", "R_MATCH",
	"R_WHILE", "R_LOOP", "R_FOR", "R_IN", "R_BREAK", "R_CONTINUE", "R_RETURN",
	"R_MAIN", "R_FUNCTION", "R_POW", "R_POWF", "R_ABS", "R_AS", "R_INT", "R_FLOAT",
	"R_STRING", "R_BOOL", "R_CHAR", "R_STR", "NUMBER", "DOUBLE", "CHAR", "STRING",
	"BOOLEAN", "ID", "TK_DOBLEPUNTO", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA",
	"TK_DOSPUNTOS", "TK_IGUAL", "TK_IGUALIGUAL", "TK_MAYORIGUAL", "TK_IGUALMAYOR",
	"TK_MENOSMAYOR", "TK_MENORIGUAL", "TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR",
	"TK_MULT", "TK_DIV", "TK_MODULO", "TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC",
	"TK_LLAVEA", "TK_LLAVEC", "TK_CORA", "TK_CORC", "TK_AND", "TK_OR", "TK_BARRA",
	"TK_NOT", "WHITESPACE", "TK_MULTI", "TK_LINE", "ESC_SEQ",
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
	ChemsLexerR_PRINT       = 1
	ChemsLexerR_PRINTLN     = 2
	ChemsLexerR_LET         = 3
	ChemsLexerR_MUT         = 4
	ChemsLexerR_IF          = 5
	ChemsLexerR_ELSE        = 6
	ChemsLexerR_MATCH       = 7
	ChemsLexerR_WHILE       = 8
	ChemsLexerR_LOOP        = 9
	ChemsLexerR_FOR         = 10
	ChemsLexerR_IN          = 11
	ChemsLexerR_BREAK       = 12
	ChemsLexerR_CONTINUE    = 13
	ChemsLexerR_RETURN      = 14
	ChemsLexerR_MAIN        = 15
	ChemsLexerR_FUNCTION    = 16
	ChemsLexerR_POW         = 17
	ChemsLexerR_POWF        = 18
	ChemsLexerR_ABS         = 19
	ChemsLexerR_AS          = 20
	ChemsLexerR_INT         = 21
	ChemsLexerR_FLOAT       = 22
	ChemsLexerR_STRING      = 23
	ChemsLexerR_BOOL        = 24
	ChemsLexerR_CHAR        = 25
	ChemsLexerR_STR         = 26
	ChemsLexerNUMBER        = 27
	ChemsLexerDOUBLE        = 28
	ChemsLexerCHAR          = 29
	ChemsLexerSTRING        = 30
	ChemsLexerBOOLEAN       = 31
	ChemsLexerID            = 32
	ChemsLexerTK_DOBLEPUNTO = 33
	ChemsLexerTK_PUNTO      = 34
	ChemsLexerTK_PUNTOCOMA  = 35
	ChemsLexerTK_COMA       = 36
	ChemsLexerTK_DOSPUNTOS  = 37
	ChemsLexerTK_IGUAL      = 38
	ChemsLexerTK_IGUALIGUAL = 39
	ChemsLexerTK_MAYORIGUAL = 40
	ChemsLexerTK_IGUALMAYOR = 41
	ChemsLexerTK_MENOSMAYOR = 42
	ChemsLexerTK_MENORIGUAL = 43
	ChemsLexerTK_DIFIGUAL   = 44
	ChemsLexerTK_MAYOR      = 45
	ChemsLexerTK_MENOR      = 46
	ChemsLexerTK_MULT       = 47
	ChemsLexerTK_DIV        = 48
	ChemsLexerTK_MODULO     = 49
	ChemsLexerTK_MAS        = 50
	ChemsLexerTK_MENOS      = 51
	ChemsLexerTK_PARA       = 52
	ChemsLexerTK_PARC       = 53
	ChemsLexerTK_LLAVEA     = 54
	ChemsLexerTK_LLAVEC     = 55
	ChemsLexerTK_CORA       = 56
	ChemsLexerTK_CORC       = 57
	ChemsLexerTK_AND        = 58
	ChemsLexerTK_OR         = 59
	ChemsLexerTK_BARRA      = 60
	ChemsLexerTK_NOT        = 61
	ChemsLexerWHITESPACE    = 62
	ChemsLexerTK_MULTI      = 63
	ChemsLexerTK_LINE       = 64
)
