# namespace sessionmgr_mysql

create database if not exists ldzpb_session;

use ldzpb_session;

create table if not exists t_session_info (
	id bigint primary key auto_increment,
 	sessionid varchar(64),
 	timeouttime bigint,
 	-- 有效期
 	losevaildtime bigint,
 	-- 用户自定义字段
 	userdata varchar(128)
);

-- 添加 session
drop procedure if exists pro_add_session;
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] timeoutTime: uint64;
	[in] losevaildTime: uint64;
	[in] userdata: string;
	[out] sessionId: string
*/
delimiter $$
create procedure pro_add_session(
in in_timeouttime bigint,
in in_losevaildtime bigint,
in in_userdata varchar(128)
)
begin
	declare var_sessionid varchar(64) default null;
	start transaction;
	select uuid() into var_sessionid;
	insert into t_session_info values(null, var_sessionid, in_timeouttime, in_losevaildtime, in_userdata);
	select var_sessionid;
	commit;
end
$$
delimiter ;

-- call pro_add_session(10 * 60 * 1000, 1000);

-- 删除 session
drop procedure if exists pro_delete_session;
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] sessionId: string;
*/
delimiter $$
create procedure pro_delete_session(
in in_sessionid varchar(64)
)
begin
	start transaction;
	delete from t_session_info where sessionid = in_sessionid;
	commit;
end
$$
delimiter ;

-- 判断session是否存在
drop procedure if exists pro_session_isexist;
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] sessionId: string;
	[out] isexist: bool
*/
delimiter $$
create procedure pro_session_isexist(
in in_sessionid varchar(64)
)
begin
	declare var_count int default 0;

	select count(*) into var_count from t_session_info
	where sessionid = in_sessionid limit 1;

	if var_count = 0 then
		select 0;
	else
		select 1;
	end if;
end
$$
delimiter ;

-- 更新失效时间
drop procedure if exists pro_update_losevaildtime;
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] sessionId: string;
	[in] nowTimeStamp: uint64
*/
delimiter $$
create procedure pro_update_losevaildtime(
in in_sessionid varchar(64),
in in_nowtimestamp bigint
)
begin
	declare var_timeout bigint default 0;

	start transaction;
	select timeouttime into var_timeout from t_session_info
	where sessionid = in_sessionid;

	update t_session_info set losevaildtime = in_nowtimestamp + var_timeout
	where sessionid = in_sessionid;
	commit;
end
$$
delimiter ;

-- 根据sessionId获取session信息
drop procedure if exists pro_get_sessioninfo_by_sessionid;
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] sessionId: string;
	[out] id: uint64;
	[out] sessionId: string;
	[out] timeoutTime: uint64;
	[out] losevaildTime: uint64;
	[out] userdata: string
*/
delimiter $$
create procedure pro_get_sessioninfo_by_sessionid(
in in_sessionid varchar(64)
)
begin
	select * from t_session_info
	where sessionid = in_sessionid;
end
$$
delimiter ;

-- 传入当前的时间戳, 在数据库中查找所有session的时间, 找出失效时间在给定的时间之前的, 并删除
drop procedure pro_delete_losevaild_sessions;
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] nowTimeStamp: uint64;
*/
delimiter $$
create procedure pro_delete_losevaild_sessions(
in in_nowtimestamp bigint
)
begin
	start transaction;
	delete from t_session_info
	where id in (
	select tmp.id from (
	select id from t_session_info
	where losevaildtime < in_nowtimestamp
	) as tmp
	);
	commit;
end
$$
delimiter ;

