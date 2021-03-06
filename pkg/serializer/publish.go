package serializer

import "github.com/YouthCampProj/douyin/model"

const (
	CodePublishTokenInvalid = 3000 + iota
	CodePublishFileSizeError
	CodePublishFileInvalid
	CodePublishNameInvalid
	CodePublishUploadError
	CodeGetPublishListError
)

var CodePublishMessages = map[int]string{
	CodePublishTokenInvalid:  "publish token invalid",
	CodePublishFileSizeError: "file size error",
	CodePublishFileInvalid:   "file invalid",
	CodePublishNameInvalid:   "publish name invalid",
	CodePublishUploadError:   "upload error",
}

// PublishActionResponse 投稿操作响应
type PublishActionResponse struct {
	Response
}

// PublishListResponse 投稿列表响应
type PublishListResponse struct {
	Response
	VideoList []*Video `json:"video_list"` // 用户发布的视频列表
}

func BuildPublishActionResponse(code int) *PublishActionResponse {
	res := &PublishActionResponse{}
	res.Response = NewResponse(code, CodePublishMessages[code])
	return res
}

func BuildPublishListResponse(code int, bundle []*model.VideoAuthorBundle, msg ...string) *PublishListResponse {
	res := &PublishListResponse{}
	res.Response = NewResponse(code, CodePublishMessages[code])
	res.VideoList = make([]*Video, len(bundle))
	if code != CodeSuccess {
		if len(msg) > 0 {
			res.StatusMsg = msg[0]
		}
		return res
	}
	for i, v := range bundle {
		res.VideoList[i] = BuildVideoResponse(v)
	}
	return res

}
