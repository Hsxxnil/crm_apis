alter table crm_customers
    rename column c_id to customer_id;

alter table crm_customers
    add updated_by uuid not null;

create index idx_crm_customers_updated_by
    on crm_customers using hash (updated_by);
