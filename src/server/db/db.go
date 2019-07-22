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
	//DBDriver = "root:A845240287a@tcp(rm-wz9sw694mi8020vigo.mysql.rds.aliyuncs.com:3306)/wxgame?charset=utf8&&parseTime=true"
	DBDriver = "root:A845240287a@tcp(rm-wz9sw694mi8020vigo.mysql.rds.aliyuncs.com:3306)/wxgame_test?charset=utf8&&parseTime=true"
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
	log.Debug("persistent start ==================================== ")
	m.PersistentUser()
	m.PersistentSign()
	m.PersistentCloth()
	m.PersistentSnap()
	log.Debug("persistent end ==================================== ")
}

// Rank 排序
func (m *Module) Rank() {
	m.RankPlayer()
}

// RankPlayer 固化用户信息
func (m *Module) RankPlayer() {
	var userBaseInfos []*UserBaseInfo
	m.db.Order("star desc").Order("uid").Find(&userBaseInfos)

	m.rankPlayers = m.rankPlayers[0:0]

	var index = 0
	for _, baseInfo := range userBaseInfos {
		p := new(entry.Player)
		p.UserId = baseInfo.Uid
		p.Star = baseInfo.Star
		p.Name = baseInfo.Name
		p.HeadUrl = baseInfo.HeadUrl
		if index < 100 {
			m.rankPlayers = append(m.rankPlayers, p)
		}

		index++
		if p, ok := m.players[baseInfo.Uid]; ok {
			p.Order = int32(index)
		}
	}
}

// PersistentUser 固化用户信息
func (m *Module) PersistentUser() {
	cnt := 0
	for _, player := range m.players {
		user := User{
			Uid:        player.UserId,
			Account:    player.Account,
			Password:   player.Password,
			OpenId:     player.OpenId,
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
			HeadUrl: player.HeadUrl,
			Star:    player.Star,
			Exp:     player.Exp,
			LvChao:  player.LvChao,
			Diamond: player.Diamond,
		}

		// log.Debug("PersistentUser   ==== > %v ", userBaseInfo)

		var oldUserInfo UserBaseInfo
		m.db.Where("uid = ?", userBaseInfo.Uid).First(&oldUserInfo)
		if userBaseInfo.Uid != oldUserInfo.Uid {
			m.db.Create(&userBaseInfo)
		} else {
			m.db.Model(&userBaseInfo).Where("uid = ?", userBaseInfo.Uid).Updates(userBaseInfo)
		}

		userExtendInfo := UserExtendInfo{
			Uid:      player.UserId,
			Level:    player.Level,
			Scene:    player.Scene,
			Hair:     player.Hair,
			Coat:     player.Coat,
			Trouser:  player.Trouser,
			Neck:     player.Neck,
			Shoe:     player.Shoe,
			MaxCloth: player.MaxCloth,
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
	// for k := range m.signs {
	// 	delete(m.signs, k)
	// }
}

// PersistentCloth 固化衣服快照
func (m *Module) PersistentCloth() {
	cnt := 0
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

		cnt++
	}

	log.Debug("persistent cloth %v ", cnt)
}

// PersistentSnap 固化快照
func (m *Module) PersistentSnap() {
	for uid, v := range m.snaps {
		snap := UserSnapInfo{
			Uid:    uid,
			LvChao: v.LvChao,
		}

		var oldSnap UserSnapInfo
		m.db.Where("uid = ?", snap.Uid).First(&oldSnap)
		if snap.Uid != oldSnap.Uid {
			m.db.Create(&snap)
		} else {
			m.db.Model(&snap).Where("uid = ?", snap.Uid).Updates(snap)
		}
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
	if !m.db.HasTable(&UserSnapInfo{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserSnapInfo{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&ConfigCloth{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ConfigCloth{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&ConfigScene{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ConfigScene{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&ConfigLevel{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ConfigLevel{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&ConfigSign{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ConfigSign{}).Error; err != nil {
			panic(err)
		}
	}
}

// LoadFromDB 加载数据
func (m *Module) LoadFromDB() {
	log.Debug("LoadFromDB start ==================================== ")
	m.loadPlayer()
	m.loadCloth()
	m.loadSign()
	m.loadSnap()
	m.loadClothConfigs()
	m.loadSceneConfigs()
	m.loadLevelConfigs()
	m.loadSignConfigs()
	log.Debug("LoadFromDB end ==================================== ")
}

func (m *Module) loadPlayer() {
	tempPlayers := make(map[string]*entry.Player)
	var users []*User
	m.db.Find(&users)
	for _, user := range users {
		player := new(entry.Player)
		player.UserId = user.Uid
		player.Account = user.Account
		player.Password = user.Password
		player.OpenId = user.OpenId
		player.LoginTime = user.LoginTime
		player.LogoutTime = user.LogoutTime
		player.CreateTime = user.CreatedAt

		tempPlayers[player.UserId] = player
	}

	var userBaseInfos []*UserBaseInfo
	m.db.Find(&userBaseInfos)
	for _, baseInfo := range userBaseInfos {
		if tempPlayers[baseInfo.Uid] == nil {
			continue
		}
		// log.Debug("userbaseInfo ==== %v %v %v", baseInfo.Uid, baseInfo.Name, baseInfo.HeadUrl)
		tempPlayers[baseInfo.Uid].Name = baseInfo.Name
		tempPlayers[baseInfo.Uid].HeadUrl = baseInfo.HeadUrl
		tempPlayers[baseInfo.Uid].Star = baseInfo.Star
		tempPlayers[baseInfo.Uid].Exp = baseInfo.Exp
		tempPlayers[baseInfo.Uid].Diamond = baseInfo.Diamond
		tempPlayers[baseInfo.Uid].LvChao = baseInfo.LvChao
	}

	var userExtendInfos []*UserExtendInfo
	m.db.Find(&userExtendInfos)
	for _, extendInfo := range userExtendInfos {
		if tempPlayers[extendInfo.Uid] == nil {
			continue
		}
		// log.Debug("userbaseInfo ==== %v", baseInfo.Uid)
		tempPlayers[extendInfo.Uid].Level = extendInfo.Level
		tempPlayers[extendInfo.Uid].Scene = extendInfo.Scene
		tempPlayers[extendInfo.Uid].Coat = extendInfo.Coat
		tempPlayers[extendInfo.Uid].Trouser = extendInfo.Trouser
		tempPlayers[extendInfo.Uid].Hair = extendInfo.Hair
		tempPlayers[extendInfo.Uid].Neck = extendInfo.Neck
		tempPlayers[extendInfo.Uid].Shoe = extendInfo.Shoe
		tempPlayers[extendInfo.Uid].MaxCloth = extendInfo.MaxCloth
	}

	for _, player := range tempPlayers {
		m.SavePlayer(player)
	}

	// for _, player := range m.players {
	// 	log.Debug("userbaseInfo ==== %v %v %v", player.UserId, player.Name, player.HeadUrl)
	// }

	log.Debug("load players  db %v  mem %v", len(users), len(m.players))
}

// LoadCloth 加载衣服快照
func (m *Module) loadCloth() {
	var clothInfos []*UserClothInfo
	m.db.Find(&clothInfos)
	for _, cloth := range clothInfos {
		// log.Debug("LoadCloth %v %v", cloth.Uid, cloth.Snap)
		m.SaveCloth(cloth.Uid, cloth.Snap)
	}
	log.Debug("load cloths  db %v  mem %v", len(clothInfos), len(m.cloths))
}

// LoadSign 加载签到
func (m *Module) loadSign() {
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

func (m *Module) loadSnap() {
	var snapInfos []*UserSnapInfo
	m.db.Find(&snapInfos)
	for _, snap := range snapInfos {
		if s, ok := m.snaps[snap.Uid]; ok {
			s.LvChao = snap.LvChao
		} else {
			s := new(entry.Snap)
			s.LvChao = snap.LvChao
			m.snaps[snap.Uid] = s
		}
	}

	// for k, v := range m.snaps {
	// 	log.Debug("loadSnap  %v %v", k, v.LvChao)
	// }

	log.Debug("Load Snaps  db %v  mem %v", len(snapInfos), len(m.snaps))
}

func (m *Module) loadClothConfigs() {
	var clothConfigs []*ConfigCloth
	m.db.Find(&clothConfigs)

	m.clothConfigs = m.clothConfigs[0:0]
	for _, config := range clothConfigs {
		cloth := new(entry.ConfigCloth)
		cloth.Name = config.Name
		cloth.Icon = config.Icon
		cloth.Type = config.Type
		cloth.Level = config.Level
		cloth.Exp = config.Exp
		cloth.Cost = config.Cost
		m.clothConfigs = append(m.clothConfigs, cloth)
	}

	log.Debug("Load ClothConfigs  db %v  mem %v", len(clothConfigs), len(m.clothConfigs))
}

func (m *Module) loadSceneConfigs() {
	var sceneConfigs []*ConfigScene
	m.db.Find(&sceneConfigs)

	m.sceneConfigs = m.sceneConfigs[0:0]
	for _, config := range sceneConfigs {
		cloth := new(entry.ConfigScene)
		cloth.Id = config.No
		cloth.Name = config.Name
		cloth.Icon = config.Icon
		cloth.Level = config.Level
		cloth.Star = config.Star
		m.sceneConfigs = append(m.sceneConfigs, cloth)
	}

	log.Debug("Load SceneConfigs  db %v  mem %v", len(sceneConfigs), len(m.sceneConfigs))
}

func (m *Module) loadLevelConfigs() {
	var levelConfigs []*ConfigLevel
	m.db.Find(&levelConfigs)

	m.levelConfigs = m.levelConfigs[0:0]
	for _, config := range levelConfigs {
		cloth := new(entry.ConfigLevel)
		cloth.Id = config.No
		cloth.Name = config.Name
		cloth.Icon = config.Icon
		cloth.Level = config.Level
		cloth.Star = config.Star
		m.levelConfigs = append(m.levelConfigs, cloth)
	}

	log.Debug("Load LevelConfigs  db %v  mem %v", len(levelConfigs), len(m.levelConfigs))
}

func (m *Module) loadSignConfigs() {
	var signConfigs []*ConfigSign
	m.db.Find(&signConfigs)

	m.signConfigs = m.signConfigs[0:0]
	for _, config := range signConfigs {
		cloth := new(entry.ConfigSign)
		cloth.Id = config.No
		cloth.Day = config.Day
		cloth.Num = config.Num
		m.signConfigs = append(m.signConfigs, cloth)
	}

	log.Debug("Load SignConfigs  db %v  mem %v", len(signConfigs), len(m.signConfigs))
}
