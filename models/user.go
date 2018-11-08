package models

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"-"`
}

func (this *User) GenPwd(password []byte) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	this.Password = string(hash)
}

func (this *User) CheckPwd(password []byte) bool {
	byteHash := []byte(this.Password)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		return false
	}
	return true
}
