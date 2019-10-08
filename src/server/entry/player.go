package entry

import (
	"bytes"
	"encoding/json"
	"server/tool"
	"time"
)

type Player struct {
	UserId     string
	Account    string
	Password   string
	OpenId     string
	HeadUrl    string
	Name       string
	LoginTime  time.Time
	LogoutTime time.Time
	CreateTime time.Time

	Star       int32
	Exp        int32
	LvChao     string
	Diamond    int32
	Level      int32
	Scene      int32
	CP         int32
	Hair       int32
	Coat       int32
	Trouser    int32
	Neck       int32
	Shoe       int32
	Pet        int32
	MaxCloth   int32
	MaxCoat    int32
	MaxTrouser int32
	MaxShoe    int32

	Order int32

	HasFavour bool
	Favour    int32
	Constri   int32
}

// Cal 计算
func (p *Player) Cal() {
	// log.Debug("cal ===> start %v %v ", p.Star, p.LvChao)
	var s []int32
	json.Unmarshal([]byte(p.LvChao), &s)
	// log.Debug("%v", s)

	otherNum := new(tool.BigNumber)
	otherNum.Raw(p.Star * 10)

	// log.Debug("%v", otherNum)

	bNum := new(tool.BigNumber)
	bNum.FromArr(s)
	bNum.Add(otherNum)

	bs, _ := json.Marshal(bNum.ToArr())
	p.LvChao = bytes.NewBuffer(bs).String()

	// log.Debug("p.LvChao ====== %v", p.LvChao)
}

// BuyLvChao 购买绿钞
func (p *Player) BuyLvChao(num int32) {
	if p.Diamond >= num {
		p.Diamond -= num

		var s []int32
		json.Unmarshal([]byte(p.LvChao), &s)

		otherNum := new(tool.BigNumber)
		otherNum.Raw(num * 10000)

		// log.Debug("%v", otherNum)

		bNum := new(tool.BigNumber)
		bNum.FromArr(s)
		bNum.Add(otherNum)

		bs, _ := json.Marshal(bNum.ToArr())
		p.LvChao = bytes.NewBuffer(bs).String()
	}
}

// ExpendCloth  扩展衣橱
func (p *Player) ExpendCloth(num int32) {
	if p.Diamond >= num {
		p.Diamond -= num

		p.MaxCloth++
	}
}

// SellCloth  出售衣服
func (p *Player) SellCloth(t int32, level int32, cost string) {
	var s []int32
	json.Unmarshal([]byte(p.LvChao), &s)

	otherNum := new(tool.BigNumber)
	var otherArr []int32
	json.Unmarshal([]byte(cost), &otherArr)
	otherNum.FromArr(otherArr)

	// log.Debug("%v", otherNum)

	bNum := new(tool.BigNumber)
	bNum.FromArr(s)
	bNum.Add(otherNum)

	bs, _ := json.Marshal(bNum.ToArr())
	p.LvChao = bytes.NewBuffer(bs).String()
}
