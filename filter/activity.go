package filter

import "time"

type ActivitySifter struct {
	startTimestamp int64
	endTimestamp   int64
	next           Sifter
}

func (a *ActivitySifter) execute() bool {
	now := time.Now().UnixMilli()
	if now < a.startTimestamp || now > a.endTimestamp {
		return false
	}
	if a.next == nil {
		return true
	}
	return a.next.execute()
}

func (a *ActivitySifter) setNext(s Sifter) {
	a.next = s
}
