package business

type User struct {
	Region string
}

func NewSgpUser() User {
	return User{
		Region: "SGP",
	}
}

func (u *User) LoginTimes() int {
	return 5
}

func (u *User) OrderTimes() int {
	return 1
}

func (u *User) CommentTimes() int {
	return 1
}

func (u *User) BrowseTimes() int {
	return 15
}

func (u *User) PostTimes() int {
	return 3
}

func (u *User) FollowTimes() int {
	return 8
}

func (u *User) InvitationTimes() int {
	return 2
}
