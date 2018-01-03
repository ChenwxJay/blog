package web

import (
	"net/http"
	"fmt"
	"strconv"
	"io/ioutil"
	"errors"
)

func Interface2Error(data interface{}) error {
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

func httpGet(url string) (result string, resultErr error) {
	defer func() {
		err := recover()
		if err != nil {
			result = ""
			resultErr = Interface2Error(err)
		}
	}()
	client := &http.Client{};
	//client.Timeout = time.Second * 5
	request, err := http.NewRequest("GET", url, nil);
	if err != nil {
		return "", err;
	}
	response, err := client.Do(request);
	defer response.Body.Close();
	if err != nil {
		return "", err;
	}
	if response.StatusCode >= 400 || response.StatusCode < 200 {
		return "", fmt.Errorf("请求出错， 状态码：" + strconv.Itoa(response.StatusCode))
	}
	content, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return "", err;
	}
	return string(content), nil;
}

func GetIpInfo(response http.ResponseWriter, request *http.Request )  {
	var ip = request.FormValue("ip")
	var taobaoIpUrl = "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip
	var result,_ = httpGet(taobaoIpUrl)
	response.Header().Add("Access-Control-Allow-Origin","*")
	response.Write([]byte(result))
}