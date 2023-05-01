drop index idx_orders_activated_by;
drop index idx_orders_activated_at;

alter table orders
    drop column activated_at;

alter table orders
    drop column activated_by;