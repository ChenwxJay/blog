package model

import (
	"../common"
	"../db_helper"
	"strconv"
)

type Article struct {

}

func (self *Article) AddArticle(title string, cate string, author string, content string)  {
	var sql = "insert into article (title, author, content, show_time, add_time, read_count, is_commend ) values (?,?,?,CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP(),0,0)";
	DbHelper.GetDataBase().ExecuteSql(sql,title,author,content);
	var identitySql = "select max(id) from article"
	var articleId = DbHelper.GetDataBase().GetSingle(identitySql)
	var addCateSql = "insert into article_categories(article_id,cate_id) values (?,?)"
	DbHelper.GetDataBase().ExecuteSql(addCateSql,articleId,cate)
}

func (self *Article) EditArticle(articleId int , title string, cate string, author string, content string)  {
	var sql = "update article set title = ?, author = ?, content = ? where id = ?";
	DbHelper.GetDataBase().ExecuteSql(sql,title,author,content,articleId);
	var delSql = "delete from article_categories where article_id = ?";
	DbHelper.GetDataBase().ExecuteSql(delSql,articleId)
	var addCateSql = "insert into article_categories(article_id,cate_id) values (?,?)"
	DbHelper.GetDataBase().ExecuteSql(addCateSql,articleId,cate)
}

func (self *Article) DelArticle( id int)  {
	var sql = "delete from article where id = ?"
	DbHelper.GetDataBase().ExecuteSql(sql,id)
}

func (self *Article) GetArticle( id int) (map[string]string)  {
	sql := "select * from article where id = " + strconv.Itoa(id);
	result := DbHelper.GetDataBase().Query(sql);
	if(  len(result) > 0 ) {
		var topCateIDs = self.getTopCateIDs(id)
		if len(topCateIDs) > 0  {
			result[0]["cate_id"] = strconv.Itoa( topCateIDs[0] )
		} else {
			result[0]["cate_id"] = "0"
		}
		return result[0];
	} else {
		return nil;
	}
}

func (self *Article) getTopCateId( cateId int ) int {
	var sql = "select id, pid from article_category where id =  " + strconv.Itoa( cateId )
	var result = DbHelper.GetDataBase().Query(sql)
	if( len(result) == 0 ) {
		return cateId
	}
	var item = result[0]
	var id,_ = strconv.Atoi( item["id"] )
	var pid,_ = strconv.Atoi( item["pid"] )
	if pid == 0  {
		return id
	} else {
		return self.getTopCateId(pid)
	}
}


func (self *Article) getTopCateIDs( articleId int ) []int {
	var sql = "select cate_id from article_categories where article_id = " + strconv.Itoa( articleId )
	var cateIDs = DbHelper.GetDataBase().Query(sql)
	var topCateIDs = make([]int,0)
	for _, item := range cateIDs {
		var cateId,_ = strconv.Atoi( item["cate_id"] )
		topCateIDs = append(topCateIDs, self.getTopCateId(cateId))
	}
	return topCateIDs
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
