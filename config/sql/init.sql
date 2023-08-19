create table tiktok.user
(
    `id`               bigint auto_increment not null,
    `username`         varchar(255)                                                                 not null unique,
    `password`         varchar(255)                                                                 not null,
    `avatar`           varchar(255) default 'https://files.ozline.icu/images/avatar.jpg'            not null comment 'url',
    `background_image` varchar(255) default 'https://files.ozline.icu/images/BannerImg_221116.jpeg' not null comment 'url',
    `signature`        varchar(255) default 'NOT NULL BUT SEEMS NULL'                               not null comment '255charmax',
    `created_at`       timestamp    default current_timestamp                                       not null,
    `updated_at`       timestamp    default current_timestamp                                       not null on update current_timestamp comment 'update profile time',
    `deleted_at`       timestamp    default null null,
    constraint id
        primary key (id)
) engine=InnoDB auto_increment=10000 default charset=utf8mb4;

/*
TODO: follow_count, follower_count, is_follow, work_count, favorite_count, total_favorited(string)
很迷的, 我觉得很多东西不需要
*/

create table tiktok.favorite
(
    `id`         bigint auto_increment not null,
    `user_id`    bigint                              not null,
    `video_id`   bigint                              not null,
    `status`     tinyint   default 1                 not null comment '1: like, 0: dislike',
    `created_at` timestamp default current_timestamp not null,
    `updated_at` timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    `deleted_at` timestamp default null null,
    constraint id
        primary key (id)
) engine=InnoDB auto_increment=10000 default charset=utf8mb4;

create table tiktok.comment
(
    `id`         bigint auto_increment not null,
    `user_id`    bigint                              not null,
    `video_id`   bigint                              not null,
    `content`    varchar(255)                        not null,
    `created_at` timestamp default current_timestamp not null,
    `updated_at` timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    `deleted_at` timestamp default null null,
    constraint id
        primary key (id)
) engine=InnoDB auto_increment=10000 default charset=utf8mb4;

create table tiktok.video
(
    `id`              bigint auto_increment               not null,
    `user_id`         bigint                              not null,
    `favorite_count`  bigint   default 0                  not null,
    `comment_count`   bigint   default 0                  not null,
    `play_url`        varchar(255)                        not null comment 'url',
    `cover_url`       varchar(255)                        not null comment 'url',
    `title`           varchar(255)                        not null,
    `created_at`      timestamp default current_timestamp not null,
    `updated_at`      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    `deleted_at`      timestamp default null null,
    constraint id
        primary key (id)
)engine=InnoDB auto_increment=10000 default charset=utf8mb4;

create table tiktok.follow
(
    `id`          bigint auto_increment               not null,
    `status`      bigint    default 1                 not null,
    `user_id`     bigint                              not null comment 'user id',
    `to_user_id`  bigint                              not null comment 'target user id',
    `status`      bigint    default 1                 not null comment 'status',
    `created_at`  timestamp default current_timestamp not null,
    `updated_at`  timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    `deleted_at`  timestamp default null null,
    constraint id
        primary key (id)
) engine=InnoDB auto_increment=10000 default charset=utf8mb4;

create table `tiktok`.`message`(
    `id`                  bigint auto_increment NOT NULL
    `to_user_id`          bigint NOT NULL
    `from_user_id`        bigint NOT NULL
    `content`             longtext NOT NULL
    `created_at`          timestamp     default current_timestamp                                        not null,
    `updated_at`          timestamp     default current_timestamp                                        not null on update current_timestamp comment 'update profile time',
    `deleted_at`          timestamp     default null
)engine=InnoDB auto_increment=10000 default charset=utf8mb4;
