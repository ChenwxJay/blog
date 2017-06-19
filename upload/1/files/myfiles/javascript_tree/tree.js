//事件基础类
( function() {
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
			
    window['EventBase'] = EventBase;			
	/**
	 * 获得对象所拥有监听类型的所有监听器
	 * @public
	 * @function
	 * @param {Object} obj  查询监听器的对象
	 * @param {String} type 事件类型
	 * @param {Boolean} force  为true且当前所有type类型的侦听器不存在时，创建一个空监听器数组
	 * @returns {Array} 监听器数组
	 */
	function getListener( obj, type, force ) {
		var allListeners;
		type = type.toLowerCase();
		return ( ( allListeners = ( obj.__allListeners || force && ( obj.__allListeners = {} ) ) )
			&& ( allListeners[type] || force && ( allListeners[type] = [] ) ) );
	}

})();


;( function() {
	var magic = '$ROOT';
	var root = window[ magic ] = {};
	var uidCount = 0;

	var UIBase = function() {
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
			id = 'UI-NODE-' + ( ++uidCount );
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

	UIBase.getMagic = function() {
		return magic;
	}

	UIBase.get = function( id ) {
		return root[ id ];
	}

	window[ 'UIBase' ] = UIBase;
} )();  


;( function() {
	//是否可编辑
	var editable;
	//是否可拖动
	var dragable;

	//是否有节点在拖动
	var isDrag = false;
	//拖动类型
	var POS_TYPE_IN = 'in';
	var POS_TYPE_NEXT  = 'next';
	//节点接口
	var INode = function( id, value ) {
		UIBase.call( this );

		var self = this;
		var childNodes = [];
		var placeholder;

		this.add = function( node ) {
			childNodes.push( node );
		}

	    this.getChild = function() {
	    	return childNodes;
	    }

	    this.initialize = function() {
	    	this.fireEvent( 'onrender' );
	    	for( var i = 0; i < childNodes.length; i++ ) {
	    		childNodes[ i ].initialize();
	    	}
	    }

	    this.addListener( 'onrender', function() {
	    	if( !editable ) {
	    		$( '#' + this.getId() + '-add' ).hide();
	    		$( '#' + this.getId() + '-edit' ).hide();
	    		$( '#' + this.getId() + '-remove' ).hide();
	    		return;
	    	}

	    	var button = self.getJDom().children( 'a' );

			$( '#' + this.getId() + '-add' ).bind( 'click', function() {
				self.fireEvent( 'onAddNode' );
			} );

			$( '#' + this.getId() + '-edit' ).bind( 'click', function() {
				button.after( '<input type="text" id="'+ self.getId() +'-input" value="'+ button.text() +'" size="12"/>' );
				button.hide();
				var input = $( '#' + self.getId() + '-input' );
				input.select();
				input.bind( 'blur', function() {
					button.show().text( this.value );
					input.remove();
					self.fireEvent( 'onEditNode', this.value );
				} );
			} );

			$( '#' + this.getId() + '-remove' ).bind( 'click', function() {
				self.fireEvent( 'onRemoveNode' );
			} );

			this.getJDom().bind( 'mouseover', function( e ) {
				var t = e.target || e.srcElement;
				var ul = $( this ).children( 'ul' ).get( 0 );
				if( ul && ul.contains( t ) ) return;

				$( '#' + self.getId() + '-add' ).show();
				$( '#' + self.getId() + '-edit' ).show();
				$( '#' + self.getId() + '-remove' ).show();
			} );

		    this.getJDom().bind( 'mouseout', function( e ) {
				var t = e.target || e.srcElement;
				var ul = $( this ).children( 'ul' ).get( 0 );
				if( ul && ul.contains( t ) ) {
					$( '#' + self.getId() + '-add' ).hide();
					$( '#' + self.getId() + '-edit' ).hide();
					$( '#' + self.getId() + '-remove' ).hide();
				}
			} );

		    this.getJDom().bind( 'mouseleave', function( e ) {
				$( '#' + self.getId() + '-add' ).hide();
				$( '#' + self.getId() + '-edit' ).hide();
				$( '#' + self.getId() + '-remove' ).hide();
			} );
	    } );

		//拖动部份
		this.addListener( 'onrender', function() {
		    var title = this.getJDom().children( 'a' );			
	    	var element;
			var elementX, elementY, pointX, pointY;
			var left, top;
			var self = this;
			var timer;
			var ref;
			var posType;

			var mouseDown = function( e ) {
				timer = setTimeout( function() {
					//alert( value )
					 isDrag = true;
					 element = $( '<span style="font-size:12px;color:#fff;background:#000;padding:2px;">'+ value +'</span>' );
					 $( 'body' ).append( element );	
					 element.css(  {
					 	position: 'absolute',
					 	left: e.pageX,
					 	top: e.pageY
					 } );

					 $( document ).bind( 'mousemove', mouseMove );
			         $( document ).bind( 'mouseup', mouseUp )		 
				     elementX = parseFloat( element.css( 'left' ) );
			         elementY = parseFloat( element.css( 'top' ) );
			         pointX = e.pageX;
			         pointY = e.pageY;		 	
	         		
	         		 self.close();

					 if( this.setCapture ) this.setCapture();		         		 
				}, 1000 );
			 	 e.stopPropagation();
				 e.preventDefault();
			}

			var mouseMove = function( e ) {
				if( !isDrag ) return;
				var btns = $( '#'+ placeholder +' a' );
				btns.each( function() {
					if( $( this ).is( ':hidden' ) ) return;

					var point = $( this ).offset();

					if( point.left < e.pageX && e.pageX < point.left + $( this ).width() &&
						point.top < e.pageY && e.pageY < point.top + $( this ).height() ) {

						if( point.top < e.pageY && e.pageY < point.top + $( this ).height() / 2 ) {
							btns.removeClass( 'selected' );
							btns.css( 'border', '' )

							$( this ).addClass( 'selected' );
							posType = POS_TYPE_IN;
						} else {
							btns.removeClass( 'selected' );
							btns.css( 'border', '' )
							
							$( this ).css( 'borderBottom', '2px solid blue' );
							posType = POS_TYPE_NEXT;
						}
						ref = $( this ).attr( 'nodeid' );
					}
				} );

				left = elementX + e.pageX - pointX;
				top = elementY + e.pageY - pointY;
			    element.css( 'left', left );
			    element.css( 'top', top );
				e.stopPropagation();
				e.preventDefault();
			}
			
			var mouseUp = function( e ) {
				isDrag = false;
			    element.remove();
			    $( document ).unbind( 'mousemove', mouseMove );
				$( document ).unbind( 'mouseup', mouseUp );
				e.stopPropagation();
				e.preventDefault();
			    if( this.releaseCapture ) this.releaseCapture();

			    if( self == UIBase.get( ref ) ) {
			    	return;
			    }

			    self.fireEvent( 'onDragUp', {
			    	from: self,
			    	to: UIBase.get( ref ),
			    	posType: posType
			    } );
			}

			//拖动节点功能
			if( dragable ) {
				title.bind( 'mousedown', mouseDown );
				title.bind( 'mouseup', function() {
					window.clearTimeout( timer );
				} );
			}
		} );

		this.getSource = function() {
			return { id: id };
		}

		this.setTreeDomID = function( place ) {
			placeholder = place;
		}

	    this.open = new Function();
	    this.close = new Function();
	    this.isOpen = new Function();
	}

	//普通节点
	var TreeNodeBase = function( id, value ) {
		INode.call( this, id, value );

		var self = this;

		this.getHtmlTpl = function() {
			var childNodes = this.getChild();		
			var html = '';
			for( var i = 0; i < childNodes.length; i++ ) {
				html += childNodes[ i ].renderHtml();
			}
			html = '<li id="##" class="'+ this.nodeClass +'">'+
						'<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="'+ this.iconClass1 +'">' +
						'<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="'+ this.iconClass2 +'">' +
						'<a href="javascript:;" nodeid="##">'+ value +'</a>'+
						'<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="add" id="##-add">' +
						'<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="edit" id="##-edit">' +
						'<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="remove" id="##-remove">' +										
						'<ul style="display:none;">' + html + '</ul>'+
					'</li>';
			return html;
		}

		this.replace = function( newNode ) {
			this.getJDom().replaceWith( newNode.renderHtml() );
			newNode.initialize();
			newNode.open();
		}

		this.open = function() {
			var jNode = this.getJDom();
			var list = jNode.children( 'ul' );
			var icon1 = jNode.children( 'img' ).get( 0 );
			var icon2 = jNode.children( 'img' ).get( 1 )
			list.show();
			icon1.className = this.iconClass1 + '-open';
			icon2.className = this.iconClass2 + '-open'

			this.fireEvent( 'onOpen' );
		}

		this.close = function() {
			var jNode = this.getJDom();		
			var list = jNode.children( 'ul' );
			var icon1 = jNode.children( 'img' ).get( 0 );
			var icon2 = jNode.children( 'img' ).get( 1 )
			list.hide();
			icon1.className = this.iconClass1;
			icon2.className = this.iconClass2;

			this.fireEvent( 'onClose' );		
		}

		this.isOpen = function() {
			var list = this.getJDom().children( 'ul' );
			return !list.is( ':hidden' );
		}

		this.addListener( 'onrender', function() {
			var jNode = this.getJDom();
			jNode.children( 'a' ).bind( 'click', function() {
				if( !self.isOpen() ) {
					self.open();
				} else {
					self.close();
				}

				self.fireEvent( 'onclick' );			
			} );
		} );
	}

	var TreeNode = function( id, value ) {
		TreeNodeBase.apply( this, [ id, value ] );

		this.nodeClass = 'folder'
		this.iconClass1 = 'dash-f';
		this.iconClass2 = 'icon-f';

	}

	var TreeNodeLast = function( id, value ) {
		TreeNodeBase.apply( this, [ id, value ] );

		this.nodeClass = '';
		this.iconClass1 = 'dash-fl';
		this.iconClass2 = 'icon-fl';	
	}

	//叶节点
	var TreeLeafBase = function( id, value ) {
		INode.call( this, id, value );

		var self = this;

		this.getHtmlTpl = function() {
			return '<li id="##">'+
						'<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="'+ this.iconClass1 +'">'+
						'<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="'+ this.iconClass2 +'">'+
						'<a href="javascript:;" class="leaf" nodeid="##">'+ value +'</a>'+	
						'<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="add" id="##-add">' +
						'<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="edit" id="##-edit">' +
						'<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="remove" id="##-remove">' +											
					'</li>'
		}

		this.replace = function( newNode ) {
			this.getJDom().replaceWith( newNode.renderHtml() );
			newNode.initialize();
			newNode.open();
		}	

		this.addListener( 'onrender', function() {
			this.getJDom().children( 'a' ).bind( 'click', function() {
				self.fireEvent( 'onclick' );
			} );		
		} );
	}

	var TreeLeaf = function( id, value ) {
		TreeLeafBase.apply( this, [ id, value ] );

		this.iconClass1 = 'dash-t';
		this.iconClass2 = 'icon-t';
	}

	var TreeLeafLast = function( id, value ) {
		TreeLeafBase.apply( this, [ id, value ] );

		this.iconClass1 = 'dash-tl';
		this.iconClass2 = 'icon-tl';
	}

	//根节点
	var RootNode = function() {
		INode.call( this );

		this.getHtmlTpl = function() {
			var childNodes = this.getChild();	
			var html = '';
			for( var i = 0; i < childNodes.length; i++ ) {
				html += childNodes[ i ].renderHtml();
			}
			html = '<div id="##" class="tree">\
						<ul>\
							<li class="root">\
								<img src="http://www.chhblog.com/upload/1/files/88f382d2373941e38231a41367913b71.gif" class="icon-open">\
								<a href="javascript:;">树的根节点</a>\
								<ul style="margin:0;padding:0">' + html + '</ul>\
							</li>\
						</ul>\
					</div>';
			return html;
		}
	}

	//自定义根节点(单个目录)
	var CustomRootNode = function( id, value ) {
		TreeNodeBase.apply( this, [ id, value ] );

		var superGetHtmlTpl = this.getHtmlTpl;

		this.getHtmlTpl = function() {
			return '<div class="tree" id="##-root"><ul>'+ superGetHtmlTpl.call( this ) +'</ul></div>';
		}

		this.replace = function( newNode ) {
			$( document.getElementById( this.getId() + '-root' ) ).replaceWith( newNode.renderHtml() );
			newNode.initialize();
			newNode.open();
		}

		this.nodeClass = '';
		this.iconClass1 = 'dash-f-s';
		this.iconClass2 = 'icon-f';	
	}


	//节点类型
	var NODE_TYPE_CATE = 'category';
	var NODE_TYPE_ATTR = 'attr';
	var NODE_TYPE_ATTRVALUE = 'attr-value';

	//树
	var Tree = function( data ) {
		EventBase.call( this );

		var newNodeCount = 0;
		var self = this;

		var insertAfter = function( refid, nextid ) {
			var ref = findItem( refid );
			var next = findItem( nextid );
			next.pid = ref.pid;
			next.sort = ref.sort + 1;
		}

		var findItem = function( id ) {
			for( var i = 0; i < data.length; i++ ) {
				if( data[ i ].id == id ) {
					return data[ i ];
				}
			}
		}

		var findItems = function( pid ) {
			var nodelist = [];
			for( var i = 0; i < data.length; i++ ) {
				if( data[ i ].pid == pid ) {
					nodelist.push( data[ i ] );
				}
			}
			return nodelist;
		}

		var removeItem = function( id ) {
			for( var i = 0; i < data.length; i++ ) {
				if( data[ i ].id == id ) {
					data.splice( i, 1 );
					break;
				}
			}

			var childItems = findItems( id );
			for( var i = 0; i < childItems.length; i++ ) {
				removeItem( childItems[ i ].id );
			}
		}	

		var isLeafNode = function( node ) {
			var childItems = findItems( node.id );
			return childItems.length == 0;
		}

		var genTreeNode = function( ds, islast ) {
			var node;
			if( ds.pid == -1 ) {
				var root = new RootNode( ds.id, ds.name );
				var childItems = findItems( ds.id );
				for( var i = 0; i < childItems.length; i++ ) {
					root.add( genTreeNode( childItems[ i ], i == childItems.length - 1 ) );
				}  
			    node = root;
			} else {
				var isleaf = isLeafNode( ds );
				if( !isleaf ) {
					node = !islast ?  new TreeNode( ds.id, ds.name ) : new TreeNodeLast( ds.id, ds.name );
					var items = findItems( ds.id );
					for( var i = 0; i < items.length; i++ ) {
						node.add( genTreeNode( items[ i ], i == items.length - 1 ) );
					}
				} else {
					ds.isOpen = false;
					node = !islast ? new TreeLeaf( ds.id, ds.name ) : new TreeLeafLast( ds.id, ds.name );
				}
			}

			node.setTreeDomID( placeholder );
			
			node.addListener( 'onclick', function() {
				self.fireEvent( 'onNodeClick', ds );
			} );

			node.addListener( 'onAddNode', function() {
				if( !confirm( '确定要添加吗？' ) ) {
					return;
				}
				
				ds.newName = '结点' + ( ++newNodeCount );
				self.fireEvent( 'onAddNode', ds, function( newNodeData ) {
					data.push( newNodeData );
					var newNode = genTreeNode( ds, islast );
					node.replace( newNode );
				} );
			} );

			node.addListener( 'onEditNode', function( eventType, eventArg ) {
	
				ds.newName = eventArg;
				self.fireEvent( 'onEditNode', ds, function() {
					ds.name = eventArg;
				} );		
			} );

			node.addListener( 'onRemoveNode', function() {
				if( !confirm( '确定要删除节点吗？' ) ) {
					return;
				}

				self.fireEvent( 'onRemoveNode', ds, function() {
					removeItem( ds.id );
					self.reRender();
				} );
			} );

			node.addListener( 'onrender', function() {
				if( ds.isOpen ) {
					node.open();
				}
			} );

			node.addListener( 'onOpen', function() {
				ds.isOpen = true;
			} );

			node.addListener( 'onClose', function() {
				ds.isOpen = false;
			} );

			node.addListener( 'onDragUp', function( t, e ) {
				var posType = e.posType;
				var fromid = e.from.getSource().id;
				var toid = e.to.getSource().id;
				if( posType == POS_TYPE_IN ) {
					var fromitem = findItem( fromid );
					fromitem.pid = toid;
				} else {
					insertAfter( toid, fromid );
				}		
				self.reRender();
				self.fireEvent( 'onDrag', data );				
			} );

			return node;		
		}

		var placeholder;

		this.render = function( p ) {
			var rootData = data[ 0 ];
			data.sort( function( a, b ) {
				return a.sort - b.sort;
			} );

			placeholder = p;
			$( document.getElementById( p ) ).empty();
			if( data.length == 0 ) return;

			var root = genTreeNode( rootData );
			$( document.getElementById( placeholder ) ).html( root.renderHtml() );
			root.initialize();
		}  

		//重新渲染过后不可调用原对象方法
		this.reRender = function() {
			this.render( placeholder );
		}

		this.getDataSource = function() {
			return data;
		}

		this.setEditable = function( value ) {
			editable = value;
		}

		this.setDragable = function( value ) {
			dragable = value;
		}
	}

	window[ 'Tree' ]  = Tree;
} )();
