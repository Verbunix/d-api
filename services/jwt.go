package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

//Creating Access Token
func CreateToken(userId uint) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 60 * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func CreateShaHash(str string) string {
	key := []byte(str)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(os.Getenv("PASSWORD_SALT")))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
