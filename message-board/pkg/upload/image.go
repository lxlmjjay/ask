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
	"os"
	"path"
	"strings"
)

//func GetImageFullUrl(name string) string {
//	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
//}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

//func GetImagePath() string {
//	return setting.AppSetting.ImageSavePath + time.Now().Format("20060102") + "/"
//}

//func GetImageFullPath() string {
//	return setting.AppSetting.RuntimeRootPath + GetImagePath()
//}

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

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
}

	return nil
}

func GetImageUrl(c *gin.Context) (url string, err error) {
	//data := make(map[string]string)
	_, image, err := c.Request.FormFile("image")
	if err != nil || image == nil{
		return
	}
	imageName := GetImageName(image.Filename)
	savePath := setting.AppSetting.ImageSavePath + imageName
	//imageName := GetImageName(image.Filename)
	//fullPath := GetImageFullPath()
	//savePath := GetImagePath()

	//src := fullPath + imageName
	//if !CheckImageExt(imageName){
	//	err = fmt.Errorf("image wrong format")
	//	return
	//	}
	//if !CheckImageSize(file) {
	//	err = fmt.Errorf("image max size is %vM", setting.AppSetting.ImageMaxSize)
	//	return
	//	}
	//
	//err = upload.CheckImage(fullPath)
	//if err != nil {
	//	logs.Warn(err)
	//	code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL}
	err = c.SaveUploadedFile(image, savePath)
	if err != nil {
		logs.Warn(err)
		return
		}
	//data["image_url"] = upload.GetImageFullUrl(imageName)
	//data["image_save_url"] = savePath + imageName
	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": data,
	//})
	url = setting.AppSetting.ImagePrefixUrl + "api/v1/images/" + imageName
	return
}
