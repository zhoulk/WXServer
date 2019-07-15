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
	Uid     string `gorm:"size:64;unique;not null"`
	Name    string `gorm:"size:64"`
	headUrl string `gorm:"size:128"`
	Star    int32
	Exp     int32
	LvChao  string `gorm:"size:1024"`
	Diamond int32

	gorm.Model
}

type UserExtendInfo struct {
	Uid     string `gorm:"size:64;unique;not null"`
	Level   int32
	Scene   int32
	Hair    int32
	Coat    int32
	Trouser int32
	Neck    int32
	Shoe    int32

	gorm.Model
}

type UserSignInfo struct {
	Uid      string `gorm:"size:64;not null"`
	Day      string `gorm:"size:16"`
	SignTime time.Time

	gorm.Model
}

// UserClothInfo 用户衣服合成信息
type UserClothInfo struct {
	Uid  string `gorm:"size:64;unique;not null"`
	Snap string `gorm:"size:1024"`

	gorm.Model
}

// UserSnapInfo 最近一次登出时的数据
type UserSnapInfo struct {
	Uid    string `gorm:"size:64;unique;not null"`
	LvChao string `gorm:"size:1024"`

	gorm.Model
}
