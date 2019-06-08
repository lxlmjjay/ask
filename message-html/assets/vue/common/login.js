new Vue({
    el: '#app',
    data: {
        user : '',
        upass : '',
        ucode : '',
        captcha_img : '',
        captcha_code : ''
    },
    mounted: function (){
            this.$nextTick(function () {
                this.logined();
            })
    },
    methods: {
        logined : function(){
            if (cookie.get('jwt_report') != null){
                window.location.href ='/main.html';
            }
            this.captcha();
        },
        login : function(){
            var _this = this, _data;

            if (_this.user == '' || _this.upass == '' || _this.ucode == '')
            {
                layer.msg('请输入用户名、密码和验证码');
                return false;
            }

            _data = {   user:_this.user, 
                        upass:_this.upass, 
                        ucode:_this.ucode, 
                        captcha_code:_this.captcha_code
                    }

            API('/login/auth', _data, function(d){
                if (d.status == 200)
                {
                    var re = d.data;
                    if (re){
                        cookie.set('jwt_report', re.jwt, 10);
                        cookie.set('user_report', re.user, 10.2);
                        //cookie.set('menu', JSON.stringify(re.menu), 2);
                        //localStorage.setItem('menu_report', JSON.stringify(re.menu));

                        window.location.href ='/main.html';
                    }                    
                }
                else{
                    layer.msg(d.statusText);
                }
            },false);
            return false;
        },
        captcha : function(){
            var _this = this;
            API('/login/captcha', {}, function(d){
                if (d.status == 200)
                {
                    var re = d.data;
                    if (re){
                        _this.captcha_img = re.img_html;
                        _this.captcha_code = re.jwt_captcha;
                    }
                }
            },false);
        }
    }
})
