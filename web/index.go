package web

import(
	"net/http"
	//"../db_helper"
	"../common"
	"strings"
	"strconv"
	"time"
)

var articleCates = make([]map[string]string,0)

func InitSyncArticleCates()  {
	articleCates = articleCateModel.GetEnabledArticleCates()
	go func() {
		ticker := time.NewTicker(time.Second * 60)
		for range ticker.C {
			articleCates = articleCateModel.GetEnabledArticleCates()
		}
	}()
}

type Index struct{
	http.Handler
}

func ( self * Index ) ServeHTTP( response http.ResponseWriter, request *http.Request ) {
	handleContent( response, request,"html/default.html" )
}

func handleContent(response http.ResponseWriter, request *http.Request , templateString string )  {
	pageContent := common.GetFileContent( templateString )
	pageContent = attachCateHtml( pageContent )
	pageContent = attachArticleList(pageContent,request)
	pageContent = setClientResourceVersion(pageContent)
	response.Write([]byte(pageContent))
}

func getQuery(  request *http.Request) (int, string)  {
	cateIdParam := request.FormValue("cate_id")
	kwParam := request.FormValue("kw")
	var cateId = 0
	if !common.IsNullOrEmpty(cateIdParam) {
		cateId,_ = strconv.Atoi(cateIdParam)
	}
	return cateId, kwParam
}

func attachArticleList( pageHtml string,   request *http.Request) string  {
	pager := common.CreatePager(request,common.INDEX_PAGE_SIZE,common.NUMERIC_BUTTON_COUNT)
	cateId, kw := getQuery(request)
	dataList, dataCount := articleModel.GetEnableArticleList( pager, cateId, kw)
	pager.SetDataCount(dataCount)
	var listHtml = ""
	for _, item := range dataList {
		var title = item["title"]
		listHtml += `<li>
				    <a class="big" href="article_view?id=`+ item["id"] +`">`+ title +`</a>
				</li>`
	}
	pageHtml = strings.Replace(pageHtml,"${article_list}",listHtml,-1)
	pageHtml = strings.Replace(pageHtml,"${page}",pager.ToHtml(),-1)
	pageHtml = strings.Replace(pageHtml, "${cate_name}",articleCateModel.GetCateName(cateId),-1)
	return pageHtml
}

func attachCateHtml( pageHtml string )  string {
	var html = ""
	for _, item := range articleCates {
		html += `<li>
               			<a link-id=`+ item["id"] +` href="index?cate_id=`+ item["id"] +`" class="small">`+ item["name"] +`</a>（`+ item["article_count"] +`）
            		</li>`
	}
	pageHtml = strings.Replace(pageHtml,"${cates}",html,-1)
	return pageHtml
}

