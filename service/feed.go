package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"log"
)

type FeedService struct {
	LatestTime int64
}

func (fs *FeedService) GetFeed() *serializer.FeedResponse {
	var feedList []*serializer.FeedResponseBuilder
	videoList, err := model.GetVideoByTime(fs.LatestTime)
	if err != nil {
		log.Println(err)
		return serializer.BuildFeedResponse(serializer.CodeFailedGetFeed, nil, 0)
	}
	if len(videoList) == 0 {
		return serializer.BuildFeedResponse(serializer.CodeNoMoreFeed, nil, 0)
	}
	for i := range videoList {
		author, err := model.GetUserByID(videoList[i].AuthorID)
		if err != nil {
			continue
		}
		feed := &serializer.FeedResponseBuilder{
			Video:  videoList[i],
			Author: author,
		}
		feedList = append(feedList, feed)

	}
	latestTime := videoList[len(videoList)-1].CreatedAt.Unix()
	return serializer.BuildFeedResponse(serializer.CodeSuccess, feedList, latestTime)
}
