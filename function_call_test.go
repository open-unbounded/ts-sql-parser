package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitFunctionArg(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("1")
		accept := sqlParser.FunctionArg().Accept(visitor)
		assert.EqualValues(t, ConstantDecimal{Val: 1}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("a")
		accept := sqlParser.FunctionArg().Accept(visitor)
		assert.EqualValues(t, FullColumnName("a"), accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("3", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("a")
		accept := sqlParser.FunctionArg().Accept(visitor)
		assert.EqualValues(t, FullColumnName("a"), accept)
		fmt.Print(listener.errString.String())
	})
}

func Test_parseTreeVisitor_VisitFunctionCall(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("COUNT(*)")
		accept := sqlParser.FunctionCall().Accept(visitor)
		assert.EqualValues(t, AggregateWindowedFunction{
			Function: "COUNT",
			StarArg:  true,
		}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("2", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("COUNT(1)")
		accept := sqlParser.FunctionCall().Accept(visitor)
		assert.EqualValues(t, AggregateWindowedFunction{
			Function:    "COUNT",
			FunctionArg: ConstantDecimal{Val: 1},
		}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("3", func(t *testing.T) {
		sqlParser, visitor, listener := createParser("COUNT(a)")
		accept := sqlParser.FunctionCall().Accept(visitor)
		assert.EqualValues(t, AggregateWindowedFunction{
			Function:    "COUNT",
			FunctionArg: FullColumnName("a"),
		}, accept)
		fmt.Print(listener.errString.String())
	})

	t.Run("4", func(t *testing.T) {
		for _, s := range []string{"AVG", "MAX", "MIN", "SUM", "STD", "STDDEV", "COUNT"} {
			sqlParser, visitor, listener := createParser(fmt.Sprintf("%s(a)", s))
			accept := sqlParser.FunctionCall().Accept(visitor)
			assert.EqualValues(t, AggregateWindowedFunction{
				Function:    s,
				FunctionArg: FullColumnName("a"),
			}, accept)
			fmt.Print(listener.errString.String())
		}
	})

}
