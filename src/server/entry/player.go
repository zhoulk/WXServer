package entry

import (
	"bytes"
	"encoding/json"
	"server/tool"
	"time"

	"github.com/name5566/leaf/log"
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
	//log.Debug("cal ===> start %v %v", p.UserId, p.LogoutTime)

	now := time.Now()             //取到当前的时间点
	subM := now.Sub(p.LogoutTime) //通过这个方法我们可以将两个事件差值计算出来
	// fmt.Println(int(subM.Hours()), "Hours") //我们打印一下相距的小时数
	//log.Debug("cal ===> start %v", subM.Hours())
	// 离线最多8小时的收益
	if subM.Hours() > 8 {
		return
	}

	var s []int64
	json.Unmarshal([]byte(p.LvChao), &s)
	// log.Debug("%v", s)

	otherNum := new(tool.BigNumber)
	//otherNum.Raw(p.Star * 30)
	otherNum.Raw(int64(p.Star))

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

		var s []int64
		json.Unmarshal([]byte(p.LvChao), &s)

		otherNum := new(tool.BigNumber)
		otherNum.Raw(int64(num * 10000))

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
	log.Debug("%v", cost)

	var s []int64
	json.Unmarshal([]byte(p.LvChao), &s)

	otherNum := new(tool.BigNumber)
	var otherArr []int64
	json.Unmarshal([]byte(cost), &otherArr)
	otherNum.FromArr(otherArr)

	// log.Debug("%v", otherNum)

	bNum := new(tool.BigNumber)
	bNum.FromArr(s)
	bNum.Add(otherNum)

	bs, _ := json.Marshal(bNum.ToArr())
	p.LvChao = bytes.NewBuffer(bs).String()
}
