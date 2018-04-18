package cache

import (
	"time"
	"../model"
	"strconv"
    "github.com/go-redis/redis"
	"encoding/json"
	"../common/config_manager"
)

var modelArticle = model.Article{}
var modelArticleCate = model.ArticleCate{}

var client *redis.Client

const KEY_ARTICLE_CATES  =  "article_cates"
const KEY_ARTICLE  =  "article"

func InitClientCache()  {
	client = redis.NewClient(&redis.Options{
		Addr:    config_manager.GetConfig().Redis,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetArticleCates() []map[string]string {
	var cacheData, err = client.Get(KEY_ARTICLE_CATES).Result()
	if err != nil {
		var fromDatabase = modelArticleCate.GetEnabledArticleCates()
		var jsonBytes,_ = json.Marshal(fromDatabase)
		var jsonString = string(jsonBytes)
		client.Set(KEY_ARTICLE_CATES, jsonString,time.Minute * 10)
		return fromDatabase
	} else {
		var data = make([]map[string]string ,0)
		var err = json.Unmarshal([]byte(cacheData), &data)
		if err == nil {
			return data
		} else {
			return nil
		}
	}
}

func GetArticle( articleId int ) map[string]string  {
	var articleCache, err = client.Get( articleItemKey(articleId) ).Result()
	if err != nil {
		var articleInfo = modelArticle.GetArticle(articleId)
		var jsonBytes ,_ = json.Marshal(articleInfo)
		var jsonString = string(jsonBytes)
		client.Set(articleItemKey(articleId),jsonString, time.Minute * 10)
		return articleInfo
	} else {
		var data = make(map[string]string)
		var err = json.Unmarshal([]byte(articleCache),&data)
		if err == nil{
			return data
		} else {
			return nil
		}}
}

func UpdateArticle( articleId int )  {
	var key = articleItemKey(articleId)
	client.Expire(key,0)
}

func articleItemKey( articleId int) string {
	var sArticleId = strconv.Itoa( articleId )
	return KEY_ARTICLE +"-"+ sArticleId
}