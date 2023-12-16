package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"shop_dev/common"
	"shop_dev/config"
	"shop_dev/dto/request"
	"shop_dev/entity"
	"shop_dev/repository"
)

type IUserService interface {
	CreateNewUser(ctx context.Context, request *request.CreateUserRequest) (*entity.User, error)
	LoginUser(ctx context.Context, request *request.LoginRequest) (string, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type UserService struct {
	config              *config.Config
	databaseTransaction IDatabaseTransaction
	userRepository      repository.IUserRepository
}

func NewUserService(
	config *config.Config,
	databaseTransaction IDatabaseTransaction,
	userRepository repository.IUserRepository) IUserService {
	return &UserService{
		config:              config,
		databaseTransaction: databaseTransaction,
		userRepository:      userRepository}
}

func (u UserService) CreateNewUser(ctx context.Context, request *request.CreateUserRequest) (*entity.User, error) {
	//check existed user
	existedUser, err := u.GetUserByUsername(ctx, request.Username)
	if existedUser != nil {
		log.Error("Username is existed")
		return nil, err
	}
	hashPassword, err := common.HashPassword(request.Password)
	if err != nil {
		log.Error("Hash password fail")
		return nil, err
	}
	newUser := request.ToEntityUser()
	newUser.Password = hashPassword
	userRsp, err := u.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}
	log.Info("Create new user successfully")
	return userRsp, nil
}
func (u UserService) LoginUser(ctx context.Context, request *request.LoginRequest) (string, error) {
	existedUser, err := u.GetUserByUsername(ctx, request.Username)
	if err != nil || existedUser == nil {
		log.Error("Invalid login")
		return "", err
	}
	ok := common.CheckPasswordHash(request.Password, existedUser.Password)
	if !ok {
		log.Error("Invalid password")
		return "", err
	}
	token, err := common.GenerateJwtToken(u.config, int(existedUser.ID))
	if err != nil {
		log.Error("Cannot generate token ", err)
		return "", err
	}
	return token, nil
}
func (u UserService) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := u.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
