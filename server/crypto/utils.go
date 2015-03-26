package crypto

import (
  "crypto/sha256"
  "encoding/base64"
  "github.com/dgrijalva/jwt-go"
  "time"
)

var privateKey = []byte("CHANGE ME")

func EncodePassword(value string) string {
  hash := sha256.New()
  hash.Write([]byte(value))
  return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func CreateToken(data map[string]interface{}) string {
  token := jwt.New(jwt.GetSigningMethod("HS256"))
  token.Claims["data"] = data
  token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
  tokenString, err := token.SignedString(privateKey)
  if err != nil {
    return ""
  }
  return tokenString
}