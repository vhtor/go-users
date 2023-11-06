package service

import (
	resterr "github.com/vhtor/metaifrn-simulados-api/src/configuration/rest_err"
	"github.com/vhtor/metaifrn-simulados-api/src/model"
)

func (user *userService) UpdateUser(
	id string,
	userDomain model.UserInterface,
) *resterr.RestErr {
	return nil
}
