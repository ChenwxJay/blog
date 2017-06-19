//事件基础类
( function() {
	window[ 'EventBase' ] = function() {

		this.addListener = function ( type, listener ) {
			getListener( this, type, true ).push( listener );
		}	

		this.removeListener = function ( type, listener ) {
			var listeners = getListener( this, type );
			for( var i = 0; i < listeners.length; i++ ) {
				if( listeners[ i ] == listener ) {
					 listeners.splice( i, 1 );
					 return;
				}
			}
		}	
		
		this.fireEvent = function ( type ) {
			var listeners = getListener( this, type ),
				r, t, k;
			if ( listeners ) {
				k = listeners.length;
				while ( k -- ) {
					t = listeners[k].apply( this, arguments );
					if ( t !== undefined ) {
						r = t;
					}
				}    
			}
			if ( t = this['on' + type.toLowerCase()] ) {
				r = t.apply( this, arguments );
			}
			return r;
		}		
	}
		
	function getListener( obj, type, force ) {
		var allListeners;
		type = type.toLowerCase();
		return ( ( allListeners = ( obj.__allListeners || force && ( obj.__allListeners = {} ) ) )
			&& ( allListeners[type] || force && ( allListeners[type] = [] ) ) );
	}

} )();


//文档操作基础类
(function() {
	var magic = '$ROOT';
	var root = window[ magic ] = {};
	var uidCount = 0;
    
	window[ 'UIBase' ] = function() {
	    EventBase.call( this );
		
		var self = this;
		var id;
		var globalKey;
		
		var setGlobal = function( id, obj ) {
			root[ id ] = obj;
			return magic + '[\''+ id +'\']';
		}

		var unsetGlobal = function( id ) {
			delete root[ id ];
		}
		
		var formatHtml = function( tpl ) {
			return ( tpl
				.replace( /##/g, id )
				.replace( /\$\$/g, globalKey ) );
		}    
			
		var initUIBase = function() {
			id = '$UI' + ( ++uidCount );
			globalKey = setGlobal( id, self );
		}
	
		
		this.render = function( holder ) {
			var html = this.renderHtml();
			var jquery = $( html );
			var target;
			if( holder == null ) {
				target = $( document.body );
			} else if( typeof( holder ) == 'string' ) {
				target = $( document.getElementById( holder ) );
			} else if( holder.constructor == $ ) {
				target = holder;
			} else {
			   target = $( holder );
			}
			target.append( jquery );
		}
			
		this.renderReplace = function( target ) {
		   target.replaceWith( this.renderHtml() ); 
		}
				
		this.getJDom = function() {
		   return $( document.getElementById( id ) );
		}
			
		this.getHtmlTpl = function() {
			return '';
		}
			
		this.renderHtml = function() {
			return formatHtml( this.getHtmlTpl() );
		}
			
		this.dispose = function() {
			var box = this.getJDom();
			box.remove();
			unsetGlobal( id );
		}
		
		this.getId = function() {
			return id;
		}
		
		this.getGlobalKey = function() {
			return globalKey;
		}
		
		initUIBase();
	}
})();
	


//弹窗基础类
var PopupBase = function( openAndClose ) {
    UIBase.call( this );
    
    var self = this;

    this.open = function() {
        openAndClose.open();
    }
    
    this.close = function() {
        openAndClose.close();
    }
    
    this._close = function() {
        this.fireEvent( 'mouseClickClose' );
        this.close();    
    }
    
    
    var supperDispose = this.dispose;
    this.dispose = function() {
        supperDispose.apply( this, arguments );
        this.fireEvent( 'disposed' );
    }
    
    this.initialize = function() {
        this.render( document.body );
        openAndClose.setJDom( this.getJDom() );
        openAndClose.addListener( 'opened', function() {
            self.fireEvent( 'opened' );
        });
        openAndClose.addListener( 'closed', function() {
            self.fireEvent( 'closed' );
        });
        
        this.fireEvent( 'initialized' );
    }
}


//常规弹窗
var PopupLayer = function( openAndClose, framework, content ) {
    PopupBase.call( this, openAndClose );
    framework.accep( this );
    
    var self = this;

    this.getHtmlTpl = function() {
        return framework.getContent( content );
    }
}

//会话弹窗(iframe)
var PopupDialog = function( openAndClose, framework, url ) {
    PopupBase.call( this, openAndClose );
    framework.accep( this );
    
    var self = this;
    
    this.getHtmlTpl = function() {
        var content = '<iframe style="width:100%; height:100%;" frameborder="0" scrolling="no" marginwidth="0" marginheight="0" id="##-framedialog" src="' + url + '"></iframe>';
        return framework.getContent( content );
    }
    
    this.getDialogWindow = function() {
        var iframe = this.getJDom().find( 'iframe' ).get( 0 );
        return iframe.contentWindow;    
    }
}    

//html框架
var HtmlFramework = function() {
   var self = this;
    
   this.accep = function( pop ) {
       //如需对popup对象进行操作， 可在此处进行
   }    
  
   this.getContent = function( content ) {
        throw '未实现';
   }   
}



//打开关闭形为基础类
var OpenAndCloseBase = function() {
    EventBase.call( this );
    
    var jdom;
    
    this.setJDom = function( jquery ) {
        jdom = jquery;
    }
    
    this.getJDom = function() {
        return jdom;
    }
    
    this.open = function() {
        throw new Error( '未实现' );
    }
    
    this.close = function() {
        throw new Error( '未实现' );
    }
}


//打开关闭方式类
var OpenAndClose = function() {
    OpenAndCloseBase.call( this );
    
    var self = this;
    
    var setPosition = function() {
        var jdom = self.getJDom();
        var screenWidth = $( window ).width();
        var screenHeight = $( window ).height();
        var popWidth = jdom.width();
        var popHeight = jdom.height();
        var marginLeft = popWidth / 2 ;
        var marginTop = popHeight / 2 > screenHeight / 2 ? screenHeight / 2 : popHeight / 2;
        jdom.css( {
            position: 'absolute',
            zIndex: 999,
            left: '50%',
            top: '50%',
            marginLeft: -marginLeft,
            marginTop: -marginTop
        } );    

        jdom.css( 'margin-top', '+=' + parseInt( $( document ).scrollTop() ) );
    }
    
    this.open = function() {
        setPosition();
        self.getJDom().show();
        this.fireEvent( 'opened' );
    }
    
    this.close = function() {
        self.getJDom().hide();
        this.fireEvent( 'closed' );
    }
} 

//具有移动效果的打开关闭类
var MoveOpenAndClose = function() {
    OpenAndCloseBase.call( this );
    
    var self = this;
    var moveCount = 30;
    
    var setPosition = function() {
        var jdom = self.getJDom();
        var screenWidth = $( window ).width();
        var screenHeight = $( window ).height();
        var popWidth = jdom.width();
        var popHeight = jdom.height();
        var marginLeft = popWidth / 2 ;
        var marginTop = popHeight / 2 > screenHeight / 2 ? screenHeight / 2 : popHeight / 2;
        jdom.css( {
            position: 'absolute',
            zIndex: 999,		
            left: '50%',
            top: '50%',
            marginLeft: -marginLeft,
            marginTop: -marginTop - moveCount
        } );    
    }
    
    this.open = function() {
        setPosition();
        self.getJDom().show().css( 'opacity', 0 ).animate( { opacity: 1, marginTop: '+=' + moveCount }, 300 );
        this.fireEvent( 'opened' );
    }
    
    this.close = function() {
        self.getJDom().animate( { opacity: 0, marginTop: '-=' + moveCount }, 300, 
            function() { $( this ).hide() } );
        this.fireEvent( 'closed' );
    }
} 


//遮罩类
;( function() {
    window[ 'Mask' ] = function() {
        UIBase.call( this );
        
        this.getHtmlTpl = function() {
            var html = '<div id="##"></div>'
            return html;
        }
        
        var getPageSize = function() {
            var jdoc = $( document );
            var size = { width: parseInt( jdoc.width() ), height: parseInt( jdoc.height() ) };    
            size.height += 150;
            return size;
        }
        
        this.show = function() {
            this.getJDom().show();    
        }
        
        this.hide = function() {
            this.getJDom().hide();    
        }
        
        var oldRender = this.render;
        this.render = function() {
            oldRender.apply( this );
            var size = getPageSize();
            this.getJDom().css( { 
                opacity: 0.4, 
                width: size.width, 
                height: size.height, 
                position: 'absolute',
                left: 0,
                top: 0,
                backgroundColor: 'black',
                zIndex: 998,
                display: 'none' 
            } );    
        }
    }
    
} )();


//遮罩层调用函数
( function() {
	var mask;
	window[ 'initMask' ] = function() {
		if( !mask ) {
			mask = new Mask();
			mask.render();	
		}		
	}
	
	window[ 'showMask' ] = function() {
		if( mask ) {
			mask.show();
		} else {
			alert( '请选择调用initMask函数进行初始化' );	
		}
	}
	
	window[ 'hideMask' ] = function() {
		if( mask ) {
			mask.hide();
		} else {
			alert( '请选择调用initMask函数进行初始化' );	
		}	
	}
} )();