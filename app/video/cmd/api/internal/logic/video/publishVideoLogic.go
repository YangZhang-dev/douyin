package video

import (
	"context"
	filepb "douyin/app/file/cmd/rpc/pb"
	videopb "douyin/app/video/cmd/rpc/pb"
	"douyin/common/ctxdata"
	"douyin/common/tool"
	"douyin/common/xerr"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"strconv"
	"time"

	"douyin/app/video/cmd/api/internal/svc"
	"douyin/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUploadFileError = xerr.NewErrCode(xerr.UPLOAD_FILE_ERROR)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	file   *multipart.FileHeader
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext, file *multipart.FileHeader) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		file:   file,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PublishReq) (resp *types.PublishResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	open, err := l.file.Open()
	if err != nil {
		return nil, errors.Wrapf(ErrUploadFileError, "读取文件错误 err:%+v", err)
	}
	bytes, err := io.ReadAll(open)
	if err != nil {
		return nil, errors.Wrapf(ErrUploadFileError, "读取文件错误 err:%+v", err)
	}
	bytes, err = tool.GZipBytes(bytes)
	if err != nil {
		return nil, errors.Wrapf(ErrUploadFileError, "读取文件错误 err:%+v", err)
	}

	fileName := uuid.New().String() + "-" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + strconv.FormatInt(userId, 10) + ".mp4"

	uploadVideoByOssResp, err := l.svcCtx.FileRpc.UploadVideoByOss(l.ctx, &filepb.UploadVideoByOssReq{
		VideoName: fileName,
		Data:      bytes,
	})
	if err == nil {
		err := l.saveVideoData(uploadVideoByOssResp.PlayUrl, uploadVideoByOssResp.CoverUrl, req.Title, userId)
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
		return &types.PublishResp{
			Status: types.Status{
				StatusCode: xerr.OK,
				StatusMsg:  xerr.MapErrMsg(xerr.OK),
			},
		}, nil
	}
	logx.Error(err)
	uploadVideoByLocalResp, err := l.svcCtx.FileRpc.UploadVideoByLocal(l.ctx, &filepb.UploadVideoByLocalReq{
		VideoName: fileName,
		Data:      bytes,
	})
	if err == nil {
		err := l.saveVideoData(uploadVideoByLocalResp.PlayUrl, uploadVideoByLocalResp.CoverUrl, req.Title, userId)
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
		return &types.PublishResp{
			Status: types.Status{
				StatusCode: xerr.OK,
				StatusMsg:  xerr.MapErrMsg(xerr.OK),
			},
		}, nil
	}
	logx.Error(err)
	return nil, ErrUploadFileError
}

func (l *PublishVideoLogic) saveVideoData(playUrl, coverUrl, title string, userId int64) error {
	_, err := l.svcCtx.VideoRpc.UploadVideo(l.ctx, &videopb.UploadVideoReq{
		Title:    title,
		UserId:   userId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	})
	if err != nil {
		return err
	}
	return nil
}
