alter table orders
    add code serial;

create index idx_orders_code
    on orders (code);