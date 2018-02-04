package model

import (
	"../common"
	"../db_helper"
	"strconv"
	"math"
	"sort"
	"strings"
)

type Article struct {}

func (self *Article) Disabled( articleId int ) {
	var sql = "update article set disabled = if(disabled = 1, 0, 1) where id = ?"
	DbHelper.GetDataBase().ExecuteSql(sql, articleId)
}

func (self *Article) AddArticle(title string, cate string, author string, content string) int {
	var sql = "insert into article (title, author, content, show_time, add_time, read_count, is_commend ) values (?,?,?,CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP(),0,0)"
	DbHelper.GetDataBase().ExecuteSql(sql,title,author,content)
	var identitySql = "select max(id) from article"
	var articleId = DbHelper.GetDataBase().GetSingle(identitySql)
	var iArticle,_ = strconv.Atoi(articleId)
	if cate == "" {
		return iArticle
	}
	go func() {
		var cateIds = strings.Split(cate,",")
		for _, cateId := range cateIds {
			var addCateSql = "insert into article_categories(article_id,cate_id) values (?,?)"
			DbHelper.GetDataBase().ExecuteSql(addCateSql,articleId,cateId)
		}
	}()
	return iArticle
}

func (self *Article) EditArticle(articleId int , title string, cate string, author string, content string)  {
	var sql = "update article set title = ?, author = ?, content = ? where id = ?"
	DbHelper.GetDataBase().ExecuteSql(sql,title,author,content,articleId)
	var delSql = "delete from article_categories where article_id = ?"
	DbHelper.GetDataBase().ExecuteSql(delSql,articleId)
	go func() {
		var cateIds = strings.Split(cate,",")
		for _, cateId := range cateIds {
			var addCateSql = "insert into article_categories(article_id,cate_id) values (?,?)"
			DbHelper.GetDataBase().ExecuteSql(addCateSql,articleId,cateId)
		}
	}()
}

func (self *Article) SetArticleCate(articleId int, cate string)  {
	var delSql = "delete from article_categories where article_id = ?"
	DbHelper.GetDataBase().ExecuteSql(delSql,articleId)
	go func() {
		var cateIds = strings.Split(cate,",")
		for _, cateId := range cateIds {
			var addCateSql = "insert into article_categories(article_id,cate_id) values (?,?)"
			DbHelper.GetDataBase().ExecuteSql(addCateSql,articleId,cateId)
		}
	}()
}

func (self *Article) DelArticle( id int)  {
	var sql = "delete from article where id = ?"
	DbHelper.GetDataBase().ExecuteSql(sql,id)
}

func (self *Article) GetEnableArticle( id int) (map[string]string) {
	return self.GetArticleDataItem(id,"v_enabled_article")
}

func (self *Article) GetArticle( id int) (map[string]string)  {
	return self.GetArticleDataItem(id,"article")
}

func (self *Article) GetArticleDataItem( id int, desc string ) (map[string]string)  {
	var db = DbHelper.GetDataBase()
	var articleOut = make(chan map[string]string)
	go func() {
		var sql = "select *, date(add_time) as add_date from "+ desc +" where id = " + strconv.Itoa(id)
		var result = db.Query(sql)
		if len(result) > 0 {
			articleOut <- result[0]
		}
		close(articleOut)
	}()

	var topCateIDsOut = make(chan string)
	go func() {
		var topCateIDs = self.getLevel0CateIDs(id)
		var result = "0"
		if len(topCateIDs) > 0  {
			result = strings.Join(common.IntArray2StringArray(topCateIDs),",")
		}
		topCateIDsOut <- result
		close(topCateIDsOut)
	}()

	var result = <- articleOut
	if len(result) > 0 {
		result["cate_id"] = <- topCateIDsOut
		return result
	} else {
		return nil
	}
}

func (self *Article) getLevel0CateId( cateId int ) int {
	var sql = "select id, pid from article_category where id =  " + strconv.Itoa( cateId )
	var result = DbHelper.GetDataBase().Query(sql)
	if len(result) == 0 {
		return cateId
	}
	var item = result[0]
	var id,_ = strconv.Atoi( item["id"] )
	var pid,_ = strconv.Atoi( item["pid"] )
	if pid == 0  {
		return id
	} else {
		return self.getLevel0CateId(pid)
	}
}

func (self *Article) getLevel0CateIDs( articleId int ) []int {
	var sql = "select cate_id from article_categories where article_id = " + strconv.Itoa( articleId )
	var cateIDs = DbHelper.GetDataBase().Query(sql)
	var topCateIDs = make([]int,0)
	for _, item := range cateIDs {
		var cateId,_ = strconv.Atoi( item["cate_id"] )
		topCateIDs = append(topCateIDs, self.getLevel0CateId(cateId))
	}
	return topCateIDs
}

func (self *Article) GetAllEnabled() []map[string]string  {
	var articleSql = "select id, title, author ,add_time, disabled from v_enabled_article order by id desc"
	return DbHelper.GetDataBase().Query(articleSql)
}

func (self *Article) GetAll(cateId string,  keyword string,disabled string) []map[string]string {
	var where = self.getWhere(cateId,keyword,disabled)
	var articleCateSql = `select a.article_id,
									group_concat(  b.name ) as cate_name,
									group_concat(b.id) as cate_id
							from article_categories as  a
							LEFT JOIN  article_category as b
							on a.cate_id = b.id
							GROUP BY article_id`
	var articleSql = `select id, title, author ,add_time, disabled, 
							(select count(1) from visit_info where article_id = article.id) as visit_count 
							from article where ` + where + ` order by id desc`
	var articleCateOut = make(chan []map[string]string)
	var articleListOut = make(chan []map[string]string)
	go func() {
		var cateList = DbHelper.GetDataBase().Query(articleCateSql)
		articleCateOut <- cateList
		close(articleCateOut)
	}()
	go func() {
		var articleList = DbHelper.GetDataBase().Query(articleSql)
		articleListOut <- articleList
		close(articleListOut)
	}()
	var articleCateList = <- articleCateOut
	var articleList = <- articleListOut
	for index, articleItem := range articleList {
		for _, cateItem := range articleCateList {
			if articleItem["id"] == cateItem["article_id"] {
				articleList[index]["cates"] = cateItem["cate_name"]
				articleList[index]["cate_ids"] = cateItem["cate_id"]
			}
		}
	}
	return articleList
}

func (self *Article) lexemeInCondition( articleIdInfo []common.LexemeDataItem  ) string {
	sort.Sort( common.LexemeDataItemSortList(articleIdInfo) )
	var ids = make([]string,0)
	for _, item := range articleIdInfo {
		ids = append(ids, strconv.Itoa( item.ID ) )
	}
	if len(ids) == 0 {
		return ""
	}
	return strings.Join(ids,",")
}

func (self *Article) sortResult( dataList []map[string]string, idListRefable string ) []map[string]string {
	if idListRefable == "" {
		return dataList
	}
	var result = make([]map[string]string,0)
	var ids = strings.Split(idListRefable,",")
	for _, id := range ids {
		for _, dataItem := range dataList {
			if dataItem["id"] == id {
				result = append(result, dataItem)
			}
		}
	}
	return result
}

func (self *Article) getWhere(cateIdString string, kw string, disabled string) string  {
	var where = " 1=1 "
	if cateIdString != "" {
		where += " and id in (select article_id from article_categories where cate_id in (select DISTINCT(id) from article_category where pid = "+ cateIdString +"  or id = "+ cateIdString +"))"
	}

	if kw != "" {
		var ids = ""
		var lexemeResult = common.LexemeFind( kw, common.ARTICLE_LEXEME_TYPE, 1, math.MaxInt32 )
		if lexemeResult.DataCount > 0  {
			ids = self.lexemeInCondition(lexemeResult.DataList)
			if ids != "" {
				where += " and id in (" + ids + ") "
			}
		} else {
			where += " and 1 = 2 "
		}
	}
	if disabled != "" {
		var _ , err = strconv.Atoi(disabled)
		if err == nil {
			where += " and disabled = " + disabled
		}
	}
	return where
}

func (self *Article) GetEnableArticleList( pager common.Pager, cateId int, kw string) ([]map[string]string, int)  {
	var start, _ = pager.GetRange()
	var pageSize = pager.GetPageSize()
	kw = common.Trim(kw)
	var limit = strconv.Itoa(start) + "," + strconv.Itoa(pageSize)
	//var where = " 1=1"
	var orderBy = " order by id desc "
	var ids = ""
	args := make([]interface{},0)
	var cateIdString = strconv.Itoa( cateId )
	var where = self.getWhere(cateIdString,kw,"")
	//if cateId != 0 {
	//	where += " and id in (select article_id from article_categories where cate_id in (select DISTINCT(id) from article_category where pid = "+ cateIdString +"  or id = "+ cateIdString +"))"
	//}
	//if kw != "" {
	//	var lexemeResult = common.LexemeFind( kw, common.ARTICLE_LEXEME_TYPE, 1, math.MaxInt32 )
	//	if lexemeResult.DataCount > 0  {
	//		ids = self.lexemeInCondition(lexemeResult.DataList)
	//		if ids != "" {
	//			where += " and id in (" + ids + ")"
	//			orderBy = ""
	//		}
	//	} else {
	//		where += " and 1 = 2 "
	//	}
	//}
	var countSql = "select count(1) from v_enabled_article where " + where
	var sql = "select id , title from v_enabled_article where "+ where +  orderBy + " limit " + limit
	dataCount,_:= strconv.Atoi( DbHelper.GetDataBase().GetSingle(countSql,args...) )
	dataList := DbHelper.GetDataBase().Query(sql,args...)
	dataList = self.sortResult(dataList, ids)
	return dataList,dataCount
}

