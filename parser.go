package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	parser "github.com/open-unbounded/ts-sql-parser/internal/gen"
)

var ErrEmptySql = errors.New("empty sql")

type parseTreeVisitor struct {
	parser.BaseTsSqlParserVisitor
}

func Parse(sql string) (result interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	if strings.TrimSpace(sql) == "" {
		return nil, ErrEmptySql
	}

	inputStream := antlr.NewInputStream(sql)
	lexer := parser.NewTsSqlLexer(inputStream)
	lexer.RemoveErrorListeners()

	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	mysqlParser := parser.NewTsSqlParser(tokens)
	mysqlParser.RemoveErrorListeners()

	lis := &errorListener{}
	mysqlParser.AddErrorListener(lis)

	visitor := &parseTreeVisitor{}
	accept := mysqlParser.Root().Accept(visitor)

	if lis.errString.Len() != 0 {
		return nil, errors.New(lis.errString.String())
	}

	return accept, nil
}

func (v *parseTreeVisitor) VisitRoot(ctx *parser.RootContext) interface{} {
	allSelectStmt := ctx.AllSelectStmt()

	selectStmts := make([]SelectStmt, 0, len(allSelectStmt))
	for _, stmtContext := range allSelectStmt {
		selectStmt := stmtContext.Accept(v).(SelectStmt)
		selectStmts = append(selectStmts, selectStmt)
	}

	return selectStmts
}

// -----------------

func (v *parseTreeVisitor) VisitFullColumnName(ctx *parser.FullColumnNameContext) interface{} {
	return FullColumnName(ctx.Uid().Accept(v).(string))
}

// -----------------

func (v *parseTreeVisitor) VisitUid(ctx *parser.UidContext) interface{} {
	return ctx.GetText()
}

// -----------------

type TableName struct {
	TableName  string
	SourceType string
}

func (v *parseTreeVisitor) VisitTableName(ctx *parser.TableNameContext) interface{} {
	var sourceType string
	if ctx.DOT_PROPERTY() != nil {
		sourceType = "PROPERTY"
	} else if ctx.DOT_SERVICE() != nil {
		sourceType = "SERVICE"
	} else if ctx.DOT_EVENT() != nil {
		sourceType = "EVENT"
	}

	return TableName{
		TableName:  ctx.Uid().Accept(v).(string),
		SourceType: sourceType,
	}
}

// -----------------

type (
	FromClause struct {
		TableName  TableName
		Expression Expression
	}
)

func (v *parseTreeVisitor) VisitFromClause(ctx *parser.FromClauseContext) interface{} {
	fromClause := FromClause{}
	tableNameContext := ctx.TableName()
	if tableNameContext != nil {
		fromClause.TableName = tableNameContext.Accept(v).(TableName)
	}

	expressionContext := ctx.Expression()
	if expressionContext != nil {
		fromClause.Expression = expressionContext.Accept(v).(Expression)
	}

	return fromClause
}

// -----------------

type (
	Expression interface {
		isExpression()
	}
)

// -----------------
var _ Expression = (*LogicalExpression)(nil)

type (
	LogicalExpression struct {
		LeftExpression  Expression
		Op              string
		RightExpression Expression
	}
)

func (l LogicalExpression) isExpression() {}

func (v *parseTreeVisitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) interface{} {
	leftCtx := ctx.GetLeft()
	rightCtx := ctx.GetRight()
	return LogicalExpression{
		LeftExpression:  leftCtx.Accept(v).(Expression),
		Op:              ctx.LogicalOperator().Accept(v).(string),
		RightExpression: rightCtx.Accept(v).(Expression),
	}
}

// -----------------
var _ ExpressionAtom = (*ConstantExpressionAtom)(nil)
var _ ExpressionAtom = (*ColumnNameExpressionAtom)(nil)
var _ Expression = (*PredicateExpression)(nil)

type (
	ExpressionAtom interface {
		isExpressionAtom()
	}
	PredicateExpression struct {
		Predicate Predicate
	}

	ColumnNameExpressionAtom struct {
		ColumnName string
	}
)

func (p PredicateExpression) isFunctionArg()         {}
func (p PredicateExpression) isExpression()          {}
func (c ColumnNameExpressionAtom) isExpressionAtom() {}

func (v *parseTreeVisitor) VisitPredicateExpression(ctx *parser.PredicateExpressionContext) interface{} {
	return PredicateExpression{Predicate: ctx.Predicate().Accept(v).(Predicate)}
}

// -----------------

type ConstantExpressionAtom struct {
	Constant Constant
}

func (c ConstantExpressionAtom) isExpressionAtom() {}

func (v *parseTreeVisitor) VisitConstantExpressionAtom(ctx *parser.ConstantExpressionAtomContext) interface{} {
	return ConstantExpressionAtom{Constant: ctx.Constant().Accept(v).(Constant)}
}

// -----------------

type LimitClause struct {
	Valid  bool
	Offset int
	Limit  int
}

func (v *parseTreeVisitor) VisitLimitClause(ctx *parser.LimitClauseContext) interface{} {
	limitClause := LimitClause{}
	limitCtx := ctx.GetLimit()
	if limitCtx != nil {
		limitClause.Limit = limitCtx.Accept(v).(int)
		limitClause.Valid = true
	}

	offsetCtx := ctx.GetOffset()
	if offsetCtx != nil {
		limitClause.Offset = offsetCtx.Accept(v).(int)
		limitClause.Valid = true
	}

	return limitClause
}

func (v *parseTreeVisitor) VisitLimitClauseAtom(ctx *parser.LimitClauseAtomContext) interface{} {
	return toInt(ctx.GetText())
}

// -----------------

type WindowClause struct {
	Duration string
}

func (v *parseTreeVisitor) VisitWindowClause(ctx *parser.WindowClauseContext) interface{} {
	windowClause := WindowClause{}
	timeIntervalCtx := ctx.TIME_INTERVAL()
	if timeIntervalCtx != nil {
		windowClause.Duration = timeIntervalCtx.GetText()
	}

	return windowClause
}

// -----------------

func (v *parseTreeVisitor) VisitLogicalOperator(ctx *parser.LogicalOperatorContext) interface{} {
	return ctx.GetText()
}

// -----------------
var _ Predicate = (*BinaryComparisonPredicate)(nil)
var _ ExpressionAtom = (*BinaryComparisonPredicate)(nil)

type (
	Predicate interface {
		Expression
		isPredicate()
	}
	BinaryComparisonPredicate struct {
		Left  Expression
		Op    string
		Right Expression
	}
)

func (b BinaryComparisonPredicate) isExpressionAtom() {}
func (b BinaryComparisonPredicate) isExpression()     {}
func (b BinaryComparisonPredicate) isPredicate()      {}

func (v *parseTreeVisitor) VisitBinaryComparisonPredicate(ctx *parser.BinaryComparisonPredicateContext) interface{} {
	binaryComparisonPredicate := BinaryComparisonPredicate{}
	leftCtx := ctx.GetLeft()
	if leftCtx != nil {
		binaryComparisonPredicate.Left = leftCtx.Accept(v).(Expression)
	}

	comparisonOperatorCtx := ctx.ComparisonOperator()
	if comparisonOperatorCtx != nil {
		binaryComparisonPredicate.Op = comparisonOperatorCtx.Accept(v).(string)
	}

	rightCtx := ctx.GetRight()
	if rightCtx != nil {
		binaryComparisonPredicate.Right = rightCtx.Accept(v).(Expression)
	}

	return binaryComparisonPredicate
}

func (v *parseTreeVisitor) VisitComparisonOperator(ctx *parser.ComparisonOperatorContext) interface{} {
	return ctx.GetText()
}

// -----------------
var _ Predicate = (*ExpressionAtomPredicate)(nil)
var _ ExpressionAtom = (*ExpressionAtomPredicate)(nil)

type (
	ExpressionAtomPredicate struct {
		ExpressionAtom ExpressionAtom
	}
)

func (e ExpressionAtomPredicate) isExpressionAtom() {}
func (e ExpressionAtomPredicate) isExpression()     {}
func (e ExpressionAtomPredicate) isPredicate()      {}

func (v *parseTreeVisitor) VisitExpressionAtomPredicate(ctx *parser.ExpressionAtomPredicateContext) interface{} {
	return ExpressionAtomPredicate{ExpressionAtom: ctx.ExpressionAtom().Accept(v).(ExpressionAtom)}
}

// -----------------

func (v *parseTreeVisitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	return ctx.GetText()
}

// -----------------

func (v *parseTreeVisitor) VisitColumnNameExpressionAtom(ctx *parser.ColumnNameExpressionAtomContext) interface{} {
	return ColumnNameExpressionAtom{
		ColumnName: ctx.ColumnName().Accept(v).(string),
	}
}

func (v *parseTreeVisitor) VisitColumnName(ctx *parser.ColumnNameContext) interface{} {
	return ctx.Uid().Accept(v).(string)
}

// -----------------

func toInt(s string) int {
	a, _ := strconv.Atoi(s)

	return a
}

func toDecimal(v string) float64 {
	float, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(err)
	}
	return float
}
