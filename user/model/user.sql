create table users
(
    id                 bigint unsigned auto_increment,
    name               varchar(255) not null,
    phone              varchar(255) null,
    email              varchar(255) null,
    email_verified_at  timestamp    null,
    password           varchar(255) null,
    weixin_openid      varchar(255) null,
    weixin_unionid     varchar(255) null,
    remember_token     varchar(100) null,
    created_at         timestamp    null default current_timestamp,
    updated_at         timestamp    null default current_timestamp on update current_timestamp,
    avatar             varchar(255) null,
    introduction       varchar(255) null,
    notification_count int unsigned      default '0' not null,
    last_actived_at    timestamp    null,
    registration_id    varchar(255) null,
    primary key (id),
    constraint users_email_unique
        unique (email),
    constraint users_phone_unique
        unique (phone),
    constraint users_weixin_openid_unique
        unique (weixin_openid),
    constraint users_weixin_unionid_unique
        unique (weixin_unionid)
)
    collate = utf8mb4_unicode_ci;

