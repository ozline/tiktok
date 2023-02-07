create table user
(
    id             bigint                              not null comment 'user_id',
    username       text                                not null comment 'username',
    password       text                                not null comment 'password',
    follow_count   int       default 0                 not null comment 'follow count',
    follower_count int       default 0                 not null comment 'follower count',
    create_at      timestamp default current_timestamp not null comment 'create time',
    constraint id
        primary key (id)
);