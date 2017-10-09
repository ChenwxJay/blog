package admin

import (
	"net/http"
)


func Login( response http.ResponseWriter, request *http.Request ) {
	userName := request.FormValue("user_name");
	password := request.FormValue("password");
	success := loginModel.CallLogin( userName, password, response );
	var result []byte;
	if success {
		result = JsonData( 0, "登陆成功", nil );
	} else {
		result = JsonData( 1,"用户名或密码错误", nil );
	}
	response.Write(result);
}

func CheckLogin( response http.ResponseWriter, request *http.Request ) {
	var isLogin = loginModel.IsLogin( request )
	var result  []byte = nil
	if isLogin {
		result = JsonData( 0, "登录中", nil );
	} else {
		result = JsonData( 1,"登录失效", nil );
	}
	response.Write(result);
}

