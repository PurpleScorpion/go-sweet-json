package jsonutil

import (
	"fmt"
	"time"
)

func (js *JSONObject) GetTime(key string) time.Time {
	value, exists := js.jsonMap[key]
	if !exists {
		return time.Time{}
	}
	switch t := value.(type) {
	case time.Time:
		val := value.(time.Time)
		return val
	default:
		panic(fmt.Sprintf("The value is of an unknown type: %T\n", t))
	}
}

func (js *JSONObject) GetString(key string) string {
	value, exists := js.jsonMap[key]
	if !exists {
		return ""
	}
	switch t := value.(type) {
	case string:
		val := value.(string)
		return val
	default:
		panic(fmt.Sprintf("The value is of an unknown type: %T\n", t))
	}
}

func (js *JSONObject) GetInt(key string) int {
	value, exists := js.jsonMap[key]
	if !exists {
		return 0
	}
	switch t := value.(type) {
	case int:
		val := value.(int)
		return val
	case int16:
		val := value.(int16)
		return int(val)
	case int32:
		val := value.(int32)
		return int(val)
	case int64:
		val := value.(int64)
		return int(val)
	default:
		panic(fmt.Sprintf("The value is of an unknown type: %T\n", t))
	}
}

func (js *JSONObject) GetInt16(key string) int16 {
	value, exists := js.jsonMap[key]
	if !exists {
		return 0
	}
	switch t := value.(type) {
	case int:
		val := value.(int)
		return int16(val)
	case int16:
		val := value.(int16)
		return int16(val)
	case int32:
		val := value.(int32)
		return int16(val)
	case int64:
		val := value.(int64)
		return int16(val)
	default:
		panic(fmt.Sprintf("The value is of an unknown type: %T\n", t))
	}
}

func (js *JSONObject) GetInt32(key string) int32 {
	value, exists := js.jsonMap[key]
	if !exists {
		return 0
	}
	switch t := value.(type) {
	case int:
		val := value.(int)
		return int32(val)
	case int16:
		val := value.(int16)
		return int32(val)
	case int32:
		val := value.(int32)
		return int32(val)
	case int64:
		val := value.(int64)
		return int32(val)
	default:
		panic(fmt.Sprintf("The value is of an unknown type: %T\n", t))
	}
}

func (js *JSONObject) GetInt64(key string) int64 {
	value, exists := js.jsonMap[key]
	if !exists {
		return 0
	}

	switch t := value.(type) {
	case int:
		val := value.(int)
		return int64(val)
	case int16:
		val := value.(int16)
		return int64(val)
	case int32:
		val := value.(int32)
		return int64(val)
	case int64:
		return value.(int64)
	default:
		panic(fmt.Sprintf("The value is of an unknown type: %T\n", t))
	}
}

func (js *JSONObject) GetFloat32(key string) float32 {
	value, exists := js.jsonMap[key]
	if !exists {
		return 0
	}
	switch t := value.(type) {
	case int:
		val := value.(int)
		return float32(val)
	case int16:
		val := value.(int16)
		return float32(val)
	case int32:
		val := value.(int32)
		return float32(val)
	case int64:
		val := value.(int64)
		return float32(val)
	case float32:
		return value.(float32)
	case float64:
		val := value.(float64)
		return float32(val)
	default:
		panic(fmt.Sprintf("The value is of an unknown type: %T\n", t))
	}
}

func (js *JSONObject) GetFloat64(key string) float64 {
	value, exists := js.jsonMap[key]
	if !exists {
		return 0
	}
	switch t := value.(type) {
	case int:
		val := value.(int)
		return float64(val)
	case int16:
		val := value.(int16)
		return float64(val)
	case int32:
		val := value.(int32)
		return float64(val)
	case int64:
		val := value.(int64)
		return float64(val)
	case float32:
		val := value.(float32)
		return float64(val)
	case float64:
		return value.(float64)
	default:
		panic(fmt.Sprintf("The value is of an unknown type: %T\n", t))
	}
}

func (js *JSONObject) GetBool(key string) bool {
	value, exists := js.jsonMap[key]
	if !exists {
		panic("The key is not exists")
	}
	switch t := value.(type) {
	case bool:
		val := value.(bool)
		return val
	default:
		panic(fmt.Sprintf("The value is of an unknown type: %T\n", t))
	}
}

func (js *JSONObject) GetJSONObject(key string) JSONObject {
	_, exists := js.jsonMap[key]
	if !exists {
		panic("The key is not exists")
	}
	obj := js.jsonMap[key]
	// 如果是切片类型
	if isSlice(obj) {
		// 抛出异常
		panic("Cannot convert object type to array type")
	}

	// 如果不是结构体类型
	if !isStruct(obj) {
		// 抛出异常
		panic("Cannot convert non object type to object type")
	}
	jsonObj := NewJSONObject()
	switch obj.(type) {
	// 如果是JSONObject类型
	case JSONObject:
		val := obj.(JSONObject)
		jsonObj.jsonMap = val.jsonMap
	// 如果是map类型
	case map[string]interface{}:
		val := obj.(map[string]interface{})
		jsonObj.jsonMap = val
	default:
		// 否则一般对象类型不能转换为JSONObject类型
		panic("Other object types cannot be converted to JSONObject type objects")
	}
	return jsonObj
}

func (js *JSONObject) GetJSONArray(key string) JSONArray {
	_, exists := js.jsonMap[key]
	if !exists {
		panic("The key is not exists")
	}
	obj := js.jsonMap[key]
	arr := NewJSONArray()
	//如果是结构体类型
	if isStruct(obj) {
		switch obj.(type) {
		// 如果该结构体类型是JSONArray类型
		case JSONArray:
			val := obj.(JSONArray)
			arr.jsonArray = val.jsonArray
			return arr
		}
		// 否则其他对象类型不能转换成切片类型
		panic("Cannot convert object type to array type")
	}
	// 如果不是切片类型
	if !isSlice(obj) {
		panic("Cannot convert non array types to array types")
	}
	// 否则强转成切片类型
	val := js.jsonMap[key].([]interface{})
	arr.jsonArray = val
	return arr
}

func (jr *JSONArray) GetJSONObject(index int32) JSONObject {
	if index < 0 || index >= jr.Size() {
		panic("The index is out of range")
	}
	obj := jr.jsonArray[index]
	if isSlice(obj) {
		panic("Cannot convert object type to array type")
	}

	if !isStruct(obj) {
		panic("Cannot convert non object type to object type")
	}
	jsonObj := NewJSONObject()
	switch obj.(type) {
	case JSONObject:
		val := obj.(JSONObject)
		jsonObj.jsonMap = val.jsonMap
	case map[string]interface{}:
		val := obj.(map[string]interface{})
		jsonObj.jsonMap = val
	default:
		panic("Other object types cannot be converted to JSONObject type objects")
	}
	return jsonObj
}

func (jr *JSONArray) GetJSONArray(index int32) JSONArray {
	if index < 0 || index >= jr.Size() {
		panic("The index is out of range")
	}
	obj := jr.jsonArray[index]
	arr := NewJSONArray()
	if isStruct(obj) {
		switch obj.(type) {
		case JSONArray:
			val := obj.(JSONArray)
			arr.jsonArray = val.jsonArray
			return arr
		}
		panic("Cannot convert object type to array type")
	}

	if !isSlice(obj) {
		panic("Cannot convert non array types to array types")
	}
	val := obj.([]interface{})
	arr.jsonArray = val
	return arr
}
