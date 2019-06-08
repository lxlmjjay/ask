var tree_style = {
    data: {
        simpleData: {
            enable: true
        }
    },
    check: {
        enable: true,
        // chkStyle: "radio",
        // radioType: "all"
    }
},modal_id = 'companyLevel' ,ul_id = 'treeDemo';
var companyLevel = {
	template:'<div class="modal" :id="modal_id" tabindex="-1" role="dialog" aria-hidden="true">\
                <div class="modal-dialog">\
                    <div class="modal-content animated bounceInRight">\
                        <div class="modal-body">\
                            <p>&nbsp;</p>\
                            <div class="ibox-tools">\
                                <form class="form-inline text-right">\
                                    <div class="form-group form-group-sm"><input type="text" v-model.trim="com_keyword" placeholder="门店名称" class="form-control"></div>\
                                    <div class="input-group">\
                                        <button type="button" v-on:click="com_search" class="btn btn-sm btn-primary">搜索</button>&nbsp;\
                                        <button type="button" v-on:click="com_old" class="btn btn-sm btn-primary">还原</button>\
                                    </div>\
                                </form>\
                            </div>\
                            <p>&nbsp;</p>\
                            <form class="form-horizontal">\
                                <div class="form-group clearfix">\
                                    <ul :id="ul_id" class="ztree" style="width:100%"></ul>\
                                </div>\
                            </form>\
                        </div>\
                        <div class="modal-footer">\
                            </p><button type="button" class="btn btn-white" data-dismiss="modal">取消</button>\
                            <button type="button" class="btn btn-primary" v-on:click="get_com_save">确定</button>\
                        </div>\
                    </div>\
                </div>\
            </div>',
    props:['comInfo','func','pModalId','pUlId'],
    data:function(){
    	return {
    		com_data:[],
    		com_keyword:'',
    		com_info:this.comInfo,
            modal_id: this.pModalId ? this.pModalId : modal_id,
            ul_id: this.pUlId ? this.pUlId : ul_id,
    	}
    },
    methods:{
        //获取数据
    	index: function () {
            var _this = this;
            API(_this.func,{},function(d){
                if (d.status == 200)
                {
                    _this.com_data = d.data;
                    $.fn.zTree.init($("#"+_this.ul_id), tree_style,_this.com_data);
                    var zTree = $.fn.zTree.getZTreeObj(_this.ul_id);
                    zTree.setting.check.chkboxType = { "Y": "s", "N": "s" }
                }else{
                    layer.msg(d.statusText);
                }
            })
            $('#'+_this.modal_id).modal('show');
        },
        //搜索
        com_search:function(){
            var _this = this;
            if(_this.com_keyword==''){
                layer.msg('请输入门店名称');
                return ;
            }
            var treeObj = $.fn.zTree.getZTreeObj(_this.ul_id),
                nodes = treeObj.getNodesByParamFuzzy("name", _this.com_keyword, null),
                temp = [];

            if(nodes.length > 0){
                for(i in nodes){
                    temp.push({'name':nodes[i]['name'],'open':nodes[i]['open'],'pId':nodes[i]['pId'],'id':nodes[i]['id']});
                }
                $.fn.zTree.init($("#"+_this.ul_id), tree_style, temp);
            }
            
        },
        //还原
        com_old:function(){
            var _this = this;
            $.fn.zTree.init($("#"+_this.ul_id), tree_style,_this.com_data );
        },
        //保存
        get_com_save: function () {
            _this = this;
            var treeObj = $.fn.zTree.getZTreeObj(_this.ul_id),
                nodes = treeObj.getCheckedNodes(true),
                temp = [],
                com_name = '默认当前门店（包括下级）';
                nodes_length = nodes.length;
            if(nodes_length > 0){
                com_name = '已选择门店数：'+nodes_length
                for(i in nodes){
                    temp.push(nodes[i]['id']);
                }  
            }

            _this.com_info.com_name = com_name;

            _this.com_info.com_key_id = temp
            $('#'+_this.modal_id).modal('hide');
        },
    }
}