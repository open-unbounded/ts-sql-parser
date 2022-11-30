// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // TsSqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseTsSqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTsSqlParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitSelectStmt(ctx *SelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitSelectElements(ctx *SelectElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitSelectElement(ctx *SelectElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitLimitClauseAtom(ctx *LimitClauseAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitUid(ctx *UidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitPredicateExpression(ctx *PredicateExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitColumnNameExpressionAtom(ctx *ColumnNameExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitLogicalOperator(ctx *LogicalOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitFullColumnName(ctx *FullColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsSqlParserVisitor) VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{} {
	return v.VisitChildren(ctx)
}
