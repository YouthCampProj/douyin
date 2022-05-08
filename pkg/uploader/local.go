package uploader

import (
	"github.com/YouthCampProj/douyin/pkg/config"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

// UploadToLocal 将文件上传到本地 返回指向该文件的URL
func UploadToLocal(fileHeader *multipart.FileHeader, path string) (string, error) {
	ctx := &gin.Context{}
	err := ctx.SaveUploadedFile(fileHeader, path)
	if err != nil {
		return "", err
	}
	return config.Conf.Site.Domain + "/" + path, nil
}
