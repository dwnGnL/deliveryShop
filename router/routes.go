package router

import (
	"context"
	"fmt"
	"foodShop/dbr"
	"foodShop/middleware"
	"foodShop/router/store"
	"foodShop/router/user"
	"foodShop/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB
var config *utils.Config

func Init() {
	config = utils.ReadConfig()

	log.SetOutput(gin.DefaultWriter)

	logger := logrus.New()
	logger.Level = logrus.TraceLevel
	logger.SetOutput(gin.DefaultWriter)
	log.Println(config.DBUri)
	db = dbr.Init(config.DBUri, logger)
	defer db.Close()

	jwtMiddleware := &middleware.GinJWTMiddleware{
		Realm:         config.Realm,
		Key:           []byte(config.Key),
		Timeout:       time.Second * time.Duration(config.Timeout),
		MaxRefresh:    time.Hour * 24,
		Authenticator: user.Authenticate,
		PayloadFunc:   user.Payload,
	}

	// Initialize default gin router
	defaultRouter := gin.Default()

	defaultRouter.Use(middleware.Logger(logger), gin.Recovery())
	defaultRouter.Use(middleware.CORSMiddleware())
	defaultRouter.POST("/v1/newUser", user.CreateNewUser)
	defaultRouter.POST("/v1/login", jwtMiddleware.LoginHandler)
	defaultRouter.GET("/v1/store/count", store.GetCountOfStore)
	defaultRouter.POST("/v1/store", store.GetStores)

	defaultRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.PortRun),
		WriteTimeout:   60 * time.Second,
		ReadTimeout:    40 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        defaultRouter,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 20 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
