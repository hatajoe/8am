package model

type User struct {
	id    string
	email string
}

func NewUser(id, email string) *User {
	return &User{
		id:    id,
		email: email,
	}
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetEmail() string {
	return u.email
}
