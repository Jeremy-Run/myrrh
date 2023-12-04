package filter

import "myrrh/business"

type UserSifter struct {
	region string
	next   Sifter
}

func (u *UserSifter) execute() bool {
	user := business.NewSgpUser()
	if u.region != user.Region {
		return false
	}
	if u.next == nil {
		return true
	}
	return u.next.execute()
}

func (u *UserSifter) setNext(s Sifter) {
	u.next = s
}
