package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"server/entry"
	"time"

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
	http.HandleFunc("/logout", LoginHandler)
	http.HandleFunc("/sign", SignHandler)
	http.HandleFunc("/getSign", GetSignHandler)
	http.HandleFunc("/cloth", ClothHandler)
	http.HandleFunc("/getCloth", GetClothHandler)
	http.HandleFunc("/userInfo", LoginHandler)
	http.HandleFunc("/getUserInfo", GetUserInfoHandler)

	// 监听绑定
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

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
		fmt.Println(bytes.NewBuffer(result).String())
		var str = bytes.NewBuffer(result).String()

		var s GetUserInfoRequest
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid)

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

		fmt.Println(bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
}

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
		fmt.Println(s.Uid)

		res := new(GetClothResponse)
		res.Code = 200
		res.Snap = "{\"aaa\" : \"asasa\"}"

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		fmt.Println(bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
}

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

		res := new(GetSignResponse)
		res.Code = 200
		res.Days = []bool{true, false, false, true}

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		fmt.Println(bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
}

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

		res := new(SignResponse)
		res.Code = 200
		res.Days = []bool{true, false, false, true}

		// 给客户端回复数据
		resBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println("生成json字符串错误")
		}

		fmt.Println(bytes.NewBuffer(resBytes).String())
		w.Write(resBytes)
	}
}

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
