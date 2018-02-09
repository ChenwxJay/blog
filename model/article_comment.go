package model

import (
	"../db_helper"
)

type ArticleComment struct {}


func ( self * ArticleComment ) Add ( articleId int, content string, nick string, email string, ip string )  {
	var sql = "insert into article_comment(nick_name, content, from_ip, email, add_time, article_id) values(?,?,?,?,current_timestamp(),?)"
	DbHelper.GetDataBase().ExecuteSql(sql,nick,content,ip,email,articleId)
}