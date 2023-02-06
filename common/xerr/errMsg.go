package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_PARSE_ERROR] = "token错误，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[DATA_FORMAT_ERROR] = "数据格式错误"

	message[USER_NOT_EXIST_ERROR] = "用户不存在"
	message[USER_ALREADY_EXIST_ERROR] = "用户已存在"
	message[USER_OR_PASSWORD_ERROR] = "用户名或密码错误"
	message[USER_ALREADY_FOLLOW_ERROR] = "用户已关注"
	message[USER_NOT_FOLLOW_ERROR] = "用户未关注"
	message[REGISTER_ERROR] = "注册失败"
	message[LOGIN_ERROR] = "登录失败"
	message[USER_CANNOT_ACTION_SELF] = "用户不能操作自己"
	message[USER_IS_NOT_FRIEND_ERROR] = "当前用户不是好友"

	message[FORM_PRASE_ERROR] = "表单校验错误"
	message[UPLOAD_FILE_TYPE_ERROR] = "上传文件类型错误"
	message[UPLOAD_FILE_LIMIT_EXCEEDED] = "上传文件大小超出限制"
	message[UPLOAD_FILE_ERROR] = "上传文件失败"
	message[UPLOAD_FILE_NOT_FOUND_ERROR] = "请上传文件"
	message[COMMENT_NOT_EXIST_ERROR] = "评论不存在"

	message[VIDEO_NOT_EXIST_ERROR] = "视频不存在"
	message[VIDEO_ALREADY_FAVORITE_ERROR] = "视频已点赞"
	message[VIDEO_NOT_FAVORITE_ERROR] = "视频未点赞"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
