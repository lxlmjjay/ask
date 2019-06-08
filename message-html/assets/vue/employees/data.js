index()

function index() {
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
                if(app.page != 1){
                    app.page = 1;
                }else{
                    app.index();
                }
            }
        }
    });

    app = new Vue({
        'el': '#app',
        data: {
            page : 1,
        //    page_count : 1,
            items:[],
            list:[],
            com_info:[],
            count_info:{},
            search_info:{
                start_date:'',
                end_date:'', 
                select_com:[],
            },
        },
        watch:{
            page : 'index',
          //  page_count : 'count_btn'
        },
        mounted: function () {
            this.$nextTick(function () {
                this.index();
            })
        },
        filters:{
            work_age:function(v){
                var y  = 365*24*3600,m = y/12,year='',month='';
                if(v >= y){
                     year =parseInt(v / y)+'年';
                    v  = v%y;
                }
                if(v >= m){
                     month =parseInt(v / m)+'个月';
                }
                return year+month;
            }
        },
        methods: {
            index : function(){
                var _this = this,_page = _this.page;
                // if(_this.search_info.start_date == '' || _this.search_info.end_date == ''){
                //     layer.msg('请选择开始时间和结束时间');
                //     return 
                // }
                if(_this.search_info.start_date == ''){
                    layer.msg('请选择开始时间');
                    return 
                }
                var _data = {page : _page, "start_date":_this.search_info.start_date,"end_date":_this.search_info.end_date,"select_com":_this.search_info.select_com};
                API('reports/employee/get_data',_data, function(d){
                    if (d.status == 200){
                        _this.items = d.data.data;
                        page_set(d.data.last_page, _page, _this);
                    }else{
                        layer.msg(d.statusText);
                    }
                })
            },
            count_btn:function(){
                var _this = this,_page = _this.page_count;
                // if(_this.search_info.start_date == '' || _this.search_info.end_date == ''){
                //     layer.msg('请选择开始时间和结束时间');
                //     return
                // }
                if(_this.search_info.start_date == '' ){
                    layer.msg('请选择开始时间');
                    return
                }
                var _data = {page : _page,"start_date":_this.search_info.start_date,"end_date":_this.search_info.end_date,"select_com":_this.search_info.select_com};
                API('reports/employee/get_count',_data, function(d){
                    if (d.status == 200){
                        _this.list = d.data;
                       // _this.list = d.data.data;
                       // page_set_cont(d.data.last_page, _page, _this,'.page_count','page_count');
                    }else{
                        layer.msg(d.statusText);
                    }
                })
            }
        }
    }); 
}
