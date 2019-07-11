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

type LoginRequest struct {
	Uid  string
	Name string
}

type LoginResponse struct {
	Code   int
	Result bool
}

type LogoutRequest struct {
	Uid string
}

type LogoutResponse struct {
	Code   int
	Result bool
}

type GetUserInfoRequest struct {
	Uid string
}

type GetUserInfoResponse struct {
	Code int

	Name    string
	Star    int32
	LvChao  int32
	Diamond int32
	Level   int32
	Scene   int32
	Hair    int32
	Coat    int32
	Trouser int32
	Neck    int32
}

type UserInfoRequest struct {
	Uid string

	Name    string
	Star    int32
	LvChao  int32
	Diamond int32
	Level   int32
	Scene   int32
	Hair    int32
	Coat    int32
	Trouser int32
	Neck    int32
}

type UserInfoResponse struct {
	Code int
}

type GetClothRequest struct {
	Uid string
}

type GetClothResponse struct {
	Code int

	Snap string
}

type ClothRequest struct {
	Uid string

	Snap string
}

type ClothResponse struct {
	Code int
}

type GetSignRequest struct {
	Uid string
}

type GetSignResponse struct {
	Code int

	Days []bool
}

type SignRequest struct {
	Uid string
}

type SignResponse struct {
	Code int

	Days []bool
}

func main() {
	fmt.Println("start main")

	m := db.GetInstance()
	m.ConnectDB()
	m.CreateTables()

	m.LoadFromDB()

	// 注册函数，用户连接， 自动调用指定处理函数
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/sign", SignHandler)
	http.HandleFunc("/getSign", GetSignHandler)
	http.HandleFunc("/cloth", ClothHandler)
	http.HandleFunc("/getCloth", GetClothHandler)
	http.HandleFunc("/userInfo", UserInfoHandler)
	http.HandleFunc("/getUserInfo", GetUserInfoHandler)

	// 监听绑定
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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

		m.PersistentData()
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

		m.PersistentData()
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

		res := new(GetUserInfoResponse)
		res.Code = 200
		res.Name = player.Name
		res.Star = player.Star
		res.LvChao = player.LvChao
		res.Diamond = player.Diamond
		res.Level = player.Level
		res.Scene = player.Scene
		res.Hair = player.Hair
		res.Coat = player.Coat
		res.Trouser = player.Trouser
		res.Neck = player.Neck

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
		for i := days - 1; i >= 0; i-- {
			d := time.Now().AddDate(0, 0, -i).Format("2006/1/2")
			log.Debug("dddddd  %v ", d)
			if _, ok := signDic[d]; ok {
				bools[j] = true
			}
			j++
		}

		res := new(GetSignResponse)
		res.Code = 200
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
		m.PersistentData()

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
		m.PersistentData()

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
			if s.Star > 0 {
				player.Star = s.Star
			}
			if s.LvChao > 0 {
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
				player.LvChao = s.Hair
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
		}
		m.PersistentData()

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
