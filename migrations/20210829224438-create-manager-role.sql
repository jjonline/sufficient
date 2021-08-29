-- +migrate Up
CREATE TABLE `sf_manager_role` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `manager_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
    `role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
    `sort` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '排序，数字越小越靠前',
    `remark` varchar(191) NOT NULL DEFAULT '' COMMENT '备注描述信息',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `manager_id` (`manager_id`),
    KEY `role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='管理员所属角色表';

-- +migrate Down
DROP table if exists `sf_manager_role`;
