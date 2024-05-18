package parser

import (
	"terascript/src/ast"
	"terascript/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentToken().Kind]

	if exists {
		return stmt_fn(p)
	}

	return parse_expression_stmt(p)
}

func parse_expression_stmt(p *parser) ast.ExprStmt {
	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMICOLON)

	return ast.ExprStmt{
		Expression: expression,
	}
}

func parse_var_decl_stmt(p *parser) ast.Stmt {
	var explicitType ast.Type
	var assignedValue ast.Expr

	isConstant := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Expected variable name").Value

	if p.currentToken().Kind == lexer.COLON {
		p.advance()
		explicitType = parse_type(p, default_bp)
	}

	if p.currentToken().Kind != lexer.SEMICOLON {
		p.expect(lexer.ASSIGNMENT)
		assignedValue = parse_expr(p, assignment)
	} else if isConstant {
		p.Error("Must assign value to constant declaration")
	}

	p.expect(lexer.SEMICOLON)

	p.sprite.Vars = append(p.sprite.Vars, varName)

	return ast.VarDeclStmt{
		IsConstant:    isConstant,
		VariableName:  varName,
		AssignedValue: assignedValue,
		ExplicitType:  explicitType,
	}
}
