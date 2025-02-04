package function

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

func newNumber(num float64) *parser.NumberLiteral {
	return &parser.NumberLiteral{
		Val: num,
	}
}

func newString(s string) *parser.StringLiteral {
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

func AbsentOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("absent_over_time", matrix)
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

func AvgOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("avg_over_time", matrix)
}

func Ceil(vector parser.Expr) *parser.Call {
	return newFunction("ceil", vector)
}

func Changes(matrix *matrix.Builder) *parser.Call {
	return newFunction("changes", matrix)
}

func Clamp(vector parser.Expr, min float64, max float64) *parser.Call {
	return newFunction("clamp", vector, newNumber(min), newNumber(max))
}

func ClampMax(vector parser.Expr, max float64) *parser.Call {
	return newFunction("clamp_max", vector, newNumber(max))
}

func ClampMin(vector parser.Expr, min float64) *parser.Call {
	return newFunction("clamp_min", vector, newNumber(min))
}

func Cos(vector parser.Expr) *parser.Call {
	return newFunction("cos", vector)
}

func Cosh(vector parser.Expr) *parser.Call {
	return newFunction("cosh", vector)
}

func CountOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("count_over_time", matrix)
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

func Delta(matrix *matrix.Builder) *parser.Call {
	return newFunction("delta", matrix)
}

func Deriv(matrix *matrix.Builder) *parser.Call {
	return newFunction("deriv", matrix)
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
	return newFunction("histogram_fraction", newNumber(lower), newNumber(upper), vector)
}

func HistogramQuantile(quantile float64, vector parser.Expr) *parser.Call {
	return newFunction("histogram_quantile", newNumber(quantile), vector)
}

func DoubleExponentialSmoothing(matrix *matrix.Builder, smoothingFactor float64, trendFactor float64) *parser.Call {
	return newFunction("double_exponential_smoothing", matrix, newNumber(smoothingFactor), newNumber(trendFactor))
}

func Hour(vector parser.Expr) *parser.Call {
	return newFunction("hour", vector)
}

func IDelta(matrix *matrix.Builder) *parser.Call {
	return newFunction("idelta", matrix)
}

func Increase(matrix *matrix.Builder) *parser.Call {
	return newFunction("increase", matrix)
}

func Info(vector parser.Expr, dataLabelSelector parser.Expr) *parser.Call {
	return newFunction("info", vector, dataLabelSelector)
}

func IRate(matrix *matrix.Builder) *parser.Call {
	return newFunction("irate", matrix)
}

func LabelReplace(vector parser.Expr, destinationLabel string, replacement string, sourceLabel string, regexp string) *parser.Call {
	return newFunction("label_replace", vector, newString(destinationLabel), newString(replacement), newString(sourceLabel), newString(regexp))
}

func LabelJoin(vector parser.Expr, destinationLabel string, replacement string, srcLabels ...string) *parser.Call {
	args := []parser.Expr{vector, newString(destinationLabel), newString(replacement)}
	for _, label := range srcLabels {
		args = append(args, newString(label))
	}
	return newFunction("label_join", args...)
}

func LastOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("last_over_time", matrix)
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

func MadOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("mad_over_time", matrix)
}

func MaxOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("max_over_time", matrix)
}

func MinOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("min_over_time", matrix)
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

func PredictLinear(matrix *matrix.Builder, t float64) *parser.Call {
	return newFunction("predict_linear", matrix, newNumber(t))
}

func PresentOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("present_over_time", matrix)
}

func QuantileOverTime(t float64, matrix *matrix.Builder) *parser.Call {
	return newFunction("quantile_over_time", newNumber(t), matrix)
}

func Rad(vector parser.Expr) *parser.Call {
	return newFunction("rad", vector)
}

func Rate(matrix *matrix.Builder) *parser.Call {
	return newFunction("rate", matrix)
}

func Resets(matrix *matrix.Builder) *parser.Call {
	return newFunction("resets", matrix)
}

func Round(vector parser.Expr, t float64) *parser.Call {
	return newFunction("round", vector, newNumber(t))
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
		args = append(args, newString(label))
	}
	return newFunction("sort_by_label", args...)
}

func SortByLabelDesc(vector parser.Expr, labels ...string) *parser.Call {
	args := []parser.Expr{vector}
	for _, label := range labels {
		args = append(args, newString(label))
	}
	return newFunction("sort_by_label_desc", args...)
}

func Sqrt(vector parser.Expr) *parser.Call {
	return newFunction("sqrt", vector)
}

func StddevOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("stddev_over_time", matrix)
}

func StdvarOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("stdvar_over_time", matrix)
}

func SumOverTime(matrix *matrix.Builder) *parser.Call {
	return newFunction("sum_over_time", matrix)
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
	return newFunction("vector", newNumber(scalar))
}

func Year(vector parser.Expr) *parser.Call {
	return newFunction("year", vector)
}
