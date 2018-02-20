package web

import (
	"../model"
	"strings"
	"../common/config_manager"
	"net/http"
	"../common"
)

var articleModel = model.Article{}
var articleCateModel = model.ArticleCate{}
var articleVisitInfoModel = model.ArticleVisitInfo{}
var loginModel = model.Login{}
var articleCommentModel = model.ArticleComment{}

func setClientResourceVersion( pageContent string ) string {
	var version = config_manager.GetConfig().ClientResourceVersion
	if version == "" {
		version = "19871106"
	}
	return strings.Replace(pageContent,`${client_resource_version}`, version,-1)
}

func getIpAddr(  request *http.Request  ) string {
	var xForwardedFor = request.Header.Get("x-forwarded-for")
	var xRealIp = request.Header.Get("X-Real-IP")
	if xForwardedFor != "" {
		return xForwardedFor
	}
	if xRealIp != "" {
		return xRealIp
	}
	var ipAndPort = request.RemoteAddr
	var items = strings.Split(ipAndPort,":")
	if len(items) == 2 {
		return items[0]
	} else {
		return "localhost"
	}
}

func PutUrlToBaidu ( response http.ResponseWriter, request *http.Request)  {
	var articleList = articleModel.GetAllEnabled()
	var urlList = make([]string,0)
	for _, item := range articleList {
		var url = "https://www.chhblog.com/article_view?id=" + item["id"]
		urlList = append(urlList, url)
	}
	var bm = common.Baidu{}
	var result = bm.PutUrl(urlList)
	response.Write([]byte(result))
}