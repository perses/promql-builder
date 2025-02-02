package vector

import (
	"time"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/promql/parser"
)

type Option func(vector *Builder)

type Builder parser.VectorSelector

func New(options ...Option) *parser.VectorSelector {
	b := &Builder{}
	for _, opt := range options {
		opt(b)
	}
	return (*parser.VectorSelector)(b)
}

func WithMetricName(name string) Option {
	return func(vector *Builder) {
		vector.Name = name
	}
}

func WithLabelMatchers(matchers ...*labels.Matcher) Option {
	return func(vector *Builder) {
		vector.LabelMatchers = matchers
	}
}

func Offset(duration time.Duration) Option {
	return func(vector *Builder) {
		vector.Offset = duration
	}
}

func AtStart() Option {
	return func(vector *Builder) {
		vector.StartOrEnd = parser.START
	}
}

func AtEnd() Option {
	return func(vector *Builder) {
		vector.StartOrEnd = parser.END
	}
}

func AtSpecificTimeStamp(timestamp int64) Option {
	return func(vector *Builder) {
		vector.Timestamp = &timestamp
	}
}
