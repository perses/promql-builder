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
