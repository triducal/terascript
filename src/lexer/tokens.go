package lexer

import "fmt"

type TokenKind string

const (
	EOF        TokenKind = "EOF"
	NUMBER     TokenKind = "NUMBER"
	STRING     TokenKind = "STRING"
	IDENTIFIER TokenKind = "IDENTIFIER"

	OPEN_BRACKET  TokenKind = "OPEN_BRACKET"
	CLOSE_BRACKET TokenKind = "CLOSE_BRACKET"
	OPEN_CURLY    TokenKind = "OPEN_CURLY"
	CLOSE_CURLY   TokenKind = "CLOSE_CURLY"
	OPEN_PAREN    TokenKind = "OPEN_PAREN"
	CLOSE_PAREN   TokenKind = "CLOSE_PAREN"

	ASSIGNMENT TokenKind = "ASSIGNMENT"
	EQUALS     TokenKind = "EQUALS"
	NOT        TokenKind = "NOT"
	NOT_EQUALS TokenKind = "NOT_EQUALS"

	LESS           TokenKind = "LESS"
	LESS_EQUALS    TokenKind = "LESS_EQUALS"
	GREATER        TokenKind = "GREATER"
	GREATER_EQUALS TokenKind = "GREATER_EQUALS"

	OR  TokenKind = "OR"
	AND TokenKind = "AND"

	DOT       TokenKind = "DOT"
	DOT_DOT   TokenKind = "DOT_DOT"
	SEMICOLON TokenKind = "SEMICOLON"
	COLON     TokenKind = "COLON"
	QUESTION  TokenKind = "QUESTION"
	COMMA     TokenKind = "COMMA"

	PLUS_PLUS      TokenKind = "PLUS_PLUS"
	MINUS_MINUS    TokenKind = "MINUS_MINUS"
	PLUS_EQUALS    TokenKind = "PLUS_EQUALS"
	MINUS_EQUALS   TokenKind = "MINUS_EQUALS"
	SLASH_EQUALS   TokenKind = "SLASH_EQUALS"
	STAR_EQUALS    TokenKind = "STAR_EQUALS"
	PERCENT_EQUALS TokenKind = "PERCENT_EQUALS"

	PLUS    TokenKind = "PLUS"
	MINUS   TokenKind = "MINUS"
	SLASH   TokenKind = "SLASH"
	STAR    TokenKind = "STAR"
	PERCENT TokenKind = "PERCENT"

	// RESERVED KEYWORDS
	VAR    TokenKind = "VAR"
	CONST  TokenKind = "CONST"
	IMPORT TokenKind = "IMPORT"
	FN     TokenKind = "FN"
	IF     TokenKind = "IF"
	ELSE   TokenKind = "ELSE"
	ELIF   TokenKind = "ELIF"
	FOR    TokenKind = "FOR"
	EXPORT TokenKind = "EXPORT"
)

var keywords map[string]TokenKind = map[string]TokenKind{
	"var":    VAR,
	"const":  CONST,
	"import": IMPORT,
	"fn":     FN,
	"if":     IF,
	"else":   ELSE,
	"elif":   ELIF,
	"for":    FOR,
	"export": EXPORT,
}

type Token struct {
	Kind  TokenKind
	Value string
	Pos   int
}

func NewToken(kind TokenKind, value string, pos int) Token {
	return Token{kind, value, pos}
}

func (token Token) isOneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if expected == token.Kind {
			return true
		}
	}

	return false
}

func (token Token) Debug() {
	if token.isOneOfMany(IDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", token.Kind, token.Value)
	} else {
		fmt.Printf("%s ()\n", token.Kind)
	}
}
