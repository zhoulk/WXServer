package db

import (
	"encoding/json"
)

// InitializeConfigs  初始化
func (m *Module) InitializeConfigs() {
	m.InitializeClothConfig()
}

// InitializeClothConfig 初始化衣服配置
func (m *Module) InitializeClothConfig() {
	m.db.Unscoped().Delete(&ConfigCloth{})

	var configs = []string{
		"{\"name\" : \"鞋子1\", \"icon\" : \"player_shoes1\", \"type\" : 1, \"level\" : 1, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"name\" : \"鞋子2\", \"icon\" : \"player_shoes2\", \"type\" : 1, \"level\" : 2, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"name\" : \"鞋子3\", \"icon\" : \"player_shoes3\", \"type\" : 1, \"level\" : 3, \"exp\" : 10, \"cost\" : \"[0,0,1]\"}",
		"{\"name\" : \"鞋子4\", \"icon\" : \"player_shoes4\", \"type\" : 1, \"level\" : 4, \"exp\" : 15, \"cost\" : \"[0,0,0,1]\"}",

		"{\"name\" : \"上衣1\", \"icon\" : \"item_coat8\", \"type\" : 2, \"level\" : 1, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"name\" : \"上衣2\", \"icon\" : \"item_coat13\", \"type\" : 2, \"level\" : 2, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"name\" : \"下衣1\", \"icon\" : \"player_dress3\", \"type\" : 3, \"level\" : 1, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"name\" : \"下衣2\", \"icon\" : \"player_dress6\", \"type\" : 3, \"level\" : 2, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"name\" : \"头发1\", \"icon\" : \"player_hair1\", \"type\" : 4, \"level\" : 1, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"name\" : \"头发2\", \"icon\" : \"player_hair2\", \"type\" : 4, \"level\" : 2, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"name\" : \"围巾1\", \"icon\" : \"player_scarf1\", \"type\" : 5, \"level\" : 1, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"name\" : \"围巾2\", \"icon\" : \"player_scarf2\", \"type\" : 5, \"level\" : 2, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"name\" : \"包包1\", \"icon\" : \"player_bag1\", \"type\" : 6, \"level\" : 1, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"name\" : \"包包2\", \"icon\" : \"player_bag2\", \"type\" : 6, \"level\" : 2, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
	}

	for _, configStr := range configs {
		var s ConfigCloth
		json.Unmarshal([]byte(configStr), &s)
		m.db.Create(&s)
	}
}
