package main

import (
	"./web"
	"./admin"
	"net/http"
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
	http.ListenAndServe(":9529", nil);
}

