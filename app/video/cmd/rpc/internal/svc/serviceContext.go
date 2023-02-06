package svc

import (
	"douyin/app/video/cmd/rpc/internal/config"
	"douyin/app/video/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	VideoModel    model.VideoModel
	FavoriteModel model.FavoriteVideoModel
	CommentModel  model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		VideoModel:    model.NewVideoModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		FavoriteModel: model.NewFavoriteVideoModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		CommentModel:  model.NewCommentModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
