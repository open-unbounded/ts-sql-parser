// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // SqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseSqlParserListener is a complete listener for a parse tree produced by SqlParser.
type BaseSqlParserListener struct{}

var _ SqlParserListener = &BaseSqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseSqlParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseSqlParserListener) ExitRoot(ctx *RootContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseSqlParserListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseSqlParserListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterSelectElements is called when production selectElements is entered.
func (s *BaseSqlParserListener) EnterSelectElements(ctx *SelectElementsContext) {}

// ExitSelectElements is called when production selectElements is exited.
func (s *BaseSqlParserListener) ExitSelectElements(ctx *SelectElementsContext) {}

// EnterSelectElement is called when production selectElement is entered.
func (s *BaseSqlParserListener) EnterSelectElement(ctx *SelectElementContext) {}

// ExitSelectElement is called when production selectElement is exited.
func (s *BaseSqlParserListener) ExitSelectElement(ctx *SelectElementContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseSqlParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseSqlParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseSqlParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseSqlParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterLimitClauseAtom is called when production limitClauseAtom is entered.
func (s *BaseSqlParserListener) EnterLimitClauseAtom(ctx *LimitClauseAtomContext) {}

// ExitLimitClauseAtom is called when production limitClauseAtom is exited.
func (s *BaseSqlParserListener) ExitLimitClauseAtom(ctx *LimitClauseAtomContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseSqlParserListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseSqlParserListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterUid is called when production uid is entered.
func (s *BaseSqlParserListener) EnterUid(ctx *UidContext) {}

// ExitUid is called when production uid is exited.
func (s *BaseSqlParserListener) ExitUid(ctx *UidContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseSqlParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseSqlParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseSqlParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseSqlParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterPredicateExpression is called when production predicateExpression is entered.
func (s *BaseSqlParserListener) EnterPredicateExpression(ctx *PredicateExpressionContext) {}

// ExitPredicateExpression is called when production predicateExpression is exited.
func (s *BaseSqlParserListener) ExitPredicateExpression(ctx *PredicateExpressionContext) {}

// EnterExpressionAtomPredicate is called when production expressionAtomPredicate is entered.
func (s *BaseSqlParserListener) EnterExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// ExitExpressionAtomPredicate is called when production expressionAtomPredicate is exited.
func (s *BaseSqlParserListener) ExitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// EnterBinaryComparisonPredicate is called when production binaryComparisonPredicate is entered.
func (s *BaseSqlParserListener) EnterBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {
}

// ExitBinaryComparisonPredicate is called when production binaryComparisonPredicate is exited.
func (s *BaseSqlParserListener) ExitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) {
}

// EnterConstantExpressionAtom is called when production constantExpressionAtom is entered.
func (s *BaseSqlParserListener) EnterConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// ExitConstantExpressionAtom is called when production constantExpressionAtom is exited.
func (s *BaseSqlParserListener) ExitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// EnterColumnNameExpressionAtom is called when production columnNameExpressionAtom is entered.
func (s *BaseSqlParserListener) EnterColumnNameExpressionAtom(ctx *ColumnNameExpressionAtomContext) {}

// ExitColumnNameExpressionAtom is called when production columnNameExpressionAtom is exited.
func (s *BaseSqlParserListener) ExitColumnNameExpressionAtom(ctx *ColumnNameExpressionAtomContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseSqlParserListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseSqlParserListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BaseSqlParserListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BaseSqlParserListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSqlParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSqlParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterFullColumnName is called when production fullColumnName is entered.
func (s *BaseSqlParserListener) EnterFullColumnName(ctx *FullColumnNameContext) {}

// ExitFullColumnName is called when production fullColumnName is exited.
func (s *BaseSqlParserListener) ExitFullColumnName(ctx *FullColumnNameContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseSqlParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseSqlParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseSqlParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseSqlParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseSqlParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseSqlParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSqlParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSqlParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BaseSqlParserListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BaseSqlParserListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}
