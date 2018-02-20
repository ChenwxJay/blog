package cache

import (
	"time"
	"../model"
	"strconv"
)

var modelArtile = model.Article{}
var modelArticleCate = model.ArticleCate{}


type ArticleCacheItem struct {
	Time time.Time
	Data map[string]string
}

var dataArticleCates = make([]map[string]string,0)
var dataArticleItems =  make(map[string]ArticleCacheItem)

func initArticleCates()  {
	dataArticleCates = modelArticleCate.GetEnabledArticleCates()
	go func() {
		ticker := time.NewTicker(time.Second * 60)
		for range ticker.C {
			dataArticleCates = modelArticleCate.GetEnabledArticleCates()
		}
	}()
}

func initCheckDataArticleItems()  {
	go func() {
		ticker := time.NewTicker(time.Second * 10)
		for range ticker.C {
			for sArticleId, item := range dataArticleItems{
				var diff = time.Now().Sub(item.Time).Seconds()
				if diff >= 60 {
					delete(dataArticleItems,sArticleId)
				}
			}
		}
	}()
}

func InitClientCache()  {
	initArticleCates()
	initCheckDataArticleItems()
}

func GetArticleCates() []map[string]string {
	return dataArticleCates
}

func GetArticle( articleId int ) map[string]string  {
	var sArticleId = strconv.Itoa( articleId )
	var result, found = dataArticleItems[sArticleId]
	if found {
		return result.Data
	}
	var articleItem = modelArtile.GetEnableArticle(articleId)

	var cacheItem = ArticleCacheItem{}
	cacheItem.Time = time.Now()
	cacheItem.Data = articleItem
	dataArticleItems[sArticleId] = cacheItem
	return articleItem
}