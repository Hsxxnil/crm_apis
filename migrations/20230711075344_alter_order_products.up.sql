alter table order_products
    add code text not null;

create index idx_order_products_code
    on order_products using gin (code gin_trgm_ops);