alter table crm_customers
    rename column c_id to customer_id;

alter table crm_customers
    rename column mail to email;

alter table crm_customers
    add updated_by uuid not null;

alter index idx_crm_customers_c_id rename to idx_crm_customers_customer_id;

alter index idx_crm_customers_mail rename to idx_crm_customers_email;

create index idx_crm_customers_updated_by
    on crm_customers using hash (updated_by);
