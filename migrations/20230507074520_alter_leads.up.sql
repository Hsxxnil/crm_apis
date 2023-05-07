alter table leads
    add salesperson_id uuid not null;

create index idx_leads_salesperson_id
    on leads using hash (salesperson_id);