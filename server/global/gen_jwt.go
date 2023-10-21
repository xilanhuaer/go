package global

import (
	"errors"
	"interface/model/common/claim"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Path: go/server/global/gen_jwt.go

// 生成Jwt
func GenJwt(userId uint, userName string) (string, error) {
	claims := claim.UserClaim{
		UserId:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "interface",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(UserSecret)
	return tokenString, err
}

// 解析Jwt
func ParseJwt(tokenString string) (*claim.UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claim.UserClaim{}, Security())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*claim.UserClaim); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("couldn't handle this token")
	}
}

func Security() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return UserSecret, nil
	}
}
