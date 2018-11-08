# namespace collectionmgr

create database if not exists ldzpb_dynamic_config;

use ldzpb_dynamic_config;

-- 清空表
truncate t_commodity_collect_info;

create table if not exists t_commodity_collect_info(
uuid varchar(64),
useruuid varchar(64),
commodityuuid varchar(64)

