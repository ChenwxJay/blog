package admin

import (
	"net/http"
	"encoding/json"
	"strconv"
)

func ArticleList(response http.ResponseWriter, request *http.Request) {
	articleList := articleModel.GetAll();
	json,_ := json.Marshal(articleList);
	response.Write(json);
}

func ArticleAdd(response http.ResponseWriter, request *http.Request)   {
	var title = request.FormValue("title")
	var cate = request.FormValue("cate");
	var author = request.FormValue("author")
	var content = request.FormValue("content")
	articleModel.AddArticle(title,cate,author,content)
	response.Write(JsonData(0,"添加成功",nil))
}


func ArticleEdit(response http.ResponseWriter, request *http.Request)   {
	var title = request.FormValue("title")
	var cate = request.FormValue("cate");
	var author = request.FormValue("author")
	var content = request.FormValue("content")
	var id,_ =  strconv.Atoi( request.FormValue("id") )
	articleModel.EditArticle(id,title,cate,author,content)
	response.Write(JsonData(0,"编辑成功",nil))
}

func ArticleDel(response http.ResponseWriter, request *http.Request)   {
	var id = request.FormValue("id")
	var intId ,_ = strconv.Atoi(id)
	articleModel.DelArticle(intId)
	response.Write(JsonData(0,"删除成功",nil))
}

func ArticleCates(response http.ResponseWriter, request *http.Request)   {
	var cates = articleModel.GetCates()
	var json,_ = json.Marshal(cates)
	response.Write(json)
}


func ArticleView(response http.ResponseWriter, request *http.Request)   {
	var id,err = strconv.Atoi( request.FormValue("id") )
	if( err != nil ) {
		return
	}
	var articleInfo = articleModel.GetArticle(id)
	var json,_ = json.Marshal(articleInfo)
	response.Write(json)
}

