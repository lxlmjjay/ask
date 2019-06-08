switch(module){
    case 'sales':
        index(2);
        break;

    case 'consume':
        index(3);
        break;

    default:
        index(1);
}

function index(flag) {
    search_app = new Vue({
        el: '#search_app',
        data:{       
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
                app.search_info.select_com = _this.com_info.com_key_id;
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
            suggest_info:[],
            search_info:{
                select_com:[]
            }
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
                    _data = {page : _page, info : _this.search_info};

                if (_this.search_info.start_date=='' || _this.search_info.start_date=='') return ;

                API('/reports/employee/workage', _data, function(d){
                    if(d.status == 200){
                        _this.items = d.data.data;
                        _this.records = d.data.records;
                        page_set(d.data.last_page, _page, _this);
                    }else{
                        layer.msg(d.statusText);
                        //setTimeout(function(){location.reload();},800);
                    }
                })
            }

        }
    }); 
};
