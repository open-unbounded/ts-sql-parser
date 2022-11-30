// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // SqlParser

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

type SqlParser struct {
	*antlr.BaseParser
}

var sqlparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sqlparserParserInit() {
	staticData := &sqlparserParserStaticData
	staticData.literalNames = []string{
		"", "", "", "", "", "'.PROPERTY'", "'.SERVICE'", "'.EVENT'", "'FALSE'",
		"'TRUE'", "'SELECT'", "'FROM'", "'WHERE'", "'INTERVAL'", "'AND'", "'IN'",
		"'IS'", "'OR'", "'AS'", "'XOR'", "'JOIN'", "'LEFT'", "'LIKE'", "'LIMIT'",
		"'NOT'", "'*'", "'/'", "'%'", "'+'", "'-'", "'DIV'", "'MOD'", "'='",
		"'>'", "'<'", "'!'", "'~'", "'|'", "'&'", "'^'", "", "'.'", "'('", "')'",
		"','", "';'", "'@'", "'0'", "'1'", "'2'", "'''", "'\"'", "'`'", "':'",
		"", "'NULL'", "", "", "'PROPERTY'", "'SERVICE'",
	}
	staticData.symbolicNames = []string{
		"", "SPACE", "SPEC_MYSQL_COMMENT", "COMMENT_INPUT", "LINE_COMMENT",
		"DOT_PROPERTY", "DOT_SERVICE", "DOT_EVENT", "FALSE", "TRUE", "SELECT",
		"FROM", "WHERE", "INTERVAL", "AND", "IN", "IS", "OR", "AS", "XOR", "JOIN",
		"LEFT", "LIKE", "LIMIT", "NOT", "STAR", "DIVIDE", "MODULE", "PLUS",
		"MINUS", "DIV", "MOD", "EQUAL_SYMBOL", "GREATER_SYMBOL", "LESS_SYMBOL",
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
		"expression", "predicate", "expressionAtom", "columnName", "logicalOperator",
		"comparisonOperator", "fullColumnName", "constant", "stringLiteral",
		"decimalLiteral", "booleanLiteral", "emptyStatement_",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 62, 195, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 1, 0, 3, 0, 47, 8, 0, 1, 0, 5, 0, 50, 8, 0, 10, 0, 12,
		0, 53, 9, 0, 1, 0, 1, 0, 3, 0, 57, 8, 0, 1, 0, 3, 0, 60, 8, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 66, 8, 1, 1, 1, 3, 1, 69, 8, 1, 1, 2, 1, 2, 3, 2,
		73, 8, 2, 1, 2, 1, 2, 5, 2, 77, 8, 2, 10, 2, 12, 2, 80, 9, 2, 1, 3, 1,
		3, 3, 3, 84, 8, 3, 1, 3, 3, 3, 87, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		93, 8, 4, 1, 4, 1, 4, 3, 4, 97, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 103,
		8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		126, 8, 10, 10, 10, 12, 10, 129, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 5, 11, 138, 8, 11, 10, 11, 12, 11, 141, 9, 11, 1, 12,
		1, 12, 3, 12, 145, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 3, 14, 156, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 172,
		8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 182,
		8, 17, 1, 17, 3, 17, 185, 8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 0, 2, 20, 22, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 0, 5, 2, 0, 47, 49,
		60, 60, 1, 0, 5, 7, 1, 0, 55, 56, 2, 0, 47, 49, 60, 61, 1, 0, 8, 9, 205,
		0, 51, 1, 0, 0, 0, 2, 61, 1, 0, 0, 0, 4, 72, 1, 0, 0, 0, 6, 81, 1, 0, 0,
		0, 8, 88, 1, 0, 0, 0, 10, 98, 1, 0, 0, 0, 12, 106, 1, 0, 0, 0, 14, 108,
		1, 0, 0, 0, 16, 113, 1, 0, 0, 0, 18, 115, 1, 0, 0, 0, 20, 118, 1, 0, 0,
		0, 22, 130, 1, 0, 0, 0, 24, 144, 1, 0, 0, 0, 26, 146, 1, 0, 0, 0, 28, 155,
		1, 0, 0, 0, 30, 171, 1, 0, 0, 0, 32, 173, 1, 0, 0, 0, 34, 184, 1, 0, 0,
		0, 36, 186, 1, 0, 0, 0, 38, 188, 1, 0, 0, 0, 40, 190, 1, 0, 0, 0, 42, 192,
		1, 0, 0, 0, 44, 46, 3, 2, 1, 0, 45, 47, 5, 45, 0, 0, 46, 45, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 50, 3, 42, 21, 0, 49, 44, 1,
		0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51,
		52, 1, 0, 0, 0, 52, 59, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 56, 3, 2, 1,
		0, 55, 57, 5, 45, 0, 0, 56, 55, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 60,
		1, 0, 0, 0, 58, 60, 3, 42, 21, 0, 59, 54, 1, 0, 0, 0, 59, 58, 1, 0, 0,
		0, 60, 1, 1, 0, 0, 0, 61, 62, 5, 10, 0, 0, 62, 63, 3, 4, 2, 0, 63, 65,
		3, 8, 4, 0, 64, 66, 3, 10, 5, 0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0,
		66, 68, 1, 0, 0, 0, 67, 69, 3, 14, 7, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1,
		0, 0, 0, 69, 3, 1, 0, 0, 0, 70, 73, 5, 25, 0, 0, 71, 73, 3, 6, 3, 0, 72,
		70, 1, 0, 0, 0, 72, 71, 1, 0, 0, 0, 73, 78, 1, 0, 0, 0, 74, 75, 5, 44,
		0, 0, 75, 77, 3, 6, 3, 0, 76, 74, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76,
		1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 5, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0,
		81, 86, 3, 32, 16, 0, 82, 84, 5, 18, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 3, 16, 8, 0, 86, 83, 1, 0, 0, 0,
		86, 87, 1, 0, 0, 0, 87, 7, 1, 0, 0, 0, 88, 89, 5, 11, 0, 0, 89, 92, 3,
		18, 9, 0, 90, 91, 5, 18, 0, 0, 91, 93, 3, 16, 8, 0, 92, 90, 1, 0, 0, 0,
		92, 93, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 95, 5, 12, 0, 0, 95, 97, 3,
		20, 10, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 9, 1, 0, 0, 0, 98,
		102, 5, 23, 0, 0, 99, 100, 3, 12, 6, 0, 100, 101, 5, 44, 0, 0, 101, 103,
		1, 0, 0, 0, 102, 99, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 1, 0,
		0, 0, 104, 105, 3, 12, 6, 0, 105, 11, 1, 0, 0, 0, 106, 107, 7, 0, 0, 0,
		107, 13, 1, 0, 0, 0, 108, 109, 5, 13, 0, 0, 109, 110, 5, 42, 0, 0, 110,
		111, 5, 40, 0, 0, 111, 112, 5, 43, 0, 0, 112, 15, 1, 0, 0, 0, 113, 114,
		5, 54, 0, 0, 114, 17, 1, 0, 0, 0, 115, 116, 3, 16, 8, 0, 116, 117, 7, 1,
		0, 0, 117, 19, 1, 0, 0, 0, 118, 119, 6, 10, -1, 0, 119, 120, 3, 22, 11,
		0, 120, 127, 1, 0, 0, 0, 121, 122, 10, 2, 0, 0, 122, 123, 3, 28, 14, 0,
		123, 124, 3, 20, 10, 3, 124, 126, 1, 0, 0, 0, 125, 121, 1, 0, 0, 0, 126,
		129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 21, 1,
		0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 131, 6, 11, -1, 0, 131, 132, 3, 24,
		12, 0, 132, 139, 1, 0, 0, 0, 133, 134, 10, 2, 0, 0, 134, 135, 3, 30, 15,
		0, 135, 136, 3, 22, 11, 3, 136, 138, 1, 0, 0, 0, 137, 133, 1, 0, 0, 0,
		138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140,
		23, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 145, 3, 34, 17, 0, 143, 145,
		3, 26, 13, 0, 144, 142, 1, 0, 0, 0, 144, 143, 1, 0, 0, 0, 145, 25, 1, 0,
		0, 0, 146, 147, 3, 16, 8, 0, 147, 27, 1, 0, 0, 0, 148, 156, 5, 14, 0, 0,
		149, 150, 5, 38, 0, 0, 150, 156, 5, 38, 0, 0, 151, 156, 5, 19, 0, 0, 152,
		156, 5, 17, 0, 0, 153, 154, 5, 37, 0, 0, 154, 156, 5, 37, 0, 0, 155, 148,
		1, 0, 0, 0, 155, 149, 1, 0, 0, 0, 155, 151, 1, 0, 0, 0, 155, 152, 1, 0,
		0, 0, 155, 153, 1, 0, 0, 0, 156, 29, 1, 0, 0, 0, 157, 172, 5, 32, 0, 0,
		158, 172, 5, 33, 0, 0, 159, 172, 5, 34, 0, 0, 160, 161, 5, 34, 0, 0, 161,
		172, 5, 32, 0, 0, 162, 163, 5, 33, 0, 0, 163, 172, 5, 32, 0, 0, 164, 165,
		5, 34, 0, 0, 165, 172, 5, 33, 0, 0, 166, 167, 5, 35, 0, 0, 167, 172, 5,
		32, 0, 0, 168, 169, 5, 34, 0, 0, 169, 170, 5, 32, 0, 0, 170, 172, 5, 33,
		0, 0, 171, 157, 1, 0, 0, 0, 171, 158, 1, 0, 0, 0, 171, 159, 1, 0, 0, 0,
		171, 160, 1, 0, 0, 0, 171, 162, 1, 0, 0, 0, 171, 164, 1, 0, 0, 0, 171,
		166, 1, 0, 0, 0, 171, 168, 1, 0, 0, 0, 172, 31, 1, 0, 0, 0, 173, 174, 3,
		16, 8, 0, 174, 33, 1, 0, 0, 0, 175, 185, 3, 36, 18, 0, 176, 185, 3, 38,
		19, 0, 177, 178, 5, 29, 0, 0, 178, 185, 3, 38, 19, 0, 179, 185, 3, 40,
		20, 0, 180, 182, 5, 24, 0, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0,
		0, 182, 183, 1, 0, 0, 0, 183, 185, 7, 2, 0, 0, 184, 175, 1, 0, 0, 0, 184,
		176, 1, 0, 0, 0, 184, 177, 1, 0, 0, 0, 184, 179, 1, 0, 0, 0, 184, 181,
		1, 0, 0, 0, 185, 35, 1, 0, 0, 0, 186, 187, 5, 57, 0, 0, 187, 37, 1, 0,
		0, 0, 188, 189, 7, 3, 0, 0, 189, 39, 1, 0, 0, 0, 190, 191, 7, 4, 0, 0,
		191, 41, 1, 0, 0, 0, 192, 193, 5, 45, 0, 0, 193, 43, 1, 0, 0, 0, 21, 46,
		49, 51, 56, 59, 65, 68, 72, 78, 83, 86, 92, 96, 102, 127, 139, 144, 155,
		171, 181, 184,
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

// SqlParserInit initializes any static state used to implement SqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SqlParserInit() {
	staticData := &sqlparserParserStaticData
	staticData.once.Do(sqlparserParserInit)
}

// NewSqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSqlParser(input antlr.TokenStream) *SqlParser {
	SqlParserInit()
	this := new(SqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sqlparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// SqlParser tokens.
const (
	SqlParserEOF                = antlr.TokenEOF
	SqlParserSPACE              = 1
	SqlParserSPEC_MYSQL_COMMENT = 2
	SqlParserCOMMENT_INPUT      = 3
	SqlParserLINE_COMMENT       = 4
	SqlParserDOT_PROPERTY       = 5
	SqlParserDOT_SERVICE        = 6
	SqlParserDOT_EVENT          = 7
	SqlParserFALSE              = 8
	SqlParserTRUE               = 9
	SqlParserSELECT             = 10
	SqlParserFROM               = 11
	SqlParserWHERE              = 12
	SqlParserINTERVAL           = 13
	SqlParserAND                = 14
	SqlParserIN                 = 15
	SqlParserIS                 = 16
	SqlParserOR                 = 17
	SqlParserAS                 = 18
	SqlParserXOR                = 19
	SqlParserJOIN               = 20
	SqlParserLEFT               = 21
	SqlParserLIKE               = 22
	SqlParserLIMIT              = 23
	SqlParserNOT                = 24
	SqlParserSTAR               = 25
	SqlParserDIVIDE             = 26
	SqlParserMODULE             = 27
	SqlParserPLUS               = 28
	SqlParserMINUS              = 29
	SqlParserDIV                = 30
	SqlParserMOD                = 31
	SqlParserEQUAL_SYMBOL       = 32
	SqlParserGREATER_SYMBOL     = 33
	SqlParserLESS_SYMBOL        = 34
	SqlParserEXCLAMATION_SYMBOL = 35
	SqlParserBIT_NOT_OP         = 36
	SqlParserBIT_OR_OP          = 37
	SqlParserBIT_AND_OP         = 38
	SqlParserBIT_XOR_OP         = 39
	SqlParserTIME_INTERVAL      = 40
	SqlParserDOT                = 41
	SqlParserLR_BRACKET         = 42
	SqlParserRR_BRACKET         = 43
	SqlParserCOMMA              = 44
	SqlParserSEMI               = 45
	SqlParserAT_SIGN            = 46
	SqlParserZERO_DECIMAL       = 47
	SqlParserONE_DECIMAL        = 48
	SqlParserTWO_DECIMAL        = 49
	SqlParserSINGLE_QUOTE_SYMB  = 50
	SqlParserDOUBLE_QUOTE_SYMB  = 51
	SqlParserREVERSE_QUOTE_SYMB = 52
	SqlParserCOLON_SYMB         = 53
	SqlParserID                 = 54
	SqlParserNULL_LITERAL       = 55
	SqlParserNULL_SPEC_LITERAL  = 56
	SqlParserSTRING_LITERAL     = 57
	SqlParserPROPERTY           = 58
	SqlParserSERVICE            = 59
	SqlParserDECIMAL_LITERAL    = 60
	SqlParserREAL_LITERAL       = 61
	SqlParserERROR_RECONGNIGION = 62
)

// SqlParser rules.
const (
	SqlParserRULE_root               = 0
	SqlParserRULE_selectStmt         = 1
	SqlParserRULE_selectElements     = 2
	SqlParserRULE_selectElement      = 3
	SqlParserRULE_fromClause         = 4
	SqlParserRULE_limitClause        = 5
	SqlParserRULE_limitClauseAtom    = 6
	SqlParserRULE_windowClause       = 7
	SqlParserRULE_uid                = 8
	SqlParserRULE_tableName          = 9
	SqlParserRULE_expression         = 10
	SqlParserRULE_predicate          = 11
	SqlParserRULE_expressionAtom     = 12
	SqlParserRULE_columnName         = 13
	SqlParserRULE_logicalOperator    = 14
	SqlParserRULE_comparisonOperator = 15
	SqlParserRULE_fullColumnName     = 16
	SqlParserRULE_constant           = 17
	SqlParserRULE_stringLiteral      = 18
	SqlParserRULE_decimalLiteral     = 19
	SqlParserRULE_booleanLiteral     = 20
	SqlParserRULE_emptyStatement_    = 21
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
	p.RuleIndex = SqlParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_root

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
	return s.GetTokens(SqlParserSEMI)
}

func (s *RootContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserSEMI, i)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Root() (localctx IRootContext) {
	this := p
	_ = this

	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SqlParserRULE_root)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(49)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case SqlParserSELECT:
				{
					p.SetState(44)
					p.SelectStmt()
				}
				p.SetState(46)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(45)
						p.Match(SqlParserSEMI)
					}

				}

			case SqlParserSEMI:
				{
					p.SetState(48)
					p.EmptyStatement_()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserSELECT:
		{
			p.SetState(54)
			p.SelectStmt()
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserSEMI {
			{
				p.SetState(55)
				p.Match(SqlParserSEMI)
			}

		}

	case SqlParserSEMI:
		{
			p.SetState(58)
			p.EmptyStatement_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = SqlParserRULE_selectStmt
	return p
}

func (*SelectStmtContext) IsSelectStmtContext() {}

func NewSelectStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStmtContext {
	var p = new(SelectStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectStmt

	return p
}

func (s *SelectStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStmtContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SqlParserSELECT, 0)
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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSelectStmt(s)
	}
}

func (s *SelectStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSelectStmt(s)
	}
}

func (s *SelectStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSelectStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SelectStmt() (localctx ISelectStmtContext) {
	this := p
	_ = this

	localctx = NewSelectStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SqlParserRULE_selectStmt)
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
		p.SetState(61)
		p.Match(SqlParserSELECT)
	}
	{
		p.SetState(62)
		p.SelectElements()
	}
	{
		p.SetState(63)
		p.FromClause()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserLIMIT {
		{
			p.SetState(64)
			p.LimitClause()
		}

	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserINTERVAL {
		{
			p.SetState(67)
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
	p.RuleIndex = SqlParserRULE_selectElements
	return p
}

func (*SelectElementsContext) IsSelectElementsContext() {}

func NewSelectElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementsContext {
	var p = new(SelectElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectElements

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
	return s.GetToken(SqlParserSTAR, 0)
}

func (s *SelectElementsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *SelectElementsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *SelectElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSelectElements(s)
	}
}

func (s *SelectElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSelectElements(s)
	}
}

func (s *SelectElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSelectElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SelectElements() (localctx ISelectElementsContext) {
	this := p
	_ = this

	localctx = NewSelectElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SqlParserRULE_selectElements)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserSTAR:
		{
			p.SetState(70)

			var _m = p.Match(SqlParserSTAR)

			localctx.(*SelectElementsContext).star = _m
		}

	case SqlParserID:
		{
			p.SetState(71)
			p.SelectElement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(74)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(75)
			p.SelectElement()
		}

		p.SetState(80)
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
	p.RuleIndex = SqlParserRULE_selectElement
	return p
}

func (*SelectElementContext) IsSelectElementContext() {}

func NewSelectElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementContext {
	var p = new(SelectElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectElement

	return p
}

func (s *SelectElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementContext) FullColumnName() IFullColumnNameContext {
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

func (s *SelectElementContext) Uid() IUidContext {
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

func (s *SelectElementContext) AS() antlr.TerminalNode {
	return s.GetToken(SqlParserAS, 0)
}

func (s *SelectElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSelectElement(s)
	}
}

func (s *SelectElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSelectElement(s)
	}
}

func (s *SelectElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSelectElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SelectElement() (localctx ISelectElementContext) {
	this := p
	_ = this

	localctx = NewSelectElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SqlParserRULE_selectElement)
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
		p.SetState(81)
		p.FullColumnName()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserAS || _la == SqlParserID {
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserAS {
			{
				p.SetState(82)
				p.Match(SqlParserAS)
			}

		}
		{
			p.SetState(85)
			p.Uid()
		}

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
	p.RuleIndex = SqlParserRULE_fromClause
	return p
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(SqlParserFROM, 0)
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

func (s *FromClauseContext) AS() antlr.TerminalNode {
	return s.GetToken(SqlParserAS, 0)
}

func (s *FromClauseContext) Uid() IUidContext {
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

func (s *FromClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SqlParserWHERE, 0)
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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitFromClause(s)
	}
}

func (s *FromClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitFromClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) FromClause() (localctx IFromClauseContext) {
	this := p
	_ = this

	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SqlParserRULE_fromClause)
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
		p.SetState(88)
		p.Match(SqlParserFROM)
	}
	{
		p.SetState(89)
		p.TableName()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserAS {
		{
			p.SetState(90)
			p.Match(SqlParserAS)
		}
		{
			p.SetState(91)
			p.Uid()
		}

	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserWHERE {
		{
			p.SetState(94)
			p.Match(SqlParserWHERE)
		}
		{
			p.SetState(95)
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
	p.RuleIndex = SqlParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) GetOffset() ILimitClauseAtomContext { return s.offset }

func (s *LimitClauseContext) GetLimit() ILimitClauseAtomContext { return s.limit }

func (s *LimitClauseContext) SetOffset(v ILimitClauseAtomContext) { s.offset = v }

func (s *LimitClauseContext) SetLimit(v ILimitClauseAtomContext) { s.limit = v }

func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SqlParserLIMIT, 0)
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
	return s.GetToken(SqlParserCOMMA, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) LimitClause() (localctx ILimitClauseContext) {
	this := p
	_ = this

	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SqlParserRULE_limitClause)

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
		p.SetState(98)
		p.Match(SqlParserLIMIT)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(99)

			var _x = p.LimitClauseAtom()

			localctx.(*LimitClauseContext).offset = _x
		}
		{
			p.SetState(100)
			p.Match(SqlParserCOMMA)
		}

	}
	{
		p.SetState(104)

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
	p.RuleIndex = SqlParserRULE_limitClauseAtom
	return p
}

func (*LimitClauseAtomContext) IsLimitClauseAtomContext() {}

func NewLimitClauseAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseAtomContext {
	var p = new(LimitClauseAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_limitClauseAtom

	return p
}

func (s *LimitClauseAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseAtomContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, 0)
}

func (s *LimitClauseAtomContext) TWO_DECIMAL() antlr.TerminalNode {
	return s.GetToken(SqlParserTWO_DECIMAL, 0)
}

func (s *LimitClauseAtomContext) ONE_DECIMAL() antlr.TerminalNode {
	return s.GetToken(SqlParserONE_DECIMAL, 0)
}

func (s *LimitClauseAtomContext) ZERO_DECIMAL() antlr.TerminalNode {
	return s.GetToken(SqlParserZERO_DECIMAL, 0)
}

func (s *LimitClauseAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLimitClauseAtom(s)
	}
}

func (s *LimitClauseAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLimitClauseAtom(s)
	}
}

func (s *LimitClauseAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitLimitClauseAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) LimitClauseAtom() (localctx ILimitClauseAtomContext) {
	this := p
	_ = this

	localctx = NewLimitClauseAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SqlParserRULE_limitClauseAtom)
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
		p.SetState(106)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1153906667025334272) != 0) {
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
	p.RuleIndex = SqlParserRULE_windowClause
	return p
}

func (*WindowClauseContext) IsWindowClauseContext() {}

func NewWindowClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowClauseContext {
	var p = new(WindowClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_windowClause

	return p
}

func (s *WindowClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowClauseContext) INTERVAL() antlr.TerminalNode {
	return s.GetToken(SqlParserINTERVAL, 0)
}

func (s *WindowClauseContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(SqlParserLR_BRACKET, 0)
}

func (s *WindowClauseContext) TIME_INTERVAL() antlr.TerminalNode {
	return s.GetToken(SqlParserTIME_INTERVAL, 0)
}

func (s *WindowClauseContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(SqlParserRR_BRACKET, 0)
}

func (s *WindowClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterWindowClause(s)
	}
}

func (s *WindowClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitWindowClause(s)
	}
}

func (s *WindowClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitWindowClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) WindowClause() (localctx IWindowClauseContext) {
	this := p
	_ = this

	localctx = NewWindowClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SqlParserRULE_windowClause)

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
		p.SetState(108)
		p.Match(SqlParserINTERVAL)
	}
	{
		p.SetState(109)
		p.Match(SqlParserLR_BRACKET)
	}
	{
		p.SetState(110)
		p.Match(SqlParserTIME_INTERVAL)
	}
	{
		p.SetState(111)
		p.Match(SqlParserRR_BRACKET)
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
	p.RuleIndex = SqlParserRULE_uid
	return p
}

func (*UidContext) IsUidContext() {}

func NewUidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UidContext {
	var p = new(UidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_uid

	return p
}

func (s *UidContext) GetParser() antlr.Parser { return s.parser }

func (s *UidContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *UidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterUid(s)
	}
}

func (s *UidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitUid(s)
	}
}

func (s *UidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitUid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Uid() (localctx IUidContext) {
	this := p
	_ = this

	localctx = NewUidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SqlParserRULE_uid)

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
		p.SetState(113)
		p.Match(SqlParserID)
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
	p.RuleIndex = SqlParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableName

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
	return s.GetToken(SqlParserDOT_PROPERTY, 0)
}

func (s *TableNameContext) DOT_SERVICE() antlr.TerminalNode {
	return s.GetToken(SqlParserDOT_SERVICE, 0)
}

func (s *TableNameContext) DOT_EVENT() antlr.TerminalNode {
	return s.GetToken(SqlParserDOT_EVENT, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (s *TableNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitTableName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TableName() (localctx ITableNameContext) {
	this := p
	_ = this

	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SqlParserRULE_tableName)
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
		p.SetState(115)
		p.Uid()
	}
	{
		p.SetState(116)
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
	p.RuleIndex = SqlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_expression

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
	right IExpressionContext
	left  IExpressionContext
}

func NewLogicalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *LogicalExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *LogicalExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *LogicalExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterPredicateExpression(s)
	}
}

func (s *PredicateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitPredicateExpression(s)
	}
}

func (s *PredicateExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitPredicateExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *SqlParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, SqlParserRULE_expression, _p)

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
		p.SetState(119)
		p.predicate(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*LogicalExpressionContext).right = _prevctx

			p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expression)
			p.SetState(121)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(122)
				p.LogicalOperator()
			}
			{
				p.SetState(123)

				var _x = p.expression(3)

				localctx.(*LogicalExpressionContext).left = _x
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.RuleIndex = SqlParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_predicate

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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterExpressionAtomPredicate(s)
	}
}

func (s *ExpressionAtomPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitExpressionAtomPredicate(s)
	}
}

func (s *ExpressionAtomPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitExpressionAtomPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryComparisonPredicateContext struct {
	*PredicateContext
	right IPredicateContext
	left  IPredicateContext
}

func NewBinaryComparisonPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryComparisonPredicateContext {
	var p = new(BinaryComparisonPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *BinaryComparisonPredicateContext) GetRight() IPredicateContext { return s.right }

func (s *BinaryComparisonPredicateContext) GetLeft() IPredicateContext { return s.left }

func (s *BinaryComparisonPredicateContext) SetRight(v IPredicateContext) { s.right = v }

func (s *BinaryComparisonPredicateContext) SetLeft(v IPredicateContext) { s.left = v }

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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterBinaryComparisonPredicate(s)
	}
}

func (s *BinaryComparisonPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitBinaryComparisonPredicate(s)
	}
}

func (s *BinaryComparisonPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitBinaryComparisonPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Predicate() (localctx IPredicateContext) {
	return p.predicate(0)
}

func (p *SqlParser) predicate(_p int) (localctx IPredicateContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPredicateContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, SqlParserRULE_predicate, _p)

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
		p.SetState(131)
		p.ExpressionAtom()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBinaryComparisonPredicateContext(p, NewPredicateContext(p, _parentctx, _parentState))
			localctx.(*BinaryComparisonPredicateContext).right = _prevctx

			p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_predicate)
			p.SetState(133)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(134)
				p.ComparisonOperator()
			}
			{
				p.SetState(135)

				var _x = p.predicate(3)

				localctx.(*BinaryComparisonPredicateContext).left = _x
			}

		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.RuleIndex = SqlParserRULE_expressionAtom
	return p
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_expressionAtom

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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterColumnNameExpressionAtom(s)
	}
}

func (s *ColumnNameExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitColumnNameExpressionAtom(s)
	}
}

func (s *ColumnNameExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterConstantExpressionAtom(s)
	}
}

func (s *ConstantExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitConstantExpressionAtom(s)
	}
}

func (s *ConstantExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitConstantExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ExpressionAtom() (localctx IExpressionAtomContext) {
	this := p
	_ = this

	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SqlParserRULE_expressionAtom)

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

	p.SetState(144)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserFALSE, SqlParserTRUE, SqlParserNOT, SqlParserMINUS, SqlParserZERO_DECIMAL, SqlParserONE_DECIMAL, SqlParserTWO_DECIMAL, SqlParserNULL_LITERAL, SqlParserNULL_SPEC_LITERAL, SqlParserSTRING_LITERAL, SqlParserDECIMAL_LITERAL, SqlParserREAL_LITERAL:
		localctx = NewConstantExpressionAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Constant()
		}

	case SqlParserID:
		localctx = NewColumnNameExpressionAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.ColumnName()
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
	p.RuleIndex = SqlParserRULE_columnName
	return p
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnName

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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (s *ColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ColumnName() (localctx IColumnNameContext) {
	this := p
	_ = this

	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SqlParserRULE_columnName)

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
		p.SetState(146)
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
	p.RuleIndex = SqlParserRULE_logicalOperator
	return p
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(SqlParserAND, 0)
}

func (s *LogicalOperatorContext) AllBIT_AND_OP() []antlr.TerminalNode {
	return s.GetTokens(SqlParserBIT_AND_OP)
}

func (s *LogicalOperatorContext) BIT_AND_OP(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserBIT_AND_OP, i)
}

func (s *LogicalOperatorContext) XOR() antlr.TerminalNode {
	return s.GetToken(SqlParserXOR, 0)
}

func (s *LogicalOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(SqlParserOR, 0)
}

func (s *LogicalOperatorContext) AllBIT_OR_OP() []antlr.TerminalNode {
	return s.GetTokens(SqlParserBIT_OR_OP)
}

func (s *LogicalOperatorContext) BIT_OR_OP(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserBIT_OR_OP, i)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitLogicalOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	this := p
	_ = this

	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SqlParserRULE_logicalOperator)

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
	case SqlParserAND:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Match(SqlParserAND)
		}

	case SqlParserBIT_AND_OP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Match(SqlParserBIT_AND_OP)
		}
		{
			p.SetState(150)
			p.Match(SqlParserBIT_AND_OP)
		}

	case SqlParserXOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.Match(SqlParserXOR)
		}

	case SqlParserOR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.Match(SqlParserOR)
		}

	case SqlParserBIT_OR_OP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(153)
			p.Match(SqlParserBIT_OR_OP)
		}
		{
			p.SetState(154)
			p.Match(SqlParserBIT_OR_OP)
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
	p.RuleIndex = SqlParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) EQUAL_SYMBOL() antlr.TerminalNode {
	return s.GetToken(SqlParserEQUAL_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) GREATER_SYMBOL() antlr.TerminalNode {
	return s.GetToken(SqlParserGREATER_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) LESS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(SqlParserLESS_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) EXCLAMATION_SYMBOL() antlr.TerminalNode {
	return s.GetToken(SqlParserEXCLAMATION_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	this := p
	_ = this

	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SqlParserRULE_comparisonOperator)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Match(SqlParserEQUAL_SYMBOL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Match(SqlParserGREATER_SYMBOL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.Match(SqlParserLESS_SYMBOL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(160)
			p.Match(SqlParserLESS_SYMBOL)
		}
		{
			p.SetState(161)
			p.Match(SqlParserEQUAL_SYMBOL)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(162)
			p.Match(SqlParserGREATER_SYMBOL)
		}
		{
			p.SetState(163)
			p.Match(SqlParserEQUAL_SYMBOL)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(164)
			p.Match(SqlParserLESS_SYMBOL)
		}
		{
			p.SetState(165)
			p.Match(SqlParserGREATER_SYMBOL)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(166)
			p.Match(SqlParserEXCLAMATION_SYMBOL)
		}
		{
			p.SetState(167)
			p.Match(SqlParserEQUAL_SYMBOL)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(168)
			p.Match(SqlParserLESS_SYMBOL)
		}
		{
			p.SetState(169)
			p.Match(SqlParserEQUAL_SYMBOL)
		}
		{
			p.SetState(170)
			p.Match(SqlParserGREATER_SYMBOL)
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
	p.RuleIndex = SqlParserRULE_fullColumnName
	return p
}

func (*FullColumnNameContext) IsFullColumnNameContext() {}

func NewFullColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullColumnNameContext {
	var p = new(FullColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_fullColumnName

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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterFullColumnName(s)
	}
}

func (s *FullColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitFullColumnName(s)
	}
}

func (s *FullColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitFullColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) FullColumnName() (localctx IFullColumnNameContext) {
	this := p
	_ = this

	localctx = NewFullColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SqlParserRULE_fullColumnName)

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
		p.SetState(173)
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
	p.RuleIndex = SqlParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_constant

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
	return s.GetToken(SqlParserMINUS, 0)
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
	return s.GetToken(SqlParserNULL_LITERAL, 0)
}

func (s *ConstantContext) NULL_SPEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserNULL_SPEC_LITERAL, 0)
}

func (s *ConstantContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SqlParserRULE_constant)
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

	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.StringLiteral()
		}

	case SqlParserZERO_DECIMAL, SqlParserONE_DECIMAL, SqlParserTWO_DECIMAL, SqlParserDECIMAL_LITERAL, SqlParserREAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.DecimalLiteral()
		}

	case SqlParserMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)
			p.Match(SqlParserMINUS)
		}
		{
			p.SetState(178)
			p.DecimalLiteral()
		}

	case SqlParserFALSE, SqlParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(179)
			p.BooleanLiteral()
		}

	case SqlParserNOT, SqlParserNULL_LITERAL, SqlParserNULL_SPEC_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(180)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(183)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ConstantContext).nullLiteral = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserNULL_LITERAL || _la == SqlParserNULL_SPEC_LITERAL) {
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
	p.RuleIndex = SqlParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) StringLiteral() (localctx IStringLiteralContext) {
	this := p
	_ = this

	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SqlParserRULE_stringLiteral)

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
		p.SetState(186)
		p.Match(SqlParserSTRING_LITERAL)
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
	p.RuleIndex = SqlParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, 0)
}

func (s *DecimalLiteralContext) ZERO_DECIMAL() antlr.TerminalNode {
	return s.GetToken(SqlParserZERO_DECIMAL, 0)
}

func (s *DecimalLiteralContext) ONE_DECIMAL() antlr.TerminalNode {
	return s.GetToken(SqlParserONE_DECIMAL, 0)
}

func (s *DecimalLiteralContext) TWO_DECIMAL() antlr.TerminalNode {
	return s.GetToken(SqlParserTWO_DECIMAL, 0)
}

func (s *DecimalLiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserREAL_LITERAL, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	this := p
	_ = this

	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SqlParserRULE_decimalLiteral)
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
		p.SetState(188)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3459749676239028224) != 0) {
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
	p.RuleIndex = SqlParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SqlParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SqlParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	this := p
	_ = this

	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SqlParserRULE_booleanLiteral)
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
		p.SetState(190)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SqlParserFALSE || _la == SqlParserTRUE) {
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
	p.RuleIndex = SqlParserRULE_emptyStatement_
	return p
}

func (*EmptyStatement_Context) IsEmptyStatement_Context() {}

func NewEmptyStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_emptyStatement_

	return p
}

func (s *EmptyStatement_Context) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatement_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(SqlParserSEMI, 0)
}

func (s *EmptyStatement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatement_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitEmptyStatement_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) EmptyStatement_() (localctx IEmptyStatement_Context) {
	this := p
	_ = this

	localctx = NewEmptyStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SqlParserRULE_emptyStatement_)

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
		p.SetState(192)
		p.Match(SqlParserSEMI)
	}

	return localctx
}

func (p *SqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *SqlParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) Predicate_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
