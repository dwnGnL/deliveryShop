package user

import (
	"foodShop/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func Authenticate(userName string, password string, c *gin.Context) (string, bool) {
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
