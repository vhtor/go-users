package service

import (
	resterr "github.com/vhtor/metaifrn-simulados-api/src/configuration/rest_err"
	"github.com/vhtor/metaifrn-simulados-api/src/model"
)

func NewUserService() UserService {
	return &userService{}
}

type userService struct {
}

type UserService interface {
	CreateUser(model.UserInterface) *resterr.RestErr
	UpdateUser(string, model.UserInterface) *resterr.RestErr
	FindUser(string) (*model.UserInterface, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}
