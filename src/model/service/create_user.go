package service

import (
	"fmt"

	logger "github.com/vhtor/metaifrn-simulados-api/src/configuration/log"
	resterr "github.com/vhtor/metaifrn-simulados-api/src/configuration/rest_err"
	model "github.com/vhtor/metaifrn-simulados-api/src/model"

	"go.uber.org/zap"
)

func (user *userService) CreateUser(
	userDomain model.UserInterface,
) *resterr.RestErr {
	logger.Info("Init create domain User", zap.String("journey", "create domain User"))

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())

	return nil
}