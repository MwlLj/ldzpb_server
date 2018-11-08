# namespace serverconfig

create database if not exists ldzpb_configs;

use ldzpb_configs

-- 清空表
truncate t_server_info;

create table if not exists t_server_info(
uuid varchar(64),
servername varchar(128),
servertype varchar(64),
serverip varchar(64),
serverport int,
serverdomainname varchar(128)
);
create unique index UK_serverinfo_servertype on t_server_info (servertype);

-- procedures
drop procedure if exists pro_add_server_info;
-- 添加服务信息
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] servername: string;
	[in] servertype: string;
	[in] serverip: string;
	[in] serverport: int;
	[in] serverdomainname: string;
	[out] serveruuid: string
*/
delimiter $$
create procedure pro_add_server_info(
in in_servername varchar(128),
in in_servertype varchar(64),
in in_serverip varchar(64),
in in_serverport int,
in in_serverdomainname varchar(128)
)
begin
	declare var_uuid varchar(64) default null;

	select uuid() into var_uuid;

	start transaction;
	insert into t_server_info values(var_uuid, in_servername, in_servertype, in_serverip, in_serverport, in_serverdomainname);
	commit;

	select var_uuid;
end
$$
delimiter ;

drop procedure if exists pro_get_server_info;
-- 获取服务信息
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] servertype: string;
	[out] serveruuid: string;
	[out] servername: string;
	[out] servertype: string;
	[out] serverip: string;
	[out] serverport: int;
	[out] serverdomainname: string
*/
delimiter $$
create procedure pro_get_server_info(
in in_servertype varchar(64)
)
begin
	select * from t_server_info where servertype = in_servertype;
end
$$
delimiter ;
