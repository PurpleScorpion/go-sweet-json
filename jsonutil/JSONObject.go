package jsonutil

import "encoding/json"

type JSONObject struct {
	jsonMap map[string]interface{}
}

func NewJSONObject() JSONObject {
	var jsonObj JSONObject
	jsonObj.jsonMap = make(map[string]interface{})
	return jsonObj
}

func (js *JSONObject) GetData() map[string]interface{} {
	res := make(map[string]interface{})
	return convertJSONObject2Map(res, js)
}

func (js *JSONObject) ParseObject(str string) *JSONObject {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		panic("An unexpected error occurred during the conversion process: " + err.Error())
	}
	js.jsonMap = data
	return js
}

func (js *JSONObject) ToJsonString() string {
	data := js.GetData()
	val, _ := json.Marshal(data)
	return string(val)
}

func (js *JSONObject) FluentPut(key string, value interface{}) *JSONObject {
	js.jsonMap[key] = value
	return js
}

func (js *JSONObject) HasKey(key string) bool {
	_, exists := js.jsonMap[key]
	return exists
}
