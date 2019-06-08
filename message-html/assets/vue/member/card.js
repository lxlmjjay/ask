switch (module) {
    case 'info':
        info();
        break;
    default:
        index();
}

//---------------------------------------------------------
var app, app2;
//会员卡资料列表
function index() {
    Vue.prototype.member_card_status = member_card_status;
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
            page: 1,
            items: [],
            keyword: ''
        },
        mounted: function () {
            this.$nextTick(function () {
                this.view();
            })
        },
        watch: {
            page: 'view'
        },
        methods: {
            view: function () {
                _this = this;
                var _this = this,
                    _page = _this.page;
                API('/card/index', {page: _page, keyword: this.keyword}, function (d) {
                    if (d.status == 200) {
                        var re = d.data;
                        _this.items = re.data;
                        var time = Vue.filter('time');
                        _this.items.forEach(function (v) {

                            v.cardsale_date = time(v.cardsale_date, 'yyyy-MM-dd');
                            v.card_validdate = time(v.card_validdate, 'yyyy-MM-dd');

                        });

                        page_set(re.last_page, _page, _this);

                    } else {
                        layer.msg(d.statusText);
                    }
                })
            }
        }
    });
}
function info() {
    Vue.prototype.card_status = member_card_status;
    Vue.prototype.card_change_type = card_change_type;
    Vue.prototype.operate_type = operate_type;
    Vue.prototype.orders_status = orders_status;
    Vue.prototype.price_type = price_type;
    
    app = new Vue({
        el: '#app',
        data: {
            card: {},
            level: [],
            total_balance: 0,
            total_amount: 0,
            account_history:[],
            product:[],
            project:[],
            packages:[],
            project_use:[],
            com_code:[],
            product_unit_price:[],
            trade_history:[],
            change_history:[],
            integral_record:[],
            storein_history:[],
            storeout_history:[],
            in_num:1,
            out_num:1,
            coupon_history:[],
        },
        mounted: function () {
            this.$nextTick(function () {
                this.view();
            })
        },
        filters:{
            //如果是折扣型那么用50%如果是价格型那么直接系显示
            set_coupon_gwqmoneystart:function(val,price_type){
                return price_type == 1 ? Math.floor(val*10000)/100+'%':val;
            }
        },
        methods: {
            view: function () {
                var _this = this;
                API('card/info', {key_id: id}, function (d) {
                    if (d.status == 200) {
                        _this.card = d.data.card;
                        var time = Vue.filter('time');
                        _this.card.card_validdate = time(_this.card.card_validdate,'yyyy-MM-dd');
                        _this.level = d.data.level;
                        _this.get_total();
                    } else {
                        layer.msg(d.statusText);
                    }
                })
            },
            get_total: function () {
                _this = this;
                _this.level.forEach(function (v) {
                    _this.total_balance += parseFloat(v.account_balance);
                    _this.total_amount += parseFloat(v.payable_amount);
                });
            },
            //账户历史
            account_history_btn: function () {
                var _this = this;
                if(_this.account_history.length<=0){
                    API('card/account_history', {key_id: id}, function (d) {
                        if (d.status == 200) {
                            _this.account_history = d.data;
                        } else {
                            layer.msg(d.statusText);
                        }
                    });
                }
            },
            //单品
            product_btn: function () {
                var _this = this;
                if(_this.product.length<=0){
                    API('card/product', {key_id: id}, function (d) {
                        if (d.status == 200) {
                            _this.product = d.data;
                        } else {
                            layer.msg(d.statusText);
                        }
                    });
                }
            },
            //项目
            project_btn: function () {
                var _this = this;
                if(_this.project.length<=0){
                    API('card/project', {key_id: id}, function (d) {
                        if (d.status == 200) {
                            _this.project = d.data;
                        } else {
                            layer.msg(d.statusText);
                        }
                    });
                }
            },
            //套餐
            package_btn: function () {
                var _this = this;
                if(_this.packages.length<=0){
                    API('card/package', {key_id: id}, function (d) {
                        if (d.status == 200) {
                            _this.packages = d.data;
                        } else {
                            layer.msg(d.statusText);
                        }
                    });
                }
            },
            //项目消耗历史
            project_use_btn: function () {
                var _this = this;
                if(_this.project_use.length<=0){
                    API('card/project_use', {key_id: id}, function (d) {
                        if (d.status == 200) {
                            _this.project_use = d.data.data;
                            _this.com_code = d.data.com_code;
                            _this.product_unit_price = d.data.product_unit_price;
                        } else {
                            layer.msg(d.statusText);
                        }
                    });
                }
            },
            //交易历史
            trade_history_btn: function () {
                var _this = this;
                if(_this.trade_history.length<=0){
                    API('card/trade_history', {key_id: id}, function (d) {
                        if (d.status == 200) {
                            _this.trade_history = d.data;
                        } else {
                            layer.msg(d.statusText);
                        }
                    });
                }
            },
            //异动历史
            change_history_btn: function () {
                var _this = this;
                if(_this.change_history.length<=0){
                    API('card/change_history', {key_id: id}, function (d) {
                        if (d.status == 200) {
                            _this.change_history = d.data;
                        } else {
                            layer.msg(d.statusText);
                        }
                    });
                }
            },
           integral_record_btn:function(){
               var _this = this;
               if(_this.trade_history.length<=0){
                   API('card/integral_record', {key_id: id}, function (d) {
                       if (d.status == 200) {
                           _this.integral_record = d.data;
                       } else {
                           layer.msg(d.statusText);
                       }
                   });
               }
            },
            //客户拥有的优惠券
            coupon_history_btn: function () {
                var _this = this;
                if(_this.coupon_history.length<=0){
                    API('card/get_coupon_code',{"str":id+'_'},function(d){
                        if(d.status == 200){
                            CouponAPI('getYHMUserAll.php',{AccountID:id,md5hash:d.data},function(d){
                                if(d.result == 200){
                                    _this.coupon_history = d.data;
                                }   else {
                                    layer.msg(d.message);
                                }
                            })
                        }else{
                            layer.msg(d.statusText);    
                        }
                    })

                }
            },
            //欠货历史
            storein_history_btn: function () {
                var _this = this;
                if(_this.storein_history.length<=0){
                    API('card/storeio_history', {key_id: id,type:4}, function (d) {
                        if (d.status == 200) {
                            _this.storein_history = d.data;
                        } else {
                            layer.msg(d.statusText);
                        }
                    });
                }
            },
            //领货历史
            storeout_history_btn: function () {
                var _this = this;
                if(_this.storeout_history.length<=0){
                    API('card/storeio_history', {key_id: id,type:3}, function (d) {
                        if (d.status == 200) {
                            _this.storeout_history = d.data;
                        } else {
                            layer.msg(d.statusText);
                        }
                    });
                }
            },
            ff:function(){
                alert(this.in_num)
            }
        }
    });
}