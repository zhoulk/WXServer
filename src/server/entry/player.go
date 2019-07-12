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
	Name       string
	LoginTime  time.Time
	LogoutTime time.Time
	CreateTime time.Time

	Star    int32
	LvChao  string
	Diamond int32
	Level   int32
	Scene   int32
	Hair    int32
	Coat    int32
	Trouser int32
	Neck    int32
}

// Cal 计算
func (p *Player) Cal() {
	// log.Debug("cal ===> start %v %v ", p.Star, p.LvChao)
	var s []int32
	json.Unmarshal([]byte(p.LvChao), &s)
	// log.Debug("%v", s)

	otherNum := new(tool.BigNumber)
	otherNum.Raw(p.Star * 2)

	// log.Debug("%v", otherNum)

	bNum := new(tool.BigNumber)
	bNum.FromArr(s)
	bNum.Add(otherNum)

	bs, _ := json.Marshal(bNum.ToArr())
	p.LvChao = bytes.NewBuffer(bs).String()

	// log.Debug("p.LvChao ====== %v", p.LvChao)
}
