create table `forum_topic_t`(
    topic_id bigint(20) NOT NULL,
    uid bigint(20) NOT NULL,
    reply_total int(11) not null default '0',
    last_reply_uid bigint(20) NOT NULL default '0',
    last_reply_ts datetime not null default '0',
    create_ts datetime not null default '0',
    status int(11) not null default '0',
    primary key (`topic_id`),
    key idx_uid (`uid`)
) ENGINE = InnoDB CHARSET = utf8mb4 COLLATE utf8mb4_bin;

create table `forum_content_t`(
    id bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    topic_id bigint(20) NOT NULL,
    content varchar(20000) not null default '',
    primary key (`id`),
    unique uniq_topic_id (`topic_id`)
) ENGINE = InnoDB CHARSET = utf8mb4 COLLATE utf8mb4_bin;

create table `forum_reply_t`(
    id bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    uid bigint(20) NOT NULL,
    reply_id bigint(20) NOT NULL,
    topic_id bigint(20) NOT NULL,
    content varchar(500) not null default '',
    reply_ts datetime not null default '0',
    primary key (`id`),
    key idx_uid (`uid`),
    unique uniq_reply_id (`reply_id`),
    unique uniq_topic_id (`topic_id`)
) ENGINE = InnoDB CHARSET = utf8mb4 COLLATE utf8mb4_bin;

create table `forum_notify_t`(
    id bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    uid bigint(20) NOT NULL,
    reply_id bigint(20) NOT NULL,
    topic_id bigint(20) NOT NULL,
    reply_ts datetime not null default '0',
    status int(11) not null default '0' comment '消息状态 1-未读 2-已读',
    primary key (`id`),
    key idx_uid (`uid`)
) ENGINE = InnoDB CHARSET = utf8mb4 COLLATE utf8mb4_bin;
