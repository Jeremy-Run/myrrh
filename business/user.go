package business

type User struct {
	Region string
}

func (u *User) LoginTimes() int {
	return 5
}

func NewSgpUser() User {
	return User{
		Region: "SGP",
	}
}
