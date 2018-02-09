package web

import "net/http"
import (
	"../common"
	"strconv"
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
	articleCommentModel.Add(iArticleId,content,nick,email,fromIp)
	response.Write(common.JsonData(1,"评论成功",nil))
}