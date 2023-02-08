// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id            int64  `json:"id"`
	Username      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type Status struct {
	StatusCode uint32 `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type RegisterReq struct {
	Username string `form:"username" validate:"required,max=32,min=1"`
	Password string `form:"password" validate:"required,max=32,min=1"`
}

type RegisterResp struct {
	Status
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type LoginReq struct {
	Username string `form:"username" validate:"required,max=32,min=1"`
	Password string `form:"password" validate:"required,max=32,min=1"`
}

type LoginResp struct {
	Status
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserInfoReq struct {
	UserId int64  `form:"user_id" validate:"required,gte=0"`
	Token  string `form:"token,optional"`
}

type UserInfoResp struct {
	Status
	User User `json:"user,omitempty"`
}

type FollowOrUnfollowReq struct {
	Token      string `form:"token" validate:"required"`
	ToUserId   int64  `form:"to_user_id" validate:"required,gte=0"`
	ActionType int64  `form:"action_type" validate:"required,gte=1,lte=2"`
}

type FollowOrUnfollowResp struct {
	Status
}

type FollowListReq struct {
	Token  string `form:"token" validate:"required"`
	UserId int64  `form:"user_id" validate:"required,gte=0"`
}

type FollowListResp struct {
	Status
	UserList []User `json:"user_list,omitempty"`
}

type FansListReq struct {
	Token  string `form:"token" validate:"required"`
	UserId int64  `form:"user_id" validate:"required,gte=0"`
}

type FansListResp struct {
	Status
	UserList []User `json:"user_list,omitempty"`
}

type FriendsListReq struct {
	Token  string `form:"token" validate:"required"`
	UserId int64  `form:"user_id" validate:"required,gte=0"`
}

type FriendsListResp struct {
	Status
	UserList []User `json:"user_list,omitempty"`
}

type Message struct {
	Id         int64  `json:"id"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

type SendMessageReq struct {
	Token      string `form:"token" validate:"required"`
	ToUserId   int64  `form:"to_user_id" validate:"required,gte=0"`
	ActionType int64  `form:"action_type" validate:"required,gte=1,lte=2"`
	Content    string `form:"content" validate:"required"`
}

type SendMessageResp struct {
	Status
}

type HistoryMessageReq struct {
	Token    string `form:"token" validate:"required"`
	ToUserId int64  `form:"to_user_id" validate:"required,gte=0"`
}

type HistoryMessageResp struct {
	Status
	Message []Message `json:"message_list,omitempty"`
}
