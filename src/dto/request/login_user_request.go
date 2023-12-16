package request

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"gte=6,lte=64,required"`
}
