package parser

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	parser "github.com/open-unbounded/ts-sql-parser/internal/gen"
	"github.com/stretchr/testify/assert"
)

func createParser(sql string) (*parser.TsSqlParser, *parseTreeVisitor, *VerboseListener) {
	inputStream := antlr.NewInputStream(sql)
	lexer := parser.NewTsSqlLexer(inputStream)
	lexer.RemoveErrorListeners()

	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	tsSqlParser := parser.NewTsSqlParser(tokens)
	tsSqlParser.RemoveErrorListeners()
	listener := &VerboseListener{}
	tsSqlParser.AddErrorListener(listener)
	visitor := &parseTreeVisitor{}
	return tsSqlParser, visitor, listener
}

func Test_parseTreeVisitor_VisitLogicalOperator(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("AND")
		accept := sqlParser.LogicalOperator().Accept(visitor)
		assert.EqualValues(t, "AND", accept)
		fmt.Printf(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("OR")
		accept := sqlParser.LogicalOperator().Accept(visitor)
		assert.EqualValues(t, "OR", accept)
		fmt.Printf(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitConstant(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("1")
		accept := sqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantDecimal{Val: 1}, accept)
		fmt.Printf(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("'1'")
		accept := sqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantString{Val: "'1'"}, accept)
		fmt.Printf(listener.errString.String())
	})

	t.Run("3", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("FALSE")
		accept := sqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantBool{Val: false}, accept)
		fmt.Printf(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitUid(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("a")
		accept := sqlParser.Uid().Accept(visitor)
		assert.EqualValues(t, "a", accept)
		fmt.Printf(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitTableName(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("a.property")
		accept := sqlParser.TableName().Accept(visitor)
		assert.EqualValues(t, TableName{
			TableName:  "a",
			SourceType: "PROPERTY",
		}, accept)
		fmt.Printf(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("A.SERVICE")
		accept := sqlParser.TableName().Accept(visitor)
		assert.EqualValues(t, TableName{
			TableName:  "A",
			SourceType: "SERVICE",
		}, accept)
		fmt.Printf(listener.errString.String())
	})

	t.Run("3", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("A.EVENT")
		accept := sqlParser.TableName().Accept(visitor)
		assert.EqualValues(t, TableName{
			TableName:  "A",
			SourceType: "EVENT",
		}, accept)
		fmt.Printf(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitFromClause(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("FROM a.EVENT WHERE a>1 and 1>=b")
		accept := sqlParser.FromClause().Accept(visitor)
		assert.EqualValues(t, accept, FromClause{
			TableName: TableName{
				TableName:  "a",
				SourceType: "EVENT",
			},
			Alias: "",
			Expression: LogicalExpression{
				LeftExpression: PredicateExpression{ExpressionAtom: BinaryComparisonPredicate{
					Left:  ExpressionAtomPredicate{ExpressionAtom: ColumnNameExpressionAtom{ColumnName: "b"}},
					Op:    ">=",
					Right: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				}},
				op: "and",
				RightExpression: PredicateExpression{ExpressionAtom: BinaryComparisonPredicate{
					Left:  ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
					Op:    ">",
					Right: ExpressionAtomPredicate{ExpressionAtom: ColumnNameExpressionAtom{ColumnName: "a"}},
				}},
			},
		})
		fmt.Printf(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitLimitClause(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("LIMIT 1,2")
		accept := sqlParser.LimitClause().Accept(visitor)
		assert.EqualValues(t, LimitClause{
			offset: 1,
			limit:  2,
		}, accept)
		fmt.Printf(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("LIMIT 2")
		accept := sqlParser.LimitClause().Accept(visitor)
		assert.EqualValues(t, LimitClause{
			limit: 2,
		}, accept)
		fmt.Printf(listener.errString.String())
	})

	t.Run("ERROR", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("LIMIT 1.22")
		accept := sqlParser.LimitClause().Accept(visitor)
		assert.EqualValues(t, LimitClause{offset: 0, limit: 0}, accept)
		fmt.Printf(listener.errString.String())
	})
}
