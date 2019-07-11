package db

import (
	"fmt"
	"server/entry"
	"time"

	"github.com/name5566/leaf/log"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// DBDriver 数据库地址
	DBDriver = "root:A845240287a@tcp(rm-wz9sw694mi8020vigo.mysql.rds.aliyuncs.com:3306)/wxgame?charset=utf8&&parseTime=true"
)

// ConnectDB 连接数据库
func (m *Module) ConnectDB() {
	db, err := gorm.Open("mysql", DBDriver)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	// defer db.Close()
	m.db = db
}

// PersistentData 数据库固化
func (m *Module) PersistentData() {
	m.PersistentUser()
	m.PersistentSign()
	m.PersistentCloth()
}

// PersistentUser 固化用户信息
func (m *Module) PersistentUser() {
	cnt := 0
	for _, player := range m.players {
		user := User{
			Uid:        player.UserId,
			Account:    player.Account,
			Password:   player.Password,
			LoginTime:  player.LoginTime,
			LogoutTime: player.LogoutTime,
		}

		var oldUser User
		m.db.Where("uid = ?", user.Uid).First(&oldUser)
		if user.Uid != oldUser.Uid {
			m.db.Create(&user)
		} else {
			m.db.Model(&user).Where("uid = ?", user.Uid).Updates(user)
		}

		userBaseInfo := UserBaseInfo{
			Uid:     player.UserId,
			Name:    player.Name,
			Star:    player.Star,
			LvChao:  player.LvChao,
			Diamond: player.Diamond,
		}

		var oldUserInfo UserBaseInfo
		m.db.Where("uid = ?", userBaseInfo.Uid).First(&oldUserInfo)
		if userBaseInfo.Uid != oldUserInfo.Uid {
			m.db.Create(&userBaseInfo)
		} else {
			m.db.Model(&userBaseInfo).Where("uid = ?", userBaseInfo.Uid).Updates(userBaseInfo)
		}

		userExtendInfo := UserExtendInfo{
			Uid:     player.UserId,
			Level:   player.Level,
			Scene:   player.Scene,
			Hair:    player.Hair,
			Coat:    player.Coat,
			Trouser: player.Trouser,
			Neck:    player.Neck,
		}

		var oldExtendInfo UserExtendInfo
		m.db.Where("uid = ?", userExtendInfo.Uid).First(&oldExtendInfo)
		if userExtendInfo.Uid != oldExtendInfo.Uid {
			m.db.Create(&userExtendInfo)
		} else {
			m.db.Model(&userExtendInfo).Where("uid = ?", userExtendInfo.Uid).Updates(userExtendInfo)
		}

		cnt++
	}
	log.Debug("persistent user %v ", cnt)
}

// PersistentSign 固化签到信息
func (m *Module) PersistentSign() {
	for uid, v := range m.signs {
		for day, v := range v {
			sign := UserSignInfo{
				Uid:      uid,
				Day:      day,
				SignTime: v,
			}

			var oldSign UserSignInfo
			m.db.Where("uid = ? and day = ?", sign.Uid, sign.Day).First(&oldSign)
			if sign.Uid != oldSign.Uid {
				m.db.Create(&sign)
			} else {
				m.db.Model(&sign).Where("uid = ? and day = ?", sign.Uid, sign.Day).Updates(sign)
			}
		}
	}
	for k := range m.signs {
		delete(m.signs, k)
	}
}

// PersistentCloth 固化衣服快照
func (m *Module) PersistentCloth() {
	for uid, v := range m.cloths {
		cloth := UserClothInfo{
			Uid:  uid,
			Snap: v,
		}

		var oldCloth UserClothInfo
		m.db.Where("uid = ?", cloth.Uid).First(&oldCloth)
		if cloth.Uid != oldCloth.Uid {
			m.db.Create(&cloth)
		} else {
			m.db.Model(&cloth).Where("uid = ?", cloth.Uid).Updates(cloth)
		}
	}
	for k := range m.cloths {
		delete(m.cloths, k)
	}
}

// CreateTables 创建表
func (m *Module) CreateTables() {
	if !m.db.HasTable(&User{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserBaseInfo{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserBaseInfo{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserExtendInfo{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserExtendInfo{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserSignInfo{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserSignInfo{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserClothInfo{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserClothInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

// LoadFromDB 加载数据
func (m *Module) LoadFromDB() {
	log.Debug("loading data from db start ...")
	m.LoadPlayer()
	m.LoadCloth()
	m.LoadSign()
	log.Debug("loading data from db end ...")
}

// LoadPlayer 加载用户信息
func (m *Module) LoadPlayer() {
	var users []*User
	m.db.Find(&users)
	tempPlayers := make(map[string]*entry.Player)
	for _, user := range users {
		player := new(entry.Player)
		player.UserId = user.Uid
		player.Account = user.Account
		player.Password = user.Password
		player.LoginTime = user.LoginTime
		player.LogoutTime = user.LogoutTime
		player.CreateTime = user.CreatedAt
		m.SavePlayer(player)

		tempPlayers[user.Uid] = player
	}

	var userBaseInfos []*UserBaseInfo
	m.db.Find(&userBaseInfos)
	for _, baseInfo := range userBaseInfos {
		if tempPlayers[baseInfo.Uid] == nil {
			continue
		}
		tempPlayers[baseInfo.Uid].Name = baseInfo.Name
		tempPlayers[baseInfo.Uid].Star = baseInfo.Star
		tempPlayers[baseInfo.Uid].Diamond = baseInfo.Diamond
		tempPlayers[baseInfo.Uid].LvChao = baseInfo.LvChao
	}

	tempPlayers = nil

	log.Debug("load players  db %v  mem %v", len(users), len(m.players))
}

// LoadCloth 加载衣服快照
func (m *Module) LoadCloth() {
	var clothInfos []*UserClothInfo
	m.db.Find(&clothInfos)
	for _, cloth := range clothInfos {
		// log.Debug("LoadCloth %v %v", cloth.Uid, cloth.Snap)
		m.SaveCloth(cloth.Uid, cloth.Snap)
	}
	log.Debug("load cloths  db %v  mem %v", len(clothInfos), len(m.cloths))
}

// LoadSign 加载签到
func (m *Module) LoadSign() {
	var signInfos []*UserSignInfo
	m.db.Find(&signInfos)
	for _, sign := range signInfos {
		if dayDic, ok := m.signs[sign.Uid]; ok {
			dayDic[sign.Day] = sign.SignTime
		} else {
			dayDic := make(map[string]time.Time)
			dayDic[sign.Day] = sign.SignTime
			m.signs[sign.Uid] = dayDic
		}
	}

	for k, v := range m.signs {
		for day, t := range v {
			log.Debug("LoadSign  %v %v %v", k, day, t)
		}
	}

	log.Debug("Load Signs  db %v  mem %v", len(signInfos), len(m.signs))
}
