package promqlbuilder

import (
	"testing"
	"time"

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
		{
			name:     "custom function",
			expected: "xincrease(metric[1m])",
			expr: NewFunction("xincrease",
				matrix.New(
					vector.New(vector.WithMetricName("metric")),
					matrix.WithRange(time.Minute),
				),
			),
		},
		{
			name:     "count values with label",
			expected: "count_values(\"config_hash\", alertmanager_config_hash)",
			expr:     CountValues("config_hash", vector.New(vector.WithMetricName("alertmanager_config_hash"))),
		},
		{
			name:     "parenthesis",
			expected: "(time() - foo[5d]) / 100",
			expr: Div(
				Parenthesis(
					Sub(
						Time(),
						matrix.New(
							vector.New(
								vector.WithMetricName("foo")),
							matrix.WithRangeAsString("5d"),
						)),
				),
				NewNumber(100),
			),
		},
	}
	for _, test := range testSuite {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.expr.String())
		})
	}
}

func TestPretty(t *testing.T) {
	testSuite := []struct {
		name     string
		expected string
		expr     parser.Expr
	}{
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
			name: "long promql expression with range vector and binary op",
			expected: `    sum by (namespace, job, code) (
      rate(
        http_requests_total{code=~"5..",handler="query",job=~"thanos-query-example-query",namespace="thanos-operator-system"}[5m]
      )
    )
  / ignoring (code) group_left ()
    sum by (namespace, job) (
      rate(
        http_requests_total{handler="query",job=~"thanos-query-example-query",namespace="thanos-operator-system"}[5m]
      )
    )
*
  100`,
			expr: Mul(
				Div(
					Sum(
						Rate(
							matrix.New(
								vector.New(
									vector.WithMetricName("http_requests_total"),
									vector.WithLabelMatchers(
										label.New("code").EqualRegexp("5.."),
										label.New("handler").Equal("query"),
										label.New("job").EqualRegexp("thanos-query-example-query"),
										label.New("namespace").Equal("thanos-operator-system"),
									),
								),
								matrix.WithRangeAsString("5m"),
							),
						),
					).By("namespace", "job", "code"),
					Sum(
						Rate(
							matrix.New(
								vector.New(
									vector.WithMetricName("http_requests_total"),
									vector.WithLabelMatchers(
										label.New("handler").Equal("query"),
										label.New("job").EqualRegexp("thanos-query-example-query"),
										label.New("namespace").Equal("thanos-operator-system"),
									),
								),
								matrix.WithRangeAsString("5m"),
							),
						),
					).By("namespace", "job"),
				).Ignoring("code").GroupLeft(),
				&parser.NumberLiteral{Val: 100},
			),
		},
		{
			name: "long promql expression with range vector and binary op and matrix variable",
			expected: `    sum by (namespace, job, code) (
      rate(
        http_requests_total{code=~"5..",handler="query",job=~"thanos-query-example-query",namespace="thanos-operator-system"}[$__rate_interval]
      )
    )
  / ignoring (code) group_left ()
    sum by (namespace, job) (
      rate(
        http_requests_total{handler="query",job=~"thanos-query-example-query",namespace="thanos-operator-system"}[$__rate_interval]
      )
    )
*
  100`,
			expr: Mul(
				Div(
					Sum(
						Rate(
							matrix.New(
								vector.New(
									vector.WithMetricName("http_requests_total"),
									vector.WithLabelMatchers(
										label.New("code").EqualRegexp("5.."),
										label.New("handler").Equal("query"),
										label.New("job").EqualRegexp("thanos-query-example-query"),
										label.New("namespace").Equal("thanos-operator-system"),
									),
								),
								matrix.WithRangeAsVariable("$__rate_interval"),
							),
						),
					).By("namespace", "job", "code"),
					Sum(
						Rate(
							matrix.New(
								vector.New(
									vector.WithMetricName("http_requests_total"),
									vector.WithLabelMatchers(
										label.New("handler").Equal("query"),
										label.New("job").EqualRegexp("thanos-query-example-query"),
										label.New("namespace").Equal("thanos-operator-system"),
									),
								),
								matrix.WithRangeAsVariable("$__rate_interval"),
							),
						),
					).By("namespace", "job"),
				).Ignoring("code").GroupLeft(),
				&parser.NumberLiteral{Val: 100},
			),
		},
		{
			name: "parenthesis with range vector",
			expected: `  (
      time()
    -
      thanos_objstore_last_successful_upload_time{job=~"thanos-compactor",namespace="thanos-operator-system"}[5d]
  )
/
  100`,
			expr: Div(
				Parenthesis(
					Sub(
						Time(),
						matrix.New(
							vector.New(
								vector.WithMetricName("thanos_objstore_last_successful_upload_time"),
								vector.WithLabelMatchers(
									label.New("job").EqualRegexp("thanos-compactor"),
									label.New("namespace").Equal("thanos-operator-system"),
								),
							),
							matrix.WithRangeAsString("5d"),
						),
					),
				),
				NewNumber(100),
			),
		},
	}
	for _, test := range testSuite {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.expr.Pretty(0))
		})
	}
}
