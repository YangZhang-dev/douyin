-- 授权 root 用户可以远程链接
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'root';
flush privileges;

create database douyin_user default character set utf8mb4 collate utf8mb4_general_ci;
create database douyin_video default character set utf8mb4 collate utf8mb4_general_ci;


use douyin_user;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `chat`;
CREATE TABLE `chat`
(
    `id`          int          NOT NULL AUTO_INCREMENT COMMENT '娑堟伅id',
    `user_id`     int          NOT NULL COMMENT '发送用户id',
    `to_user_id`  int          NOT NULL COMMENT '接收消息用户id',
    `content`     varchar(200) NOT NULL,
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '鍒涘缓鏃堕棿',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_chat` (`user_id`, `to_user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`
(
    `id`          int        NOT NULL AUTO_INCREMENT COMMENT '鍏虫敞id',
    `user_id`     int        NOT NULL COMMENT '关注用户id',
    `to_user_id`  int        NOT NULL COMMENT '被关注用户id',
    `create_time` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '鍒涘缓鏃堕棿',
    `update_time` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_friend`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '0代表没有互相关注，1代表互相关注',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_to_user_id` (`to_user_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 22
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`             int                                                     NOT NULL AUTO_INCREMENT COMMENT ' 用户id',
    `username`       varchar(32)                                             NOT NULL COMMENT '用户名',
    `password`       varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
    `follow_count`   int                                                     NOT NULL DEFAULT '0' COMMENT '关注数量',
    `follower_count` int                                                     NOT NULL DEFAULT '0' COMMENT '粉丝数量',
    `create_time`    timestamp                                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '鍒涘缓鏃堕棿',
    `update_time`    timestamp                                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 34
  DEFAULT CHARSET = utf8mb4;



use douyin_video;
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`          int          NOT NULL AUTO_INCREMENT,
    `user_id`     int          NOT NULL COMMENT '评论用户id',
    `video_id`    int          NOT NULL COMMENT '被评论的视频id',
    `content`     varchar(200) NOT NULL COMMENT '评论内容',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '鍒涘缓鏃堕棿',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_video_id` (`video_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `favorite_video`;
CREATE TABLE `favorite_video`
(
    `id`          int       NOT NULL AUTO_INCREMENT COMMENT '点赞id',
    `user_id`     int       NOT NULL COMMENT '点赞用户id',
    `video_id`    int       NOT NULL COMMENT '视频id',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '鍒涘缓鏃堕棿',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`
(
    `id`             int          NOT NULL AUTO_INCREMENT COMMENT '视频id',
    `user_id`        int          NOT NULL COMMENT '发布作者id',
    `title`          varchar(32)  NOT NULL COMMENT '视频标题',
    `play_url`       varchar(200) NOT NULL COMMENT '视频播放地址',
    `cover_url`      varchar(200) NOT NULL COMMENT '封面地址',
    `favorite_count` int          NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `comment_count`  int          NOT NULL DEFAULT '0' COMMENT '评论数量',
    `create_time`    timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '鍒涘缓鏃堕棿',
    `update_time`    timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_create_time` (`create_time`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

