package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type PublishActionService struct {
	User       *model.User
	FileHeader *multipart.FileHeader
}

func (p *PublishActionService) Publish() *serializer.PublishActionResponse {
	res := &serializer.PublishActionResponse{}
	ctx := &gin.Context{}
	err := ctx.SaveUploadedFile(p.FileHeader, "./uploads/"+p.FileHeader.Filename)
	if err != nil {
		res.Response = serializer.NewResponse(serializer.CodePublishUploadError, err.Error())
		return res
	}
	res.Response = serializer.NewResponse(serializer.CodeSuccess, "")
	return res
	// TODO: 在数据库中创建该视频对应的记录
}
