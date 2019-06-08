index()
function index() {
    search_app = new Vue({
        el: '#search_app',
        data:{
            com_info:{
                com_key_id:[],//选中门店key_id
                com_name:'请选择门店',
            },
        },
        components:{
            'my-component':companyLevel
        },
        methods:{
            search: function(){
                var _this = this;
                app.search_info.select_com = _this.com_info.com_key_id;
                app.index(1);
            }
        }
    });

   app = new Vue({
        'el': '#app',
        data: {
            data:[],
            search_info:{
                select_com:[],
            },
            page: 1,
        },
        mounted: function () {
            this.$nextTick(function () {
                this.index(1);
            })
        },
        watch:{
            'page':'index_ajax',
        },
        methods: {

            index : function(page){
                var _this = this
                if(_this.search_info.select_com == ''){
                    layer.msg('请选择门店');
                    return ;
                }
                API('reports/employee/results',{"select_com":_this.search_info.select_com,"page":page}, function(d){
                    if (d.status == 200){   
                        // data = d.data,
                        // _this.data = data.data;
                         _this.data = d.data;
                        //page_set(data.last_page,page,_this);
                    }else{
                        layer.msg(d.statusText);
                    }
                })
            },
            // index_ajax : function(page){
            //     var _this = this
            //     API('reports/employee/results',{"start_date":_this.search_info.start_date,"end_date":_this.search_info.end_date,"select_com":_this.search_info.select_com,"page":page}, function(d){
            //         if (d.status == 200){   
            //             _this.data = d.data
            //         }else{
            //             layer.msg(d.statusText);
            //         }
            //     })
            // },
        }
    }); 
};
