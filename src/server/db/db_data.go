package db

import (
	"encoding/json"
)

// InitializeConfigs  初始化
func (m *Module) InitializeConfigs() {
	m.InitializeClothConfig()
	m.InitializeSceneConfig()
	m.InitializeLevelConfig()
	m.InitializeSignConfig()
}

// InitializeSignConfig 初始化咔位配置
func (m *Module) InitializeSignConfig() {
	m.db.Unscoped().Delete(&ConfigSign{})

	var configs = []string{
		"{\"no\" : 1, \"day\" : 1, \"num\" : 30}",
		"{\"no\" : 2, \"day\" : 2, \"num\" : 50}",
		"{\"no\" : 3, \"day\" : 3, \"num\" : 100}",
		"{\"no\" : 4, \"day\" : 4, \"num\" : 150}",
		"{\"no\" : 5, \"day\" : 5, \"num\" : 200}",
		"{\"no\" : 6, \"day\" : 6, \"num\" : 200}",
		"{\"no\" : 7, \"day\" : 7, \"num\" : 200}",
	}

	for _, configStr := range configs {
		var s ConfigSign
		json.Unmarshal([]byte(configStr), &s)
		m.db.Create(&s)
	}
}

// InitializeLevelConfig 初始化咔位配置
func (m *Module) InitializeLevelConfig() {
	m.db.Unscoped().Delete(&ConfigLevel{})

	var configs = []string{
		"{\"no\" : 1, \"name\" : \"咔位1\", \"icon\" : \"\", \"star\" : 10, \"level\" : 1}",
		"{\"no\" : 2, \"name\" : \"咔位2\", \"icon\" : \"\", \"star\" : 20, \"level\" : 2}",
		"{\"no\" : 3, \"name\" : \"咔位3\", \"icon\" : \"\", \"star\" : 30, \"level\" : 3}",
		"{\"no\" : 4, \"name\" : \"咔位4\", \"icon\" : \"\", \"star\" : 40, \"level\" : 4}",
		"{\"no\" : 5, \"name\" : \"咔位5\", \"icon\" : \"\", \"star\" : 50, \"level\" : 5}",
		"{\"no\" : 6, \"name\" : \"咔位6\", \"icon\" : \"\", \"star\" : 60, \"level\" : 6}",
		"{\"no\" : 7, \"name\" : \"咔位7\", \"icon\" : \"\", \"star\" : 70, \"level\" : 7}",
		"{\"no\" : 8, \"name\" : \"咔位8\", \"icon\" : \"\", \"star\" : 80, \"level\" : 8}",
		"{\"no\" : 9, \"name\" : \"咔位9\", \"icon\" : \"\", \"star\" : 90, \"level\" : 9}",
	}

	for _, configStr := range configs {
		var s ConfigLevel
		json.Unmarshal([]byte(configStr), &s)
		m.db.Create(&s)
	}
}

// InitializeSceneConfig 初始化场景配置
func (m *Module) InitializeSceneConfig() {
	m.db.Unscoped().Delete(&ConfigScene{})

	var configs = []string{
		"{\"no\" : 1, \"name\" : \"场景1\", \"icon\" : \"\", \"star\" : 10, \"level\" : 1}",
		"{\"no\" : 2, \"name\" : \"场景2\", \"icon\" : \"\", \"star\" : 20, \"level\" : 2}",
		"{\"no\" : 3, \"name\" : \"场景3\", \"icon\" : \"\", \"star\" : 30, \"level\" : 3}",
		"{\"no\" : 4, \"name\" : \"场景4\", \"icon\" : \"\", \"star\" : 40, \"level\" : 4}",
		"{\"no\" : 5, \"name\" : \"场景5\", \"icon\" : \"\", \"star\" : 50, \"level\" : 5}",
		"{\"no\" : 6, \"name\" : \"场景6\", \"icon\" : \"\", \"star\" : 60, \"level\" : 6}",
		"{\"no\" : 7, \"name\" : \"场景7\", \"icon\" : \"\", \"star\" : 70, \"level\" : 7}",
		"{\"no\" : 8, \"name\" : \"场景8\", \"icon\" : \"\", \"star\" : 80, \"level\" : 8}",
		"{\"no\" : 9, \"name\" : \"场景9\", \"icon\" : \"\", \"star\" : 90, \"level\" : 9}",
	}

	for _, configStr := range configs {
		var s ConfigScene
		json.Unmarshal([]byte(configStr), &s)
		m.db.Create(&s)
	}
}

// InitializeClothConfig 初始化衣服配置
func (m *Module) InitializeClothConfig() {
	m.db.Unscoped().Delete(&ConfigCloth{})

	var configs = []string{
		"{\"no\" : 1, \"name\" : \"鞋子1\", \"icon\" : \"player_shoes1\", \"type\" : 1, \"level\" : 1, \"star\":10, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 2, \"name\" : \"鞋子2\", \"icon\" : \"player_shoes2\", \"type\" : 1, \"level\" : 2, \"star\":20, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 3, \"name\" : \"鞋子3\", \"icon\" : \"player_shoes3\", \"type\" : 1, \"level\" : 3, \"star\":30, \"exp\" : 10, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 4, \"name\" : \"鞋子4\", \"icon\" : \"player_shoes4\", \"type\" : 1, \"level\" : 4, \"star\":40, \"exp\" : 15, \"cost\" : \"[0,0,0,1]\"}",

		"{\"no\" : 100, \"name\" : \"上衣1\", \"icon\" : \"item_coat8\", \"type\" : 2, \"level\" : 1, \"star\":20, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 102, \"name\" : \"上衣2\", \"icon\" : \"item_coat13\", \"type\" : 2, \"level\" : 2, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"no\" : 201, \"name\" : \"下衣1\", \"icon\" : \"player_dress3\", \"type\" : 3, \"level\" : 1, \"star\":20, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 202, \"name\" : \"下衣2\", \"icon\" : \"player_dress6\", \"type\" : 3, \"level\" : 2, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"no\" : 301, \"name\" : \"头发1\", \"icon\" : \"player_hair1\", \"type\" : 4, \"level\" : 1, \"star\":20, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 302, \"name\" : \"头发2\", \"icon\" : \"player_hair2\", \"type\" : 4, \"level\" : 2, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"no\" : 401, \"name\" : \"围巾1\", \"icon\" : \"player_scarf1\", \"type\" : 5, \"level\" : 1, \"star\":20, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 402, \"name\" : \"围巾2\", \"icon\" : \"player_scarf2\", \"type\" : 5, \"level\" : 2, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"no\" : 501, \"name\" : \"包包1\", \"icon\" : \"player_bag1\", \"type\" : 6, \"level\" : 1, \"star\":20, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 502, \"name\" : \"包包2\", \"icon\" : \"player_bag2\", \"type\" : 6, \"level\" : 2, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
	}

	for _, configStr := range configs {
		// log.Debug("Load ClothConfigs  db %v ", configStr)
		var s ConfigCloth
		json.Unmarshal([]byte(configStr), &s)
		m.db.Create(&s)
	}
}
