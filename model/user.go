package model

type User struct {
	Id  string
	Key string
}

type Anonym struct {
	Id string
}

func NewUser() *User {
	return &User{}
}

func (u User) ToAnonym() *Anonym {
	return &Anonym{Id: u.Id}
}
