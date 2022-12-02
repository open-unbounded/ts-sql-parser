package parser

import (
	parser "github.com/open-unbounded/ts-sql-parser/internal/gen"
)

var (
	_ FunctionCall = (*AggregateWindowedFunction)(nil)
)

type (
	FunctionCall interface {
		FunctionArg
		isFunctionCall()
	}
	AggregateWindowedFunction struct {
		Function    string
		StarArg     bool
		FunctionArg FunctionArg
	}
	FunctionArg interface {
		isFunctionArg()
	}
)

func (a AggregateWindowedFunction) isFunctionArg()  {}
func (a AggregateWindowedFunction) isFunctionCall() {}

func (v *parseTreeVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	var functionCall FunctionCall
	if aggregateWindowedFunctionContext := ctx.AggregateWindowedFunction(); aggregateWindowedFunctionContext != nil {
		functionCall = aggregateWindowedFunctionContext.Accept(v).(AggregateWindowedFunction)
	}

	return functionCall
}
func (v *parseTreeVisitor) VisitAggregateWindowedFunction(ctx *parser.AggregateWindowedFunctionContext) interface{} {
	aggregateWindowedFunction := AggregateWindowedFunction{}
	if ctx.AVG() != nil {
		aggregateWindowedFunction.Function = "AVG"
	} else if ctx.MAX() != nil {
		aggregateWindowedFunction.Function = "MAX"
	} else if ctx.MAX() != nil {
		aggregateWindowedFunction.Function = "MAX"
	} else if ctx.MIN() != nil {
		aggregateWindowedFunction.Function = "MIN"
	} else if ctx.SUM() != nil {
		aggregateWindowedFunction.Function = "SUM"
	} else if ctx.STD() != nil {
		aggregateWindowedFunction.Function = "STD"
	} else if ctx.STDDEV() != nil {
		aggregateWindowedFunction.Function = "STDDEV"
	} else if ctx.COUNT() != nil {
		aggregateWindowedFunction.Function = "COUNT"
	}

	functionArgContext := ctx.FunctionArg()
	if functionArgContext != nil {
		aggregateWindowedFunction.FunctionArg = functionArgContext.Accept(v).(FunctionArg)
	}

	aggregateWindowedFunction.StarArg = ctx.STAR() != nil

	return aggregateWindowedFunction

}

func (v *parseTreeVisitor) VisitFunctionArg(ctx *parser.FunctionArgContext) interface{} {
	var functionArg FunctionArg
	if constantContext := ctx.Constant(); constantContext != nil {
		functionArg = constantContext.Accept(v).(Constant)
	} else if columnNameContext := ctx.FullColumnName(); columnNameContext != nil {
		functionArg = columnNameContext.Accept(v).(FullColumnName)
	} else if functionCallContext := ctx.FunctionCall(); functionCallContext != nil {
		functionArg = functionCallContext.Accept(v).(FunctionCall)
	}

	return functionArg
}
