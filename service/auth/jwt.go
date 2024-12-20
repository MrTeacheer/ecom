package auth

import (
	"strconv"
	"time"

	"github.com/MrTeacheer/ecom/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, user_id int) (string, error) {
	exparation := time.Second * time.Duration(config.Envs.JWTexparation)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(user_id),
		"expiredAt": time.Now().Add(exparation).Unix(),
	})
	token_string, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return token_string, nil

}
