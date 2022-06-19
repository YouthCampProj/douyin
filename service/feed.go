package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/auth"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"log"
	"time"
)

type FeedService struct {
	LatestTime int64
	Token      string
}

func (fs *FeedService) GetFeed() *serializer.FeedResponse {
	var (
		latestTime time.Time
		feedList   []*model.VideoAuthorBundle
		err        error
	)
	if fs.Token != "" {
		user, err := auth.ParseToken(fs.Token)
		if err != nil {
			return serializer.BuildFeedResponse(serializer.CodePublishTokenInvalid, nil, 0)
		}
		latestTime, feedList, err = model.GetFeedListByTime(time.UnixMilli(fs.LatestTime), user.ID)
		if len(feedList) == 0 {
			latestTime, feedList, err = model.GetFeedListByTime(time.Now(), user.ID)
		}
	} else {
		latestTime, feedList, err = model.GetFeedListByTime(time.UnixMilli(fs.LatestTime))
		if len(feedList) == 0 {
			latestTime, feedList, err = model.GetFeedListByTime(time.Now())
		}
	}
	if err != nil {
		log.Println(err)
		return serializer.BuildFeedResponse(serializer.CodeFailedGetFeed, nil, time.Now().UnixMilli())
	}
	return serializer.BuildFeedResponse(serializer.CodeSuccess, feedList, latestTime.UnixMilli())
}
