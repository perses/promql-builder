package promqlbuilder

import "github.com/prometheus/prometheus/promql/parser"

// Parenthesis wraps an expression in parentheses.
func Parenthesis(expr parser.Expr) *parser.ParenExpr {
	return &parser.ParenExpr{
		Expr: expr,
	}
}
