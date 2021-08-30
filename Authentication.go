package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication(c *gin.Context) {
	// cookie, err := c.Cookie("jwt")

	// if err != nil {
	// 	log.Printf("There is no cookie present:%v", err)
	// 	c.JSON(http.StatusNotFound, gin.H{"err": "No cookis prestent"})
	// 	c.Abort()
	// 	return
	// }
	auth := c.GetHeader("Authorization")
	log.Println(auth)
	token, err := jwt.ParseWithClaims(auth, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) { return []byte(SecretKey), nil })
	if err != nil {
		log.Printf("Erro in token:%v", err)
		c.JSON(http.StatusNotFound, gin.H{"err": "token error token is not found"})
		c.Abort()
		return
	}
	claims := token.Claims.(*jwt.StandardClaims)
	us, err := FindOnebyemail(claims.Issuer)
	log.Printf("Issuer:%v", us.Email)
	if err != nil {
		log.Printf("Not valid user:%v", err)
		c.JSON(http.StatusNotFound, gin.H{"err": "Not valid user login"})
		c.Abort()
		return
	}
	c.Next()
}
