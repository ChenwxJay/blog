package web

import "net/http"
import (
	"../common"
	"strconv"
	"encoding/json"
	"html"
)

func AddArticleComment( response http.ResponseWriter, request *http.Request )  {
	var articleId = request.FormValue("article_id")
	var content = request.FormValue("content")
	var email = request.FormValue("email")
	var nick = request.FormValue("nick")
	var fromIp = getIpAddr(request)
	if articleId == "" {
		response.Write(common.JsonData(1,"文章ID为空",nil))
		return
	}
	if content == "" {
		response.Write(common.JsonData(1,"文章内容为空",nil))
		return
	}
	var iArticleId,_ = strconv.Atoi(articleId)
	var success, msg = articleCommentModel.Add(iArticleId,content,nick,email,fromIp)
	var code = 1
	if success {
		code = 0
	}
	response.Write(common.JsonData(code,msg,nil))
}

func ListArticleComment( response http.ResponseWriter, request *http.Request )  {
	var articleId = request.FormValue("article_id")
	var iArticleId,err = strconv.Atoi(articleId)
	var result = make([]map[string]string,0)
	if err == nil {
		result = articleCommentModel.List(iArticleId)
		for _,item := range result {
			var addTime = common.String2DateTime( item["add_time"] )
			item["add_time"] = common.DateTime2String(addTime)
			item["content"] = html.EscapeString(item["content"])
		}
	}
	var listJson,_ = json.Marshal(result)
	response.Write(listJson)
}