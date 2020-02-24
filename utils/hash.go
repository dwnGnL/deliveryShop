package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"log"
)

func CheckPasswordHash(password, salt, hash string) bool {

	var array []byte
	sha512h := sha512.New()

	array = append(array, []byte(password)...)
	array = append(array, []byte(salt)...)

	sha512h.Write(array)
	if base64.RawStdEncoding.EncodeToString(sha512h.Sum(nil)) == hash {

		return true
	}
	return false
}

func HashPassword(password string) (string, string) {

	salt := generateSalt()
	var array []byte

	sha512h := sha512.New()

	array = append(array, []byte(password)...)
	array = append(array, []byte(salt)...)

	sha512h.Write(array)

	return salt, base64.RawStdEncoding.EncodeToString(sha512h.Sum(nil))
}

func generateSalt() string {

	const SaltLength = 5
	data := make([]byte, SaltLength)
	_, err := rand.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to a string

	return base64.RawStdEncoding.EncodeToString(data[:])[:5]
}
