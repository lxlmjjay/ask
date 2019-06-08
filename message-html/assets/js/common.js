var i  = 0, //循环常用变量
    id = 0,
    _tmp , //作为临时变量
    uriObj = {} , //url后所有的参数(对象)
    module = 'index',//当前页所属模块,默认列表页
    baseURL = 'http://dfm-api:3202',//后台地址  'http://118.31.248.114:1201'
    layerLoadShade = {shade: ['0.1','#676a6c']},
    layerLoadType = 0


//---------页面打开时取URL参数-----------------------------------------------
var _url = {
	parameter : function(){
 		var _uri = window.location.search , _paraArr;
		if (_uri != ''){
		    _uri = _uri.replace('?','');
		    _paraArr = _uri.split('&');
		    for(; i < _paraArr.length; i++){
		        _tmp = _paraArr[i].split('=');
		        uriObj[_tmp[0]] = _tmp[1];
		    }
		    module = typeof(uriObj.do) == 'undefined' ? module : uriObj.do;
		    id = typeof(uriObj.id) == 'undefined' ? id : uriObj.id;
		}
	}
}
_url.parameter();
//console.log(uriObj)
//console.log(uriObj.id)
//console.log(uriObj.参数名)


//---------用户交互验证本地存储-----------------------------------------------
var cookie = {
	data:function(item){
		var _arr = [], _obj = {}, u = cookie.get('user_report');
		if (!u) return '';
		_arr = u.split('||');
		if (_arr.length != 7) return '';

		_obj['code'] = _arr[0];
		_obj['name']  = _arr[1];
		_obj['header'] = _arr[2];
		_obj['jobs']  	= _arr[3];
		_obj['group']  	 = _arr[4];
		_obj['coupon']   = _arr[5];
		_obj['com_name'] = _arr[6];

		return _obj[item];
	},
    set:function(key,val,time){//设置cookie方法
        var date = new Date(); //获取当前时间
        	date.setTime(date.getTime() + time * 3600 * 1000); //格式化为cookie识别的时间(小时)
        document.cookie = key + "=" + val + ";expires="+ date.toGMTString();  //设置cookie
    },
    get:function(key){
        /*获取cookie参数*/
        var arr,reg = new RegExp("(^| )"+ key +"=([^;]*)(;|$)");
		if(arr = document.cookie.match(reg))
			return unescape(arr[2]);
		else
			return null;
	},
    delete:function(key){
        var date = new Date(); //获取当前时间
        	date.setTime(date.getTime() - 10000); //将date设置为过去的时间
        var val = cookie.get(key);
		if(val != null)
		document.cookie = key + "="+ val +";expires="+ date.toGMTString();
    }
}
 
//-----------接口交互--------------------------------------------------------
function API(url, data, fun,status_layer){
	if(status_layer !== false){
		var index_layer = layer.load(layerLoadType,layerLoadShade);
	}
	
	//先验证,　超时或未登录则转向
	_tmp = cookie.get('jwt_report');
	//arguments参数数组以外面传进来为主
	if (arguments.length == 3 && !_tmp)
	{
		window.location.href = '/index.html';
		return false;
	}

	var config = {
		method: 'post',
		url: url,
		//timeout: 1000*10,
		data: data || {}
	};

	// if (_tmp == null){
	// 	//用户姓名不为空说明已经登录过
	// 	if (cookie.get('user_name') != null){
	// 		//重新验证
	// 		//layer.open()
	// 	}
	// 	else{
	// 		//window.location.href = '/login.html';
	// 	}
	// 	//return false;
	// }
	config.data['auth'] = _tmp;
	//交互初始设置
	axios.defaults.baseURL = baseURL;
	//axios.defaults.headers.post['ssssssssssssss'] = 123;
	axios.defaults.transformRequest=[function (data) {
	    // Do whatever you want to transform the data
	    return Qs.stringify(data)
	}];

	axios(config).then(function(response){
		/*
		console.log(response.data);
	    console.log(response.status);
	    console.log(response.statusText);
	    console.log(response.headers);
	    cosole.log(response.config);
	    */		
		if(status_layer !== false){
			layer.close(index_layer)
		}
		fun(response.data);
	}).catch(function (response){
	    console.log(response);
	});	
}

/**
 * [idate description]
 * @param   _this [description]
 * @param   app   [description]
 * @param   param [description]
 * @param   type  [year、month、time、datetime]
 * @return        [description]
 */
 function idate(_this,app,param,type) {
 	if(!type)
 		type = 'datetime';
	//时间选择器
    laydate.render({
        elem: _this
        , show: true //直接显示
        , type: type
        , done: function (value, date, endDate) {
            app[param] = value 
        }
    });

}

//分页
function page_set(pages, page, _this){
	var config = {
		cont: $('.pagination'),
		pages: pages,
		skin:'#00AA91',
		curr: function(){ //通过url获取当前页，也可以同上(pages)方式获取
			return page ? page : 1;
		}(),
		jump: function(e, first){ //触发分页后的回调
			if(!first){ //一定要加此判断，否则初始时会无限刷新
				//_this.index(e.curr); //调用当前对象中的index方法
				_this.page = e.curr;
			}
		}
	}
	//合并对象
	//if (typeof arguments[3] == 'Object')
	//config.extend(config, _obj);

	laypage(config);
}
  

//保留小数位
function round2(number,fractionDigits){  
    with(Math){  
        return round(number*pow(10,fractionDigits))/pow(10,fractionDigits);  
    }  
}

//-----------Vue-Js日期格式化--------------------------------------------------------
//yyyy-MM-dd hh:mm:ss
Vue.filter('time', function (value, fmt,str) {
    if(typeof(str) == 'string' && !Boolean(value)){
        return  str;
    }
	if(typeof(value) == 'undefined'){
		fmt = '--';
	}else{
		 var date = new Date(value * 1000);
	    var o = {
	        "M+": date.getMonth() + 1, //月份
	        "d+": date.getDate(), //日
	        "h+": date.getHours(), //小时
	        "m+": date.getMinutes(), //分
	        "s+": date.getSeconds(), //秒
	        "q+": Math.floor((date.getMonth() + 3) / 3), //季度
	        "S": date.getMilliseconds() //毫秒
	    };
	    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (date.getFullYear() + "").substr(4 - RegExp.$1.length));
	    for (var k in o)
	        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
	}
    return fmt;
});

function number_set(e){
	var value = parseFloat($(e).val());
	if(value < 0){
		value = 0
	}
	$(e).val(value)
} 