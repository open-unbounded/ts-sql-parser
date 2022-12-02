// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // TsSqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseTsSqlParserListener is a complete listener for a parse tree produced by TsSqlParser.
type BaseTsSqlParserListener struct{}

var _ TsSqlParserListener = &BaseTsSqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTsSqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTsSqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTsSqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTsSqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseTsSqlParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseTsSqlParserListener) ExitRoot(ctx *RootContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseTsSqlParserListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseTsSqlParserListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterSelectElements is called when production selectElements is entered.
func (s *BaseTsSqlParserListener) EnterSelectElements(ctx *SelectElementsContext) {}

// ExitSelectElements is called when production selectElements is exited.
func (s *BaseTsSqlParserListener) ExitSelectElements(ctx *SelectElementsContext) {}

// EnterSelectColumnElement is called when production selectColumnElement is entered.
func (s *BaseTsSqlParserListener) EnterSelectColumnElement(ctx *SelectColumnElementContext) {}

// ExitSelectColumnElement is called when production selectColumnElement is exited.
func (s *BaseTsSqlParserListener) ExitSelectColumnElement(ctx *SelectColumnElementContext) {}

// EnterSelectFunctionElement is called when production selectFunctionElement is entered.
func (s *BaseTsSqlParserListener) EnterSelectFunctionElement(ctx *SelectFunctionElementContext) {}

// ExitSelectFunctionElement is called when production selectFunctionElement is exited.
func (s *BaseTsSqlParserListener) ExitSelectFunctionElement(ctx *SelectFunctionElementContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseTsSqlParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseTsSqlParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseTsSqlParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseTsSqlParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterLimitClauseAtom is called when production limitClauseAtom is entered.
func (s *BaseTsSqlParserListener) EnterLimitClauseAtom(ctx *LimitClauseAtomContext) {}

// ExitLimitClauseAtom is called when production limitClauseAtom is exited.
func (s *BaseTsSqlParserListener) ExitLimitClauseAtom(ctx *LimitClauseAtomContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseTsSqlParserListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseTsSqlParserListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterUid is called when production uid is entered.
func (s *BaseTsSqlParserListener) EnterUid(ctx *UidContext) {}

// ExitUid is called when production uid is exited.
func (s *BaseTsSqlParserListener) ExitUid(ctx *UidContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseTsSqlParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseTsSqlParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseTsSqlParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseTsSqlParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterPredicateExpression is called when production predicateExpression is entered.
func (s *BaseTsSqlParserListener) EnterPredicateExpression(ctx *PredicateExpressionContext) {}

// ExitPredicateExpression is called when production predicateExpression is exited.
func (s *BaseTsSqlParserListener) ExitPredicateExpression(ctx *PredicateExpressionContext) {}

// EnterExpressionAtomPredicate is called when production expressionAtomPredicate is entered.
func (s *BaseTsSqlParserListener) EnterExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// ExitExpressionAtomPredicate is called when production expressionAtomPredicate is exited.
func (s *BaseTsSqlParserListener) ExitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// EnterBinaryComparisonPredicate is called when production binaryComparisonPredicate is entered.
func (s *BaseTsSqlParserListener) EnterBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {
}

// ExitBinaryComparisonPredicate is called when production binaryComparisonPredicate is exited.
func (s *BaseTsSqlParserListener) ExitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {
}

// EnterConstantExpressionAtom is called when production constantExpressionAtom is entered.
func (s *BaseTsSqlParserListener) EnterConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// ExitConstantExpressionAtom is called when production constantExpressionAtom is exited.
func (s *BaseTsSqlParserListener) ExitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// EnterColumnNameExpressionAtom is called when production columnNameExpressionAtom is entered.
func (s *BaseTsSqlParserListener) EnterColumnNameExpressionAtom(ctx *ColumnNameExpressionAtomContext) {
}

// ExitColumnNameExpressionAtom is called when production columnNameExpressionAtom is exited.
func (s *BaseTsSqlParserListener) ExitColumnNameExpressionAtom(ctx *ColumnNameExpressionAtomContext) {
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseTsSqlParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseTsSqlParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAggregateWindowedFunction is called when production aggregateWindowedFunction is entered.
func (s *BaseTsSqlParserListener) EnterAggregateWindowedFunction(ctx *AggregateWindowedFunctionContext) {
}

// ExitAggregateWindowedFunction is called when production aggregateWindowedFunction is exited.
func (s *BaseTsSqlParserListener) ExitAggregateWindowedFunction(ctx *AggregateWindowedFunctionContext) {
}

// EnterFunctionArg is called when production functionArg is entered.
func (s *BaseTsSqlParserListener) EnterFunctionArg(ctx *FunctionArgContext) {}

// ExitFunctionArg is called when production functionArg is exited.
func (s *BaseTsSqlParserListener) ExitFunctionArg(ctx *FunctionArgContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseTsSqlParserListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseTsSqlParserListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BaseTsSqlParserListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BaseTsSqlParserListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseTsSqlParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseTsSqlParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterFullColumnName is called when production fullColumnName is entered.
func (s *BaseTsSqlParserListener) EnterFullColumnName(ctx *FullColumnNameContext) {}

// ExitFullColumnName is called when production fullColumnName is exited.
func (s *BaseTsSqlParserListener) ExitFullColumnName(ctx *FullColumnNameContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseTsSqlParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseTsSqlParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseTsSqlParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseTsSqlParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseTsSqlParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseTsSqlParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseTsSqlParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseTsSqlParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BaseTsSqlParserListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BaseTsSqlParserListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}
