package serializer

// RelationActionResponse 关注操作响应
type RelationActionResponse struct {
	Response
}

// RelationFollowListResponse 获取关注列表响应
type RelationFollowListResponse struct {
	Response
	UserList []User `json:"user_list,omitempty"` // 用户信息列表
}

// RelationFollowerListResponse 获取粉丝列表响应
type RelationFollowerListResponse struct {
	Response
	UserList []User `json:"user_list,omitempty"` // 用户列表
}
