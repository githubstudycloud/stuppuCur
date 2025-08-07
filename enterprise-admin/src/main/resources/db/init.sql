-- 创建数据库
CREATE DATABASE IF NOT EXISTS `enterprise` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `enterprise`;

-- 用户表
CREATE TABLE IF NOT EXISTS `sys_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(100) NOT NULL COMMENT '密码',
  `nickname` varchar(50) DEFAULT NULL COMMENT '昵称',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态（0：禁用，1：启用）',
  `gender` int(11) DEFAULT '0' COMMENT '性别（0：未知，1：男，2：女）',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门ID',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(50) DEFAULT NULL COMMENT '最后登录IP',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_by` varchar(50) DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(50) DEFAULT NULL COMMENT '更新人',
  `deleted` int(11) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识（0：未删除，1：已删除）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_email` (`email`),
  KEY `idx_phone` (`phone`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted` (`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 插入默认管理员用户（密码：admin123）
INSERT INTO `sys_user` (`username`, `password`, `nickname`, `email`, `status`, `create_by`) VALUES
('admin', '$2a$10$7JB720yubVSOfvVWdBYoOeymQxVqHhqHhqHhqHhqHhqHhqHhqHhqH', '系统管理员', 'admin@example.com', 1, 'system'),
('test', '$2a$10$7JB720yubVSOfvVWdBYoOeymQxVqHhqHhqHhqHhqHhqHhqHhqHhqH', '测试用户', 'test@example.com', 1, 'system');

-- 角色表
CREATE TABLE IF NOT EXISTS `sys_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(50) NOT NULL COMMENT '角色名称',
  `role_code` varchar(50) NOT NULL COMMENT '角色编码',
  `description` varchar(200) DEFAULT NULL COMMENT '角色描述',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态（0：禁用，1：启用）',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_by` varchar(50) DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(50) DEFAULT NULL COMMENT '更新人',
  `deleted` int(11) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识（0：未删除，1：已删除）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_code` (`role_code`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted` (`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 插入默认角色
INSERT INTO `sys_role` (`role_name`, `role_code`, `description`, `create_by`) VALUES
('超级管理员', 'SUPER_ADMIN', '系统超级管理员，拥有所有权限', 'system'),
('管理员', 'ADMIN', '系统管理员，拥有大部分权限', 'system'),
('普通用户', 'USER', '普通用户，拥有基本权限', 'system');

-- 用户角色关联表
CREATE TABLE IF NOT EXISTS `sys_user_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_role` (`user_id`, `role_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';

-- 插入默认用户角色关联
INSERT INTO `sys_user_role` (`user_id`, `role_id`) VALUES
(1, 1), -- admin -> 超级管理员
(2, 3); -- test -> 普通用户

-- 菜单表
CREATE TABLE IF NOT EXISTS `sys_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父菜单ID',
  `menu_name` varchar(50) NOT NULL COMMENT '菜单名称',
  `menu_code` varchar(50) NOT NULL COMMENT '菜单编码',
  `menu_type` varchar(10) NOT NULL COMMENT '菜单类型（M：目录，C：菜单，F：按钮）',
  `path` varchar(200) DEFAULT NULL COMMENT '路由地址',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `perms` varchar(100) DEFAULT NULL COMMENT '权限标识',
  `icon` varchar(100) DEFAULT NULL COMMENT '菜单图标',
  `sort_order` int(11) DEFAULT '0' COMMENT '排序',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态（0：禁用，1：启用）',
  `visible` int(11) NOT NULL DEFAULT '1' COMMENT '是否可见（0：隐藏，1：显示）',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_by` varchar(50) DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(50) DEFAULT NULL COMMENT '更新人',
  `deleted` int(11) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识（0：未删除，1：已删除）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_menu_code` (`menu_code`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted` (`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单表';

-- 插入默认菜单
INSERT INTO `sys_menu` (`parent_id`, `menu_name`, `menu_code`, `menu_type`, `path`, `component`, `perms`, `icon`, `sort_order`, `create_by`) VALUES
(0, '系统管理', 'system', 'M', '/system', NULL, NULL, 'system', 1, 'system'),
(1, '用户管理', 'system:user', 'C', '/system/user', 'system/user/index', 'system:user:list', 'user', 1, 'system'),
(2, '用户查询', 'system:user:query', 'F', NULL, NULL, 'system:user:query', NULL, 1, 'system'),
(2, '用户新增', 'system:user:add', 'F', NULL, NULL, 'system:user:add', NULL, 2, 'system'),
(2, '用户修改', 'system:user:edit', 'F', NULL, NULL, 'system:user:edit', NULL, 3, 'system'),
(2, '用户删除', 'system:user:remove', 'F', NULL, NULL, 'system:user:remove', NULL, 4, 'system'),
(1, '角色管理', 'system:role', 'C', '/system/role', 'system/role/index', 'system:role:list', 'role', 2, 'system'),
(7, '角色查询', 'system:role:query', 'F', NULL, NULL, 'system:role:query', NULL, 1, 'system'),
(7, '角色新增', 'system:role:add', 'F', NULL, NULL, 'system:role:add', NULL, 2, 'system'),
(7, '角色修改', 'system:role:edit', 'F', NULL, NULL, 'system:role:edit', NULL, 3, 'system'),
(7, '角色删除', 'system:role:remove', 'F', NULL, NULL, 'system:role:remove', NULL, 4, 'system'),
(1, '菜单管理', 'system:menu', 'C', '/system/menu', 'system/menu/index', 'system:menu:list', 'menu', 3, 'system'),
(12, '菜单查询', 'system:menu:query', 'F', NULL, NULL, 'system:menu:query', NULL, 1, 'system'),
(12, '菜单新增', 'system:menu:add', 'F', NULL, NULL, 'system:menu:add', NULL, 2, 'system'),
(12, '菜单修改', 'system:menu:edit', 'F', NULL, NULL, 'system:menu:edit', NULL, 3, 'system'),
(12, '菜单删除', 'system:menu:remove', 'F', NULL, NULL, 'system:menu:remove', NULL, 4, 'system');

-- 角色菜单关联表
CREATE TABLE IF NOT EXISTS `sys_role_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_menu` (`role_id`, `menu_id`),
  KEY `idx_role_id` (`role_id`),
  KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色菜单关联表';

-- 插入默认角色菜单关联（超级管理员拥有所有权限）
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES
(1, 1), (1, 2), (1, 3), (1, 4), (1, 5), (1, 6), (1, 7), (1, 8), (1, 9), (1, 10), (1, 11), (1, 12), (1, 13), (1, 14), (1, 15), (1, 16), (1, 17),
(2, 1), (2, 2), (2, 3), (2, 4), (2, 5), (2, 6), (2, 7), (2, 8), (2, 9), (2, 10), (2, 11), (2, 12), (2, 13), (2, 14), (2, 15), (2, 16), (2, 17),
(3, 1), (3, 2), (3, 3); -- 普通用户只有查询权限