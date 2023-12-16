package convert

import (
	"github.com/vhtor/metaifrn-simulados-api/src/controller/model/response"
	"github.com/vhtor/metaifrn-simulados-api/src/model"
)

func UserDomainToResponse(userDomain model.UserInterface) response.UserResponse {
	return response.UserResponse{
		ID: userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}