package model

import (
	"../db_helper"
	"time"
	"../common"
)

type ArticleComment struct {}


func ( self * ArticleComment ) Add ( articleId int, content string, nick string, email string, ip string ) (success bool, message string) {
	var articeExistsOut = make(chan bool)
	go func() {
		var sql = "select id from v_enabled_article where id = ?"
		var result = DbHelper.GetDataBase().Query(sql,articleId)
		if len(result) == 0{
			articeExistsOut <- false
		} else {
			articeExistsOut <- true
		}
		close(articeExistsOut)
	}()

	var canCommentOut = make(chan bool)
	go func() {
		var sql = "select id,add_time from article_comment where from_ip = ? order by id desc"
		var result = DbHelper.GetDataBase().Query(sql,ip)
		if  len(result) >= 10{
			canCommentOut <- false
		} else {
			if len(result) == 0 {
				canCommentOut <- true
			} else {
				var item = result[0]
				var addTime = item["add_time"]
				var dAddTime = common.String2DateTime(addTime)
				var diff = time.Now().Sub(dAddTime)
				if diff.Seconds() < 60 {
					canCommentOut <- false
				}else {
					canCommentOut <- true
				}
			}
		}
		close(canCommentOut)
	}()

	var articleExists = <- articeExistsOut
	var canComment = <- canCommentOut
	if !articleExists  {
		return false,"文章不存在"
	}
	if !canComment {
		return false,"提交评论过于频繁或者评论次数达到上限"
	}
	var sql = "insert into article_comment(nick_name, content, from_ip, email, add_time, article_id) values(?,?,?,?,current_timestamp(),?)"
	DbHelper.GetDataBase().ExecuteSql(sql,nick,content,ip,email,articleId)
	return true ,"评论成功"
}

func ( self * ArticleComment ) List ( articleId int ) []map[string]string {
	var sql = "select * from article_comment where article_id = ? order by id desc"
	var result = DbHelper.GetDataBase().Query(sql,articleId)
	return result
}

func ( self * ArticleComment ) ListAll () []map[string]string {
	var sql = `select *,
					(select title from article where id = article_comment.article_id) as article_title
                     from article_comment order by id desc`
	var result = DbHelper.GetDataBase().Query(sql)
	return result
}

func ( self * ArticleComment ) Delete ( commentId int ){
	var sql = "delete from article_comment where id = ?"
	DbHelper.GetDataBase().Query(sql, commentId)
}