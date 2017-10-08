package main

import (
	"./web"
	"./admin"
	"net/http"
	"./admin/ueditor"
)


func main() {
	//前端
	http.Handle("/log/", http.FileServer(http.Dir(".")))
	http.Handle("/html/", http.FileServer(http.Dir(".")))
	http.Handle("/upload/", http.FileServer(http.Dir(".")))
	http.Handle("/ueditor/", http.FileServer(http.Dir(".")))
	http.Handle("/", new(web.Index));
	http.Handle("/index", new(web.Index));
	http.Handle("/article_view", new(web.ArticleView));

	//后端
	http.Handle("/admin/login", new(admin.Login));
	http.HandleFunc("/admin/article_list", admin.ArticleList)
	http.HandleFunc("/admin/article_add", admin.ArticleAdd)
	http.HandleFunc("/admin/article_del", admin.ArticleDel)
	http.HandleFunc("/admin/article_cates", admin.ArticleCates)
	http.HandleFunc("/admin/article_view", admin.ArticleView)

	//UEditor上传图片
	http.HandleFunc("/ueditor/go/controller", ueditor.Controller)
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":9529", nil);
}

