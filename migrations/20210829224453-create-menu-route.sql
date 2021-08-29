-- +migrate Up
CREATE TABLE `sf_menu_route` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `menu_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
    `method` varchar(10) NOT NULL COMMENT 'http请求方法',
    `path` varchar(191) NOT NULL COMMENT '后端api菜单路由',
    `sort` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '排序，数字越小越靠前',
    `remark` varchar(191) NOT NULL DEFAULT '' COMMENT '备注描述信息',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单所有路由表';

-- +migrate Down
DROP table if exists `sf_menu_route`;