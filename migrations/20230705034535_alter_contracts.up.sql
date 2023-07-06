alter table contracts
    add opportunity_id uuid not null;

create index idx_contracts_opportunity_id
    on contracts using hash (opportunity_id);