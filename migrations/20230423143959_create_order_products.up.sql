create table order_products
(
    order_product_id uuid      default uuid_generate_v4() not null
        primary key,
    order_id         uuid                                 not null,
    product_id       uuid                                 not null,
    quantity         int                                  not null,
    unit_price       decimal                              not null,
    sub_total        decimal                              not null,
    description      text,
    created_at       timestamp default now()              not null,
    created_by       uuid                                 not null,
    updated_at       timestamp default now()              not null,
    updated_by       uuid                                 not null
);

create index idx_order_products_order_product_id
    on order_products using hash (order_product_id);

create index idx_order_products_order_id
    on order_products using hash (order_id);

create index idx_order_products_product_id
    on order_products using hash (product_id);

create index idx_order_products_description
    on order_products using gin (description gin_trgm_ops);

create index idx_order_products_created_at
    on order_products (created_at desc);

create index idx_order_products_created_by
    on order_products using hash (created_by);

create index idx_order_products_updated_at
    on order_products (updated_at desc);

create index idx_order_products_updated_by
    on order_products using hash (updated_by);
