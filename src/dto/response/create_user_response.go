package response

import "shop_dev/entity"

type CreateUserResponse struct {
	ID        uint
	FirstName string
	LastName  string
	Username  string
	CreatedAt int64
	UpdatedAt int64
}

func ToUserResponse(user *entity.User) *CreateUserResponse {
	return &CreateUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
