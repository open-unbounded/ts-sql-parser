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
		fmt.Print(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("OR")
		accept := sqlParser.LogicalOperator().Accept(visitor)
		assert.EqualValues(t, "OR", accept)
		fmt.Print(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitConstant(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("1")
		accept := sqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantDecimal{Val: 1}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("-1")
		accept := sqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantDecimal{Val: -1}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("3", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("'1'")
		accept := sqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantString{Val: "'1'"}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("4", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("FALSE")
		accept := sqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantBool{Val: false}, accept)
		fmt.Print(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitUid(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("a")
		accept := sqlParser.Uid().Accept(visitor)
		assert.EqualValues(t, "a", accept)
		fmt.Print(listener.errString.String())
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
		fmt.Print(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("A.SERVICE")
		accept := sqlParser.TableName().Accept(visitor)
		assert.EqualValues(t, TableName{
			TableName:  "A",
			SourceType: "SERVICE",
		}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("3", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("A.EVENT")
		accept := sqlParser.TableName().Accept(visitor)
		assert.EqualValues(t, TableName{
			TableName:  "A",
			SourceType: "EVENT",
		}, accept)
		fmt.Print(listener.errString.String())
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
			Expression: LogicalExpression{
				LeftExpression: PredicateExpression{ExpressionAtom: BinaryComparisonPredicate{
					Left:  ExpressionAtomPredicate{ExpressionAtom: ColumnNameExpressionAtom{ColumnName: "b"}},
					Op:    ">=",
					Right: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				}},
				Op: "and",
				RightExpression: PredicateExpression{ExpressionAtom: BinaryComparisonPredicate{
					Left:  ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
					Op:    ">",
					Right: ExpressionAtomPredicate{ExpressionAtom: ColumnNameExpressionAtom{ColumnName: "a"}},
				}},
			},
		})
		fmt.Print(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitLimitClause(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("LIMIT 1,2")
		accept := sqlParser.LimitClause().Accept(visitor)
		assert.EqualValues(t, LimitClause{
			Offset: 1,
			Limit:  2,
		}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("LIMIT 2")
		accept := sqlParser.LimitClause().Accept(visitor)
		assert.EqualValues(t, LimitClause{
			Limit: 2,
		}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("ERROR", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("LIMIT 1.22")
		accept := sqlParser.LimitClause().Accept(visitor)
		assert.EqualValues(t, LimitClause{Offset: 0, Limit: 0}, accept)
		fmt.Print(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitWindowClause(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		for _, s := range []string{"1s", "1h", "1d", "3d", "30h", "300s"} {
			sqlParser, visitor, listener := createParser(fmt.Sprintf("INTERVAL(%s)", s))
			accept := sqlParser.WindowClause().Accept(visitor)
			assert.EqualValues(t, WindowClause{
				Duration: s,
			}, accept)
			fmt.Print(listener.errString.String())
		}
	})
}

func Test_parseTreeVisitor_VisitSelectElement(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("a b")
		accept := sqlParser.SelectElement().Accept(visitor)
		assert.EqualValues(t, SelectElement{
			FullColumnName: "a",
			Alias:          "b",
		}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("a")
		accept := sqlParser.SelectElement().Accept(visitor)
		assert.EqualValues(t, SelectElement{
			FullColumnName: "a",
		}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("3", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("a AS B")
		accept := sqlParser.SelectElement().Accept(visitor)
		assert.EqualValues(t, SelectElement{
			FullColumnName: "a",
			Alias:          "B",
		}, accept)
		fmt.Print(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitSelectElements(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("*")
		accept := sqlParser.SelectElements().Accept(visitor)
		assert.EqualValues(t, SelectElements{
			Star: true,
		}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("*, a b, c AS D")
		accept := sqlParser.SelectElements().Accept(visitor)
		assert.EqualValues(t, SelectElements{
			Star: true,
			SelectElements: []SelectElement{
				{
					FullColumnName: "a",
					Alias:          "b",
				},
				{
					FullColumnName: "c",
					Alias:          "D",
				},
			},
		}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("3", func(t *testing.T) {
		sqlParser, visitor, listener := createParser(" a b, C AS D")
		accept := sqlParser.SelectElements().Accept(visitor)
		assert.EqualValues(t, SelectElements{
			SelectElements: []SelectElement{
				{
					FullColumnName: "a",
					Alias:          "b",
				},
				{
					FullColumnName: "C",
					Alias:          "D",
				},
			},
		}, accept)
		fmt.Print(listener.errString.String())
	})
}
