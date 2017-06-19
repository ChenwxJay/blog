package admin

import (
	"../model"
	"net/http"
	"encoding/json"
)

func ArticleList(response http.ResponseWriter, request *http.Request) {
	articleModel := new (model.Article);
	articleList := articleModel.GetAll();
	json,_ := json.Marshal(articleList);
	response.Write(json);
}
