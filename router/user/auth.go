package user

import (
	"fmt"
	"foodShop/dbr"
	"foodShop/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func Authenticate(userName string, password string, c *gin.Context) (string, bool) {
	Db = dbr.GetDb()
	var user User
	log.Println("--------------------->", password)
	if result := Db.Where("userName = ? and disabled = false", userName).
		Find(&user).
		Limit(1); result.Error == nil {
		if utils.CheckPasswordHash(password, user.Salt, user.Password) {
			log.Println("--------------------->", userName)
			Db.Callback().Update().Remove("gorm:update_time_stamp")

			return userName, true
		}
	}

	return "", false
}

func Payload(userName string) map[string]interface{} {
	// var userR models.Usersinroles
	var user User

	if result := Db.Table("tusers u,troles r").
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
