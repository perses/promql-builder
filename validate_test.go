// Copyright The Perses Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package promqlbuilder

import (
	"testing"

	"github.com/perses/promql-builder/duration"
	"github.com/perses/promql-builder/label"
	"github.com/perses/promql-builder/matrix"
	"github.com/perses/promql-builder/vector"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("valid simple expression", func(t *testing.T) {
		expr := vector.New(vector.WithMetricName("up"))
		assert.NoError(t, Validate(expr))
	})

	t.Run("valid rate expression", func(t *testing.T) {
		expr := Rate(
			matrix.New(
				vector.New(vector.WithMetricName("http_requests_total")),
				matrix.WithRangeAsString("5m"),
			),
		)
		assert.NoError(t, Validate(expr))
	})

	t.Run("valid expression with labels", func(t *testing.T) {
		expr := Sum(
			Rate(
				matrix.New(
					vector.New(
						vector.WithMetricName("http_requests_total"),
						vector.WithLabelMatchers(
							label.New("namespace").Equal("monitoring"),
							label.New("job").EqualRegexp("thanos-.+"),
						),
					),
					matrix.WithRangeAsString("5m"),
				),
			),
		).By("namespace", "job")
		assert.NoError(t, Validate(expr))
	})

	t.Run("valid expression with $__rate_interval", func(t *testing.T) {
		expr := Rate(
			matrix.New(
				vector.New(vector.WithMetricName("up")),
				matrix.WithRangeAsVariable("$__rate_interval"),
			),
		)
		assert.NoError(t, Validate(expr))
	})

	t.Run("valid expression with $__interval", func(t *testing.T) {
		expr := Rate(
			matrix.New(
				vector.New(vector.WithMetricName("up")),
				matrix.WithRangeAsVariable("$__interval"),
			),
		)
		assert.NoError(t, Validate(expr))
	})

	t.Run("valid expression with user variable in label", func(t *testing.T) {
		expr := vector.New(
			vector.WithMetricName("up"),
			vector.WithLabelMatchers(
				label.New("namespace").Equal("$namespace"),
			),
		)
		assert.NoError(t, Validate(expr))
	})

	t.Run("valid expression with user variable via ValidateWithVars", func(t *testing.T) {
		expr := vector.New(
			vector.WithMetricName("up"),
			vector.WithLabelMatchers(
				label.New("cluster").Equal("$cluster"),
			),
		)
		assert.NoError(t, ValidateWithVars(expr, []string{"cluster"}))
	})

	t.Run("invalid expression returns error", func(t *testing.T) {
		expr := &parser.BinaryExpr{
			Op:  parser.ADD,
			LHS: vector.New(vector.WithMetricName("foo")),
		}
		err := Validate(expr)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid PromQL expression")
	})

	t.Run("unknown $__ variable in range fails validation", func(t *testing.T) {
		expr := Rate(
			matrix.New(
				vector.New(vector.WithMetricName("up")),
				matrix.WithRangeAsVariable("$__nonexistent"),
			),
		)
		err := ValidateWithVars(expr, nil)
		assert.Error(t, err)
	})

	t.Run("user variable auto-substituted without extraVars", func(t *testing.T) {
		expr := vector.New(
			vector.WithMetricName("up"),
			vector.WithLabelMatchers(
				label.New("ns").Equal("$ns"),
			),
		)
		assert.NoError(t, Validate(expr))
	})
}

func TestLabelRegexValidation(t *testing.T) {
	t.Run("valid regex succeeds", func(t *testing.T) {
		m := label.New("job").EqualRegexp("thanos-.+")
		assert.NotNil(t, m)
	})

	t.Run("invalid regex panics", func(t *testing.T) {
		assert.Panics(t, func() {
			label.New("job").EqualRegexp("[invalid")
		})
	})

	t.Run("bare quantifier panics", func(t *testing.T) {
		assert.Panics(t, func() {
			label.New("job").EqualRegexp("*")
		})
	})

	t.Run("empty regex succeeds", func(t *testing.T) {
		m := label.New("job").EqualRegexp("")
		assert.NotNil(t, m)
	})

	t.Run("valid not-equal regex succeeds", func(t *testing.T) {
		m := label.New("code").NotEqualRegexp("2..")
		assert.NotNil(t, m)
	})

	t.Run("invalid not-equal regex panics", func(t *testing.T) {
		assert.Panics(t, func() {
			label.New("code").NotEqualRegexp("(unclosed")
		})
	})
}

func TestDurationMustParse(t *testing.T) {
	t.Run("valid duration", func(t *testing.T) {
		d := duration.MustParse("5m")
		assert.NotZero(t, d)
	})

	t.Run("invalid duration panics", func(t *testing.T) {
		assert.Panics(t, func() {
			duration.MustParse("not-a-duration")
		})
	})

	t.Run("empty string panics", func(t *testing.T) {
		assert.Panics(t, func() {
			duration.MustParse("")
		})
	})

	t.Run("number without unit panics", func(t *testing.T) {
		assert.Panics(t, func() {
			duration.MustParse("5")
		})
	})
}
