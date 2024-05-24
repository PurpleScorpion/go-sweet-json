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

func testJson() {
	js := jsonutil.NewJSONObject()
	js.FluentPut("name", "Jaina")
	js.FluentPut("age", 18)
	js.FluentPut("birthday", "2024-01-01 00:12:13")

	js2 := jsonutil.NewJSONObject()
	js2.FluentPut("testKey", 123)
	js2.FluentPut("phone", "110")
	js.FluentPut("testObject", js2)

	jr1 := jsonutil.NewJSONArray()
	jr1.FluentPut("胡桃")
	jr1.FluentPut("心海")
	jr1.FluentPut("椰羊")
	jr1.FluentPut("影")

	js.FluentPut("arr1", jr1)

	jsa1 := jsonutil.NewJSONObject()
	jsa1.FluentPut("name", "提莫")
	jsa1.FluentPut("age", 18)

	jsa2 := jsonutil.NewJSONObject()
	jsa2.FluentPut("name", "阿狸")
	jsa2.FluentPut("age", 19)

	jr2 := jsonutil.NewJSONArray()
	jr2.FluentPut(jsa1)
	jr2.FluentPut(jsa2)

	js.FluentPut("arr2", jr2)

	fmt.Println(js.ToJsonString())
}

func testJson2() {
	json := `{"age":18,"arr1":["胡桃","心海","椰羊","影"],"arr2":[{"age":18,"name":"提莫"},{"age":19,"name":"阿狸"}],"birthday":"2024-01-01 00:12:13","name":"Jaina","testObject":{"phone":"110","testKey":123}}`
	js := jsonutil.NewJSONObject()
	js.ParseObject(json)
	fmt.Println(js.ToJsonString())
}
