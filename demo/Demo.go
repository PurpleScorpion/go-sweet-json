package demo

import (
	"fmt"
	"github.com/PurpleScorpion/go-sweet-json/jsonutil"
)

type Demo struct {
}

// 获取json对象设置值
func demo1() {
	js := jsonutil.NewJSONObject()
	js.FluentPut("name", "Jaina")
	fmt.Println(js.ToJsonString())
}

// 设置多个值
func demo2() {
	js := jsonutil.NewJSONObject()
	js.FluentPut("name", "Jaina")
	js.FluentPut("age", 18)
	js.FluentPut("birthday", "2024-01-01 00:12:13")
	fmt.Println(js.ToJsonString())
}
