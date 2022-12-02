package parser

import (
	parser "github.com/open-unbounded/ts-sql-parser/internal/gen"
)

var (
	_ SelectElement = (*SelectColumnElement)(nil)
	_ SelectElement = (*SelectFunctionElement)(nil)
)

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
	SelectElement interface {
		isSelectElement()
	}

	FullColumnName string

	SelectColumnElement struct {
		FullColumnName FullColumnName
		Alias          string
	}
	SelectFunctionElement struct {
		FunctionCall FunctionCall
		Alias        string
	}
)

func (f FullColumnName) isFunctionArg()          {}
func (s SelectFunctionElement) isSelectElement() {}
func (s SelectColumnElement) isSelectElement()   {}

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

// -----------------

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

func (v *parseTreeVisitor) VisitSelectColumnElement(ctx *parser.SelectColumnElementContext) interface{} {
	element := SelectColumnElement{}

	fullColumnNameContext := ctx.FullColumnName()
	if fullColumnNameContext != nil {
		element.FullColumnName = fullColumnNameContext.Accept(v).(FullColumnName)
	}

	uidContext := ctx.Uid()
	if uidContext != nil {
		element.Alias = uidContext.Accept(v).(string)
	}

	return element
}

func (v *parseTreeVisitor) VisitSelectFunctionElement(ctx *parser.SelectFunctionElementContext) interface{} {
	selectFunctionElement := SelectFunctionElement{}
	functionCallContext := ctx.FunctionCall()
	if functionCallContext != nil {
		selectFunctionElement.FunctionCall = functionCallContext.Accept(v).(FunctionCall)
	}

	uidContext := ctx.Uid()
	if uidContext != nil {
		selectFunctionElement.Alias = uidContext.Accept(v).(string)
	}

	return selectFunctionElement
}
