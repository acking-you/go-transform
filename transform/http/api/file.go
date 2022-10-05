package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"logger/lg"
	"path/filepath"
	"transform/http/util"
)

func PostFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		util.BadListResponse(c, "客户端传输出错")
		return
	}
	files := form.File["files"]
	for _, file := range files {
		dst := filepath.Join("../static", file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			lg.L.Errorf("文件上传出错: %v", err)
			util.BadResponse(c, fmt.Sprintf("%s上传失败", file.Filename))
			return
		}
		util.GoodResponse(c, fmt.Sprintf("%s上传成功", file.Filename))
	}
}

func GetFiles(c *gin.Context) {
	files := getFilesName()
	if files == nil {
		util.BadListResponse(c, "获取文件出错，或者服务端文件为空")
		return
	}

	util.GoodListResponse(c, files)
}

func getFilesName() []string {
	root := "../static"
	fs, err := ioutil.ReadDir(root)
	if err != nil {
		lg.L.Error("读取文件路径出错")
		return nil
	}
	var ret []string
	for _, f := range fs {
		ret = append(ret, f.Name())
	}
	return ret
}
