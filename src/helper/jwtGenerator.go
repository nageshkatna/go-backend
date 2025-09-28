package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte("mySecretKey")

type jwtPasswordClaims struct {
	Password string `json:"password"`
	RoleId uint `json:roleId`
	jwt.RegisteredClaims
}

func GenerateJWT(stringToSign string, userId string, roleId uint) (string, error) {
	claim := jwtPasswordClaims{
		Password: stringToSign,
		RoleId: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: userId,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString((secretkey))

	if(err != nil){
		return err.Error(), err
	}

	return signedToken, nil
}

func VeifyToken(signedToken string) (bool, error, *jwtPasswordClaims) {
	claims := &jwtPasswordClaims{}

	token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretkey, nil
	})

	if err != nil {
		return false, err, claims
	}

	if !token.Valid {
		return false, fmt.Errorf("Invalid token"), claims
	}

	return true, nil, claims
}

// func main() {
// 	token, err := GenerateJWT("myPassword123", "email.com")
// 	if err != nil {
// 		fmt.Println("Error generating token:", err)
// 		return
// 	}

// 	isValid, err := VeifyToken(token, "myPassword123")
// 	if err != nil {
// 		fmt.Println("Error verifying token:", err)
// 		return
// 	}

// 	fmt.Println("Is token valid?", isValid)
// }