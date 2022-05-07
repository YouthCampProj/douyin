package serializer

// PublishActionResponse 投稿操作响应
type PublishActionResponse struct {
	Response
}

// PublishListResponse 投稿列表响应
type PublishListResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"` // 用户发布的视频列表
}
