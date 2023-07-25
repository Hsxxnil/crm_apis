create table event_user_attendees
(
    event_user_attendee_id uuid      default uuid_generate_v4() not null
        primary key,
    event_id               uuid                                 not null,
    attendee_id            uuid                                 not null,
    is_deleted             bool      default false              not null,
    created_at             timestamp default now()              not null,
    created_by             uuid                                 not null,
    updated_at             timestamp default now()              not null,
    updated_by             uuid                                 not null
);

create index idx_event_user_attendees_event_user_attendee_id
    on event_user_attendees using hash (event_user_attendee_id);

create index idx_event_user_attendees_event_id
    on event_user_attendees using hash (event_id);

create index idx_event_user_attendees_attendee_id
    on event_user_attendees using hash (attendee_id);

create index idx_event_user_attendees_created_at
    on event_user_attendees (created_at desc);

create index idx_event_user_attendees_created_by
    on event_user_attendees using hash (created_by);

create index idx_event_user_attendees_updated_at
    on event_user_attendees (updated_at desc);

create index idx_event_user_attendees_updated_by
    on event_user_attendees using hash (updated_by);
