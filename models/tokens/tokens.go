package tokens

type TokenType int64

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

var keywords = map[string]TokenType{
	"var":    VariableDeclaration,
	"fun":    Func,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
	"while":  While,
	"import": Import,
	"as":     As,
	"export": Export,
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
	Assign:              "=",
	Asterisk:            "*",
	Slash:               "/",
	Minus:               "-",
	LeftParenthesis:     "(",
	RightParenthesis:    ")",
	Func:                "fun",
	LeftBrace:           "{",
	RightBrace:          "}",
	Comma:               ",",
	True:                "true",
	False:               "false",
	Equals:              "==",
	If:                  "if",
	Else:                "else",
	GreaterThan:         ">",
	LesserThan:          "<",
	GreaterEqualsThan:   ">=",
	LesserEqualsThan:    "<=",
	NotEquals:           "!=",
	LeftBracket:         "[",
	RightBracket:        "]",
	Return:              "return",
	String:              "string",
	While:               "while",
	Colon:               "colon",
	Import:              "import",
	As:                  "as",
	Export:              "export",
	Bang:                "!",
}

const (
	EOF TokenType = iota
	Integer
	Plus
	Semicolon
	VariableDeclaration
	Identifier
	Assign
	Asterisk
	Slash
	Minus
	LeftParenthesis
	RightParenthesis
	Func
	LeftBrace
	RightBrace
	Comma
	True
	False
	Equals
	If
	Else
	GreaterThan
	LesserThan
	GreaterEqualsThan
	LesserEqualsThan
	LeftBracket
	RightBracket
	Return
	String
	While
	Colon
	Import
	As
	Export
	Bang
	NotEquals
)

func FindKeyword(keyword string) TokenType {
	if t, ok := keywords[keyword]; ok {
		return t
	}

	return Identifier
}
