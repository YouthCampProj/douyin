package serializer

import "github.com/YouthCampProj/douyin/model"

const (
	CodeFavoriteLikeFailed = 4000 + iota
	CodeFavoriteUnLikeFailed
	FavoriteActionFailed
	CodeFavoriteTokenInvalid
	CodeFavoriteUserIDInvalid
	CodeFavoriteVideoIDInvalid
	CodeFavoriteActionTypeInvalid
	CodeFavoriteExists
	CodeFavoriteDBFailed
	CodeFavoriteNotExists
	CodeFavoriteListDBFailed
)

var CodeFavoriteMessage = map[int]string{
	CodeFavoriteLikeFailed:        "点赞失败",
	CodeFavoriteUnLikeFailed:      "取消点赞失败",
	FavoriteActionFailed:          "操作失败",
	CodeFavoriteTokenInvalid:      "Token无效",
	CodeFavoriteUserIDInvalid:     "用户ID无效",
	CodeFavoriteVideoIDInvalid:    "视频ID无效",
	CodeFavoriteActionTypeInvalid: "操作类型无效",
	CodeFavoriteExists:            "已经点赞啦",
	CodeFavoriteDBFailed:          "数据库操作失败",
	CodeFavoriteNotExists:         "还没有点赞呢",
	CodeFavoriteListDBFailed:      "数据库操作失败",
}

// FavoriteActionResponse 点赞操作响应
type FavoriteActionResponse struct {
	Response
}

// FavoriteListResponse 点赞列表响应
type FavoriteListResponse struct {
	Response
	VideoList []*Video `json:"video_list"`
}

func BuildFavoriteActionResponse(code int) *FavoriteActionResponse {
	res := &FavoriteActionResponse{}
	res.Response = NewResponse(code, CodeFavoriteMessage[code])
	return res
}

func BuildFavoriteListResponse(code int, videoList []*model.VideoAuthorBundle) *FavoriteListResponse {
	res := &FavoriteListResponse{}
	res.Response = NewResponse(code, CodeFavoriteMessage[code])
	res.VideoList = make([]*Video, len(videoList))
	if code != CodeSuccess {
		return res
	}
	for i, v := range videoList {
		res.VideoList[i] = BuildVideoResponse(v)
	}
	return res

}
