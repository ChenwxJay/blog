( function() {
	var calendarID = '';
	var instanceCount = 0;

    function getStartDate( y, m ) {
        var dt = new Date( y, m, 1 );
        var week = dt.getDay();
        dt.setDate( dt.getDate() - week );
        return dt;
    }

    function getEndDate( y, m ) {
        var dt = new Date( y, m ,1 );
        dt.setMonth( dt.getMonth() + 1 );
        dt.setDate( 0 );
        return dt;
    }

    function addEventHandler(target, type, func) {
        if (target.addEventListener)
            target.addEventListener(type, func, false);
        else if (target.attachEvent)
            target.attachEvent("on" + type, func);
        else target["on" + type] = func;
    }

    function removeEventHandler(target, type, func) {
        if (target.removeEventListener)
            target.removeEventListener(type, func, false);
        else if (target.detachEvent)
            target.detachEvent("on" + type, func);
        else delete target["on" + type];
    }

    function getOffset( element ) {
        var left = element.offsetLeft;
        var top = element.offsetTop;
        element = element.offsetParent;
        while (element) {
            left += element.offsetLeft;
            top += element.offsetTop;
            element = element.offsetParent;
        }
      
        return { left: left, 
            top: top }
    }

    function checkDate(datestr) {
        var regex = /^(\d{1,4})(-|\/|.)(\d{1,2})\2(\d{1,2})$/;       
        if(!regex.test(datestr) || !isDate(datestr)) return false
            else return true;
        
        function isDate(s){
            var s    = s.replace(/\-|\/|\./g,"/");
            var p    = s.split("/");
            var y    = parseInt(p[0]);
            var m    = parseInt(p[1]) - 1;
            var d    = parseInt(p[2]);
            var a    = new Date(y,m,d);
            if(a.getFullYear()!=y || a.getMonth()!=m || a.getDate() != d) return false;
            else return true;
        }
    }    

	function Calender() {
		calendarID = 'calender_' + ( ++instanceCount );
        
        var self = this;
        var calendarEl;
        var defaultYear;
        var defaultMonth;
        var defaultDate;

		var getContentHtml = function( y, m ) {
			var startDate = getStartDate( y, m );
			var endDate = getEndDate( y, m );
			var title = '<dl class="calendar_title"><dd>日</dd><dd>一</dd><dd>二</dd><dd>三</dd><dd>四</dd><dd>五</dd><dd>六</dd></dl>';
            var content = '<ul>';
			for( var i = 0; i < 38; i++ ) {
                if( ( i + 1 ) % 7 == 0 ) {
                    content += '</ul><ul>';
                }

                if( startDate.getDate() == defaultDate && startDate.getMonth() == defaultMonth && startDate.getFullYear() == defaultYear ) {
                    content += '<li class="selected">';
                } else if( m == startDate.getMonth() ) {
                    content += '<li class="c">'
                } else {
                    content += '<li>';
                }
                content += '<a href="javascript:;">' + startDate.getDate() +'</a></li>';
                startDate.setDate( startDate.getDate() + 1 );
            }
            content += '</ul>';
            return title + content;
		}

        var getCalendarHtml = function() {
            var html =  '    <div class="calendar_tool" id="'+ calendarID +'_tool">'+
                        '        <div class="calendar_month">'+
                        '            <select id="'+ calendarID +'_month_select"><option value="0">1月</option>'+
                        '                <option value="1">2月</option>'+
                        '                <option value="2">3月</option>'+
                        '                <option value="3">4月</option>'+
                        '                <option value="4">5月</option>'+
                        '                <option value="5">6月</option>'+
                        '                <option value="6">7月</option>'+
                        '                <option value="7">8月</option>'+
                        '                <option value="8">9月</option>'+
                        '                <option value="9">10月</option>'+
                        '                <option value="10">11月</option>'+
                        '                <option value="11">12月</option></select>'+
                        '        </div>'+
                        '        <div class="calendar_year">'+
                        '            <input type="button" value="&lt;" class="calendar_year_left" id="'+ calendarID +'_year_prev"><input'+
                        '                type="text" class="calendar_year_input" id="'+ calendarID +'_year_input"><input type="button"'+
                        '                value="&gt;" class="calendar_year_right" id="'+ calendarID +'_year_next">'+
                        '        </div>'+
                        '    </div>'; 

            html += '<div class="calendar_content" id="'+ calendarID +'_date_list"></div>';

            html += '<div class="calendar_action">' +
                        '<input type="button" value="清空" id="'+ calendarID +'_clear">' +
                        '<input type="button" value="今天" id="'+ calendarID +'_today">'+
                    '</div>';

            return '<div class="calendar_body">' +  html + '</div>';
        }

        var getElement = function( id ) {
            return document.getElementById( calendarID + '_' + id );
        }

        var monthChangeAction = function() {
            var monthSelect = getElement( 'month_select' );
            var yearInput = getElement( 'year_input' );
            
            addEventHandler( monthSelect, 'click', function() {
                var month = this.value;
                var year = yearInput.value;
                getElement( 'date_list' ).innerHTML = getContentHtml( year, month );
            } );
        }

        var yearChangeAction = function() {
            var monthSelect = getElement( 'month_select' );
            var yearInput = getElement( 'year_input' );
            var yearPrev = getElement( 'year_prev' );
            var yearNext = getElement( 'year_next' );

            addEventHandler( yearInput, 'blur', function() {
                if( /[^\d]+/.test( this.value ) ) {
                    this.value = this.value.replace( /[^\d]+/g, '' );
                } 
                getElement( 'date_list' ).innerHTML = getContentHtml( yearInput.value, monthSelect.value ); 
            } );

            addEventHandler( yearPrev, 'click', function() {
                var year = yearInput.value;
                var month = monthSelect.value;
                getElement( 'date_list' ).innerHTML = getContentHtml( --year, month ); 
                yearInput.value = year;
            } );

            addEventHandler( yearNext, 'click', function() {
                var year = yearInput.value;
                var month = monthSelect.value;
                getElement( 'date_list' ).innerHTML = getContentHtml( ++year, month ); 
                yearInput.value = year;
            } );           
        }

        var initCalendar = function( y, m, d, placeholder ) {
            calendarEl = document.createElement( 'div' );
            calendarEl.id = calendarID;
            calendarEl.className = 'aspwebchh';
            calendarEl.innerHTML = getCalendarHtml();
            
            placeholder = placeholder ? document.getElementById( placeholder ) : document.body;
            placeholder.appendChild( calendarEl );

            refreshCalender( y, m, d );

            monthChangeAction();
            yearChangeAction();

            addEventHandler( getElement( 'date_list' ), 'click', function( e ) {
                e = e || window.event;
                var t = e.target || e.srcElement;
                if( t.tagName != 'A' ) {
                    return;
                }

                var year = getElement( 'year_input' ).value;
                var month = getElement( 'month_select' ).value;
                var date = t.innerHTML;

                if( typeof( self.onSelected ) == 'function' ) {
                    self.onSelected( parseInt( year ), parseInt( month ), parseInt( date ) );
                }
            } );

            addEventHandler( getElement( 'clear' ), 'click', function() {
                if( typeof( self.onClear ) == 'function' ) {
                    self.onClear();
                }
            } );

            addEventHandler( getElement( 'today' ), 'click', function() {
                if( typeof( self.onToday ) == 'function' ) {
                    self.onToday();
                }
            } );            
        }

        var refreshCalender = function( y, m, d ) {
            getElement( 'month_select' ).value = m;
            getElement( 'year_input' ).value = y;            
            getElement( 'date_list' ).innerHTML = getContentHtml( y, m );
        }

        this.onClear = function() {};
        this.onSetToday = function() {};
        this.onSelected = function( y, m, d ) {};

        this.setDate = function( y, m, d ) {
            defaultYear = y;
            defaultMonth = m;
            defaultDate = d;
            refreshCalender( y, m, d );
        }

        this.position = function( left, top ) {
             calendarEl.style.display = 'block';
             calendarEl.style.position = 'absolute';
             calendarEl.style.left = left + 'px';
             calendarEl.style.top = top + 'px';
        }

        this.render = function( placeholder ) {
            var now = new Date();
            defaultYear = now.getFullYear();
            defaultMonth = now.getMonth();
            defaultDate = now.getDate();
            initCalendar( defaultYear, defaultMonth, defaultDate, placeholder );
        }

        this.hide = function() {
            calendarEl.style.display = 'none';
        }

        this.contains = function( target ) {
            return calendarEl.contains( target );
        }
	}

    window.Calender = Calender;


    function closeHandler( e ) {
        e = e || window.event;
        var t = e.target || e.srcElement;
        if( single.contains( t ) ) {
            return;
        }
        if( t == element ) {
            return;
        }
        single.hide();
    }

    var single;
    var element;

    function calendar() {
        var e = calendar.caller.arguments[0] || window.event;
        element = e.target || e.srcElement; 
        var offset = getOffset( element );

        removeEventHandler( document.documentElement, 'click', closeHandler );

        if( !element.checkDateAction ) {
            addEventHandler( element, 'blur', function() {
                if( this.value != '' && this.value != undefined && !checkDate( this.value ) ) {
                    alert( '日期格式不正确' );
                    this.value = '';
                }
            } );
            element.checkDateAction = true;
        }

        if( !single ) {
            single = new Calender();
            single.onSelected = function( y, m, d ) {
                var datestr = y + '-' + ( m + 1 )  + '-' + d;
                element.value = datestr;
                this.hide();
            }
            single.onToday = function() {
                var now = new Date();
                element.value = now.getFullYear() + '-' + ( now.getMonth() + 1 ) + '-' + now.getDate();
                this.hide();
            }
            single.onClear = function() {
                element.value = '';
                this.hide();
            }
            single.render();
        }

        var date = element.value;
        if( checkDate( date ) ) {
            var date    = date.replace(/\-|\/|\./g,"/");
            var ymd    = date.split("/");
            var y    = parseInt(ymd[0]);
            var m    = parseInt(ymd[1]) - 1;
            var d    = parseInt(ymd[2]);
            single.setDate( y, m, d );
        }

        single.position( offset.left, offset.top + element.offsetHeight );
        addEventHandler( document.documentElement, 'click', closeHandler );
    }
    window.calendar = calendar;
} )();

