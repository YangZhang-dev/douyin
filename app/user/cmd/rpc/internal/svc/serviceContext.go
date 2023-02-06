package svc

import (
	"douyin/app/user/cmd/rpc/internal/config"
	"douyin/app/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	UserModel   model.UserModel
	FollowModel model.FollowModel
	ChatModel   model.ChatModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserModel:   model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		FollowModel: model.NewFollowModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ChatModel:   model.NewChatModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
