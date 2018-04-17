package ueditor

import (
	"path"
	"os"
	"io"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"strconv"
)

func newFileName() string  {
	var now = time.Now()
	var fileName = now.Format("200601021504")
	fileName = fileName + strconv.Itoa( now.Nanosecond() )
	return fileName
}

func upload(pathString string, response http.ResponseWriter, request *http.Request) {
	file, header, err := request.FormFile("upfile")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var filename = newFileName() +  path.Ext(header.Filename)

	err = os.MkdirAll(pathString, 0775)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(path.Join(pathString, filename))
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	io.Copy(outFile, file)

	b, err := json.Marshal(map[string]string{
		"url":      fmt.Sprintf("/"+ pathString +"/%s", filename), //保存后的文件路径
		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
		"original": header.Filename,                            //原始文件名
		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		panic(err)
	}
	response.Write(b)
}
