package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"shop_dev/dto/request"
	"shop_dev/entity"
	"shop_dev/repository"
)

type IUserService interface {
	CreateNewUser(ctx context.Context, request *request.CreateUserRequest) (*entity.User, error)
}

type UserService struct {
	databaseTransaction IDatabaseTransaction
	userRepository      repository.IUserRepository
}

func NewUserService(databaseTransaction IDatabaseTransaction,
	userRepository repository.IUserRepository) IUserService {
	return &UserService{
		databaseTransaction: databaseTransaction,
		userRepository:      userRepository}
}

func (u UserService) CreateNewUser(ctx context.Context, request *request.CreateUserRequest) (*entity.User, error) {
	tx := u.databaseTransaction.StartTransaction()
	newUser := request.ToEntityUser()
	userRsp, err := u.userRepository.CreateUser(ctx, tx, newUser)
	if err != nil {
		return nil, err
	}
	err = u.databaseTransaction.CommitTransaction(tx)
	if err != nil {
		log.Error("Cannot commit transaction: ", err)
		return nil, err
	}
	return userRsp, nil
}
