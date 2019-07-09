package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"server/entry"

	"server/db"
)

type loginParam struct {
	Uid  string
	Name string
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type") //header的类型
	w.Header().Set("content-type", "application/json")                              //返回数据格式是json

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

		var s loginParam
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s.Uid + "  " + s.Name)

		m := db.GetInstance()
		p := new(entry.Player)
		p.UserId = s.Uid
		p.Name = s.Name
		m.SavePlayer(p)

		m.PersistentData()
	}

	// 给客户端回复数据
	io.WriteString(w, "{\"a\" : 11}")
	// w.Write([]byte("lisa"))
}

func main() {
	fmt.Println("start main")

	m := db.GetInstance()
	m.ConnectDB()
	m.CreateTables()

	m.LoadFromDB()

	// 注册函数，用户连接， 自动调用指定处理函数
	http.HandleFunc("/login", LoginHandler)

	// 监听绑定
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
