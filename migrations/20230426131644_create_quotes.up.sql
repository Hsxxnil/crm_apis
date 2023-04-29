create table quotes
(
    quote_id              uuid      default uuid_generate_v4() not null
        primary key,
    name                  text      default '':: text not null,
    status                text      default '':: text not null,
    is_syncing            bool                                 not null,
    opportunity_id        uuid                                 not null,
    account_id            uuid                                 not null,
    expiration_date       date,
    description           text,
    tax                   numeric,
    shipping_and_handling numeric,
    created_at            timestamp default now()              not null,
    created_by            uuid                                 not null,
    updated_at            timestamp default now()              not null,
    updated_by            uuid                                 not null
);

create index idx_quotes_quote_id
    on quotes using hash (quote_id);

create index idx_quotes_name
    on quotes using gin (name gin_trgm_ops);

create index idx_quotes_status
    on quotes using gin (status gin_trgm_ops);

create index idx_quotes_opportunity_id
    on quotes using hash (opportunity_id);

create index idx_quotes_account_id
    on quotes using hash (account_id);

create index idx_quotes_expiration_date
    on quotes (expiration_date asc);

create index idx_quotes_description
    on quotes using gin (description gin_trgm_ops);

create index idx_quotes_tax
    on quotes (tax);

create index idx_quotes_budget_shipping_and_handling
    on quotes (shipping_and_handling);

create index idx_quotes_created_at
    on quotes (created_at desc);

create index idx_quotes_created_by
    on quotes using hash (created_by);

create index idx_quotes_updated_at
    on quotes (updated_at desc);

create index idx_quotes_updated_by
    on quotes using hash (updated_by);
