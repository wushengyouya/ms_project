/*
 Navicat Premium Data Transfer

 Source Server         : my_ms_project
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : localhost:3309
 Source Schema         : msproject

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 25/04/2025 10:30:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ms_member
-- ----------------------------
DROP TABLE IF EXISTS `ms_member`;
CREATE TABLE `ms_member`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '系统前台用户表',
  `account` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户登陆账号',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登陆密码',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '用户昵称',
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
  `realname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `create_time` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态',
  `last_login_time` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '上次登录时间',
  `sex` tinyint NULL DEFAULT 0 COMMENT '性别',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '头像',
  `idcard` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '身份证',
  `province` int NULL DEFAULT 0 COMMENT '省',
  `city` int NULL DEFAULT 0 COMMENT '市',
  `area` int NULL DEFAULT 0 COMMENT '区',
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '所在地址',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '备注',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `dingtalk_openid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '钉钉openid',
  `dingtalk_unionid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '钉钉unionid',
  `dingtalk_userid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '钉钉用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1003 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_member
-- ----------------------------
INSERT INTO `ms_member` VALUES (1003, 'mikasa', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'mikasa', '15545556677', '', '1744965913349', 1, '1744965913349', 0, '', '', 0, 0, 0, '', '', '100aaa@qq.com', '', '', '');

-- ----------------------------
-- Table structure for ms_organization
-- ----------------------------
DROP TABLE IF EXISTS `ms_organization`;
CREATE TABLE `ms_organization`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `avatar` varchar(511) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `member_id` bigint NULL DEFAULT NULL COMMENT '拥有者',
  `create_time` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `personal` tinyint(1) NULL DEFAULT 0 COMMENT '是否个人项目',
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `province` int NULL DEFAULT 0 COMMENT '省',
  `city` int NULL DEFAULT 0 COMMENT '市',
  `area` int NULL DEFAULT 0 COMMENT '区',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组织表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_organization
-- ----------------------------
INSERT INTO `ms_organization` VALUES (8, 'mikasa个人项目', 'https://www.baidu.com/img/flexible/logo/pc/index.png', '', 1000, 1744945464392, 1, '', 0, 0, 0);
INSERT INTO `ms_organization` VALUES (9, 'mikasa个人项目', 'https://www.baidu.com/img/flexible/logo/pc/index.png', '', 1001, 1744965521283, 1, '', 0, 0, 0);
INSERT INTO `ms_organization` VALUES (10, 'mikasa个人项目', 'https://www.baidu.com/img/flexible/logo/pc/index.png', '', 1002, 1744965707306, 1, '', 0, 0, 0);
INSERT INTO `ms_organization` VALUES (11, 'mikasa个人项目', 'https://www.baidu.com/img/flexible/logo/pc/index.png', '', 1003, 1744965913364, 1, '', 0, 0, 0);

-- ----------------------------
-- Table structure for ms_project
-- ----------------------------
DROP TABLE IF EXISTS `ms_project`;
CREATE TABLE `ms_project`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `cover` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '封面',
  `name` varchar(90) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '描述',
  `access_control_type` tinyint NULL DEFAULT 0 COMMENT '访问控制l类型',
  `white_list` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '可以访问项目的权限组（白名单）',
  `order` int UNSIGNED NULL DEFAULT 0 COMMENT '排序',
  `deleted` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  `template_code` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '项目类型',
  `schedule` double(5, 2) NULL DEFAULT 0.00 COMMENT '进度',
  `create_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `organization_code` bigint NULL DEFAULT NULL COMMENT '组织id',
  `deleted_time` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '删除时间',
  `private` tinyint(1) NULL DEFAULT 1 COMMENT '是否私有',
  `prefix` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '项目前缀',
  `open_prefix` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启项目前缀',
  `archive` tinyint(1) NULL DEFAULT 0 COMMENT '是否归档',
  `archive_time` bigint NULL DEFAULT NULL COMMENT '归档时间',
  `open_begin_time` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启任务开始时间',
  `open_task_private` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启新任务默认开启隐私模式',
  `task_board_theme` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'default' COMMENT '看板风格',
  `begin_time` bigint NULL DEFAULT NULL COMMENT '项目开始日期',
  `end_time` bigint NULL DEFAULT NULL COMMENT '项目截止日期',
  `auto_update_schedule` tinyint(1) NULL DEFAULT 0 COMMENT '自动更新项目进度',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project`(`order` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_project
-- ----------------------------

-- ----------------------------
-- Table structure for ms_project_member
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_member`;
CREATE TABLE `ms_project_member`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `project_code` bigint NULL DEFAULT NULL COMMENT '项目id',
  `member_code` bigint NULL DEFAULT NULL COMMENT '成员id',
  `join_time` bigint NULL DEFAULT NULL COMMENT '加入时间',
  `is_owner` bigint NULL DEFAULT 0 COMMENT '拥有者',
  `authorize` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique`(`project_code` ASC, `member_code` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目-成员表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_project_member
-- ----------------------------

-- ----------------------------
-- Table structure for ms_project_menu
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_menu`;
CREATE TABLE `ms_project_menu`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '父id',
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `url` varchar(400) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '链接',
  `file_path` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件路径',
  `params` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '链接参数',
  `node` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '#' COMMENT '权限节点',
  `sort` int UNSIGNED NULL DEFAULT 0 COMMENT '菜单排序',
  `status` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态(0:禁用,1:启用)',
  `create_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `is_inner` tinyint(1) NULL DEFAULT 0 COMMENT '是否内页',
  `values` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '参数默认值',
  `show_slider` tinyint(1) NULL DEFAULT 1 COMMENT '是否显示侧栏',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 168 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ms_project_menu
-- ----------------------------
INSERT INTO `ms_project_menu` VALUES (120, 0, '工作台', 'appstore-o', 'home', 'home', ':org', '#', 0, 1, 0, 0, '', 0);
INSERT INTO `ms_project_menu` VALUES (121, 0, '项目管理', 'project', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (122, 121, '项目列表', 'branches', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (124, 0, '系统设置', 'setting', '#', '#', '', '#', 100, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (125, 124, '成员管理', 'unlock', '#', '#', '', '#', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (126, 125, '账号列表', '', 'system/account', 'system/account', '', 'project/account/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (127, 122, '我的组织', '', 'organization', 'organization', '', 'project/organization/index', 30, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (130, 125, '访问授权', '', 'system/account/auth', 'system/account/auth', '', 'project/auth/index', 20, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (131, 125, '授权页面', '', 'system/account/apply', 'system/account/apply', ':id', 'project/auth/apply', 30, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menu` VALUES (138, 121, '消息提醒', 'info-circle-o', '#', '#', '', '#', 30, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (139, 138, '站内消息', '', 'notify/notice', 'notify/notice', '', 'project/notify/index', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (140, 138, '系统公告', '', 'notify/system', 'notify/system', '', 'project/notify/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (143, 124, '系统管理', 'appstore', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (144, 143, '菜单路由', '', 'system/config/menu', 'system/config/menu', '', 'project/menu/menuadd', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (145, 143, '访问节点', '', 'system/config/node', 'system/config/node', '', 'project/node/save', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (148, 124, '个人管理', 'user', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (149, 148, '个人设置', '', 'account/setting/base', 'account/setting/base', '', 'project/index/editpersonal', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (150, 148, '安全设置', '', 'account/setting/security', 'account/setting/security', '', 'project/index/editpersonal', 0, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menu` VALUES (151, 122, '我的项目', '', 'project/list', 'project/list', ':type', 'project/project/index', 0, 1, 0, 0, 'my', 1);
INSERT INTO `ms_project_menu` VALUES (152, 122, '回收站', '', 'project/recycle', 'project/recycle', '', 'project/project/index', 20, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (153, 121, '项目空间', 'heat-map', 'project/space/task', 'project/space/task', ':code', '#', 20, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menu` VALUES (154, 153, '任务详情', '', 'project/space/task/:code/detail', 'project/space/taskdetail', ':code', 'project/task/read', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (155, 122, '我的收藏', '', 'project/list', 'project/list', ':type', 'project/project/index', 10, 1, 0, 0, 'collect', 1);
INSERT INTO `ms_project_menu` VALUES (156, 121, '基础设置', 'experiment', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (157, 156, '项目模板', '', 'project/template', 'project/template', '', 'project/project_template/index', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (158, 156, '项目列表模板', '', 'project/template/taskStages', 'project/template/taskStages', ':code', 'project/task_stages_template/index', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (159, 122, '已归档项目', '', 'project/archive', 'project/archive', '', 'project/project/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (160, 0, '团队成员', 'team', '#', '#', '', '#', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (161, 153, '项目概况', '', 'project/space/overview', 'project/space/overview', ':code', 'project/index/info', 20, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (162, 153, '项目文件', '', 'project/space/files', 'project/space/files', ':code', 'project/index/info', 10, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (163, 122, '项目分析', '', 'project/analysis', 'project/analysis', '', 'project/index/info', 5, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (164, 160, '团队成员', '', '#', '#', '', '#', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (166, 164, '团队成员', '', 'members', 'members', '', 'project/department/index', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (167, 164, '成员信息', '', 'members/profile', 'members/profile', ':code', 'project/department/read', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (168, 153, '版本管理', '', 'project/space/features', 'project/space/features', ':code', 'project/index/info', 20, 1, 0, 1, '', 0);

SET FOREIGN_KEY_CHECKS = 1;
