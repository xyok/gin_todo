package models

import (
	"gin_todo/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)
var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, setting.DatabaseSetting.Host)

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}


func GetDB() *gorm.DB{
	return db
}


func Migrate(){
	db.AutoMigrate(&User{},&Todo{})

	db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")
}