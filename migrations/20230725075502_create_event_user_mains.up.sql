create table event_user_mains
(
    event_user_main_id uuid      default uuid_generate_v4() not null
        primary key,
    event_id           uuid                                 not null,
    main_id            uuid                                 not null,
    is_deleted         bool      default false              not null,
    created_at         timestamp default now()              not null,
    created_by         uuid                                 not null,
    updated_at         timestamp default now()              not null,
    updated_by         uuid                                 not null
);

create index idx_event_user_mains_event_user_main_id
    on event_user_mains using hash (event_user_main_id);

create index idx_event_user_mains_event_id
    on event_user_mains using hash (event_id);

create index idx_event_user_mains_main_id
    on event_user_mains using hash (main_id);

create index idx_event_user_mains_created_at
    on event_user_mains (created_at desc);

create index idx_event_user_mains_created_by
    on event_user_mains using hash (created_by);

create index idx_event_user_mains_updated_at
    on event_user_mains (updated_at desc);

create index idx_event_user_mains_updated_by
    on event_user_mains using hash (updated_by);
