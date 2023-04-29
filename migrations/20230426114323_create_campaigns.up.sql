create table campaigns
(
    campaign_id        uuid      default uuid_generate_v4() not null
        primary key,
    name               text      default '':: text not null,
    status             text      default '':: text not null,
    is_enable          bool                                 not null,
    type               text,
    parent_campaign_id uuid                                 not null,
    start_date         date,
    end_date           date,
    description        text,
    sent               int,
    budget_cost        numeric,
    expected_responses int,
    actual_cost        numeric,
    expected_income    numeric,
    created_at         timestamp default now()              not null,
    created_by         uuid                                 not null,
    updated_at         timestamp default now()              not null,
    updated_by         uuid                                 not null
);

create index idx_campaigns_campaign_id
    on campaigns using hash (campaign_id);

create index idx_campaigns_name
    on campaigns using gin (name gin_trgm_ops);

create index idx_campaigns_status
    on campaigns using gin (status gin_trgm_ops);

create index idx_campaigns_type
    on campaigns using gin (type gin_trgm_ops);

create index idx_campaigns_parent_campaign_id
    on campaigns using hash (parent_campaign_id);

create index idx_campaigns_start_date
    on campaigns (start_date asc);

create index idx_campaigns_end_date
    on campaigns (end_date asc);

create index idx_campaigns_description
    on campaigns using gin (description gin_trgm_ops);

create index idx_campaigns_sent
    on campaigns (sent);

create index idx_campaigns_budget_cost
    on campaigns (budget_cost);

create index idx_campaigns_expected_responses
    on campaigns (expected_responses);

create index idx_campaigns_actual_cost
    on campaigns (actual_cost);

create index idx_campaigns_expected_income
    on campaigns (expected_income);

create index idx_campaigns_created_at
    on campaigns (created_at desc);

create index idx_campaigns_created_by
    on campaigns using hash (created_by);

create index idx_campaigns_updated_at
    on campaigns (updated_at desc);

create index idx_campaigns_updated_by
    on campaigns using hash (updated_by);
