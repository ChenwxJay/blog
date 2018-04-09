package ueditor

import (
	"net/http"
	"github.com/google/uuid"
	"strings"
	"path"
	"os"
	"io"
	"encoding/json"
	"fmt"
)

func uploadImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("upfile")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var _uuid, _ = uuid.NewUUID()
	var _uuid_string = _uuid.String()
	filename := strings.Replace(_uuid_string, "-", "", -1) + path.Ext(header.Filename)

	err = os.MkdirAll(path.Join("static", "upload"), 0775)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(path.Join("static", "upload", filename))
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	io.Copy(outFile, file)

	b, err := json.Marshal(map[string]string{
		"url":      fmt.Sprintf("/static/upload/%s", filename), //保存后的文件路径
		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
		"original": header.Filename,                            //原始文件名
		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(b))
	w.Write(b)
}
