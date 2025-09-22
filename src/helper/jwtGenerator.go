package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte("mySecretKey")

type jwtPasswordClaims struct {
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func GenerateJWT(stringToSign string, userEmail string) (string, error) {
	claim := jwtPasswordClaims{
		Password: stringToSign,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: userEmail,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString((secretkey))

	if(err != nil){
		return err.Error(), err
	}

	return signedToken, nil
}

func VeifyToken(signedToken string, stringToVerify string) (bool, error) {
	claims := &jwtPasswordClaims{}

	token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretkey, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("Invalid token")
	}

	if claims.Password != stringToVerify {
		return false, fmt.Errorf("Password mismatch")
	}

	return true, nil
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