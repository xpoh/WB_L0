\c base10
SET ROLE akaddr;

create schema wb_l0;

alter schema wb_l0 owner to akaddr;

create table wb_l0.wb_orders
(
    order_uid uuid,
    jsonData varchar
);
