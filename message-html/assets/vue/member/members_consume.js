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
            //临时变量
            placeholder:'请填写编号',
            which_type : 1,
        },
        components:{
            'my-component':companyLevel
        },
        watch:{
            which_type : 'icode'
        },
        methods:{
            search: function(){ 
                var _this = this;
                app.search_info.select_com  = _this.com_info.com_key_id;
                app.search_info.start_date = _this.start_date;
                app.search_info.end_date = _this.end_date;
                app.search_info.which_type = _this.which_type;
                app.search_info.mem_code = _this.mem_code;
                app.page = 1;

                app.index();
            },
            icode : function(){
                if (this.which_type == 2){
                    this.placeholder = '请填写卡号';
                }else{
                    this.placeholder = '请填写编号';
                }
            }
        }
    });

    var app = new Vue({
        el: '#app',
        data: {
            page : 1,
            flag : 1,//默认单品
            consume: [],
            active1 : 'active',
            active2 : '',
            active3 : '',
            search_info:{
                start_date:'',
                end_date:'', 
                select_com:[],
                which_type:0,
                mem_code:'',
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

                API('/reports/members/consume', _data, function(d){
                    if(d.status == 200){
                        _this.consume = d.data.data;
                        page_set(d.data.last_page, _page, _this);
                    }else{
                        layer.msg(d.statusText);
                        //setTimeout(function(){location.reload();},800);
                    }
                })
            },

            sale_product:function(){
                var _t = this;
                _t.active1 = 'active';
                _t.active2 = '';
                _t.active3 = '';
                _t.flag = 1;
                _t.page = 1;
                _t.index();
            },
            sale_project:function(){
                var _t = this;
                _t.active1 = '';
                _t.active2 = 'active';
                _t.active3 = '';
                _t.flag = 2;
                _t.page = 1;
                _t.index();
            },
            sale_package:function(){
                var _t = this;
                _t.active1 = '';
                _t.active2 = '';
                _t.active3 = 'active';
                _t.flag = 3;
                _t.page = 1;
                _t.index();
            },
        }
    }); 
};
