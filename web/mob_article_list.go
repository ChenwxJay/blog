package web

import (
	"net/http"
	"../common"
	"strings"
)

//移动端
type IndexForMobile struct{
	http.Handler
}

func ( self * IndexForMobile ) ServeHTTP( response http.ResponseWriter, request *http.Request ) {
	var dataList = articleModel.GetAllEnabled()
	var pageContent = common.GetFileContent( "html/mob_article_list.html"  )
	var dataListString = data2String(dataList)
	pageContent = strings.Replace(pageContent,"${article_list}", dataListString, -1)
	response.Write([]byte(pageContent))
}

func splitDateString(  dateString string ) (date string, time string)  {
	var items = strings.Split(dateString," " )
	return items[0], items[1]
}

func data2String(dataList []map[string]string) string  {
	var htmlString = ""
	for _, item := range dataList {
		var id = item["id"]
		var title = item["title"]
		var addTime = item["add_time"]
		var date,_ = splitDateString(addTime)
		htmlString += `<li>
							<a href="/mobile/article_view?id=`+ id +`">`+ title +`</a>
							<span>`+ date +`</span>
                        </li>`
	}
	return htmlString
}