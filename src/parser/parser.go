package parser

import (
	"fmt"
	"os"
	"strings"

	"terascript/src/ast"
	"terascript/src/lexer"

	"github.com/fatih/color"
)

type parser struct {
	lexer  *lexer.Lexer
	errors []string
	pos    int
}

func createParser(lexer *lexer.Lexer) *parser {
	createTokenTypeLookups()
	createTokenLookups()
	return &parser{
		lexer: lexer,
	}
}

func Parse(lexer *lexer.Lexer) (ast.File, []string) {
	Body := make([]ast.Stmt, 0)
	p := createParser(lexer)

	for p.hasTokens() {
		Body = append(Body, parse_stmt(p))
	}

	return ast.File{
		Statements: Body,
	}, p.errors
}

func (p *parser) currentToken() lexer.Token {
	return p.lexer.Tokens[p.pos]
}

func (p *parser) advance() lexer.Token {
	token := p.currentToken()
	p.pos++
	return token
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.lexer.Tokens) && p.currentToken().Kind != lexer.EOF
}

func (p *parser) expectError(expectedKind lexer.TokenKind, err string) lexer.Token {
	token := p.currentToken()
	kind := token.Kind

	if kind != expectedKind {
		lineNum, collumn, lineText := p.lexer.ErrorLine(token.Pos)
		p.errors = append(p.errors, fmt.Sprintf(`[%d:%d] %s`, lineNum, collumn-1, lineText))
		p.errors = append(p.errors, fmt.Sprintf(`%s^`, strings.Repeat(" ", collumn+len(fmt.Sprint(lineNum))+len(fmt.Sprint(collumn))+2)))
		if err == "" {
			p.errors = append(p.errors, fmt.Sprintf(`Expected %s But Received %s`, expectedKind, kind))
		} else {
			p.errors = append(p.errors, err)
		}

		p.printAllErrors()
		os.Exit(99)

	}

	return p.advance()
}

func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, "")
}

func (p *parser) Error(err string) {
	lineNum, collumn, lineText := p.lexer.ErrorLine(p.currentToken().Pos)
	p.errors = append(p.errors, fmt.Sprintf(`[%d:%d] %s`, lineNum, collumn-1, lineText))
	p.errors = append(p.errors, fmt.Sprintf(`%s^`, strings.Repeat(" ", collumn+len(fmt.Sprint(lineNum))+len(fmt.Sprint(collumn))+2)))
	p.errors = append(p.errors, err)

	p.printAllErrors()
	os.Exit(99)
}

func (p *parser) printAllErrors() {
	if len(p.errors) > 0 {
		for _, err := range p.errors {
			c := color.New(color.FgRed).Add(color.Bold)
			c.Println(err)
		}
	}
}
