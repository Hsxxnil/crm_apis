create table opportunity_campaigns
(
    opportunity_campaign_id uuid      default uuid_generate_v4() not null
        primary key,
    opportunity_id          uuid                                 not null,
    campaign_id             uuid                                 not null,
    created_at              timestamp default now()              not null,
    created_by              uuid                                 not null,
    updated_at              timestamp default now()              not null,
    updated_by              uuid                                 not null
);

create index idx_opportunity_campaigns_opportunity_campaign_id
    on opportunity_campaigns using hash (opportunity_campaign_id);

create index idx_opportunity_campaigns_opportunity_id
    on opportunity_campaigns using hash (opportunity_id);

create index idx_opportunity_campaigns_campaign_id
    on opportunity_campaigns using hash (campaign_id);

create index idx_opportunity_campaigns_created_at
    on opportunity_campaigns (created_at desc);

create index idx_opportunity_campaigns_created_by
    on opportunity_campaigns using hash (created_by);

create index idx_opportunity_campaigns_updated_at
    on opportunity_campaigns (updated_at desc);

create index idx_opportunity_campaigns_updated_by
    on opportunity_campaigns using hash (updated_by);
