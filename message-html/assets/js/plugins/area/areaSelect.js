
var city_json = '';
function begin(p_data ,c_data,d_data){
  p_data = p_data == ''?0:p_data;
  c_data = c_data == ''?0:c_data;
  d_data = d_data == ''?0:d_data;
  $.getJSON('/assets/js/plugins/area/areas.js',function(json){
    city_json = json;
    var province = $('.province-select');
    var ciy = $('.city-select');
    var district = $('.district-select');
    area_select(0,province,p_data,'选择省');
    if(typeof(p_data) != 'undefined' && p_data >0){
      area_select(p_data,ciy,c_data,'选择市');
    } 
    if(typeof(c_data) != 'undefined' && c_data >0){
      area_select(c_data,district,d_data,'选择区/县');
    } 
  });
}

//pid:父id，e当前区域选择对象，value 当前的select是否有选中的默认值 0为无
function area_select(pid,e,value,text){
  //alert(city_json);
  var option = city_json[pid];
  var length = option.length;
  if (length){ 
    var option_obj = '',selected='';
    for(var i=0 ; i < length ;i++){
        if(value==option[i]['id']){
            selected = 'selected';
        }else{
            selected= '';
        }
      option_obj += '<option value="'+option[i]['id']+'"  '+selected+'>'+option[i]['name']+'</option>';
    }
    if(e.children().length >=1)
       e.empty();
    e.append('<option value="0" >'+text+'</option>'+option_obj);
    // if(value >=0)
    //   e.val(value);
  }
}

$(document).on('change','.province-select',function(){

  var _this = $(this);
  var city =  $('.city-select');
  var district = $('.district-select')
  if(_this.val() > 0){
     area_select(_this.val(),city,0,'选择市');
  }else{
    city.empty();
    city.append('<option value="0">选择市</>');
  }
  district.empty();
  district.append('<option value="0">选择区/县</>');
  app.city = 0;
  app.district = 0;
})

$(document).on('change','.city-select',function(){

  var _this = $(this);
  var district = $('.district-select');
  if(_this.val() > 0){
     area_select(_this.val(),district,0,'选择区/县');
  }else{
    district.empty();
    district.append('<option value="0">选择区/县</>');
  }
  app.district = 0;
})


