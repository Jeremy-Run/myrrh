package filter

import "myrrh/business"

type CompletionLimitSifter struct {
	completionLimit string
	next            Sifter
}

func (c *CompletionLimitSifter) execute() bool {
	limit := 0
	switch c.completionLimit {
	case "OnceDuringActivity":
		limit = 1
	}
	activity := business.Activity{}
	if activity.FinishTimes() >= limit {
		return false
	}
	if c.next == nil {
		return true
	}
	return c.next.execute()
}

func (c *CompletionLimitSifter) setNext(s Sifter) {
	c.next = s
}
