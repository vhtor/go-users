package service

import (
	logger "github.com/vhtor/metaifrn-simulados-api/src/configuration/log"
	resterr "github.com/vhtor/metaifrn-simulados-api/src/configuration/rest_err"
	"github.com/vhtor/metaifrn-simulados-api/src/model"

	"go.uber.org/zap"
)

func (user *userService) CreateUser(
	userDomain model.UserInterface,
) (model.UserInterface, *resterr.RestErr) {
	logger.Info("Init create domain User", zap.String("journey", "create domain User"))

	userDomain.EncryptPassword()

	userRepository, err := user.userRepository.CreateUser(userDomain)
	if err != nil {
		
	}

	return userRepository, err
}