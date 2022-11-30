lexer grammar TsSqlLexer;

options { caseInsensitive = true; }

channels { SQLCOMMENT, ERRORCHANNEL }

// SKIP

SPACE:                               [ \t\r\n]+    -> channel(HIDDEN);
SPEC_MYSQL_COMMENT:                  '/*!' .+? '*/' -> channel(SQLCOMMENT);
COMMENT_INPUT:                       '/*' .*? '*/' -> channel(HIDDEN);
LINE_COMMENT:                        (
                                       ('--' [ \t] | '#') ~[\r\n]* ('\r'? '\n' | EOF)
                                       | '--' ('\r'? '\n' | EOF)
                                     ) -> channel(HIDDEN);



DOT_PROPERTY: '.PROPERTY';
DOT_SERVICE: '.SERVICE';
DOT_EVENT: '.EVENT';
FALSE:                               'FALSE';
TRUE:                                'TRUE';

SELECT:                              'SELECT';
FROM:                                'FROM';
WHERE:                               'WHERE';
INTERVAL:                            'INTERVAL';
AND:                                 'AND';
IN:                                  'IN';
IS:                                  'IS';
OR:                                  'OR';
AS:                                  'AS';
XOR:                                 'XOR';

JOIN:                                'JOIN';
LEFT:                                'LEFT';
LIKE:                                'LIKE';
LIMIT:                               'LIMIT';
NOT:                                 'NOT';



STAR:                                '*';
DIVIDE:                              '/';
MODULE:                              '%';
PLUS:                                '+';
MINUS:                               '-';
DIV:                                 'DIV';
MOD:                                 'MOD';



EQUAL_SYMBOL:                        '=';
GREATER_SYMBOL:                      '>';
LESS_SYMBOL:                         '<';
EXCLAMATION_SYMBOL:                  '!';

BIT_NOT_OP:                          '~';
BIT_OR_OP:                           '|';
BIT_AND_OP:                          '&';
BIT_XOR_OP:                          '^';

TIME_INTERVAL:  DEC_DIGIT+ ('S'|'M'|'H'|'D');

DOT:                                 '.';
LR_BRACKET:                          '(';
RR_BRACKET:                          ')';
COMMA:                               ',';
SEMI:                                ';';
AT_SIGN:                             '@';
ZERO_DECIMAL:                        '0';
ONE_DECIMAL:                         '1';
TWO_DECIMAL:                         '2';
SINGLE_QUOTE_SYMB:                   '\'';
DOUBLE_QUOTE_SYMB:                   '"';
REVERSE_QUOTE_SYMB:                  '`';
COLON_SYMB:                          ':';
ID:                                  ID_LITERAL;

NULL_LITERAL:                        'NULL';
NULL_SPEC_LITERAL:                   '\\' 'N';
STRING_LITERAL:                      DQUOTA_STRING | SQUOTA_STRING | BQUOTA_STRING;
fragment DQUOTA_STRING:              '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
fragment SQUOTA_STRING:              '\'' ('\\'. | '\'\'' | ~('\'' | '\\'))* '\'';
fragment BQUOTA_STRING:              '`' ( '\\'. | '``' | ~('`'|'\\'))* '`';
PROPERTY:'PROPERTY';
SERVICE:'SERVICE';
fragment ID_LITERAL:                 [A-Z_$0-9\u0080-\uFFFF]*?[A-Z_$\u0080-\uFFFF]+?[A-Z_$0-9\u0080-\uFFFF]*;
fragment DEC_DIGIT:                  [0-9];
DECIMAL_LITERAL:                     DEC_DIGIT+;
REAL_LITERAL:                        (DEC_DIGIT+)? '.' DEC_DIGIT+;

ERROR_RECONGNIGION:                  .    -> channel(ERRORCHANNEL);