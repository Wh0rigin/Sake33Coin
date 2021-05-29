package main

import(
	"fmt"
	"crypto/sha256"
)

func GetSHA256HashCode(msg string) string {
	hashcode := sha256.New()
	hashcode.Write([]byte(msg))
	return fmt.Sprintf("%x",hashcode.Sum(nil))
}
