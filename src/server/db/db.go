package db

import (
	"fmt"
	"server/entry"

	"github.com/name5566/leaf/log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_Driver = "root:A845240287a@tcp(rm-wz9sw694mi8020vigo.mysql.rds.aliyuncs.com:3306)/wxgame?charset=utf8&&parseTime=true"
)

func (m *Module) ConnectDB() {
	db, err := gorm.Open("mysql", DB_Driver)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	// defer db.Close()
	m.db = db
}

func (m *Module) PersistentData() {
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "dota_" + defaultTableName
	// }

	// m.ConnectDB()
	// m.CreateTables()
	// m.IntializeTables()

	m.PersistentUser()
}

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

		cnt++
	}
	log.Debug("persistent user %v ", cnt)
}

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
}

func (m *Module) IntializeTables() {

}

func (m *Module) LoadFromDB() {
	log.Debug("loading data from db start ...")
	m.LoadPlayer()
	log.Debug("loading data from db end ...")
}

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
