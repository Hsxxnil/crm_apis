alter table contracts
    add code serial;

create index idx_contracts_code
    on contracts (code);