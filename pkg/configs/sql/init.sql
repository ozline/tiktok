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