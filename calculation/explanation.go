package calculation

func ExplainRelation(op string) func(l, r int) bool {
	c := Calculation{}
	switch op {
	case ">=":
		return c.GreaterThanOrEqual
	case "<=":
		return c.LessThanOrEqual
	}
	return nil
}

func ExplainLogic(op string) func(bools ...bool) bool {
	c := Calculation{}
	switch op {
	case "&&":
		return c.AND
	case "||":
		return c.OR
	}
	return nil
}
