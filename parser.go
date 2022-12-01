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
			err = errors.New(fmt.Sprintf("%v", e))
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

type (
	SelectStmt struct {
		SelectElements SelectElements
		FromClause     FromClause
		LimitClause    LimitClause
		WindowClause   WindowClause
	}

	SelectElements struct {
		Star           bool
		SelectElements []SelectElement
	}
	SelectElement struct {
		FullColumnName string
		Alias          string
	}
)

func (v *parseTreeVisitor) VisitSelectStmt(ctx *parser.SelectStmtContext) interface{} {
	stmt := SelectStmt{}

	elementsContext := ctx.SelectElements()
	if elementsContext != nil {
		stmt.SelectElements = elementsContext.Accept(v).(SelectElements)
	}

	fromClauseCtx := ctx.FromClause()
	if fromClauseCtx != nil {
		stmt.FromClause = fromClauseCtx.Accept(v).(FromClause)
	}

	limitClauseContext := ctx.LimitClause()
	if limitClauseContext != nil {
		stmt.LimitClause = limitClauseContext.Accept(v).(LimitClause)
	}

	windowClauseContext := ctx.WindowClause()
	if windowClauseContext != nil {
		stmt.WindowClause = windowClauseContext.Accept(v).(WindowClause)
	}

	return stmt
}

func (v *parseTreeVisitor) VisitSelectElements(ctx *parser.SelectElementsContext) interface{} {
	selectElementCtxs := ctx.AllSelectElement()

	var selectElements []SelectElement
	for _, elementCtx := range selectElementCtxs {
		selectElements = append(selectElements, elementCtx.Accept(v).(SelectElement))
	}

	return SelectElements{
		Star:           ctx.STAR() != nil,
		SelectElements: selectElements,
	}

}

func (v *parseTreeVisitor) VisitSelectElement(ctx *parser.SelectElementContext) interface{} {
	element := SelectElement{}

	fullColumnNameContext := ctx.FullColumnName()
	if fullColumnNameContext != nil {
		element.FullColumnName = fullColumnNameContext.Accept(v).(string)
	}

	uidContext := ctx.Uid()
	if uidContext != nil {
		element.Alias = uidContext.Accept(v).(string)
	}

	return element
}

// -----------------

func (v *parseTreeVisitor) VisitFullColumnName(ctx *parser.FullColumnNameContext) interface{} {
	return ctx.Uid().Accept(v).(string)
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

type (
	Constant interface {
		isConstant()
	}
	ConstantBool struct {
		Val bool
	}
	ConstantString struct {
		Val string
	}
	ConstantNull struct {
		Val string
	}
	ConstantDecimal struct {
		Val float64
	}
)

func (c ConstantDecimal) isConstant() {}
func (c ConstantNull) isConstant()    {}
func (c ConstantBool) isConstant()    {}
func (c ConstantString) isConstant()  {}

func (v *parseTreeVisitor) VisitConstant(ctx *parser.ConstantContext) interface{} {
	stringLiteralContext := ctx.StringLiteral()
	if stringLiteralContext != nil {
		return ConstantString{Val: stringLiteralContext.Accept(v).(string)}
	}

	decimalContext := ctx.DecimalLiteral()
	if decimalContext != nil {
		var decimal string
		minusCtx := ctx.MINUS()
		if minusCtx != nil {
			decimal = "-"
		}
		decimal += decimalContext.GetText()

		return ConstantDecimal{Val: toDecimal(decimal)}
	}

	booleanLiteralContext := ctx.BooleanLiteral()
	if booleanLiteralContext != nil {
		b := booleanLiteralContext.GetText()
		return ConstantBool{Val: b == "TRUE"}
	}

	nullLiteralContext := ctx.GetNullLiteral()
	if nullLiteralContext != nil {
		notContext := ctx.NOT()
		var constant string
		if notContext != nil {
			constant = notContext.GetText()
		}
		if constant != "" {
			constant += " "
		}
		constant += nullLiteralContext.GetText()

		return ConstantNull{Val: constant}
	}

	return nil
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

type LimitClause struct {
	Offset int
	Limit  int
}

func (v *parseTreeVisitor) VisitLimitClause(ctx *parser.LimitClauseContext) interface{} {
	limitClause := LimitClause{}
	limitCtx := ctx.GetLimit()
	if limitCtx != nil {
		limitClause.Limit = limitCtx.Accept(v).(int)
	}

	offsetCtx := ctx.GetOffset()
	if offsetCtx != nil {
		limitClause.Offset = offsetCtx.Accept(v).(int)
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
