\c base10
SET ROLE akaddr;

create schema wb_l0;

alter schema wb_l0 owner to akaddr;

create table wb_l0.wb_orders
(
    order_uid uuid,
    track_number varchar,
    entry varchar,
    delivery_id uuid,
    payment_id uuid,
    items_id uuid,
    locale varchar,
    internal_signature varchar,
    customer_id uuid,
    delivery_service varchar,
    shardkey varchar,
    sm_id integer,
    date_created date,
    oof_shard varchar
);

create table wb_l0.wb_delivery
(
    delivery_id uuid,
    name varchar,
    phone varchar,
    zip varchar,
    city varchar,
    address varchar,
    region varchar,
    email varchar
);

create table wb_l0.wb_payment
(
    transaction varchar,
    request_id uuid,
    currency varchar,
    provider varchar,
    amount integer,
    payment_dt float,
    bank varchar,
    delivery_cost float,
    goods_total integer,
    custom_fee float
);
create table wb_l0.wb_items
(
    chrt_id integer,
    track_number varchar,
    price float,
    rid varchar,
    name varchar,
    sale float,
    size varchar,
    total_price float,
    nm_id uuid,
    brand varchar,
    status  integer
);