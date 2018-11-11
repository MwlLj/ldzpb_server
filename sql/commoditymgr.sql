#namespace commoditymgr

create database if not exists ldzpb_commodity_ress;

use ldzpb_commodity_ress;

-- 清空表
truncate t_commodity_classifition;
truncate t_object_detail;
truncate t_commodity_info;

#create tables
/*
create table if not exists t_object_detail(
uuid varchar(64),
-- text / picture / video
detailtype varchar(64),
-- desc / picurl / ...
detailvalue text,
detailno int,
objectuuid varchar(64)
);
create index UK_objectdetail_uuid on t_object_detail (uuid);

create table if not exists t_commodity_classifition(
uuid varchar(64),
parentuuid varchar(64),
name varchar(256)
);
create index UK_commodityclassifition_uuid on t_commodity_classifition (uuid);

create table if not exists t_commodity_info(
uuid varchar(64),
name varchar(256),
price decimal(10, 2),
classifyuuid varchar(64)
);
create index UK_commodityinfo_uuid on t_commodity_info (uuid);
*/
#end

/*
	@bref 添加商品分类信息
	@is_brace true
	@in_isarr true
	@out_isarr false
	@in uuid: string
	@in parentUuid: string
	@in name: string
*/
#define addCommodityClassifition
insert into t_commodity_classifition values({0}, {1}, {2});
#end

/*
	@添加商品分类详情信息
	@is_brace true
	@in_isarr true
	@out_isarr false
	@in detailUuid: string
	@in detailType: string
	@in detailValue: string
	@in detailNo: int
	@in classifyUuid: string
*/
#define addCommodityClassifitionDetailInfo
insert into t_object_detail values({0}, {1}, {2}, {3}, {4});
#end

/*
	@bref 添加商品信息
	@is_brace true
	@in_isarr true
	@out_isarr false
	@in uuid: string
	@in name: string
	@in price: string
	@in classifyUuid: string
*/
#define addCommodityInfo
insert into t_commodity_info values({0}, {1}, {2}, {3});
#end

/*
	@bref 添加商品详情信息
	@is_brace true
	@in_isarr true
	@out_isarr false
	@in uuid: string
	@in detailType: string
	@in detailValue: string
	@in detialNo: int
	@in commodityUuid: string
*/
#define addCommodityDetailInfo
insert into t_object_detail values({0}, {1}, {2}, {3}, {4});
#end

-- procedure
drop procedure if exists pro_add_commodity_classifition;
-- 添加商品分类信息
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] parentUuid: string;
	[in] name: string;
	[out] uuid: string
*/
delimiter $$
create procedure pro_add_commodity_classifition(
in in_parentuuid varchar(64),
in in_name varchar(256)
)
begin
	insert into t_commodity_classifition values(uuid(), in_parentuuid, in_name);
end
$$
delimiter ;

drop procedure if exists pro_add_commodity_classifition_detail;
-- 添加商品分类详情信息
/*
	[in_isarr]: true;
	[out_isarr]: false;
	[in] classifitionUuid: string;
	[in] detailType: string;
	[in] detailValue: string;
	[in] detailNo: int
*/
delimiter $$
create procedure pro_add_commodity_classifition_detail(
in in_classifitionUuid varchar(64),
in in_detailType varchar(64),
in in_detailValue text,
in in_detailNo int
)
begin
	declare total int default 0;

	select count(*) into total from t_commodity_classifition where uuid = in_classifitionUuid;
	if total = 0 then
		-- classifition is not exist
		signal sqlstate "HY000" set message_text = "commodity classifition is not exist";
	end if;

	insert into t_object_detail values(uuid(), in_detailType, in_detailValue, in_detailNo, in_classifitionUuid);
end
$$
delimiter ;

drop procedure if exists pro_add_commodity;
/*
	[in_isarr]: false;
	[out_isarr]: false;
	[in] name: string;
	[in] price: double;
	[in] classifitionUuid: string
*/
-- 添加商品信息
delimiter $$
create procedure pro_add_commodity(
in in_name varchar(256),
in in_price decimal(10, 2),
in in_classifyuuid varchar(64)
)
begin
	insert into t_commodity_info values(uuid(), in_name, in_price, in_classifyuuid);
end
$$
delimiter ;

drop procedure if exists pro_add_commodity_detail;
-- 添加商品详情信息
/*
	[in_isarr]: true;
	[out_isarr]: false;
	[in] commodityUuid: string;
	[in] detailType: string;
	[in] detailValue: string;
	[in] detailNo: int
*/
delimiter $$
create procedure pro_add_commodity_detail(
in in_commodityUuid varchar(64),
in in_detailType varchar(64),
in in_detailValue text,
in in_detailNo int
)
begin
	declare total int default 0;

	select count(*) into total from t_commodity_info where uuid = in_commodityUuid;
	if total = 0 then
		-- commodity is not exist
		signal sqlstate "HY000" set message_text = "commodity is not exist";
	end if;

	insert into t_object_detail values(uuid(), in_detailType, in_detailValue, in_detailNo, in_commodityUuid);
end
$$
delimiter ;
