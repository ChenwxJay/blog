/// <reference path="jquery/jquery.d.ts" />
namespace chhblog{
    export function queryString() : Map<string,string>{
        let urlItems = window.location.search.split("?");
        let values = new Map();
        if( urlItems.length == 0 ) {
            return values;
        }
        let params = urlItems[urlItems.length - 1];
        let result  = params.split("&").reduce((rt,item)=> {
            let [key,val] = item.split("=");
            rt.set(key, decodeURIComponent( val ) );
            return rt;
        },values);
        return result;
    }

    export async function ajaxAsync(url = "", data:any, method = "GET", dataType = "json") {
        return new Promise<any>((resolve, reject) => {
            let ajaxSetting = {} as JQueryAjaxSettings;
            ajaxSetting.url = url;
            ajaxSetting.type = method;
            ajaxSetting.dataType = dataType;
            ajaxSetting.data = data;
            ajaxSetting.success = (response) => {
                resolve(response);
            }
            ajaxSetting.error = (xhr) => {
                reject("请求出错，请联系网站管理员");
            };
            $.ajax(ajaxSetting);
        });
    }

    export function mapToObject( map: Map<string,any> ) {
        let result = {};
        for(let [key,val] of map.entries() ) {
            result[key] =  val;
        }
        return result;
    }

    export function jquery2HtmlElements( jquery : JQuery ) : HTMLElement[] {
        let result:HTMLElement[] = [];
        jquery.each( (index,item) => result.push((item as HTMLElement)));
        return result;
    }

    export function isInteger( val : string ) {
        return /^\d+$/.test( val );
    }

    export function checkLogin() {
        if( window.location.href.indexOf("\/html\/admin\/") != -1 ) {
            $.get("/admin/check_login", function( response ){
                let data = eval("("+ response +")");
                if( data.Code == 1 ) {
                    window.top.location.href = "/html/admin/login.html";
                }
            });
        }
    }
}

chhblog.checkLogin();