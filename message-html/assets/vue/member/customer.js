switch(module){
    case 'info_base':
        info_base();
        break;

    case 'info_lives':
        info_lives();
        break;

    case 'info_history':
        info_history();
        break;

    case 'info_experience':
        info_experience();
        break;

    default:
        index();
}

//---------------------------------------------------------
var app, app2;
//会员资料
function index(){
    Vue.prototype.group = cookie.data('group')
    //列表页搜索和分配操作按钮
    app2 = new Vue({
        el: '#app2',
        data: {
            emp_info:{},
            emp_id:'',  
            keyword:''          
        },
        methods: {
            btn_search : function(){
                if (!this.keyword){
                    layer.msg('请输入关键词!');
                    return false;
                }
                app.keyword = this.keyword;
                if(app.page != 1){
                    app.page = 1;
                }else{
                    app.view();
                }
            },
            btn_distribution : function(){
                if (this.emp_id == ''){
                    layer.msg('请选择员工！');
                    return false;
                }
                if (app.mem_ids.length == 0){
                    layer.msg('请选择会员！');
                    return false;
                }
                app.emp_id = this.emp_id;
                app.distribution();
            }
        }
    });

    //列表页
    app = new Vue({
        el: '#app',
        data: {
            page:1,
            items:[],
            emp_id:'',//分配哪位员工
            mem_ids:[], //选定哪些会员
            keyword:'',
            url:'customer_info.html?do=info_base&id='
        },
        mounted: function () {
            this.$nextTick(function () {
                this.view();
            })
        },
        watch:{
            page:'view'
        },
        filters: {
            age: function (value) {
                if (!value) return '';
                var birthday = new Date(value*1000).getFullYear();
                var nian = new Date().getFullYear();
                return nian-birthday;
            }
        },
        methods: {
            view : function(){
                var _this = this,
                    _page = _this.page,
                    _keyword = _this.keyword;

                API('/member/customer/index', {page:_page, keyword:_keyword}, function(d){
                    if(d.status==200){
                        var re = d.data, _info = [],
                            pages = re.last_page,
                            page = _this.page;
                        _this.items = re.data;
                        _info = re.user_info;

                        for (var i = _info.length - 1; i >= 0; i--) {
                            switch(_info[i].emp_duty){
                                case '003':
                                _info[i].duty_name = '经理';break;
                                case '004':
                                _info[i].duty_name = '项目组长';break;
                                default:
                                _info[i].duty_name = '员工';
                            }
                        }

                        app2.emp_info = _info;

                        page_set(pages,page,_this);
                    }else{
                        layer.msg(d.statusText);
                    }
                })
            },
            //分配员工
            distribution : function(){
                var _data = {emp_id:this.emp_id, mem_id:this.mem_ids}, 
                    idata = app2.emp_info;

                //取出此员工姓名
                for(i;i < idata.length;i++){
                    if (_data.emp_id == idata[i].key_id){
                        _data.emp_name = idata[i].emp_name;
                        break;
                    }
                }
                //return false;
                var _this = this;
                API('/member/customer/distribution', _data, function(d){
                    if(d.status==200){
                        _this.view();
                        //layer.msg(d.statusText);
                    }
                    else{

                    }

                    layer.msg(d.statusText);
                })
            }
        }
    })
}


//基本信息
function info_base(){
    Vue.prototype.sex_arr = sex_arr;
    Vue.prototype.id_type = id_type;
    Vue.prototype.marital_status = marital_status;
    if (id ==null) {
        layer.msg('id为空!');
        return false;
    }
    app = new Vue({
        el: '#app',
        data: {
           item:{}
        },
        mounted: function () {
            this.$nextTick(function () {
                this.view();
            })
        },
        methods: {
            view : function(){
                _this = this;
                API('/member/customer/info_base', {key_id:id}, function(d){
                    if(d.status == 200){
                        var time = Vue.filter('time');
                        _this.item = d.data;
                        _this.item.id_type = _this.item.id_type?_this.item.id_type:1;
                        _this.item.marital_status = _this.item.marital_status?_this.item.marital_status:2;
                        _this.item.profession_syscate_keyid = _this.item.profession_syscate_keyid?_this.item.profession_syscate_keyid:'';
                        _this.item.addtime = _this.item.addtime?time(_this.item.addtime, "yyyy-MM-dd"):'';
                        _this.item.birthday =   _this.item.birthday?time(_this.item.birthday, "yyyy-MM-dd"):'';
                        _this.item.lunar_birthday = _this.item.lunar_birthday?time(_this.item.lunar_birthday, "yyyy-MM-dd"):'';
                    }else{
                        layer.msg(d.statusText);
                    }
                })
            }
        }
    });

    app2 = new Vue({
        el: '#app2',
        methods: {
            save : function(){
                API('/member/customer/info_base_save', app.item,function(d){
                    if(d.status == 200){
                        layer.msg(d.statusText);
                    }
                    else{
                        layer.msg(d.statusText);
                    }
                })
            }
        }
    });
}

//生活状况
function info_lives(){
    Vue.prototype.sms_status = sms_status;
    if (id ==null) {
        layer.msg('id为空!');
        return false;
    }
    app = new Vue({
        el: '#app',
        data: {
            item:{},
            page_id : id
        },
        mounted: function () {
            this.$nextTick(function () {
                this.view();
            })
        },
        methods: {
            view : function(){
                _this = this;
                API('/member/customer/info_lives', {key_id:id}, function(d){
                    if(d.status == 200){
                        _this.item = d.data;
                    }else{
                        layer.msg(d.statusText);
                    }
                })
            }
        }
    });
    app2 = new Vue({
        el: '#app2',
        methods: {
            save : function(){
                API('/member/customer/info_lives_save', app.item,function(d){
                    layer.msg(d.statusText);
                })
            }
        }
    });
}

//咨询历史
function info_history(){
    Vue.prototype.body_info = member_body;
    app = new Vue({
        el: '#app',
        data: {
            page_id:id,
            items:[],
            info: {
                member_keyid:id,
                part_syscate_keyid   :'',
                consult_project  :'',
                description  :'',
                consult_date :'',
                emp_name     :''
            },
            edit_flag : false,
            my_name: cookie.data('name')
        },
        mounted: function () {
            this.$nextTick(function () {
                this.view();
            })
        },
        methods: {
            view : function(){
                var _this = this;
                
                API('/member/customer/consult', {mem_keyid: id}, function(d){
                    if(d.status == 200){
                        var time = Vue.filter('time'),
                            _exp = d.data;
                        
                        for(i in _exp)
                        _exp[i].consult_date = time(_exp[i].consult_date, "yyyy-MM-dd");
                        _this.items = _exp;

                    }else{
                        layer.msg(d.statusText);
                    }
                })
            },
            btn_add : function(){
                var _this = this,
                    _data = _this.info;

                if (!_this.info.member_keyid) return ;

                API('/member/customer/consult_save', _data, function(d){
                    if(d.status == 200){
                        _this.view();
                        //layer.msg(d.statusText);

                        //清空input
                        for(i in _this.info) _this.info[i] = '';
                    }else{                        
                    }

                    layer.msg(d.statusText);
                });
            },
            btn_edit : function(){
                var _this = this,
                    _data = _this.info;

                API('/member/customer/consult_save_edit', _data, function(d){
                    if(d.status == 200){
                        _this.view();                        

                        //取消编辑
                        _this.edit_flag = false;

                        //清空input
                        for(i in _this.info) _this.info[i] = '';
                    }else{
                        //alert(d.statusText);
                    }

                    layer.msg(d.statusText);
                });
            },
            a_del : function(i){
                var _this = this, 
                    _data = { key_id : this.items[i].key_id };

                layer.confirm('确定要删除吗？',function(){
                    API('/member/customer/consult_del', _data, function(d){
                        if(d.status == 200){
                            _this.items.splice(i,1);
                        }
                        else{
                            //alert(d.statusText);
                        }
                        layer.msg(d.statusText);
                    });
                },function(){
                    layer.close();
                });
            },
            a_edit : function(i){
                this.info = this.items[i];
                this.edit_flag = true;
            }
        }
    });
}

//美容经历
function info_experience(){
    app = new Vue({
        el: '#app',
        data: {
            page_id:id,
            items:[],
            info: {
                member_keyid:id,
                store_name  :'',
                brand_name  :'',
                header      :'',
                work_date   :'',
                description :'',
                mem_keyid   :''
            },
            edit_flag : false,
            my_name: cookie.data('name')
        },
        mounted: function () {
            this.$nextTick(function () {
                this.view();
            })
        },
        methods: {
            view : function(){
                var _this = this;
                
                API('/member/customer/experience', {mem_keyid: id}, function(d){
                    if(d.status == 200){
                        var time = Vue.filter('time'),
                            _exp = d.data;
                        
                        for(i in _exp)
                        _exp[i].work_date = time(_exp[i].work_date, "yyyy-MM-dd");
                        _this.items = _exp;

                    }else{
                        layer.msg(d.statusText);
                    }
                })
            },
            btn_add : function(){
                var _this = this,
                    _data = _this.info;

                if (!_this.info.member_keyid) return ;

                API('/member/customer/experience_save', _data, function(d){
                    if(d.status == 200){
                        _this.view();                        

                        //清空input
                        for(i in _this.info) _this.info[i] = '';
                    }else{
                        //alert(d.statusText);
                    }
                    layer.msg(d.statusText);
                });
            },
            btn_edit : function(){
                var _this = this,
                    _data = _this.info;

                API('/member/customer/experience_save_edit', _data, function(d){
                    if(d.status == 200){
                        _this.view();                        

                        //取消编辑
                        _this.edit_flag = false;

                        //清空input
                        for(i in _this.info) _this.info[i] = '';
                    }else{
                        //alert(d.statusText);
                    }
                    layer.msg(d.statusText);
                });
            },
            a_del : function(i){
                var _this = this, 
                    _data = { key_id : this.items[i].key_id };

                layer.confirm('确定要删除吗？',function(){
                    API('/member/customer/experience_del', _data, function(d){
                        if(d.status == 200){
                            _this.items.splice(i,1);                        
                        }
                        else{
                            //alert(d.statusText);
                        }
                        layer.msg(d.statusText);
                    });
                },function(){
                    layer.close();
                });
            },
            a_edit : function(i){
                this.info = this.items[i];
                this.edit_flag = true;
            }
        }
    });
}