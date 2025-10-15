PromQL Builder
==============

A library to build PromQL expression fully in Golang.

## Usage

### Create an instant vector

To handle this use case, we are providing the package `vector` that proposes various options to create your instant
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

	"github.com/perses/promql-builder/matrix"
	"github.com/perses/promql-builder/vector"
)

func main() {
	m := matrix.New(
		vector.New(vector.WithMetricName("foo")),
		matrix.WithRangeAsString("1h2m4s"),
	)
	fmt.Print(m.String())
}
```

It will give the following output:

```text
foo[1h2m4s]
```

### Use PromQL function

All functions, aggregations and binary operations are available at the root of this package `promqlbuilder`.

For example, with the function `rate`:

```go
package main

import (
	"fmt"

	promqlbuilder "github.com/perses/promql-builder"
	"github.com/perses/promql-builder/matrix"
	"github.com/perses/promql-builder/vector"
)

func main() {
	m := promqlbuilder.Rate(
		matrix.New(
			vector.New(vector.WithMetricName("foo")),
			matrix.WithRangeAsString("1h2m4s"),
		),
	)
	fmt.Print(m.String())
}
```

It will give the following output:

```text
rate(foo[1h2m4s])
```

### Create your custom/non-standard PromQL function

If you want to use a function that is not provided by this library, you can create your own function using the
`promqlbuilder.NewFunction` function.

For example, with a custom function `my_custom_function`:

```go
package main
import (
    "fmt"

    promqlbuilder "github.com/perses/promql-builder"
    "github.com/perses/promql-builder/matrix"
    "github.com/perses/promql-builder/vector"
)
func main() {
    m := promqlbuilder.NewFunction("my_custom_function",
        matrix.New(
            vector.New(vector.WithMetricName("foo")),
            matrix.WithRangeAsString("1h2m4s"),
        ),
        promqlbuilder.NewNumberLiteral(123),
        promqlbuilder.NewStringLiteral("bar"),
    )
    fmt.Print(m.String())
}
```

It will give the following output:

```text
my_custom_function(foo[1h2m4s], 123, "bar")
```

### Use aggregation function

As you may know, all aggregation functions can be combined with the keywords `by` or `without` in order to specify on
which labels you would like to aggregate the timeseries.

In this builder, these keywords are a function that you can use once you have created the aggregation function.

For example, with the aggregation function `sum`:

```go
package main

import (
	"fmt"

	promqlbuilder "github.com/perses/promql-builder"
	"github.com/perses/promql-builder/matrix"
	"github.com/perses/promql-builder/vector"
)

func main() {
	m :=
		promqlbuilder.Sum(
			promqlbuilder.Rate(
				matrix.New(
					vector.New(vector.WithMetricName("foo")),
					matrix.WithRangeAsString("1h2m4s"),
				),
			),
		).By("namespace")
	fmt.Print(m.String())
}
```

It will give the following output:

```text
sum by("namespace") (rate(foo[1h2m4s]))
```

### Use binary operation

Binary operation can be used with the vector matching keywords `on`, `ignoring`.
Same like the aggregation function, the usage of these keywords can be done once you have built your binary operation.

```go
package main

import (
	"fmt"

	promqlbuilder "github.com/perses/promql-builder"
	"github.com/perses/promql-builder/matrix"
	"github.com/perses/promql-builder/vector"
)

func main() {
	m :=
		promqlbuilder.Add(
			promqlbuilder.Rate(
				matrix.New(
					vector.New(vector.WithMetricName("foo")),
					matrix.WithRangeAsString("1h2m4s"),
				),
			),
			promqlbuilder.Vector(123),
		).On("namespace")
	fmt.Print(m.String())
}

```

```text
rate(foo[1h2m4s]) + on("namespace") vector(123)
```

Note: the Group modifiers (`group_left` or `group_right`) can be used once the vector matching keywords are used.

### Iterate through PromQL AST

This lib also provides Prometheus-inspired PromQL AST iteration methods such as `Inspect`, `Walk`, `Children`, that can handle the 
custom nodes of expressions constructed using this lib, as well as PromQL expression deep copying using `DeepCopyExpr`.

This allows constructing utilities like,
```go
func SetLabelMatchers(query parser.Expr, matchers []*labels.Matcher) parser.Expr {
	copy := promqlbuilder.DeepCopyExpr(query)
	for _, l := range matchers {
		copy = LabelsSetPromQL(copy, l.Type, l.Name, l.Value)
	}
	return copy
}

func LabelsSetPromQL(query parser.Expr, matchType labels.MatchType, name, value string) parser.Expr {
	if name == "" || value == "" {
		return query
	}

	promqlbuilder.Inspect(query, func(node parser.Node, path []parser.Node) error {
		if n, ok := node.(*parser.VectorSelector); ok {
			var found bool
			for i, l := range n.LabelMatchers {
				if l.Name == name {
					n.LabelMatchers[i].Type = matchType
					n.LabelMatchers[i].Value = value
					found = true
				}
			}
			if !found {
				n.LabelMatchers = append(n.LabelMatchers, &labels.Matcher{
					Type:  matchType,
					Name:  name,
					Value: value,
				})
			}
		}
		return nil
	})

	return query
}
```
