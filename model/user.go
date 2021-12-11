package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"varchar(20);not null;" form:"name" binding:"required"`
	Telephone string `gorm:"varchar(20);not null;unique" form:"telephone" binding:"required"`
	Password  string `gorm:"size:255;not null" form:"password" binding:"required"`
	PublicKey string `gorm:"size255;unique" `
	Kind      string    `gorm:"not null" form:"kind" binding:"required"`
}
