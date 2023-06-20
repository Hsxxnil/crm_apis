create table historical_records
(
    historical_record_id uuid      default uuid_generate_v4() not null
        primary key,
    source_id            uuid                                 not null,
    content              text      default '':: text not null,
    action               text      default '':: text not null,
    modified_at          timestamp default now()              not null,
    modified_by          uuid                                 not null
);

create index idx_historical_records_historical_record_id
    on historical_records using hash (historical_record_id);

create index idx_historical_records_content
    on historical_records using gin (content gin_trgm_ops);

create index idx_historical_records_action
    on historical_records using gin (action gin_trgm_ops);

create index idx_historical_records_modified_at
    on historical_records (modified_at desc);

create index idx_historical_records_modified_by
    on historical_records using hash (modified_by);
