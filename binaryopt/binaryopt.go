package binaryopt

import "github.com/prometheus/prometheus/promql/parser"

type Builder parser.BinaryExpr

type VectorMatching parser.VectorMatching

func (v *VectorMatching) GroupLeft(labels ...string) *VectorMatching {
	v.Include = labels
	v.Card = parser.CardManyToOne
	return v
}

func (v *VectorMatching) GroupRight(labels ...string) *VectorMatching {
	v.Include = labels
	v.Card = parser.CardOneToMany
	return v
}

func Ignoring(labels ...string) *VectorMatching {
	return &VectorMatching{
		MatchingLabels: labels,
		On:             false,
	}
}

func On(labels ...string) *VectorMatching {
	return &VectorMatching{
		MatchingLabels: labels,
		On:             true,
	}
}

func create(itemType parser.ItemType, left parser.Expr, right parser.Expr, matchers []*VectorMatching) *parser.BinaryExpr {
	b := &Builder{
		Op:  itemType,
		LHS: left,
		RHS: right,
	}

	if len(matchers) > 0 {
		b.VectorMatching = (*parser.VectorMatching)(matchers[0])
	}

	return (*parser.BinaryExpr)(b)
}

func Pow(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.POW, left, right, matchers)
}

func Mul(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.MUL, left, right, matchers)
}

func Div(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.DIV, left, right, matchers)
}

func Mod(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.MOD, left, right, matchers)
}

func Atan2(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.ATAN2, left, right, matchers)
}

func Add(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.ADD, left, right, matchers)
}

func Sub(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.SUB, left, right, matchers)
}

func Eql(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.EQL, left, right, matchers)
}

func Gte(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.GTE, left, right, matchers)
}

func Gtr(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.GTR, left, right, matchers)
}

func Lte(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.LTE, left, right, matchers)
}

func Lss(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.LSS, left, right, matchers)
}
func Neq(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.NEQ, left, right, matchers)
}
func And(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.LAND, left, right, matchers)
}
func Unless(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.LUNLESS, left, right, matchers)
}
func Or(left parser.Expr, right parser.Expr, matchers ...*VectorMatching) *parser.BinaryExpr {
	return create(parser.LOR, left, right, matchers)
}
