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
	NextTime  uint64   `json:"next_time"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []*Video `json:"video_list"` // 视频列表
}

func BuildFeedResponse(code int, feedList []*model.VideoAuthorBundle, nextTime int64) *FeedResponse {
	res := &FeedResponse{}
	if code != CodeSuccess {
		res.Response = NewResponse(code, CodeFeedMessages[code])
		return res
	}
	res.Response = NewResponse(CodeSuccess, "")
	res.NextTime = uint64(nextTime)
	res.VideoList = make([]*Video, len(feedList))
	for i, feed := range feedList {
		res.VideoList[i] = BuildVideoResponse(feed)
	}
	return res
}
