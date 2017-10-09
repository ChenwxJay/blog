package admin

import (
	"net/http"
	"fmt"
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

func Logout( response http.ResponseWriter, request *http.Request)  {
	fmt.Println(123)
	userIdCookie,e1 := request.Cookie("user_id")
	userNameCookie,e2 := request.Cookie("user_name")
	userKeyCookie,e3 := request.Cookie("user_key")

	if e1 == nil && e2 == nil && e3 == nil {
		userIdCookie.MaxAge = -1
		userIdCookie.Path = "/"

		userNameCookie.MaxAge = -1
		userNameCookie.Path = "/"

		userKeyCookie.MaxAge = -1
		userKeyCookie.Path = "/"

		http.SetCookie( response,userIdCookie )
		http.SetCookie( response,userNameCookie )
		http.SetCookie( response,userKeyCookie )
	}
	http.Redirect( response, request, "/html/admin/login.html", http.StatusFound )
}