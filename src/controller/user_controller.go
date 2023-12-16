package controller

import (
	"github.com/gin-gonic/gin"
	"shop_dev/common"
	request "shop_dev/dto/request"
	"shop_dev/dto/response"
	"shop_dev/service"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}

}
func (c *UserController) CreateNewUser(ctx *gin.Context) {
	var request request.CreateUserRequest
	if err := ctx.BindJSON(&request); err != nil {
		common.AbortErrorHandleCustomMessage(ctx, common.CannotBindJson, err.Error())
		return
	}
	newUser, err := c.userService.CreateNewUser(ctx, &request)
	if err != nil {
		common.AbortErrorHandleCustomMessage(ctx, common.CannotCreateNewUser, err.Error())
		return
	}
	common.SuccessfulHandle(ctx, response.ToUserResponse(newUser))
}
func (c *UserController) LoginUser(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		common.AbortErrorHandleCustomMessage(ctx, common.CannotBindJson, err.Error())
		return
	}
	token, err := c.userService.LoginUser(ctx, &req)
	if err != nil || len(token) == 0 {
		common.AbortErrorHandleCustomMessage(ctx, common.InvalidLogin, err.Error())
		return
	}
	common.SuccessfulHandle(ctx, response.ToLoginResponse(token))
}
