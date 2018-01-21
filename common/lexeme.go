package common

import (
	"strings"
	"net/http"
	"fmt"
	"strconv"
	"io/ioutil"
	"errors"
	"regexp"
)

func interface2Error(data interface{}) error {
	resultError, ok := data.(error)
	if !ok {
		str, ok := data.(string)
		if !ok {
			return errors.New("未知错误")
		}
		return errors.New(str)
	}
	return resultError
}

func extractValidCharacters(input string) string {
	//过滤代码(pre标签中的内容)
	var preRegex = regexp.MustCompile(`<pre[^>]*>[\s\S]*?<\/pre>`)
	var output = preRegex.ReplaceAllString(input,"")
	//过滤html标签
	var  htmlTagRegex = regexp.MustCompile("<[^>]*>")
	output = htmlTagRegex.ReplaceAllString(output,"")
	//过滤html实体
	var entitiesRegex = regexp.MustCompile(`&[^;]+;`)
	output = entitiesRegex.ReplaceAllString(output,"")
	//过滤非语言文字字符
	var charRegex = regexp.MustCompile("[^\u4e00-\u9fa5a-zA-Z0-9]")
	output = charRegex.ReplaceAllString(output,"")
	return output
}

func map2HttpParams(data map[string]string) string {
	var result = make([]string,0)
	for key, val := range data {
		var item = key + "=" + extractValidCharacters(val)
		result  = append(result, item)
	}
	var str = strings.Join(result,"&")
	fmt.Println(str)
	return str
}

func HttpPost(url string, data map[string]string) (result string, resultErr error) {
	defer func() {
		err := recover()
		if err != nil {
			result = ""
			resultErr = interface2Error(err)
		}
	}()
	var params = map2HttpParams(data);
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err;
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("请求出错， 状态码：" + strconv.Itoa(resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil;
}

const LEXEME_SYS_KEY  =   "chhblog"

func LexemeCreate(dataId int, _type int, text string) bool {
	var data = make(map[string]string)
	data["text"] = text
	data["data_id"] = strconv.Itoa( dataId )
	data["type"] = strconv.Itoa( _type )
	data["sys_key"] = LEXEME_SYS_KEY
	HttpPost("https://www.chhblog.com", data)
	return true;
}