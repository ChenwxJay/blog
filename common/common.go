package common

import (
	"strings"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
	"io"
	"encoding/base64"
	"crypto/rand"
	"regexp"
	"os"
	"io/ioutil"
	"../common/logger"
	"strconv"
	"time"
)

func IsNullOrEmpty( str string ) bool{
	str = strings.Trim(str," ")
	if len(str) == 0 {
		return true
	}
	return false
}

const (
	Code_Success = 0
	Code_Error = -1
)

type ResponseData struct {
	Code int
	Message string
	Data interface{}
}

func JsonData( code int , message string , data interface{} ) []byte {
	responeData := ResponseData{code, message, data}
	result , _ := json.Marshal(responeData)
	return result
}

func JsonpData( callback string, code int , message string , data interface{} ) []byte {
	responeData := ResponseData{code, message, data}
	result , _ := json.Marshal(responeData)
	javasScript := callback + "("+ string( result ) +");"
	return []byte( javasScript )
}

func NotData() []byte  {
	msg := "{\"code\": 0,\"message\": \"no data\"}"
	return []byte(msg)
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func PageSplice( data []map[string]string, page int, pageSize int )  []map[string]string{
	count := len(data)
	start := page * pageSize
	end := start + pageSize
	if start >= count {
		return []map[string]string{}
	}
	if end > count {
		end = count
	}
	return data[start:end]
}

func Trim( str string ) string {
	regex,_ := regexp.Compile("(^\\s*)|(\\s*$)")
	return regex.ReplaceAllString(str,"")
}

func GetFileContent( path string ) string  {
	file,err := os.Open(path)
	if err != nil {
		logger.OpenFileLogger().Println(err)
		logger.CloseFileLogger()
		return ""
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	result := string(data)
	return result
}

func IntArray2StringArray( intArray []int) []string  {
	var result = make([]string,0)
	for _, item := range intArray {
		result = append(result,strconv.Itoa(item))
	}
	return result
}

func String2DateTime( dateTimeString string ) time.Time  {
	var loc, _ = time.LoadLocation("Local")
	var time ,_ = time.ParseInLocation("2006-01-02 15:04:05", dateTimeString,loc)
	return time
}

func DateTime2String( dateTime time.Time) string  {
	return dateTime.Format("2006年01月02日 15:04")
}