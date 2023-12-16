package service

import (
	resterr "github.com/vhtor/metaifrn-simulados-api/src/configuration/rest_err"
	"github.com/vhtor/metaifrn-simulados-api/src/model"
	"github.com/vhtor/metaifrn-simulados-api/src/model/repository"
)

func NewUserService(
	userRepository repository.UserRepository,
) UserService {
	return &userService{ userRepository }
}

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	CreateUser(model.UserInterface) (model.UserInterface, *resterr.RestErr)
	UpdateUser(string, model.UserInterface) *resterr.RestErr
	FindUser(string) (*model.UserInterface, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}
