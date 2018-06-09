package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/google/uuid"
	"time"
	"fmt"
	"path/filepath"
	"github.com/astaxie/beego"
	"os"
)

type FileUploadController struct {
	beego.Controller
}

//定义Kindeditor特殊的返回json。
type UploadMessage struct {
	Error   int64  `json:"error"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//https://github.com/google/uuid
//
func (c *FileUploadController) Upload() {
	fileNameKey := "imgFile"
	uploadDir := beego.AppConfig.String("uploadDir")
	uploadBase := beego.AppConfig.String("uploadBase")
	file, header, err := c.GetFile(fileNameKey) // where <<this>> is the controller and <<file>> the id of your form field
	defer c.ServeJSON()
	if file != nil {
		// get the filename
		fileName := header.Filename
		fileNameUuid, _ := uuid.NewUUID()
		fileNameExt := filepath.Ext(fileName)
		t := time.Now()

		pathDir := fmt.Sprintf("%d/%d-%02d", t.Year(), t.Year(), t.Month())
		logs.Info("########### pathDir : %v ", uploadDir+pathDir)
		if !Exists(uploadDir + pathDir) {
			// 创建文件夹
			os.Mkdir(uploadDir+pathDir, os.ModePerm)
		}
		pathUrl := fmt.Sprintf("%s/%s%s", pathDir, fileNameUuid.String(), fileNameExt)
		logs.Info("########### file : %v baseDir : %v pathUrl : %v", fileName, pathDir, pathUrl)
		// save to server
		err := c.SaveToFile(fileNameKey, uploadDir+pathUrl)
		if err != nil {
			c.Data["json"] = &UploadMessage{Error: 1, Message: err.Error()}
		} else {
			logs.Info("########### upload url : %v", uploadBase+pathUrl)
			c.Data["json"] = &UploadMessage{Error: 0, Message: "", Url: uploadBase + pathUrl}
		}

	}
	logs.Info("########### file %v eror: %v", file, err)
}
