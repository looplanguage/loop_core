package precedence

import "github.com/looplanguage/loop/models/tokens"

const (
	LOWEST int = iota
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
	INDEX
)

var Precedences = map[tokens.TokenType]int{
	tokens.Plus:              SUM,
	tokens.Minus:             SUM,
	tokens.Asterisk:          PRODUCT,
	tokens.Slash:             PRODUCT,
	tokens.LeftParenthesis:   CALL,
	tokens.Equals:            EQUALS,
	tokens.LesserEqualsThan:  LESSGREATER,
	tokens.LesserThan:        LESSGREATER,
	tokens.GreaterThan:       LESSGREATER,
	tokens.GreaterEqualsThan: LESSGREATER,
	tokens.LeftBracket:       INDEX,
}
