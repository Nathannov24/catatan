package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FullName string `gorm:"type:varchar(255)" json:"fullname" form:"fullname"`
	Username string `gorm:"type:varchar(100);unique;not null" json:"username" form:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
	// Orders         []Orders         `gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Products       []Products       `gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Shopping_Carts []Shopping_Carts `gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
