-- +migrate Up
CREATE TABLE `sf_role` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(191) NOT NULL COMMENT '角色名',
    `enable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '启用禁用标记：1启用0禁用',
    `sort` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '排序，数字越小越靠前',
    `remark` varchar(191) NOT NULL DEFAULT '' COMMENT '备注描述信息',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色表';

-- +migrate Down
DROP table if exists `sf_role`;
