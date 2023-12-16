package response

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func ToLoginResponse(token string) *LoginResponse {
	return &LoginResponse{AccessToken: token}
}
