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

package label

import "github.com/prometheus/prometheus/model/labels"

type Builder struct {
	name string
}

func New(labelName string) *Builder {
	return &Builder{
		name: labelName,
	}
}

func (b *Builder) Equal(labelValue string) *labels.Matcher {
	m, _ := labels.NewMatcher(labels.MatchEqual, b.name, labelValue)
	return m
}

// EqualRegexp creates a regexp matcher. Panics if the pattern is invalid
func (b *Builder) EqualRegexp(labelValue string) *labels.Matcher {
	m, err := labels.NewMatcher(labels.MatchRegexp, b.name, labelValue)
	if err != nil {
		panic(err)
	}
	return m
}

func (b *Builder) NotEqual(labelValue string) *labels.Matcher {
	m, _ := labels.NewMatcher(labels.MatchNotEqual, b.name, labelValue)
	return m
}

// NotEqualRegexp creates a negative regexp matcher. Panics if the pattern is invalid
func (b *Builder) NotEqualRegexp(labelValue string) *labels.Matcher {
	m, err := labels.NewMatcher(labels.MatchNotRegexp, b.name, labelValue)
	if err != nil {
		panic(err)
	}
	return m
}
