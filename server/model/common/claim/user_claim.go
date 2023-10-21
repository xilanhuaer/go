package claim

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}
