package matrix

import (
	"time"

	"github.com/prometheus/prometheus/promql/parser"
)

func New(vector *parser.VectorSelector, duration time.Duration) *parser.MatrixSelector {
	return &parser.MatrixSelector{
		VectorSelector: vector,
		Range:          duration,
	}
}
