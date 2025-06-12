package promqlbuilder

import (
	"github.com/perses/promql-builder/matrix"
	"github.com/prometheus/prometheus/promql/parser"
)

func newFunction(name string, args ...parser.Expr) *parser.Call {
	return &parser.Call{
		Func: parser.Functions[name],
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
	return newFunction("abs", vector)
}

func Absent(vector parser.Expr) *parser.Call {
	return newFunction("absent", vector)
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
	return newFunction("absent_over_time", convertToExpr(input))
}

func Acos(vector parser.Expr) *parser.Call {
	return newFunction("acos", vector)
}

func Acosh(vector parser.Expr) *parser.Call {
	return newFunction("acosh", vector)
}

func Asin(vector parser.Expr) *parser.Call {
	return newFunction("asin", vector)
}

func Asinh(vector parser.Expr) *parser.Call {
	return newFunction("asinh", vector)
}

func Atan(vector parser.Expr) *parser.Call {
	return newFunction("atan", vector)
}

func Atanh(vector parser.Expr) *parser.Call {
	return newFunction("atanh", vector)
}

func AvgOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("avg_over_time", convertToExpr(input))
}

func Ceil(vector parser.Expr) *parser.Call {
	return newFunction("ceil", vector)
}

func Changes[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("changes", convertToExpr(input))
}

func Clamp(vector parser.Expr, min float64, max float64) *parser.Call {
	return newFunction("clamp", vector, NewNumber(min), NewNumber(max))
}

func ClampMax(vector parser.Expr, max float64) *parser.Call {
	return newFunction("clamp_max", vector, NewNumber(max))
}

func ClampMin(vector parser.Expr, min float64) *parser.Call {
	return newFunction("clamp_min", vector, NewNumber(min))
}

func Cos(vector parser.Expr) *parser.Call {
	return newFunction("cos", vector)
}

func Cosh(vector parser.Expr) *parser.Call {
	return newFunction("cosh", vector)
}

func CountOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("count_over_time", convertToExpr(input))
}

func DaysInMonth(vector parser.Expr) *parser.Call {
	return newFunction("days_in_month", vector)
}

func DaysOfMonth(vector parser.Expr) *parser.Call {
	return newFunction("days_of_month", vector)
}

func DaysOfWeek(vector parser.Expr) *parser.Call {
	return newFunction("days_of_week", vector)
}

func DaysOfYear(vector parser.Expr) *parser.Call {
	return newFunction("days_of_year", vector)
}

func Deg(vector parser.Expr) *parser.Call {
	return newFunction("deg", vector)
}

func Delta[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("delta", convertToExpr(input))
}

func Deriv[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("deriv", convertToExpr(input))
}

func Exp(vector parser.Expr) *parser.Call {
	return newFunction("exp", vector)
}

func Floor(vector parser.Expr) *parser.Call {
	return newFunction("floor", vector)
}

func HistogramAvg(vector parser.Expr) *parser.Call {
	return newFunction("histogram_avg", vector)
}

func HistogramCount(vector parser.Expr) *parser.Call {
	return newFunction("histogram_count", vector)
}

func HistogramSum(vector parser.Expr) *parser.Call {
	return newFunction("histogram_sum", vector)
}

func HistogramStddev(vector parser.Expr) *parser.Call {
	return newFunction("histogram_stddev", vector)
}

func HistogramStdvar(vector parser.Expr) *parser.Call {
	return newFunction("histogram_stdvar", vector)
}

func HistogramFraction(lower float64, upper float64, vector parser.Expr) *parser.Call {
	return newFunction("histogram_fraction", NewNumber(lower), NewNumber(upper), vector)
}

func HistogramQuantile(quantile float64, vector parser.Expr) *parser.Call {
	return newFunction("histogram_quantile", NewNumber(quantile), vector)
}

func DoubleExponentialSmoothing[T RangeVectorBuilder](input T, smoothingFactor float64, trendFactor float64) *parser.Call {
	return newFunction("double_exponential_smoothing", convertToExpr(input), NewNumber(smoothingFactor), NewNumber(trendFactor))
}

func Hour(vector parser.Expr) *parser.Call {
	return newFunction("hour", vector)
}

func IDelta[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("idelta", convertToExpr(input))
}

func Increase[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("increase", convertToExpr(input))
}

func Info(vector parser.Expr, dataLabelSelector parser.Expr) *parser.Call {
	return newFunction("info", vector, dataLabelSelector)
}

func IRate[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("irate", convertToExpr(input))
}

func LabelReplace(vector parser.Expr, destinationLabel string, replacement string, sourceLabel string, regexp string) *parser.Call {
	return newFunction("label_replace", vector, NewString(destinationLabel), NewString(replacement), NewString(sourceLabel), NewString(regexp))
}

func LabelJoin(vector parser.Expr, destinationLabel string, replacement string, srcLabels ...string) *parser.Call {
	args := []parser.Expr{vector, NewString(destinationLabel), NewString(replacement)}
	for _, label := range srcLabels {
		args = append(args, NewString(label))
	}
	return newFunction("label_join", args...)
}

func LastOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("last_over_time", convertToExpr(input))
}

func Ln(vector parser.Expr) *parser.Call {
	return newFunction("ln", vector)
}

func Log10(vector parser.Expr) *parser.Call {
	return newFunction("log10", vector)
}

func Log2(vector parser.Expr) *parser.Call {
	return newFunction("log2", vector)
}

func MadOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("mad_over_time", convertToExpr(input))
}

func MaxOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("max_over_time", convertToExpr(input))
}

func MinOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("min_over_time", convertToExpr(input))
}

func Minute(vector parser.Expr) *parser.Call {
	return newFunction("minute", vector)
}

func Month(vector parser.Expr) *parser.Call {
	return newFunction("month", vector)
}

func PI() *parser.Call {
	return newFunction("pi")
}

func PredictLinear[T RangeVectorBuilder](input T, t float64) *parser.Call {
	return newFunction("predict_linear", convertToExpr(input), NewNumber(t))
}

func PresentOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("present_over_time", convertToExpr(input))
}

func QuantileOverTime[T RangeVectorBuilder](t float64, input T) *parser.Call {
	return newFunction("quantile_over_time", NewNumber(t), convertToExpr(input))
}

func Rad(vector parser.Expr) *parser.Call {
	return newFunction("rad", vector)
}

func Rate[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("rate", convertToExpr(input))
}

func Resets[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("resets", convertToExpr(input))
}

func Round(vector parser.Expr, t float64) *parser.Call {
	return newFunction("round", vector, NewNumber(t))
}

func Scalar(vector parser.Expr) *parser.Call {
	return newFunction("scalar", vector)
}

func Sgn(vector parser.Expr) *parser.Call {
	return newFunction("sgn", vector)
}

func Sin(vector parser.Expr) *parser.Call {
	return newFunction("sin", vector)
}

func Sinh(vector parser.Expr) *parser.Call {
	return newFunction("sinh", vector)
}

func Sort(vector parser.Expr) *parser.Call {
	return newFunction("sort", vector)
}

func SortDesc(vector parser.Expr) *parser.Call {
	return newFunction("sort_desc", vector)
}

func SortByLabel(vector parser.Expr, labels ...string) *parser.Call {
	args := []parser.Expr{vector}
	for _, label := range labels {
		args = append(args, NewString(label))
	}
	return newFunction("sort_by_label", args...)
}

func SortByLabelDesc(vector parser.Expr, labels ...string) *parser.Call {
	args := []parser.Expr{vector}
	for _, label := range labels {
		args = append(args, NewString(label))
	}
	return newFunction("sort_by_label_desc", args...)
}

func Sqrt(vector parser.Expr) *parser.Call {
	return newFunction("sqrt", vector)
}

func StddevOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("stddev_over_time", convertToExpr(input))
}

func StdvarOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("stdvar_over_time", convertToExpr(input))
}

func SumOverTime[T RangeVectorBuilder](input T) *parser.Call {
	return newFunction("sum_over_time", convertToExpr(input))
}

func Tan(vector parser.Expr) *parser.Call {
	return newFunction("tan", vector)
}

func Tanh(vector parser.Expr) *parser.Call {
	return newFunction("tanh", vector)
}

func Time() *parser.Call {
	return newFunction("time")
}

func Timestamp(vector parser.Expr) *parser.Call {
	return newFunction("timestamp", vector)
}

func Vector(scalar float64) *parser.Call {
	return newFunction("vector", NewNumber(scalar))
}

func Year(vector parser.Expr) *parser.Call {
	return newFunction("year", vector)
}
