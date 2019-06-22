package upload

import (
"fmt"
"github.com/astaxie/beego/logs"
"github.com/gin-gonic/gin"
"log"
"message-board/pkg/file"
"message-board/pkg/setting"
"message-board/pkg/util"
"mime/multipart"
"net/http"
"path"
"strings"
)

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logs.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

func GetImageUrl(c *gin.Context) (url string, code int, err error, msg string) {
	file, image, _:= c.Request.FormFile("image")
	if image == nil{
		return
	}
	imageName := GetImageName(image.Filename)
	savePath := setting.AppSetting.ImageSavePath + imageName

	if !CheckImageExt(imageName){
		return "", http.StatusBadRequest, fmt.Errorf("图片格式错误"), fmt.Sprintf("允许图片格式为:%s", setting.AppSetting.ImageAllowExts)
	}
	if !CheckImageSize(file) {
		return "", http.StatusBadRequest, fmt.Errorf("图片尺寸过大"), fmt.Sprintf("图片尺寸不能超过:%vM", setting.AppSetting.ImageMaxSize)
	}

	err = c.SaveUploadedFile(image, savePath)
	if err != nil {
		logs.Warn(err)
		return "", http.StatusInternalServerError, fmt.Errorf("服务器处理时发生了意外"),"请稍后再试"
	}

	url = setting.AppSetting.ImagePrefixUrl + "api/v1/images/" + imageName
	return
}
