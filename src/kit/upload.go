package kit

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/Scfy-Code/scfy-im/app"
)

// Upload 将上传的文件保存在指定的目录中并返回文件id
func Upload(file multipart.File, fileHeader *multipart.FileHeader) string {
	defer file.Close()
	fileName := fileHeader.Filename
	suffixName := fileName[strings.LastIndex(fileName, ".")-1:]
	filePath := app.UploadDir + time.Now().Format("2006/01/02/15") + "/"
	err0 := os.MkdirAll(filePath, os.ModePerm)
	if err0 != nil {
		app.WarnLogger.Printf("新建上传目录出错！错误信息：%s", err0.Error())
	}
	rFile, err1 := os.Create(filePath + fmt.Sprintf("%d", time.Now().UnixNano()) + suffixName)
	defer rFile.Close()
	if err1 != nil {
		app.WarnLogger.Printf("新建文件失败！失败原因：%s", err1.Error())
	}
	io.Copy(rFile, file)
	return filePath + rFile.Name()
}
