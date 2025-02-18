package jsonutil

import (
	"encoding/json"
	"reflect"
)

func ParseEntity(jsonStr string, v any) error {
	err := json.Unmarshal([]byte(jsonStr), v)
	if err != nil {
		return err
	}
	return nil
}

func ToJsonEntityString(data any) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func convertJSONObject2Map(res map[string]interface{}, js JSONObject) map[string]interface{} {
	jsonMap := js.jsonMap
	if jsonMap == nil {
		return res
	}
	for key, value := range jsonMap {
		if value == nil {
			res[key] = nil
			continue
		}
		ifaceType := reflect.ValueOf(value).Type()
		if ifaceType == reflect.TypeOf(JSONObject{}) {
			jm := make(map[string]interface{})
			res[key] = convertJSONObject2Map(jm, value.(JSONObject))
		} else if ifaceType == reflect.TypeOf(JSONArray{}) {
			sliceOfMaps := make([]interface{}, 0)
			res[key] = convertJSONArray2Map(sliceOfMaps, value.(JSONArray))
		} else {
			res[key] = value
		}
	}
	return res
}

func convertJSONArray2Map(list []interface{}, array JSONArray) []interface{} {
	if array.Size() == 0 {
		return list
	}
	for i := 0; i < len(array.jsonArray); i++ {
		js := array.jsonArray[i]
		if js == nil {
			list = append(list, nil)
			continue
		}
		ifaceType := reflect.ValueOf(js).Type()
		if ifaceType == reflect.TypeOf(JSONObject{}) {
			jm := make(map[string]interface{})
			list = append(list, convertJSONObject2Map(jm, js.(JSONObject)))
		} else if ifaceType == reflect.TypeOf(JSONArray{}) {
			sliceOfMaps := make([]interface{}, 0)
			mp := convertJSONArray2Map(sliceOfMaps, js.(JSONArray))
			list = append(list, mp)
		} else {
			list = append(list, js)
		}
	}
	return list
}

func isSlice(obj interface{}) bool {
	value := reflect.ValueOf(obj)
	kind := value.Kind()

	return kind == reflect.Slice
}
func isStruct(obj interface{}) bool {
	value := reflect.ValueOf(obj)
	kind := value.Kind()

	return kind == reflect.Struct
}
