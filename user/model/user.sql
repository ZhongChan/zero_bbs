create table users
(
    id                 int unsigned auto_increment,
    name               varchar(255)             not null,
    email              varchar(255)             not null,
    email_verified_at  timestamp                null,
    password           varchar(255)             not null,
    remember_token     varchar(100)             null,
    created_at         timestamp                null,
    updated_at         timestamp                null,
    avatar             varchar(255)             null,
    introduction       varchar(255)             null,
    notification_count int unsigned default '0' not null,
    last_actived_at    timestamp                null,
    primary key (id),
    constraint users_email_unique unique (email)
) collate = utf8mb4_unicode_ci;

