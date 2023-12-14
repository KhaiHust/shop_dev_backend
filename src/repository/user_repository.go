package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"shop_dev/entity"
	"shop_dev/repository/mapper"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, tx *gorm.DB, user *entity.User) (*entity.User, error)
}
type UserRepository struct {
	base
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{base: base{db: db}}
}
func (u UserRepository) CreateUser(ctx context.Context, tx *gorm.DB, user *entity.User) (*entity.User, error) {
	userModel := mapper.UserEntityToModel(user)
	if err := tx.WithContext(ctx).Create(userModel).Error; err != nil {
		log.Error("Can not create new user")
		return nil, err
	}
	return mapper.UserModelToEntity(userModel), nil
}
