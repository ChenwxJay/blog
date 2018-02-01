package common

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

type Baidu struct {

}

func (self *Baidu) httpPost(url string , content string) string {
	resp, err := http.Post(url,
		"text/plain",
		strings.NewReader(content))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	} else {
		return string(body)
	}
}

func (self *Baidu) PutUrl( urlList []string ) string  {
	var baiduServiceUrl = "http://data.zz.baidu.com/urls?site=https://www.chhblog.com&token=sXJx5ag4QVAveCMm"
	var urlString = strings.Join(urlList,"\n")
	var result = self.httpPost(baiduServiceUrl, urlString)
	return result
}

