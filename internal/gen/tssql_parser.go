// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // TsSqlParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TsSqlParser struct {
	*antlr.BaseParser
}

var tssqlparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tssqlparserParserInit() {
	staticData := &tssqlparserParserStaticData
	staticData.literalNames = []string{
		"", "", "", "", "", "'.PROPERTY'", "'.SERVICE'", "'.EVENT'", "'FALSE'",
		"'TRUE'", "'SELECT'", "'FROM'", "'WHERE'", "'INTERVAL'", "'AND'", "'IN'",
		"'IS'", "'OR'", "'AS'", "'XOR'", "'JOIN'", "'LEFT'", "'LIKE'", "'LIMIT'",
		"'NOT'", "'*'", "'/'", "'%'", "'+'", "'-'", "'DIV'", "'MOD'", "'ALL'",
		"'AVG'", "'COUNT'", "'MAX'", "'MIN'", "'SUM'", "'STD'", "'STDDEV'",
		"'DISTINCT'", "'='", "'>'", "'<'", "'!'", "'~'", "'|'", "'&'", "'^'",
		"", "'.'", "'('", "')'", "','", "';'", "'@'", "'0'", "'1'", "'2'", "'''",
		"'\"'", "'`'", "':'", "", "'NULL'", "", "", "'PROPERTY'", "'SERVICE'",
	}
	staticData.symbolicNames = []string{
		"", "SPACE", "SPEC_MYSQL_COMMENT", "COMMENT_INPUT", "LINE_COMMENT",
		"DOT_PROPERTY", "DOT_SERVICE", "DOT_EVENT", "FALSE", "TRUE", "SELECT",
		"FROM", "WHERE", "INTERVAL", "AND", "IN", "IS", "OR", "AS", "XOR", "JOIN",
		"LEFT", "LIKE", "LIMIT", "NOT", "STAR", "DIVIDE", "MODULE", "PLUS",
		"MINUS", "DIV", "MOD", "ALL", "AVG", "COUNT", "MAX", "MIN", "SUM", "STD",
		"STDDEV", "DISTINCT", "EQUAL_SYMBOL", "GREATER_SYMBOL", "LESS_SYMBOL",
		"EXCLAMATION_SYMBOL", "BIT_NOT_OP", "BIT_OR_OP", "BIT_AND_OP", "BIT_XOR_OP",
		"TIME_INTERVAL", "DOT", "LR_BRACKET", "RR_BRACKET", "COMMA", "SEMI",
		"AT_SIGN", "ZERO_DECIMAL", "ONE_DECIMAL", "TWO_DECIMAL", "SINGLE_QUOTE_SYMB",
		"DOUBLE_QUOTE_SYMB", "REVERSE_QUOTE_SYMB", "COLON_SYMB", "ID", "NULL_LITERAL",
		"NULL_SPEC_LITERAL", "STRING_LITERAL", "PROPERTY", "SERVICE", "DECIMAL_LITERAL",
		"REAL_LITERAL", "ERROR_RECONGNIGION",
	}
	staticData.ruleNames = []string{
		"root", "selectStmt", "selectElements", "selectElement", "fromClause",
		"limitClause", "limitClauseAtom", "windowClause", "uid", "tableName",
		"expression", "predicate", "expressionAtom", "functionCall", "aggregateWindowedFunction",
		"functionArg", "columnName", "logicalOperator", "comparisonOperator",
		"fullColumnName", "constant", "stringLiteral", "decimalLiteral", "booleanLiteral",
		"emptyStatement_",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 226, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 3, 0,
		53, 8, 0, 1, 0, 3, 0, 56, 8, 0, 1, 0, 1, 0, 3, 0, 60, 8, 0, 1, 0, 5, 0,
		63, 8, 0, 10, 0, 12, 0, 66, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 72, 8,
		1, 1, 1, 3, 1, 75, 8, 1, 1, 2, 1, 2, 3, 2, 79, 8, 2, 1, 2, 1, 2, 5, 2,
		83, 8, 2, 10, 2, 12, 2, 86, 9, 2, 1, 3, 1, 3, 3, 3, 90, 8, 3, 1, 3, 3,
		3, 93, 8, 3, 1, 3, 1, 3, 3, 3, 97, 8, 3, 1, 3, 3, 3, 100, 8, 3, 3, 3, 102,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 108, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 114, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 137, 8, 10, 10, 10, 12, 10, 140, 9, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 149, 8, 11, 10, 11, 12, 11, 152,
		9, 11, 1, 12, 1, 12, 3, 12, 156, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 169, 8, 14, 1, 14,
		3, 14, 172, 8, 14, 1, 15, 1, 15, 3, 15, 176, 8, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 187, 8, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 3, 18, 203, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 213, 8, 20, 1, 20, 3, 20, 216, 8, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 0, 2, 20, 22, 25,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 0, 6, 2, 0, 56, 58, 69, 69, 1, 0, 5, 7, 2, 0, 33,
		33, 35, 39, 1, 0, 64, 65, 2, 0, 56, 58, 69, 70, 1, 0, 8, 9, 238, 0, 55,
		1, 0, 0, 0, 2, 67, 1, 0, 0, 0, 4, 78, 1, 0, 0, 0, 6, 101, 1, 0, 0, 0, 8,
		103, 1, 0, 0, 0, 10, 109, 1, 0, 0, 0, 12, 117, 1, 0, 0, 0, 14, 119, 1,
		0, 0, 0, 16, 124, 1, 0, 0, 0, 18, 126, 1, 0, 0, 0, 20, 129, 1, 0, 0, 0,
		22, 141, 1, 0, 0, 0, 24, 155, 1, 0, 0, 0, 26, 157, 1, 0, 0, 0, 28, 171,
		1, 0, 0, 0, 30, 175, 1, 0, 0, 0, 32, 177, 1, 0, 0, 0, 34, 186, 1, 0, 0,
		0, 36, 202, 1, 0, 0, 0, 38, 204, 1, 0, 0, 0, 40, 215, 1, 0, 0, 0, 42, 217,
		1, 0, 0, 0, 44, 219, 1, 0, 0, 0, 46, 221, 1, 0, 0, 0, 48, 223, 1, 0, 0,
		0, 50, 52, 3, 2, 1, 0, 51, 53, 5, 54, 0, 0, 52, 51, 1, 0, 0, 0, 52, 53,
		1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 56, 3, 48, 24, 0, 55, 50, 1, 0, 0,
		0, 55, 54, 1, 0, 0, 0, 56, 64, 1, 0, 0, 0, 57, 59, 3, 2, 1, 0, 58, 60,
		5, 54, 0, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0,
		61, 63, 3, 48, 24, 0, 62, 57, 1, 0, 0, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1,
		0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 1, 1, 0, 0, 0, 66,
		64, 1, 0, 0, 0, 67, 68, 5, 10, 0, 0, 68, 69, 3, 4, 2, 0, 69, 71, 3, 8,
		4, 0, 70, 72, 3, 10, 5, 0, 71, 70, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72,
		74, 1, 0, 0, 0, 73, 75, 3, 14, 7, 0, 74, 73, 1, 0, 0, 0, 74, 75, 1, 0,
		0, 0, 75, 3, 1, 0, 0, 0, 76, 79, 5, 25, 0, 0, 77, 79, 3, 6, 3, 0, 78, 76,
		1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 84, 1, 0, 0, 0, 80, 81, 5, 53, 0, 0,
		81, 83, 3, 6, 3, 0, 82, 80, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1,
		0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 5, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87,
		92, 3, 38, 19, 0, 88, 90, 5, 18, 0, 0, 89, 88, 1, 0, 0, 0, 89, 90, 1, 0,
		0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 3, 16, 8, 0, 92, 89, 1, 0, 0, 0, 92,
		93, 1, 0, 0, 0, 93, 102, 1, 0, 0, 0, 94, 99, 3, 26, 13, 0, 95, 97, 5, 18,
		0, 0, 96, 95, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 100,
		3, 16, 8, 0, 99, 96, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 102, 1, 0, 0,
		0, 101, 87, 1, 0, 0, 0, 101, 94, 1, 0, 0, 0, 102, 7, 1, 0, 0, 0, 103, 104,
		5, 11, 0, 0, 104, 107, 3, 18, 9, 0, 105, 106, 5, 12, 0, 0, 106, 108, 3,
		20, 10, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 9, 1, 0, 0,
		0, 109, 113, 5, 23, 0, 0, 110, 111, 3, 12, 6, 0, 111, 112, 5, 53, 0, 0,
		112, 114, 1, 0, 0, 0, 113, 110, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 116, 3, 12, 6, 0, 116, 11, 1, 0, 0, 0, 117, 118,
		7, 0, 0, 0, 118, 13, 1, 0, 0, 0, 119, 120, 5, 13, 0, 0, 120, 121, 5, 51,
		0, 0, 121, 122, 5, 49, 0, 0, 122, 123, 5, 52, 0, 0, 123, 15, 1, 0, 0, 0,
		124, 125, 5, 63, 0, 0, 125, 17, 1, 0, 0, 0, 126, 127, 3, 16, 8, 0, 127,
		128, 7, 1, 0, 0, 128, 19, 1, 0, 0, 0, 129, 130, 6, 10, -1, 0, 130, 131,
		3, 22, 11, 0, 131, 138, 1, 0, 0, 0, 132, 133, 10, 2, 0, 0, 133, 134, 3,
		34, 17, 0, 134, 135, 3, 20, 10, 3, 135, 137, 1, 0, 0, 0, 136, 132, 1, 0,
		0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0,
		139, 21, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 6, 11, -1, 0, 142,
		143, 3, 24, 12, 0, 143, 150, 1, 0, 0, 0, 144, 145, 10, 2, 0, 0, 145, 146,
		3, 36, 18, 0, 146, 147, 3, 22, 11, 3, 147, 149, 1, 0, 0, 0, 148, 144, 1,
		0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0,
		0, 151, 23, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 156, 3, 40, 20, 0, 154,
		156, 3, 32, 16, 0, 155, 153, 1, 0, 0, 0, 155, 154, 1, 0, 0, 0, 156, 25,
		1, 0, 0, 0, 157, 158, 3, 28, 14, 0, 158, 27, 1, 0, 0, 0, 159, 160, 7, 2,
		0, 0, 160, 161, 5, 51, 0, 0, 161, 162, 3, 30, 15, 0, 162, 163, 5, 52, 0,
		0, 163, 172, 1, 0, 0, 0, 164, 165, 5, 34, 0, 0, 165, 168, 5, 51, 0, 0,
		166, 169, 5, 25, 0, 0, 167, 169, 3, 30, 15, 0, 168, 166, 1, 0, 0, 0, 168,
		167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 172, 5, 52, 0, 0, 171, 159,
		1, 0, 0, 0, 171, 164, 1, 0, 0, 0, 172, 29, 1, 0, 0, 0, 173, 176, 3, 38,
		19, 0, 174, 176, 3, 40, 20, 0, 175, 173, 1, 0, 0, 0, 175, 174, 1, 0, 0,
		0, 176, 31, 1, 0, 0, 0, 177, 178, 3, 16, 8, 0, 178, 33, 1, 0, 0, 0, 179,
		187, 5, 14, 0, 0, 180, 181, 5, 47, 0, 0, 181, 187, 5, 47, 0, 0, 182, 187,
		5, 19, 0, 0, 183, 187, 5, 17, 0, 0, 184, 185, 5, 46, 0, 0, 185, 187, 5,
		46, 0, 0, 186, 179, 1, 0, 0, 0, 186, 180, 1, 0, 0, 0, 186, 182, 1, 0, 0,
		0, 186, 183, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 187, 35, 1, 0, 0, 0, 188,
		203, 5, 41, 0, 0, 189, 203, 5, 42, 0, 0, 190, 203, 5, 43, 0, 0, 191, 192,
		5, 43, 0, 0, 192, 203, 5, 41, 0, 0, 193, 194, 5, 42, 0, 0, 194, 203, 5,
		41, 0, 0, 195, 196, 5, 43, 0, 0, 196, 203, 5, 42, 0, 0, 197, 198, 5, 44,
		0, 0, 198, 203, 5, 41, 0, 0, 199, 200, 5, 43, 0, 0, 200, 201, 5, 41, 0,
		0, 201, 203, 5, 42, 0, 0, 202, 188, 1, 0, 0, 0, 202, 189, 1, 0, 0, 0, 202,
		190, 1, 0, 0, 0, 202, 191, 1, 0, 0, 0, 202, 193, 1, 0, 0, 0, 202, 195,
		1, 0, 0, 0, 202, 197, 1, 0, 0, 0, 202, 199, 1, 0, 0, 0, 203, 37, 1, 0,
		0, 0, 204, 205, 3, 16, 8, 0, 205, 39, 1, 0, 0, 0, 206, 216, 3, 42, 21,
		0, 207, 216, 3, 44, 22, 0, 208, 209, 5, 29, 0, 0, 209, 216, 3, 44, 22,
		0, 210, 216, 3, 46, 23, 0, 211, 213, 5, 24, 0, 0, 212, 211, 1, 0, 0, 0,
		212, 213, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 216, 7, 3, 0, 0, 215,
		206, 1, 0, 0, 0, 215, 207, 1, 0, 0, 0, 215, 208, 1, 0, 0, 0, 215, 210,
		1, 0, 0, 0, 215, 212, 1, 0, 0, 0, 216, 41, 1, 0, 0, 0, 217, 218, 5, 66,
		0, 0, 218, 43, 1, 0, 0, 0, 219, 220, 7, 4, 0, 0, 220, 45, 1, 0, 0, 0, 221,
		222, 7, 5, 0, 0, 222, 47, 1, 0, 0, 0, 223, 224, 5, 54, 0, 0, 224, 49, 1,
		0, 0, 0, 26, 52, 55, 59, 62, 64, 71, 74, 78, 84, 89, 92, 96, 99, 101, 107,
		113, 138, 150, 155, 168, 171, 175, 186, 202, 212, 215,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TsSqlParserInit initializes any static state used to implement TsSqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTsSqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TsSqlParserInit() {
	staticData := &tssqlparserParserStaticData
	staticData.once.Do(tssqlparserParserInit)
}

// NewTsSqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTsSqlParser(input antlr.TokenStream) *TsSqlParser {
	TsSqlParserInit()
	this := new(TsSqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tssqlparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// TsSqlParser tokens.
const (
	TsSqlParserEOF                = antlr.TokenEOF
	TsSqlParserSPACE              = 1
	TsSqlParserSPEC_MYSQL_COMMENT = 2
	TsSqlParserCOMMENT_INPUT      = 3
	TsSqlParserLINE_COMMENT       = 4
	TsSqlParserDOT_PROPERTY       = 5
	TsSqlParserDOT_SERVICE        = 6
	TsSqlParserDOT_EVENT          = 7
	TsSqlParserFALSE              = 8
	TsSqlParserTRUE               = 9
	TsSqlParserSELECT             = 10
	TsSqlParserFROM               = 11
	TsSqlParserWHERE              = 12
	TsSqlParserINTERVAL           = 13
	TsSqlParserAND                = 14
	TsSqlParserIN                 = 15
	TsSqlParserIS                 = 16
	TsSqlParserOR                 = 17
	TsSqlParserAS                 = 18
	TsSqlParserXOR                = 19
	TsSqlParserJOIN               = 20
	TsSqlParserLEFT               = 21
	TsSqlParserLIKE               = 22
	TsSqlParserLIMIT              = 23
	TsSqlParserNOT                = 24
	TsSqlParserSTAR               = 25
	TsSqlParserDIVIDE             = 26
	TsSqlParserMODULE             = 27
	TsSqlParserPLUS               = 28
	TsSqlParserMINUS              = 29
	TsSqlParserDIV                = 30
	TsSqlParserMOD                = 31
	TsSqlParserALL                = 32
	TsSqlParserAVG                = 33
	TsSqlParserCOUNT              = 34
	TsSqlParserMAX                = 35
	TsSqlParserMIN                = 36
	TsSqlParserSUM                = 37
	TsSqlParserSTD                = 38
	TsSqlParserSTDDEV             = 39
	TsSqlParserDISTINCT           = 40
	TsSqlParserEQUAL_SYMBOL       = 41
	TsSqlParserGREATER_SYMBOL     = 42
	TsSqlParserLESS_SYMBOL        = 43
	TsSqlParserEXCLAMATION_SYMBOL = 44
	TsSqlParserBIT_NOT_OP         = 45
	TsSqlParserBIT_OR_OP          = 46
	TsSqlParserBIT_AND_OP         = 47
	TsSqlParserBIT_XOR_OP         = 48
	TsSqlParserTIME_INTERVAL      = 49
	TsSqlParserDOT                = 50
	TsSqlParserLR_BRACKET         = 51
	TsSqlParserRR_BRACKET         = 52
	TsSqlParserCOMMA              = 53
	TsSqlParserSEMI               = 54
	TsSqlParserAT_SIGN            = 55
	TsSqlParserZERO_DECIMAL       = 56
	TsSqlParserONE_DECIMAL        = 57
	TsSqlParserTWO_DECIMAL        = 58
	TsSqlParserSINGLE_QUOTE_SYMB  = 59
	TsSqlParserDOUBLE_QUOTE_SYMB  = 60
	TsSqlParserREVERSE_QUOTE_SYMB = 61
	TsSqlParserCOLON_SYMB         = 62
	TsSqlParserID                 = 63
	TsSqlParserNULL_LITERAL       = 64
	TsSqlParserNULL_SPEC_LITERAL  = 65
	TsSqlParserSTRING_LITERAL     = 66
	TsSqlParserPROPERTY           = 67
	TsSqlParserSERVICE            = 68
	TsSqlParserDECIMAL_LITERAL    = 69
	TsSqlParserREAL_LITERAL       = 70
	TsSqlParserERROR_RECONGNIGION = 71
)

// TsSqlParser rules.
const (
	TsSqlParserRULE_root                      = 0
	TsSqlParserRULE_selectStmt                = 1
	TsSqlParserRULE_selectElements            = 2
	TsSqlParserRULE_selectElement             = 3
	TsSqlParserRULE_fromClause                = 4
	TsSqlParserRULE_limitClause               = 5
	TsSqlParserRULE_limitClauseAtom           = 6
	TsSqlParserRULE_windowClause              = 7
	TsSqlParserRULE_uid                       = 8
	TsSqlParserRULE_tableName                 = 9
	TsSqlParserRULE_expression                = 10
	TsSqlParserRULE_predicate                 = 11
	TsSqlParserRULE_expressionAtom            = 12
	TsSqlParserRULE_functionCall              = 13
	TsSqlParserRULE_aggregateWindowedFunction = 14
	TsSqlParserRULE_functionArg               = 15
	TsSqlParserRULE_columnName                = 16
	TsSqlParserRULE_logicalOperator           = 17
	TsSqlParserRULE_comparisonOperator        = 18
	TsSqlParserRULE_fullColumnName            = 19
	TsSqlParserRULE_constant                  = 20
	TsSqlParserRULE_stringLiteral             = 21
	TsSqlParserRULE_decimalLiteral            = 22
	TsSqlParserRULE_booleanLiteral            = 23
	TsSqlParserRULE_emptyStatement_           = 24
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) AllSelectStmt() []ISelectStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectStmtContext); ok {
			len++
		}
	}

	tst := make([]ISelectStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectStmtContext); ok {
			tst[i] = t.(ISelectStmtContext)
			i++
		}
	}

	return tst
}

func (s *RootContext) SelectStmt(i int) ISelectStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStmtContext)
}

func (s *RootContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *RootContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *RootContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(TsSqlParserSEMI)
}

func (s *RootContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(TsSqlParserSEMI, i)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) Root() (localctx IRootContext) {
	this := p
	_ = this

	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TsSqlParserRULE_root)
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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TsSqlParserSELECT:
		{
			p.SetState(50)
			p.SelectStmt()
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(51)
				p.Match(TsSqlParserSEMI)
			}

		}

	case TsSqlParserSEMI:
		{
			p.SetState(54)
			p.EmptyStatement_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TsSqlParserSELECT || _la == TsSqlParserSEMI {
		p.SetState(62)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case TsSqlParserSELECT:
			{
				p.SetState(57)
				p.SelectStmt()
			}
			p.SetState(59)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(58)
					p.Match(TsSqlParserSEMI)
				}

			}

		case TsSqlParserSEMI:
			{
				p.SetState(61)
				p.EmptyStatement_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectStmtContext is an interface to support dynamic dispatch.
type ISelectStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectStmtContext differentiates from other interfaces.
	IsSelectStmtContext()
}

type SelectStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStmtContext() *SelectStmtContext {
	var p = new(SelectStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_selectStmt
	return p
}

func (*SelectStmtContext) IsSelectStmtContext() {}

func NewSelectStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStmtContext {
	var p = new(SelectStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_selectStmt

	return p
}

func (s *SelectStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStmtContext) SELECT() antlr.TerminalNode {
	return s.GetToken(TsSqlParserSELECT, 0)
}

func (s *SelectStmtContext) SelectElements() ISelectElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectElementsContext)
}

func (s *SelectStmtContext) FromClause() IFromClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFromClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFromClauseContext)
}

func (s *SelectStmtContext) LimitClause() ILimitClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *SelectStmtContext) WindowClause() IWindowClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindowClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWindowClauseContext)
}

func (s *SelectStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterSelectStmt(s)
	}
}

func (s *SelectStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitSelectStmt(s)
	}
}

func (s *SelectStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitSelectStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) SelectStmt() (localctx ISelectStmtContext) {
	this := p
	_ = this

	localctx = NewSelectStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TsSqlParserRULE_selectStmt)
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
	{
		p.SetState(67)
		p.Match(TsSqlParserSELECT)
	}
	{
		p.SetState(68)
		p.SelectElements()
	}
	{
		p.SetState(69)
		p.FromClause()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TsSqlParserLIMIT {
		{
			p.SetState(70)
			p.LimitClause()
		}

	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TsSqlParserINTERVAL {
		{
			p.SetState(73)
			p.WindowClause()
		}

	}

	return localctx
}

// ISelectElementsContext is an interface to support dynamic dispatch.
type ISelectElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStar returns the star token.
	GetStar() antlr.Token

	// SetStar sets the star token.
	SetStar(antlr.Token)

	// IsSelectElementsContext differentiates from other interfaces.
	IsSelectElementsContext()
}

type SelectElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	star   antlr.Token
}

func NewEmptySelectElementsContext() *SelectElementsContext {
	var p = new(SelectElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_selectElements
	return p
}

func (*SelectElementsContext) IsSelectElementsContext() {}

func NewSelectElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementsContext {
	var p = new(SelectElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_selectElements

	return p
}

func (s *SelectElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementsContext) GetStar() antlr.Token { return s.star }

func (s *SelectElementsContext) SetStar(v antlr.Token) { s.star = v }

func (s *SelectElementsContext) AllSelectElement() []ISelectElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectElementContext); ok {
			len++
		}
	}

	tst := make([]ISelectElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectElementContext); ok {
			tst[i] = t.(ISelectElementContext)
			i++
		}
	}

	return tst
}

func (s *SelectElementsContext) SelectElement(i int) ISelectElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectElementContext)
}

func (s *SelectElementsContext) STAR() antlr.TerminalNode {
	return s.GetToken(TsSqlParserSTAR, 0)
}

func (s *SelectElementsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TsSqlParserCOMMA)
}

func (s *SelectElementsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TsSqlParserCOMMA, i)
}

func (s *SelectElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterSelectElements(s)
	}
}

func (s *SelectElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitSelectElements(s)
	}
}

func (s *SelectElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitSelectElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) SelectElements() (localctx ISelectElementsContext) {
	this := p
	_ = this

	localctx = NewSelectElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TsSqlParserRULE_selectElements)
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
	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TsSqlParserSTAR:
		{
			p.SetState(76)

			var _m = p.Match(TsSqlParserSTAR)

			localctx.(*SelectElementsContext).star = _m
		}

	case TsSqlParserAVG, TsSqlParserCOUNT, TsSqlParserMAX, TsSqlParserMIN, TsSqlParserSUM, TsSqlParserSTD, TsSqlParserSTDDEV, TsSqlParserID:
		{
			p.SetState(77)
			p.SelectElement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TsSqlParserCOMMA {
		{
			p.SetState(80)
			p.Match(TsSqlParserCOMMA)
		}
		{
			p.SetState(81)
			p.SelectElement()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectElementContext is an interface to support dynamic dispatch.
type ISelectElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectElementContext differentiates from other interfaces.
	IsSelectElementContext()
}

type SelectElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectElementContext() *SelectElementContext {
	var p = new(SelectElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_selectElement
	return p
}

func (*SelectElementContext) IsSelectElementContext() {}

func NewSelectElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementContext {
	var p = new(SelectElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_selectElement

	return p
}

func (s *SelectElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementContext) CopyFrom(ctx *SelectElementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SelectElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelectFunctionElementContext struct {
	*SelectElementContext
}

func NewSelectFunctionElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectFunctionElementContext {
	var p = new(SelectFunctionElementContext)

	p.SelectElementContext = NewEmptySelectElementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectElementContext))

	return p
}

func (s *SelectFunctionElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectFunctionElementContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *SelectFunctionElementContext) Uid() IUidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUidContext)
}

func (s *SelectFunctionElementContext) AS() antlr.TerminalNode {
	return s.GetToken(TsSqlParserAS, 0)
}

func (s *SelectFunctionElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterSelectFunctionElement(s)
	}
}

func (s *SelectFunctionElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitSelectFunctionElement(s)
	}
}

func (s *SelectFunctionElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitSelectFunctionElement(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectColumnElementContext struct {
	*SelectElementContext
}

func NewSelectColumnElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectColumnElementContext {
	var p = new(SelectColumnElementContext)

	p.SelectElementContext = NewEmptySelectElementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectElementContext))

	return p
}

func (s *SelectColumnElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectColumnElementContext) FullColumnName() IFullColumnNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullColumnNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullColumnNameContext)
}

func (s *SelectColumnElementContext) Uid() IUidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUidContext)
}

func (s *SelectColumnElementContext) AS() antlr.TerminalNode {
	return s.GetToken(TsSqlParserAS, 0)
}

func (s *SelectColumnElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterSelectColumnElement(s)
	}
}

func (s *SelectColumnElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitSelectColumnElement(s)
	}
}

func (s *SelectColumnElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitSelectColumnElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) SelectElement() (localctx ISelectElementContext) {
	this := p
	_ = this

	localctx = NewSelectElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TsSqlParserRULE_selectElement)
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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TsSqlParserID:
		localctx = NewSelectColumnElementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.FullColumnName()
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TsSqlParserAS || _la == TsSqlParserID {
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == TsSqlParserAS {
				{
					p.SetState(88)
					p.Match(TsSqlParserAS)
				}

			}
			{
				p.SetState(91)
				p.Uid()
			}

		}

	case TsSqlParserAVG, TsSqlParserCOUNT, TsSqlParserMAX, TsSqlParserMIN, TsSqlParserSUM, TsSqlParserSTD, TsSqlParserSTDDEV:
		localctx = NewSelectFunctionElementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.FunctionCall()
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TsSqlParserAS || _la == TsSqlParserID {
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == TsSqlParserAS {
				{
					p.SetState(95)
					p.Match(TsSqlParserAS)
				}

			}
			{
				p.SetState(98)
				p.Uid()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFromClauseContext is an interface to support dynamic dispatch.
type IFromClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFromClauseContext differentiates from other interfaces.
	IsFromClauseContext()
}

type FromClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromClauseContext() *FromClauseContext {
	var p = new(FromClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_fromClause
	return p
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(TsSqlParserFROM, 0)
}

func (s *FromClauseContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *FromClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(TsSqlParserWHERE, 0)
}

func (s *FromClauseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FromClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitFromClause(s)
	}
}

func (s *FromClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitFromClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) FromClause() (localctx IFromClauseContext) {
	this := p
	_ = this

	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TsSqlParserRULE_fromClause)
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
	{
		p.SetState(103)
		p.Match(TsSqlParserFROM)
	}
	{
		p.SetState(104)
		p.TableName()
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TsSqlParserWHERE {
		{
			p.SetState(105)
			p.Match(TsSqlParserWHERE)
		}
		{
			p.SetState(106)
			p.expression(0)
		}

	}

	return localctx
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOffset returns the offset rule contexts.
	GetOffset() ILimitClauseAtomContext

	// GetLimit returns the limit rule contexts.
	GetLimit() ILimitClauseAtomContext

	// SetOffset sets the offset rule contexts.
	SetOffset(ILimitClauseAtomContext)

	// SetLimit sets the limit rule contexts.
	SetLimit(ILimitClauseAtomContext)

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	offset ILimitClauseAtomContext
	limit  ILimitClauseAtomContext
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) GetOffset() ILimitClauseAtomContext { return s.offset }

func (s *LimitClauseContext) GetLimit() ILimitClauseAtomContext { return s.limit }

func (s *LimitClauseContext) SetOffset(v ILimitClauseAtomContext) { s.offset = v }

func (s *LimitClauseContext) SetLimit(v ILimitClauseAtomContext) { s.limit = v }

func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(TsSqlParserLIMIT, 0)
}

func (s *LimitClauseContext) AllLimitClauseAtom() []ILimitClauseAtomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILimitClauseAtomContext); ok {
			len++
		}
	}

	tst := make([]ILimitClauseAtomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILimitClauseAtomContext); ok {
			tst[i] = t.(ILimitClauseAtomContext)
			i++
		}
	}

	return tst
}

func (s *LimitClauseContext) LimitClauseAtom(i int) ILimitClauseAtomContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitClauseAtomContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitClauseAtomContext)
}

func (s *LimitClauseContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TsSqlParserCOMMA, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) LimitClause() (localctx ILimitClauseContext) {
	this := p
	_ = this

	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TsSqlParserRULE_limitClause)

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
		p.SetState(109)
		p.Match(TsSqlParserLIMIT)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(110)

			var _x = p.LimitClauseAtom()

			localctx.(*LimitClauseContext).offset = _x
		}
		{
			p.SetState(111)
			p.Match(TsSqlParserCOMMA)
		}

	}
	{
		p.SetState(115)

		var _x = p.LimitClauseAtom()

		localctx.(*LimitClauseContext).limit = _x
	}

	return localctx
}

// ILimitClauseAtomContext is an interface to support dynamic dispatch.
type ILimitClauseAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitClauseAtomContext differentiates from other interfaces.
	IsLimitClauseAtomContext()
}

type LimitClauseAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseAtomContext() *LimitClauseAtomContext {
	var p = new(LimitClauseAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_limitClauseAtom
	return p
}

func (*LimitClauseAtomContext) IsLimitClauseAtomContext() {}

func NewLimitClauseAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseAtomContext {
	var p = new(LimitClauseAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_limitClauseAtom

	return p
}

func (s *LimitClauseAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseAtomContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserDECIMAL_LITERAL, 0)
}

func (s *LimitClauseAtomContext) TWO_DECIMAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserTWO_DECIMAL, 0)
}

func (s *LimitClauseAtomContext) ONE_DECIMAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserONE_DECIMAL, 0)
}

func (s *LimitClauseAtomContext) ZERO_DECIMAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserZERO_DECIMAL, 0)
}

func (s *LimitClauseAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterLimitClauseAtom(s)
	}
}

func (s *LimitClauseAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitLimitClauseAtom(s)
	}
}

func (s *LimitClauseAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitLimitClauseAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) LimitClauseAtom() (localctx ILimitClauseAtomContext) {
	this := p
	_ = this

	localctx = NewLimitClauseAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TsSqlParserRULE_limitClauseAtom)
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
	{
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-56)) & ^0x3f) == 0 && ((int64(1)<<(_la-56))&8199) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IWindowClauseContext is an interface to support dynamic dispatch.
type IWindowClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindowClauseContext differentiates from other interfaces.
	IsWindowClauseContext()
}

type WindowClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowClauseContext() *WindowClauseContext {
	var p = new(WindowClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_windowClause
	return p
}

func (*WindowClauseContext) IsWindowClauseContext() {}

func NewWindowClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowClauseContext {
	var p = new(WindowClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_windowClause

	return p
}

func (s *WindowClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowClauseContext) INTERVAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserINTERVAL, 0)
}

func (s *WindowClauseContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(TsSqlParserLR_BRACKET, 0)
}

func (s *WindowClauseContext) TIME_INTERVAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserTIME_INTERVAL, 0)
}

func (s *WindowClauseContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(TsSqlParserRR_BRACKET, 0)
}

func (s *WindowClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterWindowClause(s)
	}
}

func (s *WindowClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitWindowClause(s)
	}
}

func (s *WindowClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitWindowClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) WindowClause() (localctx IWindowClauseContext) {
	this := p
	_ = this

	localctx = NewWindowClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TsSqlParserRULE_windowClause)

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
		p.SetState(119)
		p.Match(TsSqlParserINTERVAL)
	}
	{
		p.SetState(120)
		p.Match(TsSqlParserLR_BRACKET)
	}
	{
		p.SetState(121)
		p.Match(TsSqlParserTIME_INTERVAL)
	}
	{
		p.SetState(122)
		p.Match(TsSqlParserRR_BRACKET)
	}

	return localctx
}

// IUidContext is an interface to support dynamic dispatch.
type IUidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUidContext differentiates from other interfaces.
	IsUidContext()
}

type UidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUidContext() *UidContext {
	var p = new(UidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_uid
	return p
}

func (*UidContext) IsUidContext() {}

func NewUidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UidContext {
	var p = new(UidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_uid

	return p
}

func (s *UidContext) GetParser() antlr.Parser { return s.parser }

func (s *UidContext) ID() antlr.TerminalNode {
	return s.GetToken(TsSqlParserID, 0)
}

func (s *UidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterUid(s)
	}
}

func (s *UidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitUid(s)
	}
}

func (s *UidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitUid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) Uid() (localctx IUidContext) {
	this := p
	_ = this

	localctx = NewUidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TsSqlParserRULE_uid)

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
		p.SetState(124)
		p.Match(TsSqlParserID)
	}

	return localctx
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) Uid() IUidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUidContext)
}

func (s *TableNameContext) DOT_PROPERTY() antlr.TerminalNode {
	return s.GetToken(TsSqlParserDOT_PROPERTY, 0)
}

func (s *TableNameContext) DOT_SERVICE() antlr.TerminalNode {
	return s.GetToken(TsSqlParserDOT_SERVICE, 0)
}

func (s *TableNameContext) DOT_EVENT() antlr.TerminalNode {
	return s.GetToken(TsSqlParserDOT_EVENT, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (s *TableNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitTableName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) TableName() (localctx ITableNameContext) {
	this := p
	_ = this

	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TsSqlParserRULE_tableName)
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
	{
		p.SetState(126)
		p.Uid()
	}
	{
		p.SetState(127)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&224) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LogicalExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewLogicalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *LogicalExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *LogicalExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *LogicalExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) LogicalOperator() ILogicalOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOperatorContext)
}

func (s *LogicalExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitLogicalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type PredicateExpressionContext struct {
	*ExpressionContext
}

func NewPredicateExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PredicateExpressionContext {
	var p = new(PredicateExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PredicateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateExpressionContext) Predicate() IPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *PredicateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterPredicateExpression(s)
	}
}

func (s *PredicateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitPredicateExpression(s)
	}
}

func (s *PredicateExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitPredicateExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TsSqlParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, TsSqlParserRULE_expression, _p)

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
	localctx = NewPredicateExpressionContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(130)
		p.predicate(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*LogicalExpressionContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, TsSqlParserRULE_expression)
			p.SetState(132)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(133)
				p.LogicalOperator()
			}
			{
				p.SetState(134)

				var _x = p.expression(3)

				localctx.(*LogicalExpressionContext).right = _x
			}

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) CopyFrom(ctx *PredicateContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionAtomPredicateContext struct {
	*PredicateContext
}

func NewExpressionAtomPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionAtomPredicateContext {
	var p = new(ExpressionAtomPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *ExpressionAtomPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomPredicateContext) ExpressionAtom() IExpressionAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionAtomPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterExpressionAtomPredicate(s)
	}
}

func (s *ExpressionAtomPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitExpressionAtomPredicate(s)
	}
}

func (s *ExpressionAtomPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitExpressionAtomPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryComparisonPredicateContext struct {
	*PredicateContext
	left  IPredicateContext
	right IPredicateContext
}

func NewBinaryComparisonPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryComparisonPredicateContext {
	var p = new(BinaryComparisonPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *BinaryComparisonPredicateContext) GetLeft() IPredicateContext { return s.left }

func (s *BinaryComparisonPredicateContext) GetRight() IPredicateContext { return s.right }

func (s *BinaryComparisonPredicateContext) SetLeft(v IPredicateContext) { s.left = v }

func (s *BinaryComparisonPredicateContext) SetRight(v IPredicateContext) { s.right = v }

func (s *BinaryComparisonPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryComparisonPredicateContext) ComparisonOperator() IComparisonOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *BinaryComparisonPredicateContext) AllPredicate() []IPredicateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPredicateContext); ok {
			len++
		}
	}

	tst := make([]IPredicateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPredicateContext); ok {
			tst[i] = t.(IPredicateContext)
			i++
		}
	}

	return tst
}

func (s *BinaryComparisonPredicateContext) Predicate(i int) IPredicateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredicateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *BinaryComparisonPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterBinaryComparisonPredicate(s)
	}
}

func (s *BinaryComparisonPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitBinaryComparisonPredicate(s)
	}
}

func (s *BinaryComparisonPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitBinaryComparisonPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) Predicate() (localctx IPredicateContext) {
	return p.predicate(0)
}

func (p *TsSqlParser) predicate(_p int) (localctx IPredicateContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPredicateContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, TsSqlParserRULE_predicate, _p)

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
	localctx = NewExpressionAtomPredicateContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(142)
		p.ExpressionAtom()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBinaryComparisonPredicateContext(p, NewPredicateContext(p, _parentctx, _parentState))
			localctx.(*BinaryComparisonPredicateContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, TsSqlParserRULE_predicate)
			p.SetState(144)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(145)
				p.ComparisonOperator()
			}
			{
				p.SetState(146)

				var _x = p.predicate(3)

				localctx.(*BinaryComparisonPredicateContext).right = _x
			}

		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionAtomContext is an interface to support dynamic dispatch.
type IExpressionAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAtomContext differentiates from other interfaces.
	IsExpressionAtomContext()
}

type ExpressionAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAtomContext() *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_expressionAtom
	return p
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_expressionAtom

	return p
}

func (s *ExpressionAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAtomContext) CopyFrom(ctx *ExpressionAtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ColumnNameExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewColumnNameExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnNameExpressionAtomContext {
	var p = new(ColumnNameExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *ColumnNameExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameExpressionAtomContext) ColumnName() IColumnNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ColumnNameExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterColumnNameExpressionAtom(s)
	}
}

func (s *ColumnNameExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitColumnNameExpressionAtom(s)
	}
}

func (s *ColumnNameExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitColumnNameExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstantExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewConstantExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantExpressionAtomContext {
	var p = new(ConstantExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *ConstantExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantExpressionAtomContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ConstantExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterConstantExpressionAtom(s)
	}
}

func (s *ConstantExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitConstantExpressionAtom(s)
	}
}

func (s *ConstantExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitConstantExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) ExpressionAtom() (localctx IExpressionAtomContext) {
	this := p
	_ = this

	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TsSqlParserRULE_expressionAtom)

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

	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TsSqlParserFALSE, TsSqlParserTRUE, TsSqlParserNOT, TsSqlParserMINUS, TsSqlParserZERO_DECIMAL, TsSqlParserONE_DECIMAL, TsSqlParserTWO_DECIMAL, TsSqlParserNULL_LITERAL, TsSqlParserNULL_SPEC_LITERAL, TsSqlParserSTRING_LITERAL, TsSqlParserDECIMAL_LITERAL, TsSqlParserREAL_LITERAL:
		localctx = NewConstantExpressionAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Constant()
		}

	case TsSqlParserID:
		localctx = NewColumnNameExpressionAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.ColumnName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) AggregateWindowedFunction() IAggregateWindowedFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateWindowedFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateWindowedFunctionContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) FunctionCall() (localctx IFunctionCallContext) {
	this := p
	_ = this

	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TsSqlParserRULE_functionCall)

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
		p.SetState(157)
		p.AggregateWindowedFunction()
	}

	return localctx
}

// IAggregateWindowedFunctionContext is an interface to support dynamic dispatch.
type IAggregateWindowedFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStarArg returns the starArg token.
	GetStarArg() antlr.Token

	// SetStarArg sets the starArg token.
	SetStarArg(antlr.Token)

	// IsAggregateWindowedFunctionContext differentiates from other interfaces.
	IsAggregateWindowedFunctionContext()
}

type AggregateWindowedFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	starArg antlr.Token
}

func NewEmptyAggregateWindowedFunctionContext() *AggregateWindowedFunctionContext {
	var p = new(AggregateWindowedFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_aggregateWindowedFunction
	return p
}

func (*AggregateWindowedFunctionContext) IsAggregateWindowedFunctionContext() {}

func NewAggregateWindowedFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateWindowedFunctionContext {
	var p = new(AggregateWindowedFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_aggregateWindowedFunction

	return p
}

func (s *AggregateWindowedFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateWindowedFunctionContext) GetStarArg() antlr.Token { return s.starArg }

func (s *AggregateWindowedFunctionContext) SetStarArg(v antlr.Token) { s.starArg = v }

func (s *AggregateWindowedFunctionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(TsSqlParserLR_BRACKET, 0)
}

func (s *AggregateWindowedFunctionContext) FunctionArg() IFunctionArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgContext)
}

func (s *AggregateWindowedFunctionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(TsSqlParserRR_BRACKET, 0)
}

func (s *AggregateWindowedFunctionContext) AVG() antlr.TerminalNode {
	return s.GetToken(TsSqlParserAVG, 0)
}

func (s *AggregateWindowedFunctionContext) MAX() antlr.TerminalNode {
	return s.GetToken(TsSqlParserMAX, 0)
}

func (s *AggregateWindowedFunctionContext) MIN() antlr.TerminalNode {
	return s.GetToken(TsSqlParserMIN, 0)
}

func (s *AggregateWindowedFunctionContext) SUM() antlr.TerminalNode {
	return s.GetToken(TsSqlParserSUM, 0)
}

func (s *AggregateWindowedFunctionContext) STD() antlr.TerminalNode {
	return s.GetToken(TsSqlParserSTD, 0)
}

func (s *AggregateWindowedFunctionContext) STDDEV() antlr.TerminalNode {
	return s.GetToken(TsSqlParserSTDDEV, 0)
}

func (s *AggregateWindowedFunctionContext) COUNT() antlr.TerminalNode {
	return s.GetToken(TsSqlParserCOUNT, 0)
}

func (s *AggregateWindowedFunctionContext) STAR() antlr.TerminalNode {
	return s.GetToken(TsSqlParserSTAR, 0)
}

func (s *AggregateWindowedFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateWindowedFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateWindowedFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterAggregateWindowedFunction(s)
	}
}

func (s *AggregateWindowedFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitAggregateWindowedFunction(s)
	}
}

func (s *AggregateWindowedFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitAggregateWindowedFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) AggregateWindowedFunction() (localctx IAggregateWindowedFunctionContext) {
	this := p
	_ = this

	localctx = NewAggregateWindowedFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TsSqlParserRULE_aggregateWindowedFunction)
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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TsSqlParserAVG, TsSqlParserMAX, TsSqlParserMIN, TsSqlParserSUM, TsSqlParserSTD, TsSqlParserSTDDEV:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1073741824000) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(160)
			p.Match(TsSqlParserLR_BRACKET)
		}
		{
			p.SetState(161)
			p.FunctionArg()
		}
		{
			p.SetState(162)
			p.Match(TsSqlParserRR_BRACKET)
		}

	case TsSqlParserCOUNT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(TsSqlParserCOUNT)
		}
		{
			p.SetState(165)
			p.Match(TsSqlParserLR_BRACKET)
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case TsSqlParserSTAR:
			{
				p.SetState(166)

				var _m = p.Match(TsSqlParserSTAR)

				localctx.(*AggregateWindowedFunctionContext).starArg = _m
			}

		case TsSqlParserFALSE, TsSqlParserTRUE, TsSqlParserNOT, TsSqlParserMINUS, TsSqlParserZERO_DECIMAL, TsSqlParserONE_DECIMAL, TsSqlParserTWO_DECIMAL, TsSqlParserID, TsSqlParserNULL_LITERAL, TsSqlParserNULL_SPEC_LITERAL, TsSqlParserSTRING_LITERAL, TsSqlParserDECIMAL_LITERAL, TsSqlParserREAL_LITERAL:
			{
				p.SetState(167)
				p.FunctionArg()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(170)
			p.Match(TsSqlParserRR_BRACKET)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionArgContext is an interface to support dynamic dispatch.
type IFunctionArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgContext differentiates from other interfaces.
	IsFunctionArgContext()
}

type FunctionArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgContext() *FunctionArgContext {
	var p = new(FunctionArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_functionArg
	return p
}

func (*FunctionArgContext) IsFunctionArgContext() {}

func NewFunctionArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgContext {
	var p = new(FunctionArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_functionArg

	return p
}

func (s *FunctionArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgContext) FullColumnName() IFullColumnNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullColumnNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullColumnNameContext)
}

func (s *FunctionArgContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FunctionArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterFunctionArg(s)
	}
}

func (s *FunctionArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitFunctionArg(s)
	}
}

func (s *FunctionArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitFunctionArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) FunctionArg() (localctx IFunctionArgContext) {
	this := p
	_ = this

	localctx = NewFunctionArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TsSqlParserRULE_functionArg)

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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TsSqlParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.FullColumnName()
		}

	case TsSqlParserFALSE, TsSqlParserTRUE, TsSqlParserNOT, TsSqlParserMINUS, TsSqlParserZERO_DECIMAL, TsSqlParserONE_DECIMAL, TsSqlParserTWO_DECIMAL, TsSqlParserNULL_LITERAL, TsSqlParserNULL_SPEC_LITERAL, TsSqlParserSTRING_LITERAL, TsSqlParserDECIMAL_LITERAL, TsSqlParserREAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Constant()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_columnName
	return p
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) Uid() IUidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUidContext)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (s *ColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) ColumnName() (localctx IColumnNameContext) {
	this := p
	_ = this

	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TsSqlParserRULE_columnName)

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
		p.SetState(177)
		p.Uid()
	}

	return localctx
}

// ILogicalOperatorContext is an interface to support dynamic dispatch.
type ILogicalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalOperatorContext differentiates from other interfaces.
	IsLogicalOperatorContext()
}

type LogicalOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOperatorContext() *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_logicalOperator
	return p
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(TsSqlParserAND, 0)
}

func (s *LogicalOperatorContext) AllBIT_AND_OP() []antlr.TerminalNode {
	return s.GetTokens(TsSqlParserBIT_AND_OP)
}

func (s *LogicalOperatorContext) BIT_AND_OP(i int) antlr.TerminalNode {
	return s.GetToken(TsSqlParserBIT_AND_OP, i)
}

func (s *LogicalOperatorContext) XOR() antlr.TerminalNode {
	return s.GetToken(TsSqlParserXOR, 0)
}

func (s *LogicalOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(TsSqlParserOR, 0)
}

func (s *LogicalOperatorContext) AllBIT_OR_OP() []antlr.TerminalNode {
	return s.GetTokens(TsSqlParserBIT_OR_OP)
}

func (s *LogicalOperatorContext) BIT_OR_OP(i int) antlr.TerminalNode {
	return s.GetToken(TsSqlParserBIT_OR_OP, i)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitLogicalOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	this := p
	_ = this

	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TsSqlParserRULE_logicalOperator)

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

	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TsSqlParserAND:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.Match(TsSqlParserAND)
		}

	case TsSqlParserBIT_AND_OP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.Match(TsSqlParserBIT_AND_OP)
		}
		{
			p.SetState(181)
			p.Match(TsSqlParserBIT_AND_OP)
		}

	case TsSqlParserXOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Match(TsSqlParserXOR)
		}

	case TsSqlParserOR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.Match(TsSqlParserOR)
		}

	case TsSqlParserBIT_OR_OP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(184)
			p.Match(TsSqlParserBIT_OR_OP)
		}
		{
			p.SetState(185)
			p.Match(TsSqlParserBIT_OR_OP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) EQUAL_SYMBOL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserEQUAL_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) GREATER_SYMBOL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserGREATER_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) LESS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserLESS_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) EXCLAMATION_SYMBOL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserEXCLAMATION_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	this := p
	_ = this

	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TsSqlParserRULE_comparisonOperator)

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

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(TsSqlParserEQUAL_SYMBOL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Match(TsSqlParserGREATER_SYMBOL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(190)
			p.Match(TsSqlParserLESS_SYMBOL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(191)
			p.Match(TsSqlParserLESS_SYMBOL)
		}
		{
			p.SetState(192)
			p.Match(TsSqlParserEQUAL_SYMBOL)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(193)
			p.Match(TsSqlParserGREATER_SYMBOL)
		}
		{
			p.SetState(194)
			p.Match(TsSqlParserEQUAL_SYMBOL)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(195)
			p.Match(TsSqlParserLESS_SYMBOL)
		}
		{
			p.SetState(196)
			p.Match(TsSqlParserGREATER_SYMBOL)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(197)
			p.Match(TsSqlParserEXCLAMATION_SYMBOL)
		}
		{
			p.SetState(198)
			p.Match(TsSqlParserEQUAL_SYMBOL)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(199)
			p.Match(TsSqlParserLESS_SYMBOL)
		}
		{
			p.SetState(200)
			p.Match(TsSqlParserEQUAL_SYMBOL)
		}
		{
			p.SetState(201)
			p.Match(TsSqlParserGREATER_SYMBOL)
		}

	}

	return localctx
}

// IFullColumnNameContext is an interface to support dynamic dispatch.
type IFullColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullColumnNameContext differentiates from other interfaces.
	IsFullColumnNameContext()
}

type FullColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullColumnNameContext() *FullColumnNameContext {
	var p = new(FullColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_fullColumnName
	return p
}

func (*FullColumnNameContext) IsFullColumnNameContext() {}

func NewFullColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullColumnNameContext {
	var p = new(FullColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_fullColumnName

	return p
}

func (s *FullColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FullColumnNameContext) Uid() IUidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUidContext)
}

func (s *FullColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterFullColumnName(s)
	}
}

func (s *FullColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitFullColumnName(s)
	}
}

func (s *FullColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitFullColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) FullColumnName() (localctx IFullColumnNameContext) {
	this := p
	_ = this

	localctx = NewFullColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TsSqlParserRULE_fullColumnName)

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
		p.SetState(204)
		p.Uid()
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNullLiteral returns the nullLiteral token.
	GetNullLiteral() antlr.Token

	// SetNullLiteral sets the nullLiteral token.
	SetNullLiteral(antlr.Token)

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	nullLiteral antlr.Token
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) GetNullLiteral() antlr.Token { return s.nullLiteral }

func (s *ConstantContext) SetNullLiteral(v antlr.Token) { s.nullLiteral = v }

func (s *ConstantContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ConstantContext) DecimalLiteral() IDecimalLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimalLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *ConstantContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TsSqlParserMINUS, 0)
}

func (s *ConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ConstantContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserNULL_LITERAL, 0)
}

func (s *ConstantContext) NULL_SPEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserNULL_SPEC_LITERAL, 0)
}

func (s *ConstantContext) NOT() antlr.TerminalNode {
	return s.GetToken(TsSqlParserNOT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TsSqlParserRULE_constant)
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

	p.SetState(215)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TsSqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.StringLiteral()
		}

	case TsSqlParserZERO_DECIMAL, TsSqlParserONE_DECIMAL, TsSqlParserTWO_DECIMAL, TsSqlParserDECIMAL_LITERAL, TsSqlParserREAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.DecimalLiteral()
		}

	case TsSqlParserMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(208)
			p.Match(TsSqlParserMINUS)
		}
		{
			p.SetState(209)
			p.DecimalLiteral()
		}

	case TsSqlParserFALSE, TsSqlParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(210)
			p.BooleanLiteral()
		}

	case TsSqlParserNOT, TsSqlParserNULL_LITERAL, TsSqlParserNULL_SPEC_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TsSqlParserNOT {
			{
				p.SetState(211)
				p.Match(TsSqlParserNOT)
			}

		}
		{
			p.SetState(214)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ConstantContext).nullLiteral = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == TsSqlParserNULL_LITERAL || _la == TsSqlParserNULL_SPEC_LITERAL) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ConstantContext).nullLiteral = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) StringLiteral() (localctx IStringLiteralContext) {
	this := p
	_ = this

	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TsSqlParserRULE_stringLiteral)

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
		p.SetState(217)
		p.Match(TsSqlParserSTRING_LITERAL)
	}

	return localctx
}

// IDecimalLiteralContext is an interface to support dynamic dispatch.
type IDecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalLiteralContext differentiates from other interfaces.
	IsDecimalLiteralContext()
}

type DecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalLiteralContext() *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserDECIMAL_LITERAL, 0)
}

func (s *DecimalLiteralContext) ZERO_DECIMAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserZERO_DECIMAL, 0)
}

func (s *DecimalLiteralContext) ONE_DECIMAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserONE_DECIMAL, 0)
}

func (s *DecimalLiteralContext) TWO_DECIMAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserTWO_DECIMAL, 0)
}

func (s *DecimalLiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(TsSqlParserREAL_LITERAL, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	this := p
	_ = this

	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TsSqlParserRULE_decimalLiteral)
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
	{
		p.SetState(219)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-56)) & ^0x3f) == 0 && ((int64(1)<<(_la-56))&24583) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TsSqlParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TsSqlParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	this := p
	_ = this

	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TsSqlParserRULE_booleanLiteral)
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
	{
		p.SetState(221)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TsSqlParserFALSE || _la == TsSqlParserTRUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEmptyStatement_Context is an interface to support dynamic dispatch.
type IEmptyStatement_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyStatement_Context differentiates from other interfaces.
	IsEmptyStatement_Context()
}

type EmptyStatement_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatement_Context() *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TsSqlParserRULE_emptyStatement_
	return p
}

func (*EmptyStatement_Context) IsEmptyStatement_Context() {}

func NewEmptyStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsSqlParserRULE_emptyStatement_

	return p
}

func (s *EmptyStatement_Context) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatement_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(TsSqlParserSEMI, 0)
}

func (s *EmptyStatement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatement_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.EnterEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsSqlParserListener); ok {
		listenerT.ExitEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsSqlParserVisitor:
		return t.VisitEmptyStatement_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsSqlParser) EmptyStatement_() (localctx IEmptyStatement_Context) {
	this := p
	_ = this

	localctx = NewEmptyStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TsSqlParserRULE_emptyStatement_)

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
		p.SetState(223)
		p.Match(TsSqlParserSEMI)
	}

	return localctx
}

func (p *TsSqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 11:
		var t *PredicateContext = nil
		if localctx != nil {
			t = localctx.(*PredicateContext)
		}
		return p.Predicate_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TsSqlParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TsSqlParser) Predicate_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
