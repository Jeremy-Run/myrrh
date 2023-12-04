package filter

type Sifter interface {
	execute() bool
	setNext(Sifter)
}

type SifterParams struct {
	StartTimestamp  int64
	EndTimestamp    int64
	Region          string
	CompletionLimit string
}

func ExecuteSifter(p SifterParams) bool {
	cls := CompletionLimitSifter{
		completionLimit: p.CompletionLimit,
	}

	us := UserSifter{
		region: p.Region,
	}
	us.setNext(&cls)

	as := ActivitySifter{
		startTimestamp: p.StartTimestamp,
		endTimestamp:   p.EndTimestamp,
	}
	as.setNext(&us)

	return as.execute()
}
