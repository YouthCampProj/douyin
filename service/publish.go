package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/auth"
	"github.com/YouthCampProj/douyin/pkg/fileprocess"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/YouthCampProj/douyin/pkg/uploader"
	"mime/multipart"
)

type PublishActionService struct {
	Token      string
	FileHeader *multipart.FileHeader
}

type PublishListService struct {
	Token string
}

func (p *PublishActionService) Publish() *serializer.PublishActionResponse {
	res := &serializer.PublishActionResponse{}
	user, err := auth.ParseToken(p.Token)
	if err != nil {
		return serializer.BuildPublishActionResponse(serializer.CodePublishTokenInvalid)
	}
	videoPath := "uploads/" + p.FileHeader.Filename
	videoURL, err := uploader.UploadToLocal(p.FileHeader, videoPath)
	if err != nil {
		res.Response = serializer.NewResponse(serializer.CodePublishUploadError, err.Error())
		return res
	}
	coverPath := "uploads/" + p.FileHeader.Filename + ".jpg"
	coverURL, err := fileprocess.GetCoverFromLocal(videoPath, coverPath)
	if err != nil {
		res.Response = serializer.NewResponse(serializer.CodePublishUploadError, err.Error())
		return res
	}
	video := model.NewVideo()
	video.AuthorID = user.ID
	video.PlayURL = videoURL
	video.CoverURL = coverURL
	err = video.Save()
	if err != nil {
		res.Response = serializer.NewResponse(serializer.CodePublishUploadError, err.Error())
		return res
	}
	res.Response = serializer.NewResponse(serializer.CodeSuccess, "")
	return res
}

func (p *PublishListService) GetPublishList() *serializer.PublishListResponse {
	res := &serializer.PublishListResponse{}
	user, err := auth.ParseToken(p.Token)
	if err != nil {
		return serializer.BuildPublishListResponse(serializer.CodePublishTokenInvalid)
	}
	videos, err := model.GetVideoByAuthorID(user.ID)
	if err != nil {
		res.Response = serializer.NewResponse(serializer.CodeGetPublishListError, err.Error())
	}
	author := serializer.BuildUserResponse(user, false)
	res.VideoList = make([]*serializer.Video, 0)
	for _, video := range videos {
		res.VideoList = append(res.VideoList, serializer.BuildVideoResponse(video, author, false))
	}
	return res
}
