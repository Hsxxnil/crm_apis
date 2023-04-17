create table crm_industries
(
    industry_id uuid default uuid_generate_v4() not null
        primary key,
    name        text default '':: text not null
);

create index idx_crm_industries_industry_id
    on crm_industries using hash (industry_id);

create index idx_crm_industries_name
    on crm_industries using gin (name gin_trgm_ops);
