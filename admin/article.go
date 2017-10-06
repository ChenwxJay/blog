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
	var author = request.FormValue("author")
	var content = request.FormValue("content")
	articleModel.AddArticle(title,author,content)
	response.Write(JsonData(0,"添加成功",nil))
}

func ArticleDel(response http.ResponseWriter, request *http.Request)   {
	var id = request.FormValue("id")
	var intId ,_ = strconv.Atoi(id)
	articleModel.DelArticle(intId)
	response.Write(JsonData(0,"删除成功",nil))
}

