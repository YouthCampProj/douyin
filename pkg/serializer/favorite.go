package serializer

// FavoriteActionResponse 点赞操作响应
type FavoriteActionResponse struct {
	Response
}

// FavoriteListResponse 点赞列表响应
type FavoriteListResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
}
