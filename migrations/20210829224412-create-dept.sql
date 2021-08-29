-- +migrate Up
CREATE TABLE `sf_dept` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
     `name` varchar(191) NOT NULL COMMENT '部门名称',
     `pid` int(11) unsigned DEFAULT NULL COMMENT '父级部门ID，为0则是顶级部门',
     `level` tinyint(4) unsigned NOT NULL COMMENT '部门层级：1->2->3逐次降低',
     `sort` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '排序，数字越小越靠前',
     `remark` varchar(191) NOT NULL DEFAULT '' COMMENT '备注描述信息',
     `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
     `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
     PRIMARY KEY (`id`),
     KEY `pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='多级部门表';

-- +migrate Down
DROP table if exists `sf_dept`;
