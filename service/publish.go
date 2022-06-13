package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/auth"
	"github.com/YouthCampProj/douyin/pkg/config"
	"github.com/YouthCampProj/douyin/pkg/fileprocess"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/YouthCampProj/douyin/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

type PublishActionService struct {
	Token      string
	FileHeader *multipart.FileHeader
	Title      string
}

type PublishListService struct {
	Token  string
	UserID uint64
}

func (p *PublishActionService) Publish() *serializer.PublishActionResponse {
	res := &serializer.PublishActionResponse{}
	user, err := auth.ParseToken(p.Token)
	if err != nil {
		return serializer.BuildPublishActionResponse(serializer.CodePublishTokenInvalid)
	}
	videoPath := "uploads/" + strconv.FormatUint(user.ID, 10) + "/" + time.Now().Format("2006-01-02")
	err = os.MkdirAll(videoPath, 0755)
	if err != nil && !os.IsExist(err) {
		res.Response = serializer.NewResponse(serializer.CodePublishUploadError, err.Error())
		return res
	}
	ctx := &gin.Context{}
	err = ctx.SaveUploadedFile(p.FileHeader, videoPath+"/"+p.FileHeader.Filename)
	if err != nil {
		res.Response = serializer.NewResponse(serializer.CodePublishUploadError, err.Error())
		return res
	}
	videoPath += "/" + p.FileHeader.Filename
	videoURL := config.Conf.Site.Domain + videoPath
	coverPath := videoPath + ".jpg"
	coverURL, err := fileprocess.GetCoverFromLocal(videoPath, coverPath)
	if err != nil {
		res.Response = serializer.NewResponse(serializer.CodePublishUploadError, err.Error())
		return res
	}
	video := model.NewVideo()
	// 生成视频ID
	video.ID = uint64(snowflake.GenID())
	video.AuthorID = user.ID
	video.PlayURL = videoURL
	video.CoverURL = coverURL
	video.Title = p.Title
	err = video.Save()
	if err != nil {
		res.Response = serializer.NewResponse(serializer.CodePublishUploadError, err.Error())
		return res
	}
	res.Response = serializer.NewResponse(serializer.CodeSuccess, "")
	return res
}

func (p *PublishListService) GetPublishList() *serializer.PublishListResponse {
	_, err := auth.ParseToken(p.Token)
	if err != nil {
		return serializer.BuildPublishListResponse(serializer.CodePublishTokenInvalid, nil)
	}

	videoAuthorBundle, err := model.GetPublishListByAuthorID(p.UserID)
	if err != nil {
		return serializer.BuildPublishListResponse(serializer.CodeGetPublishListError, nil, err.Error())
	}
	res := serializer.BuildPublishListResponse(serializer.CodeSuccess, videoAuthorBundle)
	return res
}
