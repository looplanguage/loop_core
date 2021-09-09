package precedence

import "git.kanersps.pw/loop/models/tokens"

const (
	LOWEST int = iota
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
)

var Precedences = map[tokens.TokenType]int{
	tokens.Plus: SUM,
}
