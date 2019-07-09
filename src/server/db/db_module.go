package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Uid        string `gorm:"size:64;unique;not null"`
	Account    string `gorm:"size:128"`
	Password   string `gorm:"size:64"`
	LoginTime  time.Time
	LogoutTime time.Time

	gorm.Model
}

type UserBaseInfo struct {
	Uid        string `gorm:"size:64;unique;not null"`
	Name       string `gorm:"size:64"`
	Star       int32
	LvChao     int32
	Diamond    int32

	gorm.Model
}