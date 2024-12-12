USE mqtt;

CREATE TABLE `user` (
    `uuid` BINARY(32) NOT NULL COMMENT 'uuid',
    `id` INT NOT NULL COMMENT 'id',
    `username` VARCHAR(64) NOT NULL COMMENT '用户名',
    `password` VARCHAR(128) NOT NULL COMMENT '密码',
    `mobile` VARCHAR(128) DEFAULT NULL COMMENT '手机号',
    `email` VARCHAR(64) DEFAULT NULL COMMENT '邮箱',
    `status` INT NOT NULL DEFAULT 0 COMMENT '身份',
    PRIMARY KEY (`uuid`)
)