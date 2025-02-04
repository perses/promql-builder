package promqlbuilder

import (
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/prometheus/prometheus/promql/parser/posrange"
)

type AggregationBuilder struct {
	parser.Expr
	internal *parser.AggregateExpr
}

func (a *AggregationBuilder) Type() parser.ValueType {
	return a.internal.Type()
}
func (a *AggregationBuilder) PromQLExpr() {
	a.internal.PromQLExpr()
}
func (a *AggregationBuilder) String() string {
	return a.internal.String()
}
func (a *AggregationBuilder) Pretty(level int) string {
	return a.internal.Pretty(level)
}
func (a *AggregationBuilder) PositionRange() posrange.PositionRange {
	return a.internal.PositionRange()
}

func (a *AggregationBuilder) By(labels ...string) *AggregationBuilder {
	a.internal.Without = false
	a.internal.Grouping = labels
	return a
}

func (a *AggregationBuilder) Without(labels ...string) *AggregationBuilder {
	a.internal.Without = true
	a.internal.Grouping = labels
	return a
}

func create(aggregateOp parser.ItemType, vector parser.Expr) *AggregationBuilder {
	b := &AggregationBuilder{
		internal: &parser.AggregateExpr{},
	}
	b.internal.Expr = vector
	b.internal.Op = aggregateOp
	return b
}

func Avg(vector parser.Expr) *AggregationBuilder {
	return create(parser.AVG, vector)
}

func Bottomk(vector parser.Expr) *AggregationBuilder {
	return create(parser.BOTTOMK, vector)
}

func Count(vector parser.Expr) *AggregationBuilder {
	return create(parser.COUNT, vector)
}

func CountValues(vector parser.Expr) *AggregationBuilder {
	return create(parser.COUNT_VALUES, vector)
}

func Group(vector parser.Expr) *AggregationBuilder {
	return create(parser.GROUP, vector)
}

func Max(vector parser.Expr) *AggregationBuilder {
	return create(parser.MAX, vector)
}

func Min(vector parser.Expr) *AggregationBuilder {
	return create(parser.MIN, vector)
}

func Quantile(vector parser.Expr) *AggregationBuilder {
	return create(parser.QUANTILE, vector)
}

func LimitK(vector parser.Expr) *AggregationBuilder {
	return create(parser.LIMITK, vector)
}

func LimitRatio(vector parser.Expr) *AggregationBuilder {
	return create(parser.LIMIT_RATIO, vector)
}

func Stddev(vector parser.Expr) *AggregationBuilder {
	return create(parser.STDDEV, vector)
}

func Stdvar(vector parser.Expr) *AggregationBuilder {
	return create(parser.STDVAR, vector)
}

func Sum(vector parser.Expr) *AggregationBuilder {
	return create(parser.SUM, vector)
}

func Topk(vector parser.Expr) *AggregationBuilder {
	return create(parser.TOPK, vector)
}
