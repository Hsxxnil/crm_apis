drop index idx_accounts_salesperson_id;

alter table accounts
    drop column salesperson_id;