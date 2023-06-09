alter table accounts
    alter column type drop default;

drop index idx_accounts_type;

alter table accounts
    alter column type type text[] using type::text[];

alter table accounts
    alter column type set default '{}'::text[];

create index idx_accounts_type
    on accounts using gin (type);