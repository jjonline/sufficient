-- +migrate Up
CREATE TABLE `sf_role_menu` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
    `menu_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `role_id` (`role_id`),
    KEY `menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色所属菜单表';

-- +migrate Down
DROP table if exists `sf_role_menu`;
