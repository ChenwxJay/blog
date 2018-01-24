package web

import(
	"net/http"
	//"../db_helper"
	"../common"
	"../model"
	"strings"
	"strconv"
)



type Index struct{
	http.Handler
}

func ( self * Index ) ServeHTTP( response http.ResponseWriter, request *http.Request ) {
	pageContent := common.GetFileContent("html/default.html");
	pageContent = attachCateHtml( pageContent );
	pageContent = attachArticleList(pageContent,request);
	response.Write([]byte(pageContent));
}

func getQuery(  request *http.Request) (int, string)  {
	cateIdParam := request.FormValue("cate_id");
	kwParam := request.FormValue("kw");
	var cateId = 0;
	if( !common.IsNullOrEmpty(cateIdParam)) {
		cateId,_ = strconv.Atoi(cateIdParam);
	}
	return cateId, kwParam;
}

func attachArticleList( pageHtml string,   request *http.Request) string  {
	pager := common.CreatePager(request,common.INDEX_PAGE_SIZE,common.NUMERIC_BUTTON_COUNT);
	cateId, kw := getQuery(request);
	dataList, dataCount := articleModel.GetArticleList( pager, cateId, kw);
	pager.SetDataCount(dataCount);
	var listHtml = "";
	for _, item := range dataList {
		var title = item["title"];
		listHtml += `<li>
				    <a class="big" href="article_view?id=`+ item["id"] +`">`+ title +`</a>
				</li>`;
	}
	pageHtml = strings.Replace(pageHtml,"${article_list}",listHtml,-1);
	pageHtml = strings.Replace(pageHtml,"${page}",pager.ToHtml(),-1);
	pageHtml = strings.Replace(pageHtml, "${cate_name}",articleModel.GetCateName(cateId),-1);
	return pageHtml;
}

func attachCateHtml( pageHtml string )  string {
	article := model.Article{};
	cates := article.GetCates();
	var html string = "";
	for _, item := range cates {
		html += `<li>
               			<a link-id=`+ item["id"] +` href="index?cate_id=`+ item["id"] +`" class="small">`+ item["name"] +`</a>
            		</li>`;
	}
	pageHtml = strings.Replace(pageHtml,"${cates}",html,-1);
	return pageHtml;
}
