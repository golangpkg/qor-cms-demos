package models

import "github.com/jinzhu/gorm"

type UserInfo struct {
	gorm.Model
	UserName string
}

func (userInfo UserInfo) DisplayName() string {
	return userInfo.UserName
}