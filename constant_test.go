package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
