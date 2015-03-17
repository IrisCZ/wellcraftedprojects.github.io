package crypto

import (
  "crypto/sha256"
  "encoding/base64"
)

func EncodePassword(value string) string {
    hash := sha256.New()
    hash.Write([]byte(value))
    return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
