alter table campaigns
    add salesperson_id uuid not null;

create index idx_campaigns_salesperson_id
    on campaigns using hash (salesperson_id);