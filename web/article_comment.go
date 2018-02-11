package web

import "net/http"
import (
	"../common"
	"strconv"
	"encoding/json"
)

func AddArticleComment( response http.ResponseWriter, request *http.Request )  {
	var articleId = request.FormValue("article_id")
	var content = request.FormValue("content")
	var email = request.FormValue("email")
	var nick = request.FormValue("nick")
	var fromIp = getIpAddr(request)
	if articleId == "" {
		response.Write(common.JsonData(0,"文章ID为空",nil))
		return
	}
	if content == "" {
		response.Write(common.JsonData(0,"文章内容为空",nil))
		return
	}
	var iArticleId,_ = strconv.Atoi(articleId)
	var success, msg = articleCommentModel.Add(iArticleId,content,nick,email,fromIp)
	var code = 0
	if success {
		code = 1
	}
	response.Write(common.JsonData(code,msg,nil))
}

func ListArticleComment( response http.ResponseWriter, request *http.Request )  {
	var articleId = request.FormValue("article_id")
	var iArticleId,err = strconv.Atoi(articleId)
	var result = make([]map[string]string,0)
	if err == nil {
		result = articleCommentModel.List(iArticleId)
	}
	var listJson,_ = json.Marshal(result)
	response.Write(listJson)
}