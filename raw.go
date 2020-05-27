package goquery

type RawExpression struct {
	exp string
}

func Raw(exp string) RawExpression {
	return RawExpression{exp: exp}
}
