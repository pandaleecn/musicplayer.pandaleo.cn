/*
 Navicat Premium Data Transfer

 Source Server         : iris
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost
 Source Database       : musicplayer

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : utf-8

 Date: 08/05/2020 09:59:46 AM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `casbin_rule`
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `casbin_rule`
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES ('p', '4', '/v1/admin/roles', 'POST', '', '', ''), ('p', '4', '/v1/admin/roles/{id:uint}', 'PUT', '', '', ''), ('p', '4', '/v1/admin/roles/{id:uint}', 'GET', '', '', ''), ('p', '4', '/v1/admin/roles/{id:uint}', 'DELETE', '', '', ''), ('p', '4', '/v1/admin/permissions', 'GET', '', '', ''), ('p', '4', '/v1/admin/permissions/{id:uint}', 'GET', '', '', ''), ('p', '4', '/v1/admin/roles', 'GET', '', '', ''), ('p', '4', '/v1/admin/permissions/import', 'POST', '', '', ''), ('p', '4', '/v1/admin/users/profile', 'GET', '', '', ''), ('p', '4', '/v1/admin/permissions', 'POST', '', '', ''), ('p', '4', '/v1/admin/permissions/{id:uint}', 'PUT', '', '', ''), ('p', '4', '/v1/admin/users/{id:uint}', 'DELETE', '', '', ''), ('p', '4', '/v1/admin/permissions/{id:uint}', 'DELETE', '', '', ''), ('p', '4', '/v1/admin/songs', 'GET', '', '', ''), ('p', '4', '/v1/admin/songs/{id:uint}', 'GET', '', '', ''), ('p', '4', '/v1/admin/users/{id:uint}', 'PUT', '', '', ''), ('p', '4', '/v1/admin/songs', 'POST', '', '', ''), ('p', '4', '/v1/admin/songs/{id:uint}', 'PUT', '', '', ''), ('p', '4', '/v1/admin/users', 'POST', '', '', ''), ('p', '4', '/v1/admin/songs/{id:uint}', 'DELETE', '', '', ''), ('p', '4', '/v1/admin/playlists', 'GET', '', '', ''), ('p', '4', '/v1/admin/playlists/{id:uint}', 'GET', '', '', ''), ('p', '4', '/v1/admin/users/{id:uint}', 'GET', '', '', ''), ('p', '4', '/v1/admin/playlists', 'POST', '', '', ''), ('p', '4', '/v1/admin/playlists/{id:uint}', 'PUT', '', '', ''), ('p', '4', '/v1/admin/users', 'GET', '', '', ''), ('p', '4', '/v1/admin/playlists/{id:uint}', 'DELETE', '', '', ''), ('p', '4', '/v1/admin/albums', 'GET', '', '', ''), ('p', '4', '/v1/admin/albums/{id:uint}', 'GET', '', '', ''), ('p', '4', '/v1/admin/logout', 'GET', '', '', ''), ('p', '4', '/v1/admin/albums', 'POST', '', '', ''), ('p', '4', '/v1/admin/albums/{id:uint}', 'PUT', '', '', ''), ('p', '4', '/v1/admin/albums/{id:uint}', 'DELETE', '', '', ''), ('p', '4', '/v1/admin/artists', 'GET', '', '', ''), ('p', '4', '/v1/admin/artists/{id:uint}', 'GET', '', '', ''), ('p', '4', '/v1/admin/artists', 'POST', '', '', ''), ('p', '4', '/v1/admin/artists/{id:uint}', 'PUT', '', '', ''), ('p', '4', '/v1/admin/artists/{id:uint}', 'DELETE', '', '', ''), ('p', '4', '/v1/admin/lyrics', 'GET', '', '', ''), ('p', '4', '/v1/admin/lyrics/{id:uint}', 'GET', '', '', ''), ('p', '4', '/v1/admin/lyrics', 'POST', '', '', ''), ('p', '4', '/v1/admin/lyrics/{id:uint}', 'PUT', '', '', ''), ('p', '4', '/v1/admin/lyrics/{id:uint}', 'DELETE', '', '', ''), ('p', '4', '/v1/admin/qiniutoken', 'GET', '', '', ''), ('g', '5', '4', '', '', '', '');
COMMIT;

-- ----------------------------
--  Table structure for `iris_albums`
-- ----------------------------
DROP TABLE IF EXISTS `iris_albums`;
CREATE TABLE `iris_albums` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `cover` varchar(255) DEFAULT NULL,
  `artist_id` int DEFAULT NULL,
  `create_user_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_iris_albums_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `iris_albums`
-- ----------------------------
BEGIN;
INSERT INTO `iris_albums` VALUES ('7', '2020-08-04 10:17:35', '2020-08-04 14:55:50', null, '11月的肖邦', '', '0', '5'), ('8', '2020-08-04 10:17:49', '2020-08-05 09:30:07', null, '魔杰座', '', '0', '5'), ('9', '2020-08-04 11:03:18', '2020-08-05 09:30:19', null, '七里香', '', '0', '5'), ('10', '2020-08-05 09:53:08', '2020-08-05 09:53:08', null, '宠爱', '', '0', '5'), ('11', '2020-08-05 09:53:48', '2020-08-05 09:53:48', null, 'Summer Romance\'87', '', '0', '5'), ('12', '2020-08-05 09:54:06', '2020-08-05 09:54:06', null, '陪你倒数', '', '0', '5');
COMMIT;

-- ----------------------------
--  Table structure for `iris_artists`
-- ----------------------------
DROP TABLE IF EXISTS `iris_artists`;
CREATE TABLE `iris_artists` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `poster` varchar(255) DEFAULT NULL,
  `create_user_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_iris_artists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `iris_artists`
-- ----------------------------
BEGIN;
INSERT INTO `iris_artists` VALUES ('4', '2020-08-04 10:18:27', '2020-08-04 14:15:48', null, '周杰伦', '', '5'), ('5', '2020-08-05 09:49:14', '2020-08-05 09:49:14', null, '张国荣', '', '5');
COMMIT;

-- ----------------------------
--  Table structure for `iris_lyrics`
-- ----------------------------
DROP TABLE IF EXISTS `iris_lyrics`;
CREATE TABLE `iris_lyrics` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `song_id` int unsigned DEFAULT NULL,
  `create_user_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_iris_lyrics_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `iris_lyrics`
-- ----------------------------
BEGIN;
INSERT INTO `iris_lyrics` VALUES ('1', '2020-08-04 11:17:38', '2020-08-04 11:17:38', '2020-08-05 09:58:35', '333', '', '0', '5'), ('2', '2020-08-04 14:12:06', '2020-08-04 14:12:06', '2020-08-05 09:58:38', '111', '', '0', '5'), ('3', '2020-08-04 14:12:37', '2020-08-04 14:12:37', null, '666', '', '0', '5');
COMMIT;

-- ----------------------------
--  Table structure for `iris_oauth_tokens`
-- ----------------------------
DROP TABLE IF EXISTS `iris_oauth_tokens`;
CREATE TABLE `iris_oauth_tokens` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `user_id` int unsigned DEFAULT NULL,
  `secret` varchar(255) DEFAULT NULL,
  `express_in` bigint DEFAULT NULL,
  `revoked` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_iris_oauth_tokens_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `iris_oauth_tokens`
-- ----------------------------
BEGIN;
INSERT INTO `iris_oauth_tokens` VALUES ('15', '2020-08-04 10:17:28', '2020-08-04 10:17:28', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MDM4NDgsImlhdCI6MTU5NjUwNzQ0OH0.JvER25mmJ99VdLYxqllV_dKYTR293paUrdLftoHs3TY', '5', 'secret', '1596511048', '0'), ('16', '2020-08-04 10:20:40', '2020-08-04 10:20:40', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MDQwNDAsImlhdCI6MTU5NjUwNzY0MH0.qFx0IlKqUkC-Y4VGLwyXKF2W1zTPnvZNvDMTHt6Hjko', '5', 'secret', '1596511240', '0'), ('17', '2020-08-04 10:30:01', '2020-08-04 10:30:01', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MDQ2MDAsImlhdCI6MTU5NjUwODIwMH0.c_7JiVZgrnoTRM0wgELW1Ybhb5NpuvhqtI5_NEQP-Gc', '5', 'secret', '1596511800', '0'), ('18', '2020-08-04 11:01:21', '2020-08-04 11:01:21', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MDY0ODAsImlhdCI6MTU5NjUxMDA4MH0.10KKmKhZPvBi3fcrg1i7dAT1Ds8t5N9Z7nw8eO3VNks', '5', 'secret', '1596513680', '0'), ('19', '2020-08-04 11:13:49', '2020-08-04 11:13:49', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MDcyMjksImlhdCI6MTU5NjUxMDgyOX0.zGXO_Vdbpnz0-r1IX3_aFMAuFLq_ciLcGH_caVfvL10', '5', 'secret', '1596514429', '0'), ('20', '2020-08-04 12:22:30', '2020-08-04 12:22:30', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MTEzNTAsImlhdCI6MTU5NjUxNDk1MH0.gV6zAf2Qm6Z73Zeu8GSMP6TY8KFvkxEjyA5TQOt7mS4', '5', 'secret', '1596518550', '0'), ('21', '2020-08-04 14:11:46', '2020-08-04 14:11:46', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MTc5MDUsImlhdCI6MTU5NjUyMTUwNX0.HO-cGDX9DdDrrCp-41SxEgEiVOfIxnadV9Xl7Vo3laQ', '5', 'secret', '1596525105', '0'), ('22', '2020-08-04 14:12:26', '2020-08-04 14:12:26', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MTc5NDYsImlhdCI6MTU5NjUyMTU0Nn0.cGkurGgDC1b2BngLY6hrJxf7TsY6bYOMNYy0b6UJYQo', '5', 'secret', '1596525146', '0'), ('23', '2020-08-04 15:27:38', '2020-08-04 15:27:38', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MjI0NTcsImlhdCI6MTU5NjUyNjA1N30.WQvmkkqiq3C-lvaKe-_gX8g3yFtHpXfdzrVMxVxYte0', '5', 'secret', '1596529657', '0'), ('24', '2020-08-04 16:31:18', '2020-08-04 16:31:18', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MjYyNzgsImlhdCI6MTU5NjUyOTg3OH0.D_4l_Hoo_rl3nRRgpynruXKvejIZPxSpP93VQQCVOaA', '5', 'secret', '1596533478', '0'), ('25', '2020-08-04 17:41:36', '2020-08-04 17:41:36', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MzA0OTYsImlhdCI6MTU5NjUzNDA5Nn0.e7rLAhrXoq1v6D43Arl2ok2FwwmwadQ8xCnPbmovI9s', '5', 'secret', '1596537696', '0'), ('26', '2020-08-04 18:41:57', '2020-08-04 18:41:57', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1MzQxMTYsImlhdCI6MTU5NjUzNzcxNn0.2Gm3N11dYGbZuw5-1DbFBaZxdEYdJn9WZyUazNgPQF4', '5', 'secret', '1596541316', '0'), ('27', '2020-08-05 09:28:29', '2020-08-05 09:28:29', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5NTY1ODczMDksImlhdCI6MTU5NjU5MDkwOX0.B76Ls_FmwvHTnEmyaUUgB2YoFBQO826EqQXlrlvvF0k', '5', 'secret', '1596594509', '0');
COMMIT;

-- ----------------------------
--  Table structure for `iris_permissions`
-- ----------------------------
DROP TABLE IF EXISTS `iris_permissions`;
CREATE TABLE `iris_permissions` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `display_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `act` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_iris_permissions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=138 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `iris_permissions`
-- ----------------------------
BEGIN;
INSERT INTO `iris_permissions` VALUES ('94', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/roles', '创建角色', '创建角色', 'POST'), ('95', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/roles/{id:uint}', '编辑角色', '编辑角色', 'PUT'), ('96', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/roles/{id:uint}', '角色详情', '角色详情', 'GET'), ('97', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/roles/{id:uint}', '删除角色', '删除角色', 'DELETE'), ('98', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/permissions', '权限列表', '权限列表', 'GET'), ('99', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/permissions/{id:uint}', '权限详情', '权限详情', 'GET'), ('100', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/roles', '角色列表', '角色列表', 'GET'), ('101', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/permissions/import', '导入权限', '导入权限', 'POST'), ('102', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/users/profile', '个人信息', '个人信息', 'GET'), ('103', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/permissions', '创建权限', '创建权限', 'POST'), ('104', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/permissions/{id:uint}', '编辑权限', '编辑权限', 'PUT'), ('105', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/users/{id:uint}', '删除用户', '删除用户', 'DELETE'), ('106', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/permissions/{id:uint}', '删除权限', '删除权限', 'DELETE'), ('107', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/songs', '用户歌曲列表', '用户歌曲列表', 'GET'), ('108', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/songs/{id:uint}', '歌曲详情', '歌曲详情', 'GET'), ('109', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/users/{id:uint}', '编辑用户', '编辑用户', 'PUT'), ('110', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/songs', '新增歌曲', '新增歌曲', 'POST'), ('111', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/songs/{id:uint}', '编辑歌曲', '编辑歌曲', 'PUT'), ('112', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/users', '创建用户', '创建用户', 'POST'), ('113', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/songs/{id:uint}', '删除歌曲', '删除歌曲', 'DELETE'), ('114', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/playlists', '歌单列表', '歌单列表', 'GET'), ('115', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/playlists/{id:uint}', '歌单详情', '歌单详情', 'GET'), ('116', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/users/{id:uint}', '用户详情', '用户详情', 'GET'), ('117', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/playlists', '新增歌单', '新增歌单', 'POST'), ('118', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/playlists/{id:uint}', '编辑歌单', '编辑歌单', 'PUT'), ('119', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/users', '用户列表', '用户列表', 'GET'), ('120', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/playlists/{id:uint}', '删除歌单', '删除歌单', 'DELETE'), ('121', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/albums', '专辑列表', '专辑列表', 'GET'), ('122', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/albums/{id:uint}', '专辑详情', '专辑详情', 'GET'), ('123', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/logout', '退出', '退出', 'GET'), ('124', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/albums', '新增专辑', '新增专辑', 'POST'), ('125', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/albums/{id:uint}', '编辑专辑', '编辑专辑', 'PUT'), ('126', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/albums/{id:uint}', '删除专辑', '删除专辑', 'DELETE'), ('127', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/artists', '歌手列表', '歌手列表', 'GET'), ('128', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/artists/{id:uint}', '歌手详情', '歌手详情', 'GET'), ('129', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/artists', '新增歌手', '新增歌手', 'POST'), ('130', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/artists/{id:uint}', '编辑歌手', '编辑歌手', 'PUT'), ('131', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/artists/{id:uint}', '删除歌手', '删除歌手', 'DELETE'), ('132', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/lyrics', '歌词列表', '歌词列表', 'GET'), ('133', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/lyrics/{id:uint}', '歌词详情', '歌词详情', 'GET'), ('134', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/lyrics', '新增歌词', '新增歌词', 'POST'), ('135', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/lyrics/{id:uint}', '编辑歌词', '编辑歌词', 'PUT'), ('136', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/lyrics/{id:uint}', '删除歌词', '删除歌词', 'DELETE'), ('137', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, '/v1/admin/qiniutoken', '七牛token', '七牛token', 'GET');
COMMIT;

-- ----------------------------
--  Table structure for `iris_playlists`
-- ----------------------------
DROP TABLE IF EXISTS `iris_playlists`;
CREATE TABLE `iris_playlists` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `user_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_iris_playlists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `iris_playlists`
-- ----------------------------
BEGIN;
INSERT INTO `iris_playlists` VALUES ('3', '2020-08-04 10:18:18', '2020-08-05 09:37:39', null, '学习', '5'), ('4', '2020-08-04 16:08:05', '2020-08-05 09:37:39', null, '运动', '5'), ('5', '2020-08-04 16:08:10', '2020-08-05 09:32:29', null, '日常', '5');
COMMIT;

-- ----------------------------
--  Table structure for `iris_playlists_songs`
-- ----------------------------
DROP TABLE IF EXISTS `iris_playlists_songs`;
CREATE TABLE `iris_playlists_songs` (
  `song_id` int unsigned NOT NULL,
  `playlist_id` int unsigned NOT NULL,
  PRIMARY KEY (`song_id`,`playlist_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `iris_playlists_songs`
-- ----------------------------
BEGIN;
INSERT INTO `iris_playlists_songs` VALUES ('5', '3'), ('5', '5'), ('6', '3'), ('6', '4'), ('7', '3'), ('7', '4');
COMMIT;

-- ----------------------------
--  Table structure for `iris_roles`
-- ----------------------------
DROP TABLE IF EXISTS `iris_roles`;
CREATE TABLE `iris_roles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `display_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_iris_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `iris_roles`
-- ----------------------------
BEGIN;
INSERT INTO `iris_roles` VALUES ('4', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, 'admin', '管理员', '管理员');
COMMIT;

-- ----------------------------
--  Table structure for `iris_songs`
-- ----------------------------
DROP TABLE IF EXISTS `iris_songs`;
CREATE TABLE `iris_songs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `cover` varchar(255) DEFAULT NULL,
  `artist_id` int unsigned DEFAULT NULL,
  `lrc` varchar(255) DEFAULT NULL,
  `upload_user_id` int unsigned DEFAULT NULL,
  `album_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_iris_songs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `iris_songs`
-- ----------------------------
BEGIN;
INSERT INTO `iris_songs` VALUES ('4', '2020-08-04 14:15:34', '2020-08-04 18:32:07', '2020-08-05 09:28:40', '周杰伦 - 珊瑚海', 'http://files.pandaleo.cn/周杰伦 - 珊瑚海-2020-08-04下午2-15-28.mp3', '', '4', '', '5', '7'), ('5', '2020-08-05 09:30:49', '2020-08-05 09:32:29', null, '周杰伦 - 发如雪', 'http://files.pandaleo.cn/周杰伦 - 发如雪-2020-08-05上午9-30-28.mp3', '', '4', '', '5', '7'), ('6', '2020-08-05 09:33:55', '2020-08-05 09:34:06', null, '周杰伦 - 枫', 'http://files.pandaleo.cn/周杰伦 - 枫-2020-08-05上午9-33-45.mp3', '', '4', '', '5', '7'), ('7', '2020-08-05 09:34:59', '2020-08-05 09:37:39', null, '周杰伦 - 黑色毛衣', 'http://files.pandaleo.cn/周杰伦 - 黑色毛衣-2020-08-05上午9-34-47.mp3', '', '4', '', '5', '7'), ('8', '2020-08-05 09:38:25', '2020-08-05 09:41:50', null, '周杰伦 - 蓝色风暴', 'http://files.pandaleo.cn/周杰伦 - 蓝色风暴-2020-08-05上午9-38-06.mp3', '', '4', '', '5', '7'), ('9', '2020-08-05 09:45:52', '2020-08-05 09:45:52', null, '周杰伦 - 浪漫手机', 'http://files.pandaleo.cn/周杰伦 - 浪漫手机-2020-08-05上午9-43-35.mp3', '', '4', '', '5', '7'), ('10', '2020-08-05 09:46:15', '2020-08-05 09:46:15', null, '周杰伦 - 飘移', 'http://files.pandaleo.cn/周杰伦 - 飘移-2020-08-05上午9-46-04.mp3', '', '4', '', '5', '7'), ('11', '2020-08-05 09:46:32', '2020-08-05 09:46:32', null, '周杰伦 - 珊瑚海', 'http://files.pandaleo.cn/周杰伦 - 珊瑚海-2020-08-05上午9-46-20.mp3', '', '4', '', '5', '7'), ('12', '2020-08-05 09:46:48', '2020-08-05 09:46:48', null, '周杰伦 - 夜曲', 'http://files.pandaleo.cn/周杰伦 - 夜曲-2020-08-05上午9-46-41.mp3', '', '4', '', '5', '7'), ('13', '2020-08-05 09:47:03', '2020-08-05 09:47:03', null, '周杰伦 - 一路向北', 'http://files.pandaleo.cn/周杰伦 - 一路向北-2020-08-05上午9-46-54.mp3', '', '4', '', '5', '7'), ('14', '2020-08-05 09:47:34', '2020-08-05 09:47:34', null, '周杰伦 - 稻香', 'http://files.pandaleo.cn/周杰伦 - 稻香-2020-08-05上午9-47-22.mp3', '', '4', '', '5', '8'), ('15', '2020-08-05 09:47:55', '2020-08-05 09:47:55', null, '周杰伦 - 给我一首歌的时间', 'http://files.pandaleo.cn/周杰伦 - 给我一首歌的时间-2020-08-05上午9-47-41.mp3', '', '4', '', '5', '8'), ('16', '2020-08-05 09:48:08', '2020-08-05 09:48:08', null, '周杰伦 - 花海', 'http://files.pandaleo.cn/周杰伦 - 花海-2020-08-05上午9-48-00.mp3', '', '4', '', '5', '8'), ('17', '2020-08-05 09:48:28', '2020-08-05 09:48:28', null, '周杰伦 - 兰亭序', 'http://files.pandaleo.cn/周杰伦 - 兰亭序-2020-08-05上午9-48-17.mp3', '', '4', '', '5', '8'), ('18', '2020-08-05 09:48:44', '2020-08-05 09:48:44', null, '周杰伦 - 时光机', 'http://files.pandaleo.cn/周杰伦 - 时光机-2020-08-05上午9-48-36.mp3', '', '4', '', '5', '8'), ('19', '2020-08-05 09:49:01', '2020-08-05 09:49:01', null, '周杰伦 - 说好的幸福呢', 'http://files.pandaleo.cn/周杰伦 - 说好的幸福呢-2020-08-05上午9-48-50.mp3', '', '4', '', '5', '8'), ('20', '2020-08-05 09:54:37', '2020-08-05 09:54:37', null, '张国荣 - 当爱已成往事（高品质）', 'http://files.pandaleo.cn/张国荣 - 当爱已成往事（高品质）-2020-08-05上午9-54-24.mp3', '', '5', '', '5', '10'), ('21', '2020-08-05 09:56:58', '2020-08-05 09:56:58', null, '张国荣-倩女幽魂 (电影《倩女幽魂》主题曲)', 'http://files.pandaleo.cn/张国荣-倩女幽魂 (电影《倩女幽魂》主题曲)-2020-08-05上午9-56-48.mp3', '', '5', '', '5', '11'), ('22', '2020-08-05 09:57:16', '2020-08-05 09:57:16', null, '张国荣 - 春夏秋冬 [mqms2]', 'http://files.pandaleo.cn/张国荣 - 春夏秋冬 [mqms2]-2020-08-05上午9-57-04.mp3', '', '5', '', '5', '12'), ('23', '2020-08-05 09:57:46', '2020-08-05 09:57:46', null, '周杰倫 - 藉口', 'http://files.pandaleo.cn/周杰倫 - 藉口-2020-08-05上午9-57-38.mp3', '', '4', '', '5', '9'), ('24', '2020-08-05 09:58:09', '2020-08-05 09:58:09', null, '周杰倫 - 七里香', 'http://files.pandaleo.cn/周杰倫 - 七里香-2020-08-05上午9-57-55.mp3', '', '4', '', '5', '9'), ('25', '2020-08-05 09:58:23', '2020-08-05 09:58:23', null, '周杰倫 - 園游會', 'http://files.pandaleo.cn/周杰倫 - 園游會-2020-08-05上午9-58-14.mp3', '', '4', '', '5', '9');
COMMIT;

-- ----------------------------
--  Table structure for `iris_songs_playlists`
-- ----------------------------
DROP TABLE IF EXISTS `iris_songs_playlists`;
CREATE TABLE `iris_songs_playlists` (
  `song_id` int unsigned NOT NULL,
  `playlist_id` int unsigned NOT NULL,
  PRIMARY KEY (`song_id`,`playlist_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Table structure for `iris_streams`
-- ----------------------------
DROP TABLE IF EXISTS `iris_streams`;
CREATE TABLE `iris_streams` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `url` varchar(256) DEFAULT NULL,
  `custom_path` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `url` (`url`),
  KEY `idx_iris_streams_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Table structure for `iris_users`
-- ----------------------------
DROP TABLE IF EXISTS `iris_users`;
CREATE TABLE `iris_users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_iris_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
--  Records of `iris_users`
-- ----------------------------
BEGIN;
INSERT INTO `iris_users` VALUES ('5', '2020-08-04 10:17:26', '2020-08-04 10:17:26', null, 'name', 'username', '$2a$10$V.5X756TEZx228Ui9LTzquaNuyn6LZSf2y.7xuSwIWQ1YsMJ46kFy');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
