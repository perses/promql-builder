PromQL Builder
==============

A library to build PromQL expression fully in Golang.

## Usage

### Create an instant vector

To handle this usecase, we are providing the package `vector` that proposes various options to create your instant
vector.

For example:

```go
package main

import (
	"fmt"

	"github.com/perses/promql-builder/label"
	"github.com/perses/promql-builder/vector"
)

func main() {
	v1 := vector.New(
		vector.WithMetricName("foo"),
		vector.WithLabelMatchers(
			label.New("namespace").Equal("monitoring"),
			label.New("pod-name").EqualRegexp("prom-.+"),
		),
	)
	fmt.Print(v1.String())
}
```

It will give the following output:

```text
foo{namespace="monitoring",pod-name=~"prom-.+"}
```

### Create a range vector

To handle this usecase, we are providing the package `matrix` that proposes various options to create your range vector.

For example:

```go
package main

import (
	"fmt"

	"github.com/perses/promql-builder/label"
	"github.com/perses/promql-builder/matrix"
	"github.com/perses/promql-builder/vector"
)

func main() {
	v1 := vector.New(
		vector.WithMetricName("foo"),
		vector.WithLabelMatchers(
			label.New("namespace").Equal("monitoring"),
			label.New("pod-name").EqualRegexp("prom-.+"),
		),
	)
	m := matrix.New(v1, matrix.WithRangeAsVariable("$__rate_interval"))
	fmt.Print(m.String())
}
```

It will give the following output:

```text
foo{namespace="monitoring",pod-name=~"prom-.+"}[$__rate_interval]
```

Note that as a duration we are using a variable.
This is useful when you are using this library in a Dashboard As Code context,
because likely the range duration will use the building variable coming with the dashboard tools like Grafana or Perses.

Of course, you can define a proper range duration using the option `WithRangeAsString`:

```go
package main

import (
	"fmt"

	"github.com/perses/promql-builder/label"
	"github.com/perses/promql-builder/matrix"
	"github.com/perses/promql-builder/vector"
)

func main() {
	v1 := vector.New(
		vector.WithMetricName("foo"),
		vector.WithLabelMatchers(
			label.New("namespace").Equal("monitoring"),
			label.New("pod-name").EqualRegexp("prom-.+"),
		),
	)
	m := matrix.New(v1, matrix.WithRangeAsString("1h2m4s"))
	fmt.Print(m.String())
}
```

It will give the following output:

```text
foo{namespace="monitoring",pod-name=~"prom-.+"}[1h2m4s]
```
