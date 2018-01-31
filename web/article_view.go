package web

import(
	"net/http"
	"../common"
	"../model"
	"strconv"
	"strings"
	"regexp"
)


type ArticleView struct{
	http.Handler
}

func ( self * ArticleView ) ServeHTTP( response http.ResponseWriter, request *http.Request ) {
	id, error := strconv.Atoi( request.FormValue("id") )
	if error != nil {
		response.Write([]byte("参数错误"))
		return 
	}
	var pageContentOut = make(chan string)
	go func() {
		var pageContent = common.GetFileContent("html/article_view.html")
		pageContentOut <- pageContent
		close(pageContentOut)
	}()
	var articleDataOut = make( chan map[string]string )
	go func() {
		var article = model.Article{}
		var data = article.GetArticle(id)
		articleDataOut <- data
		close(articleDataOut)
	}()

	var pageContent = <- pageContentOut
	var data = <- articleDataOut

	pageContent = strings.Replace(pageContent, "${title}",data["title"],-1)
	pageContent = strings.Replace(pageContent, "${content}",data["content"],-1)
	pageContent = code(pageContent)
	pageContent = setClientResourceVersion(pageContent)
	response.Write([]byte(pageContent))

	go func() {
		articleVisitInfoModel.Add(id, getIpAddr(request))
	}()
}

func code( content string ) string {
	var regex = regexp.MustCompile(`<pre[^>]*>([\s\S]+?)<\/pre>`)
	return regex.ReplaceAllStringFunc(content, func(str string) string {
		var replaceHtmlRegex = regexp.MustCompile(`<\/?(p|span)(>|[\s|\/][^>]*>)`)
		str = replaceHtmlRegex.ReplaceAllString(str,"")
		start := regexp.MustCompile(`<pre[^>]*>`)
		end := regexp.MustCompile(`<\/pre>`)
		str = start.ReplaceAllStringFunc(str, func(str string) string  {
			return str + "<code>"
		})
		str = end.ReplaceAllStringFunc(str, func(str string) string  {
			return "</code>" + str
		})
		return str
	})
}