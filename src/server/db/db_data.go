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
		"{\"no\" : 1, \"name\" : \"场景1\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/1.jpg\", \"star\" : 10, \"level\" : 1}",
		"{\"no\" : 2, \"name\" : \"场景2\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/2.jpg\", \"star\" : 20, \"level\" : 2}",
		"{\"no\" : 3, \"name\" : \"场景3\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/3.jpg\", \"star\" : 30, \"level\" : 3}",
		"{\"no\" : 4, \"name\" : \"场景4\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/4.jpg\", \"star\" : 40, \"level\" : 4}",
		"{\"no\" : 5, \"name\" : \"场景5\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/5.jpg\", \"star\" : 50, \"level\" : 5}",
		"{\"no\" : 6, \"name\" : \"场景6\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/6.jpg\", \"star\" : 60, \"level\" : 6}",
		"{\"no\" : 7, \"name\" : \"场景7\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/7.jpg\", \"star\" : 70, \"level\" : 7}",
		"{\"no\" : 8, \"name\" : \"场景8\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/8.jpg\", \"star\" : 80, \"level\" : 8}",
		"{\"no\" : 9, \"name\" : \"场景9\", \"icon\" : \"https://mudgame.com.cn/wx/res/WXGame/101/scene/9.jpg\", \"star\" : 90, \"level\" : 9}",
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

		"{\"no\" : 601, \"name\" : \"留言板\", \"icon\" : \"gift0\", \"type\" : 10, \"diamond\":0, \"favour\" : 1, \"reward\" : 0}",
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
