package admin

import (
	"net/http"
	"../model"
)

type Login struct{
	http.Handler
}

func ( self * Login ) ServeHTTP( response http.ResponseWriter, request *http.Request ) {
	userName := request.FormValue("user_name");
	password := request.FormValue("password");
	loginModel := new (model.Login);
	success := loginModel.CallLogin( userName, password, response );
	var result []byte;
	if success {
		result = JsonData( 0, "登陆成功", nil );
	} else {
		result = JsonData( 1,"用户名或密码错误", nil );
	}
	response.Write(result);
}