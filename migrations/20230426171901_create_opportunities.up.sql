create table opportunities
(
    opportunity_id    uuid      default uuid_generate_v4() not null
        primary key,
    stage             text      default '':: text not null,
    forecast_category text      default '':: text not null,
    close_date        date                                 not null,
    account_id        uuid                                 not null,
    amount            numeric,
    created_at        timestamp default now()              not null,
    created_by        uuid                                 not null,
    updated_at        timestamp default now()              not null,
    updated_by        uuid                                 not null
);

create index idx_opportunities_opportunity_id
    on opportunities using hash (opportunity_id);

create index idx_opportunities_stage
    on opportunities using gin (stage gin_trgm_ops);

create index idx_opportunities_forecast_category
    on opportunities using gin (forecast_category gin_trgm_ops);

create index idx_opportunities_close_date
    on opportunities (close_date asc);

create index idx_opportunities_account_id
    on opportunities using hash (account_id);

create index idx_opportunities_amount
    on opportunities (amount);

create index idx_opportunities_created_at
    on opportunities (created_at desc);

create index idx_opportunities_created_by
    on opportunities using hash (created_by);

create index idx_opportunities_updated_at
    on opportunities (updated_at desc);

create index idx_opportunities_updated_by
    on opportunities using hash (updated_by);
