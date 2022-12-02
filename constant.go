package parser

import (
	parser "github.com/open-unbounded/ts-sql-parser/internal/gen"
)

var (
	_ Constant = (*ConstantBool)(nil)
	_ Constant = (*ConstantString)(nil)
	_ Constant = (*ConstantNull)(nil)
	_ Constant = (*ConstantDecimal)(nil)
)

type (
	Constant interface {
		isConstant()
		FunctionArg
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

func (c ConstantDecimal) isFunctionArg() {}
func (c ConstantNull) isFunctionArg()    {}
func (c ConstantString) isFunctionArg()  {}
func (c ConstantBool) isFunctionArg()    {}
func (c ConstantDecimal) isConstant()    {}
func (c ConstantNull) isConstant()       {}
func (c ConstantBool) isConstant()       {}
func (c ConstantString) isConstant()     {}

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
