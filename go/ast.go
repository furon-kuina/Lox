package main

type Expr interface {
	accept(visitor) string
	aExpr()
}

type BinaryExpr struct {
	op          Token
	left, right Expr
}

type GroupingExpr struct {
	expr Expr
}

type Printable interface {
	ToString() string
}

type LiteralExpr struct {
	value Printable
}

type UnaryExpr struct {
	right Expr
	op    Token
}

func (be *BinaryExpr) aExpr()   {}
func (ge *GroupingExpr) aExpr() {}
func (le *LiteralExpr) aExpr()  {}
func (ue *UnaryExpr) aExpr()    {}

type visitor interface {
	visitBinaryExpr(*BinaryExpr) string
	visitGroupingExpr(*GroupingExpr) string
	visitLiteralExpr(*LiteralExpr) string
	visitUnaryExpr(*UnaryExpr) string
}

func (be *BinaryExpr) accept(v visitor) string {
	return v.visitBinaryExpr(be)
}

func (ge *GroupingExpr) accept(v visitor) string {
	return v.visitGroupingExpr(ge)
}

func (le *LiteralExpr) accept(v visitor) string {
	return v.visitLiteralExpr(le)
}

func (ue *UnaryExpr) accept(v visitor) string {
	return v.visitUnaryExpr(ue)
}

type AstPrinter struct{}

func (ap AstPrinter) print(expr Expr) string {
	return expr.accept(ap)
}

func (ap AstPrinter) visitBinaryExpr(be *BinaryExpr) string {
	return ap.parenthesize(be.op.lexeme, be.left, be.right)
}

func (ap AstPrinter) visitGroupingExpr(ge *GroupingExpr) string {
	return ap.parenthesize("group", ge.expr)
}

func (ap AstPrinter) visitLiteralExpr(le *LiteralExpr) string {
	if le.value == nil {
		return "nil"
	} else {
		return le.value.ToString()
	}
}

func (ap AstPrinter) visitUnaryExpr(ue *UnaryExpr) string {
	return ap.parenthesize(ue.op.lexeme, ue.right)
}

func (ap AstPrinter) parenthesize(name string, exprs ...Expr) (res string) {
	res += "(" + name
	for _, expr := range exprs {
		res += " "
		res += expr.accept(ap)
	}
	res += ")"
	return
}
