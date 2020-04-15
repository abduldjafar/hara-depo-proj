package user

import (
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
)

func Finduser(db *gorm.DB, hp string) error {
	account := &mobile.UserOtlet{}
	err := db.Where("Hp = ?", hp).First(account).Error
	return err
}

func FindUserStatus(db *gorm.DB, hp string, status string) error {
	account := &mobile.UserOtlet{}
	err := db.Where("status = ? AND hp = ?", status, hp).First(account).Error
	return err
}
