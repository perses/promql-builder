// Copyright The Perses Authors
// Licensed under the Apache License, Version 2.0 (the \"License\");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an \"AS IS\" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func createWithParam(aggregateOp parser.ItemType, vector parser.Expr, param parser.Expr) *AggregationBuilder {
	b := &AggregationBuilder{
		internal: &parser.AggregateExpr{},
	}
	b.internal.Expr = vector
	b.internal.Op = aggregateOp
	b.internal.Param = param
	return b
}

func Avg(vector parser.Expr) *AggregationBuilder {
	return create(parser.AVG, vector)
}

func BottomK(vector parser.Expr, k float64) *AggregationBuilder {
	return createWithParam(parser.BOTTOMK, vector, NewNumber(k))
}

func Count(vector parser.Expr) *AggregationBuilder {
	return create(parser.COUNT, vector)
}

func CountValues(label string, vector parser.Expr) *AggregationBuilder {
	return createWithParam(parser.COUNT_VALUES, vector, NewString(label))
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

func Quantile(vector parser.Expr, quantile float64) *AggregationBuilder {
	return createWithParam(parser.QUANTILE, vector, NewNumber(quantile))
}

func LimitK(vector parser.Expr, k float64) *AggregationBuilder {
	return createWithParam(parser.LIMITK, vector, NewNumber(k))
}

func LimitRatio(vector parser.Expr, ratio float64) *AggregationBuilder {
	return createWithParam(parser.LIMIT_RATIO, vector, NewNumber(ratio))
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

func TopK(vector parser.Expr, k float64) *AggregationBuilder {
	return createWithParam(parser.TOPK, vector, NewNumber(k))
}
