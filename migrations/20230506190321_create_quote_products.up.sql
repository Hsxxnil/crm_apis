create table quote_products
(
    quote_product_id uuid      default uuid_generate_v4() not null
        primary key,
    quote_id         uuid                                 not null,
    product_id       uuid                                 not null,
    quantity         int                                  not null,
    unit_price       numeric                              not null,
    sub_total        numeric                              not null,
    discount         numeric                              not null,
    created_at       timestamp default now()              not null,
    created_by       uuid                                 not null,
    updated_at       timestamp default now()              not null,
    updated_by       uuid                                 not null
);

create index idx_quote_products_quote_product_id
    on quote_products using hash (quote_product_id);

create index idx_quote_products_quote_id
    on quote_products using hash (quote_id);

create index idx_quote_products_product_id
    on quote_products using hash (product_id);

create index idx_quote_products_created_at
    on quote_products (created_at desc);

create index idx_quote_products_created_by
    on quote_products using hash (created_by);

create index idx_quote_products_updated_at
    on quote_products (updated_at desc);

create index idx_quote_products_updated_by
    on quote_products using hash (updated_by);
