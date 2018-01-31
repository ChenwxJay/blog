package model

import (
	"../db_helper"
)

type ArticleVisitInfo struct {}

func (self ArticleVisitInfo) Add( articleId int,  ip string)  {
	var sql = "insert visit_info(article_id, ip, time)  VALUES (?,?,current_timestamp())  ON DUPLICATE KEY UPDATE time = current_timestamp()"
	var db = DbHelper.GetDataBase()
	db.ExecuteSql(sql,articleId,ip)
}