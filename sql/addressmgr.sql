# namespace addressmgr

create database if not exists ldzpb_configs;

use ldzpb_configs;

-- 清空表
truncate t_address_detail;
truncate t_address_info;

create table if not exists t_address_detail(
uuid varchar(64),
country varchar(64),
provice varchar(64),
city varchar(64),
town varchar(64),
countryside varchar(64),
housingestate varchar(64),
block varchar(64),
unit varchar(64),
floor varchar(64),
room varchar(64)
);

create table if not exists t_address_info(
uuid varchar(64),
useruuid varchar(64),
addrdetailuuid varchar(64),
addressdesc text
);
