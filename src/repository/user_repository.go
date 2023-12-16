package repository

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"shop_dev/common"
	"shop_dev/entity"
	"shop_dev/repository/mapper"
	"shop_dev/repository/model"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserById(ctx context.Context, ID int) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}
type UserRepository struct {
	base
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{base: base{db: db}}
}
func (u UserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	userModel := mapper.UserEntityToModel(user)
	if err := u.db.WithContext(ctx).Create(userModel).Error; err != nil {
		log.Error("Can not create new user with err: ", err.Error())
		return nil, err
	}
	return mapper.UserModelToEntity(userModel), nil
}
func (u UserRepository) GetUserById(ctx context.Context, ID int) (*entity.User, error) {
	var userModel model.User
	if err := u.db.WithContext(ctx).Where("ID = ?", ID).First(&userModel).Error; err != nil {
		log.Error("Can not find user with ID: ", ID, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(common.RECORD_NOT_FOUND)
		}
		return nil, err
	}
	return mapper.UserModelToEntity(&userModel), nil
}
func (u UserRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var userModel model.User
	if err := u.db.WithContext(ctx).Where("username = ?", username).First(&userModel).Error; err != nil {
		log.Error("Can not find user with username: ", username, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(common.RECORD_NOT_FOUND)
		}
		return nil, err
	}
	return mapper.UserModelToEntity(&userModel), nil
}
