/* user service */
create table user
(
    id              bigint                              not null comment 'user_id',
    username        text                                not null comment 'username',
    password        text                                not null comment 'password',
    follow_count    int       default 0                 not null comment 'follow count',
    follower_count  int       default 0                 not null comment 'follower count',
    created_at      timestamp default current_timestamp not null comment 'create time',
    updated_at      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    deleted_at      timestamp default null              null comment 'user delete time',
    constraint id
        primary key (id)
);

/* video service */
create table video
(
    id              bigint                              not null comment 'video id',
    user_id         bigint                              not null comment 'user id',
    title           text                                not null comment 'videotitle',
    play_url        text                                not null comment 'play URL',
    cover_url       text                                not null comment 'cover URL',
    created_at      timestamp default current_timestamp not null comment 'create time',
    updated_at      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    deleted_at      timestamp default null              null comment 'delete time',
    constraint id
        primary key (id)
);

/* follow service */
create table follow
(
    user_id         bigint                              not null comment 'user id',
    to_user_id      bigint                              not null comment 'follow user id',
)

/* comment service */
create table comment
(
    id                  bigint                                  not null comment 'comment id',
    user_id             bigint                                  not null comment 'commentor',
    video_id            bigint                                  not null comment 'video id',
    content             text                                    not null comment 'comment content',
    is_uploader_like    bool        default false               not null comment 'uploader like',
    like_count          int         default 0                   not null comment 'like count',
    created_at          timestamp   default current_timestamp   not null comment 'create time',
    deleted_at          timestamp   default null                null comment 'delete time',
    constraint id
        primary key (id)
)

create table comment_like
(
    user_id             bigint                                  not null comment 'user id',
    comment_id          bigint                                  not null comment 'comment id',
)

create table video_favorite
(
    user_id             bigint                                  not null comment 'user id',
    video_id            bigint                                  not null comment 'video id',
)

/* chat service */
create table message
(
    id                  bigint                                 not null comment 'message id',
    from_user_id        bigint                                 not null comment 'from user id',
    to_user_id          bigint                                 not null comment 'to user id',
    content             text                                   not null comment 'message content',
    created_time        text       default current_timestamp   not null comment 'create time',
    constraint id
        primary key (id)
)