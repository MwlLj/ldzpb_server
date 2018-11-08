# namespace usermgr

create database if not exists ldzpb_configs;

use ldzpb_configs;

-- 清空表
truncate t_user_info;
truncate t_user_secrecy_info;
truncate t_user_email_info;
truncate t_user_phone_info;

create table if not exists t_user_info (
user_id bigint primary key auto_increment,
user_uuid varchar(64),
user_secrecy_info_id bigint
);
alter table t_user_info add unique (user_uuid);

create table if not exists t_user_secrecy_info (
user_secrecy_info_id bigint primary key auto_increment,
user_name varchar(128),
user_pwd varchar(128)
);

create table if not exists t_user_email_info (
user_email_id bigint primary key auto_increment,
user_secrecy_info_id bigint,
user_email_no varchar(128)
);
alter table t_user_email_info add unique (user_email_no);

create table if not exists t_user_phone_info (
user_phone_id bigint primary key auto_increment,
user_secrecy_info_id bigint,
user_phone_no varchar(128)
);
alter table t_user_phone_info add unique (user_phone_no);

-- procedures
drop procedure if exists pro_user_isexists_by_email;
-- 判断用户是否存在
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] userEmail: string;
	[out] userUuid: string
*/
delimiter $$
create procedure pro_user_isexists_by_email(
in in_user_email_no varchar(128)
)
begin
	declare var_uuid varchar(64) default null;

	select ui.user_uuid into var_uuid from t_user_info as ui
	inner join
	(select user_secrecy_info_id as usii from t_user_email_info as uei
	where uei.user_email_no = in_user_email_no) as email_result
	on ui.user_secrecy_info_id = email_result.usii;

	select var_uuid;
end
$$
delimiter ;

-- call pro_user_isexists_by_email("731025894@qq.com");

drop procedure if exists pro_user_isexists_by_phone;
-- 判断用户是否存在
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] userPhone: string;
	[out] userUuid: string
*/
delimiter $$
create procedure pro_user_isexists_by_phone(
in in_user_phone_no varchar(128)
)
begin
	declare var_uuid varchar(64) default null;

	select ui.user_uuid into var_uuid from t_user_info as ui
	inner join
	(select user_secrecy_info_id as usii from t_user_phone_info as upi
	where upi.user_phone_no = in_user_phone_no) as phone_result
	on ui.user_secrecy_info_id = phone_result.usii;

	select var_uuid;
end
$$
delimiter ;

-- call pro_user_isexists_by_phone("15771342867");

drop procedure if exists pro_user_isexists_by_email_or_phone;
-- 根据邮箱或者手机号码判断用户是否存在
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] userNo: string;
	[out] userUuid: string
*/
delimiter $$
create procedure pro_user_isexists_by_email_or_phone(
in in_no varchar(128)
)
begin
	declare var_uuid varchar(64) default null;

	select ui.user_uuid into var_uuid from t_user_info as ui
	inner join
	(select user_secrecy_info_id as usii from t_user_email_info as uei
	where uei.user_email_no = in_no) as email_result
	on ui.user_secrecy_info_id = email_result.usii;

	if var_uuid is null then
		select ui.user_uuid into var_uuid from t_user_info as ui
		inner join
		(select user_secrecy_info_id as usii from t_user_phone_info as upi
		where upi.user_phone_no = in_no) as phone_result
		on ui.user_secrecy_info_id = phone_result.usii;
	end if;

	select var_uuid;
end
$$
delimiter ;

-- call pro_user_isexists_by_email_or_phone("15771342867");

drop procedure if exists pro_add_user;
-- 添加一个用户
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] userName: string;
	[in] userPwd: string;
	[in] userEmailNo: string;
	[in] userPhoneNo: string;
	[out] userUuid: string
*/
delimiter $$
create procedure pro_add_user(
in in_user_name varchar(128),
in in_user_pwd varchar(128),
in in_user_email_no varchar(128),
in in_user_phone_no varchar(128)
)
begin
	declare var_secrecy_info_id bigint default 0;
	declare var_uuid varchar(64) default null;

	start transaction;

	insert into t_user_secrecy_info values(null, in_user_name, in_user_pwd);
	select last_insert_id() into var_secrecy_info_id;

	insert into t_user_email_info values(null, var_secrecy_info_id, in_user_email_no);
	insert into t_user_phone_info values(null, var_secrecy_info_id, in_user_phone_no);

	select uuid() into var_uuid;
	insert into t_user_info values(null, var_uuid, var_secrecy_info_id);

	select var_uuid;

	commit;
end
$$
delimiter ;

-- call pro_add_user("admin", "123456", "731025894@qq.com", "15771342867");

drop procedure if exists pro_password_istrue_by_useruuid;
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] userUuid: string;
	[in] userPwd: string;
	[out] isTrue: bool
*/
-- 根据用户uuid判断密码是否正确
delimiter $$
create procedure pro_password_istrue_by_useruuid(
in in_user_uuid varchar(64),
in in_user_pwd varchar(128)
)
begin
	declare var_user_pwd varchar(128) default null;

	select user_pwd into var_user_pwd
	from t_user_secrecy_info as usi
	inner join
	(select user_secrecy_info_id as usii
	from t_user_info as ui
	where ui.user_uuid = in_user_uuid) as secrecy_result
	on usi.user_secrecy_info_id = secrecy_result.usii;

	if var_user_pwd != in_user_pwd then
		select 0;
	else
		select 1;
	end if;
end
$$
delimiter ;

-- call pro_password_istrue_by_useruuid("0874a2a3-a378-11e8-8483-00e01a680285", "123456");

drop procedure if exists pro_password_istrue_by_no;
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] no: string;
	[in] userPwd: string;
	[out] isTrue: bool;
	[out] userUuid: string
*/
-- 通过邮箱或手机号码判断用户密码是否正确
delimiter $$
create procedure pro_password_istrue_by_no(
in in_no varchar(128),
in in_user_pwd varchar(128)
)
begin
	declare var_user_uuid varchar(128) default null;
	declare var_user_pwd varchar(128) default null;
	declare var_pwd_is_true int default 0;

	select ui.user_uuid, t.user_pwd into var_user_uuid, var_user_pwd
	from t_user_info as ui
	inner join
	(select usi.user_secrecy_info_id as user_secrecy_id, usi.user_pwd as user_pwd
	from t_user_secrecy_info as usi
	inner join t_user_email_info as uei
	on usi.user_secrecy_info_id = uei.user_secrecy_info_id
	where uei.user_email_no = in_no) as t
	on ui.user_secrecy_info_id = t.user_secrecy_id;

	if var_user_uuid is null then
		select ui.user_uuid, t.user_pwd into var_user_uuid, var_user_pwd
		from t_user_info as ui
		inner join
		(select usi.user_secrecy_info_id as user_secrecy_id, usi.user_pwd as user_pwd
		from t_user_secrecy_info as usi
		inner join t_user_phone_info as uei
		on usi.user_secrecy_info_id = uei.user_secrecy_info_id
		where uei.user_phone_no = in_no) as t
		on ui.user_secrecy_info_id = t.user_secrecy_id;
	end if;

	if var_user_pwd is not null and var_user_pwd = in_user_pwd then
		set var_pwd_is_true = 1;
	end if;

	select var_pwd_is_true, var_user_uuid;
end
$$
delimiter ;

-- call pro_password_istrue_by_no("731025894@qq.com", "e10adc3949ba59abbe56e057f20f883e");

-- 通过用户UUID删除用户
drop procedure if exists pro_delete_user;
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] userUuid: string
*/
delimiter $$
create procedure pro_delete_user(
in in_useruuid varchar(64)
)
begin
	declare var_secrecy_info_id bigint default null;

	select user_secrecy_info_id into var_secrecy_info_id from t_user_info
	where user_uuid = in_useruuid;

	if var_secrecy_info_id = null then
	-- error
	select * from unknow;
	end if;

	start transaction;
	delete from t_user_info where user_uuid = in_useruuid;
	delete from t_user_secrecy_info where user_secrecy_info_id = var_secrecy_info_id;
	delete from t_user_email_info where user_secrecy_info_id = var_secrecy_info_id;
	delete from t_user_phone_info where user_secrecy_info_id = var_secrecy_info_id;
	commit;
end
$$
delimiter ;
