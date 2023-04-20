create table account_types
(
    account_type_id uuid default uuid_generate_v4() not null
        primary key,
    name        text default '':: text not null
);

create index idx_account_types_account_type_id
    on account_types using hash (account_type_id);

create index idx_account_types_name
    on account_types using gin (name gin_trgm_ops);
