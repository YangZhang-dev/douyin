package logic

import (
	"bytes"
	"context"
	"douyin/app/file/cmd/rpc/internal/svc"
	"douyin/app/file/cmd/rpc/pb"
	"douyin/common/globalkey"
	"douyin/common/tool"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"strings"
)

type UploadVideoByLocalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoByLocalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoByLocalLogic {
	return &UploadVideoByLocalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UploadVideoByLocal 上传视频到本地
func (l *UploadVideoByLocalLogic) UploadVideoByLocal(in *pb.UploadVideoByLocalReq) (*pb.UploadVideoByLocalResp, error) {
	videoPath := globalkey.LocalVideoPath + in.VideoName
	data, err := tool.UGZipBytes(in.Data)
	if err != nil {
		return nil, errors.Wrapf(ErrUploadFileError, "读取文件错误 err:%+v", err)
	}
	err = os.WriteFile(videoPath, data, 0666)
	if err != nil {
		return nil, errors.Wrapf(ErrCommonError, "local 上传错误 err:%+v", err)
	}
	coverPath := globalkey.LocalCoverPath + in.VideoName
	coverPath, err = GetSnapshot(videoPath, coverPath, 1)
	if err != nil {
		return nil, errors.Wrapf(ErrCommonError, "生成封面失败：videoPath:%+v,err:%+v", videoPath, err)
	}
	return &pb.UploadVideoByLocalResp{
		PlayUrl:  globalkey.StaticFileServiceIP + videoPath,
		CoverUrl: globalkey.StaticFileServiceIP + coverPath,
	}, nil
}
func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return "", err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		return "", err
	}
	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		return "", err
	}
	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}
