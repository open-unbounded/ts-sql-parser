package parser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type VerboseListener struct {
	antlr.ConsoleErrorListener
	errString bytes.Buffer
}

func (l *VerboseListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.ConsoleErrorListener.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
	underLineError(&l.errString, recognizer, offendingSymbol.(antlr.Token), line, column)
}

func underLineError(errString *bytes.Buffer, recognizer antlr.Recognizer, offendingToken antlr.Token, line, column int) {
	inputStream := recognizer.(*antlr.BaseParser).GetTokenStream().(*antlr.CommonTokenStream).GetTokenSource().GetInputStream()
	input := inputStream.(*antlr.InputStream).String()
	lines := strings.Split(input, "\n")
	errorLine := lines[line-1]
	fmt.Fprintln(errString, errorLine)
	for i := 0; i < column; i++ {
		fmt.Fprintf(errString, " ")
	}
	start := offendingToken.GetStart()
	stop := offendingToken.GetStop()
	if start >= 0 && stop >= 0 {
		for i := start; i <= stop; i++ {
			fmt.Fprintf(errString, "^")
		}
	}
	fmt.Fprintln(errString)
}
