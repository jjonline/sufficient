-- +migrate Up
CREATE TABLE `sf_manager_open` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `manager_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
    `open_name` varchar(64) NOT NULL DEFAULT '' COMMENT '开放平台账号的名称',
    `open_type` tinyint(4) NOT NULL COMMENT '开放平台类型，枚举值代码层面指定',
    `open_id` varchar(191) NOT NULL DEFAULT '' COMMENT '开放平台的OpenID',
    `access_token` varchar(512) NOT NULL DEFAULT '' COMMENT 'AccessToken',
    `token_expire_in` int(11) NOT NULL DEFAULT '0' COMMENT 'Token过期时间-linux时间戳',
    `refresh_token` varchar(512) NOT NULL DEFAULT '' COMMENT 'RefreshToken-大部分时候用不上',
    `figure` varchar(191) NOT NULL DEFAULT '' COMMENT '头像图src-可能需要本地化',
    `union_id` varchar(191) NOT NULL DEFAULT '' COMMENT 'UnionID',
    `enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '标记是否有效1-有效0-禁用',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `open_id` (`open_id`,`open_type`),
    KEY `manager_id` (`manager_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='管理员开放平台登录支持表';

-- +migrate Down
DROP table if exists `sf_manager_open`;
