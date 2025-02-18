package promqlbuilder

import (
	"testing"

	"github.com/perses/promql-builder/label"
	"github.com/perses/promql-builder/matrix"
	"github.com/perses/promql-builder/vector"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/stretchr/testify/assert"
)

func TestPromQLBuilder(t *testing.T) {
	testSuite := []struct {
		name     string
		expected string
		expr     parser.Expr
	}{
		{
			name:     "simple instant vector",
			expected: "foo",
			expr:     vector.New(vector.WithMetricName("foo")),
		},
		{
			name:     "instant vector with label",
			expected: `foo{namespace="monitoring",podName=~"prom-.+"}`,
			expr: vector.New(
				vector.WithMetricName("foo"),
				vector.WithLabelMatchers(
					label.New("namespace").Equal("monitoring"),
					label.New("podName").EqualRegexp("prom-.+"),
				),
			),
		},
		{
			name:     "range vector",
			expected: "foo[5d]",
			expr: matrix.New(
				vector.New(vector.WithMetricName("foo")),
				matrix.WithRangeAsString("5d"),
			),
		},
		{
			name:     "range vector with variable",
			expected: "foo[$__rate_interval]",
			expr: matrix.New(
				vector.New(vector.WithMetricName("foo")),
				matrix.WithRangeAsVariable("$__rate_interval"),
			),
		},
		{
			name:     "rate function",
			expected: "rate(foo[5d])",
			expr: Rate(
				matrix.New(
					vector.New(vector.WithMetricName("foo")),
					matrix.WithRangeAsString("5d"),
				),
			),
		},
		{
			name:     "binary opt",
			expected: "vector(234) + vector(123)",
			expr: Add(
				Vector(234),
				Vector(123),
			),
		},
		{
			name:     "binary opt with ignoring and group left",
			expected: "sum by (namespace) (rate(foo[$__rate_interval])) - ignoring (podName) group_left (namespace) perses_info",
			expr: Sub(
				Sum(
					Rate(
						matrix.New(
							vector.New(vector.WithMetricName("foo")),
							matrix.WithRangeAsVariable("$__rate_interval"),
						),
					),
				).By("namespace"),
				vector.New(vector.WithMetricName("perses_info")),
			).Ignoring("podName").GroupLeft("namespace"),
		},
	}
	for _, test := range testSuite {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.expr.String())
		})
	}
}
