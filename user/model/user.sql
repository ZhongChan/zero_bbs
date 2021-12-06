CREATE TABLE `users`
(
    `id`                 bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `name`               varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `phone`              varchar(255) COLLATE utf8mb4_unicode_ci                       DEFAULT NULL,
    `email`              varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `email_verified_at`  timestamp                               NULL                  DEFAULT NULL,
    `password`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `weixin_openid`      varchar(255) COLLATE utf8mb4_unicode_ci                       DEFAULT NULL,
    `weixin_unionid`     varchar(255) COLLATE utf8mb4_unicode_ci                       DEFAULT NULL,
    `remember_token`     varchar(100) COLLATE utf8mb4_unicode_ci                       DEFAULT NULL,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `avatar`             varchar(255) COLLATE utf8mb4_unicode_ci                       DEFAULT NULL,
    `introduction`       varchar(255) COLLATE utf8mb4_unicode_ci                       DEFAULT NULL,
    `notification_count` int unsigned                            NOT NULL              DEFAULT '0',
    `last_actived_at`    timestamp                               NULL                  DEFAULT NULL,
    `registration_id`    varchar(255) COLLATE utf8mb4_unicode_ci                       DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_email_unique` (`email`),
    UNIQUE KEY `users_phone_unique` (`phone`),
    UNIQUE KEY `users_weixin_openid_unique` (`weixin_openid`),
    UNIQUE KEY `users_weixin_unionid_unique` (`weixin_unionid`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci

