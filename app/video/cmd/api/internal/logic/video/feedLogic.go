package video

import (
	"context"
	userpb "douyin/app/user/cmd/rpc/pb"
	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"
	videopb "douyin/app/video/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrDataFormatError = xerr.NewErrCode(xerr.DATA_FORMAT_ERROR)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedResp, err error) {
	// 1. 准备数据
	userId := ctxdata.GetUidFromCtx(l.ctx)
	var curUserId *int64
	if userId != 0 {
		curUserId = &userId
	}
	var lastTime *int64
	if req.LastTime != "" {
		*lastTime, err = strconv.ParseInt(req.LastTime, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(ErrDataFormatError, "数据转换错误 str:%+v,err:%+v", req.LastTime, err)
		}
	}

	// 2. 请求视频列表
	feedResp, err := l.svcCtx.VideoRpc.Feed(l.ctx, &videopb.FeedReq{
		UserId:   curUserId,
		LastTime: lastTime,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	// 3. 获取作者信息
	videos := feedResp.VideoList
	authorsId, res := DrawAuthorId(videos)

	userListByIdsResp, err := l.svcCtx.UserRpc.GetUserListByIds(l.ctx, &userpb.GetUserListByIdsReq{
		UserId: curUserId,
		Ids:    authorsId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	// 4. 将作者信息封装到视频信息中
	authors := userListByIdsResp.Users
	SetAuthorInfo(res, authors)

	return &types.FeedResp{
		Status: types.Status{
			StatusCode: xerr.OK,
			StatusMsg:  xerr.MapErrMsg(xerr.OK),
		},
		VideoList: res,
		NextTime:  feedResp.NextTime,
	}, nil
}

func SetAuthorInfo(res []types.Video, authors []*userpb.User) {
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(authors); j++ {
			if authors[j].Id == res[i].Author.Id {
				_ = copier.Copy(&res[i].Author, authors[j])
				break
			}
		}
	}
}

func DrawAuthorId(videos []*videopb.Video) ([]int64, []types.Video) {
	var authorsId []int64
	res := make([]types.Video, len(videos))
	_ = copier.Copy(&res, videos)

	for i := 0; i < len(videos); i++ {
		authorsId = append(authorsId, videos[i].UserId)
		for j := 0; j < len(res); j++ {
			if res[j].Id == videos[i].Id {
				_ = copier.Copy(&(res[j].Author.Id), videos[i].UserId)
			}
		}
	}
	return authorsId, res
}
