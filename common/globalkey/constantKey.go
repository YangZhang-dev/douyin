package globalkey

/**
global constant key
*/

// 关注选项
const Follow int64 = 1
const Unfollow int64 = 2

// 点赞选项
const Favorite int64 = 1
const UnFavorite int64 = 2

// 是否互相关注
const MutualAttention int64 = 1
const NotMutualAttention int64 = 0

// 视频前缀
const OssVideoPath = "video/"
const LocalVideoPath = "static/video/"

// Oss 封面后缀
const OssCoverPath = "?x-oss-process=video/snapshot,t_1,m_fast"

// 视频上传大小限制
const MaxVideoSize int64 = 52428800

// 本地封面前缀
const LocalCoverPath = "static/cover/"

// 静态文件服务ip
const StaticFileServiceIP = "http://101.42.168.126:8888/"

//const StaticFileServiceIP = "http://localhost:8888/"

// 视频流数目
const FeedVideoNum uint64 = 10

// 发布评论
const PublishComment int64 = 1

// 删除评论
const DeleteComment int64 = 2
