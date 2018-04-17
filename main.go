package main

import (
	"./web"
	"./admin"
	"./admin/ueditor"
	"./common/config_manager"
	"net/http"
	"./cache"
)

func init() {
	//检查配置文件
	config_manager.CheckConfig()
	//同步前缓存的类别数据
	cache.InitClientCache()
}


func main() {
	//前端
	http.HandleFunc("/put_url_to_baidu", web.PutUrlToBaidu )
	http.Handle("/log/", http.FileServer(http.Dir(".")))
	http.Handle("/html/", http.FileServer(http.Dir(".")))
	http.Handle("/upload/", http.FileServer(http.Dir(".")))
	http.Handle("/ueditor/", http.FileServer(http.Dir(".")))
	http.Handle("/", new(web.Index))
	http.Handle("/index", new(web.Index))
	http.Handle("/article_view", new(web.ArticleView))
	http.HandleFunc("/add_article_comment", web.AddArticleComment)
	http.HandleFunc("/list_article_comment", web.ListArticleComment)
	//手机端
	http.Handle("/mobile/", new(web.IndexForMobile))
	http.Handle("/mobile/index", new(web.IndexForMobile))
	http.Handle("/mobile/article_view", new(web.ArticleViewForMobile))

	//后端
	http.HandleFunc("/admin/login", admin.Login )
	http.HandleFunc("/admin/logout", admin.Logout )
	http.HandleFunc("/admin/check_login", admin.CheckLogin )

	http.HandleFunc("/admin/article_list", admin.AdminHandleFuncFactory(  admin.ArticleList )                                                            )
	http.HandleFunc("/admin/article_add",admin.AdminHandleFuncFactory(  admin.ArticleAdd ) )
	http.HandleFunc("/admin/article_del", admin.AdminHandleFuncFactory( admin.ArticleDel ) )
	http.HandleFunc("/admin/article_cates", admin.AdminHandleFuncFactory( admin.ArticleCates ) )
	http.HandleFunc("/admin/article_view", admin.AdminHandleFuncFactory( admin.ArticleView ) )
	http.HandleFunc("/admin/article_edit", admin.AdminHandleFuncFactory( admin.ArticleEdit ) )
	http.HandleFunc("/admin/article_disabled", admin.AdminHandleFuncFactory( admin.ArticleDisabled ) )
	http.HandleFunc("/admin/set_article_cate", admin.AdminHandleFuncFactory( admin.SetArticleCate ) )
	http.HandleFunc("/admin/cate_add",admin.AdminHandleFuncFactory(  admin.CateAdd ) )
	http.HandleFunc("/admin/cate_edit", admin.AdminHandleFuncFactory( admin.CateEdit ) )
	http.HandleFunc("/admin/cate_del",admin.AdminHandleFuncFactory(  admin.CateDel ) )
	http.HandleFunc("/admin/cate_item", admin.AdminHandleFuncFactory( admin.CateItem ) )
	http.HandleFunc("/admin/cate_list", admin.AdminHandleFuncFactory( admin.CateList ) )
	http.HandleFunc("/admin/comment_list",admin.AdminHandleFuncFactory(  admin.ListAllForArticleComment ) )
	http.HandleFunc("/admin/comment_del",admin.AdminHandleFuncFactory(  admin.DeleteArticleComment ) )


	http.HandleFunc("/get_ip_info", web.GetIpInfo)

	//UEditor上传相关
	http.HandleFunc("/ueditor/go", ueditor.UEditor)
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":" + config_manager.GetConfig().Port, nil)
}

