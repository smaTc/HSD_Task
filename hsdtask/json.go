package hsdtask

import "encoding/json"

func ObjectToJson[T any](data T) ([]byte, error) {
	return json.Marshal(data)
}

func JsonToObject[T any](data []byte) (T, error) {
	var obj T
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return obj, err
	}
	return obj, nil
}

func GenerateJsonResponse(key, message string) []byte {
	return []byte("{\"" + key + "\":\"" + message + "\"}")
}
