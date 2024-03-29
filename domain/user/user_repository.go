package user

type UserRepository interface {
	Create(user *User) error
	FindAll() ([]User, error)
	Authenticate(email, password string) (*User, error)
}
