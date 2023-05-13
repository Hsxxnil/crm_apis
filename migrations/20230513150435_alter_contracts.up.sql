alter table contracts
    add end_date date not null;

create index idx_contracts_end_date
    on contracts (end_date asc);