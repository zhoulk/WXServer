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
		"{\"no\" : 1, \"name\" : \"鞋子1\", \"icon\" : \"shoe/xiezi_01\", \"type\" : 1, \"level\" : 1, \"star\":10, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 2, \"name\" : \"鞋子2\", \"icon\" : \"shoe/xiezi_02\", \"type\" : 1, \"level\" : 2, \"star\":20, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 3, \"name\" : \"鞋子3\", \"icon\" : \"shoe/xiezi_03\", \"type\" : 1, \"level\" : 3, \"star\":30, \"exp\" : 10, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 4, \"name\" : \"鞋子4\", \"icon\" : \"shoe/xiezi_04\", \"type\" : 1, \"level\" : 4, \"star\":40, \"exp\" : 15, \"cost\" : \"[0,0,0,1]\"}",
		"{\"no\" : 5, \"name\" : \"鞋子5\", \"icon\" : \"shoe/xiezi_05\", \"type\" : 1, \"level\" : 5, \"star\":50, \"exp\" : 20, \"cost\" : \"[0,0,0,0,1]\"}",

		"{\"no\" : 101, \"name\" : \"上衣1\", \"icon\" : \"coat/yifu_01\", \"type\" : 2, \"level\" : 1, \"star\":5, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 102, \"name\" : \"上衣2\", \"icon\" : \"coat/yifu_02\", \"type\" : 2, \"level\" : 2, \"star\":10, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 103, \"name\" : \"上衣3\", \"icon\" : \"coat/yifu_03\", \"type\" : 2, \"level\" : 3, \"star\":15, \"exp\" : 10, \"cost\" : \"[0,5]\"}",
		"{\"no\" : 104, \"name\" : \"上衣4\", \"icon\" : \"coat/yifu_04\", \"type\" : 2, \"level\" : 4, \"star\":20, \"exp\" : 15, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 105, \"name\" : \"上衣5\", \"icon\" : \"coat/yifu_05\", \"type\" : 2, \"level\" : 5, \"star\":25, \"exp\" : 20, \"cost\" : \"[0,0,5]\"}",

		"{\"no\" : 201, \"name\" : \"下衣1\", \"icon\" : \"dress/kuzi_01\", \"type\" : 3, \"level\" : 1, \"star\":10, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 202, \"name\" : \"下衣2\", \"icon\" : \"dress/kuzi_02\", \"type\" : 3, \"level\" : 2, \"star\":20, \"exp\" : 10, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 203, \"name\" : \"下衣3\", \"icon\" : \"dress/kuzi_03\", \"type\" : 3, \"level\" : 3, \"star\":30, \"exp\" : 20, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 204, \"name\" : \"下衣4\", \"icon\" : \"dress/kuzi_04\", \"type\" : 3, \"level\" : 4, \"star\":40, \"exp\" : 30, \"cost\" : \"[0,0,0,1]\"}",
		"{\"no\" : 205, \"name\" : \"下衣5\", \"icon\" : \"dress/kuzi_05\", \"type\" : 3, \"level\" : 5, \"star\":50, \"exp\" : 40, \"cost\" : \"[0,0,0,0,1]\"}",

		"{\"no\" : 301, \"name\" : \"头发1\", \"icon\" : \"faxing/faxing_01\", \"type\" : 4, \"level\" : 1, \"star\":10, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 302, \"name\" : \"头发2\", \"icon\" : \"faxing/faxing_02\", \"type\" : 4, \"level\" : 2, \"star\":20, \"exp\" : 5, \"cost\" : \"[5]\"}",
		"{\"no\" : 303, \"name\" : \"头发3\", \"icon\" : \"faxing/faxing_03\", \"type\" : 4, \"level\" : 3, \"star\":30, \"exp\" : 10, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 304, \"name\" : \"头发4\", \"icon\" : \"faxing/faxing_04\", \"type\" : 4, \"level\" : 4, \"star\":40, \"exp\" : 15, \"cost\" : \"[5,1]\"}",
		"{\"no\" : 305, \"name\" : \"头发5\", \"icon\" : \"faxing/faxing_05\", \"type\" : 4, \"level\" : 5, \"star\":50, \"exp\" : 20, \"cost\" : \"[0,2]\"}",
		"{\"no\" : 306, \"name\" : \"头发6\", \"icon\" : \"faxing/faxing_06\", \"type\" : 4, \"level\" : 6, \"star\":60, \"exp\" : 25, \"cost\" : \"[5,2]\"}",
		"{\"no\" : 307, \"name\" : \"头发7\", \"icon\" : \"faxing/faxing_07\", \"type\" : 4, \"level\" : 7, \"star\":70, \"exp\" : 30, \"cost\" : \"[0,3]\"}",
		"{\"no\" : 308, \"name\" : \"头发8\", \"icon\" : \"faxing/faxing_08\", \"type\" : 4, \"level\" : 8, \"star\":80, \"exp\" : 35, \"cost\" : \"[5,3]\"}",
		"{\"no\" : 309, \"name\" : \"头发9\", \"icon\" : \"faxing/faxing_09\", \"type\" : 4, \"level\" : 9, \"star\":90, \"exp\" : 40, \"cost\" : \"[0,4]\"}",
		"{\"no\" : 310, \"name\" : \"头发10\", \"icon\" : \"faxing/faxing_10\", \"type\" : 4, \"level\" : 10, \"star\":100, \"exp\" : 45, \"cost\" : \"[5,4]\"}",
		"{\"no\" : 311, \"name\" : \"头发11\", \"icon\" : \"faxing/faxing_11\", \"type\" : 4, \"level\" : 11, \"star\":110, \"exp\" : 50, \"cost\" : \"[0,5]\"}",
		"{\"no\" : 312, \"name\" : \"头发12\", \"icon\" : \"faxing/faxing_12\", \"type\" : 4, \"level\" : 12, \"star\":120, \"exp\" : 55, \"cost\" : \"[5,5]\"}",
		"{\"no\" : 313, \"name\" : \"头发13\", \"icon\" : \"faxing/faxing_13\", \"type\" : 4, \"level\" : 13, \"star\":130, \"exp\" : 60, \"cost\" : \"[0,6]\"}",
		"{\"no\" : 314, \"name\" : \"头发14\", \"icon\" : \"faxing/faxing_14\", \"type\" : 4, \"level\" : 14, \"star\":140, \"exp\" : 65, \"cost\" : \"[5,6]\"}",
		"{\"no\" : 315, \"name\" : \"头发15\", \"icon\" : \"faxing/faxing_15\", \"type\" : 4, \"level\" : 15, \"star\":160, \"exp\" : 70, \"cost\" : \"[0,7]\"}",
		"{\"no\" : 316, \"name\" : \"头发16\", \"icon\" : \"faxing/faxing_16\", \"type\" : 4, \"level\" : 16, \"star\":160, \"exp\" : 75, \"cost\" : \"[5,7]\"}",
		"{\"no\" : 317, \"name\" : \"头发17\", \"icon\" : \"faxing/faxing_17\", \"type\" : 4, \"level\" : 17, \"star\":170, \"exp\" : 80, \"cost\" : \"[0,8]\"}",
		"{\"no\" : 318, \"name\" : \"头发18\", \"icon\" : \"faxing/faxing_18\", \"type\" : 4, \"level\" : 18, \"star\":180, \"exp\" : 85, \"cost\" : \"[5,8]\"}",
		"{\"no\" : 319, \"name\" : \"头发19\", \"icon\" : \"faxing/faxing_19\", \"type\" : 4, \"level\" : 19, \"star\":190, \"exp\" : 90, \"cost\" : \"[0,9]\"}",
		"{\"no\" : 320, \"name\" : \"头发20\", \"icon\" : \"faxing/faxing_20\", \"type\" : 4, \"level\" : 20, \"star\":200, \"exp\" : 95, \"cost\" : \"[5,9]\"}",

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
