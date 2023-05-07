alter table opportunities
    add salesperson_id uuid not null;

create index idx_opportunities_salesperson_id
    on opportunities using hash (salesperson_id);