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

	Name     string
	Star     int32
	Exp      int32
	LvChao   string
	Diamond  int32
	Level    int32
	Scene    int32
	Hair     int32
	Coat     int32
	Trouser  int32
	Neck     int32
	Shoe     int32
	MaxCloth int32

	OffLineLvChao string
}

// UserInfoRequest ..
type UserInfoRequest struct {
	Uid string

	Name    string
	Star    int32
	Exp     int32
	LvChao  string
	Diamond int32
	Level   int32
	Scene   int32
	Hair    int32
	Coat    int32
	Trouser int32
	Neck    int32
	Shoe    int32
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
	Order    int32
	Uid      string
	NickName string
	Star     int32
	HeadUrl  string
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
	for range time.Tick(time.Duration(10) * time.Second) {
		m := db.GetInstance()
		m.PersistentData()
	}
}

// Rank  排序
func Rank() {
	m := db.GetInstance()
	m.Rank()
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

		m := db.GetInstance()
		ranks := m.GetRank()

		res := new(RankResponse)
		res.Code = 200
		res.Players = make([]*RankInfo, 0)

		for i := 0; i < len(ranks); i++ {
			p := ranks[i]
			rankInfo := new(RankInfo)
			rankInfo.Order = int32(i + 1)
			rankInfo.NickName = p.Name
			rankInfo.HeadUrl = p.HeadUrl
			rankInfo.Star = p.Star
			res.Players = append(res.Players, rankInfo)
		}

		me := m.GetPlayer(s.Uid)
		// log.Debug("  asas  %v %v ", me.UserId, me.HeadUrl)
		meRankInfo := new(RankInfo)
		meRankInfo.Order = me.Order
		meRankInfo.NickName = me.Name
		meRankInfo.HeadUrl = me.HeadUrl
		meRankInfo.Star = me.Star
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
					p.Name = s.Name
					p.HeadUrl = s.HeadUrl
					p.Diamond = 100
					p.MaxCloth = 12
					p.LoginTime = time.Now()
					m.SavePlayer(p)
				} else {
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
		res.Level = player.Level
		res.Scene = player.Scene
		res.Hair = player.Hair
		res.Coat = player.Coat
		res.Trouser = player.Trouser
		res.Neck = player.Neck
		res.Shoe = player.Shoe
		res.MaxCloth = player.MaxCloth

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

// setCrossHeader 设置跨域访问
func setCrossHeader(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json
}
