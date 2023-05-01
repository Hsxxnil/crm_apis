alter table orders
    add activated_by uuid;

alter table orders
    add activated_at timestamp default now();

create index idx_orders_activated_at
    on orders (activated_at desc);

create index idx_orders_activated_by
    on orders using hash (activated_by);
