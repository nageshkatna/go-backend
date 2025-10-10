package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte("mySecretKey")

type jwtPasswordClaims struct {
	jwt.RegisteredClaims
	Password string `json:"password"`
	RoleId   uint   `json:"roleId"`
}

func GenerateJWT(stringToSign string, userId string, roleId uint) (string, error) {
	claim := jwtPasswordClaims{
		Password: stringToSign,
		RoleId:   roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    userId,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString((secretkey))

	if err != nil {
		return err.Error(), err
	}

	return signedToken, nil
}

func VeifyToken(signedToken string) (bool, *jwtPasswordClaims, error) {
	claims := &jwtPasswordClaims{}

	token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretkey, nil
	})

	if err != nil {
		return false, claims, err
	}

	if !token.Valid {
		return false, claims, fmt.Errorf("invalid token")
	}

	return true, claims, nil
}
