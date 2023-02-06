package logic

import (
	"context"
	"douyin/app/user/cmd/rpc/internal/svc"
	"douyin/app/user/cmd/rpc/pb"
	"douyin/app/user/model"
	"douyin/common/xerr"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrDBError = xerr.NewErrCode(xerr.DB_ERROR)
var ErrDataFormatError = xerr.NewErrCode(xerr.DATA_FORMAT_ERROR)

type GetUserListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListByIdsLogic {
	return &GetUserListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListByIdsLogic) GetUserListByIds(in *pb.GetUserListByIdsReq) (*pb.GetUserListByIdsResp, error) {
	ids := in.Ids
	var query squirrel.SelectBuilder

	query = l.svcCtx.UserModel.RowBuilder().Where(squirrel.Eq{"id": ids})
	// TODO 加缓存
	users, err := l.svcCtx.UserModel.FindAll(l.ctx, query, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(ErrDBError, "通过ids查找用户 ids:%+v,err:%v", ids, err)
	}

	res := make([]*pb.User, len(users))
	_ = copier.Copy(&res, users)

	if in.UserId != nil {
		id := *in.UserId
		getFollowInfoLogic := NewGetFollowInfoLogic(l.ctx, l.svcCtx)
		for _, user := range res {
			followInfoResp, err := getFollowInfoLogic.GetFollowInfo(&pb.GetFollowInfoReq{
				UserId:   id,
				ToUserId: user.Id,
			})
			if err != nil {
				return nil, err
			}
			user.IsFollow = followInfoResp.IsFollow
		}
	}

	return &pb.GetUserListByIdsResp{Users: res}, nil
}
