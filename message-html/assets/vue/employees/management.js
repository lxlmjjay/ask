type = typeof(uriObj.type) == 'undefined' ? 1 : uriObj.type;
show_str = type == 1 ?  '消费' : '消耗';
index()

function index() {
    Vue.prototype.type = type;
    Vue.prototype.show_str = show_str;
    search_app = new Vue({
        el: '#search_app',
        data:{
            start_date:'',
            end_date:'',
            com_info:{
                com_key_id:[],//选中门店key_id
                com_name:'默认当前门店（包括下级）',
            },
        },
        components:{
            'my-component':companyLevel
        },
        methods:{
            search: function(){
                var _this = this;
                app.search_info.start_date = _this.start_date;
                app.search_info.end_date = _this.end_date;
                app.search_info.select_com = _this.com_info.com_key_id;
                app.index();
            }
        }
    });

   app = new Vue({
        'el': '#app',
        data: {
            page:1,
            data:[],
            membeer_consumption_arr:[],
            membeer_consume_arr:[],
            search_info:{
                start_date:'',
                end_date:'', 
                select_com:[],
            },
        },
        mounted: function () {
            this.$nextTick(function () {
                this.index();
            })
        },
        watch:{
            'page':'index_ajax',
        },
        methods: {
            index : function(){
                var _this = this;
                if(_this.search_info.start_date == ''){
                    layer.msg('请选择开始时间');
                    return ;
                }
                API('reports/employee/management',{"start_date":_this.search_info.start_date,"end_date":_this.search_info.end_date,'select_com':_this.search_info.select_com,"page":1}, function(d){
                    if (d.status == 200){   
                        var data = d.data;
                        _this.data = data.data.data,
                        _this.membeer_consumption_arr = data.membeer_consumption_arr,
                        _this.membeer_consume_arr = data.membeer_consume_arr;
                        page_set(data.data.last_page,1,_this);

                    }else{
                        layer.msg(d.statusText);
                    }
                })
            },
            index_ajax : function(){
                var _this = this;
                if(_this.search_info.start_date == ''){
                    layer.msg('请选择开始时间');
                    return ;
                }
                API('reports/employee/management',{"start_date":_this.search_info.start_date,"end_date":_this.search_info.end_date,'select_com':_this.search_info.select_com,"page":_this.page}, function(d){
                    if (d.status == 200){   
                        var data = d.data;
                        _this.data = data.data.data;
                    }else{
                        layer.msg(d.statusText);
                    }
                })
            },
        }
    }); 
};
