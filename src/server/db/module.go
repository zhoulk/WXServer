package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"server/tool"
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
	snaps   map[string]*entry.Snap

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
	GetInstance().snaps = make(map[string]*entry.Snap)
}

// SavePlayer 保存用户信息
func (m *Module) SavePlayer(s *entry.Player) error {
	if s == nil || len(s.UserId) == 0 {
		return errors.New("player is nil or userId length is 0")
	}
	if player, ok := m.players[s.UserId]; ok {
		if len(s.Name) > 0 {
			player.Name = s.Name
		}
		if s.Star > 0 {
			player.Star = s.Star
		} else {
			player.Star = 1
		}
		if len(s.LvChao) > 0 {
			player.LvChao = s.LvChao
		}
		if s.Diamond > 0 {
			player.Diamond = s.Diamond
		}
		if s.Level > 0 {
			player.Level = s.Level
		}
		if s.Scene > 0 {
			player.Scene = s.Scene
		}
		if s.Hair > 0 {
			player.Hair = s.Hair
		}
		if s.Coat > 0 {
			player.Coat = s.Coat
		}
		if s.Trouser > 0 {
			player.Trouser = s.Trouser
		}
		if s.Neck > 0 {
			player.Neck = s.Neck
		}
	} else {
		m.players[player.UserId] = player
	}
	return nil
}

// SaveSnap 保存快照
func (m *Module) SaveSnap(player *entry.Player) error {
	if _, ok := m.snaps[player.UserId]; ok {
		m.snaps[player.UserId].LvChao = player.LvChao
	} else {
		snap := new(entry.Snap)
		snap.LvChao = player.LvChao
		m.snaps[player.UserId] = snap
	}

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

// GetOffLineLvChao 获取离线绿钞
func (m *Module) GetOffLineLvChao(uid string) string {
	// for k, v := range m.snaps {
	// 	log.Debug("GetOffLineLvChao %v %v %v", uid, k, v.LvChao)
	// }

	var lastLvChao = "[]"
	if snap, ok := m.snaps[uid]; ok {
		lastLvChao = snap.LvChao
	}
	lvChao := m.players[uid].LvChao

	// log.Debug("GetOffLineLvChao %v %v ", uid, lastLvChao)

	var s []int32
	json.Unmarshal([]byte(lastLvChao), &s)
	lastNum := new(tool.BigNumber)
	lastNum.FromArr(s)

	var s2 []int32
	json.Unmarshal([]byte(lvChao), &s2)
	curNum := new(tool.BigNumber)
	curNum.FromArr(s2)

	// log.Debug("GetOffLineLvChao ===> start\n %v\n %v ", curNum, lastNum)

	curNum.Minus(lastNum)

	// log.Debug("GetOffLineLvChao ===> end\n %v\n %v ", curNum, lastNum)

	bs, _ := json.Marshal(curNum.ToArr())
	return bytes.NewBuffer(bs).String()
}

// Cal 计算
func (m *Module) Cal() {
	for _, p := range m.players {
		p.Cal()
	}
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

	if sign, ok := m.signs[uid]; ok {
		sign[day] = time.Now()
	} else {
		m.signs[uid] = make(map[string]time.Time)
		m.signs[uid][day] = time.Now()
	}

}

// GetSign 获取签到信息
func (m *Module) GetSign(uid string) map[string]time.Time {
	return m.signs[uid]
}
