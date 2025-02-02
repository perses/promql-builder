package aggregation

import "github.com/prometheus/prometheus/promql/parser"

type Builder parser.AggregateExpr

type Option func(b *Builder)

func create(aggregateOp parser.ItemType, vector parser.Expr, options []Option) *parser.AggregateExpr {
	b := &Builder{}
	for _, opt := range options {
		opt(b)
	}
	b.Expr = vector
	b.Op = aggregateOp
	return (*parser.AggregateExpr)(b)
}

func By(labels ...string) Option {
	return func(b *Builder) {
		b.Grouping = labels
		b.Without = false
	}
}

func Without(labels ...string) Option {
	return func(b *Builder) {
		b.Grouping = labels
		b.Without = true
	}
}

func Avg(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.AVG, vector, options)
}

func Bottomk(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.BOTTOMK, vector, options)
}

func Count(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.COUNT, vector, options)
}

func CountValues(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.COUNT_VALUES, vector, options)
}

func Group(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.GROUP, vector, options)
}

func Max(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.MAX, vector, options)
}

func Min(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.MIN, vector, options)
}

func Quantile(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.QUANTILE, vector, options)
}

func LimitK(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.LIMITK, vector, options)
}

func LimitRatio(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.LIMIT_RATIO, vector, options)
}

func Stddev(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.STDDEV, vector, options)
}

func Stdvar(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.STDVAR, vector, options)
}

func Sum(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.SUM, vector, options)
}

func Topk(vector parser.Expr, options ...Option) *parser.AggregateExpr {
	return create(parser.TOPK, vector, options)
}
