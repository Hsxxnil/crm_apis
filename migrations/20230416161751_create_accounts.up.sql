create table accounts
(
    account_id        uuid      default uuid_generate_v4() not null
        primary key,
    name              text      default '':: text not null,
    phone_number      text,
    type              text      default '':: text not null,
    industry_id       uuid                                 not null,
    parent_account_id uuid                                 not null,
    created_at        timestamp default now()              not null,
    created_by        uuid                                 not null,
    updated_at        timestamp default now()              not null,
    updated_by        uuid                                 not null
);

create index idx_accounts_account_id
    on accounts using hash (account_id);

create index idx_accounts_name
    on accounts using gin (name gin_trgm_ops);

create index idx_accounts_phone_number
    on accounts using gin (phone_number gin_trgm_ops);

create index idx_accounts_type
    on accounts using gin (type gin_trgm_ops);

create index idx_accounts_industry_id
    on accounts using hash (industry_id);

create index idx_accounts_parent_account_id
    on accounts using hash (parent_account_id);

create index idx_accounts_created_at
    on accounts (created_at desc);

create index idx_accounts_created_by
    on accounts using hash (created_by);

create index idx_accounts_updated_at
    on accounts (updated_at desc);

create index idx_accounts_updated_by
    on accounts using hash (updated_by);
