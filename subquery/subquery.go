package subquery

import (
	"time"

	"github.com/perses/promql-builder/duration"
	"github.com/prometheus/prometheus/promql/parser"
)

type Builder parser.SubqueryExpr

type Option func(subquery *Builder)

func New(options ...Option) *parser.SubqueryExpr {
	b := &Builder{}
	for _, opt := range options {
		opt(b)
	}
	return (*parser.SubqueryExpr)(b)
}

func WithExpr(expr parser.Expr) Option {
	return func(subquery *Builder) {
		subquery.Expr = expr
	}
}

func WithRangeAsString(d string) Option {
	return func(subquery *Builder) {
		subquery.Range = time.Duration(duration.MustParse(d))
	}
}

func WithRange(duration time.Duration) Option {
	return func(subquery *Builder) {
		subquery.Range = duration
	}
}

func WithOffset(duration time.Duration) Option {
	return func(vector *Builder) {
		vector.Offset = duration
	}
}

func WithOffsetAsString(d string) Option {
	return func(vector *Builder) {
		vector.Offset = time.Duration(duration.MustParse(d))
	}
}

func WithAtStart() Option {
	return func(vector *Builder) {
		vector.StartOrEnd = parser.START
	}
}

func WithAtEnd() Option {
	return func(vector *Builder) {
		vector.StartOrEnd = parser.END
	}
}

func WithAtSpecificTimeStamp(timestamp int64) Option {
	return func(vector *Builder) {
		vector.Timestamp = &timestamp
	}
}
