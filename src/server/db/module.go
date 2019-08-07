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
	players          map[string]*entry.Player
	signs            map[string]map[string]time.Time
	cloths           map[string]string
	snaps            map[string]*entry.Snap
	rankPlayers      []*entry.Player
	clothConfigs     []*entry.ConfigCloth
	sceneConfigs     []*entry.ConfigScene
	levelConfigs     []*entry.ConfigLevel
	signConfigs      []*entry.ConfigSign
	favourLogs       []*entry.FavourLog
	favourFlag       map[string]bool
	favourReportLogs map[string]*entry.FavourReport

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
	GetInstance().rankPlayers = make([]*entry.Player, 0)
	GetInstance().clothConfigs = make([]*entry.ConfigCloth, 0)
	GetInstance().sceneConfigs = make([]*entry.ConfigScene, 0)
	GetInstance().levelConfigs = make([]*entry.ConfigLevel, 0)
	GetInstance().signConfigs = make([]*entry.ConfigSign, 0)
	GetInstance().favourLogs = make([]*entry.FavourLog, 0)
	GetInstance().favourFlag = make(map[string]bool)
	GetInstance().favourReportLogs = make(map[string]*entry.FavourReport)

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
		if len(s.HeadUrl) > 0 {
			player.HeadUrl = s.HeadUrl
		}
		star := s.Star
		if star == 0 {
			star = 1
		}
		if player.Star < star {
			player.Star = star
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
		if s.Shoe > 0 {
			player.Shoe = s.Shoe
		}
	} else {
		m.players[s.UserId] = s
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
	// for k, v := range m.cloths {
	// 	log.Debug("SaveCloth  %v  %v  %v", uid, k, v)
	// }
}

// GetPlayer 获取用户信息
func (m *Module) GetPlayer(uid string) *entry.Player {
	return m.players[uid]
}

// FindPlayerByOpenID 根据openId查找
func (m *Module) FindPlayerByOpenID(openID string) *entry.Player {
	if len(openID) == 0 {
		return nil
	}
	var res *entry.Player
	for _, v := range m.players {
		if v.OpenId == openID {
			res = v
		}
	}
	return res
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

// GetRank 获取排行榜
func (m *Module) GetRank(uid string) []*entry.Player {

	day := time.Now().Format("2006/1/2")
	for _, p := range m.rankPlayers {
		if flag, ok := m.favourFlag[uid+p.UserId+day]; ok {
			if flag {
				p.HasFavour = true
			}
		}
		if report, ok := m.favourReportLogs["all"+p.UserId]; ok {
			p.Favour = report.Num
		}
	}

	return m.rankPlayers
}

// Heart 心跳
func (m *Module) Heart(uid string) {
	if player, ok := m.players[uid]; ok {
		player.LogoutTime = time.Now()

		m.SaveSnap(player)
	}
}

// Buy 购买
func (m *Module) Buy(uid string, t int32, num int32) {
	if player, ok := m.players[uid]; ok {
		if t == 1 {
			player.BuyLvChao(num)
		} else if t == 2 {
			player.ExpendCloth(num)
		}
	}
}

// SellClothParams  ...
type SellClothParams struct {
	Type  int32
	Level int32
}

// Sell 购买
func (m *Module) Sell(uid string, t int32, params string) {
	if player, ok := m.players[uid]; ok {
		if t == 1 {
			var s SellClothParams
			json.Unmarshal([]byte(params), &s)
			cost := m.CostOfCloth(s.Type, s.Level)
			player.SellCloth(s.Type, s.Level, cost)
		}
	}
}

// Favour 点赞
func (m *Module) Favour(uid string, toUID string, num int32) {
	l := new(entry.FavourLog)
	l.Uid = uid
	l.ToUID = toUID
	l.Day = time.Now().Format("2006/1/2")
	l.Num = num
	m.favourLogs = append(m.favourLogs, l)

	m.favourFlag[l.Uid+l.ToUID+l.Day] = true

	if report, ok := m.favourReportLogs[uid+toUID]; ok {
		report.Num = report.Num + 1
		m.favourReportLogs[uid+toUID] = report
	} else {
		report := new(entry.FavourReport)
		report.From = uid
		report.To = toUID
		report.Num = 1
		m.favourReportLogs[uid+toUID] = report
	}

	if report, ok := m.favourReportLogs["all"+toUID]; ok {
		report.Num = report.Num + 1
		m.favourReportLogs["all"+toUID] = report
	} else {
		report := new(entry.FavourReport)
		report.From = "all"
		report.To = toUID
		report.Num = 1
		m.favourReportLogs["all"+toUID] = report
	}
}

// func (m *Module) Rank() {
// 	// log.Debug("rank  ====>  %v %v %v", p.Name, p.Star, p.UserId)
// 	var insertIndex = -1
// 	var existIndex = -1
// 	if len(m.rankPlayers) == 0 {
// 		insertIndex = 0
// 	} else {
// 		for i := len(m.rankPlayers) - 1; i >= 0; i-- {
// 			if p.Star >= m.rankPlayers[i].Star {
// 				insertIndex = i
// 			}
// 			if m.rankPlayers[i].UserId == p.UserId {
// 				existIndex = i
// 			}
// 		}
// 	}

// 	// log.Debug("rank  ====>  %v %v", insertIndex, existIndex)

// 	if insertIndex >= 0 {
// 		if existIndex >= 0 {
// 			m.rankPlayers = append(m.rankPlayers[:existIndex], m.rankPlayers[existIndex+1:]...)
// 		}
// 		last := append([]*entry.Player{}, m.rankPlayers[insertIndex:]...)
// 		m.rankPlayers = append(append(m.rankPlayers[:insertIndex], p), last...)
// 	} else {
// 		if existIndex == -1 && len(m.rankPlayers) < 100 {
// 			m.rankPlayers = append(m.rankPlayers, p)
// 		}
// 	}

// 	// for i := 0; i < len(m.rankPlayers); i++ {
// 	// 	log.Debug("rankPlayers  ====>  %v %v", m.rankPlayers[i].UserId, m.rankPlayers[i].Star)
// 	// }
// }
