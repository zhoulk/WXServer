package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"server/config"
	"server/entry"
	"server/tool"
	"strings"
	"time"

	"github.com/name5566/leaf/log"

	"server/db"
)

// HeartRequest ..
type HeartRequest struct {
	Uid string
}

// HeartResponse ..
type HeartResponse struct {
	Code int
}

// LoginRequest ..
type LoginRequest struct {
	Uid     string
	Name    string
	Code    string
	HeadUrl string
}

// LoginResponse ..
type LoginResponse struct {
	Code   int
	Result bool
	Uid    string
}

// LogoutRequest ..
type LogoutRequest struct {
	Uid string
}

// LogoutResponse ..
type LogoutResponse struct {
	Code   int
	Result bool
}

// GetUserInfoRequest ..
type GetUserInfoRequest struct {
	Uid string
}

// GetUserInfoResponse ..
type GetUserInfoResponse struct {
	Code int

	Name       string
	Star       int32
	Exp        int32
	LvChao     string
	Diamond    int32
	Favour     int32
	Level      int32
	Scene      int32
	CP         int32
	Hair       int32
	Coat       int32
	Trouser    int32
	Neck       int32
	Shoe       int32
	MaxCloth   int32
	HeadUrl    string
	MaxCoat    int32
	MaxTrouser int32
	MaxShoe    int32

	OffLineLvChao string
}

// UserInfoRequest ..
type UserInfoRequest struct {
	Uid string

	Name       string
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
	MaxCoat    int32
	MaxTrouser int32
	MaxShoe    int32
}

// UserInfoResponse ..
type UserInfoResponse struct {
	Code int
}

// GetClothRequest ..
type GetClothRequest struct {
	Uid string
}

// GetClothResponse ..
type GetClothResponse struct {
	Code int

	Snap string
}

// ClothRequest ..
type ClothRequest struct {
	Uid string

	Snap string
}

// ClothResponse ..
type ClothResponse struct {
	Code int
}

// GetSignRequest ..
type GetSignRequest struct {
	Uid string
}

// GetSignResponse ..
type GetSignResponse struct {
	Code int

	StartDay string
	Days     []bool
}

// SignRequest ..
type SignRequest struct {
	Uid string
}

// SignResponse ..
type SignResponse struct {
	Code int

	Days []bool
}

// GetConfigRequest ..
type GetConfigRequest struct {
	Uid string

	Type int32 // 1 衣服   2 场景   3 咔位   4 签到  5 礼物
}

// GetConfigResponse ..
type GetConfigResponse struct {
	Code int

	Config string
}

// BuyRequest ..
type BuyRequest struct {
	Uid string

	Type int32 // 1 钻石买绿钞
	Num  int32
}

// BuyResponse ..
type BuyResponse struct {
	Code int

	Diamond  int32
	LvChao   string
	MaxCloth int32
}

// SellRequest ..
type SellRequest struct {
	Uid string

	Type   int32 // 1 卖衣服
	Params string
}

// SellResponse ..
type SellResponse struct {
	Code int

	LvChao string
}

// RankRequest ..
type RankRequest struct {
	Uid string

	Type int32 // 1 世界排行   2 好友排行   3 粉丝榜
}

// RankResponse ..
type RankResponse struct {
	Code int

	Players []*RankInfo
	Me      *RankInfo
}

// RankInfo ...
type RankInfo struct {
	Order     int32
	Uid       string
	NickName  string
	Star      int32
	HeadUrl   string
	HasFavour bool
	Favour    int32
	Constri   int32
}

// FavourRequest ...
type FavourRequest struct {
	Uid string

	ToUid string
}

// FavourResponse ...
type FavourResponse struct {
	Code int
}

// RewardRequest ...
type RewardRequest struct {
	Uid    string
	ToUid  string
	Msg    string
	GiftId int32
}

// RewardResponse ...
type RewardResponse struct {
	Code int
}

// GetBarrageRequest ...
type GetBarrageRequest struct {
	Uid string
}

// GetBarrageResponse ...
type GetBarrageResponse struct {
	Code int

	Barrages []*BarrageInfo
}

// ExtraMoneyRequest ...
type ExtraMoneyRequest struct {
	Uid     string
	LvChao  string
	Diamond int32
}

// ExtraMoneyResponse ...
type ExtraMoneyResponse struct {
	Code int
}

// OpenFromRequest ...
type OpenFromRequest struct {
	Uid     string
	FromUid string
	Type    int32
}

// OpenFromResponse ...
type OpenFromResponse struct {
	Code int
}

// GetPreUserRequest ...
type GetPreUserRequest struct {
	Uid string
}

// GetPreUserResponse ...
type GetPreUserResponse struct {
	Code int

	Name    string
	Star    int32
	Exp     int32
	HeadUrl string
}

// DailyGiftRequest ...
type DailyGiftRequest struct {
	Uid string
}

// DailyGiftResponse ...
type DailyGiftResponse struct {
	Code int

	Users []*UserInfo
	Flag  bool
}

type GainDailyGiftRequest struct {
	Uid string
}

type GainDailyGiftResponse struct {
	Code   int
	LvChao string
}

type UserInfo struct {
	Uid     string
	HeadUrl string
}

// BarrageInfo ...
type BarrageInfo struct {
	From string
	Msg  string
}

// WXCode2SessionResponse ...
type WXCode2SessionResponse struct {
	Openid      string //用户唯一标识
	Session_key string //会话密钥
	Unionid     string //用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
	Errcode     int32  //错误码
	Errmsg      string //错误信息
}

// WXGetTokenResponse ...
type WXGetTokenResponse struct {
	Access_token string //获取到的凭证
	Expires_in   int32  //凭证有效时间，单位：秒。目前是7200秒之内的值。
	Errcode      int32  //错误码
	Errmsg       string //错误信息
}

func main() {
	fmt.Println("start main")

	prices := make([]int64, 40)
	//power := int64(0)
	total := int64(0)
	for i := 0; i < 40; i++ {
		prices[i] = int64(50 * math.Pow(1.49, float64(i+1)))

		total += prices[i]
	}
	fmt.Println("price  ", total, prices)

	m := db.GetInstance()
	m.ConnectDB()
	m.CreateTables()
	m.InitializeConfigs()

	m.LoadFromDB()
	m.Rank()

	go Cal()
	go Persistent()
	go Rank()

	GetWXAccessToken()

	// 注册函数，用户连接， 自动调用指定处理函数
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/sign", SignHandler)
	http.HandleFunc("/getSign", GetSignHandler)
	http.HandleFunc("/cloth", ClothHandler)
	http.HandleFunc("/getCloth", GetClothHandler)
	http.HandleFunc("/userInfo", UserInfoHandler)
	http.HandleFunc("/getUserInfo", GetUserInfoHandler)
	http.HandleFunc("/heart", HeartHandler)
	http.HandleFunc("/rank", RankHandler)
	http.HandleFunc("/buy", BuyHandler)
	http.HandleFunc("/sell", SellHandler)
	http.HandleFunc("/getConfig", GetConfigHandler)
	http.HandleFunc("/favour", FavourHandler)
	http.HandleFunc("/reward", RewardHandler)
	http.HandleFunc("/getBarrage", GetBarrageHandler)
	http.HandleFunc("/extraMoney", ExtraMoneyHandler)
	http.HandleFunc("/minusMoney", MinusMonetHandler)
	http.HandleFunc("/openFrom", OpenFromHandler)
	http.HandleFunc("/getPreUserInfo", GetPreUserHandler)
	http.HandleFunc("/dailyGift", DailyGiftHandler)
	http.HandleFunc("/gainDailyGift", GainDailyGiftHandler)

	// 监听绑定
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Cal 计算
func Cal() {
	for range time.Tick(time.Duration(1) * time.Second) {
		m := db.GetInstance()
		m.Cal()
	}
}

// Persistent 固化
func Persistent() {
	for range time.Tick(time.Duration(60) * time.Second) {
		m := db.GetInstance()
		m.PersistentData()
	}
}

// Rank  排序
func Rank() {
	for range time.Tick(time.Duration(600) * time.Second) {
		m := db.GetInstance()
		m.Rank()
	}
}

// GetWXAccessToken 获取 微信token
func GetWXAccessToken() {
	log.Debug("GetWXAccessToken start ==================================== ")

	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + config.WXAPPId + "&secret=" + config.WXSecretKey
	res, err := http.Get(url)
	if err != nil {
		log.Debug("request err %v %v %v", url, res, err)
	}
	defer res.Body.Close()

	log.Debug("GetWXAccessToken request %v ", url)

	web, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debug("GetWXAccessToken request err %v ", err)
	} else {
		var result WXGetTokenResponse
		if err := json.Unmarshal(web, &result); err != nil {
			log.Debug("GetWXAccessToken request err %v", err)
		}
		log.Debug("GetWXAccessToken response %v ", result)
	}
	log.Debug("GetWXAccessToken end ==================================== ")
}

// BuyHandler  购买处理
func BuyHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("BuyHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	var s BuyRequest

	if err != nil {
		res := new(BuyResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		log.Debug("BuyHandler request %v", bytes.NewBuffer(result).String())

		var str = bytes.NewBuffer(result).String()
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		m.Buy(s.Uid, s.Type, s.Num)
		p := m.GetPlayer(s.Uid)

		res := new(BuyResponse)
		res.Code = 200
		res.MaxCloth = p.MaxCloth
		res.LvChao = p.LvChao
		res.Diamond = p.Diamond

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("BuyHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("BuyHandler end ===================================")
}

// GetConfigHandler  获取配置
func GetConfigHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("GetConfigHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	var s GetConfigRequest

	if err != nil {
		res := new(GetConfigResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		log.Debug("GetConfigHandler request %v", bytes.NewBuffer(result).String())

		var str = bytes.NewBuffer(result).String()
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		configStr := m.GetConfigStr(s.Type)

		res := new(GetConfigResponse)
		res.Code = 200
		res.Config = configStr

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("GetConfigHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("GetConfigHandler end ===================================")
}

// SellHandler  销售处理
func SellHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("SellHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	var s SellRequest

	if err != nil {
		res := new(SellResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		log.Debug("SellHandler request %v", bytes.NewBuffer(result).String())

		var str = bytes.NewBuffer(result).String()
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		m.Sell(s.Uid, s.Type, s.Params)
		p := m.GetPlayer(s.Uid)

		res := new(SellResponse)
		res.Code = 200
		res.LvChao = p.LvChao

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("SellHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("SellHandler end ===================================")
}

// RankHandler  排行榜
func RankHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("RankHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	var s RankRequest

	if err != nil {
		res := new(GetUserInfoResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		log.Debug("RankHandler request %v", bytes.NewBuffer(result).String())

		var str = bytes.NewBuffer(result).String()
		json.Unmarshal([]byte(str), &s)

		res := new(RankResponse)
		res.Code = 200
		res.Players = make([]*RankInfo, 0)

		m := db.GetInstance()
		if s.Type == 3 {
			ranks := m.GetFansRank(s.Uid)

			for i := 0; i < len(ranks); i++ {
				p := ranks[i]
				rankInfo := new(RankInfo)
				rankInfo.Uid = p.UserId
				rankInfo.Order = int32(i + 1)
				rankInfo.NickName = p.Name
				rankInfo.HeadUrl = p.HeadUrl
				rankInfo.Star = p.Star
				rankInfo.HasFavour = p.HasFavour
				rankInfo.Favour = p.Favour
				rankInfo.Constri = p.Constri
				res.Players = append(res.Players, rankInfo)
			}
		} else {
			ranks := m.GetRank(s.Uid)

			for i := 0; i < len(ranks); i++ {
				p := ranks[i]
				rankInfo := new(RankInfo)
				rankInfo.Uid = p.UserId
				rankInfo.Order = int32(i + 1)
				rankInfo.NickName = p.Name
				rankInfo.HeadUrl = p.HeadUrl
				rankInfo.Star = p.Star
				rankInfo.HasFavour = p.HasFavour
				rankInfo.Favour = p.Favour
				rankInfo.Constri = p.Constri
				res.Players = append(res.Players, rankInfo)
			}
		}

		me := m.GetPlayer(s.Uid)
		// log.Debug("  asas  %v %v ", me.UserId, me.HeadUrl)
		meRankInfo := new(RankInfo)
		meRankInfo.Order = me.Order
		meRankInfo.NickName = me.Name
		meRankInfo.HeadUrl = me.HeadUrl
		meRankInfo.Star = me.Star
		meRankInfo.Favour = me.Favour
		res.Me = meRankInfo

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("RankHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("RankHandler end ===================================")
}

// HeartHandler  心跳处理
func HeartHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("HeartHandler start ===================================")
	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("HeartHandler request %v", str)

		var s HeartRequest
		json.Unmarshal([]byte(str), &s)
		// fmt.Println(s.Uid)

		m := db.GetInstance()
		m.Heart(s.Uid)

		res := new(HeartResponse)
		res.Code = 200
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("HeartHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("HeartHandler end ===================================")
}

// LoginHandler  登录处理
func LoginHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("LoginHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
	} else {
		log.Debug("LoginHandler request %v", bytes.NewBuffer(result).String())
		// fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s LoginRequest
		json.Unmarshal([]byte(str), &s)
		// fmt.Println(s.Uid + "  " + s.Name + "  " + s.Code)

		url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + config.WXAPPId + "&secret=" + config.WXSecretKey + "&js_code=" + s.Code + "&grant_type=authorization_code"
		res, err := http.Get(url)
		if err != nil {
			log.Debug("LoginHandler  request err %v %v %v", url, res, err)
		}
		defer res.Body.Close()

		// log.Debug("LoginHandler  request %v %v %v", url, res, err)

		web, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Debug("LoginHandler  request code2session err %v ", err)
		} else {
			var result WXCode2SessionResponse
			if err := json.Unmarshal(web, &result); err != nil {
				log.Debug("LoginHandler  request code2session err %v %v %v", url, res, err)
			} else {
				log.Debug("LoginHandler  request code2session result %v ", result)

				m := db.GetInstance()
				p := m.FindPlayerByOpenID(result.Openid)
				if p == nil {
					p = new(entry.Player)
					if strings.HasPrefix(s.Uid, "test") {
						p.UserId = s.Uid
					} else {
						p.UserId = tool.UniqueId()
					}
					p.OpenId = result.Openid
					p.LvChao = "[0,10]"
					p.Name = s.Name
					p.HeadUrl = s.HeadUrl
					// p.Diamond = 100
					p.MaxCloth = 12
					p.LoginTime = time.Now()
					m.SavePlayer(p)
				} else {
					p.MaxCloth = 12
					p.LoginTime = time.Now()
					m.SavePlayer(p)
				}

				res := new(LoginResponse)
				res.Code = 200
				res.Result = true
				res.Uid = p.UserId
				// 给客户端回复数据
				resBytes, err := json.Marshal(res)
				if err != nil {
					fmt.Println("生成json字符串错误")
				}

				log.Debug("LoginHandler response %v", bytes.NewBuffer(resBytes).String())
				// io.WriteString(w, "{\"a\" : 11}")
				w.Write(resBytes)
			}
		}
	}
	log.Debug("LoginHandler end ===================================")
}

// LogoutHandler 登出处理
func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("LogoutHandler start ===================================")
	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("LogoutHandler request %v", str)

		var s LogoutRequest
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid)

		m := db.GetInstance()
		p := m.GetPlayer(s.Uid)
		p.LogoutTime = time.Now()
		m.SavePlayer(p)
		m.SaveSnap(p)

		res := new(LogoutResponse)
		res.Code = 200
		res.Result = true
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		log.Debug("LogoutHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("LogoutHandler end ===================================")
}

// GetUserInfoHandler 获取用户信息处理
func GetUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("GetUserInfoHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(GetUserInfoResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
		log.Debug("GetUserInfoHandler end ===================================")
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("GetUserInfoHandler request %v", str)

		var s GetUserInfoRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		player := m.GetPlayer(s.Uid)

		res := new(GetUserInfoResponse)
		res.Code = 200
		res.Name = player.Name
		res.Star = player.Star
		res.Exp = player.Exp
		if len(player.LvChao) > 0 {
			res.LvChao = player.LvChao
		} else {
			res.LvChao = "[]"
		}
		res.Diamond = player.Diamond
		res.Favour = player.Favour
		res.Level = player.Level
		res.Scene = player.Scene
		res.CP = player.CP
		res.Hair = player.Hair
		res.Coat = player.Coat
		res.Trouser = player.Trouser
		res.Neck = player.Neck
		res.Shoe = player.Shoe
		res.MaxCloth = player.MaxCloth
		res.HeadUrl = player.HeadUrl
		res.MaxCoat = player.MaxCoat
		res.MaxShoe = player.MaxShoe
		res.MaxTrouser = player.MaxTrouser

		res.OffLineLvChao = m.GetOffLineLvChao(s.Uid)

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("GetUserInfoHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
		log.Debug("GetUserInfoHandler end ===================================")
	}
}

// GetClothHandler 获取衣服合成快照
func GetClothHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("GetClothHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(GetClothResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("GetClothHandler request %v", str)

		var s GetClothRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		clothSnap := m.GetCloth(s.Uid)
		if len(clothSnap) == 0 {
			clothSnap = "[]"
		}

		res := new(GetClothResponse)
		res.Code = 200
		res.Snap = clothSnap

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("GetClothHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("GetClothHandler end ===================================")
}

// GetSignHandler 获取签到信息
func GetSignHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("GetSignHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(GetSignResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("GetSignHandler request %v", str)

		var s GetSignRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		signDic := m.GetSign(s.Uid)
		player := m.GetPlayer(s.Uid)
		createTime, _ := time.Parse("2006/1/2", player.CreateTime.Format("2006/1/2"))
		offset := time.Now().Sub(createTime)
		days := int(math.Ceil(offset.Hours() / 24))
		days = days % 7

		var bools = make([]bool, 7)
		j := 0
		for i := days; i >= 0; i-- {
			d := time.Now().AddDate(0, 0, -i).Format("2006/1/2")
			if _, ok := signDic[d]; ok {
				bools[j] = true
			}
			j++
		}

		res := new(GetSignResponse)
		res.Code = 200
		res.StartDay = time.Now().AddDate(0, 0, -days).Format("2006/1/2")
		res.Days = bools

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("GetSignHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("GetSignHandler end ===================================")
}

// SignHandler 签到处理
func SignHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("SignHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(SignResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("SignHandler request %v", str)

		var s SignRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		m.Sign(s.Uid)

		res := new(SignResponse)
		res.Code = 200

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("SignHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("SignHandler end ===================================")
}

// ClothHandler 衣服快照
func ClothHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("ClothHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(ClothResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("ClothHandler request %v", str)

		var s ClothRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		m.SaveCloth(s.Uid, s.Snap)

		res := new(ClothResponse)
		res.Code = 200

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("ClothHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("ClothHandler end ===================================")
}

// UserInfoHandler 用户信息
func UserInfoHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("UserInfoHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(UserInfoResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("UserInfoHandler request %v", str)

		var s UserInfoRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		player := new(entry.Player)
		player.UserId = s.Uid
		if len(s.Name) > 0 {
			player.Name = s.Name
		}
		if s.Star > 0 && player.Star <= s.Star {
			player.Star = s.Star
			player.Exp = s.Exp
		}
		if len(s.LvChao) > 0 {
			player.LvChao = s.LvChao
		}
		if s.Diamond > 0 {
			player.Diamond = s.Diamond
		}
		if s.Level > 0 {
			player.Level = s.Level
		}
		if s.Scene > 0 {
			player.Scene = s.Scene
		}
		if s.CP > 0 {
			player.CP = s.CP
		}
		if s.Hair > 0 {
			player.Hair = s.Hair
		}
		if s.Coat > 0 {
			player.Coat = s.Coat
		}
		if s.Trouser > 0 {
			player.Trouser = s.Trouser
		}
		if s.Neck > 0 {
			player.Neck = s.Neck
		}
		if s.Shoe > 0 {
			player.Shoe = s.Shoe
		}
		if s.Pet > 0 {
			player.Pet = s.Pet
		}
		if s.MaxCoat > 0 {
			player.MaxCoat = s.MaxCoat
		}
		if s.MaxShoe > 0 {
			player.MaxShoe = s.MaxShoe
		}
		if s.MaxTrouser > 0 {
			player.MaxTrouser = s.MaxTrouser
		}
		m.SavePlayer(player)

		res := new(UserInfoResponse)
		res.Code = 200

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("UserInfoHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("UserInfoHandler end ===================================")
}

// FavourHandler ...
func FavourHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("FavourHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(FavourResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("FavourHandler request %v", str)

		var s FavourRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		m.Favour(s.Uid, s.ToUid, 1)

		res := new(FavourResponse)
		res.Code = 200

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("FavourHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("FavourHandler end ===================================")
}

// RewardHandler ...
func RewardHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("RewardHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(RewardResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("RewardHandler request %v", str)

		var s RewardRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		flag, code := m.Reward(s.Uid, s.ToUid, s.Msg, s.GiftId)

		res := new(RewardResponse)
		if flag {
			res.Code = 200
		} else {
			if code == 888 {
				res.Code = 202
			} else {
				res.Code = 201
			}
		}

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("RewardHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("RewardHandler end ===================================")
}

// GetBarrageHandler ...
func GetBarrageHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("GetBarrageHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(GetBarrageResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("GetBarrageHandler request %v", str)

		var s GetBarrageRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		barrageReports := m.GetBarrage(s.Uid)

		barrages := make([]*BarrageInfo, 0)
		for _, report := range barrageReports {
			barrage := new(BarrageInfo)
			barrage.From = report.From
			barrage.Msg = report.Msg
			barrages = append(barrages, barrage)
		}

		res := new(GetBarrageResponse)
		res.Code = 200
		res.Barrages = barrages

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("GetBarrageHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("GetBarrageHandler end ===================================")
}

// ExtraMoneyHandler ...
func ExtraMoneyHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("ExtraMoneyHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(ExtraMoneyResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("ExtraMoneyHandler request %v", str)

		var s ExtraMoneyRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		m.ExtraMoney(s.Uid, s.LvChao, s.Diamond)

		res := new(ExtraMoneyResponse)
		res.Code = 200

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("ExtraMoneyHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("ExtraMoneyHandler end ===================================")
}

// MinusMonetHandler ...
func MinusMonetHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("MinusMonetHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(ExtraMoneyResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("MinusMonetHandler request %v", str)

		var s ExtraMoneyRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		m.MinusMoney(s.Uid, s.LvChao, s.Diamond)

		res := new(ExtraMoneyResponse)
		res.Code = 200

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("ExtraMoneyHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("MinusMonetHandler end ===================================")
}

// OpenFromHandler ...
func OpenFromHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("OpenFromHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(OpenFromResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("OpenFromHandler request %v", str)

		var s OpenFromRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		m.OpenFrom(s.Uid, s.FromUid, s.Type)

		res := new(OpenFromResponse)
		res.Code = 200

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("OpenFromHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("OpenFromHandler end ===================================")
}

// GetPreUserHandler 获取上一名玩家信息
func GetPreUserHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("GetPreUserHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(GetPreUserResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("GetPreUserHandler request %v", str)

		var s GetPreUserRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		prePlayer := m.GetPrePlayer(s.Uid)

		res := new(GetPreUserResponse)
		res.Code = 200
		res.HeadUrl = prePlayer.HeadUrl
		res.Name = prePlayer.Name
		res.Star = prePlayer.Star
		res.Exp = prePlayer.Exp

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("GetPreUserHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("GetPreUserHandler end ===================================")
}

// DailyGiftHandler ...
func DailyGiftHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("DailyGiftHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(DailyGiftResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("DailyGiftHandler request %v", str)

		var s DailyGiftRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		invitePlayers := m.GetInvitePlayers(s.Uid)
		flag := m.CheckGainDailyGift(s.Uid)

		res := new(DailyGiftResponse)
		res.Code = 200
		res.Users = make([]*UserInfo, 0)
		for _, p := range invitePlayers {
			u := new(UserInfo)
			u.Uid = p.UserId
			u.HeadUrl = p.HeadUrl
			res.Users = append(res.Users, u)
		}
		res.Flag = flag

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("DailyGiftHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("DailyGiftHandler end ===================================")
}

// GainDailyGiftHandler ...
func GainDailyGiftHandler(w http.ResponseWriter, req *http.Request) {
	log.Debug("GainDailyGiftHandler start ===================================")

	setCrossHeader(w, req)

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res := new(GainDailyGiftResponse)
		res.Code = 400
		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		w.Write(resBytes)
	} else {
		var str = bytes.NewBuffer(result).String()
		log.Debug("GainDailyGiftHandler request %v", str)

		var s GainDailyGiftRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		gift := m.GetDailyGift(s.Uid)

		res := new(GainDailyGiftResponse)
		res.Code = 200
		res.LvChao = gift

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		log.Debug("GainDailyGiftHandler response %v", bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
	log.Debug("GainDailyGiftHandler end ===================================")
}

// setCrossHeader 设置跨域访问
func setCrossHeader(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json
}
