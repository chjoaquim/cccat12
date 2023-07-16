package commons

import "encoding/json"

func StructToJson(message interface{}) string {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
