package model

type UserInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(string)

	EncryptPassword()
}

func NewUser(
	email, password, name string,
	age int8,
) UserInterface {
	return &user{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}