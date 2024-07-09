package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenData struct {
	UserId    string `json:"userId"`
	BalanceId string `json:"balanceId"`
}

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(userId string, balanceId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["userId"] = userId
	claims["balanceId"] = balanceId
	claims["exp"] = time.Now().Add(time.Hour * 7 * 24).Unix()

	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func ParseToken(tokenStr string) (TokenData, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return TokenData{}, err
	}

	return TokenData{UserId: claims["userId"].(string), BalanceId: claims["balanceId"].(string)}, nil
}
