package model

import (
	"crypto/md5"
	"encoding/hex"
)

type user struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

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

func (user *user) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(user.password))

	// Exchanging original user password for the encrypted version
	// to compare it with the encrypted database password
	user.password = hex.EncodeToString(hash.Sum(nil))
}
