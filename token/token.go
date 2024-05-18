package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT        = "IDENT"  // add, foobar, x, y, ...
	NUMBER       = "NUMBER" // 1343456, 1.23456
	STRING       = "STRING" // "foobar"
	AT           = "@"      // @ At symbol
	NULL         = "NULL"   // # null
	CURRENT_ARGS = "..."    // # ... function args

	// Operators
	TILDE         = "~"
	BANG          = "!"
	ASSIGN        = "="
	PLUS          = "+"
	MINUS         = "-"
	ASTERISK      = "*"
	SLASH         = "/"
	EXPONENT      = "**"
	MODULO        = "%"
	COMP_PLUS     = "+="
	COMP_MINUS    = "-="
	COMP_ASTERISK = "*="
	COMP_SLASH    = "/="
	COMP_EXPONENT = "**="
	COMP_MODULO   = "%="
	RANGE         = ".."

	// Logical operators
	AND = "&&"
	OR  = "||"

	LT            = "<"
	LT_EQ         = "<="
	GT            = ">"
	GT_EQ         = ">="
	COMBINED_COMP = "<=>"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
	DOT      = "."
	QUESTION = "?"

	// Keywords
	FUNCTION = "FN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	WHILE    = "WHILE"
	FOR      = "FOR"
)

type Token struct {
	Type     TokenType
	Position int
	Literal  string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"while":  WHILE,
	"for":    FOR,
	"null":   NULL,
}

// NumberSeparator is a separator for numbers eg. 1_000_000
var NumberSeparator = '_'

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
