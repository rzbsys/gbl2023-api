package user

import (
	"gbl-api/data"
	"gorm.io/gorm"
	"log"
)

func IsUserExist(uid string) bool {
	db := data.GetDatabase()
	var user User
	err := db.Where("uid = ?", uid).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err == nil {
		return true
	} else {
		log.Println(err)
		return false
	}
}
