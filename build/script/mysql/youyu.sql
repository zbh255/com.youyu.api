/*
 Navicat Premium Data Transfer

 Source Server         : LocalDB
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : 192.168.1.150:3306
 Source Schema         : youyu

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 19/09/2021 16:23:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for advertisement
-- ----------------------------
DROP TABLE IF EXISTS `advertisement`;
CREATE TABLE `advertisement` (
  `advertisement_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '广告id',
  `advertisement_type` int(1) NOT NULL COMMENT '广告的类型，如内部或外部',
  `advertisement_link` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '广告链接',
  `advertisement_weight` int(4) NOT NULL COMMENT '广告的权重，权重越高的越靠前',
  `advertisement_body` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '广告的body页，可以插入图片或者视频',
  `advertisement_owner` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '广告的投放者：机构或个人名',
  PRIMARY KEY (`advertisement_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `article_id` char(32) COLLATE utf8mb4_bin NOT NULL COMMENT '文章id',
  `article_abstract` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '文章摘要',
  `article_content` longtext COLLATE utf8mb4_bin NOT NULL COMMENT '文章文本',
  `article_title` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '文章标题',
  `article_tag` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '文章的标签',
  `uid` int(11) NOT NULL COMMENT '文章的作者',
  `article_create_time` datetime NOT NULL COMMENT '文章的创建时间',
  `article_update_time` datetime NOT NULL COMMENT '文章的更新时间',
  PRIMARY KEY (`article_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for article_statistics
-- ----------------------------
DROP TABLE IF EXISTS `article_statistics`;
CREATE TABLE `article_statistics` (
  `article_id` char(32) COLLATE utf8mb4_bin NOT NULL COMMENT '文章id',
  `article_fabulous` int(11) NOT NULL COMMENT '文章的点赞数',
  `article_hot` int(11) NOT NULL COMMENT '文章的热度',
  `article_comment_num` int(11) NOT NULL COMMENT '文章的评论数',
  PRIMARY KEY (`article_id`) USING BTREE,
  CONSTRAINT `article_id` FOREIGN KEY (`article_id`) REFERENCES `article` (`article_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for comment_master
-- ----------------------------
DROP TABLE IF EXISTS `comment_master`;
CREATE TABLE `comment_master` (
  `comment_mid` int(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主评论的id',
  `type` int(1) NOT NULL COMMENT '评论的类型，比如文章下的评论',
  `text` varchar(255) NOT NULL COMMENT '评论的内容',
  `uid` int(11) unsigned NOT NULL COMMENT '评论者的id',
  `article_id` char(32) DEFAULT NULL COMMENT '文章的id',
  `fabulous` int(11) unsigned NOT NULL COMMENT '评论的点赞数量',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '评论的时间',
  `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '评论是否置顶',
  PRIMARY KEY (`comment_mid`,`uid`),
  KEY `uid` (`uid`),
  KEY `comment_mid` (`comment_mid`),
  CONSTRAINT `uid` FOREIGN KEY (`uid`) REFERENCES `user_base` (`uid`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for comment_slave
-- ----------------------------
DROP TABLE IF EXISTS `comment_slave`;
CREATE TABLE `comment_slave` (
  `comment_sid` int(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '从评论的id',
  `comment_mid` int(20) unsigned NOT NULL COMMENT '主评论的id',
  `text` varchar(255) NOT NULL COMMENT '从评论的内容',
  `fabulous` varchar(10) DEFAULT NULL COMMENT '从评论的点赞数量',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '评论的时间',
  `uid` int(11) unsigned NOT NULL COMMENT '发表从评论的用户',
  `article_id` char(32) DEFAULT NULL COMMENT '评论类型为文章时的文章id',
  `type` int(1) NOT NULL COMMENT '从评论的类型',
  `reply_id` int(20) unsigned DEFAULT NULL COMMENT '从评论回复重评论时需要填写的',
  PRIMARY KEY (`comment_sid`,`comment_mid`,`uid`) USING BTREE,
  KEY `slave_mid` (`comment_mid`),
  KEY `slave_uid` (`uid`),
  CONSTRAINT `slave_mid` FOREIGN KEY (`comment_mid`) REFERENCES `comment_master` (`comment_mid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `slave_uid` FOREIGN KEY (`uid`) REFERENCES `user_base` (`uid`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
  `tid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '标签的ID',
  `text` char(10) COLLATE utf8mb4_bin NOT NULL COMMENT '标签的内容，长度不能超过10',
  PRIMARY KEY (`tid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=28843 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user_base
-- ----------------------------
DROP TABLE IF EXISTS `user_base`;
CREATE TABLE `user_base` (
  `uid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `password` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '用户的密码',
  `salt` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'sha-512算法盐',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '用户名',
  PRIMARY KEY (`uid`) USING BTREE,
  KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '用户的登录名',
  `uid` int(10) unsigned NOT NULL COMMENT '用户ID',
  `phone` int(13) DEFAULT NULL COMMENT '用户的手机号码',
  `email` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户的邮箱',
  `phone_status` int(1) NOT NULL COMMENT '用户的手机验证状态',
  `email_status` int(1) NOT NULL COMMENT '用户的邮箱验证状态',
  `create_time` datetime NOT NULL COMMENT '用户的注册时间',
  `update_time` datetime NOT NULL COMMENT '用户信息的更新时间',
  `sex` int(1) NOT NULL COMMENT '用户的性别,1为男,2为女,3为保密',
  `age` int(3) NOT NULL COMMENT '用户的年龄,0为保密',
  `nick_name` varchar(10) COLLATE utf8mb4_bin NOT NULL COMMENT '用户的昵称',
  `explain` varchar(200) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户的简介/说明',
  `level` int(1) NOT NULL COMMENT '用户的系统级别',
  `wechat_openid` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户的微信凭证',
  `wechat_openid_status` int(1) NOT NULL COMMENT '用户的微信验证状态',
  `head_portrait` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户的头像地址',
  `country` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '国家',
  `province` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '省份',
  `city` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '城市',
  `detail_addr` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '详细住址',
  `language` varchar(10) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '语言',
  PRIMARY KEY (`name`,`uid`) USING BTREE,
  KEY `user_info_uid` (`uid`),
  CONSTRAINT `user_info_uid` FOREIGN KEY (`uid`) REFERENCES `user_base` (`uid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_name` FOREIGN KEY (`name`) REFERENCES `user_base` (`name`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
