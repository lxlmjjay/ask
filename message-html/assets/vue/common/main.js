new Vue({
    el: '#wrapper',
    data: {
        tx:'assets/img/head.jpg',
        uname: '我的姓名',
        jobs : 'Goer',
        status_commission:true,
        menu:[]
    },
    mounted: function (){
            this.$nextTick(function () {
                this.page_menu();
            })
    },
    methods: {
        page_menu : function(){
            if (cookie.get('jwt_report') == null){
                //window.location.href ='/index.html';
                return false;
            }

            
            this.uname = cookie.data('name');
            this.jobs  = cookie.data('jobs');
            this.tx  =  baseURL + '/pics'+cookie.data('header');
            /*
            //首页根据用户角色显示哪个页面
            var win = document.getElementById('dfm_win');
            switch (cookie.data('group')){
                case '001':
                    win.src = '/basic/storefront.html';
                    break;
                case '002':
                    win.src = '/basic/management.html';
                    break;
                case '010'://财务
                case '006'://前台助理
                    win.src = '/cash/users.html?do=list';
                    break;
                case '007'://人事
                    win.src = '/hr/employee.html';
                    break;
                case '008'://企划
                    win.src = '/member/customer.html';
                    break;
                case '009'://仓管
                    win.src = '/inventory/product_cost_amount.html';
                    break;
                default:
                    win.src = '/home/info.html';                    
            }

            //提成明细设置
            if(cookie.data('group') == '002' || cookie.data('group') == '001')
              this.status_commission = false
            this.menu = JSON.parse(cookie.get('menu'));            
            this.menu = JSON.parse(localStorage.getItem('menu'));
            */
        },
        a_logout : function(){
            cookie.delete('jwt_report');
            cookie.delete('user_report');
            //cookie.delete('header');
            //cookie.delete('menu');
            //localStorage.removeItem('menu_report');
            window.location.href = '/index.html';
        }
    }
})

/***********************************************************************/
//修改密码
function edit_pass (){
    layer.open({
        type: 2,
        title: '修改密码',
        skin: 'layui-layer-rim', //加上边框
        area: ['500px', '350px'], //宽高
        content: '/hr/my_employee_editpass.html?do=edit_pass'
    });
    $(document).click();
}
