
# go-sweet-json
# 一款简易的作用于go语言的JSON快速解析工具

## 用法详解
`该库的使用方式与java的fast-json很像`
### 引入库
`go get github.com/PurpleScorpion/go-sweet-json`
### 使用库
`import "github.com/PurpleScorpion/go-sweet-json/jsonutil"`

### 获取JSONObject/JSONArray对象
```text
   js := jsonutil.NewJSONObject()
   jr := jsonutil.NewJSONArray()
```
### 方法介绍

#### FluentPut

```text
   // JSONObject
   js.FluentPut("name","zhangsan")
   js.FluentPut("age",18)
   js.FluentPut("flag",true)
   
   // JSONArray
   jr.FluentPut("胡桃")
   jr.FluentPut("心海")
   jr.FluentPut("椰羊")
   
   // 组合使用
   组合使用方式请看demo包下的Demo.go文件
```

#### GetData

```text
  // 获取go原生的数据 , 可用于数据发送和数据返回等
  // JSONObject
  data := js.GetData()
  // JSONArray
  data := jr.GetData()
```

#### ToJsonString

```text
  // 获取JSON字符串
  // JSONObject
  jsonStr := js.ToJsonString()
  // JSONArray
  jsonStr := jr.ToJsonString()
  // 组合使用
  组合使用方式请看demo包下的Demo.go文件
```

#### ParseEntity
```text 
  // 解析JSON字符串为实体类或map
  jsonStr := `{"name":"Jaina","age":18,"birthday":"2024-01-01 00:12:13"}`
  myMap := make(map[string]interface{})
  jsonutil.ParseEntity(jsonStr, &myMap)
  fmt.Println(myMap)
  // 若想输出为可观测方式 , 请使用keqing工具类 https://github.com/PurpleScorpion/go-sweet-keqing
  // str := keqing.ToString(myMap)
  // fmt.Println(str)
```

#### ParseObject
```text 
  // 解析JSON字符串为JSONObject
  js := jsonutil.ParseObject(jsonStr)
  // 具体使用方式请看demo包下的Demo.go文件
```

#### ParseJSONArray
```text
  // 解析JSON字符串为JSONArray
  jr := jsonutil.ParseJSONArray(jsonStr)
  // 具体使用方式请看demo包下的Demo.go文件
```

#### Size
```text
  // 获取JSONArray的长度
  jr.Size()
  // 具体使用方式请看demo包下的Demo.go文件
```

#### HasKey
```text
  // 判断JSONObject中是否存在某个key
  js.HasKey("name")
```

#### GetJSONObject
```text
  // 获取JSONObject中的JSONObject
  js.GetJSONObject("name")
  
  // 获取JSONArray中的JSONObject
  index := 1
  jr.GetJSONObject(index)
```

#### GetJSONArray
```text
  // 获取JSONObject中的JSONArray
  js.GetJSONArray("name")

  // 获取JSONArray中的GetJSONArray
  index := 1
  jr.GetJSONArray(index)
```

#### 获取基本数据类型的数据
```text
JSONObject
    GetTime(key string)
    GetString(key string)
    GetInt(key string)
    GetInt16(key string)
    GetInt32(key string)
    GetInt64(key string)
    GetFloat32(key string)
    GetFloat64(key string)
    GetBool(key string)
注意 : 以上函数当获取不到key时会返回go的默认值


JSONArray
    GetTime(index int32)
    GetString(index int32)
    GetInt(index int32)
    GetInt16(index int32)
    GetInt32(index int32)
    GetInt64(index int32)
    GetFloat32(index int32)
    GetFloat64(index int32)
    GetBool(index int32)
注意 : 当索引大于JSONArray的长度时则抛出角标越界异常
    
    其中 
    在 Go 语言中，time.Time 类型的空值通常表示一个未设置或无效的时间值。它被称作零值（zero value）。当你声明一个 time.Time 类型的变量但没有初始化时，它的默认值就是这个空值。
    如果你想要检查一个 time.Time 是否为这个零值，你可以使用 IsZero() 方法，如下所示：
    if emptyTime.IsZero() {
        fmt.Println("The time is not set.")
    }
    
    因为bool类型的值空值为false , 所以当使用GetBool函数获取不到key时则抛出异常

    其他类型在go中的默认值
    int: 0
    float: 0
    string: ""
```