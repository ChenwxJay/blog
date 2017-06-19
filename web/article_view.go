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
	pageContent := common.GetFileContent("html/article_view.html");
	id, error := strconv.Atoi( request.FormValue("id") )
	if( error != nil ) {
		response.Write([]byte("参数错误"));
		return ;
	}
	article := model.Article{};
	data := article.GetArticle(id);
	pageContent = strings.Replace(pageContent, "${title}",data["title"],-1);
	pageContent = strings.Replace(pageContent, "${content}",data["content"],-1);
	pageContent = code(pageContent);
	response.Write([]byte(pageContent));
}

func code( content string ) string {
	regex := regexp.MustCompile(`<pre[^>]*>([\s\S]+?)<\/pre>`)
	return regex.ReplaceAllStringFunc(content, func(str string) string {
		start := regexp.MustCompile(`<pre[^>]*>`);
		end := regexp.MustCompile(`<\/pre>`);
		str = start.ReplaceAllStringFunc(str, func(str string) string  {
			return str + "<code>";
		});
		str = end.ReplaceAllStringFunc(str, func(str string) string  {
			return "</code>" + str;
		});
		return str;
	});
}