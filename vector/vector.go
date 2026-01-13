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

package vector

import (
	"time"

	"github.com/perses/promql-builder/duration"
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
