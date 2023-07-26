create table events
(
    event_id    uuid      default uuid_generate_v4() not null
        primary key,
    subject     text      default '':: text not null,
    is_whole    bool                                 not null,
    start_date  timestamp                            not null,
    end_date    timestamp                            not null,
    account_id  uuid,
    type        text,
    location    text,
    description text,
    created_at  timestamp default now()              not null,
    created_by  uuid                                 not null,
    updated_at  timestamp default now()              not null,
    updated_by  uuid                                 not null
);

create index idx_events_event_id
    on events using hash (event_id);

create index idx_events_subject
    on events using gin (subject gin_trgm_ops);

create index idx_events_start_date
    on events (start_date asc);

create index idx_events_end_date
    on events (end_date asc);

create index idx_events_account_id
    on events using hash (account_id);

create index idx_events_type
    on events using gin (type gin_trgm_ops);

create index idx_events_location
    on events using gin (location gin_trgm_ops);

create index idx_events_created_at
    on events (created_at desc);

create index idx_events_created_by
    on events using hash (created_by);

create index idx_events_updated_at
    on events (updated_at desc);

create index idx_events_updated_by
    on events using hash (updated_by);
