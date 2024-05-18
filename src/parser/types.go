package parser

import (
	"fmt"
	"terascript/src/ast"
	"terascript/src/lexer"
)

type nud_handler_type func(p *parser) ast.Type
type led_handler_type func(p *parser, left ast.Type, bp binding_power) ast.Type

type nud_lookup_type map[lexer.TokenKind]nud_handler_type
type led_lookup_type map[lexer.TokenKind]led_handler_type
type bp_lookup_type map[lexer.TokenKind]binding_power

var bp_lu_type = bp_lookup_type{}
var nud_lu_type = nud_lookup_type{}
var led_lu_type = led_lookup_type{}

func led_type(kind lexer.TokenKind, bp binding_power, led_fn led_handler_type) {
	bp_lu_type[kind] = bp
	led_lu_type[kind] = led_fn
}

func nud_type(kind lexer.TokenKind, nud_fn nud_handler_type) {
	bp_lu_type[kind] = primary
	nud_lu_type[kind] = nud_fn
}

func createTokenTypeLookups() {
	nud_type(lexer.IDENTIFIER, parse_symbol_type)
	nud_type(lexer.OPEN_BRACKET, parse_array_type)
}

func parse_symbol_type(p *parser) ast.Type {
	return ast.SymbolType{
		Name: p.expect(lexer.IDENTIFIER).Value,
	}
}

func parse_array_type(p *parser) ast.Type {
	var underlyingType = parse_type(p, default_bp)
	return ast.ArrayType{
		Underlying: underlyingType,
	}
}

func parse_type(p *parser, bp binding_power) ast.Type {

	tokenKind := p.currentToken().Kind
	nud_fn, exists := nud_lu_type[tokenKind]

	if !exists {
		panic(fmt.Sprintf("NUD HANDLER EXPECTED FOR TOKEN %s\n", p.currentToken().Kind))
	}

	left := nud_fn(p)

	for bp_lu_type[p.currentToken().Kind] > bp {
		tokenKind = p.currentToken().Kind
		led_fn, exists := led_lu_type[tokenKind]

		if !exists {
			panic(fmt.Sprintf("LED HANDLER EXPECTED FOR TOKEN %s\n", p.currentToken().Kind))
		}

		left = led_fn(p, left, bp_lu_type[p.currentToken().Kind])
	}

	return left
}
