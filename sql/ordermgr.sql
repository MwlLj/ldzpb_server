# namespace ordermgr

create database if not exists ldzpb_trades;

use ldzpb_trades;

-- 清空表
truncate t_order_commodity_info;
truncate t_order_info;
truncate t_order_pay_info;
truncate t_order_logistics_info;

create table if not exists t_order_commodity_info(
orderuuid varchar(64),
commodityuuid varchar(64)
);

create table if not exists t_order_info(
uuid varchar(64),
useruuid varchar(64),
customaddruuid varchar(64),
-- 运费
freight decimal(10, 2),
-- 发货状态
deliverstate tinyint,
-- 收货状态
recvicestate tinyint,
payuuid varchar(64),
logisticsuuid varchar(64)
);

-- 订单支付信息
create table if not exists t_order_pay_info(
uuid varchar(64),
paymode varchar(64),
paystate tinyint,
carduuid varchar(64)
);

-- 订单物流信息
create table if not exists t_order_logistics_info(
uuid varchar(64),
ordernumber varchar(64),
companyname varchar(64),
companycode varchar(64)
);
