package web

import (
	"../model"
	"strings"
	"../common/config_manger"
)

var articleModel = model.Article{}
var articleCateModel = model.ArticleCate{}

const CLIENT_RESOURCE_VERSION = "20180130"

func setClientResourceVersion( pageContent string ) string {
	var version = config_manager.GetConfig().ClientResourceVersion
	if version == "" {
		version = "19871106"
	}
	return strings.Replace(pageContent,`${client_resource_version}`, version,-1)
}