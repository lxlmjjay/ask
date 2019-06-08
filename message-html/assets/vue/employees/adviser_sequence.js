
index();

function index() {
    search_app = new Vue({
        el: '#search_app',
        data:{
            start_date:'',
            end_date:'',   
            mem_code : '',
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
                app.page = 1;

                app.index();
            }
        }
    });

    var app = new Vue({
        el: '#app',
        data: {
            page : 1,
            records : 15,
            items: [], 
            type: '',
            keyword : '',
            PM_title:'品牌销扣额',
            goods_keyid : '',
            suggest_info:[],
            search_info:{
                start_date:'',
                end_date:'', 
                select_com:[]
            }            
        },
        mounted: function () {
            this.$nextTick(function () {
                this.index(0);
            })
        },
        watch:{
            page : 'index',
            type : 'goods_select'
        },
        methods: {
            index : function(flag){
                var _this = this, 
                    _page = _this.page, _url,
                    _data = {page : _page, 
                             type : _this.type,
                             sort : flag,
                             goods_keyid : _this.goods_keyid, 
                             info : _this.search_info
                            };

                if (_this.search_info.start_date=='' || _this.search_info.start_date=='') {
                    layer.msg('请选择开始时间和结束时间');
                    return ;
                }

                if (Sequence == 'adviser'){
                    _url = '/reports/employee/adviser_sequence';
                    API(_url, _data, function(d){
                        if(d.status == 200){
                            _this.items = d.data;
                        }else{
                            layer.msg(d.statusText);
                        }
                    })
                }
                else{
                    _url = '/reports/employee/beautician_sequence';
                    API(_url, _data, function(d){
                        if(d.status == 200){
                            _this.items = d.data.data;
                            _this.records = d.data.per_page;
                            page_set(d.data.last_page, _page, _this);
                        }else{
                            layer.msg(d.statusText);
                            //setTimeout(function(){location.reload();},800);
                        }
                    })
                }                
            },
            goods_select : function(){
                var s, _this = this;
                if (_this.search_info.start_date=='' || _this.search_info.start_date=='') {
                    this.type = '';
                } ;

                if (_this.type == '0'){
                    //_this.goods_find();
                }else{
                    _this.suggest_info = [];
                }

                _this.index(6);

                switch (this.type){
                    case '1':
                        s = '单品销售额';
                        break;
                    case '2':
                        s = '项目销售额';
                        break;
                    case '3':
                        s = '套餐销扣额';
                        break;
                    case '4':
                        s = '项目销耗额';
                        break;
                    default:
                        s = '品牌销扣额';
                }
                _this.PM_title = s;
                _this.keyword = '';
                _this.goods_keyid = '';
            },
            goods_this : function(e){
                this.goods_keyid = e.key_id;
                this.keyword = e.name;
                this.suggest_info = [];
            },
            goods_find : function(){
                var _this = this,
                    _data = {type: _this.type, keyword : _this.keyword};

                API('/reports/employee/find_goods', _data, function(d){
                    if(d.status == 200){
                        _this.suggest_info = d.data;
                    }
                })
            },
            goods_sort : function(){
                if (this.goods_keyid !='')
                this.index();
            }  
        }
    });
};
