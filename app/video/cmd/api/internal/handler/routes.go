// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	comment "douyin/app/video/cmd/api/internal/handler/comment"
	favorite "douyin/app/video/cmd/api/internal/handler/favorite"
	video "douyin/app/video/cmd/api/internal/handler/video"
	"douyin/app/video/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.OptionalJWTMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: video.FeedHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/feed"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.OptionalJWTMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: video.PublishListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/publish"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseFormMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: video.PublishVideoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/publish"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: favorite.FavoriteOrUnfavoriteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/favorite"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.OptionalJWTMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: favorite.FavoriteVideosListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/favorite"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: comment.CommentOrDelCommentHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/comment"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.OptionalJWTMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: comment.CommentsListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/comment"),
	)
}
