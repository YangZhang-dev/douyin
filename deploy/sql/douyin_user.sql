/*
Navicat MySQL Data Transfer

Source Server         : 192.168.72.128_3306
Source Server Version : 80020
Source Host           : 192.168.72.128:3306
Source Database       : douyin_user

Target Server Type    : MYSQL
Target Server Version : 80020
File Encoding         : 65001

Date: 2023-01-30 23:12:49
*/

SET
FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `chat`
-- ----------------------------
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
    KEY           `idx_chat` (`user_id`,`to_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of chat
-- ----------------------------

-- ----------------------------
-- Table structure for `follow`
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`
(
    `id`          int       NOT NULL AUTO_INCREMENT COMMENT '鍏虫敞id',
    `user_id`     int       NOT NULL COMMENT '关注用户id',
    `to_user_id`  int       NOT NULL COMMENT '被关注用户id',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '鍒涘缓鏃堕棿',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_friend`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '0代表没有互相关注，1代表互相关注',
    PRIMARY KEY (`id`),
    KEY           `idx_user_id` (`user_id`) USING BTREE,
    KEY           `idx_to_user_id` (`to_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of follow
-- ----------------------------
INSERT INTO `follow`
VALUES ('1', '26', '27', '2023-01-29 01:51:49', '2023-01-29 01:53:37', '0');
INSERT INTO `follow`
VALUES ('2', '24', '25', '2023-01-17 23:43:24', '2023-01-29 01:53:38', '0');
INSERT INTO `follow`
VALUES ('6', '26', '32', '2023-01-29 23:52:25', '2023-01-29 23:52:25', '0');
INSERT INTO `follow`
VALUES ('7', '24', '26', '2023-01-30 00:08:29', '2023-01-30 06:58:44', '1');
INSERT INTO `follow`
VALUES ('14', '24', '32', '2023-01-30 02:10:46', '2023-01-30 02:10:46', '0');
INSERT INTO `follow`
VALUES ('18', '24', '27', '2023-01-30 02:14:44', '2023-01-30 02:14:44', '0');
INSERT INTO `follow`
VALUES ('19', '25', '27', '2023-01-30 02:15:12', '2023-01-30 02:15:12', '0');
INSERT INTO `follow`
VALUES ('20', '25', '32', '2023-01-30 02:15:21', '2023-01-30 02:15:21', '0');
INSERT INTO `follow`
VALUES ('21', '26', '24', '2023-01-30 06:58:32', '2023-01-30 06:58:41', '1');

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
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
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user`
VALUES ('24', 'aaa', '$2a$04$09qocQ002ifRQObuIFlWVuszD9Nzd59HK3zE82lzilZ3x.cADDTp6', '4', '0', '2023-01-25 06:00:28',
        '2023-01-30 02:14:44');
INSERT INTO `user`
VALUES ('25', 'aaaa', '$2a$04$H3KwL3cLMD.F4QUjbiT.TO4Dx24ue7xYGp972e3mcIIOK5mMNkQqe', '2', '1', '2023-01-25 20:04:03',
        '2023-01-30 02:15:21');
INSERT INTO `user`
VALUES ('26', 'aaaaa', '$2a$04$ObEtvSvVuhpy.OaWiw.eXecbnO2HKCdSFp13O2J2pfCIIvfNuk5Jm', '2', '1', '2023-01-25 21:14:25',
        '2023-01-30 02:12:59');
INSERT INTO `user`
VALUES ('27', 'aaaaab', '$2a$04$qXKEzns9UldLoV9eNrlNwOLjD.C8VbZ.I.8Pza8UN9xBTnqAS.Hcu', '0', '3', '2023-01-25 21:15:05',
        '2023-01-30 02:15:12');
INSERT INTO `user`
VALUES ('28', 'aaaaaba', '$2a$04$pimU.IlUSIP6toSfJzPf9uUh1YoAAtNguYgPNiDcNLzVEd2KWIDH6', '0', '0',
        '2023-01-25 22:01:05', '2023-01-25 22:01:05');
INSERT INTO `user`
VALUES ('32', 'bbb', '$2a$04$RPRxb6dcqr6C9RxIJjzELOqcARD5IgdpCrrPC6XJL6MFDlAkXwthS', '0', '3', '2023-01-28 18:56:09',
        '2023-01-30 02:15:21');
INSERT INTO `user`
VALUES ('33', 'bbbk', '$2a$04$3DNWMLj8rB.5B7aPMIJhIup2fki9uDf4OxUB4dwxGC7lJrRpsN9b2', '0', '0', '2023-01-29 05:29:51',
        '2023-01-29 05:29:51');
