package promqlbuilder

import (
	"github.com/perses/promql-builder/matrix"
	"github.com/prometheus/prometheus/promql/parser"
)

func NewFunction(name string, args ...parser.Expr) *parser.Call {
	fn, ok := parser.Functions[name]
	if !ok {
		fn = &parser.Function{Name: name}
	}
	return &parser.Call{
		Func: fn,
		Args: args,
	}
}

func NewNumber(num float64) *parser.NumberLiteral {
	return &parser.NumberLiteral{
		Val: num,
	}
}

func NewString(s string) *parser.StringLiteral {
	return &parser.StringLiteral{
		Val: s,
	}
}

func Abs(vector parser.Expr) *parser.Call {
	return NewFunction("abs", vector)
}

func Absent(vector parser.Expr) *parser.Call {
	return NewFunction("absent", vector)
}

type RangeVectorBuilder interface {
	*matrix.Builder | *parser.SubqueryExpr
}

func convertToExpr[T RangeVectorBuilder](input T) parser.Expr {
	switch v := any(input).(type) {
	case *matrix.Builder:
		return v
	case *parser.SubqueryExpr:
		return v
	default:
		panic("unsupported type")
	}
}

func AbsentOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("absent_over_time", convertToExpr(input))
}

func Acos(vector parser.Expr) *parser.Call {
	return NewFunction("acos", vector)
}

func Acosh(vector parser.Expr) *parser.Call {
	return NewFunction("acosh", vector)
}

func Asin(vector parser.Expr) *parser.Call {
	return NewFunction("asin", vector)
}

func Asinh(vector parser.Expr) *parser.Call {
	return NewFunction("asinh", vector)
}

func Atan(vector parser.Expr) *parser.Call {
	return NewFunction("atan", vector)
}

func Atanh(vector parser.Expr) *parser.Call {
	return NewFunction("atanh", vector)
}

func AvgOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("avg_over_time", convertToExpr(input))
}

func Ceil(vector parser.Expr) *parser.Call {
	return NewFunction("ceil", vector)
}

func Changes[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("changes", convertToExpr(input))
}

func Clamp(vector parser.Expr, min float64, max float64) *parser.Call {
	return NewFunction("clamp", vector, NewNumber(min), NewNumber(max))
}

func ClampMax(vector parser.Expr, max float64) *parser.Call {
	return NewFunction("clamp_max", vector, NewNumber(max))
}

func ClampMin(vector parser.Expr, min float64) *parser.Call {
	return NewFunction("clamp_min", vector, NewNumber(min))
}

func Cos(vector parser.Expr) *parser.Call {
	return NewFunction("cos", vector)
}

func Cosh(vector parser.Expr) *parser.Call {
	return NewFunction("cosh", vector)
}

func CountOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("count_over_time", convertToExpr(input))
}

func DaysInMonth(vector parser.Expr) *parser.Call {
	return NewFunction("days_in_month", vector)
}

func DaysOfMonth(vector parser.Expr) *parser.Call {
	return NewFunction("days_of_month", vector)
}

func DaysOfWeek(vector parser.Expr) *parser.Call {
	return NewFunction("days_of_week", vector)
}

func DaysOfYear(vector parser.Expr) *parser.Call {
	return NewFunction("days_of_year", vector)
}

func Deg(vector parser.Expr) *parser.Call {
	return NewFunction("deg", vector)
}

func Delta[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("delta", convertToExpr(input))
}

func Deriv[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("deriv", convertToExpr(input))
}

func Exp(vector parser.Expr) *parser.Call {
	return NewFunction("exp", vector)
}

func Floor(vector parser.Expr) *parser.Call {
	return NewFunction("floor", vector)
}

func HistogramAvg(vector parser.Expr) *parser.Call {
	return NewFunction("histogram_avg", vector)
}

func HistogramCount(vector parser.Expr) *parser.Call {
	return NewFunction("histogram_count", vector)
}

func HistogramSum(vector parser.Expr) *parser.Call {
	return NewFunction("histogram_sum", vector)
}

func HistogramStddev(vector parser.Expr) *parser.Call {
	return NewFunction("histogram_stddev", vector)
}

func HistogramStdvar(vector parser.Expr) *parser.Call {
	return NewFunction("histogram_stdvar", vector)
}

func HistogramFraction(lower float64, upper float64, vector parser.Expr) *parser.Call {
	return NewFunction("histogram_fraction", NewNumber(lower), NewNumber(upper), vector)
}

func HistogramQuantile(quantile float64, vector parser.Expr) *parser.Call {
	return NewFunction("histogram_quantile", NewNumber(quantile), vector)
}

func DoubleExponentialSmoothing[T RangeVectorBuilder](input T, smoothingFactor float64, trendFactor float64) *parser.Call {
	return NewFunction("double_exponential_smoothing", convertToExpr(input), NewNumber(smoothingFactor), NewNumber(trendFactor))
}

func Hour(vector parser.Expr) *parser.Call {
	return NewFunction("hour", vector)
}

func IDelta[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("idelta", convertToExpr(input))
}

func Increase[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("increase", convertToExpr(input))
}

func Info(vector parser.Expr, dataLabelSelector parser.Expr) *parser.Call {
	return NewFunction("info", vector, dataLabelSelector)
}

func IRate[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("irate", convertToExpr(input))
}

func LabelReplace(vector parser.Expr, destinationLabel string, replacement string, sourceLabel string, regexp string) *parser.Call {
	return NewFunction("label_replace", vector, NewString(destinationLabel), NewString(replacement), NewString(sourceLabel), NewString(regexp))
}

func LabelJoin(vector parser.Expr, destinationLabel string, replacement string, srcLabels ...string) *parser.Call {
	args := []parser.Expr{vector, NewString(destinationLabel), NewString(replacement)}
	for _, label := range srcLabels {
		args = append(args, NewString(label))
	}
	return NewFunction("label_join", args...)
}

func LastOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("last_over_time", convertToExpr(input))
}

func Ln(vector parser.Expr) *parser.Call {
	return NewFunction("ln", vector)
}

func Log10(vector parser.Expr) *parser.Call {
	return NewFunction("log10", vector)
}

func Log2(vector parser.Expr) *parser.Call {
	return NewFunction("log2", vector)
}

func MadOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("mad_over_time", convertToExpr(input))
}

func MaxOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("max_over_time", convertToExpr(input))
}

func MinOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("min_over_time", convertToExpr(input))
}

func Minute(vector parser.Expr) *parser.Call {
	return NewFunction("minute", vector)
}

func Month(vector parser.Expr) *parser.Call {
	return NewFunction("month", vector)
}

func PI() *parser.Call {
	return NewFunction("pi")
}

func PredictLinear[T RangeVectorBuilder](input T, t float64) *parser.Call {
	return NewFunction("predict_linear", convertToExpr(input), NewNumber(t))
}

func PresentOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("present_over_time", convertToExpr(input))
}

func QuantileOverTime[T RangeVectorBuilder](t float64, input T) *parser.Call {
	return NewFunction("quantile_over_time", NewNumber(t), convertToExpr(input))
}

func Rad(vector parser.Expr) *parser.Call {
	return NewFunction("rad", vector)
}

func Rate[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("rate", convertToExpr(input))
}

func Resets[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("resets", convertToExpr(input))
}

func Round(vector parser.Expr, t float64) *parser.Call {
	return NewFunction("round", vector, NewNumber(t))
}

func Scalar(vector parser.Expr) *parser.Call {
	return NewFunction("scalar", vector)
}

func Sgn(vector parser.Expr) *parser.Call {
	return NewFunction("sgn", vector)
}

func Sin(vector parser.Expr) *parser.Call {
	return NewFunction("sin", vector)
}

func Sinh(vector parser.Expr) *parser.Call {
	return NewFunction("sinh", vector)
}

func Sort(vector parser.Expr) *parser.Call {
	return NewFunction("sort", vector)
}

func SortDesc(vector parser.Expr) *parser.Call {
	return NewFunction("sort_desc", vector)
}

func SortByLabel(vector parser.Expr, labels ...string) *parser.Call {
	args := []parser.Expr{vector}
	for _, label := range labels {
		args = append(args, NewString(label))
	}
	return NewFunction("sort_by_label", args...)
}

func SortByLabelDesc(vector parser.Expr, labels ...string) *parser.Call {
	args := []parser.Expr{vector}
	for _, label := range labels {
		args = append(args, NewString(label))
	}
	return NewFunction("sort_by_label_desc", args...)
}

func Sqrt(vector parser.Expr) *parser.Call {
	return NewFunction("sqrt", vector)
}

func StddevOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("stddev_over_time", convertToExpr(input))
}

func StdvarOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("stdvar_over_time", convertToExpr(input))
}

func SumOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return NewFunction("sum_over_time", convertToExpr(input))
}

func Tan(vector parser.Expr) *parser.Call {
	return NewFunction("tan", vector)
}

func Tanh(vector parser.Expr) *parser.Call {
	return NewFunction("tanh", vector)
}

func Time() *parser.Call {
	return NewFunction("time")
}

func Timestamp(vector parser.Expr) *parser.Call {
	return NewFunction("timestamp", vector)
}

func Vector(scalar float64) *parser.Call {
	return NewFunction("vector", NewNumber(scalar))
}

func Year(vector parser.Expr) *parser.Call {
	return NewFunction("year", vector)
}
