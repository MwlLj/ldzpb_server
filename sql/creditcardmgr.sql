# namespace creditcardmgr

create database if not exists ldzpb_configs;

use ldzpb_configs;

-- 清空表
truncate t_credit_card_info;

create table if not exists t_credit_card_info(
uuid varchar(64),
useruuid varchar(64),
accountnumber varchar(64),
cardbelongname varchar(64),
cardbelongbank varchar(64),
cardvaildtime datetime
)
