create table crm_leads
(
    lead_id      uuid      default uuid_generate_v4() not null
        primary key,
    status       varchar   default '':: varchar not null,
    company_name varchar   default '':: varchar not null,
    source_id    uuid                                 not null,
    industry_id  uuid                                 not null,
    rating       varchar   default '':: varchar not null,
    created_at   timestamp default now()              not null,
    created_by   uuid                                 not null,
    updated_at   timestamp default now()              not null,
    updated_by   uuid                                 not null
);

create index idx_crm_leads_lead_id
    on crm_leads using hash (lead_id);

create index idx_crm_leads_status
    on crm_leads (status);

create index idx_crm_leads_company_name
    on crm_leads (company_name);

create index idx_crm_leads_source_id
    on crm_leads using hash (source_id);

create index idx_crm_leads_industry_id
    on crm_leads using hash (industry_id);

create index idx_crm_leads_rating
    on crm_leads (rating);

create index idx_crm_leads_created_at
    on crm_leads (created_at desc);

create index idx_crm_leads_created_by
    on crm_leads using hash (created_by);

create index idx_crm_leads_updated_at
    on crm_leads (updated_at desc);

create index idx_crm_leads_updated_by
    on crm_leads using hash (updated_by);
