package store

import (
	"foodShop/dbr"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func GetCountOfStore(c *gin.Context) {
	Db = dbr.GetDb()
	var count int

	Db.Model(&TStore{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func GetStores(c *gin.Context) {
	Db = dbr.GetDb()
	var req CountOfReturn
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error})
		return
	}
	var store TStore
	if req.Page <= 0 || req.Limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ты че охуел?"})
		return
	}

	query := Db.Order("id desc").Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	if err := query.Find(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": store})
}
