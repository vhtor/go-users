package convert

import (
	"github.com/vhtor/metaifrn-simulados-api/src/controller/model/response"
	"github.com/vhtor/metaifrn-simulados-api/src/model"
)

func ConvertUserDomainToResponse(userDomain model.UserInterface) response.UserResponse {
	return response.UserResponse{
		ID: "",
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}