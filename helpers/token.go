package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySigninKey = []byte("secrects")

type claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func NewToken(username, role string) *claims {
	return &claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}
}

func (c *claims) Create() (sign string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	sign, err = token.SignedString(mySigninKey)

	if err != nil {
		return
	}
	return
}

func CheckToken(token string) (bool, string) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySigninKey), nil
	})

	if err != nil {
		fmt.Print(err.Error())
		return false, err.Error()
	}

	claims := tokens.Claims.(*claims)
	return tokens.Valid, claims.Role
}
