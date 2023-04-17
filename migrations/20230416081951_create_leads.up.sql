create table leads
(
    lead_id      uuid      default uuid_generate_v4() not null
        primary key,
    status       text      default '':: text not null,
    company_name text      default '':: text not null,
    source_id    uuid                                 not null,
    industry_id  uuid                                 not null,
    rating       text      default '':: text not null,
    created_at   timestamp default now()              not null,
    created_by   uuid                                 not null,
    updated_at   timestamp default now()              not null,
    updated_by   uuid                                 not null
);

create index idx_leads_lead_id
    on leads using hash (lead_id);

create index idx_leads_status
    on leads using gin (status gin_trgm_ops);

create index idx_leads_company_name
    on leads using gin (company_name gin_trgm_ops);

create index idx_leads_source_id
    on leads using hash (source_id);

create index idx_leads_industry_id
    on leads using hash (industry_id);

create index idx_leads_rating
    on leads using gin (rating gin_trgm_ops);

create index idx_leads_created_at
    on leads (created_at desc);

create index idx_leads_created_by
    on leads using hash (created_by);

create index idx_leads_updated_at
    on leads (updated_at desc);

create index idx_leads_updated_by
    on leads using hash (updated_by);
