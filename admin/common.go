package admin

import "encoding/json"

type ResponseData struct {
	Code int
	Message string
	Data interface{}
}

func JsonData( code int , message string , data interface{} ) []byte {
	responeData := ResponseData{code, message, data}
	result , _ := json.Marshal(responeData)
	return result
}
