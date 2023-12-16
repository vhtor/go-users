package model

type user struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (user *user) GetID() string {
	return user.id
} 

func (user *user) SetID(id string) {
	user.id = id
}

func (user *user) GetEmail() string {
	return user.email
}

func (user *user) GetPassword() string {
	return user.password
}

func (user *user) GetName() string {
	return user.name
}

func (user *user) GetAge() int8 {
	return user.age
}
