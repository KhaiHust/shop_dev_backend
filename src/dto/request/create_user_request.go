package request

import "shop_dev/entity"

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"gte=6,lte=64,required"`
}

func (this *CreateUserRequest) ToEntityUser() *entity.User {
	return &entity.User{
		FirstName: this.FirstName,
		LastName:  this.LastName,
		Username:  this.Username,
		Password:  this.Password,
	}
}
