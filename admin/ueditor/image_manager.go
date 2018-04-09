package ueditor

import (
	"io/ioutil"
	"path"
	"sort"
	"net/http"
	"strconv"
	"encoding/json"
)

//排序
type ImageSortList []map[string]interface{}

func (self ImageSortList) Len() int  {
	return len(self)
}

func (self ImageSortList)  Less(i, j int) bool  {
	var iItem = self[i]
	var jItem = self[j]
	var iTime = iItem["mtime"].(int64)
	var jTime = jItem["mtime"].(int64)
	return iTime > jTime
}

func (self ImageSortList)  Swap(i, j int)   {
	self[i], self[j] = self[j], self[i]
}
//结束


func getAllImageList()  []map[string]interface{} {
	var picPathList = make([]map[string]interface{},0)
	//var picPathList = make(ImageSortList,0)
	var fileList, e = ioutil.ReadDir(path.Join("static", "upload"))
	if e != nil {
		return picPathList
	}
	for _, v := range fileList {
		v.ModTime()
		var picPath  = path.Join("/","static", "upload",v.Name())
		var mtime = v.ModTime().Unix()
		var item = make(map[string]interface{})
		item["url"] = picPath
		item["mtime"] = mtime
		picPathList = append(picPathList, item)
	}
	sort.Sort(ImageSortList(picPathList))
	return picPathList
}

func getLimit(r *http.Request) (startVal int,sizeVal int) {
	var start = r.FormValue("start")
	var size = r.FormValue("size")
	if start != "" {
		startVal,_ = strconv.Atoi(start)
	} else {
		startVal = 0
	}
	if size != "" {
		sizeVal,_ = strconv.Atoi(size)
	} else {
		sizeVal = 20
	}
	return startVal, sizeVal
}

func getImageList( r *http.Request ) (  reData []map[string]interface{}, reStart int, reTotal int ) {
	var picPathList = getAllImageList()
	var total = len(picPathList)
	var start, size = getLimit(r)
	var end =  start + size
	if end > total{
		end = total
	}
	if start >= end {
		return nil,0,0
	}
	picPathList = picPathList[start:end]
	return picPathList, start, total
}

func listImage( w http.ResponseWriter, r *http.Request )  {
	var picPathList,start, total = getImageList(r)

	var result = make(map[string]interface{})
	result["state"] = "SUCCESS"
	result["list"] = picPathList
	result["start"] = start
	result["total"] = total

	var jsonBytes,_ = json.Marshal(result)
	w.Write([]byte(jsonBytes))
}