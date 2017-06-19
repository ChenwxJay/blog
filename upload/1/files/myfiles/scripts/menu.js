( function() {
	//事件基础类
	var EventBase = function() {
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
	//--------------事件基础类代码结束---------------------------------------------------------------------

	//组合模式递归生成菜单html
	var treeNode = function() { 
		 this.nodes;
		 this.add = function(){};
		 this.getChild = function(){};
		 this.toHtml = function(){};
	}

	var NodeList = function( obj ) {
		this.id = obj && obj.id;
		this.name = obj && obj.name;
		this.url = obj && obj.url;
		this.nodes = new Array();
	}

	NodeList.prototype = new treeNode();

	NodeList.prototype.add = function( node )
	{
		this.nodes.push( node )
	}

	NodeList.prototype.getChild = function( index )
	{
		return this.nodes[index]
	}

	NodeList.prototype.toHtml = function() {
		var html = "<li class='p' category='" + this.id + "'>"
		html += "<a href='" + this.url + "' class='black'>" + this.name + "</a>";
		html += "<ul style='display:none' class='category_content' id='"+ this.id +"'>"
		for( var i = 0 ; i < this.nodes.length ; i ++ ) {
			html += this.nodes[i].toHtml();
		}
		html += "</ul>";
		html += "</li>"
		return html;
	}

	var Leaf = function( obj ) {
		this.id = obj.id;
		this.name = obj.name;
		this.url = obj.url;
	}

	Leaf.prototype = new treeNode();

	Leaf.prototype.toHtml = function() {
		var html = "<li class='c'><a href='" + this.url + "' class='blue'>" + this.name + "</a></li>";
		return html;
	}

	var TreeRoot = function() {    
		this.nodes = new Array();
		
		this.toHtml = function() {
			var html = "<div><ul>"
			for( var i = 0 ; i < this.nodes.length ; i ++ ) {
				html += this.nodes[i].toHtml();
			}
			html += "</ul>"
			html += "</div>"
			return html;
		}
	}

	TreeRoot.prototype = new NodeList();
	//-----结束--------------------------------------------------

	//菜单
	var Menu = function( dataSource ) {
		EventBase.call( this );
		this.dataSource = dataSource;
	}

	Menu.prototype.existsSub = function( id ) {
		var childs = this.getChild( id );
		return childs.length;
	}

	Menu.prototype.getChild = function( id ) {
		var nodeList = new Array();
		for( var i = 0 ; i < this.dataSource.length ; i ++ ){
			if( this.dataSource[i].pid == id ) {
				nodeList.push( this.dataSource[ i ] );
			}
		}
		return nodeList;
	}

	Menu.prototype.generateMenu = function( data ) {
		 var node ;
		 var childs = this.getChild( data.id );
		 if( childs.length > 0 ) {
			node = new NodeList( data );
		 } else {
			node = new Leaf( data );
		 }
		 for( var i = 0 ; i < childs.length ; i ++ ) {
			node.add( this.generateMenu( childs[ i ] ) );
		 }
		 return node;
	}

	Menu.prototype.initialize  = function( canvasid ) {
		var root = new TreeRoot();
		for( var i = 0 ; i < this.dataSource.length ; i ++ ) {
			var data = this.dataSource[ i ];
			if( data.pid == 0 ) {
				root.add( this.generateMenu( data ) );
			}
		}
		
		this.canvasid = canvasid; 
		this.canvas = document.getElementById( canvasid );
		this.canvas.innerHTML = root.toHtml();
		this.bindingEvent();
	}

	Menu.prototype.getDom = function() {
		return this.canvas;
	}

	Menu.prototype.bindingEvent = function( ) {
		 var self = this;
		 var menu = $( this.canvas );
		 menu.find( 'li' ).bind( 'mouseenter', function( e ) {
			var self = $( this );
			var itemWidth = 150;
			var level = getLevel( self );
			if( ( self.offset().left + ( level * itemWidth ) ) > $( window ).width() ) {
				itemWidth = -itemWidth;
			}
			$( this ).children( 'ul' ).css( { left: itemWidth} ).show();
			$( this ).addClass( 'hover' );
		 } )
		 .bind( 'mouseleave', function( ) {
			$( this ).children( 'ul' ).hide();
			$( this ).removeClass( 'hover' );
		 } );
	}

	function getLevel(  element ) {
		var level = 0;
		while( element.parents( 'ul' ).get( 0 ) && 
				 element.parents( 'div[class="categorys_list"]' ).get( 0 ) ) {
			level ++;
			element = element.parents( 'ul' );
		}
		return level;
	}

	window[ 'Menu' ] = Menu;

} )();
