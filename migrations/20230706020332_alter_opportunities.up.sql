alter table opportunities
    add lead_id uuid;

create index idx_opportunities_lead_id
    on opportunities using hash (lead_id);