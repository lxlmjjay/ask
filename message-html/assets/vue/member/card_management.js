module = 'info';
switch (module) {
    case 'info':
        info();
        break;
    default:
        index();
}

//---------------------------------------------------------
var app, app2;
//会员卡管理详情
function info() {
    Vue.prototype.member_card_status = member_card_status;
    Vue.prototype.card_change_type = card_change_type;
    app2 = new Vue({
        el: '#app2',
        data: {
            keyword: ''
        },
        methods: {
            search: function () {
                if (!this.keyword) {
                    layer.msg('请输入关键词!');
                    return false;
                }
                app.keyword = this.keyword;
                if(app.page != 1){
                    app.page = 1;
                }else{
                    app.view();
                }
            }
        }
    });
    app = new Vue({
        el: '#app',
        data: {
            show_active:false,
            keyword: '',
            item: {},
            card:'',
            page: 1,
            card_change_log:[],
            change_date:{
                key_id:'',
                card_validdate:''
            },
            account_balance:[],
            total_balance: 0,
            total_amount: 0,
            card_code:'',
            transfer:{
                key_id:'',
                card_code:'',
                name:'',
                name_keyid:''
            },
           name_active: false,
           name_list: []
        },
        watch: {
            page: 'view'
        },
        methods: {
            view: function () {
                var _this = this;
                _page = _this.page;
                API('card_management/info', {page:_page,keyword: _this.keyword}, function (d) {
                    if (d.status == 200) {
                        var re = d.data;
                        _this.item =re.member ;
                        _this.card_change_log = re.card_change_log.data;
                        _this.card = re.card;
                        var time = Vue.filter('time');
                        _this.item.cardsale_date = time(_this.item.cardsale_date, 'yyyy-MM-dd');
                        _this.item.card_validdate = time(_this.item.card_validdate, 'yyyy-MM-dd');
                        _this.show_active = true;
                        _this.change_date.key_id=_this.item.key_id;
                        var pages = re.card_change_log.last_page;
                        page_set(pages, _page, _this);
                    } else {
                        layer.msg(d.statusText);
                    }
                })
            },
            check:function(type){
                var card = this.card;
                var flag = false;
                if($.inArray(String(type),card)>-1){
                    flag = true;
                }
                return flag;
            },
            //变更截止日期
            change_date_btn: function () {
                _this = this;
                if(_this.item.card_status!=1){
                    layer.msg('该会员卡不是正常状态');
                    return;
                }
                if(!_this.check(3)){
                    layer.msg('该会员卡不允许会籍延期！');
                    return;
                }
                layer.confirm('您确定要变更截止日期吗？',function(index){
                    layer.close(index);
                    modal_show('#myModal1');
                });
            },
            //变更截止日期保存
            save_change_date: function () {
                var _this = this;
                API('card_management/save_change_date', _this.change_date, function (d) {
                    if (d.status == 200) {
                        layer.msg(d.statusText);
                        _this.view();
                    } else {
                        layer.msg(d.statusText);
                    }
                })
            },
            //会员卡挂失
            card_lose_btn: function () {
                var _this = this;
                if(_this.item.card_status!=1){
                    layer.msg('该会员卡不是正常状态');
                    return;
                }
                if(!_this.check(8)){
                    layer.msg('该会员卡不允许挂失！');
                    return;
                }
                layer.confirm('您确定要挂失吗？',function(){
                    API('card_management/card_lose', {key_id:_this.item.key_id}, function (d) {
                        if (d.status == 200) {
                            layer.msg(d.statusText);
                            _this.item.card_status = 2;
                        } else {
                            layer.msg(d.statusText);
                        }
                    })
                });
            },
            //会员卡解挂
            card_find_btn: function () {
                var _this = this;
                if(_this.item.card_status!=2){
                    layer.msg('该会员卡并未挂失');
                    return;
                }
                if(!_this.check(8)){
                    layer.msg('该会员卡不允许解挂！');
                    return;
                }
                layer.confirm('您确定要解挂吗？',function(){
                    API('card_management/card_find', {key_id:_this.item.key_id}, function (d) {
                        if (d.status == 200) {
                            layer.msg(d.statusText);
                            _this.item.card_status = 1;
                        } else {
                            layer.msg(d.statusText);
                        }
                    })
                });
            },
            //会员卡退卡时获取账户余额信息
            card_refund_btn: function () {
                var _this = this;
                if(_this.item.card_status!=1){
                    layer.msg('该会员卡不是正常状态');
                    return;
                }
                if(!_this.check(6)){
                    layer.msg('该会员卡不允许退卡！');
                    return;
                }
                layer.confirm('您确定要退卡吗？',function(index){
                    layer.close(index);
                    //获取账户余额和欠款
                    API('card_management/account_balance', {key_id:_this.item.key_id}, function (d) {
                        if (d.status == 200) {
                            _this.account_balance = d.data;
                            _this.get_total();
                            modal_show('#myModal2');
                        } else {
                            layer.msg(d.statusText);
                        }
                    })
                });
            },
            get_total:function(){
                _this = this;
                _this.total_balance = 0;
                _this.total_amount = 0;
                _this.account_balance.forEach(function(v){
                    _this.total_balance += Number(v.account_balance);
                    _this.total_amount +=  Number(v.payable_amount);
                })
            },
            //会员卡退卡
            card_refund_save: function () {
                var _this = this;
                if( _this.total_amount>0){
                    layer.msg('请先到会员账户管理中还款');
                    return;
                }
                if( _this.total_balance>0){
                    layer.msg('请先到会员账户管理中还款');
                    return;
                }
                var data =  {
                    key_id:_this.item.key_id,
                    total_balance:_this.total_balance,
                    total_amount:_this.total_amount
                };
           API('card_management/card_refund',data, function (d) {
                    if (d.status == 200) {
                        layer.msg(d.statusText);
                        _this.item.card_status = 4;
                    } else {
                        layer.msg(d.statusText);
                    }
                })
            },
            //会员卡遗失补卡
            card_replace_btn: function () {
                var _this = this;
                if(_this.item.card_status!=1){
                    layer.msg('该会员卡不是正常状态');
                    return;
                }
                if(!_this.check(5)){
                    layer.msg('该会员卡不允许补卡！');
                    return;
                }
                layer.confirm('您确定要补卡吗？',function(index){
                    layer.close(index);
                    modal_show('#myModal3');
                });
            },
            //会员卡遗失补卡保存
            card_replace_save: function () {
                var _this = this;
                    API('card_management/card_replace', {key_id:_this.item.key_id,card_code:_this.card_code}, function (d) {
                        if (d.status == 200) {
                            layer.msg(d.statusText);
                            _this.item.card_code = _this.card_code;
                        } else {
                            layer.msg(d.statusText);
                        }
                    })
            },
            //会员卡会籍转让
            card_transfer_btn: function () {
                var _this = this;
                if(_this.item.card_status!=1){
                    layer.msg('该会员卡不是正常状态');
                    return;
                }
                if(!_this.check(7)){
                    layer.msg('该会员卡不允许转让！');
                    return;
                }
                layer.confirm('您确定要转让吗？',function(index){
                    layer.close(index);
                    modal_show('#myModal4');
                });
            },
            //获取转让会员姓名列表
            name_looklist: function () {
                _this = this;
                _this.transfer.name_keyid = '';
                var mobile = _this.transfer.name;
                if (mobile == "" || mobile.length < 4) {
                    _this.name_list = [];
                    _this.name_active = false;
                    return false;
                }
                API('card_management/get_mem_name', {mobile: mobile}, function (d) {
                    if (d.status == 200) {
                        _this.name_list = d.data;
                        _this.name_active = true;
                    } else {
                        layer.msg(d.statusText);
                        _this.name_list = [];
                        _this.name_active = false;
                        _this.transfer.name_keyid = '';
                        _this.transfer.name = '';
                    }
                })
            },
            //选取转让会员姓名
            name_getlist: function (k) {
                this.transfer.name_keyid = this.name_list[k].key_id;
                this.transfer.name = this.name_list[k].name;
                this.name_list = [];
                this.name_active = false;
            },
            //会员卡会籍转让保存
            card_transfer_save: function () {
                var _this = this;
                _this.transfer.key_id = _this.item.key_id;
                API('card_management/card_transfer',_this.transfer, function (d) {
                    if (d.status == 200) {
                        layer.msg(d.statusText);
                        _this.item.card_status = 4;
                    } else {
                        layer.msg(d.statusText);
                    }
                })
            }
        }
    });
}
/********************************************************************************/
/*显示模态框*/
function modal_show(str) {
    $(str).modal('show');
}