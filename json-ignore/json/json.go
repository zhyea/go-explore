package json

import "encoding/json"

//ToJson 转为json字符串
func ToJson(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
