package common

import (
	"math"
	"strconv"
	"net/http"
	"strings"
	"regexp"
)

const INDEX_PAGE_SIZE = 30;
const NUMERIC_BUTTON_COUNT = 5;

func getPageIndex( request *http.Request ) int {
	pageIndex := request.FormValue("page");
	if pageIndex == "" {
		return 1;
	} else {
		if regexp.MustCompile("^\\d+$").MatchString(pageIndex) {
			val,_ := strconv.Atoi(pageIndex);
			return val;
		} else {
			return 1;
		}
	}
}

func CreatePager( request *http.Request, pageSize int,  numericButtonCount int ) Pager  {
	pageIndex := getPageIndex(request);
	return Pager{PageSize:pageSize, PageIndex:pageIndex, NumericButtonCount:numericButtonCount, Request:request};
}


type Pager struct {
	Request *http.Request;
	PageIndex int;
	PageSize int;
	DataCount int;
	NumericButtonCount int;
}

func (self *Pager) GetRange()( int , int )  {
	var start = self.PageSize * ( self.PageIndex - 1 );
	var end = start + self.PageSize ;
	return start, end;
}

func (self *Pager) SetDataCount( dataCount int ) {
	self.DataCount = dataCount;
}

func (self *Pager) GetPageCount() int  {
	if( self.DataCount == 0 ) {
		return 1;
	} else {
		return int( math.Ceil( float64( self.DataCount ) / float64( self.PageSize ) ) );
	}
}

func (self *Pager) getLink( page int ) string  {
	var url = self.Request.URL.String();
	var link = "";
	if strings.Index(url,"?") == -1 {
		link += url + "?page=" + strconv.Itoa(page);
	} else {
		reg := regexp.MustCompile("page=\\d*");
		if reg.MatchString(url) {
			link += reg.ReplaceAllString(url,"page=" + strconv.Itoa(page));
		} else {
			link +=  url + "&page=" + strconv.Itoa(page);
		}
	}
	return link;
}

func (self *Pager) ToHtml() string  {
	var start int = 0;
	var end int = 0;
	var pageHtml string = "";
	var interval int = int( math.Floor( float64(self.NumericButtonCount) / float64( 2 ) ) );
	if self.PageIndex <= interval  {
		start = 1;
		if self.NumericButtonCount < self.GetPageCount() {
			end = self.NumericButtonCount
		} else {
			end =  self.GetPageCount();
		}
	} else if self.PageIndex > self.GetPageCount() - interval  {
		if self.NumericButtonCount < self.GetPageCount() {
			start = self.GetPageCount() - self.NumericButtonCount + 1;
		} else {
			start = 1;
		}
		end = self.GetPageCount();
	} else {
		start = self.PageIndex - interval;
		if self.NumericButtonCount % 2 != 0 {
			end = self.PageIndex + interval;
		} else {
			end = self.PageIndex + interval - 1;
		}
	}
	if self.PageIndex == 1 {
		pageHtml += "<span>首页</span>";
		pageHtml += "<span>上页</span>";
	} else {
		pageHtml += "<a href='"+ self.getLink(1) +"'>首页</a>";
		pageHtml += "<a href='"+ self.getLink(self.PageIndex - 1) +"'>上页</a>";
	}
	for i := start; i <= end; i = i + 1 {
		if self.PageIndex == i {
			pageHtml += "<span class='current'>" +strconv.Itoa( i ) + "</span>";
		} else {
			pageHtml += "<a href='"+ self.getLink(i) +"'>" + strconv.Itoa( i ) + "</a>";
		}
	}
	if self.PageIndex == self.GetPageCount() {
		pageHtml += "<span>下页</span>";
		pageHtml += "<span>末页</span>"
	} else {
		pageHtml += "<a href='" + self.getLink(self.PageIndex + 1) + "'>下页</a>";
		pageHtml += "<a href='" + self.getLink(self.GetPageCount()) + "'>末页</a>";
	}
	return pageHtml;
}

