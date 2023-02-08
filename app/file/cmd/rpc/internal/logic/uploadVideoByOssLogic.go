package logic

import (
	"bytes"
	"context"
	"douyin/app/file/cmd/rpc/internal/svc"
	"douyin/app/file/cmd/rpc/pb"
	"douyin/common/globalkey"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrCommonError = xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)

type UploadVideoByOssLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoByOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoByOssLogic {
	return &UploadVideoByOssLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UploadVideoByOss 上传视频到OSS
func (l *UploadVideoByOssLogic) UploadVideoByOss(in *pb.UploadVideoByOssReq) (*pb.UploadVideoByOssResp, error) {
	pathName := globalkey.OssVideoPath + in.VideoName
	err := l.svcCtx.Xoss.PutObject(pathName, bytes.NewReader(in.Data))
	if err != nil {
		return nil, errors.Wrapf(ErrCommonError, "oss上传错误 err:%+v", err)
	}
	playUrl := l.svcCtx.Config.Xoss.PlayPath + pathName
	coverUrl := playUrl + globalkey.OssCoverPath

	return &pb.UploadVideoByOssResp{
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}, nil
}
