package serializer

import "github.com/YouthCampProj/douyin/model"

const (
	CodeNoMoreFeed = iota + 2000
	CodeFailedGetFeed
)

var CodeFeedMessages = map[int]string{
	CodeNoMoreFeed:    "feed is empty",
	CodeFailedGetFeed: "failed to get feed",
}

// FeedResponse 视频流响应
type FeedResponse struct {
	Response
	NextTime  uint64   `json:"next_time,omitempty"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []*Video `json:"video_list,omitempty"` // 视频列表
}

type FeedResponseBuilder struct {
	Video  *model.Video
	Author *model.User
}

func BuildFeedResponse(code int, feedList []*FeedResponseBuilder, nextTime int64) *FeedResponse {
	res := &FeedResponse{}
	if code != CodeSuccess {
		res.Response = NewResponse(code, CodeFeedMessages[code])
		return res
	}
	res.Response = NewResponse(CodeSuccess, "")
	res.NextTime = uint64(nextTime)
	for _, feed := range feedList {
		authorResponse := BuildUserResponse(feed.Author, false)
		res.VideoList = append(res.VideoList, BuildVideoResponse(feed.Video, authorResponse, false))
	}
	return res
}
