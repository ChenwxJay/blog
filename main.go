package main

import (
	"./web"
	"./admin"
	"net/http"
	"./admin/ueditor"
	"./common/config_manger"
)


func main() {
	//检查配置文件
	config_manager.CheckConfig()

	//同步前缓存的类别数据
	web.InitSyncArticleCates()
	//前端
	http.HandleFunc("/put_url_to_baidu", web.PutUrlToBaidu )
	http.Handle("/log/", http.FileServer(http.Dir(".")))
	http.Handle("/html/", http.FileServer(http.Dir(".")))
	http.Handle("/upload/", http.FileServer(http.Dir(".")))
	http.Handle("/ueditor/", http.FileServer(http.Dir(".")))
	http.Handle("/", new(web.Index))
	http.Handle("/index", new(web.Index))
	http.Handle("/article_view", new(web.ArticleView))
	http.HandleFunc("/article_comment", web.AddArticleComment)
	//手机端
	http.Handle("/mobile/", new(web.IndexForMobile))
	http.Handle("/mobile/index", new(web.IndexForMobile))
	http.Handle("/mobile/article_view", new(web.ArticleViewForMobile))

	//后端
	http.HandleFunc("/admin/login", admin.Login )
	http.HandleFunc("/admin/logout", admin.Logout )
	http.HandleFunc("/admin/check_login", admin.CheckLogin )
	http.HandleFunc("/admin/article_list", admin.ArticleList)
	http.HandleFunc("/admin/article_add", admin.ArticleAdd)
	http.HandleFunc("/admin/article_del", admin.ArticleDel)
	http.HandleFunc("/admin/article_cates", admin.ArticleCates)
	http.HandleFunc("/admin/article_view", admin.ArticleView)
	http.HandleFunc("/admin/article_edit", admin.ArticleEdit)
	http.HandleFunc("/admin/article_disabled", admin.ArticleDisabled)
	http.HandleFunc("/admin/set_article_cate", admin.SetArticleCate)
	http.HandleFunc("/admin/cate_add", admin.CateAdd)
	http.HandleFunc("/admin/cate_edit", admin.CateEdit)
	http.HandleFunc("/admin/cate_del", admin.CateDel)
	http.HandleFunc("/admin/cate_item", admin.CateItem)
	http.HandleFunc("/admin/cate_list", admin.CateList)


	http.HandleFunc("/get_ip_info", web.GetIpInfo)

	//UEditor上传图片
	http.HandleFunc("/ueditor/go/controller", ueditor.Controller)
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":9529", nil)
}

