create table products
(
    product_id  uuid      default uuid_generate_v4() not null
        primary key,
    name        text      default '':: text not null,
    code        text,
    is_enable   bool                                 not null,
    description text,
    created_at  timestamp default now()              not null,
    created_by  uuid                                 not null,
    updated_at  timestamp default now()              not null,
    updated_by  uuid                                 not null
);

create index idx_products_product_id
    on products using hash (product_id);

create index idx_products_name
    on products using gin (name gin_trgm_ops);

create index idx_products_code
    on products using gin (code gin_trgm_ops);

create index idx_products_description
    on products using gin (description gin_trgm_ops);

create index idx_products_created_at
    on products (created_at desc);

create index idx_products_created_by
    on products using hash (created_by);

create index idx_products_updated_at
    on products (updated_at desc);

create index idx_products_updated_by
    on products using hash (updated_by);
