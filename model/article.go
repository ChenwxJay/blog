package model

import (
	"../common"
	"../db_helper"
	"strconv"
	"fmt"
)

type Article struct {

}

func (self *Article) AddArticle(title string, cate string, author string, content string)  {
	var sql = "insert into article (title, author, content, show_time, add_time, read_count, is_commend ) values (?,?,?,CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP(),0,0)";
	DbHelper.GetDataBase().ExecuteSql(sql,title,author,content);
	var identitySql = "select @@identity"
	var articleId = DbHelper.GetDataBase().GetSingle(identitySql)

	var addCateSql = "insert into article_categories(article_id,cate_id) values ("+ articleId +","+ cate +")"
	DbHelper.GetDataBase().ExecuteSql(addCateSql)
}

func (self *Article) DelArticle( id int)  {
	var sql = "delete from article where id = ?"
	DbHelper.GetDataBase().ExecuteSql(sql,id)
}

func (self *Article) GetArticle( id int) (map[string]string)  {
	sql := "select * from article where id = " + strconv.Itoa(id);
	result := DbHelper.GetDataBase().Query(sql);
	if(  len(result) > 0 ) {
		return result[0];
	} else {
		return nil;
	}
}

func (self *Article) GetAll() []map[string]string {
	sql := "select id, title, author ,add_time from article order by id desc";
	return DbHelper.GetDataBase().Query(sql);
}

func (self *Article) GetArticleList( start int, end int, cateId int, kw string) ([]map[string]string, int)  {
	kw = common.Trim(kw);
	limit := strconv.Itoa(start) + "," + strconv.Itoa(end);
	where := " 1=1";
	args := make([]interface{},0);
	if( cateId != 0 ) {
		where += " and id in (select article_id from article_categories where cate_id in (select DISTINCT(id) from article_category where pid = "+ strconv.Itoa(cateId) +"))";
	}
	if( kw != "" ) {
		args = append(args, kw);
		where += " and instr(title,?)";
	}
	countSql := "select count(1) from article where " + where ;
	sql := "select id , title from article where "+ where +" order by id desc limit " + limit;
	fmt.Println(sql)
	dataCount,_:= strconv.Atoi( DbHelper.GetDataBase().GetSingle(countSql,args...) );
	dataList := DbHelper.GetDataBase().Query(sql,args...);
	return dataList,dataCount;
}


func ( self * Article ) GetCates()  []map[string]string {
	sql := "select id, name from article_category where pid = 0 order by num asc";
	result := DbHelper.GetDataBase().Query(sql);
	return result;
}

func ( self * Article ) GetCateName( cateId int ) string {
	sql := "select name from article_category where id = " + strconv.Itoa( cateId );
	result := DbHelper.GetDataBase().GetSingle(sql);
	if( result == "") {
		return "所有文章";
	}
	return result;
}
