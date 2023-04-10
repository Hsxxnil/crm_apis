create table crm_accounts
(
    account_id   uuid      default uuid_generate_v4() not null
        primary key,
    company_id   uuid      default '00000000-0000-0000-0000-000000000000'::uuid,
    account      varchar   default '':: varchar not null,
    name         varchar   default '':: varchar not null,
    password     varchar   default '':: varchar not null,
    is_deleted   bool      default false              not null,
    phone_number text,
    email        text,
    created_at   timestamp default now()              not null,
    created_by   uuid                                 not null,
    updated_at   timestamp default now()              not null,
    updated_by   uuid                                 not null
);

create index idx_crm_accounts_account_id
    on crm_accounts using hash (account_id);

create index idx_crm_accounts_account
    on crm_accounts (account);

create index idx_crm_accounts_name
    on crm_accounts (name);

create index idx_crm_accounts_phone_number
    on crm_accounts using gin (phone_number gin_trgm_ops);

create index idx_crm_accounts_email
    on crm_accounts using gin (email gin_trgm_ops);

create index idx_crm_accounts_created_at
    on crm_accounts (created_at desc);

create index idx_crm_accounts_created_by
    on crm_accounts using hash (created_by);

create index idx_crm_accounts_updated_at
    on crm_accounts (updated_at desc);

create index idx_crm_accounts_updated_by
    on crm_accounts using hash (updated_by);
