package converter

import (
	"github.com/vhtor/metaifrn-simulados-api/src/model"
	"github.com/vhtor/metaifrn-simulados-api/src/model/repository/entity"
)

func EntityToDomain(entity entity.UserEntity) model.UserInterface  {
	domain := model.NewUser(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}