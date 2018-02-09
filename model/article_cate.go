package model

import (
	"strconv"
	"errors"
	"../db_helper"
)

type orderByNum int

const(
	OrderByNumAsc orderByNum = iota
	OrderByNumDesc
)

type ArticleCate struct {}

func ( self * ArticleCate ) GetEnabledArticleCates()  []map[string]string {
	var sql = `select id, name
					from article_category
					where id in (select cate_id
                          from article_categories
                          where article_id in (select id
                                                    from v_enabled_article))
          	   order by num  ASC`
	result := DbHelper.GetDataBase().Query(sql)
	return result
}

func ( self * ArticleCate ) GetCates( orderByNum orderByNum )  []map[string]string {
	var orderByString = ""
	if orderByNum == OrderByNumAsc {
		orderByString = " order by num asc"
	} else if orderByNum == OrderByNumDesc {
		orderByString = " order by num desc"
	}
	sql := "select id, name, num, add_date from article_category where pid = 0  " + orderByString
	result := DbHelper.GetDataBase().Query(sql)
	return result
}

func ( self * ArticleCate ) GetCateName( cateId int ) string {
	sql := "select name from article_category where id = " + strconv.Itoa( cateId )
	result := DbHelper.GetDataBase().GetSingle(sql)
	if result == "" {
		return "所有文章"
	}
	return result
}

func (self *ArticleCate) AddCate( name string, orderNum int ) error  {
	var countSql = "select count(*) from article_category where name = ?"
	var count,_ = strconv.Atoi( DbHelper.GetDataBase().GetSingle(countSql,name) )
	if count > 0 {
		return errors.New("类别名称已经存在")
	}
	var insertSql = "insert into  article_category(name,num,pid,add_date) values(?,?,0,CURRENT_TIMESTAMP())"
	DbHelper.GetDataBase().Query(insertSql,name,orderNum)
	return nil
}

func  (self *ArticleCate) EditCate( id int, name string, orderNum int )  {
	var sql = "update article_category set name = ? , num = ? where id = ?"
	DbHelper.GetDataBase().ExecuteSql(sql,name,orderNum,id)
}

func  (self *ArticleCate) DelCate( id int )  {
	var sql = "delete from article_category where id = ?"
	DbHelper.GetDataBase().Query(sql,id)
}

func (self *ArticleCate ) GetCate(id int) map[string]string  {
	var sql = "select * from article_category where id = ?"
	var result = DbHelper.GetDataBase().Query(sql,id)
	if len(result) > 0 {
		return result[0]
	} else {
		return nil
	}
}
