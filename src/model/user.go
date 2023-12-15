package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type user struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}

type UserInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(string)

	GetJSONValue() (string, error)

	EncryptPassword()
}

func NewUser(
	email, password, name string,
	age int8,
) UserInterface {
	return &user{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (user *user) GetEmail() string {
	return user.Email
}

func (user *user) GetPassword() string {
	return user.Password
}

func (user *user) GetName() string {
	return user.Name
}

func (user *user) GetAge() int8 {
	return user.Age
}

func (user *user) SetID(id string) {
	user.ID = id
}

func (user *user) GetJSONValue() (string, error) {
	json, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(json), nil
}

func (user *user) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(user.Password))

	// Exchanging original user password for the encrypted version
	// to compare it with the encrypted database password
	user.Password = hex.EncodeToString(hash.Sum(nil))
}
