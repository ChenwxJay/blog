var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
/// <reference path="jquery/jquery.d.ts" />
var chhblog;
(function (chhblog) {
    function queryString() {
        let urlItems = window.location.search.split("?");
        let values = new Map();
        if (urlItems.length == 0) {
            return values;
        }
        let params = urlItems[urlItems.length - 1];
        let result = params.split("&").reduce((rt, item) => {
            let [key, val] = item.split("=");
            rt.set(key, decodeURIComponent(val));
            return rt;
        }, values);
        return result;
    }
    chhblog.queryString = queryString;
    function ajaxAsync(url = "", data, method = "GET", dataType = "json") {
        return __awaiter(this, void 0, void 0, function* () {
            return new Promise((resolve, reject) => {
                let ajaxSetting = {};
                ajaxSetting.url = url;
                ajaxSetting.type = method;
                ajaxSetting.dataType = dataType;
                ajaxSetting.data = data;
                ajaxSetting.success = (response) => {
                    resolve(response);
                };
                ajaxSetting.error = (xhr) => {
                    reject("请求出错，请联系网站管理员");
                };
                $.ajax(ajaxSetting);
            });
        });
    }
    chhblog.ajaxAsync = ajaxAsync;
    function mapToObject(map) {
        let result = {};
        for (let [key, val] of map.entries()) {
            result[key] = val;
        }
        return result;
    }
    chhblog.mapToObject = mapToObject;
    function jquery2HtmlElements(jquery) {
        let result = [];
        jquery.each((index, item) => result.push(item));
        return result;
    }
    chhblog.jquery2HtmlElements = jquery2HtmlElements;
    function isInteger(val) {
        return /^\d+$/.test(val);
    }
    chhblog.isInteger = isInteger;
    function checkLogin() {
        if (window.location.href.indexOf("\/html\/admin\/") != -1) {
            $.get("/admin/check_login", function (response) {
                let data = eval("(" + response + ")");
                if (data.Code == 1) {
                    window.top.location.href = "/html/admin/login.html";
                }
            });
        }
    }
    chhblog.checkLogin = checkLogin;
})(chhblog || (chhblog = {}));
chhblog.checkLogin();
//# sourceMappingURL=common.js.map