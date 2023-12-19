package common

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"shop_dev/config"
	"time"
)

type JWTClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(config *config.Config, ID int) (string, error) {
	expTime := time.Now().Add(time.Duration(config.JwtConfig.ExpTime) * time.Minute)
	claims := &JWTClaims{
		ID: ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	//privateKey, err := os.ReadFile("jwtRS256.rsa")
	//key, err := jwt.ParseRSAPublicKeyFromPEM(privateKey)
	//key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.JwtConfig.SecretKey))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtConfig.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ClaimToken(tokenString string, config *config.Config) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtConfig.SecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok || (claims.ExpiresAt.Unix() < time.Now().Unix()) {
		return nil, errors.New("")
	}
	return claims, nil
}
