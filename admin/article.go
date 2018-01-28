package admin

import (
	"net/http"
	"encoding/json"
	"strconv"
	"../common"
	"../model"
)

func CateItem(response http.ResponseWriter, request *http.Request) {
	var id = request.FormValue("id")
	var iid,_ = strconv.Atoi(id)
	var result = articleModel.GetCate(iid)
	json,_ := json.Marshal(result)
	response.Write(json)
}

func CateDel(response http.ResponseWriter, request *http.Request) {
	var id = request.FormValue("id")
	var iid ,_ = strconv.Atoi(id)
	articleModel.DelCate(iid)
	response.Write(JsonData(0,"删除成功",nil))
}

func CateEdit(response http.ResponseWriter, request *http.Request) {
	var name = request.FormValue("name")
	var num = request.FormValue("num")
	var id = request.FormValue("id")
	var iid ,_ = strconv.Atoi(id)
	var iNum,_ = strconv.Atoi(num)
	articleModel.EditCate(iid,name,iNum)
	response.Write(JsonData(0,"编辑成功",nil))
}

func CateAdd(response http.ResponseWriter, request *http.Request) {
	var name = request.FormValue("name")
	var num = request.FormValue("num")
	var iNum,_ = strconv.Atoi(num)
	var err = articleModel.AddCate(name, iNum)
	if err != nil {
		response.Write(JsonData(1,err.Error(),nil))
	} else {
		response.Write(JsonData(0,"添加成功",nil))
	}
}

func CateList(response http.ResponseWriter, request *http.Request) {
	var cateList = articleModel.GetCates(model.OrderByNumDesc)
	var json,_ = json.Marshal(cateList)
	response.Write(json)
}

func ArticleList(response http.ResponseWriter, request *http.Request) {
	articleList := articleModel.GetAll()
	json,_ := json.Marshal(articleList)
	response.Write(json)
}

func ArticleAdd(response http.ResponseWriter, request *http.Request) {
	if !loginModel.IsLogin(request) {
		response.Write(LoginErrorResponse())
		return
	}
	var title = request.FormValue("title")
	var cate = request.FormValue("cate")
	var author = request.FormValue("author")
	var content = request.FormValue("content")
	var articleId = articleModel.AddArticle(title,cate,author,content)
	go func() {
		common.LexemeCreate(articleId,common.ARTICLE_LEXEME_TYPE,title)
		common.LexemeCreate(articleId, common.ARTICLE_LEXEME_TYPE,content)
	}()
	response.Write(JsonData(0,"添加成功",nil))
}


func ArticleEdit(response http.ResponseWriter, request *http.Request) {
	if !loginModel.IsLogin(request) {
		response.Write(LoginErrorResponse())
		return
	}
	var title = request.FormValue("title")
	var cate = request.FormValue("cate")
	var author = request.FormValue("author")
	var content = request.FormValue("content")
	var id,_ =  strconv.Atoi( request.FormValue("id") )
	articleModel.EditArticle(id,title,cate,author,content)
	go func() {
		common.LexemeDelete(id,common.ARTICLE_LEXEME_TYPE)
		common.LexemeCreate(id,common.ARTICLE_LEXEME_TYPE,title)
		common.LexemeCreate(id,common.ARTICLE_LEXEME_TYPE,content)
	}()
	response.Write(JsonData(0,"编辑成功",nil))
}

func ArticleDel(response http.ResponseWriter, request *http.Request) {
	if !loginModel.IsLogin(request) {
		response.Write(LoginErrorResponse())
		return
	}
	var id = request.FormValue("id")
	var intId ,_ = strconv.Atoi(id)
	articleModel.DelArticle(intId)
	go func() {
		common.LexemeDelete(intId,common.ARTICLE_LEXEME_TYPE)
	}()
	response.Write(JsonData(0,"删除成功",nil))
}

func ArticleCates(response http.ResponseWriter, request *http.Request)   {
	var cates = articleModel.GetCates(model.OrderByNumAsc)
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