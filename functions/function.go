package functions

import (
	"crypto/sha256"
	"fmt"
)

func GetSHA256HashCode(msg string) string {
	hashcode := sha256.New()
	hashcode.Write([]byte(msg))
	return fmt.Sprintf("%x", hashcode.Sum(nil))
}
