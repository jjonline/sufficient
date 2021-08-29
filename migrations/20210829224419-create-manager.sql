-- +migrate Up
CREATE TABLE `sf_manager` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(32) NOT NULL COMMENT '姓名：仅显示用',
    `account` varchar(32) NOT NULL COMMENT '账号：字母和数字构成不区分大小写',
    `password` varchar(255) NOT NULL COMMENT '密码（密文）',
    `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号码：有则可作为账号代码层面控制唯一',
    `email` varchar(128) NOT NULL DEFAULT '' COMMENT '电子邮箱：有则可作为账号代码层面控制唯一',
    `is_root` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为根用户：1是0不是[根用户不受角色限制永远具备所有菜单的所有权限，只有根用户才能创建根用户]',
    `enable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '启用禁用标记：1启用0禁用',
    `remark` varchar(191) NOT NULL DEFAULT '' COMMENT '备注描述信息',
    `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `account` (`account`),
    KEY `mobile` (`mobile`),
    KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='管理员表';

-- +migrate Down
DROP table if exists `sf_manager`;
