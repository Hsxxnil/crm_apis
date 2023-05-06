alter table contracts
    add code text default contract_code() not null;

create index idx_contracts_code
    on contracts using gin (code gin_trgm_ops);