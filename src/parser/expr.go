package parser

import (
	"fmt"
	"strconv"

	"terascript/src/ast"
	"terascript/src/lexer"
)

func parse_expr(p *parser, bp binding_power) ast.Expr {

	tokenKind := p.currentToken().Kind
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		p.Error(fmt.Sprintf("INVALID PLACEMENT OF %s", p.currentToken().Kind))
	}

	left := nud_fn(p)

	for bp_lu[p.currentToken().Kind] > bp {
		tokenKind = p.currentToken().Kind
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			p.Error(fmt.Sprintf("LED HANDLER EXPECTED FOR TOKEN %s", p.currentToken().Kind))
		}

		left = led_fn(p, left, bp_lu[p.currentToken().Kind])
	}

	return left
}

func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentToken().Kind {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{
			Value: number,
		}
	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{
			Value: p.advance().Value,
		}
	default:
		p.Error(fmt.Sprintf("Cannot create expression from %s", p.currentToken().Kind))
		panic("just to keep the compiler happy")
	}
}

func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operator := p.advance()
	right := parse_expr(p, bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}

}

func parse_prefix_expr(p *parser) ast.Expr {
	operator := p.advance()
	right := parse_expr(p, default_bp)

	return ast.PrefixExpr{
		Operator: operator,
		Right:    right,
	}
}

func parse_assignment_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operator := p.advance()
	right := parse_expr(p, bp)

	return ast.AssignmentExpr{
		Assigne:  left,
		Operator: operator,
		Right:    right,
	}

}

func parse_grouping_expr(p *parser) ast.Expr {
	p.advance()
	expression := parse_expr(p, default_bp)
	p.expectError(lexer.CLOSE_PAREN, "Expected closing parenthesis")

	return expression
}
