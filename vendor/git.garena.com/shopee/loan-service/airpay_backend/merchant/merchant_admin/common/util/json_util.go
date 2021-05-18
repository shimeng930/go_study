package util

import (
	"encoding/json"
	"strings"
)

func SafeToJson(o interface{}) string {
	d, err := json.Marshal(o)
	if err != nil {
		return ""
	}
	return string(d)
}

func MapToJson(dict map[string]interface{}) string {
	jsonByte, _ := json.Marshal(dict)
	return string(jsonByte)
}

func JsonToMap(jsonStr string) map[string]interface{} {
	var m map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &m)
	return m
}

func JsonToMapStrict(jsonStr string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

func ShadowUpdateJSON(originalJson string, updateJson string) string {
	originalJson = strings.TrimSpace(originalJson)
	if originalJson == "" {
		return updateJson
	}
	updateJson = strings.TrimSpace(updateJson)
	if updateJson == "" {
		return originalJson
	}
	if originalJson == "" && updateJson == "" {
		return ""
	}
	originalMap := JsonToMap(originalJson)
	updateMap := JsonToMap(updateJson)
	for k, v := range updateMap {
		originalMap[k] = v
	}
	return MapToJson(originalMap)
}
