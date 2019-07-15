package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"server/entry"
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
	Uid  string
	Name string
}

// LoginResponse ..
type LoginResponse struct {
	Code   int
	Result bool
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

// RankRequest ..
type RankRequest struct {
	Uid string

	Type int32 // 1 世界排行   2 好友排行   3 粉丝榜
}

// RankResponse ..
type RankResponse struct {
	Code int

	Players []*RankInfo
}

// RankInfo ...
type RankInfo struct {
	Order    int32
	Uid      string
	NickName string
	Star     int32
}

func main() {
	fmt.Println("start main")

	m := db.GetInstance()
	m.ConnectDB()
	m.CreateTables()

	m.LoadFromDB()

	go Cal()
	go Persistent()

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

// RankHandler  排行榜
func RankHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息
	fmt.Println(req.Method)
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(req.URL)

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
	} else {
		fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s RankRequest
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid)

	}

	res := new(RankResponse)
	res.Code = 200
	res.Players = make([]*RankInfo, 0)

	for i := 0; i < 2; i++ {
		rankInfo := new(RankInfo)
		rankInfo.Order = int32(i + 1)
		rankInfo.NickName = fmt.Sprintf("%s%v", "asasas", i)
		rankInfo.Star = int32(200 - i)
		res.Players = append(res.Players, rankInfo)
	}

	// 给客户端回复数据
	resBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	fmt.Println(bytes.NewBuffer(resBytes).String())
	// io.WriteString(w, "{\"a\" : 11}")
	w.Write(resBytes)
}

// HeartHandler  心跳处理
func HeartHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息
	fmt.Println(req.Method)
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(req.URL)

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
	} else {
		fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s HeartRequest
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid)

		m := db.GetInstance()
		m.Heart(s.Uid)
	}

	res := new(HeartResponse)
	res.Code = 200
	// 给客户端回复数据
	resBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	fmt.Println(bytes.NewBuffer(resBytes).String())
	// io.WriteString(w, "{\"a\" : 11}")
	w.Write(resBytes)
}

// LoginHandler  登录处理
func LoginHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息
	fmt.Println(req.Method)
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(req.URL)

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
	} else {
		fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s LoginRequest
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid + "  " + s.Name)

		m := db.GetInstance()
		p := new(entry.Player)
		p.UserId = s.Uid
		p.Name = s.Name
		p.LoginTime = time.Now()
		m.SavePlayer(p)
	}

	res := new(LoginResponse)
	res.Code = 200
	res.Result = true
	// 给客户端回复数据
	resBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	fmt.Println(bytes.NewBuffer(resBytes).String())
	// io.WriteString(w, "{\"a\" : 11}")
	w.Write(resBytes)
}

// LogoutHandler 登出处理
func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息
	fmt.Println(req.Method)
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(req.URL)

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
	} else {
		fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s LogoutRequest
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid)

		m := db.GetInstance()
		p := m.GetPlayer(s.Uid)
		p.LogoutTime = time.Now()
		m.SavePlayer(p)
		m.SaveSnap(p)
	}

	res := new(LogoutResponse)
	res.Code = 200
	res.Result = true
	// 给客户端回复数据
	resBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	w.Write(resBytes)
}

// GetUserInfoHandler 获取用户信息处理
func GetUserInfoHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息
	fmt.Println(req.Method)
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(req.URL)

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
	} else {
		var str = bytes.NewBuffer(result).String()

		var s GetUserInfoRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		player := m.GetPlayer(s.Uid)

		log.Debug("GetUserInfoHandler  =====> %v   %v", s.Uid, player.Star)

		res := new(GetUserInfoResponse)
		res.Code = 200
		res.Name = player.Name
		res.Star = player.Star
		res.Exp = player.Exp
		res.LvChao = player.LvChao
		res.Diamond = player.Diamond
		res.Level = player.Level
		res.Scene = player.Scene
		res.Hair = player.Hair
		res.Coat = player.Coat
		res.Trouser = player.Trouser
		res.Neck = player.Neck
		res.Shoe = player.Shoe

		res.OffLineLvChao = m.GetOffLineLvChao(s.Uid)

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		w.Write(resBytes)
	}
}

// GetClothHandler 获取衣服合成快照
func GetClothHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息
	fmt.Println(req.Method)
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(req.URL)

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
		fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s GetClothRequest
		json.Unmarshal([]byte(str), &s)

		m := db.GetInstance()
		clothSnap := m.GetCloth(s.Uid)
		if len(clothSnap) == 0 {
			clothSnap = "[]"
		}

		// log.Debug("GetClothHandler %v", clothSnap)

		res := new(GetClothResponse)
		res.Code = 200
		res.Snap = clothSnap

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		w.Write(resBytes)
	}
}

// GetSignHandler 获取签到信息
func GetSignHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息
	fmt.Println(req.Method)
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(req.URL)

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
		fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s GetSignRequest
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid)

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
			log.Debug("dddddd  %v ", d)
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

		fmt.Println(bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
}

// SignHandler 签到处理
func SignHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息
	fmt.Println(req.Method)
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(req.URL)

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
		fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s SignRequest
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid)

		m := db.GetInstance()
		m.Sign(s.Uid)

		res := new(SignResponse)
		res.Code = 200

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		fmt.Println(bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
}

// ClothHandler 衣服快照
func ClothHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息
	fmt.Println(req.Method)
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(req.URL)

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
		fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s ClothRequest
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid + s.Snap)

		m := db.GetInstance()
		m.SaveCloth(s.Uid, s.Snap)

		res := new(ClothResponse)
		res.Code = 200

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		fmt.Println(bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
}

// UserInfoHandler 用户信息
func UserInfoHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

	if req.Method != "POST" {
		return
	}

	// 打印客户端头信息
	fmt.Println(req.Method)
	fmt.Println(req.Header)
	fmt.Println(req.Body)
	fmt.Println(req.URL)

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
		fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s UserInfoRequest
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid, "  ", s.Star, "  ", s.LvChao, "  ", len(s.Name), "  ", s.Diamond)

		m := db.GetInstance()
		player := m.GetPlayer(s.Uid)
		if player != nil {
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
		}

		res := new(UserInfoResponse)
		res.Code = 200

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		fmt.Println(bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
}
