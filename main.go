package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {
	r := gin.Default()
	gin.DefaultWriter = colorable.NewColorableStdout()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 24 * time.Hour,
	}))
	us := r.Group("/user")
	us.POST("/login", HandlerLogin)
	us.POST("/signup", HandlerInsertData)
	auth := r.Group("/userDetail", Authentication)
	auth.GET("/:id", HandlerShowdatabyID)
	auth.GET("/Alluser", HandlerShowdata)
	r.Run(URLS)
}
