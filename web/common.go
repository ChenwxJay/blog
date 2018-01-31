package web

import (
	"../model"
	"strings"
	"../common/config_manger"
	"net/http"
)

var articleModel = model.Article{}
var articleCateModel = model.ArticleCate{}
var articleVisitInfoModel = model.ArticleVisitInfo{}


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