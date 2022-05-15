package serializer

import "github.com/YouthCampProj/douyin/model"

const (
	CodeRelationActionTypeInvalid = 7000 + iota
	CodeRelationTokenInvalid
	CodeRelationAlreadyFollow
	CodeRelationNotFollow
	CodeRelationDBError
)

var CodeRelationMessage = map[int]string{
	CodeRelationActionTypeInvalid: "invalid action type",
	CodeRelationTokenInvalid:      "invalid token",
	CodeRelationAlreadyFollow:     "already follow",
	CodeRelationNotFollow:         "not follow",
	CodeRelationDBError:           "db error",
}

// RelationActionResponse 关注操作响应
type RelationActionResponse struct {
	Response
}

// RelationFollowListResponse 获取关注列表响应
type RelationFollowListResponse struct {
	Response
	UserList []*User `json:"user_list,omitempty"` // 用户信息列表
}

// RelationFollowerListResponse 获取粉丝列表响应
type RelationFollowerListResponse struct {
	Response
	UserList []*User `json:"user_list,omitempty"` // 用户列表
}

func BuildRelationActionResponse(code int) *RelationActionResponse {
	res := &RelationActionResponse{}
	res.Response = NewResponse(code, CodeRelationMessage[code])
	return res
}

func BuildRelationFollowListResponse(code int, userList []*model.UserAPI) *RelationFollowListResponse {
	res := &RelationFollowListResponse{}
	res.Response = NewResponse(code, CodeRelationMessage[code])
	res.UserList = make([]*User, len(userList))
	if code != CodeSuccess {
		return res
	}
	for i, u := range userList {
		res.UserList[i] = BuildUserResponse(u)
	}
	return res
}
