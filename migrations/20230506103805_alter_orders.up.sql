alter table orders
    add code text default order_code() not null;

create index idx_orders_code
    on orders using gin (code gin_trgm_ops);