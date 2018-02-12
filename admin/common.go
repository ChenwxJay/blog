package admin

import (
	"../model"
	"../common"
	"net/http"
)

var articleModel = model.Article{}
var loginModel = model.Login{}
var articleCateModel = model.ArticleCate{}
var articleCommentModel = model.ArticleComment{}


func JsonData( code int , message string , data interface{} ) []byte {
	return common.JsonData(code,message,data)
}

func LoginErrorResponse() []byte {
	return  JsonData( 0, "登录失效", nil )
}

func AdminHandleFuncFactory( handleFunc func(response http.ResponseWriter, request *http.Request) )  func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		if !loginModel.IsLogin(request) {
			response.Write(LoginErrorResponse())
			return
		}
		handleFunc(response,request)
	}
}