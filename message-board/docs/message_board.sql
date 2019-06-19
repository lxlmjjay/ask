/*
Source Database       : message_board
Target Server Type    : MYSQL
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL UNIQUE COMMENT '账号',
  `password` varchar(50) NOT NULL COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `message_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '标签ID',
  `username` varchar(50) NOT NULL UNIQUE COMMENT '账号',
  `title` varchar(100) NOT NULL COMMENT '文章标题',
  `content` text COMMENT '内容',
  `image_url` varchar(255) DEFAULT '' COMMENT '图片地址',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`message_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='文章管理';



