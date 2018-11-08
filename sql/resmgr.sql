# namespace resmgr

create database if not exists ldzpb_media_ress;

use ldzpb_media_ress;

-- 清空表
truncate t_resource_info;

create table if not exists t_resource_info(
resuuid varchar(64) primary key,
url varchar(256)
)
