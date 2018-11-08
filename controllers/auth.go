package controllers

import (
	"gin_todo/models"
	"gin_todo/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type LoginPost struct {
	Name     string `form:"name" json:"name" xml:"name"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {

	db := models.GetDB()
	var json LoginPost
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, badReq)
		return
	}

	var m models.User

	err := db.Where("name = ? ", json.Name).First(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, badReq)
		return
	}

	if !m.CheckPwd([]byte(json.Password)) {
		c.JSON(http.StatusOK, unauthorized)
		return
	}

	tokenString := util.GenToken(jwt.MapClaims{
		"id":   m.Id,
		"name": m.Name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	c.SetCookie("name", m.Name, 3600*24, "/", "", true, true)

	c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func SignUp(c *gin.Context) {

	db := models.GetDB()
	var json LoginPost
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, badReq)
		return
	}
	m := models.User{
		Name: json.Name,
	}
	m.GenPwd([]byte(json.Password))
	db.Create(&m)

	if ! db.NewRecord(m) {
		SuccessJson(c)
		return
	} else {
		c.JSON(http.StatusBadRequest, nameExist)
		return
	}

}

func ValidateToken(c *gin.Context) {
	SuccessJson(c)
}
