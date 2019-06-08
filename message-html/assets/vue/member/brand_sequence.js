// switch(module){
//     case 'product':
//         index(1);
//         break;
//     case 'project':
//         index(2);
//         break;
//     case 'package':
//         index(3);
//         break;
//     default:
//         index(0);//默认"品牌销售额"选项卡
// }

index();
function index() {
    search_app = new Vue({
        el: '#search_app',
        data:{
            start_date:'',
            end_date:'',
            
            com_info:{
                com_key_id:[],//选中门店key_id
                com_name:'默认当前门店（包括下级）',//门店名
            },
        },
        components:{
            'my-component':companyLevel
        },
        methods:{
            search: function(){ 
                var _this = this;
                app.search_info.select_com  = _this.com_info.com_key_id;
                app.search_info.start_date = _this.start_date;
                app.search_info.end_date = _this.end_date;

                //app.search_info.brand = _this.brand;
                app.index();
            }
        }
    });

    var app = new Vue({
        el: '#app',
        data: {
            page : 1,
            flag : 0,
            records : 15,
            items: [],
            sales_info: [],
            keyword:'', //关键词 

            active0 : 'active',
            active1 : '',
            active2 : '',
            active3 : '',
            search_info:{
                select_com:[],
                start_date:'',
                end_date:'',                
            },
        },
        mounted: function () {
            this.$nextTick(function () {
                this.index();
            })
        },
        watch:{
            page : 'index'
        },
        methods: {
            index : function(){
                var _this = this, 
                    _page = _this.page,
                    _data = {page : _page, type : _this.flag, info: _this.search_info};

                //TAB Highlight
                // for(var i=0;i<4;i++){
                //     _this['active'+i] = '';
                //     if (flag == i)
                //     _this['active'+i] = 'active';
                // }

                _this.keyword = '';
                if (_this.search_info.start_date=='' || _this.search_info.start_date=='') {
                    layer.msg('请选择开始时间和结束时间');
                    return ;
                }

                API('/reports/members/brand_sequence', _data, function(d){
                    if(d.status == 200){
                        _this.items = d.data.data;
                        _this.records = d.data.per_page;
                        page_set(d.data.last_page, _page, _this);
                        if (_this.items.length == 0){
                            _this.items = [];
                            page_set(1, 1, _this);
                        }
                    }else{
                        layer.msg(d.statusText);
                        //setTimeout(function(){location.reload();},800);
                    }
                })
            },
            //搜索查找
            search_goods : function(){
                var _this = this;
                if (_this.keyword =='') return ;
                _this.sales_info = [];

                API('/reports/members/brand_find_goods', {type:_this.flag, keyword:_this.keyword}, function(d){
                    if(d.status == 200){
                        _this.sales_info = d.data; 
                    }
                })
            },
            //当前是哪个品牌
            goods_select : function(e){
                var _this = this,
                    _data = {page : 1, type : _this.flag, brand_keyid : e.key_id, info : _this.search_info};
                    _this.keyword = e.name;
                    $('.suggestlist').hide();

                API('/reports/members/brand_sequence', _data, function(d){
                    if(d.status = 200){
                        _this.items = d.data.data;
                        _this.records = d.data.per_page;
                        page_set(d.data.last_page, _this.page, _this);
                    }
                });
            },
            s_brand:function(){
                var _t = this;
                _t.active0 = 'active';
                _t.active1 = '';
                _t.active2 = '';
                _t.active3 = '';
                _t.flag = 0;
                _t.page = 1;
                _t.index();
            },
            s_product:function(){
                var _t = this;
                _t.active0 = '';
                _t.active1 = 'active';
                _t.active2 = '';
                _t.active3 = '';
                _t.flag = 1;
                _t.page = 1;
                _t.index();
            },
            s_project:function(){
                var _t = this;
                _t.active0 = '';
                _t.active1 = '';
                _t.active2 = 'active';
                _t.active3 = '';
                _t.flag = 2;
                _t.page = 1;
                _t.index();
            },
            s_package:function(){
                var _t = this;
                _t.active0 = '';
                _t.active1 = '';
                _t.active2 = '';
                _t.active3 = 'active';
                _t.flag = 3;
                _t.page = 1;
                _t.index();
            }
        }
    });
};
