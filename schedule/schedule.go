package schedule

import (
	"myrrh/business"
	"time"
)

type Requirement struct {
	StartTimestamp  int64
	EndTimestamp    int64
	Region          string
	CompletionLimit string
	Expr            *Expression
	Action          string
}

func (r *Requirement) Execute() bool {
	now := time.Now().UnixMilli()
	if now < r.StartTimestamp || now > r.EndTimestamp {
		return false
	}

	user := business.NewSgpUser()
	if r.Region != user.Region {
		return false
	}

	limit := 0
	switch r.CompletionLimit {
	case "OnceDuringActivity":
		limit = 1
	}
	activity := business.Activity{}
	if activity.FinishTimes() >= limit {
		return false
	}

	result := ExpressionEvaluation(r.Expr)

	if result {
		business.DoAction(r.Action)
	}

	return result
}
