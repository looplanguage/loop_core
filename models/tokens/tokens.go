package tokens

type TokenType int64

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"var": VariableDeclaration,
}

func (t *Token) Name() string {
	return translator[t.Type]
}

var translator = map[TokenType]string{
	EOF:                 "EOF",
	Integer:             "Integer",
	Plus:                "+",
	Semicolon:           ";",
	VariableDeclaration: "var",
	Identifier:          "Identifier",
	Assign:              "Assign",
}

const (
	EOF TokenType = iota
	Integer
	Plus
	Semicolon
	VariableDeclaration
	Identifier
	Assign
)

func FindKeyword(keyword string) TokenType {
	if t, ok := keywords[keyword]; ok {
		return t
	}

	return Identifier
}
