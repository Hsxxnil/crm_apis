alter table opportunities
    add name text default '':: text not null;

create index idx_opportunities_name
    on opportunities using gin (name gin_trgm_ops);