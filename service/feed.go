package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"log"
	"time"
)

type FeedService struct {
	LatestTime int64
}

func (fs *FeedService) GetFeed() *serializer.FeedResponse {
	latestTime, feedList, err := model.GetFeedListByTime(time.UnixMilli(fs.LatestTime))
	if err != nil {
		log.Println(err)
		return serializer.BuildFeedResponse(serializer.CodeFailedGetFeed, nil, time.Now().UnixMilli())
	}
	return serializer.BuildFeedResponse(serializer.CodeSuccess, feedList, latestTime.UnixMilli())
}
