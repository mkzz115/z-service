create table `user_account_t`(
    uid bigint(20) NOT NULL,
    account varchar(40) NOT NULL,
    pass varchar(20) not null,
    third_id varchar(40) NOT NULL default '',
    third_type int(11) not null default '0',
    primary key (`uid`),
    unique uniq_account (`account`)
) ENGINE = InnoDB CHARSET = utf8mb4 COLLATE utf8mb4_bin;

create table `user_info_t` (
    uid bigint(20) NOT NULL,
    name varchar(50) NOT NULL,
    email varchar(50) NOT NULL default '',
    phone varchar(22) NOT NULL default '',
    addr varchar(100) NOT NULL default '',
    image varchar(100) NOT NULL default '',
    reg_ts datetime NOT NULL default '1971-01-01 00:00:00',
    reg_type int(11) NOT NULL default '0',
    primary key (`uid`)
) ENGINE = InnoDB CHARSET = utf8mb4 COLLATE utf8mb4_bin;

create table `third_account_t` (
    id bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    uid bigint(20) NOT NULL,
    third_id varchar(40) NOT NULL,
    third_type int(11) not null,
    primary key (`id`),
    unique uniq_thirdid_type (`third_id`, `third_type`)
) ENGINE = InnoDB CHARSET = utf8mb4 COLLATE utf8mb4_bin;