package promqlbuilder

import (
	"fmt"

	"github.com/perses/promql-builder/matrix"
	"github.com/prometheus/prometheus/model/labels"
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

func DeepCopyExpr(expr parser.Expr) parser.Expr {
	if expr == nil {
		return nil
	}

	switch e := expr.(type) {
	case *parser.VectorSelector:
		copy := &parser.VectorSelector{
			Name:                    e.Name,
			OriginalOffset:          e.OriginalOffset,
			Offset:                  e.Offset,
			Timestamp:               e.Timestamp,
			SkipHistogramBuckets:    e.SkipHistogramBuckets,
			StartOrEnd:              e.StartOrEnd,
			UnexpandedSeriesSet:     e.UnexpandedSeriesSet,
			Series:                  e.Series,
			BypassEmptyMatcherCheck: e.BypassEmptyMatcherCheck,
			PosRange:                e.PosRange,
		}
		copy.LabelMatchers = make([]*labels.Matcher, len(e.LabelMatchers))
		for i, m := range e.LabelMatchers {
			mCopy := *m
			copy.LabelMatchers[i] = &mCopy
		}
		return copy

	case *parser.MatrixSelector:
		return &parser.MatrixSelector{
			VectorSelector: DeepCopyExpr(e.VectorSelector).(*parser.VectorSelector),
			Range:          e.Range,
			EndPos:         e.EndPos,
		}

	case *matrix.Builder:
		return &matrix.Builder{
			Expr: DeepCopyExpr(e.Expr),
			InternalMatrix: &parser.MatrixSelector{
				VectorSelector: DeepCopyExpr(e.InternalMatrix.VectorSelector).(*parser.VectorSelector),
				Range:          e.InternalMatrix.Range,
				EndPos:         e.InternalMatrix.EndPos,
			},
			RangeAsVariable: e.RangeAsVariable,
		}

	case *parser.AggregateExpr:
		return &parser.AggregateExpr{
			Op:       e.Op,
			Expr:     DeepCopyExpr(e.Expr),
			Param:    DeepCopyExpr(e.Param),
			Grouping: e.Grouping,
			Without:  e.Without,
			PosRange: e.PosRange,
		}

	case *AggregationBuilder:
		return &AggregationBuilder{
			Expr: DeepCopyExpr(e.Expr),
			internal: &parser.AggregateExpr{
				Op:       e.internal.Op,
				Expr:     DeepCopyExpr(e.internal.Expr),
				Param:    DeepCopyExpr(e.internal.Param),
				Grouping: e.internal.Grouping,
				Without:  e.internal.Without,
				PosRange: e.internal.PosRange,
			},
		}

	case *parser.BinaryExpr:
		return &parser.BinaryExpr{
			Op:             e.Op,
			LHS:            DeepCopyExpr(e.LHS),
			RHS:            DeepCopyExpr(e.RHS),
			VectorMatching: e.VectorMatching,
			ReturnBool:     e.ReturnBool,
		}

	case *BinaryBuilder:
		return &BinaryBuilder{
			internal: &parser.BinaryExpr{
				Op:             e.internal.Op,
				LHS:            DeepCopyExpr(e.internal.LHS),
				RHS:            DeepCopyExpr(e.internal.RHS),
				VectorMatching: e.internal.VectorMatching,
				ReturnBool:     e.internal.ReturnBool,
			},
		}

	case *BinaryWithVectorMatching:
		return &BinaryWithVectorMatching{
			binaryOpt: &BinaryBuilder{
				internal: &parser.BinaryExpr{
					Op:             e.binaryOpt.internal.Op,
					LHS:            DeepCopyExpr(e.binaryOpt.internal.LHS),
					RHS:            DeepCopyExpr(e.binaryOpt.internal.RHS),
					VectorMatching: e.binaryOpt.internal.VectorMatching,
					ReturnBool:     e.binaryOpt.internal.ReturnBool,
				},
			},
		}

	case *parser.Call:
		copy := &parser.Call{
			Func: e.Func,
		}
		copy.Args = make([]parser.Expr, len(e.Args))
		for i, arg := range e.Args {
			copy.Args[i] = DeepCopyExpr(arg)
		}
		return copy

	case *parser.NumberLiteral:
		return &parser.NumberLiteral{
			Val:      e.Val,
			PosRange: e.PosRange,
		}

	case *parser.StringLiteral:
		return &parser.StringLiteral{
			Val:      e.Val,
			PosRange: e.PosRange,
		}

	case *parser.SubqueryExpr:
		return &parser.SubqueryExpr{
			Expr:           DeepCopyExpr(e.Expr),
			Range:          e.Range,
			OriginalOffset: e.OriginalOffset,
			Offset:         e.Offset,
			Timestamp:      e.Timestamp,
			StartOrEnd:     e.StartOrEnd,
			Step:           e.Step,
			EndPos:         e.EndPos,
		}

	case *parser.ParenExpr:
		return &parser.ParenExpr{
			Expr:     DeepCopyExpr(e.Expr),
			PosRange: e.PosRange,
		}

	case *parser.UnaryExpr:
		return &parser.UnaryExpr{
			Op:       e.Op,
			Expr:     DeepCopyExpr(e.Expr),
			StartPos: e.StartPos,
		}

	case *parser.StepInvariantExpr:
		return &parser.StepInvariantExpr{
			Expr: DeepCopyExpr(e.Expr),
		}

	default:
		panic("unsupported expr type in DeepCopyExpr" + fmt.Sprintf("%T", e))
	}
}
