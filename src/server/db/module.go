package db

import (
	"errors"
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	"server/entry"
)

// Module ...
type Module struct {
	players map[string]*entry.Player
	signs   map[string]map[string]time.Time
	cloths  map[string]string

	db *gorm.DB
}

var _instance *Module
var once sync.Once

// GetInstance 单例
func GetInstance() *Module {
	once.Do(func() {
		_instance = &Module{}
	})
	return _instance
}

func init() {
	GetInstance().players = make(map[string]*entry.Player)
	GetInstance().signs = make(map[string]map[string]time.Time)
	GetInstance().cloths = make(map[string]string)
}

// SavePlayer 保存用户信息
func (m *Module) SavePlayer(player *entry.Player) error {
	if player == nil || len(player.UserId) == 0 {
		return errors.New("player is nil or userId length is 0")
	}
	m.players[player.UserId] = player
	return nil
}

// SaveCloth 保存合成快照
func (m *Module) SaveCloth(uid string, snap string) {
	m.cloths[uid] = snap
}

// GetPlayer 获取用户信息
func (m *Module) GetPlayer(uid string) *entry.Player {
	return m.players[uid]
}

// GetCloth 获取合成快照
func (m *Module) GetCloth(uid string) string {
	// for k, v := range m.cloths {
	// 	log.Debug("GetCloth %v %v %v", uid, k, v)
	// }
	return m.cloths[uid]
}

// Sign  签到
func (m *Module) Sign(uid string) {
	now := time.Now()
	day := now.Format("2006/1/2")
	m.signs[uid][day] = time.Now()
}

// GetSign 获取签到信息
func (m *Module) GetSign(uid string) map[string]time.Time {
	return m.signs[uid]
}
