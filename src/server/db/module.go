package db

import (
	"errors"
	"sync"

	"github.com/jinzhu/gorm"

	"server/entry"
)

type Module struct {
	players map[string]*entry.Player

	db *gorm.DB
}

var _instance *Module
var once sync.Once

func GetInstance() *Module {
	once.Do(func() {
		_instance = &Module{}
	})
	return _instance
}

func (m *Module) init() {
	m.players = make(map[string]*entry.Player)
}

func (m *Module) SavePlayer(player *entry.Player) error {
	m.players = make(map[string]*entry.Player)
	if player == nil || len(player.UserId) == 0 {
		return errors.New("player is nil or userId length is 0")
	}
	m.players[player.UserId] = player
	return nil
}
