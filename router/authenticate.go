package router

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func authenticate(userName string, password string, c *gin.Context) (string, bool) {
	var user User
	log.Println("--------------------->", password)
	if result := db.Where("userName = ? and disabled = false", userName).
		Find(&user).
		Limit(1); result.Error == nil {
		if checkPasswordHash(password, user.Salt, user.Password) {
			log.Println("--------------------->", userName)
			db.Callback().Update().Remove("gorm:update_time_stamp")

			return userName, true
		}
	}

	return "", false
}

func checkPasswordHash(password, salt, hash string) bool {

	var array []byte
	sha512h := sha512.New()

	array = append(array, []byte(password)...)
	array = append(array, []byte(salt)...)

	sha512h.Write(array)
	log.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<", hash)
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>", base64.RawStdEncoding.EncodeToString(sha512h.Sum(nil)))
	if base64.RawStdEncoding.EncodeToString(sha512h.Sum(nil)) == hash {

		return true
	}
	return false
}

func payload(userName string) map[string]interface{} {
	// var userR models.Usersinroles
	var user User

	if result := db.Table("tusers u,troles r").
		Where("u.role = r.role_id and u.userName = ?", userName).
		Select("u.*, r.role_name as role").
		Scan(&user).Limit(1); result.Error != nil {
		fmt.Println(result.Error.Error())
		return map[string]interface{}{
			"userID":   0,
			"role":     "undefined",
			"userName": "0",
		}
	}

	log.Println("userId= ", user.ID, "\nuserName = ", user.Username, "\nrole = ", user.Role)

	return map[string]interface{}{
		"userID":   user.ID,
		"role":     user.Role,
		"userName": user.Username,
	}
}

func hashPassword(password string) (string, string) {

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
