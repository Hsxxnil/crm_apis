alter table accounts
    add salesperson_id uuid not null;

create index idx_accounts_salesperson_id
    on accounts using hash (salesperson_id);