package schedule

import (
	"myrrh/business"
	"myrrh/filter"
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
	sp := filter.SifterParams{
		StartTimestamp:  r.StartTimestamp,
		EndTimestamp:    r.EndTimestamp,
		Region:          r.Region,
		CompletionLimit: r.CompletionLimit,
	}
	isPass := filter.ExecuteSifter(sp)
	if !isPass {
		return false
	}

	result := ExpressionEvaluation(r.Expr)

	if result {
		business.DoAction(r.Action)
	}

	return result
}
