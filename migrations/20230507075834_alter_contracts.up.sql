alter table contracts
    add salesperson_id uuid not null;

create index idx_contracts_salesperson_id
    on contracts using hash (salesperson_id);