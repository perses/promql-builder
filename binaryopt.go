package promqlbuilder

import (
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/prometheus/prometheus/promql/parser/posrange"
)

type BinaryBuilder struct {
	internal *parser.BinaryExpr
}

func (b *BinaryBuilder) Type() parser.ValueType {
	return b.internal.Type()
}
func (b *BinaryBuilder) PromQLExpr() {
	b.internal.PromQLExpr()
}
func (b *BinaryBuilder) String() string {
	return b.internal.String()
}
func (b *BinaryBuilder) Pretty(level int) string {
	return b.internal.Pretty(level)
}
func (b *BinaryBuilder) PositionRange() posrange.PositionRange {
	return b.internal.PositionRange()
}
func (b *BinaryBuilder) Bool() *BinaryBuilder {
	b.internal.ReturnBool = true
	return b
}

func (b *BinaryBuilder) Ignoring(labels ...string) *BinaryWithVectorMatching {
	b.internal.VectorMatching = &parser.VectorMatching{
		MatchingLabels: labels,
		On:             false,
	}
	return &BinaryWithVectorMatching{
		binaryOpt: b,
	}
}

func (b *BinaryBuilder) On(labels ...string) *BinaryWithVectorMatching {
	b.internal.VectorMatching = &parser.VectorMatching{
		MatchingLabels: labels,
		On:             true,
	}
	return &BinaryWithVectorMatching{
		binaryOpt: b,
	}
}

type BinaryWithVectorMatching struct {
	binaryOpt *BinaryBuilder
}

func (b *BinaryWithVectorMatching) Type() parser.ValueType {
	return b.binaryOpt.Type()
}
func (b *BinaryWithVectorMatching) PromQLExpr() {
	b.binaryOpt.PromQLExpr()
}
func (b *BinaryWithVectorMatching) String() string {
	return b.binaryOpt.String()
}
func (b *BinaryWithVectorMatching) Pretty(level int) string {
	return b.binaryOpt.Pretty(level)
}
func (b *BinaryWithVectorMatching) PositionRange() posrange.PositionRange {
	return b.binaryOpt.PositionRange()
}

func (b *BinaryWithVectorMatching) Bool() *BinaryWithVectorMatching {
	b.binaryOpt.internal.ReturnBool = true
	return b
}

func (b *BinaryWithVectorMatching) GroupLeft(labels ...string) *BinaryWithVectorMatching {
	b.binaryOpt.internal.VectorMatching.Include = labels
	b.binaryOpt.internal.VectorMatching.Card = parser.CardManyToOne
	return b
}

func (b *BinaryWithVectorMatching) GroupRight(labels ...string) *BinaryWithVectorMatching {
	b.binaryOpt.internal.VectorMatching.Include = labels
	b.binaryOpt.internal.VectorMatching.Card = parser.CardOneToMany
	return b
}

func createBinaryOperation(itemType parser.ItemType, left parser.Expr, right parser.Expr) *BinaryBuilder {
	b := &BinaryBuilder{
		internal: &parser.BinaryExpr{
			Op:  itemType,
			LHS: left,
			RHS: right,
		},
	}
	return b
}

func Pow(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.POW, left, right)
}

func Mul(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.MUL, left, right)
}

func Div(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.DIV, left, right)
}

func Mod(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.MOD, left, right)
}

func Atan2(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.ATAN2, left, right)
}

func Add(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.ADD, left, right)
}

func Sub(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.SUB, left, right)
}

func Eql(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.EQL, left, right)
}

func Eqlc(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.EQLC, left, right)
}

func Gte(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.GTE, left, right)
}

func Gtr(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.GTR, left, right)
}

func Lte(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.LTE, left, right)
}

func Lss(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.LSS, left, right)
}
func Neq(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.NEQ, left, right)
}
func And(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.LAND, left, right)
}
func Unless(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.LUNLESS, left, right)
}
func Or(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.LOR, left, right)
}

func NeqRegex(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.NEQ_REGEX, left, right)
}

func EqlRegex(left parser.Expr, right parser.Expr) *BinaryBuilder {
	return createBinaryOperation(parser.EQL_REGEX, left, right)
}
