-- +migrate Up
CREATE TABLE `sf_menu` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(60) NOT NULL COMMENT '菜单名称',
    `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级菜单编码',
    `level` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '级别 1->2->3逐次降低',
    `frontend` varchar(191) NOT NULL DEFAULT '' COMMENT '前端使用的标记',
    `icon` varchar(191) NOT NULL DEFAULT '' COMMENT '菜单图标',
    `sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单排序',
    `remark` varchar(191) NOT NULL DEFAULT '' COMMENT '备注描述信息',
    `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单表';

-- +migrate Down
DROP table if exists `sf_menu`;
