package calculation

type Calculation struct{}

func (c *Calculation) GreaterThanOrEqual(l, r int) bool {
	return l >= r
}

func (c *Calculation) LessThanOrEqual(l, r int) bool {
	return l <= r
}

func (c *Calculation) GreaterThan(l, r int) bool {
	return l > r
}

func (c *Calculation) LessThan(l, r int) bool {
	return l < r
}

func (c *Calculation) Equal(l, r int) bool {
	return l == r
}

func (c *Calculation) NotEqual(l, r int) bool {
	return l != r
}

func (c *Calculation) AND(bools ...bool) bool {
	for _, b := range bools {
		if !b {
			return false
		}
	}
	return true
}

func (c *Calculation) OR(bools ...bool) bool {
	for _, b := range bools {
		if b {
			return true
		}
	}
	return false
}
