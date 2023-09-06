create table tiktok.`user`
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
    constraint `id`
        primary key (`id`)
) engine=InnoDB auto_increment=10000 default charset=utf8mb4;

create table tiktok.`video`
(
    `id`              bigint                              not null,
    `user_id`         bigint                              not null,
    `play_url`        varchar(255)                        not null comment 'url',
    `cover_url`       varchar(255)                        not null comment 'url',
    `title`           varchar(255)                        not null,
    `created_at`      timestamp default current_timestamp not null,
    `updated_at`      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    `deleted_at`      timestamp default null null,
    constraint `id`
        primary key (`id`),
    constraint `video_user`
        foreign key (`user_id`)
            references tiktok.`user` (`id`)
            on delete cascade
)engine=InnoDB default charset=utf8mb4;

create table tiktok.`favorite`
(
    `id`         bigint auto_increment not null,
    `user_id`    bigint                              not null,
    `video_id`   bigint                              not null,
    `status`     tinyint   default 1                 not null comment '1: like, 0: dislike',
    `created_at` timestamp default current_timestamp not null,
    `updated_at` timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    `deleted_at` timestamp default null null,
    constraint `id`
        primary key (`id`),
        index `uid_vid_idx` (`user_id`, `video_id`),
    constraint `favorite_user`
        foreign key (`user_id`)
            references tiktok.`user` (`id`)
            on delete cascade,
    constraint `favorite_video`
        foreign key (`video_id`)
            references tiktok.`video` (`id`)
            on delete cascade
) engine=InnoDB default charset=utf8mb4;

create table tiktok.`comment`
(
    `id`         bigint auto_increment not null,
    `user_id`    bigint                              not null,
    `video_id`   bigint                              not null,
    `content`    varchar(255)                        not null,
    `created_at` timestamp default current_timestamp not null,
    `updated_at` timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    `deleted_at` timestamp default null null,
    constraint `id`
        primary key (`id`),
        index (`video_id`),
        index (`deleted_at`),
        index (`created_at`),
    constraint `comment_user`
        foreign key (`user_id`)
            references tiktok.`user` (`id`)
            on delete cascade,
    constraint `comment_video`
        foreign key (`video_id`)
            references tiktok.`video` (`id`)
            on delete cascade
) engine=InnoDB default charset=utf8mb4;

create table tiktok.`follow`
(
    `id`          bigint auto_increment               not null,
    `user_id`     bigint                              not null comment 'user id',
    `to_user_id`  bigint                              not null comment 'target user id',
    `status`      tinyint    default 1                not null comment 'status',
    `created_at`  timestamp default current_timestamp not null,
    `updated_at`  timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    `deleted_at`  timestamp default null null,
    constraint `id`
        primary key (`id`),
    constraint `follow_user`
        foreign key (`user_id`)
            references tiktok.`user` (`id`)
            on delete cascade,
    constraint `follow_to_user`
        foreign key (`to_user_id`)
            references tiktok.`user` (`id`)
            on delete cascade
) engine=InnoDB default charset=utf8mb4;

create table tiktok.`message` (
    `id`              bigint          not null,
    `to_user_id`      bigint          not null comment 'target user id',
    `from_user_id`    bigint          not null comment 'user id',
    `content`         varchar(4000)   not null comment 'message content',
    `created_at`      timestamp       not null        default current_timestamp,
    `updated_at`      timestamp       not null        default current_timestamp on update current_timestamp,
    `deleted_at`      timestamp       null            default null,
    constraint `id`
        primary key (`id`),
        index(`to_user_id`),
        index(`from_user_id`),
    constraint `to_user`
        foreign key (`to_user_id`)
            references tiktok.`user` (`id`)
            on delete cascade,
    constraint `from-user`
        foreign key (`from_user_id`)
            references tiktok.`user` (`id`)
            on delete cascade
) engine=InnoDB default charset=utf8mb4;