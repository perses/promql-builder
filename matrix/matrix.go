package matrix

import (
	"fmt"
	"strings"
	"time"

	"github.com/perses/promql-builder/duration"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/prometheus/prometheus/promql/parser/posrange"
)

type Builder struct {
	parser.Expr
	InternalMatrix  *parser.MatrixSelector
	RangeAsVariable string
}

// Type returns the type the expression evaluates to. It does not perform
// in-depth checks as this is done at parsing-time.
func (b *Builder) Type() parser.ValueType {
	return b.InternalMatrix.Type()
}

// PromQLExpr ensures that no other types accidentally implement the interface.
func (b *Builder) PromQLExpr() {
	b.InternalMatrix.PromQLExpr()
}

// String representation of the node that returns the given node when parsed
// as part of a valid query.
func (b *Builder) String() string {
	at, offset := b.atOffset()
	// Copy the Vector selector before changing the offset
	vecSelector := *b.InternalMatrix.VectorSelector.(*parser.VectorSelector)
	// Do not print the @ and offset twice.
	offsetVal, atVal, preproc := vecSelector.OriginalOffset, vecSelector.Timestamp, vecSelector.StartOrEnd
	vecSelector.OriginalOffset = 0
	vecSelector.Timestamp = nil
	vecSelector.StartOrEnd = 0

	rangeAsString := ""
	if len(b.RangeAsVariable) > 0 {
		rangeAsString = b.RangeAsVariable
	} else {
		rangeAsString = model.Duration(b.InternalMatrix.Range).String()
	}

	str := fmt.Sprintf("%s[%s]%s%s", vecSelector.String(), rangeAsString, at, offset)

	vecSelector.OriginalOffset, vecSelector.Timestamp, vecSelector.StartOrEnd = offsetVal, atVal, preproc

	return str
}

func (b *Builder) atOffset() (string, string) {
	// Copy the Vector selector before changing the offset
	vecSelector := b.InternalMatrix.VectorSelector.(*parser.VectorSelector)
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
	return getCommonPrefixIndent(level, b)
}

func (b *Builder) PositionRange() posrange.PositionRange {
	return b.InternalMatrix.PositionRange()
}

func (b *Builder) Children() []parser.Node {
	return []parser.Node{b.InternalMatrix.VectorSelector}
}

type Option func(matrix *Builder)

func New(v *parser.VectorSelector, options ...Option) *Builder {
	b := &Builder{
		InternalMatrix: &parser.MatrixSelector{
			VectorSelector: v,
		},
	}
	for _, opt := range options {
		opt(b)
	}
	return b
}

func WithRange(d time.Duration) Option {
	return func(matrix *Builder) {
		matrix.InternalMatrix.Range = d
	}
}

// WithRangeAsString sets the range as a string like "3h2m1s".
func WithRangeAsString(d string) Option {
	return func(matrix *Builder) {
		matrix.InternalMatrix.Range = time.Duration(duration.MustParse(d))
	}
}

// WithRangeAsVariable sets the range as a variable name.
// It will be useful in case you are writing a query for a dashboard and the range is a variable like "$__rate_interval".
// Use it if your range is not correct in terms of PromQL syntax.
func WithRangeAsVariable(name string) Option {
	return func(matrix *Builder) {
		matrix.RangeAsVariable = name
	}
}

// Refer to https://github.com/prometheus/prometheus/blob/v3.4.0/promql/parser/prettier.go for below.
// The following is only used for matrix selector.
func getCommonPrefixIndent(level int, current *Builder) string {
	return fmt.Sprintf("%s%s", indent(level), current.String())
}

const indentString = "  "

// indent adds the indentString n number of times.
func indent(n int) string {
	return strings.Repeat(indentString, n)
}
