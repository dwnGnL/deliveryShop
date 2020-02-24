package user

import (
	"foodShop/dbr"
	"foodShop/utils"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

var Db *gorm.DB

func CreateNewUser(c *gin.Context) {
	Db = dbr.GetDb()
	var createdUser CreateUser
	var user User
	var userForCheck User
	if error := c.BindJSON(&createdUser); error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Fields are empty." + error.Error()})
		return
	}

	if createdUser.UserName == "" || createdUser.Password == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "login or password are empty"})
		return
	}

	salt, pass := utils.HashPassword(createdUser.Password)
	user.Username = createdUser.UserName
	user.FullName = createdUser.FIO
	user.Password = pass
	user.PhoneNumber = createdUser.PhoneNumber
	user.Salt = salt
	user.Role = 1
	count := 0
	Db.Where("username =?", user.Username).Find(&userForCheck).Count(&count)
	if count != 0 {
		c.JSON(401, gin.H{"message": "user allready created"})
		return
	}

	if err := Db.Create(&user).Scan(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create user." + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user successfully created"})
}
