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
	m.InitializeGiftConfig()
	m.InitializeBarrages()
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
		"{\"no\" : 1, \"name\" : \"巴厘岛\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/1.jpg\", \"star\" : 0, \"level\" : 1}",
		"{\"no\" : 2, \"name\" : \"丹麦古堡\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/2.jpg\", \"star\" : 20, \"level\" : 2}",
		"{\"no\" : 3, \"name\" : \"法国巴黎\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/3.jpg\", \"star\" : 30, \"level\" : 3}",
		"{\"no\" : 4, \"name\" : \"风情瑞士\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/4.jpg\", \"star\" : 40, \"level\" : 4}",
		"{\"no\" : 5, \"name\" : \"古巴\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/5.jpg\", \"star\" : 50, \"level\" : 5}",
		"{\"no\" : 6, \"name\" : \"赫尔辛基\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/6.jpg\", \"star\" : 60, \"level\" : 6}",
		"{\"no\" : 7, \"name\" : \"喀山古城\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/7.jpg\", \"star\" : 70, \"level\" : 7}",
		"{\"no\" : 8, \"name\" : \"浪漫悉尼\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/8.jpg\", \"star\" : 80, \"level\" : 8}",
		"{\"no\" : 9, \"name\" : \"马尔代夫\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/9.jpg\", \"star\" : 90, \"level\" : 9}",
		"{\"no\" : 10, \"name\" : \"圣托里尼\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/10.jpg\", \"star\" : 100, \"level\" : 10}",
		"{\"no\" : 11, \"name\" : \"泰姬陵\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/11.jpg\", \"star\" : 110, \"level\" : 11}",
		"{\"no\" : 12, \"name\" : \"威尼斯\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/12.jpg\", \"star\" : 120, \"level\" : 12}",
		"{\"no\" : 13, \"name\" : \"西雅图\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/13.jpg\", \"star\" : 130, \"level\" : 13}",
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
		"{\"no\" : 1, \"name\" : \"鞋子1\", \"icon\" : \"shoe/xiezi_1\", \"type\" : 1, \"level\" : 1, \"star\":0, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 2, \"name\" : \"鞋子2\", \"icon\" : \"shoe/xiezi_2\", \"type\" : 1, \"level\" : 2, \"star\":20, \"exp\" : 5, \"cost\" : \"[5]\"}",
		"{\"no\" : 3, \"name\" : \"鞋子3\", \"icon\" : \"shoe/xiezi_3\", \"type\" : 1, \"level\" : 3, \"star\":30, \"exp\" : 10, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 4, \"name\" : \"鞋子4\", \"icon\" : \"shoe/xiezi_4\", \"type\" : 1, \"level\" : 4, \"star\":40, \"exp\" : 15, \"cost\" : \"[0,5]\"}",
		"{\"no\" : 5, \"name\" : \"鞋子5\", \"icon\" : \"shoe/xiezi_5\", \"type\" : 1, \"level\" : 5, \"star\":50, \"exp\" : 20, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 6, \"name\" : \"鞋子6\", \"icon\" : \"shoe/xiezi_6\", \"type\" : 1, \"level\" : 6, \"star\":60, \"exp\" : 25, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 7, \"name\" : \"鞋子7\", \"icon\" : \"shoe/xiezi_7\", \"type\" : 1, \"level\" : 7, \"star\":70, \"exp\" : 30, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 8, \"name\" : \"鞋子8\", \"icon\" : \"shoe/xiezi_8\", \"type\" : 1, \"level\" : 8, \"star\":80, \"exp\" : 35, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 9, \"name\" : \"鞋子9\", \"icon\" : \"shoe/xiezi_9\", \"type\" : 1, \"level\" : 9, \"star\":90, \"exp\" : 40, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 10, \"name\" : \"鞋子10\", \"icon\" : \"shoe/xiezi_10\", \"type\" : 1, \"level\" : 10, \"star\":100, \"exp\" : 45, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 11, \"name\" : \"鞋子11\", \"icon\" : \"shoe/xiezi_11\", \"type\" : 1, \"level\" : 11, \"star\":110, \"exp\" : 50, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 12, \"name\" : \"鞋子12\", \"icon\" : \"shoe/xiezi_12\", \"type\" : 1, \"level\" : 12, \"star\":120, \"exp\" : 60, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 13, \"name\" : \"鞋子13\", \"icon\" : \"shoe/xiezi_13\", \"type\" : 1, \"level\" : 13, \"star\":130, \"exp\" : 70, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 14, \"name\" : \"鞋子14\", \"icon\" : \"shoe/xiezi_14\", \"type\" : 1, \"level\" : 14, \"star\":140, \"exp\" : 80, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 15, \"name\" : \"鞋子15\", \"icon\" : \"shoe/xiezi_15\", \"type\" : 1, \"level\" : 15, \"star\":150, \"exp\" : 90, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 16, \"name\" : \"鞋子16\", \"icon\" : \"shoe/xiezi_16\", \"type\" : 1, \"level\" : 16, \"star\":160, \"exp\" : 100, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 17, \"name\" : \"鞋子17\", \"icon\" : \"shoe/xiezi_17\", \"type\" : 1, \"level\" : 17, \"star\":170, \"exp\" : 110, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 18, \"name\" : \"鞋子18\", \"icon\" : \"shoe/xiezi_18\", \"type\" : 1, \"level\" : 18, \"star\":180, \"exp\" : 120, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 19, \"name\" : \"鞋子19\", \"icon\" : \"shoe/xiezi_19\", \"type\" : 1, \"level\" : 19, \"star\":190, \"exp\" : 130, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 20, \"name\" : \"鞋子20\", \"icon\" : \"shoe/xiezi_20\", \"type\" : 1, \"level\" : 20, \"star\":200, \"exp\" : 140, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 21, \"name\" : \"鞋子21\", \"icon\" : \"shoe/xiezi_21\", \"type\" : 1, \"level\" : 21, \"star\":210, \"exp\" : 150, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 22, \"name\" : \"鞋子22\", \"icon\" : \"shoe/xiezi_22\", \"type\" : 1, \"level\" : 22, \"star\":220, \"exp\" : 160, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 23, \"name\" : \"鞋子23\", \"icon\" : \"shoe/xiezi_23\", \"type\" : 1, \"level\" : 23, \"star\":230, \"exp\" : 170, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 24, \"name\" : \"鞋子24\", \"icon\" : \"shoe/xiezi_24\", \"type\" : 1, \"level\" : 24, \"star\":240, \"exp\" : 180, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 25, \"name\" : \"鞋子25\", \"icon\" : \"shoe/xiezi_25\", \"type\" : 1, \"level\" : 25, \"star\":250, \"exp\" : 190, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 26, \"name\" : \"鞋子26\", \"icon\" : \"shoe/xiezi_26\", \"type\" : 1, \"level\" : 26, \"star\":260, \"exp\" : 200, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 27, \"name\" : \"鞋子27\", \"icon\" : \"shoe/xiezi_27\", \"type\" : 1, \"level\" : 27, \"star\":270, \"exp\" : 220, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 28, \"name\" : \"鞋子28\", \"icon\" : \"shoe/xiezi_28\", \"type\" : 1, \"level\" : 28, \"star\":280, \"exp\" : 240, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 29, \"name\" : \"鞋子29\", \"icon\" : \"shoe/xiezi_29\", \"type\" : 1, \"level\" : 29, \"star\":290, \"exp\" : 260, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 30, \"name\" : \"鞋子30\", \"icon\" : \"shoe/xiezi_30\", \"type\" : 1, \"level\" : 30, \"star\":300, \"exp\" : 280, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 31, \"name\" : \"鞋子31\", \"icon\" : \"shoe/xiezi_31\", \"type\" : 1, \"level\" : 31, \"star\":310, \"exp\" : 300, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 32, \"name\" : \"鞋子32\", \"icon\" : \"shoe/xiezi_32\", \"type\" : 1, \"level\" : 32, \"star\":320, \"exp\" : 330, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 33, \"name\" : \"鞋子33\", \"icon\" : \"shoe/xiezi_33\", \"type\" : 1, \"level\" : 33, \"star\":330, \"exp\" : 360, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 34, \"name\" : \"鞋子34\", \"icon\" : \"shoe/xiezi_34\", \"type\" : 1, \"level\" : 34, \"star\":340, \"exp\" : 390, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 35, \"name\" : \"鞋子35\", \"icon\" : \"shoe/xiezi_35\", \"type\" : 1, \"level\" : 35, \"star\":350, \"exp\" : 420, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 36, \"name\" : \"鞋子36\", \"icon\" : \"shoe/xiezi_36\", \"type\" : 1, \"level\" : 36, \"star\":360, \"exp\" : 450, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 37, \"name\" : \"鞋子37\", \"icon\" : \"shoe/xiezi_37\", \"type\" : 1, \"level\" : 37, \"star\":370, \"exp\" : 500, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 38, \"name\" : \"鞋子38\", \"icon\" : \"shoe/xiezi_38\", \"type\" : 1, \"level\" : 38, \"star\":380, \"exp\" : 550, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 39, \"name\" : \"鞋子39\", \"icon\" : \"shoe/xiezi_39\", \"type\" : 1, \"level\" : 39, \"star\":390, \"exp\" : 600, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 40, \"name\" : \"鞋子40\", \"icon\" : \"shoe/xiezi_40\", \"type\" : 1, \"level\" : 40, \"star\":400, \"exp\" : 800, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 41, \"name\" : \"鞋子41\", \"icon\" : \"shoe/xiezi_41\", \"type\" : 1, \"level\" : 41, \"star\":410, \"exp\" : 1000, \"cost\" : \"[0,0,1]\"}",

		"{\"no\" : 101, \"name\" : \"上衣1\", \"icon\" : \"coat/yifu_1\", \"type\" : 2, \"level\" : 1, \"star\":0, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 102, \"name\" : \"上衣2\", \"icon\" : \"coat/yifu_2\", \"type\" : 2, \"level\" : 2, \"star\":20, \"exp\" : 5, \"cost\" : \"[5]\"}",
		"{\"no\" : 103, \"name\" : \"上衣3\", \"icon\" : \"coat/yifu_3\", \"type\" : 2, \"level\" : 3, \"star\":30, \"exp\" : 10, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 104, \"name\" : \"上衣4\", \"icon\" : \"coat/yifu_4\", \"type\" : 2, \"level\" : 4, \"star\":40, \"exp\" : 15, \"cost\" : \"[0,5]\"}",
		"{\"no\" : 105, \"name\" : \"上衣5\", \"icon\" : \"coat/yifu_5\", \"type\" : 2, \"level\" : 5, \"star\":50, \"exp\" : 20, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 106, \"name\" : \"上衣6\", \"icon\" : \"coat/yifu_6\", \"type\" : 2, \"level\" : 6, \"star\":60, \"exp\" : 25, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 107, \"name\" : \"上衣7\", \"icon\" : \"coat/yifu_7\", \"type\" : 2, \"level\" : 7, \"star\":70, \"exp\" : 30, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 108, \"name\" : \"上衣8\", \"icon\" : \"coat/yifu_8\", \"type\" : 2, \"level\" : 8, \"star\":80, \"exp\" : 35, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 109, \"name\" : \"上衣9\", \"icon\" : \"coat/yifu_9\", \"type\" : 2, \"level\" : 9, \"star\":90, \"exp\" : 40, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 110, \"name\" : \"上衣10\", \"icon\" : \"coat/yifu_10\", \"type\" : 2, \"level\" : 10, \"star\":100, \"exp\" : 45, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 111, \"name\" : \"上衣11\", \"icon\" : \"coat/yifu_11\", \"type\" : 2, \"level\" : 11, \"star\":110, \"exp\" : 50, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 112, \"name\" : \"上衣12\", \"icon\" : \"coat/yifu_12\", \"type\" : 2, \"level\" : 12, \"star\":120, \"exp\" : 60, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 113, \"name\" : \"上衣13\", \"icon\" : \"coat/yifu_13\", \"type\" : 2, \"level\" : 13, \"star\":130, \"exp\" : 70, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 114, \"name\" : \"上衣14\", \"icon\" : \"coat/yifu_14\", \"type\" : 2, \"level\" : 14, \"star\":140, \"exp\" : 80, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 115, \"name\" : \"上衣15\", \"icon\" : \"coat/yifu_15\", \"type\" : 2, \"level\" : 15, \"star\":150, \"exp\" : 90, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 116, \"name\" : \"上衣16\", \"icon\" : \"coat/yifu_16\", \"type\" : 2, \"level\" : 16, \"star\":160, \"exp\" : 100, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 117, \"name\" : \"上衣17\", \"icon\" : \"coat/yifu_17\", \"type\" : 2, \"level\" : 17, \"star\":170, \"exp\" : 110, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 118, \"name\" : \"上衣18\", \"icon\" : \"coat/yifu_18\", \"type\" : 2, \"level\" : 18, \"star\":180, \"exp\" : 120, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 119, \"name\" : \"上衣19\", \"icon\" : \"coat/yifu_19\", \"type\" : 2, \"level\" : 19, \"star\":190, \"exp\" : 130, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 120, \"name\" : \"上衣20\", \"icon\" : \"coat/yifu_20\", \"type\" : 2, \"level\" : 20, \"star\":200, \"exp\" : 140, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 121, \"name\" : \"上衣21\", \"icon\" : \"coat/yifu_21\", \"type\" : 2, \"level\" : 21, \"star\":210, \"exp\" : 150, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 122, \"name\" : \"上衣22\", \"icon\" : \"coat/yifu_22\", \"type\" : 2, \"level\" : 22, \"star\":220, \"exp\" : 160, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 123, \"name\" : \"上衣23\", \"icon\" : \"coat/yifu_23\", \"type\" : 2, \"level\" : 23, \"star\":230, \"exp\" : 170, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 124, \"name\" : \"上衣24\", \"icon\" : \"coat/yifu_24\", \"type\" : 2, \"level\" : 24, \"star\":240, \"exp\" : 180, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 125, \"name\" : \"上衣25\", \"icon\" : \"coat/yifu_25\", \"type\" : 2, \"level\" : 25, \"star\":250, \"exp\" : 190, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 126, \"name\" : \"上衣26\", \"icon\" : \"coat/yifu_26\", \"type\" : 2, \"level\" : 26, \"star\":260, \"exp\" : 200, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 127, \"name\" : \"上衣27\", \"icon\" : \"coat/yifu_27\", \"type\" : 2, \"level\" : 27, \"star\":270, \"exp\" : 220, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 128, \"name\" : \"上衣28\", \"icon\" : \"coat/yifu_28\", \"type\" : 2, \"level\" : 28, \"star\":280, \"exp\" : 240, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 129, \"name\" : \"上衣29\", \"icon\" : \"coat/yifu_29\", \"type\" : 2, \"level\" : 29, \"star\":290, \"exp\" : 260, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 130, \"name\" : \"上衣30\", \"icon\" : \"coat/yifu_30\", \"type\" : 2, \"level\" : 30, \"star\":300, \"exp\" : 280, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 131, \"name\" : \"上衣31\", \"icon\" : \"coat/yifu_31\", \"type\" : 2, \"level\" : 31, \"star\":310, \"exp\" : 300, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 132, \"name\" : \"上衣32\", \"icon\" : \"coat/yifu_32\", \"type\" : 2, \"level\" : 32, \"star\":320, \"exp\" : 330, \"cost\" : \"[0,0,1]\"}",
		// "{\"no\" : 133, \"name\" : \"上衣33\", \"icon\" : \"coat/yifu_33\", \"type\" : 2, \"level\" : 33, \"star\":330, \"exp\" : 360, \"cost\" : \"[0,0,1]\"}",
		// "{\"no\" : 134, \"name\" : \"上衣34\", \"icon\" : \"coat/yifu_34\", \"type\" : 2, \"level\" : 34, \"star\":340, \"exp\" : 390, \"cost\" : \"[0,0,1]\"}",
		// "{\"no\" : 135, \"name\" : \"上衣35\", \"icon\" : \"coat/yifu_35\", \"type\" : 2, \"level\" : 35, \"star\":350, \"exp\" : 420, \"cost\" : \"[0,0,1]\"}",
		// "{\"no\" : 136, \"name\" : \"上衣36\", \"icon\" : \"coat/yifu_36\", \"type\" : 2, \"level\" : 36, \"star\":360, \"exp\" : 450, \"cost\" : \"[0,0,1]\"}",
		// "{\"no\" : 137, \"name\" : \"上衣37\", \"icon\" : \"coat/yifu_37\", \"type\" : 2, \"level\" : 37, \"star\":370, \"exp\" : 500, \"cost\" : \"[0,0,1]\"}",
		// "{\"no\" : 138, \"name\" : \"上衣38\", \"icon\" : \"coat/yifu_38\", \"type\" : 2, \"level\" : 38, \"star\":380, \"exp\" : 550, \"cost\" : \"[0,0,1]\"}",
		// "{\"no\" : 139, \"name\" : \"上衣39\", \"icon\" : \"coat/yifu_39\", \"type\" : 2, \"level\" : 39, \"star\":390, \"exp\" : 600, \"cost\" : \"[0,0,1]\"}",
		// "{\"no\" : 140, \"name\" : \"上衣40\", \"icon\" : \"coat/yifu_40\", \"type\" : 2, \"level\" : 40, \"star\":400, \"exp\" : 800, \"cost\" : \"[0,0,1]\"}",
		// "{\"no\" : 141, \"name\" : \"上衣41\", \"icon\" : \"coat/yifu_41\", \"type\" : 2, \"level\" : 41, \"star\":410, \"exp\" : 1000, \"cost\" : \"[0,0,1]\"}",

		"{\"no\" : 201, \"name\" : \"下衣1\", \"icon\" : \"dress/kuzi_1\", \"type\" : 3, \"level\" : 1, \"star\":0, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 202, \"name\" : \"下衣2\", \"icon\" : \"dress/kuzi_2\", \"type\" : 3, \"level\" : 2, \"star\":20, \"exp\" : 10, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 203, \"name\" : \"下衣3\", \"icon\" : \"dress/kuzi_3\", \"type\" : 3, \"level\" : 3, \"star\":30, \"exp\" : 20, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 204, \"name\" : \"下衣4\", \"icon\" : \"dress/kuzi_4\", \"type\" : 3, \"level\" : 4, \"star\":40, \"exp\" : 30, \"cost\" : \"[0,0,0,1]\"}",
		"{\"no\" : 205, \"name\" : \"下衣5\", \"icon\" : \"dress/kuzi_5\", \"type\" : 3, \"level\" : 5, \"star\":50, \"exp\" : 40, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 206, \"name\" : \"下衣6\", \"icon\" : \"dress/kuzi_6\", \"type\" : 3, \"level\" : 6, \"star\":60, \"exp\" : 50, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 207, \"name\" : \"下衣7\", \"icon\" : \"dress/kuzi_7\", \"type\" : 3, \"level\" : 7, \"star\":70, \"exp\" : 60, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 208, \"name\" : \"下衣8\", \"icon\" : \"dress/kuzi_8\", \"type\" : 3, \"level\" : 8, \"star\":80, \"exp\" : 70, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 209, \"name\" : \"下衣9\", \"icon\" : \"dress/kuzi_9\", \"type\" : 3, \"level\" : 9, \"star\":90, \"exp\" : 80, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 210, \"name\" : \"下衣10\", \"icon\" : \"dress/kuzi_10\", \"type\" : 3, \"level\" : 10, \"star\":100, \"exp\" : 90, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 211, \"name\" : \"下衣11\", \"icon\" : \"dress/kuzi_11\", \"type\" : 3, \"level\" : 11, \"star\":110, \"exp\" : 100, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 212, \"name\" : \"下衣12\", \"icon\" : \"dress/kuzi_12\", \"type\" : 3, \"level\" : 12, \"star\":120, \"exp\" : 110, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 213, \"name\" : \"下衣13\", \"icon\" : \"dress/kuzi_13\", \"type\" : 3, \"level\" : 13, \"star\":130, \"exp\" : 120, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 214, \"name\" : \"下衣14\", \"icon\" : \"dress/kuzi_14\", \"type\" : 3, \"level\" : 14, \"star\":140, \"exp\" : 130, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 215, \"name\" : \"下衣15\", \"icon\" : \"dress/kuzi_15\", \"type\" : 3, \"level\" : 15, \"star\":150, \"exp\" : 140, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 216, \"name\" : \"下衣16\", \"icon\" : \"dress/kuzi_16\", \"type\" : 3, \"level\" : 16, \"star\":160, \"exp\" : 150, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 217, \"name\" : \"下衣17\", \"icon\" : \"dress/kuzi_17\", \"type\" : 3, \"level\" : 17, \"star\":170, \"exp\" : 160, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 218, \"name\" : \"下衣18\", \"icon\" : \"dress/kuzi_18\", \"type\" : 3, \"level\" : 18, \"star\":180, \"exp\" : 170, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 219, \"name\" : \"下衣19\", \"icon\" : \"dress/kuzi_19\", \"type\" : 3, \"level\" : 19, \"star\":190, \"exp\" : 180, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 220, \"name\" : \"下衣20\", \"icon\" : \"dress/kuzi_20\", \"type\" : 3, \"level\" : 20, \"star\":200, \"exp\" : 190, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 221, \"name\" : \"下衣21\", \"icon\" : \"dress/kuzi_21\", \"type\" : 3, \"level\" : 21, \"star\":210, \"exp\" : 200, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 222, \"name\" : \"下衣22\", \"icon\" : \"dress/kuzi_22\", \"type\" : 3, \"level\" : 22, \"star\":220, \"exp\" : 210, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 223, \"name\" : \"下衣23\", \"icon\" : \"dress/kuzi_23\", \"type\" : 3, \"level\" : 23, \"star\":230, \"exp\" : 220, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 224, \"name\" : \"下衣24\", \"icon\" : \"dress/kuzi_24\", \"type\" : 3, \"level\" : 24, \"star\":240, \"exp\" : 230, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 225, \"name\" : \"下衣25\", \"icon\" : \"dress/kuzi_25\", \"type\" : 3, \"level\" : 25, \"star\":250, \"exp\" : 240, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 226, \"name\" : \"下衣26\", \"icon\" : \"dress/kuzi_26\", \"type\" : 3, \"level\" : 26, \"star\":260, \"exp\" : 250, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 227, \"name\" : \"下衣27\", \"icon\" : \"dress/kuzi_27\", \"type\" : 3, \"level\" : 27, \"star\":270, \"exp\" : 260, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 228, \"name\" : \"下衣28\", \"icon\" : \"dress/kuzi_28\", \"type\" : 3, \"level\" : 28, \"star\":280, \"exp\" : 270, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 229, \"name\" : \"下衣29\", \"icon\" : \"dress/kuzi_29\", \"type\" : 3, \"level\" : 29, \"star\":290, \"exp\" : 280, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 230, \"name\" : \"下衣30\", \"icon\" : \"dress/kuzi_30\", \"type\" : 3, \"level\" : 30, \"star\":300, \"exp\" : 290, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 231, \"name\" : \"下衣31\", \"icon\" : \"dress/kuzi_31\", \"type\" : 3, \"level\" : 31, \"star\":310, \"exp\" : 300, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 232, \"name\" : \"下衣32\", \"icon\" : \"dress/kuzi_32\", \"type\" : 3, \"level\" : 32, \"star\":320, \"exp\" : 310, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 233, \"name\" : \"下衣33\", \"icon\" : \"dress/kuzi_33\", \"type\" : 3, \"level\" : 33, \"star\":330, \"exp\" : 320, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 234, \"name\" : \"下衣34\", \"icon\" : \"dress/kuzi_34\", \"type\" : 3, \"level\" : 34, \"star\":340, \"exp\" : 330, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 235, \"name\" : \"下衣35\", \"icon\" : \"dress/kuzi_35\", \"type\" : 3, \"level\" : 35, \"star\":350, \"exp\" : 340, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 236, \"name\" : \"下衣36\", \"icon\" : \"dress/kuzi_36\", \"type\" : 3, \"level\" : 36, \"star\":360, \"exp\" : 350, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 237, \"name\" : \"下衣37\", \"icon\" : \"dress/kuzi_37\", \"type\" : 3, \"level\" : 37, \"star\":370, \"exp\" : 360, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 238, \"name\" : \"下衣38\", \"icon\" : \"dress/kuzi_38\", \"type\" : 3, \"level\" : 38, \"star\":380, \"exp\" : 370, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 239, \"name\" : \"下衣39\", \"icon\" : \"dress/kuzi_39\", \"type\" : 3, \"level\" : 39, \"star\":390, \"exp\" : 380, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 240, \"name\" : \"下衣40\", \"icon\" : \"dress/kuzi_40\", \"type\" : 3, \"level\" : 40, \"star\":400, \"exp\" : 390, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 241, \"name\" : \"下衣41\", \"icon\" : \"dress/kuzi_41\", \"type\" : 3, \"level\" : 41, \"star\":410, \"exp\" : 400, \"cost\" : \"[0,0,0,0,1]\"}",
		// "{\"no\" : 242, \"name\" : \"下衣42\", \"icon\" : \"dress/kuzi_42\", \"type\" : 3, \"level\" : 42, \"star\":420, \"exp\" : 410, \"cost\" : \"[0,0,0,0,1]\"}",
		// "{\"no\" : 243, \"name\" : \"下衣43\", \"icon\" : \"dress/kuzi_43\", \"type\" : 3, \"level\" : 43, \"star\":430, \"exp\" : 420, \"cost\" : \"[0,0,0,0,1]\"}",
		// "{\"no\" : 244, \"name\" : \"下衣44\", \"icon\" : \"dress/kuzi_44\", \"type\" : 3, \"level\" : 44, \"star\":440, \"exp\" : 430, \"cost\" : \"[0,0,0,0,1]\"}",
		// "{\"no\" : 245, \"name\" : \"下衣45\", \"icon\" : \"dress/kuzi_45\", \"type\" : 3, \"level\" : 45, \"star\":450, \"exp\" : 440, \"cost\" : \"[0,0,0,0,1]\"}",

		"{\"no\" : 301, \"name\" : \"头发1\", \"icon\" : \"faxing/toufa_1\", \"type\" : 4, \"level\" : 1, \"star\":0, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 302, \"name\" : \"头发2\", \"icon\" : \"faxing/toufa_2\", \"type\" : 4, \"level\" : 2, \"star\":20, \"exp\" : 5, \"cost\" : \"[5]\"}",
		"{\"no\" : 303, \"name\" : \"头发3\", \"icon\" : \"faxing/toufa_3\", \"type\" : 4, \"level\" : 3, \"star\":30, \"exp\" : 10, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 304, \"name\" : \"头发4\", \"icon\" : \"faxing/toufa_4\", \"type\" : 4, \"level\" : 4, \"star\":40, \"exp\" : 15, \"cost\" : \"[5,1]\"}",
		"{\"no\" : 305, \"name\" : \"头发5\", \"icon\" : \"faxing/toufa_5\", \"type\" : 4, \"level\" : 5, \"star\":50, \"exp\" : 20, \"cost\" : \"[0,2]\"}",
		"{\"no\" : 306, \"name\" : \"头发6\", \"icon\" : \"faxing/toufa_6\", \"type\" : 4, \"level\" : 6, \"star\":60, \"exp\" : 25, \"cost\" : \"[5,2]\"}",
		"{\"no\" : 307, \"name\" : \"头发7\", \"icon\" : \"faxing/toufa_7\", \"type\" : 4, \"level\" : 7, \"star\":70, \"exp\" : 30, \"cost\" : \"[0,3]\"}",
		"{\"no\" : 308, \"name\" : \"头发8\", \"icon\" : \"faxing/toufa_8\", \"type\" : 4, \"level\" : 8, \"star\":80, \"exp\" : 35, \"cost\" : \"[5,3]\"}",
		"{\"no\" : 309, \"name\" : \"头发9\", \"icon\" : \"faxing/toufa_9\", \"type\" : 4, \"level\" : 9, \"star\":90, \"exp\" : 40, \"cost\" : \"[0,4]\"}",
		"{\"no\" : 310, \"name\" : \"头发10\", \"icon\" : \"faxing/toufa_10\", \"type\" : 4, \"level\" : 10, \"star\":100, \"exp\" : 45, \"cost\" : \"[5,4]\"}",
		"{\"no\" : 311, \"name\" : \"头发11\", \"icon\" : \"faxing/toufa_11\", \"type\" : 4, \"level\" : 11, \"star\":110, \"exp\" : 50, \"cost\" : \"[0,5]\"}",
		"{\"no\" : 312, \"name\" : \"头发12\", \"icon\" : \"faxing/toufa_12\", \"type\" : 4, \"level\" : 12, \"star\":120, \"exp\" : 55, \"cost\" : \"[5,5]\"}",
		"{\"no\" : 313, \"name\" : \"头发13\", \"icon\" : \"faxing/toufa_13\", \"type\" : 4, \"level\" : 13, \"star\":130, \"exp\" : 60, \"cost\" : \"[0,6]\"}",
		"{\"no\" : 314, \"name\" : \"头发14\", \"icon\" : \"faxing/toufa_14\", \"type\" : 4, \"level\" : 14, \"star\":140, \"exp\" : 65, \"cost\" : \"[5,6]\"}",
		"{\"no\" : 315, \"name\" : \"头发15\", \"icon\" : \"faxing/toufa_15\", \"type\" : 4, \"level\" : 15, \"star\":160, \"exp\" : 70, \"cost\" : \"[0,7]\"}",
		"{\"no\" : 316, \"name\" : \"头发16\", \"icon\" : \"faxing/toufa_16\", \"type\" : 4, \"level\" : 16, \"star\":160, \"exp\" : 75, \"cost\" : \"[5,7]\"}",
		"{\"no\" : 317, \"name\" : \"头发17\", \"icon\" : \"faxing/toufa_17\", \"type\" : 4, \"level\" : 17, \"star\":170, \"exp\" : 80, \"cost\" : \"[0,8]\"}",
		"{\"no\" : 318, \"name\" : \"头发18\", \"icon\" : \"faxing/toufa_18\", \"type\" : 4, \"level\" : 18, \"star\":180, \"exp\" : 85, \"cost\" : \"[5,8]\"}",
		"{\"no\" : 319, \"name\" : \"头发19\", \"icon\" : \"faxing/toufa_19\", \"type\" : 4, \"level\" : 19, \"star\":190, \"exp\" : 90, \"cost\" : \"[0,9]\"}",
		// "{\"no\" : 320, \"name\" : \"头发20\", \"icon\" : \"faxing/toufa_20\", \"type\" : 4, \"level\" : 20, \"star\":200, \"exp\" : 95, \"cost\" : \"[5,9]\"}",

		"{\"no\" : 401, \"name\" : \"围巾1\", \"icon\" : \"player_scarf1\", \"type\" : 5, \"level\" : 1, \"star\":20, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 402, \"name\" : \"围巾2\", \"icon\" : \"player_scarf2\", \"type\" : 5, \"level\" : 2, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"no\" : 501, \"name\" : \"包包1\", \"icon\" : \"player_bag1\", \"type\" : 6, \"level\" : 1, \"star\":20, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 502, \"name\" : \"包包2\", \"icon\" : \"player_bag2\", \"type\" : 6, \"level\" : 2, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"no\" : 701, \"name\" : \"宠物1\", \"icon\" : \"pet/chongwu_01\", \"type\" : 7, \"level\" : 1, \"star\":10, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 702, \"name\" : \"宠物2\", \"icon\" : \"pet/chongwu_02\", \"type\" : 7, \"level\" : 2, \"star\":20, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 703, \"name\" : \"宠物3\", \"icon\" : \"pet/chongwu_03\", \"type\" : 7, \"level\" : 3, \"star\":30, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 704, \"name\" : \"宠物4\", \"icon\" : \"pet/chongwu_04\", \"type\" : 7, \"level\" : 4, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 705, \"name\" : \"宠物5\", \"icon\" : \"pet/chongwu_05\", \"type\" : 7, \"level\" : 5, \"star\":50, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 706, \"name\" : \"宠物6\", \"icon\" : \"pet/chongwu_06\", \"type\" : 7, \"level\" : 6, \"star\":60, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 707, \"name\" : \"宠物7\", \"icon\" : \"pet/chongwu_07\", \"type\" : 7, \"level\" : 7, \"star\":70, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 708, \"name\" : \"宠物8\", \"icon\" : \"pet/chongwu_08\", \"type\" : 7, \"level\" : 8, \"star\":80, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"no\" : 1001, \"name\" : \"留言板\", \"icon\" : \"gift0\", \"type\" : 10, \"diamond\":0, \"favour\" : 1, \"reward\" : 0}",
	}

	for _, configStr := range configs {
		// log.Debug("Load ClothConfigs  db %v ", configStr)
		var s ConfigCloth
		json.Unmarshal([]byte(configStr), &s)
		m.db.Create(&s)
	}
}

// InitializeGiftConfig 初始化衣服配置
func (m *Module) InitializeGiftConfig() {
	m.db.Unscoped().Delete(&ConfigGift{})

	var configs = []string{
		"{\"no\" : 1, \"name\" : \"留言板\", \"icon\" : \"gift/gift0\", \"diamond\":0, \"favour\" : 1, \"reward\" : 0}",
		"{\"no\" : 2, \"name\" : \"口红\", \"icon\" : \"gift/gift1\", \"diamond\":100, \"favour\" : 10, \"reward\" : 70}",
		"{\"no\" : 3, \"name\" : \"轮船\", \"icon\" : \"gift/gift2\", \"diamond\":200, \"favour\" : 20, \"reward\" : 140}",
		"{\"no\" : 4, \"name\" : \"爱心\", \"icon\" : \"gift/gift3\", \"diamond\":300, \"favour\" : 30, \"reward\" : 210}",
		"{\"no\" : 5, \"name\" : \"鲜花\", \"icon\" : \"gift/gift4\", \"diamond\":400, \"favour\" : 40, \"reward\" : 280}",
		"{\"no\" : 6, \"name\" : \"甜品\", \"icon\" : \"gift/gift5\", \"diamond\":500, \"favour\" : 50, \"reward\" : 350}",
		"{\"no\" : 7, \"name\" : \"跑车\", \"icon\" : \"gift/gift6\", \"diamond\":600, \"favour\" : 60, \"reward\" : 420}",
		"{\"no\" : 8, \"name\" : \"钻戒\", \"icon\" : \"gift/gift7\", \"diamond\":700, \"favour\" : 70, \"reward\" : 490}",
		"{\"no\" : 9, \"name\" : \"飞机\", \"icon\" : \"gift/gift8\", \"diamond\":800, \"favour\" : 80, \"reward\" : 560}",
	}

	for _, configStr := range configs {
		// log.Debug("Load ClothConfigs  db %v ", configStr)
		var s ConfigGift
		json.Unmarshal([]byte(configStr), &s)
		m.db.Create(&s)
	}
}

// InitializeBarrages 初始化弹幕
func (m *Module) InitializeBarrages() {
	m.db.Unscoped().Delete(&BarrageReport{})

	var configs = []string{
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"新司机李浩宇不请自来，望各位带哥海涵\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"叫声龙哥，留支付宝 100红包\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"主播fps好高啊，网速这么卡怎么玩？\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"主播ping好低哦，电脑太差了吧\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"嗨呀，这是最骚的\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"好尴尬；好紧张\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"上单我只服PDD\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"卢本伟牛逼\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"弹幕掩护\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"小姐姐\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"英雄联盟抄袭王者荣耀\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"1、BGM别闹！\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"2、说出来不怕丢人，如果能重来 ，我要当明星\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"3、范冰冰邀请你加入游戏\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"4、哦买噶！这是啥神仙发色，太酷了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"5、打卡签到，日给你温暖，我给你礼物\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"6、我弃坑了，但又回来了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"7、艾玛我老婆，全身名牌，亮瞎你们狗眼\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"8、巴拉巴拉小美女，变身\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"9、真的，脖子以下全是腿\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"10、你想把衣服放头上吧\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"11、这白裙的链接有吗 在线求\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"12、穿上粗布衣，拿起大哥大\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"13、原来GM是网络管理员啊！我还以为GM是肛门的缩写呢？ \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"14、我先吐口血，你们继续刷弹幕……\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"15、我们不一样不一样不一样\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"16、送你一份纯情小礼，有效期：一辈子\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"17、魔法小樱吗？！！！\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"18、有人吗，你的礼物到了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"19、这模特腿可以上天了 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"20、这头发是一个月没洗吗？\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"21、前方高能 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"22、哦哟 不错哦\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"23、尽管周围很多刺,但是有我在,你必能安心 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"24、隔着屏幕我都觉得刺激\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"25、前方高能预警，请非单身人员撤离 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"26、我做错了什么 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"27、恍恍惚惚哈哈哈，对不起我尽力了，没绷住。。。 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"28、买 买 买！！！\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"29、打卡签到 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"30、美女能换双鞋吗，这双不好看\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"31、玫瑰玫瑰！无论你是什么形态\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"32、哈哈哈哈，那件绿衣服笑死我\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"33、老板，这里的衣服我全包了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"34、你腰间的带子是我永远的信仰哈哈哈 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"35、大吉大利，晚上吃鸡…..\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"36、欢迎新玩家\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"37、本美已上线\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"38、都是人才啊，在下实在是佩服 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"39、哎，怎么装扮不对称 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"40、我头上有电话，可以随时打给罗志祥。。。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"41、这游戏真好玩。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"42、你认我做大哥，我教你玩游戏。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"43、这游戏策划是谁，出来。。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"44、那个丸子头，笑死本宝宝了。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"45、这游戏做的真好，老夫的少女心啊\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"46、画风好萌 ，好萌啊\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"47、做的不错，支持一下 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"48、姐姐，你忘记给我零用钱了 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"49、救命，前方浪大！！波涛汹涌\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"50、我有谢霆锋电话，你信吗\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"51、游戏名字好吸引人啊\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"52、竟然可以这么好玩，偶滴天\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"53、交个朋友不\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"54、装逼你妹！\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"55、这妹子由我亲自操刀\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"56、一首凉凉送给你们。。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"57、福建人民发来贺电\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"58、黑龙江人民发来贺电\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"59、55级，求超越。。。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"60、你的青春值几块钱！\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"61、怎么不合成眼镜嘴巴耳朵。。。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"62、我放着客厅电视不看，跑来玩这游戏也是蛮拼的\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"63、蓝孩子也喜欢玩\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"64、熟悉的味道\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"65、开口脆,溜了溜了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"66、所有高级衣服穿上是什么的体验\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"67、说不丑的画画肯定不好\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"68、粉红色有点迷人\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"69、背景图是什么\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"70、画风清奇\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"71、人物挺好看的\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"72、我还在玩，老妈喊我吃饭了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"73、经费不够，品质来凑\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"74、钻石数量引起强烈不满\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"75、一进来就送钻石，靠谱\"}",
	}

	for _, configStr := range configs {
		// log.Debug("Load ClothConfigs  db %v ", configStr)
		var s BarrageReport
		json.Unmarshal([]byte(configStr), &s)
		m.db.Create(&s)
	}
}
