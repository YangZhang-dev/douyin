package logic

import (
	"context"
	"douyin/app/user/model"
	"douyin/common/globalkey"
	"douyin/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"douyin/app/user/cmd/rpc/internal/svc"
	"douyin/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNotExistError = xerr.NewErrCode(xerr.USER_NOT_EXIST_ERROR)
var ErrUserNotFollowError = xerr.NewErrCode(xerr.USER_NOT_FOLLOW_ERROR)

type UnFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFollowLogic {
	return &UnFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消关注
func (l *UnFollowLogic) UnFollow(in *pb.UnFollowReq) (*pb.UnFollowResp, error) {
	getUserListByIdsLogic := NewGetUserListByIdsLogic(l.ctx, l.svcCtx)
	getUserListByIdsResp, err := getUserListByIdsLogic.GetUserListByIds(&pb.GetUserListByIdsReq{UserId: &in.UserId, Ids: []int64{in.ToUserId, in.UserId}})
	users := getUserListByIdsResp.Users
	if err != nil {
		return nil, err
	}
	if len(users) != 2 {
		return nil, errors.Wrapf(ErrUserNotExistError, "用户不存在 users:%v,err:%v", users, err)
	}

	getFollowInfoLogic := NewGetFollowInfoLogic(l.ctx, l.svcCtx)

	err = l.svcCtx.FollowModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		// 更新用户关注和粉丝数量
		for _, user := range users {
			if user.Id == in.UserId {
				user.FollowCount--
			} else {
				if !user.IsFollow {
					return errors.Wrapf(ErrUserNotFollowError, "用户未关注 user_id:%v,to_user_id:%v", in.UserId, in.ToUserId)
				}
				user.FollowerCount--
			}
			_, err := l.svcCtx.UserModel.Update(l.ctx, session, &model.User{
				Id:            user.Id,
				Username:      user.Username,
				Password:      user.Password,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
			})
			if err != nil {
				return err
			}
		}

		// 查询对方对自己的关注情况
		getFollowInfoResp, err := getFollowInfoLogic.GetFollowInfo(&pb.GetFollowInfoReq{
			UserId:   in.ToUserId,
			ToUserId: in.UserId,
		})
		if err != nil {
			return err
		}

		var follow model.Follow
		_ = copier.Copy(&follow, getFollowInfoResp.Follow)
		if getFollowInfoResp.IsFollow {
			follow.IsFriend = globalkey.NotMutualAttention
		}

		err = l.svcCtx.FollowModel.DeleteOneByUserIdAndToUserId(l.ctx, session, in.UserId, in.ToUserId)
		if err != nil {
			return errors.Wrapf(ErrDBError, "删除关注错误 user_id:%v,to_user_id:%v", in.UserId, in.ToUserId)
		}

		_, err = l.svcCtx.FollowModel.Update(l.ctx, session, &follow)
		if err != nil {
			return errors.Wrapf(ErrDBError, "更新关注错误 user_id:%v,to_user_id:%v", in.UserId, in.ToUserId)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.UnFollowResp{}, nil
}
