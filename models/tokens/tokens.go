package tokens

type TokenType int64

type Token struct {
	Type TokenType
	Literal string
}

func (t *Token) Name() string {
	return translator[t.Type]
}

var translator = map[TokenType]string{
	EOF: "EOF",
	Integer: "Integer",
	Plus: "+",
	Semicolon: ";",
}

const (
	EOF			TokenType = iota
	Integer
	Plus
	Semicolon
)
