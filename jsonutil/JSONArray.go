package jsonutil

import (
	"encoding/json"
)

type JSONArray struct {
	jsonArray []interface{}
}

func (jr *JSONArray) Size() int {
	if jr.jsonArray == nil {
		return 0
	}
	return len(jr.jsonArray)
}

func (jr *JSONArray) FluentPut(value interface{}) *JSONArray {
	jr.jsonArray = append(jr.jsonArray, value)
	return jr
}

func (jr *JSONArray) GetData() []interface{} {
	sliceOfMaps := make([]interface{}, 0)
	arr := NewJSONArray()
	arr.jsonArray = jr.jsonArray
	return convertJSONArray2Map(sliceOfMaps, arr)
}

func (jr *JSONArray) ToJsonString() string {
	data := jr.GetData()
	val, _ := json.Marshal(data)
	return string(val)
}

func NewJSONArray() JSONArray {
	var jr JSONArray
	jr.jsonArray = make([]interface{}, 0)
	return jr
}

func (jr *JSONArray) ParseJSONArray(jsonStr string) *JSONArray {
	var result []interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		panic("An unexpected error occurred during the conversion process: " + err.Error())
	}
	jr.jsonArray = result
	return jr
}
