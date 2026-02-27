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
	"fmt"
	"testing"
	"time"

	"github.com/perses/promql-builder/matrix"
	"github.com/perses/promql-builder/subquery"
	"github.com/perses/promql-builder/vector"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/stretchr/testify/assert"
)

func TestInspect(t *testing.T) {
	testSuite := []struct {
		name      string
		expr      parser.Expr
		wantNodes []string
	}{
		{
			name:      "vector selector",
			expr:      vector.New(vector.WithMetricName("foo")),
			wantNodes: []string{"*parser.VectorSelector: foo"},
		},
		{
			name:      "number literal",
			expr:      NewNumber(123),
			wantNodes: []string{"*parser.NumberLiteral: 123"},
		},
		{
			name: "matrix builder with variable",
			expr: matrix.New(
				vector.New(vector.WithMetricName("foo")),
				matrix.WithRangeAsVariable("$__rate_interval"),
			),
			wantNodes: []string{
				"*matrix.Builder: foo[$__rate_interval]",
				"*parser.VectorSelector: foo",
			},
		},
		{
			name: "parenthesis",
			expr: Parenthesis(vector.New(vector.WithMetricName("foo"))),
			wantNodes: []string{
				"*parser.ParenExpr: (foo)",
				"*parser.VectorSelector: foo",
			},
		},
		{
			name: "unary negation",
			expr: &parser.UnaryExpr{
				Op:   parser.SUB,
				Expr: vector.New(vector.WithMetricName("foo")),
			},
			wantNodes: []string{
				"*parser.UnaryExpr: -foo",
				"*parser.VectorSelector: foo",
			},
		},
		{
			name: "subquery",
			expr: subquery.New(
				subquery.WithExpr(Sum(vector.New(vector.WithMetricName("foo")))),
				subquery.WithRange(5*time.Minute),
			),
			wantNodes: []string{
				"*parser.SubqueryExpr: sum(foo)[5m:]",
				"*promqlbuilder.AggregationBuilder: sum(foo)",
				"*parser.VectorSelector: foo",
			},
		},
		{
			name: "rate function",
			expr: Rate(
				matrix.New(
					vector.New(vector.WithMetricName("foo")),
					matrix.WithRangeAsVariable("$__rate_interval"),
				),
			),
			wantNodes: []string{
				"*parser.Call: rate(foo[$__rate_interval])",
				"*matrix.Builder: foo[$__rate_interval]",
				"*parser.VectorSelector: foo",
			},
		},
		{
			name: "sum aggregation",
			expr: Sum(vector.New(vector.WithMetricName("foo"))),
			wantNodes: []string{
				"*promqlbuilder.AggregationBuilder: sum(foo)",
				"*parser.VectorSelector: foo",
			},
		},
		{
			name: "topk aggregation with param",
			expr: TopK(vector.New(vector.WithMetricName("foo")), 5),
			wantNodes: []string{
				"*promqlbuilder.AggregationBuilder: topk(5, foo)",
				"*parser.VectorSelector: foo",
				"*parser.NumberLiteral: 5",
			},
		},
		{
			name: "binary add",
			expr: Add(
				vector.New(vector.WithMetricName("foo")),
				vector.New(vector.WithMetricName("bar")),
			),
			wantNodes: []string{
				"*promqlbuilder.BinaryBuilder: foo + bar",
				"*parser.VectorSelector: foo",
				"*parser.VectorSelector: bar",
			},
		},
		{
			name: "binary with vector matching",
			expr: Div(
				Sum(vector.New(vector.WithMetricName("foo"))),
				vector.New(vector.WithMetricName("bar")),
			).Ignoring("label").GroupLeft(),
			wantNodes: []string{
				"*promqlbuilder.BinaryWithVectorMatching: sum(foo) / ignoring (label) group_left () bar",
				"*promqlbuilder.AggregationBuilder: sum(foo)",
				"*parser.VectorSelector: foo",
				"*parser.VectorSelector: bar",
			},
		},
		{
			name: "binary with fill LHS",
			expr: Div(
				Sum(vector.New(vector.WithMetricName("foo"))),
				vector.New(vector.WithMetricName("bar")),
			).On("label").FillLHS(0),
			wantNodes: []string{
				"*promqlbuilder.BinaryWithVectorMatching: sum(foo) / on (label) fill_left (0) bar",
				"*promqlbuilder.AggregationBuilder: sum(foo)",
				"*parser.VectorSelector: foo",
				"*parser.VectorSelector: bar",
			},
		},
		{
			name: "binary with fill RHS",
			expr: Div(
				Sum(vector.New(vector.WithMetricName("foo"))),
				vector.New(vector.WithMetricName("bar")),
			).On("label").FillRHS(0),
			wantNodes: []string{
				"*promqlbuilder.BinaryWithVectorMatching: sum(foo) / on (label) fill_right (0) bar",
				"*promqlbuilder.AggregationBuilder: sum(foo)",
				"*parser.VectorSelector: foo",
				"*parser.VectorSelector: bar",
			},
		},
	}

	for _, test := range testSuite {
		t.Run(test.name, func(t *testing.T) {
			var nodes []string
			Inspect(test.expr, func(node parser.Node, _ []parser.Node) error {
				if node == nil {
					return nil
				}
				nodes = append(nodes, fmt.Sprintf("%T: %s", node, node.String()))
				return nil
			})
			assert.Equal(t, test.wantNodes, nodes)
		})
	}
}
