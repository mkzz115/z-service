create table `wish_list_info_t`(
    id bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    uid bigint(20) NOT NULL,
    commodity_id bigint(20) NOT NULL,
    commodity_total int(11) not null default '0',
    add_ts datetime not null default '0',
    primary key (`id`),
    key idx_uid (`uid`)
) ENGINE = InnoDB CHARSET = utf8mb4 COLLATE utf8mb4_bin;

