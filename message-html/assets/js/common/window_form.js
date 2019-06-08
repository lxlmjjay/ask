function home_win(w){
    var title,url;
        switch(w){
            case 1:
                title = '工作日志';
                url = '/home/plan.html';
                wh = ['600px', '410px'];
                break;
            case 2:
                title = '心得分享';
                url = '/home/share.html';
                wh = ['600px', '410px'];
                break;
            case 3:
                title = '客户评价';
                url = '/home/comment.html?do=set';
                wh = ['600px', '450px'];
                break;
        }
    layer.open({
      type: 2,
      title: title,
      skin: 'layui-layer-rim', //加上边框
      area: wh, //宽高
      content: url
    });
    return false;
}

function hr_service(w,id){
    var title,url,wh;
    if (w ==1){
        title = '考勤';
        if(id == 0){
          url = '/home/hr_service_database.html?do=add';
        }else{
          url = '/home/hr_service_database.html?do=edit&id='+id;
        }
        wh = ['620px', '300px'];
    }else{
        title = '请休假';
        if(id == 0){
          url = '/home/hr_service_timeoff.html?do=add'; 
        }else{
          url = '/home/hr_service_timeoff.html?do=edit&id='+id;
        }      
        wh = ['600px', '430px'];
    }

    layer.open({
      type: 2,
      title: title,
      skin: 'layui-layer-rim', //加上边框
      area: wh, //宽高
      content: url
    });
    return false;
}



// layer.config({
//     extend: ["extend/layer.ext.js", "skin/moon/style.css"],
//     skin: "layer-ext-moon"
// }),
// layer.ready(function() {
//     function e() {
//         parent.layer.open({
//             title: "初见倾心，再见动情",
//             type: 1,
//             area: ["700px", "auto"],
//             content: t,
//             btn: ["确定", "取消"]
//         })
//     }
//     var t = $("#welcome-template").html();
//     $("a.viewlog").click(function() {
//         return e(),
//         !1
//     }),
//     $("#pay-qrcode").click(function() {
//         var e = $(this).html();
//         parent.layer.open({
//             title: !1,
//             type: 1,
//             closeBtn: !1,
//             shadeClose: !0,
//             area: ["600px", "auto"],
//             content: e
//         })
//     }),
//     console.log("欢迎使用H+，如果您在使用的过程中有碰到问题，可以参考开发文档，感谢您的支持。")
// }); 