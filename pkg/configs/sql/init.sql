create table user
(
    user_id        bigint                              not null comment 'user_id',
    username       text                                not null comment 'username',
    password       text                                not null comment 'password',
    follow_count   int       default 0                 not null comment 'follow count',
    follower_count int       default 0                 not null comment 'follower count',
    create_date    timestamp                           not null comment 'create time',
    update_at      timestamp default current_timestamp not null on update current_timestamp comment 'update profile time',
    delete_at      timestamp default null              null comment 'user delete time',
    constraint user_id
        primary key (user_id)
);