package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Uid        string `gorm:"size:64;unique;not null"`
	Account    string `gorm:"size:128"`
	Password   string `gorm:"size:64"`
	OpenId     string `gorm:"size:64"`
	LoginTime  time.Time
	LogoutTime time.Time

	gorm.Model
}

type UserBaseInfo struct {
	Uid     string `gorm:"size:64;unique;not null"`
	Name    string `gorm:"size:64"`
	HeadUrl string `gorm:"size:128"`
	Star    int32
	Exp     int32
	LvChao  string `gorm:"size:1024"`
	Diamond int32

	gorm.Model
}

type UserExtendInfo struct {
	Uid      string `gorm:"size:64;unique;not null"`
	Level    int32
	Scene    int32
	Hair     int32
	Coat     int32
	Trouser  int32
	Neck     int32
	Shoe     int32
	MaxCloth int32

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

// ConfigCloth 衣服配置
type ConfigCloth struct {
	Name  string `gorm:"size:16"`
	Icon  string `gorm:"size:256"`
	Cost  string `gorm:"size:1024"`
	Level int32
	Type  int32
	Exp   int32

	gorm.Model
}

// ConfigScene 场景配置
type ConfigScene struct {
	No    int32
	Name  string `gorm:"size:16"`
	Icon  string `gorm:"size:256"`
	Level int32
	Star  int32

	gorm.Model
}

// ConfigLevel 咔位配置
type ConfigLevel struct {
	No    int32
	Name  string `gorm:"size:16"`
	Icon  string `gorm:"size:256"`
	Level int32
	Star  int32

	gorm.Model
}

// ConfigSign 签到配置
type ConfigSign struct {
	No  int32
	Day int32
	Num int32

	gorm.Model
}
