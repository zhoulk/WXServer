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
	m.InitializeCPConfig()
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
		"{\"no\" : 1, \"name\" : \"小萌新\", \"icon\" : \"\", \"star\" : 1, \"level\" : 1}",
		"{\"no\" : 2, \"name\" : \"校花\", \"icon\" : \"\", \"star\" : 20, \"level\" : 2}",
		"{\"no\" : 3, \"name\" : \"龙套演员\", \"icon\" : \"\", \"star\" :40, \"level\" : 3}",
		"{\"no\" : 4, \"name\" : \"平面模特\", \"icon\" : \"\", \"star\" : 60, \"level\" : 4}",
		"{\"no\" : 5, \"name\" : \"时装模特\", \"icon\" : \"\", \"star\" : 80, \"level\" : 5}",
		"{\"no\" : 6, \"name\" : \"三线艺人\", \"icon\" : \"\", \"star\" : 100, \"level\" : 6}",
		"{\"no\" : 7, \"name\" : \"二线明星\", \"icon\" : \"\", \"star\" : 120, \"level\" : 7}",
		"{\"no\" : 8, \"name\" : \"一线大咖\", \"icon\" : \"\", \"star\" : 140, \"level\" : 8}",
		"{\"no\" : 9, \"name\" : \"天后\", \"icon\" : \"\", \"star\" : 160, \"level\" : 9}",
		"{\"no\" : 10, \"name\" : \"国际巨星\", \"icon\" : \"\", \"star\" : 200, \"level\" : 109}",
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
		"{\"no\" : 1, \"name\" : \"巴厘岛\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/1.jpg\", \"star\" : 1, \"level\" : 1}",
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

// InitializeCPConfig 初始化CP配置
func (m *Module) InitializeCPConfig() {
	m.db.Unscoped().Delete(&ConfigCP{})

	var configs = []string{
		"{\"no\" : 1, \"name\" : \"温柔可人\", \"icon\" : \"cp_001\", \"image\" : \"https://mudgame.com.cn/wx/res/WXGame/101/cp/001.png\", \"star\" : 1, \"level\" : 1, \"type\" : 1, \"rate\" : \"[0,0,1]\"}",
		"{\"no\" : 2, \"name\" : \"性感靓丽\", \"icon\" : \"cp_002\", \"image\" : \"https://mudgame.com.cn/wx/res/WXGame/101/cp/002.png\", \"star\" : 50, \"level\" : 2, \"type\" : 1, \"rate\" : \"[0,0,2]\"}",
		"{\"no\" : 3, \"name\" : \"贤惠才女\", \"icon\" : \"cp_003\", \"image\" : \"https://mudgame.com.cn/wx/res/WXGame/101/cp/003.png\", \"star\" : 100, \"level\" : 3, \"type\" : 1, \"rate\" : \"[0,0,3]\"}",
		"{\"no\" : 4, \"name\" : \"知性丽人\", \"icon\" : \"cp_004\", \"image\" : \"https://mudgame.com.cn/wx/res/WXGame/101/cp/004.png\", \"star\" : 150, \"level\" : 4, \"type\" : 1, \"rate\" : \"[0,0,4]\"}",
		"{\"no\" : 5, \"name\" : \"懵懂萝莉\", \"icon\" : \"cp_005\", \"image\" : \"https://mudgame.com.cn/wx/res/WXGame/101/cp/005.png\", \"star\" : 200, \"level\" : 5, \"type\" : 1, \"rate\" : \"[0,0,5]\"}",
		"{\"no\" : 100, \"name\" : \"暖香直男\", \"icon\" : \"cp_006\", \"image\" : \"https://mudgame.com.cn/wx/res/WXGame/101/cp/006.png\", \"star\" : 1, \"level\" : 1, \"type\" : 2, \"rate\" : \"[0,0,1]\"}",
		"{\"no\" : 101, \"name\" : \"金财海归\", \"icon\" : \"cp_007\", \"image\" : \"https://mudgame.com.cn/wx/res/WXGame/101/cp/007.png\", \"star\" : 50, \"level\" : 2, \"type\" : 2, \"rate\" : \"[0,0,2]\"}",
		"{\"no\" : 102, \"name\" : \"时尚达人\", \"icon\" : \"cp_008\", \"image\" : \"https://mudgame.com.cn/wx/res/WXGame/101/cp/008.png\", \"star\" : 150, \"level\" : 3, \"type\" : 2, \"rate\" : \"[0,0,3]\"}",
		"{\"no\" : 103, \"name\" : \"霸道总裁\", \"icon\" : \"cp_009\", \"image\" : \"https://mudgame.com.cn/wx/res/WXGame/101/cp/009.png\", \"star\" : 250, \"level\" : 4, \"type\" : 2, \"rate\" : \"[0,0,4]\"}",
		"{\"no\" : 104, \"name\" : \"文艺青年\", \"icon\" : \"cp_010\", \"image\" : \"https://mudgame.com.cn/wx/res/WXGame/101/cp/010.png\", \"star\" : 350, \"level\" : 5, \"type\" : 2, \"rate\" : \"[0,0,5]\"}",
	}

	for _, configStr := range configs {
		var s ConfigCP
		json.Unmarshal([]byte(configStr), &s)
		m.db.Create(&s)
	}
}

// InitializeClothConfig 初始化衣服配置
func (m *Module) InitializeClothConfig() {
	m.db.Unscoped().Delete(&ConfigCloth{})

	var configs = []string{
		"{\"no\" : 1, \"name\" : \"水蜜桃\", \"brand\":\"范思哲\", \"icon\" : \"shoe/xiezi_001\", \"type\" : 1, \"level\" : 1, \"star\":0, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 2, \"name\" : \"金色蔷薇\", \"brand\":\"范思哲\", \"icon\" : \"shoe/xiezi_002\", \"type\" : 1, \"level\" : 2, \"star\":20, \"exp\" : 5, \"cost\" : \"[50]\"}",
		"{\"no\" : 3, \"name\" : \"花语\", \"brand\":\"范思哲\", \"icon\" : \"shoe/xiezi_003\", \"type\" : 1, \"level\" : 3, \"star\":30, \"exp\" : 10, \"cost\" : \"[100]\"}",
		"{\"no\" : 4, \"name\" : \"浅唱\", \"brand\":\"范思哲\", \"icon\" : \"shoe/xiezi_004\", \"type\" : 1, \"level\" : 4, \"star\":40, \"exp\" : 15, \"cost\" : \"[200]\"}",
		"{\"no\" : 5, \"name\" : \"万人迷\", \"brand\":\"范思哲\", \"icon\" : \"shoe/xiezi_005\", \"type\" : 1, \"level\" : 5, \"star\":50, \"exp\" : 20, \"cost\" : \"[400]\"}",
		"{\"no\" : 6, \"name\" : \"国色天香\", \"brand\":\"迪奥\", \"icon\" : \"shoe/xiezi_006\", \"type\" : 1, \"level\" : 6, \"star\":60, \"exp\" : 25, \"cost\" : \"[800]\"}",
		"{\"no\" : 7, \"name\" : \"马蹄莲\", \"brand\":\"迪奥\", \"icon\" : \"shoe/xiezi_007\", \"type\" : 1, \"level\" : 7, \"star\":70, \"exp\" : 30, \"cost\" : \"[600,1]\"}",
		"{\"no\" : 8, \"name\" : \"倾城\", \"brand\":\"迪奥\", \"icon\" : \"shoe/xiezi_008\", \"type\" : 1, \"level\" : 8, \"star\":80, \"exp\" : 35, \"cost\" : \"[200,3]\"}",
		"{\"no\" : 9, \"name\" : \"许你爱情\", \"brand\":\"迪奥\", \"icon\" : \"shoe/xiezi_009\", \"type\" : 1, \"level\" : 9, \"star\":90, \"exp\" : 40, \"cost\" : \"[400,6]\"}",
		"{\"no\" : 10, \"name\" : \"陶醉\", \"brand\":\"迪奥\", \"icon\" : \"shoe/xiezi_010\", \"type\" : 1, \"level\" : 10, \"star\":100, \"exp\" : 45, \"cost\" : \"[800,12]\"}",
		"{\"no\" : 11, \"name\" : \"昨夜星辰\", \"brand\":\"普拉达\", \"icon\" : \"shoe/xiezi_011\", \"type\" : 1, \"level\" : 11, \"star\":110, \"exp\" : 50, \"cost\" : \"[600,25]\"}",
		"{\"no\" : 12, \"name\" : \"孤舟月\", \"brand\":\"普拉达\", \"icon\" : \"shoe/xiezi_012\", \"type\" : 1, \"level\" : 12, \"star\":120, \"exp\" : 60, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 13, \"name\" : \"水晶鞋\", \"brand\":\"普拉达\", \"icon\" : \"shoe/xiezi_013\", \"type\" : 1, \"level\" : 13, \"star\":130, \"exp\" : 70, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 14, \"name\" : \"黑色玫瑰\", \"brand\":\"普拉达\", \"icon\" : \"shoe/xiezi_014\", \"type\" : 1, \"level\" : 14, \"star\":140, \"exp\" : 80, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 15, \"name\" : \"悦君颜\", \"brand\":\"普拉达\", \"icon\" : \"shoe/xiezi_015\", \"type\" : 1, \"level\" : 15, \"star\":150, \"exp\" : 90, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 16, \"name\" : \"君临\", \"brand\":\"博柏利\", \"icon\" : \"shoe/xiezi_016\", \"type\" : 1, \"level\" : 16, \"star\":160, \"exp\" : 100, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 17, \"name\" : \"话春秋\", \"brand\":\"博柏利\", \"icon\" : \"shoe/xiezi_017\", \"type\" : 1, \"level\" : 17, \"star\":170, \"exp\" : 110, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 18, \"name\" : \"故人梦\", \"brand\":\"博柏利\", \"icon\" : \"shoe/xiezi_018\", \"type\" : 1, \"level\" : 18, \"star\":180, \"exp\" : 120, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 19, \"name\" : \"柔情\", \"brand\":\"博柏利\", \"icon\" : \"shoe/xiezi_019\", \"type\" : 1, \"level\" : 19, \"star\":190, \"exp\" : 130, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 20, \"name\" : \"清风徐来\", \"brand\":\"博柏利\", \"icon\" : \"shoe/xiezi_020\", \"type\" : 1, \"level\" : 20, \"star\":200, \"exp\" : 140, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 21, \"name\" : \"青青草\", \"brand\":\"古驰\", \"icon\" : \"shoe/xiezi_021\", \"type\" : 1, \"level\" : 21, \"star\":210, \"exp\" : 150, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 22, \"name\" : \"繁华落尽\", \"brand\":\"古驰\", \"icon\" : \"shoe/xiezi_022\", \"type\" : 1, \"level\" : 22, \"star\":220, \"exp\" : 160, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 23, \"name\" : \"木婉清\", \"brand\":\"古驰\", \"icon\" : \"shoe/xiezi_023\", \"type\" : 1, \"level\" : 23, \"star\":230, \"exp\" : 170, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 24, \"name\" : \"风满楼\", \"brand\":\"古驰\", \"icon\" : \"shoe/xiezi_024\", \"type\" : 1, \"level\" : 24, \"star\":240, \"exp\" : 180, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 25, \"name\" : \"黑涩会\", \"brand\":\"古驰\", \"icon\" : \"shoe/xiezi_025\", \"type\" : 1, \"level\" : 25, \"star\":250, \"exp\" : 190, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 26, \"name\" : \"微笑\", \"brand\":\"香奈儿\", \"icon\" : \"shoe/xiezi_026\", \"type\" : 1, \"level\" : 26, \"star\":260, \"exp\" : 200, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 27, \"name\" : \"初春\", \"brand\":\"香奈儿\", \"icon\" : \"shoe/xiezi_027\", \"type\" : 1, \"level\" : 27, \"star\":270, \"exp\" : 220, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 28, \"name\" : \"寄相思\", \"brand\":\"香奈儿\", \"icon\" : \"shoe/xiezi_028\", \"type\" : 1, \"level\" : 28, \"star\":280, \"exp\" : 240, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 29, \"name\" : \"风尘四起\", \"brand\":\"香奈儿\", \"icon\" : \"shoe/xiezi_029\", \"type\" : 1, \"level\" : 29, \"star\":290, \"exp\" : 260, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 30, \"name\" : \"热恋红花\", \"brand\":\"香奈儿\", \"icon\" : \"shoe/xiezi_030\", \"type\" : 1, \"level\" : 30, \"star\":300, \"exp\" : 280, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 31, \"name\" : \"枫景\", \"brand\":\"爱马仕\", \"icon\" : \"shoe/xiezi_031\", \"type\" : 1, \"level\" : 31, \"star\":310, \"exp\" : 300, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 32, \"name\" : \"淑女郎\", \"brand\":\"爱马仕\", \"icon\" : \"shoe/xiezi_032\", \"type\" : 1, \"level\" : 32, \"star\":320, \"exp\" : 330, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 33, \"name\" : \"海的颜色\", \"brand\":\"爱马仕\", \"icon\" : \"shoe/xiezi_033\", \"type\" : 1, \"level\" : 33, \"star\":330, \"exp\" : 360, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 34, \"name\" : \"黑夜\", \"brand\":\"爱马仕\", \"icon\" : \"shoe/xiezi_034\", \"type\" : 1, \"level\" : 34, \"star\":340, \"exp\" : 390, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 35, \"name\" : \"花开半夏\", \"brand\":\"爱马仕\", \"icon\" : \"shoe/xiezi_035\", \"type\" : 1, \"level\" : 35, \"star\":350, \"exp\" : 420, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 36, \"name\" : \"白日梦\", \"brand\":\"路易威登\", \"icon\" : \"shoe/xiezi_036\", \"type\" : 1, \"level\" : 36, \"star\":360, \"exp\" : 450, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 37, \"name\" : \"听海澜\", \"brand\":\"路易威登\", \"icon\" : \"shoe/xiezi_037\", \"type\" : 1, \"level\" : 37, \"star\":370, \"exp\" : 500, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 38, \"name\" : \"白月光\", \"brand\":\"路易威登\", \"icon\" : \"shoe/xiezi_038\", \"type\" : 1, \"level\" : 38, \"star\":380, \"exp\" : 550, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 39, \"name\" : \"孤舟月\", \"brand\":\"路易威登\", \"icon\" : \"shoe/xiezi_039\", \"type\" : 1, \"level\" : 39, \"star\":390, \"exp\" : 600, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 40, \"name\" : \"兔儿洛可可\", \"brand\":\"路易威登\", \"icon\" : \"shoe/xiezi_040\", \"type\" : 1, \"level\" : 40, \"star\":400, \"exp\" : 800, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 41, \"name\" : \"浮梦鲸\", \"brand\":\"路易威登\", \"icon\" : \"shoe/xiezi_041\", \"type\" : 1, \"level\" : 41, \"star\":410, \"exp\" : 1000, \"cost\" : \"[0,0,1]\"}",

		"{\"no\" : 101, \"name\" : \"青春\", \"brand\":\"范思哲\", \"icon\" : \"coat/yifu_001\", \"type\" : 2, \"level\" : 1, \"star\":0, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 102, \"name\" : \"小蜜蜂\", \"brand\":\"范思哲\", \"icon\" : \"coat/yifu_002\", \"type\" : 2, \"level\" : 2, \"star\":20, \"exp\" : 5, \"cost\" : \"[50]\"}",
		"{\"no\" : 103, \"name\" : \"镜湖月\", \"brand\":\"范思哲\", \"icon\" : \"coat/yifu_003\", \"type\" : 2, \"level\" : 3, \"star\":30, \"exp\" : 10, \"cost\" : \"[100]\"}",
		"{\"no\" : 104, \"name\" : \"劲酷少女\", \"brand\":\"范思哲\", \"icon\" : \"coat/yifu_004\", \"type\" : 2, \"level\" : 4, \"star\":40, \"exp\" : 15, \"cost\" : \"[200]\"}",
		"{\"no\" : 105, \"name\" : \"金粉披肩\", \"brand\":\"迪奥\", \"icon\" : \"coat/yifu_005\", \"type\" : 2, \"level\" : 5, \"star\":50, \"exp\" : 20, \"cost\" : \"[400]\"}",
		"{\"no\" : 106, \"name\" : \"锦华焰火\", \"brand\":\"迪奥\", \"icon\" : \"coat/yifu_006\", \"type\" : 2, \"level\" : 6, \"star\":60, \"exp\" : 25, \"cost\" : \"[800]\"}",
		"{\"no\" : 107, \"name\" : \"热浪\", \"brand\":\"迪奥\", \"icon\" : \"coat/yifu_007\", \"type\" : 2, \"level\" : 7, \"star\":70, \"exp\" : 30, \"cost\" : \"[600,1]\"}",
		"{\"no\" : 108, \"name\" : \"遇见\", \"brand\":\"迪奥\", \"icon\" : \"coat/yifu_008\", \"type\" : 2, \"level\" : 8, \"star\":80, \"exp\" : 35, \"cost\" : \"[200,3]\"}",
		"{\"no\" : 109, \"name\" : \"素色条纹\", \"brand\":\"普拉达\", \"icon\" : \"coat/yifu_009\", \"type\" : 2, \"level\" : 9, \"star\":90, \"exp\" : 40, \"cost\" : \"[400,6]\"}",
		"{\"no\" : 110, \"name\" : \"半颗心\", \"brand\":\"普拉达\", \"icon\" : \"coat/yifu_010\", \"type\" : 2, \"level\" : 10, \"star\":100, \"exp\" : 45, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 111, \"name\" : \"彼岸花开\", \"brand\":\"普拉达\", \"icon\" : \"coat/yifu_011\", \"type\" : 2, \"level\" : 11, \"star\":110, \"exp\" : 50, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 112, \"name\" : \"青山黛\", \"brand\":\"普拉达\", \"icon\" : \"coat/yifu_012\", \"type\" : 2, \"level\" : 12, \"star\":120, \"exp\" : 60, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 113, \"name\" : \"挽清风\", \"brand\":\"博柏利\", \"icon\" : \"coat/yifu_013\", \"type\" : 2, \"level\" : 13, \"star\":130, \"exp\" : 70, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 114, \"name\" : \"格桑花\", \"brand\":\"博柏利\", \"icon\" : \"coat/yifu_014\", \"type\" : 2, \"level\" : 14, \"star\":140, \"exp\" : 80, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 115, \"name\" : \"暮光缘\", \"brand\":\"博柏利\", \"icon\" : \"coat/yifu_015\", \"type\" : 2, \"level\" : 15, \"star\":150, \"exp\" : 90, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 116, \"name\" : \"简单爱\", \"brand\":\"博柏利\", \"icon\" : \"coat/yifu_016\", \"type\" : 2, \"level\" : 16, \"star\":160, \"exp\" : 100, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 117, \"name\" : \"书香\", \"brand\":\"古驰\", \"icon\" : \"coat/yifu_017\", \"type\" : 2, \"level\" : 17, \"star\":170, \"exp\" : 110, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 118, \"name\" : \"锦风雪\", \"brand\":\"古驰\", \"icon\" : \"coat/yifu_018\", \"type\" : 2, \"level\" : 18, \"star\":180, \"exp\" : 120, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 119, \"name\" : \"寄相思\", \"brand\":\"古驰\", \"icon\" : \"coat/yifu_019\", \"type\" : 2, \"level\" : 19, \"star\":190, \"exp\" : 130, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 120, \"name\" : \"夏天\", \"brand\":\"古驰\", \"icon\" : \"coat/yifu_020\", \"type\" : 2, \"level\" : 20, \"star\":200, \"exp\" : 140, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 121, \"name\" : \"麋鹿\", \"brand\":\"香奈儿\", \"icon\" : \"coat/yifu_021\", \"type\" : 2, \"level\" : 21, \"star\":210, \"exp\" : 150, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 122, \"name\" : \"倩女丽影\", \"brand\":\"香奈儿\", \"icon\" : \"coat/yifu_022\", \"type\" : 2, \"level\" : 22, \"star\":220, \"exp\" : 160, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 123, \"name\" : \"素青衣\", \"brand\":\"香奈儿\", \"icon\" : \"coat/yifu_023\", \"type\" : 2, \"level\" : 23, \"star\":230, \"exp\" : 170, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 124, \"name\" : \"霓裳挽歌\", \"brand\":\"香奈儿\", \"icon\" : \"coat/yifu_024\", \"type\" : 2, \"level\" : 24, \"star\":240, \"exp\" : 180, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 125, \"name\" : \"白天鹅\", \"brand\":\"爱马仕\", \"icon\" : \"coat/yifu_025\", \"type\" : 2, \"level\" : 25, \"star\":250, \"exp\" : 190, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 126, \"name\" : \"梦回巴黎\", \"brand\":\"爱马仕\", \"icon\" : \"coat/yifu_026\", \"type\" : 2, \"level\" : 26, \"star\":260, \"exp\" : 200, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 127, \"name\" : \"淡若尘\", \"brand\":\"爱马仕\", \"icon\" : \"coat/yifu_027\", \"type\" : 2, \"level\" : 27, \"star\":270, \"exp\" : 220, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 128, \"name\" : \"墨染香\", \"brand\":\"爱马仕\", \"icon\" : \"coat/yifu_028\", \"type\" : 2, \"level\" : 28, \"star\":280, \"exp\" : 240, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 129, \"name\" : \"长安故事\", \"brand\":\"路易威登\", \"icon\" : \"coat/yifu_029\", \"type\" : 2, \"level\" : 29, \"star\":290, \"exp\" : 260, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 130, \"name\" : \"天国少女\", \"brand\":\"路易威登\", \"icon\" : \"coat/yifu_030\", \"type\" : 2, \"level\" : 30, \"star\":300, \"exp\" : 280, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 131, \"name\" : \"夜巴黎\", \"brand\":\"路易威登\", \"icon\" : \"coat/yifu_031\", \"type\" : 2, \"level\" : 31, \"star\":310, \"exp\" : 300, \"cost\" : \"[0,0,1]\"}",
		"{\"no\" : 132, \"name\" : \"星夜船\", \"brand\":\"路易威登\", \"icon\" : \"coat/yifu_032\", \"type\" : 2, \"level\" : 32, \"star\":320, \"exp\" : 330, \"cost\" : \"[0,0,1]\"}",

		"{\"no\" : 201, \"name\" : \"花蝴蝶\", \"brand\":\"范思哲\", \"icon\" : \"dress/kuzi_001\", \"type\" : 3, \"level\" : 1, \"star\":0, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 202, \"name\" : \"海螺裙\", \"brand\":\"范思哲\", \"icon\" : \"dress/kuzi_002\", \"type\" : 3, \"level\" : 2, \"star\":20, \"exp\" : 10, \"cost\" : \"[50]\"}",
		"{\"no\" : 203, \"name\" : \"牛仔热裤\", \"brand\":\"范思哲\", \"icon\" : \"dress/kuzi_003\", \"type\" : 3, \"level\" : 3, \"star\":30, \"exp\" : 20, \"cost\" : \"[100]\"}",
		"{\"no\" : 204, \"name\" : \"沙滩裤\", \"brand\":\"范思哲\", \"icon\" : \"dress/kuzi_004\", \"type\" : 3, \"level\" : 4, \"star\":40, \"exp\" : 30, \"cost\" : \"[200]\"}",
		"{\"no\" : 205, \"name\" : \"深蓝之吻\", \"brand\":\"范思哲\", \"icon\" : \"dress/kuzi_005\", \"type\" : 3, \"level\" : 5, \"star\":50, \"exp\" : 40, \"cost\" : \"[400]\"}",
		"{\"no\" : 206, \"name\" : \"橙黄青春\", \"brand\":\"迪奥\", \"icon\" : \"dress/kuzi_006\", \"type\" : 3, \"level\" : 6, \"star\":60, \"exp\" : 50, \"cost\" : \"[800]\"}",
		"{\"no\" : 207, \"name\" : \"少女漫舞\", \"brand\":\"迪奥\", \"icon\" : \"dress/kuzi_007\", \"type\" : 3, \"level\" : 7, \"star\":70, \"exp\" : 60, \"cost\" : \"[600,1]\"}",
		"{\"no\" : 208, \"name\" : \"粉色心情\", \"brand\":\"迪奥\", \"icon\" : \"dress/kuzi_008\", \"type\" : 3, \"level\" : 8, \"star\":80, \"exp\" : 70, \"cost\" : \"[200,3]\"}",
		"{\"no\" : 209, \"name\" : \"辰溪清\", \"brand\":\"迪奥\", \"icon\" : \"dress/kuzi_009\", \"type\" : 3, \"level\" : 9, \"star\":90, \"exp\" : 80, \"cost\" : \"[400,6]\"}",
		"{\"no\" : 210, \"name\" : \"北境玫瑰\", \"brand\":\"迪奥\", \"icon\" : \"dress/kuzi_010\", \"type\" : 3, \"level\" : 10, \"star\":100, \"exp\" : 90, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 211, \"name\" : \"性感少女\", \"brand\":\"普拉达\", \"icon\" : \"dress/kuzi_011\", \"type\" : 3, \"level\" : 11, \"star\":110, \"exp\" : 100, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 212, \"name\" : \"黑色曼陀罗\", \"brand\":\"普拉达\", \"icon\" : \"dress/kuzi_012\", \"type\" : 3, \"level\" : 12, \"star\":120, \"exp\" : 110, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 213, \"name\" : \"炽灼\", \"brand\":\"普拉达\", \"icon\" : \"dress/kuzi_013\", \"type\" : 3, \"level\" : 13, \"star\":130, \"exp\" : 120, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 214, \"name\" : \"灰色空间\", \"brand\":\"普拉达\", \"icon\" : \"dress/kuzi_014\", \"type\" : 3, \"level\" : 14, \"star\":140, \"exp\" : 130, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 215, \"name\" : \"蓝色翎羽\", \"brand\":\"普拉达\", \"icon\" : \"dress/kuzi_015\", \"type\" : 3, \"level\" : 15, \"star\":150, \"exp\" : 140, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 216, \"name\" : \"机车少女\", \"brand\":\"博柏利\", \"icon\" : \"dress/kuzi_016\", \"type\" : 3, \"level\" : 16, \"star\":160, \"exp\" : 150, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 217, \"name\" : \"夜太美\", \"brand\":\"博柏利\", \"icon\" : \"dress/kuzi_017\", \"type\" : 3, \"level\" : 17, \"star\":170, \"exp\" : 160, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 218, \"name\" : \"金粉世家\", \"brand\":\"博柏利\", \"icon\" : \"dress/kuzi_018\", \"type\" : 3, \"level\" : 18, \"star\":180, \"exp\" : 170, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 219, \"name\" : \"湛蓝颂\", \"brand\":\"博柏利\", \"icon\" : \"dress/kuzi_019\", \"type\" : 3, \"level\" : 19, \"star\":190, \"exp\" : 180, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 220, \"name\" : \"收腰格子裙\", \"brand\":\"博柏利\", \"icon\" : \"dress/kuzi_020\", \"type\" : 3, \"level\" : 20, \"star\":200, \"exp\" : 190, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 221, \"name\" : \"草绿金\", \"brand\":\"古驰\", \"icon\" : \"dress/kuzi_021\", \"type\" : 3, \"level\" : 21, \"star\":210, \"exp\" : 200, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 222, \"name\" : \"似水年华\", \"brand\":\"古驰\", \"icon\" : \"dress/kuzi_022\", \"type\" : 3, \"level\" : 22, \"star\":220, \"exp\" : 210, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 223, \"name\" : \"奶油白蝴蝶\", \"brand\":\"古驰\", \"icon\" : \"dress/kuzi_023\", \"type\" : 3, \"level\" : 23, \"star\":230, \"exp\" : 220, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 224, \"name\" : \"碎花裙\", \"brand\":\"古驰\", \"icon\" : \"dress/kuzi_024\", \"type\" : 3, \"level\" : 24, \"star\":240, \"exp\" : 230, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 225, \"name\" : \"天国少女\", \"brand\":\"古驰\", \"icon\" : \"dress/kuzi_025\", \"type\" : 3, \"level\" : 25, \"star\":250, \"exp\" : 240, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 226, \"name\" : \"白昼梦\", \"brand\":\"香奈儿\", \"icon\" : \"dress/kuzi_026\", \"type\" : 3, \"level\" : 26, \"star\":260, \"exp\" : 250, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 227, \"name\" : \"多彩少女\", \"brand\":\"香奈儿\", \"icon\" : \"dress/kuzi_027\", \"type\" : 3, \"level\" : 27, \"star\":270, \"exp\" : 260, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 228, \"name\" : \"素色水仙\", \"brand\":\"香奈儿\", \"icon\" : \"dress/kuzi_028\", \"type\" : 3, \"level\" : 28, \"star\":280, \"exp\" : 270, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 229, \"name\" : \"春花宴\", \"brand\":\"香奈儿\", \"icon\" : \"dress/kuzi_029\", \"type\" : 3, \"level\" : 29, \"star\":290, \"exp\" : 280, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 230, \"name\" : \"天青涩\", \"brand\":\"香奈儿\", \"icon\" : \"dress/kuzi_030\", \"type\" : 3, \"level\" : 30, \"star\":300, \"exp\" : 290, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 231, \"name\" : \"忆秋叶\", \"brand\":\"爱马仕\", \"icon\" : \"dress/kuzi_031\", \"type\" : 3, \"level\" : 31, \"star\":310, \"exp\" : 300, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 232, \"name\" : \"红尘恋\", \"brand\":\"爱马仕\", \"icon\" : \"dress/kuzi_032\", \"type\" : 3, \"level\" : 32, \"star\":320, \"exp\" : 310, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 233, \"name\" : \"破洞牛仔\", \"brand\":\"爱马仕\", \"icon\" : \"dress/kuzi_033\", \"type\" : 3, \"level\" : 33, \"star\":330, \"exp\" : 320, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 234, \"name\" : \"休闲长裤\", \"brand\":\"爱马仕\", \"icon\" : \"dress/kuzi_034\", \"type\" : 3, \"level\" : 34, \"star\":340, \"exp\" : 330, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 235, \"name\" : \"紫罗兰\", \"brand\":\"爱马仕\", \"icon\" : \"dress/kuzi_035\", \"type\" : 3, \"level\" : 35, \"star\":350, \"exp\" : 340, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 236, \"name\" : \"星光璀璨\", \"brand\":\"路易威登\", \"icon\" : \"dress/kuzi_036\", \"type\" : 3, \"level\" : 36, \"star\":360, \"exp\" : 350, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 237, \"name\" : \"繁星春水\", \"brand\":\"路易威登\", \"icon\" : \"dress/kuzi_037\", \"type\" : 3, \"level\" : 37, \"star\":370, \"exp\" : 360, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 238, \"name\" : \"千鹤子\", \"brand\":\"路易威登\", \"icon\" : \"dress/kuzi_038\", \"type\" : 3, \"level\" : 38, \"star\":380, \"exp\" : 370, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 239, \"name\" : \"蓝色生死恋\", \"brand\":\"路易威登\", \"icon\" : \"dress/kuzi_039\", \"type\" : 3, \"level\" : 39, \"star\":390, \"exp\" : 380, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 240, \"name\" : \"梦回凡尔赛\", \"brand\":\"路易威登\", \"icon\" : \"dress/kuzi_040\", \"type\" : 3, \"level\" : 40, \"star\":400, \"exp\" : 390, \"cost\" : \"[0,0,0,0,1]\"}",
		"{\"no\" : 241, \"name\" : \"天国嫁衣\", \"brand\":\"路易威登\", \"icon\" : \"dress/kuzi_041\", \"type\" : 3, \"level\" : 41, \"star\":410, \"exp\" : 400, \"cost\" : \"[0,0,0,0,1]\"}",

		"{\"no\" : 300, \"name\" : \"小光光\", \"icon\" : \"faxing/toufa_0\", \"type\" : 4, \"level\" : 1, \"star\":0, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 301, \"name\" : \"波波头\", \"icon\" : \"faxing/toufa_001\", \"type\" : 4, \"level\" : 2, \"star\":0, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 302, \"name\" : \"粉红初恋\", \"icon\" : \"faxing/toufa_002\", \"type\" : 4, \"level\" : 3, \"star\":20, \"exp\" : 5, \"cost\" : \"[5]\"}",
		"{\"no\" : 303, \"name\" : \"银灰光波\", \"icon\" : \"faxing/toufa_003\", \"type\" : 4, \"level\" : 4, \"star\":30, \"exp\" : 10, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 304, \"name\" : \"黄金蛋卷\", \"icon\" : \"faxing/toufa_004\", \"type\" : 4, \"level\" : 5, \"star\":40, \"exp\" : 15, \"cost\" : \"[5,1]\"}",
		"{\"no\" : 305, \"name\" : \"炫彩纹理\", \"icon\" : \"faxing/toufa_005\", \"type\" : 4, \"level\" : 6, \"star\":50, \"exp\" : 20, \"cost\" : \"[0,2]\"}",
		"{\"no\" : 306, \"name\" : \"空气刘海\", \"icon\" : \"faxing/toufa_006\", \"type\" : 4, \"level\" : 7, \"star\":60, \"exp\" : 25, \"cost\" : \"[5,2]\"}",
		"{\"no\" : 307, \"name\" : \"退潮\", \"icon\" : \"faxing/toufa_007\", \"type\" : 4, \"level\" : 8, \"star\":70, \"exp\" : 30, \"cost\" : \"[0,3]\"}",
		"{\"no\" : 308, \"name\" : \"梨花双辫\", \"icon\" : \"faxing/toufa_008\", \"type\" : 4, \"level\" : 9, \"star\":80, \"exp\" : 35, \"cost\" : \"[5,3]\"}",
		"{\"no\" : 309, \"name\" : \"白棉\", \"icon\" : \"faxing/toufa_009\", \"type\" : 4, \"level\" : 10, \"star\":90, \"exp\" : 40, \"cost\" : \"[0,4]\"}",
		"{\"no\" : 310, \"name\" : \"侠姬\", \"icon\" : \"faxing/toufa_010\", \"type\" : 4, \"level\" : 11, \"star\":100, \"exp\" : 45, \"cost\" : \"[5,4]\"}",
		"{\"no\" : 311, \"name\" : \"八爪浪\", \"icon\" : \"faxing/toufa_011\", \"type\" : 4, \"level\" : 12, \"star\":110, \"exp\" : 50, \"cost\" : \"[0,5]\"}",
		"{\"no\" : 312, \"name\" : \"艳阳天\", \"icon\" : \"faxing/toufa_012\", \"type\" : 4, \"level\" : 13, \"star\":120, \"exp\" : 55, \"cost\" : \"[5,5]\"}",
		"{\"no\" : 313, \"name\" : \"自然长发\", \"icon\" : \"faxing/toufa_013\", \"type\" : 4, \"level\" : 14, \"star\":130, \"exp\" : 60, \"cost\" : \"[0,6]\"}",
		"{\"no\" : 314, \"name\" : \"极光飞舞\", \"icon\" : \"faxing/toufa_014\", \"type\" : 4, \"level\" : 15, \"star\":140, \"exp\" : 65, \"cost\" : \"[5,6]\"}",
		"{\"no\" : 315, \"name\" : \"单马尾\", \"icon\" : \"faxing/toufa_015\", \"type\" : 4, \"level\" : 16, \"star\":160, \"exp\" : 70, \"cost\" : \"[0,7]\"}",
		"{\"no\" : 316, \"name\" : \"亚麻卷\", \"icon\" : \"faxing/toufa_016\", \"type\" : 4, \"level\" : 17, \"star\":160, \"exp\" : 75, \"cost\" : \"[5,7]\"}",
		"{\"no\" : 317, \"name\" : \"极光\", \"icon\" : \"faxing/toufa_017\", \"type\" : 4, \"level\" : 18, \"star\":170, \"exp\" : 80, \"cost\" : \"[0,8]\"}",
		"{\"no\" : 318, \"name\" : \"舒蕾\", \"icon\" : \"faxing/toufa_018\", \"type\" : 4, \"level\" : 19, \"star\":180, \"exp\" : 85, \"cost\" : \"[5,8]\"}",
		"{\"no\" : 319, \"name\" : \"鱼尾烫\", \"icon\" : \"faxing/toufa_019\", \"type\" : 4, \"level\" : 20, \"star\":190, \"exp\" : 90, \"cost\" : \"[0,9]\"}",

		"{\"no\" : 401, \"name\" : \"围巾1\", \"icon\" : \"player_scarf1\", \"type\" : 5, \"level\" : 1, \"star\":20, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 402, \"name\" : \"围巾2\", \"icon\" : \"player_scarf2\", \"type\" : 5, \"level\" : 2, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"no\" : 501, \"name\" : \"包包1\", \"icon\" : \"player_bag1\", \"type\" : 6, \"level\" : 1, \"star\":20, \"exp\" : 1, \"cost\" : \"[1]\"}",
		"{\"no\" : 502, \"name\" : \"包包2\", \"icon\" : \"player_bag2\", \"type\" : 6, \"level\" : 2, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

		"{\"no\" : 701, \"name\" : \"叮叮\", \"icon\" : \"pet/chongwu_01\", \"type\" : 7, \"level\" : 1, \"star\":10, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 702, \"name\" : \"米果\", \"icon\" : \"pet/chongwu_02\", \"type\" : 7, \"level\" : 2, \"star\":20, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 703, \"name\" : \"菲菲\", \"icon\" : \"pet/chongwu_03\", \"type\" : 7, \"level\" : 3, \"star\":30, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 704, \"name\" : \"沫沫\", \"icon\" : \"pet/chongwu_04\", \"type\" : 7, \"level\" : 4, \"star\":40, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 705, \"name\" : \"灰灰\", \"icon\" : \"pet/chongwu_05\", \"type\" : 7, \"level\" : 5, \"star\":50, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 706, \"name\" : \"嘟嘟\", \"icon\" : \"pet/chongwu_06\", \"type\" : 7, \"level\" : 6, \"star\":60, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 707, \"name\" : \"贝贝\", \"icon\" : \"pet/chongwu_07\", \"type\" : 7, \"level\" : 7, \"star\":70, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 708, \"name\" : \"奈奈\", \"icon\" : \"pet/chongwu_08\", \"type\" : 7, \"level\" : 8, \"star\":80, \"exp\" : 5, \"cost\" : \"[0,1]\"}",
		"{\"no\" : 709, \"name\" : \"小白\", \"icon\" : \"pet/chongwu_09\", \"type\" : 7, \"level\" : 9, \"star\":90, \"exp\" : 5, \"cost\" : \"[0,1]\"}",

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
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"BGM别闹！\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"说出来不怕丢人，如果能重来 ，我要当明星\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"范冰冰邀请你加入游戏\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"哦买噶！这是啥神仙发色，太酷了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"打卡签到，日给你温暖，我给你礼物\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"我弃坑了，但又回来了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"艾玛我老婆，全身名牌，亮瞎你们狗眼\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"巴拉巴拉小美女，变身\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"真的，脖子以下全是腿\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"你想把衣服放头上吧\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"这白裙的链接有吗 在线求\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"穿上粗布衣，拿起大哥大\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"原来GM是网络管理员啊！我还以为GM是肛门的缩写呢？ \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"我先吐口血，你们继续刷弹幕……\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"我们不一样不一样不一样\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"送你一份纯情小礼，有效期：一辈子\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"魔法小樱吗？！！！\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"有人吗，你的礼物到了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"这模特腿可以上天了 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"这头发是一个月没洗吗？\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"前方高能 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"哦哟 不错哦\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"尽管周围很多刺,但是有我在,你必能安心 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"隔着屏幕我都觉得刺激\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"前方高能预警，请非单身人员撤离 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"我做错了什么 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"恍恍惚惚哈哈哈，对不起我尽力了，没绷住。。。 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"买 买 买！！！\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"打卡签到 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"美女能换双鞋吗，这双不好看\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"玫瑰玫瑰！无论你是什么形态\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"哈哈哈哈，那件绿衣服笑死我\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"老板，这里的衣服我全包了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"你腰间的带子是我永远的信仰哈哈哈 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"大吉大利，晚上吃鸡…..\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"欢迎新玩家\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"本美已上线\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"都是人才啊，在下实在是佩服 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"哎，怎么装扮不对称 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"我头上有电话，可以随时打给罗志祥。。。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"这游戏真好玩。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"你认我做大哥，我教你玩游戏。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"这游戏策划是谁，出来。。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"那个丸子头，笑死本宝宝了。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"这游戏做的真好，老夫的少女心啊\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"画风好萌 ，好萌啊\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"做的不错，支持一下 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"姐姐，你忘记给我零用钱了 \"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"救命，前方浪大！！波涛汹涌\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"我有谢霆锋电话，你信吗\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"游戏名字好吸引人啊\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"竟然可以这么好玩，偶滴天\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"交个朋友不\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"装逼你妹！\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"这妹子由我亲自操刀\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"一首凉凉送给你们。。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"福建人民发来贺电\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"黑龙江人民发来贺电\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"55级，求超越。。。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"你的青春值几块钱！\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"怎么不合成眼镜嘴巴耳朵。。。\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"我放着客厅电视不看，跑来玩这游戏也是蛮拼的\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"蓝孩子也喜欢玩\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"熟悉的味道\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"开口脆,溜了溜了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"所有高级衣服穿上是什么的体验\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"说不丑的画画肯定不好\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"粉红色有点迷人\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"背景图是什么\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"画风清奇\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"人物挺好看的\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"我还在玩，老妈喊我吃饭了\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"经费不够，品质来凑\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"钻石数量引起强烈不满\"}",
		"{\"fromUid\" : \"\", \"toUid\" : \"all\", \"msg\" : \"一进来就送钻石，靠谱\"}",
	}

	for _, configStr := range configs {
		// log.Debug("Load ClothConfigs  db %v ", configStr)
		var s BarrageReport
		json.Unmarshal([]byte(configStr), &s)
		m.db.Create(&s)
	}
}
