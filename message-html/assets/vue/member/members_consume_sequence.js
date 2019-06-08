// switch(module){
//     case 'sales':
//         index(2);
//         break;

//     case 'consume':
//         index(3);
//         break;

//     default:
//         index(1);
// }

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
            flag : 1,//默认为消费额
            records : 15,
            items: [],
            active1 : 'active',
            active2 : '',
            search_info:{
                start_date:'',
                end_date:'', 
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
                    _data = {page : _page, type : _this.flag, info : _this.search_info};

                //TAB Highlight
                // for(var i=1;i<=3;i++){
                //     _this['active'+i] = '';
                //     if (flag == i)
                //     _this['active'+i] = 'active';
                // }
                
                if (_this.search_info.start_date=='' || _this.search_info.start_date=='') {
                    layer.msg('请选择开始时间和结束时间');
                    return ;
                }

                API('/reports/members/consume_sequence', _data, function(d){
                    if(d.status == 200){
                        _this.items = d.data.data;
                        _this.records = d.data.per_page;
                        page_set(d.data.last_page, _page, _this);
                    }else{
                        layer.msg(d.statusText);
                        //setTimeout(function(){location.reload();},800);
                    }
                })
            },
            s_index:function(){
                var _t = this;
                _t.active1 = 'active';
                _t.active2 = '';
                _t.flag = 1;
                _t.page = 1;
                _t.index();
            },
            s_consume:function(){
                var _t = this;
                _t.active1 = '';
                _t.active2 = 'active';
                _t.flag = 2;
                _t.page = 1;
                _t.index();
            },
        }
    }); 
};
