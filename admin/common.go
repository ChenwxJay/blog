package admin

import "encoding/json"
import (
	"../model"
)

var articleModel = model.Article{}
var loginModel = model.Login{}
var articleCateModel = model.ArticleCate{}

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

func LoginErrorResponse() []byte {
	return  JsonData( 0, "登录失效", nil )
}