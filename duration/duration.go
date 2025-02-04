package duration

import "github.com/prometheus/common/model"

func MustParse(s string) model.Duration {
	d, err := model.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}
