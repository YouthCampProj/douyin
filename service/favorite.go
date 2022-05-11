package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"log"
)

type FavoriteActionService struct {
	UserID  uint64
	VideoID uint64
	Action  int32
}

type FavoriteListService struct {
	UserID uint64
}

func (s *FavoriteActionService) Execute() *serializer.FavoriteActionResponse {
	switch s.Action {
	case 1:
		return s.Like()
	case 2:
		return s.Unlike()
	}
	return serializer.BuildFavoriteActionResponse(serializer.FavoriteActionFailed)
}

func (s *FavoriteActionService) Like() *serializer.FavoriteActionResponse {
	code := model.AddFavorite(s.UserID, s.VideoID)
	switch code {
	case 0:
		return serializer.BuildFavoriteActionResponse(serializer.CodeSuccess)
	case 1:
		return serializer.BuildFavoriteActionResponse(serializer.CodeFavoriteExists)
	case 2:
		return serializer.BuildFavoriteActionResponse(serializer.CodeFavoriteDBFailed)
	}
	return serializer.BuildFavoriteActionResponse(serializer.CodeFavoriteLikeFailed)
}

func (s *FavoriteActionService) Unlike() *serializer.FavoriteActionResponse {
	code := model.DeleteFavorite(s.UserID, s.VideoID)
	switch code {
	case 0:
		return serializer.BuildFavoriteActionResponse(serializer.CodeSuccess)
	case 1:
		return serializer.BuildFavoriteActionResponse(serializer.CodeFavoriteNotExists)
	case 2:
		return serializer.BuildFavoriteActionResponse(serializer.CodeFavoriteDBFailed)
	}
	return serializer.BuildFavoriteActionResponse(serializer.CodeFavoriteUnLikeFailed)
}

// GetFavoriteList 获取点赞列表
func (s *FavoriteListService) GetFavoriteList() *serializer.FavoriteListResponse {
	// 这里假定已经判断过token有效性 故不再验证userID是否有效
	videoAuthorBundle, err := model.GetFavoriteVideoList(s.UserID)
	if err != nil {
		log.Println(err)
		return serializer.BuildFavoriteListResponse(serializer.CodeFavoriteListDBFailed, nil)
	}
	return serializer.BuildFavoriteListResponse(serializer.CodeSuccess, videoAuthorBundle)
}
