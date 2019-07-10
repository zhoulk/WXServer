package entry

import (
	"time"
)

type Player struct {
	UserId     string
	Account    string
	Password   string
	Name       string
	LoginTime  time.Time
	LogoutTime time.Time

	Star    int32
	LvChao  int32
	Diamond int32
	Level   int32
	Scene   int32
	Hair    int32
	Coat    int32
	Trouser int32
	Neck    int32
}
