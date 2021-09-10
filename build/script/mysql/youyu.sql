/*
 Navicat Premium Data Transfer

 Source Server         : LocalRoot
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 192.168.230.128:3306
 Source Schema         : youyu

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 11/09/2021 00:25:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for advertisement
-- ----------------------------
DROP TABLE IF EXISTS `advertisement`;
CREATE TABLE `advertisement`  (
  `advertisement_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '广告id',
  `advertisement_type` int(1) NOT NULL COMMENT '广告的类型，如内部或外部',
  `advertisement_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '广告链接',
  `advertisement_weight` int(4) NOT NULL COMMENT '广告的权重，权重越高的越靠前',
  `advertisement_body` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '广告的body页，可以插入图片或者视频',
  `advertisement_owner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '广告的投放者：机构或个人名',
  PRIMARY KEY (`advertisement_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `article_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '文章id',
  `article_abstract` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '文章摘要',
  `article_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '文章文本',
  `article_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '文章标题',
  `article_tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '文章的标签',
  `uid` int(11) NOT NULL COMMENT '文章的作者',
  `article_create_time` datetime NOT NULL COMMENT '文章的创建时间',
  `article_update_time` datetime NOT NULL COMMENT '文章的更新时间',
  PRIMARY KEY (`article_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for article_statistics
-- ----------------------------
DROP TABLE IF EXISTS `article_statistics`;
CREATE TABLE `article_statistics`  (
  `article_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '文章id',
  `article_fabulous` int(11) NOT NULL COMMENT '文章的点赞数',
  `article_hot` int(11) NOT NULL COMMENT '文章的热度',
  `article_comment_num` int(11) NOT NULL COMMENT '文章的评论数',
  PRIMARY KEY (`article_id`) USING BTREE,
  CONSTRAINT `article_id` FOREIGN KEY (`article_id`) REFERENCES `article` (`article_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`  (
  `tid` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '标签的ID',
  `text` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '标签的内容，长度不能超过10',
  PRIMARY KEY (`tid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28843 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_base
-- ----------------------------
DROP TABLE IF EXISTS `user_base`;
CREATE TABLE `user_base`  (
  `uid` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户的密码',
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'sha-512算法盐',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户名',
  PRIMARY KEY (`uid`) USING BTREE,
  INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户的登录名',
  `uid` int(10) NOT NULL COMMENT '用户ID',
  `phone` int(13) NULL DEFAULT NULL COMMENT '用户的手机号码',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户的邮箱',
  `phone_status` int(1) NOT NULL COMMENT '用户的手机验证状态',
  `email_status` int(1) NOT NULL COMMENT '用户的邮箱验证状态',
  `create_time` datetime NOT NULL COMMENT '用户的注册时间',
  `update_time` datetime NOT NULL COMMENT '用户信息的更新时间',
  `sex` int(1) NOT NULL COMMENT '用户的性别,1为男,2为女,3为保密',
  `age` int(3) NOT NULL COMMENT '用户的年龄,0为保密',
  `nick_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户的昵称',
  `addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户的住址',
  `explain` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户的简介/说明',
  `level` int(1) NOT NULL COMMENT '用户的系统级别',
  PRIMARY KEY (`name`) USING BTREE,
  CONSTRAINT `user_name` FOREIGN KEY (`name`) REFERENCES `user_base` (`name`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
