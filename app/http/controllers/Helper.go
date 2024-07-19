package controllers

import "encoding/json"

func JsonToArray(jsondata []byte) map[string]interface{} {
	var err error
	var msgMapTemplate interface{}
	err = json.Unmarshal(jsondata, &msgMapTemplate)
	msgMap := msgMapTemplate.(map[string]interface{})
	if err != nil {
		panic(err)
	}

	return msgMap
}
