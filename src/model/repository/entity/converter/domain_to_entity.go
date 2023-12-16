package converter

import (
	"github.com/vhtor/metaifrn-simulados-api/src/model"
	"github.com/vhtor/metaifrn-simulados-api/src/model/repository/entity"
)

func DomainToEntity(domain model.UserInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}