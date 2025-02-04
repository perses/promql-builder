package matrix

import (
	"fmt"
	"time"

	"github.com/perses/promql-builder/duration"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/prometheus/prometheus/promql/parser/posrange"
)

type Builder struct {
	parser.Expr
	internalMatrix  *parser.MatrixSelector
	rangeAsVariable string
}

// Type returns the type the expression evaluates to. It does not perform
// in-depth checks as this is done at parsing-time.
func (b *Builder) Type() parser.ValueType {
	return b.internalMatrix.Type()
}

// PromQLExpr ensures that no other types accidentally implement the interface.
func (b *Builder) PromQLExpr() {
	b.internalMatrix.PromQLExpr()
}

// String representation of the node that returns the given node when parsed
// as part of a valid query.
func (b *Builder) String() string {
	at, offset := b.atOffset()
	// Copy the Vector selector before changing the offset
	vecSelector := *b.internalMatrix.VectorSelector.(*parser.VectorSelector)
	// Do not print the @ and offset twice.
	offsetVal, atVal, preproc := vecSelector.OriginalOffset, vecSelector.Timestamp, vecSelector.StartOrEnd
	vecSelector.OriginalOffset = 0
	vecSelector.Timestamp = nil
	vecSelector.StartOrEnd = 0

	rangeAsString := ""
	if len(b.rangeAsVariable) > 0 {
		rangeAsString = b.rangeAsVariable
	} else {
		rangeAsString = model.Duration(b.internalMatrix.Range).String()
	}

	str := fmt.Sprintf("%s[%s]%s%s", vecSelector.String(), rangeAsString, at, offset)

	vecSelector.OriginalOffset, vecSelector.Timestamp, vecSelector.StartOrEnd = offsetVal, atVal, preproc

	return str
}

func (b *Builder) atOffset() (string, string) {
	// Copy the Vector selector before changing the offset
	vecSelector := b.internalMatrix.VectorSelector.(*parser.VectorSelector)
	offset := ""
	switch {
	case vecSelector.OriginalOffset > time.Duration(0):
		offset = fmt.Sprintf(" offset %s", model.Duration(vecSelector.OriginalOffset))
	case vecSelector.OriginalOffset < time.Duration(0):
		offset = fmt.Sprintf(" offset -%s", model.Duration(-vecSelector.OriginalOffset))
	}
	at := ""
	switch {
	case vecSelector.Timestamp != nil:
		at = fmt.Sprintf(" @ %.3f", float64(*vecSelector.Timestamp)/1000.0)
	case vecSelector.StartOrEnd == parser.START:
		at = " @ start()"
	case vecSelector.StartOrEnd == parser.END:
		at = " @ end()"
	}
	return at, offset
}

func (b *Builder) Pretty(level int) string {
	return b.internalMatrix.Pretty(level)
}

func (b *Builder) PositionRange() posrange.PositionRange {
	return b.internalMatrix.PositionRange()
}

type Option func(matrix *Builder)

func New(options ...Option) *Builder {
	b := &Builder{
		internalMatrix: &parser.MatrixSelector{},
	}
	for _, opt := range options {
		opt(b)
	}
	return b
}

func WithVectorSelector(v *parser.VectorSelector) Option {
	return func(matrix *Builder) {
		matrix.internalMatrix.VectorSelector = v
	}
}

func WithRange(d time.Duration) Option {
	return func(matrix *Builder) {
		matrix.internalMatrix.Range = d
	}
}

// WithRangeAsString sets the range as a string like "3h2m1s".
func WithRangeAsString(d string) Option {
	return func(matrix *Builder) {
		matrix.internalMatrix.Range = time.Duration(duration.MustParse(d))
	}
}

// WithRangeAsVariable sets the range as a variable name.
// It will be useful in case you are writing a query for a dashboard and the range is a variable like "$__rate_interval".
// Use it if your range is not correct in terms of PromQL syntax.
func WithRangeAsVariable(name string) Option {
	return func(matrix *Builder) {
		matrix.rangeAsVariable = name
	}
}
