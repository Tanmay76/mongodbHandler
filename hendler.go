package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HandlerInsertData(c *gin.Context) {
	var buser Bloguser
	if err := c.ShouldBindJSON(&buser); err != nil {
		log.Printf("Error while binding:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	buser.Passwordh, _ = bcrypt.GenerateFromPassword([]byte(buser.Password), 14)
	buser.Password = "hellow"
	log.Printf("data:%v", buser)
	Id, err := Insertdata(&buser)
	if err != nil {
		log.Printf("Error while inserting data:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"ID": Id})
}

func HandlerShowdatabyID(c *gin.Context) {
	id := c.Param("id")
	idn, _ := strconv.ParseInt(id, 10, 32)
	buser, err := FindOneData(idn)
	if err != nil {
		log.Printf("Cannot find a data:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": buser})
}

func HandlerShowdata(c *gin.Context) {
	buser, err := GetAlldate()
	if err != nil {
		log.Printf("Cannot find a data:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": buser})
}

func HandlerLogin(c *gin.Context) {
	var us User
	err := c.ShouldBindJSON(&us)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": us})
		return
	}
	bus, _ := FindOneData(int64(us.UserID))
	if err := bcrypt.CompareHashAndPassword(bus.Passwordh, []byte(us.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Massage": "Incorrect passsword"})
		return
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    bus.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		log.Printf("%v", err.Error())
		c.String(http.StatusBadRequest, err.Error())
	}
	// c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	c.Header("Autherization", token)
	c.JSON(http.StatusOK, gin.H{"auth": token})
}
