create table contracts
(
    contract_id uuid      default uuid_generate_v4() not null
        primary key,
    status      text      default '':: text not null,
    start_date  date                                 not null,
    account_id  uuid                                 not null,
    term        int                                  not null,
    description text,
    created_at  timestamp default now()              not null,
    created_by  uuid                                 not null,
    updated_at  timestamp default now()              not null,
    updated_by  uuid                                 not null
);

create index idx_contracts_contract_id
    on contracts using hash (contract_id);

create index idx_contracts_status
    on contracts using gin (status gin_trgm_ops);

create index idx_contracts_start_date
    on contracts (start_date asc);

create index idx_contracts_account_id
    on contracts using hash (account_id);

create index idx_contracts_term
    on contracts (term);

create index idx_contracts_description
    on contracts using gin (description gin_trgm_ops);

create index idx_contracts_created_at
    on contracts (created_at desc);

create index idx_contracts_created_by
    on contracts using hash (created_by);

create index idx_contracts_updated_at
    on contracts (updated_at desc);

create index idx_contracts_updated_by
    on contracts using hash (updated_by);
