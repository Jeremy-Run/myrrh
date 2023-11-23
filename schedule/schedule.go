package schedule

type Requirement struct {
	StartTimestamp  int64
	EndTimestamp    int64
	Region          string
	CompletionLimit string
	Expr            *Expression
	Action          string
}
