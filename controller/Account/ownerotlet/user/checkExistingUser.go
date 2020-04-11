package user

import (
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model"
)

func Finduser(db *gorm.DB, hp string) error {
	account := &model.UserOtlet{}
	err := db.Where("Hp = ?", hp).First(account).Error
	return err
}

func FindUserStatus(db *gorm.DB, hp string, status string) error {
	account := &model.UserOtlet{}
	err := db.Where("status = ? AND hp = ?", status, hp).First(account).Error
	return err
}
