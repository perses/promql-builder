package promqlbuilder

import (
	"fmt"

	"github.com/perses/promql-builder/matrix"
	"github.com/prometheus/prometheus/promql/parser"
)

// Walk traverses an AST in depth-first order: It starts by calling
// v.Visit(node, path); node must not be nil. If the visitor w returned by
// v.Visit(node, path) is not nil and the visitor returns no error, Walk is
// invoked recursively with visitor w for each of the non-nil children of node,
// followed by a call of w.Visit(nil), returning an error
// As the tree is descended the path of previous nodes is provided.
func Walk(v parser.Visitor, node parser.Node, path []parser.Node) error {
	var err error
	if v, err = v.Visit(node, path); v == nil || err != nil {
		return err
	}
	path = append(path, node)

	for _, e := range Children(node) {
		if err := Walk(v, e, path); err != nil {
			return err
		}
	}

	_, err = v.Visit(nil, nil)
	return err
}

type inspector func(parser.Node, []parser.Node) error

func (f inspector) Visit(node parser.Node, path []parser.Node) (parser.Visitor, error) {
	if err := f(node, path); err != nil {
		return nil, err
	}

	return f, nil
}

// Inspect traverses an AST in depth-first order: It starts by calling
// f(node, path); node must not be nil. If f returns a nil error, Inspect invokes f
// for all the non-nil children of node, recursively.
func Inspect(node parser.Node, f inspector) {
	Walk(f, node, nil) //nolint:errcheck
}

// Children returns a list of all child nodes of a syntax tree node.
func Children(node parser.Node) []parser.Node {
	// For some reasons these switches have significantly better performance than interfaces
	switch n := node.(type) {
	case *parser.EvalStmt:
		return []parser.Node{n.Expr}
	case parser.Expressions:
		// golang cannot convert slices of interfaces
		ret := make([]parser.Node, len(n))
		for i, e := range n {
			ret[i] = e
		}
		return ret
	case *parser.AggregateExpr:
		// While this does not look nice, it should avoid unnecessary allocations
		// caused by slice resizing
		switch {
		case n.Expr == nil && n.Param == nil:
			return nil
		case n.Expr == nil:
			return []parser.Node{n.Param}
		case n.Param == nil:
			return []parser.Node{n.Expr}
		default:
			return []parser.Node{n.Expr, n.Param}
		}
	case *AggregationBuilder:
		// While this does not look nice, it should avoid unnecessary allocations
		// caused by slice resizing
		switch {
		case n.Expr == nil && n.internal.Param == nil:
			return nil
		case n.Expr == nil:
			return []parser.Node{n.internal.Param}
		case n.internal.Param == nil:
			return []parser.Node{n.Expr}
		default:
			return []parser.Node{n.Expr, n.internal.Param}
		}
	case *parser.BinaryExpr:
		return []parser.Node{n.LHS, n.RHS}
	case *BinaryBuilder:
		return []parser.Node{n.internal.LHS, n.internal.RHS}
	case *parser.Call:
		// golang cannot convert slices of interfaces
		ret := make([]parser.Node, len(n.Args))
		for i, e := range n.Args {
			ret[i] = e
		}
		return ret
	case *parser.SubqueryExpr:
		return []parser.Node{n.Expr}
	case *parser.ParenExpr:
		return []parser.Node{n.Expr}
	case *parser.UnaryExpr:
		return []parser.Node{n.Expr}
	case *parser.MatrixSelector:
		return []parser.Node{n.VectorSelector}
	case *matrix.Builder:
		return n.Children()
	case *parser.StepInvariantExpr:
		return []parser.Node{n.Expr}
	case *parser.NumberLiteral, *parser.StringLiteral, *parser.VectorSelector:
		// nothing to do
		return []parser.Node{}
	default:
		panic(fmt.Errorf("promql.Children: unhandled node type %T", node))
	}
}
