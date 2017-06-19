// JScript source code
var CsdnScriptPlugin999 = {
    /// <summary>
    /// 接口版本，必填项
    /// </summary>
    interfaceVersion: "1.0",

    /// <summary>
    /// 插件标题，显示给用户看，必填项
    /// </summary>
    caption: "炮炮兵表情",

    /// <summary>
    /// 设计者在CSDN的ID，必填项
    /// </summary>
    designer: "aspwebchh",

    /// <summary>
    /// 按钮对象，可选项
    /// </summary>
    buttons: {},

    /// <summary>
    /// 分隔条对象，可选项
    /// </summary>
    separators: {},

    /// <summary>
    /// 装载，必填项
    /// </summary>
    load: function() {
        this.separators["vertical"] = CsdnScriptWorkshop.addSeparator();

        this.buttons["vertical"] = CsdnScriptWorkshop.addButton(
             "炮炮兵表情",
             "http://hi.csdn.net/attachment/201103/19/3950093_1300499713aRAa.gif",
             function()
             {
                 var text = CsdnScriptWorkshop.getEditorText();
                 var point = absolutePoint(this);
                
                 var html = "";
                 html += "<div id='content'>";
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300455086eJCE.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465108no51.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465108I4i2.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465107K84q.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_13004651059Rw7.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_13004651043d6O.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465103oZUl.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465102FiIT.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465101PPu5.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465100QJCJ.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465097eue3.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465093b5Eo.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465090wcO4.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465088lQ0O.gif.thumb.jpg'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_130046571827v8.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_130046571782ZX.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465716ivsF.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465713fpB1.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_13004657131j83.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465711DrXG.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465709EC9d.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465708quQ8.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465708bZgX.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465707GBgn.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_13004657069Kvx.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465704VUQ7.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465703l08r.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465703ziKg.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465702q3AZ.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465108no51.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465108no51.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465109zP96.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465108I4i2.gif'/>"
                 html +=     "<img src='http://hi.csdn.net/attachment/201103/18/3950093_1300465107K84q.gif'/>"             
                 html += "</div>";
                 html += "<style type='text/css'>"
                 html += "#content{ text-align:left;top:8px;position:relative; width:97%;  margin:0 auto;  height:90%;}";
                 html += "#content img{ width:20px;height:20px; border:1px solid green;position:absolute;z-index:1;cursor:pointer}";
                 html += "</style>"
                
                 CsdnScriptWorkshop.showDialog("泡泡兵表情", html,point.x - 50, point.y + 15, 510, 253);
                 var content = document.getElementById("content");
                 var collection = content.getElementsByTagName("img")
                 var X = 0;
                 var Y = 0;
                 var W = content.offsetWidth;
                 //var H = content.offsetHeight;
                 for(var i = 0 ; i < collection.length ; i ++)
                 {
                     if(X >= W)
                     {
                         Y = Y + 30;
                         X = 0;
                     }
                     collection[i].style.left = X + "px";
                     collection[i].style.top = Y + "px";
                     X = X + 30;
                 }
                
                 var paopaobing = function()
                 {
                     this.onclick  = function(e)
                     {
                         e = e || window.event ; target = e.target || e.srcElement;
                         if(target.tagName == "IMG")
                         {
                             var src = target.src;
                             var imgText = "[img="+src+"][/img]";
                             CsdnScriptWorkshop.setEditorText(text + imgText);
                             CsdnScriptWorkshop.closeDialog();
                         }
                     }
                    
                     var startX ;
                     var startY ;
                     var endX;
                     var endY;
                     var objX;
                     var objY;
                    
                     this.onmouseover = function(e)
                     {
                         e = e || window.event ; target = e.target || e.srcElement;
                         if(target.tagName == "IMG")
                         {
                              target.style.width = "50px";
                              target.style.height = "50px";
                              target.style.zIndex = "2"
                             
                              objX = target.style.left;
                              objY = target.style.top;
                              objX = parseInt(objX);
                              objY = parseInt(objY);
                             
                              if(objX >= W - 30)
                              {
                                  target.style.left = objX - 30 + "px";
                              }
                             
                              if(Y == objY && Y != 0)
                              {
                                  target.style.top = Y - 30 + "px";
                              }
                             
                              startX = e.clientX;
                              startY = e.clientY;
                             
                              target.onmousemove = function()
                              {
                                 endX = e.clientX;
                                 endY = e.clientY;
                                 var deltaX = parseInt(endX) - parseInt(startX);
                                 var deltaY = parseInt(endY) - parseInt(startY);
                                 deltaX = Math.abs(deltaX);
                                 deltaY = Math.abs(deltaY);
                                
                                 if(deltaX > 15 || deltaY > 15)
                                 {
                                     restore(target,objX,objY);
                                 }                             
                              }
                         }
                     }
                    
                    
                     this.onmouseout = function(e)
                     {
                         e = e || window.event ; target = e.target || e.srcElement;
                         if(target.tagName == "IMG")
                         {
                             restore(target,objX,objY);
                         }
                     }
                    
                     var restore = function(target,X,Y)
                     {
                         target.style.width = "20px";
                         target.style.height = "20px";
                         target.style.zIndex = "1"
                         target.style.left = X + "px";
                         target.style.top = Y + "px";     
                         target.onmousemove = null;                    
                     }
                 }      
                 paopaobing.call(content);
             }
        )
    },

    /// <summary>
    /// 卸载，必填项
    /// </summary>
    free: function() {
        for (var button in this.buttons)
            CsdnScriptWorkshop.deleteButton(this.buttons[button]);
        for (var separator in this.separators)
            CsdnScriptWorkshop.deleteSeparator(this.separators[separator]);
    }
}