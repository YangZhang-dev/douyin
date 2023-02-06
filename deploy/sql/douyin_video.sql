/*
Navicat MySQL Data Transfer

Source Server         : 192.168.72.128_3306
Source Server Version : 80020
Source Host           : 192.168.72.128:3306
Source Database       : douyin_video

Target Server Type    : MYSQL
Target Server Version : 80020
File Encoding         : 65001

Date: 2023-01-26 12:52:54
*/

SET
FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `comment`
-- ----------------------------
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
    KEY           `idx_video_id` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for `favorite_video`
-- ----------------------------
DROP TABLE IF EXISTS `favorite_video`;
CREATE TABLE `favorite_video`
(
    `id`          int       NOT NULL AUTO_INCREMENT COMMENT '点赞id',
    `user_id`     int       NOT NULL COMMENT '点赞用户id',
    `video_id`    int       NOT NULL COMMENT '视频id',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '鍒涘缓鏃堕棿',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of favorite_video
-- ----------------------------

-- ----------------------------
-- Table structure for `video`
-- ----------------------------
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
    KEY              `idx_user_id` (`user_id`) USING BTREE,
    KEY              `idx_create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of video
-- ----------------------------
