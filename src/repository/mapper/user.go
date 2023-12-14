package mapper

import (
	"shop_dev/entity"
	"shop_dev/repository/model"
)

func UserModelToEntity(m *model.User) *entity.User {
	return &entity.User{
		ID:        m.ID,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Username:  m.Username,
		Password:  m.Password,
		CreatedAt: m.CreatedAt.Unix(),
		UpdatedAt: m.UpdatedAt.Unix(),
	}
}

func UserEntityToModel(e *entity.User) *model.User {
	return &model.User{
		BaseModel: model.BaseModel{
			ID: e.ID,
		},
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Username:  e.Username,
		Password:  e.Password,
	}
}
