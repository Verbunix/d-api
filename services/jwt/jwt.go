package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

//Creating Access Token
func CreateToken(userId uint) (error, string) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 60 * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return err, ""
	}

	return nil, token
}

func CreateShaHash(str string) string {
	key := []byte(str)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(os.Getenv("PASSWORD_SALT")))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Validate, and return a token.
// Will receive the parsed token and should return the key for validating.
func VerifyToken(tokenString string) (error, *jwt.Token) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return err, nil
	}

	return nil, token
}

func TokenValidation(tokenString string) error {
	err, token := VerifyToken(tokenString)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return err
	}

	return nil
}
