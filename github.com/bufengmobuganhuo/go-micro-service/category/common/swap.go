package common

import "encoding/json"

// SwapTo 通过json进行结构体赋值
func SwapTo(request, category interface{}) (err error) {
	// 先序列化成json格式
	dataByte, err := json.Marshal(request)
	if err != nil {
		return err
	}
	// 再反序列化
	return json.Unmarshal(dataByte, category)
}
