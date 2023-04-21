create table orders
(
    order_id    uuid      default uuid_generate_v4() not null
        primary key,
    status      text      default '':: text not null,
    start_date  timestamp                            not null,
    account_id  uuid                                 not null,
    contract_id uuid                                 not null,
    description text,
    created_at  timestamp default now()              not null,
    created_by  uuid                                 not null,
    updated_at  timestamp default now()              not null,
    updated_by  uuid                                 not null
);

create index idx_orders_order_id
    on orders using hash (order_id);

create index idx_orders_status
    on orders using gin (status gin_trgm_ops);

create index idx_orders_start_date
    on orders (start_date asc);

create index idx_orders_account_id
    on orders using hash (account_id);

create index idx_orders_contract_id
    on orders using hash (contract_id);

create index idx_orders_description
    on orders using gin (description gin_trgm_ops);

create index idx_orders_created_at
    on orders (created_at desc);

create index idx_orders_created_by
    on orders using hash (created_by);

create index idx_orders_updated_at
    on orders (updated_at desc);

create index idx_orders_updated_by
    on orders using hash (updated_by);
