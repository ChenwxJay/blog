package model

import (
	"../common"
	"../db_helper"
	"net/http"
)

type Login struct {

}

func ( self *Login)  CallLogin( userName string, password string, response http.ResponseWriter)  bool {
	sql := "select id from user_info where username = ? and password = ? "
	result  := DbHelper.GetDataBase().Query(sql, userName, password)
	if len(result) == 0 {
		return false
	}
	userInfo := result[0]
	userId := userInfo["id"]
	self.doLogin(userId, userName, response)
	return true;
}

func (self *Login) IsLogin( request *http.Request  ) bool {
	userIdCookie,e1 := request.Cookie("user_id")
	userNameCookie,e2 := request.Cookie("user_name")
	userKeyCookie,e3 := request.Cookie("user_key")
	if e1 != nil || e2 != nil || e3 != nil {
		return  false
	}
	userKey := self.createUserKey( userIdCookie.Value, userNameCookie.Value )
	return userKeyCookie.Value == userKey
}

func (self *Login) doLogin( userId string, userName string,response http.ResponseWriter ) {
	http.SetCookie(response, self.createCookie("user_id", userId))
	http.SetCookie(response, self.createCookie("user_name", userName))
	http.SetCookie(response, self.createCookie("user_key", self.createUserKey(userId,userName)))
}

func (self *Login) createUserKey( userId string, userName string ) string {
	str := userId + userName + "kdjss141@2_1F-12";
	key := common.GetMd5String(str);
	return key;
}

func (self *Login) createCookie(name string, val string)  *http.Cookie{
	return &http.Cookie{
		Name:name,
		Value:val,
		Path:"/",
	}
}