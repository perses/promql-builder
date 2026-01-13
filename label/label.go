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

type Builder labels.Matcher

func New(labelName string) *Builder {
	return &Builder{
		Name: labelName,
	}
}

func (b *Builder) Equal(labelValue string) *labels.Matcher {
	b.Type = labels.MatchEqual
	b.Value = labelValue
	return (*labels.Matcher)(b)
}

func (b *Builder) EqualRegexp(labelValue string) *labels.Matcher {
	b.Type = labels.MatchRegexp
	b.Value = labelValue
	return (*labels.Matcher)(b)
}

func (b *Builder) NotEqual(labelValue string) *labels.Matcher {
	b.Type = labels.MatchNotEqual
	b.Value = labelValue
	return (*labels.Matcher)(b)
}

func (b *Builder) NotEqualRegexp(labelValue string) *labels.Matcher {
	b.Type = labels.MatchNotRegexp
	b.Value = labelValue
	return (*labels.Matcher)(b)
}
