parser grammar TsSqlParser;

options { tokenVocab=TsSqlLexer; }


root
    : (selectStmt SEMI? | emptyStatement_) (selectStmt SEMI? | emptyStatement_)*
    ;
selectStmt
    : SELECT selectElements fromClause limitClause? windowClause?
    ;

selectElements
    : (star='*' | selectElement) (',' selectElement)*
    ;

selectElement
    : fullColumnName (AS? uid)? #selectColumnElement
    | functionCall (AS? uid)?   #selectFunctionElement
    ;

fromClause
    : FROM tableName  (WHERE expression )?
    ;

limitClause
    : LIMIT (offset=limitClauseAtom ',')? limit=limitClauseAtom
    ;

limitClauseAtom
    : DECIMAL_LITERAL
    | TWO_DECIMAL
    | ONE_DECIMAL
    | ZERO_DECIMAL
    ;

windowClause
    : INTERVAL'(' TIME_INTERVAL ')'
    ;

uid
    : ID
    ;

tableName
    : uid (DOT_PROPERTY | DOT_SERVICE | DOT_EVENT)
    ;

expression
    : left=expression logicalOperator right=expression            #logicalExpression
    | predicate                                                   #predicateExpression
    ;

predicate
    : left=predicate comparisonOperator right=predicate           #binaryComparisonPredicate
    | expressionAtom                                              #expressionAtomPredicate
    ;

expressionAtom
    : constant                                                    #constantExpressionAtom
    | columnName                                                  #columnNameExpressionAtom
    ;

functionCall
    : aggregateWindowedFunction
    ;

aggregateWindowedFunction
    :(AVG | MAX | MIN | SUM)
           '(' aggregator=(ALL | DISTINCT)? functionArg ')'
         | COUNT '(' (starArg='*' | aggregator=ALL? functionArg ) ')'
    ;

functionArg
    : constant | fullColumnName | functionCall | expression
    ;

columnName
    : uid
    ;

logicalOperator
    : AND | '&' '&' | XOR | OR | '|' '|'
    ;

comparisonOperator
    : '=' | '>' | '<' | '<' '=' | '>' '='
    | '<' '>' | '!' '=' | '<' '=' '>'
    ;

fullColumnName
    : uid
    ;

constant
    : stringLiteral | decimalLiteral
    | '-' decimalLiteral
    | booleanLiteral
    | NOT? nullLiteral=(NULL_LITERAL | NULL_SPEC_LITERAL)
    ;

stringLiteral
    : STRING_LITERAL
    ;

decimalLiteral
    : DECIMAL_LITERAL | ZERO_DECIMAL | ONE_DECIMAL | TWO_DECIMAL | REAL_LITERAL
    ;

booleanLiteral
    : TRUE | FALSE;

emptyStatement_
    : SEMI
    ;
