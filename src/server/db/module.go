package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"server/tool"
	"sort"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/name5566/leaf/log"

	"server/entry"
)

// Module ...
type Module struct {
	players           map[string]*entry.Player
	signs             map[string]map[string]time.Time
	cloths            map[string]string
	snaps             map[string]*entry.Snap
	rankPlayers       []*entry.Player
	clothConfigs      []*entry.ConfigCloth
	sceneConfigs      []*entry.ConfigScene
	levelConfigs      []*entry.ConfigLevel
	signConfigs       []*entry.ConfigSign
	giftConfigs       []*entry.ConfigGift
	CPConfigs         []*entry.ConfigCP
	favourLogs        []*entry.FavourLog
	favourFlag        map[string]bool
	favourReportLogs  map[string]*entry.FavourReport
	barrageReports    []*entry.BarrageReport
	addBarrageReports []*entry.BarrageReport
	extraMoney        []*entry.ExtraMoney
	openFroms         []*entry.OpenFrom
	gainDailyGifts    map[string]bool

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
	GetInstance().giftConfigs = make([]*entry.ConfigGift, 0)
	GetInstance().CPConfigs = make([]*entry.ConfigCP, 0)
	GetInstance().favourLogs = make([]*entry.FavourLog, 0)
	GetInstance().barrageReports = make([]*entry.BarrageReport, 0)
	GetInstance().addBarrageReports = make([]*entry.BarrageReport, 0)
	GetInstance().favourFlag = make(map[string]bool)
	GetInstance().gainDailyGifts = make(map[string]bool)
	GetInstance().favourReportLogs = make(map[string]*entry.FavourReport)
	GetInstance().extraMoney = make([]*entry.ExtraMoney, 0)
	GetInstance().openFroms = make([]*entry.OpenFrom, 0)

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
		if player.Star <= star {
			player.Star = star
			player.Exp = s.Exp
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
		if s.CP > 0 {
			player.CP = s.CP
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
		if s.Pet > 0 {
			player.Pet = s.Pet
		}
		if s.MaxCoat > 0 {
			player.MaxCoat = s.MaxCoat
		}
		if s.MaxShoe > 0 {
			player.MaxShoe = s.MaxShoe
		}
		if s.MaxTrouser > 0 {
			player.MaxTrouser = s.MaxTrouser
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
	player := m.players[uid]
	if report, ok := m.favourReportLogs["all"+uid]; ok {
		player.Favour = report.Num
	}
	return player
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

type fanPlayers []*entry.Player

func (s fanPlayers) Len() int           { return len(s) }
func (s fanPlayers) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s fanPlayers) Less(i, j int) bool { return s[i].Favour < s[j].Favour }

// GetFansRank 获取点赞排行
func (m *Module) GetFansRank(uid string) []*entry.Player {
	players := make([]*entry.Player, 0)
	for _, report := range m.favourReportLogs {
		if report.To == uid {
			p := m.players[report.From]
			if p != nil {
				p.Constri = report.Num
				players = append(players, p)
			}
		}
	}

	sort.Sort(fanPlayers(players))
	return players
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

//Reward 送礼物
func (m *Module) Reward(uid string, toUID string, msg string, giftID int32) bool {

	var curGift *entry.ConfigGift
	for _, gift := range m.giftConfigs {
		if gift.Id == giftID {
			curGift = gift
		}
	}

	if m.players[uid].Diamond >= curGift.Diamond {
		l := new(entry.FavourLog)
		l.Uid = uid
		l.ToUID = toUID
		l.Day = time.Now().Format("2006/1/2")
		l.Num = curGift.Favour
		m.favourLogs = append(m.favourLogs, l)

		if report, ok := m.favourReportLogs[uid+toUID]; ok {
			report.Num = report.Num + curGift.Favour
			m.favourReportLogs[uid+toUID] = report
		} else {
			report := new(entry.FavourReport)
			report.From = uid
			report.To = toUID
			report.Num = curGift.Favour
			m.favourReportLogs[uid+toUID] = report
		}

		// 对方加
		if report, ok := m.favourReportLogs["all"+toUID]; ok {
			report.Num = report.Num + curGift.Favour
			m.favourReportLogs["all"+toUID] = report
		} else {
			report := new(entry.FavourReport)
			report.From = "all"
			report.To = toUID
			report.Num = curGift.Favour
			m.favourReportLogs["all"+toUID] = report
		}
		// 自己加
		if report, ok := m.favourReportLogs["all"+uid]; ok {
			report.Num = report.Num + curGift.Favour
			m.favourReportLogs["all"+uid] = report
		} else {
			report := new(entry.FavourReport)
			report.From = "all"
			report.To = uid
			report.Num = curGift.Favour
			m.favourReportLogs["all"+uid] = report
		}

		// 消耗钻石
		m.players[uid].Diamond -= curGift.Diamond
		// 添加钻石
		m.players[toUID].Diamond += curGift.Reward

		if len(msg) > 0 {
			br := new(entry.BarrageReport)
			br.From = uid
			br.To = toUID
			br.Msg = msg
			m.addBarrageReports = append(m.addBarrageReports, br)
			m.barrageReports = append(m.barrageReports, br)
		}

		return true
	}

	return false
}

// GetBarrage  获取弹幕
func (m *Module) GetBarrage(uid string) []*entry.BarrageReport {
	barrages := make([]*entry.BarrageReport, 0)
	for _, report := range m.barrageReports {
		// log.Debug("GetBarrage  ====>  %v %v %v", uid, report.To, report.Msg)
		if report.To == uid || report.To == "all" {
			barrages = append(barrages, report)
		}
	}

	return barrages
}

// ExtraMoney  额外绿钞
func (m *Module) ExtraMoney(uid string, lvChao string, diamond int32) {
	otherNum := new(tool.BigNumber)
	var otherArr []int32
	json.Unmarshal([]byte(lvChao), &otherArr)
	otherNum.FromArr(otherArr)

	if player, ok := m.players[uid]; ok {
		bNum := new(tool.BigNumber)
		var bArr []int32
		json.Unmarshal([]byte(player.LvChao), &bArr)
		bNum.FromArr(bArr)
		bNum.Add(otherNum)

		bs, _ := json.Marshal(bNum.ToArr())
		player.LvChao = bytes.NewBuffer(bs).String()
	}

	extra := new(entry.ExtraMoney)
	extra.Uid = uid
	extra.LvChao = lvChao
	extra.Diamond = diamond
	extra.Reason = 1
	m.extraMoney = append(m.extraMoney, extra)
}

// OpenFrom 来自谁的分享
func (m *Module) OpenFrom(uid string, fromUid string, t int32) {
	log.Debug("OpenFrom  ====>  %v ", len(m.openFroms))

	exist := false
	for _, p := range m.openFroms {
		if p.Uid == uid && p.FromUid == fromUid {
			exist = true
			break
		}
	}

	if !exist {
		of := new(entry.OpenFrom)
		of.Uid = uid
		of.FromUid = fromUid
		of.Type = t
		m.openFroms = append(m.openFroms, of)
	}

	log.Debug("OpenFrom  ====>  %v ", len(m.openFroms))
}

// GetPrePlayer 获取前面一名玩家
func (m *Module) GetPrePlayer(uid string) *entry.Player {
	prePlayer := m.GetPlayer(uid)
	for _, p := range m.rankPlayers {
		if p.UserId == uid {
			break
		}
		prePlayer = p
	}
	return prePlayer
}

// GetInvitePlayers 获取邀请了哪些玩家玩家
func (m *Module) GetInvitePlayers(uid string) []*entry.Player {
	players := make([]*entry.Player, 0)
	for _, p := range m.openFroms {
		log.Debug("GetInvitePlayers  ====>  %v %v", p.Uid, p.FromUid)
		if p.FromUid == uid {
			player := m.GetPlayer(p.Uid)
			players = append(players, player)
		}
	}
	return players
}

func (m *Module) CheckGainDailyGift(uid string) bool {
	day := time.Now().Format("2006/1/2")
	if _, ok := m.gainDailyGifts[uid+day]; ok {
		return true
	}
	return false
}

// GetDailyGift ...
func (m *Module) GetDailyGift(uid string) string {

	day := time.Now().Format("2006/1/2")
	if _, ok := m.gainDailyGifts[uid+day]; ok {

	} else {
		m.gainDailyGifts[uid+day] = true
	}

	otherNum := new(tool.BigNumber)
	otherArr := []int32{1, 0, 0, 1, 0, 1}
	otherNum.FromArr(otherArr)

	player := m.GetPlayer(uid)

	bNum := new(tool.BigNumber)
	var bArr []int32
	json.Unmarshal([]byte(player.LvChao), &bArr)
	bNum.FromArr(bArr)
	bNum.Add(otherNum)
	bs, _ := json.Marshal(bNum.ToArr())
	player.LvChao = bytes.NewBuffer(bs).String()

	otherBs, _ := json.Marshal(otherNum.ToArr())
	lvChao := bytes.NewBuffer(otherBs).String()

	return lvChao
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
