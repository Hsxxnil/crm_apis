create table event_contacts
(
    event_contact_id uuid      default uuid_generate_v4() not null
        primary key,
    event_id         uuid                                 not null,
    contact_id       uuid                                 not null,
    is_deleted       bool      default false              not null,
    created_at       timestamp default now()              not null,
    created_by       uuid                                 not null,
    updated_at       timestamp default now()              not null,
    updated_by       uuid                                 not null
);

create index idx_event_contacts_event_contact_id
    on event_contacts using hash (event_contact_id);

create index idx_event_contacts_event_id
    on event_contacts using hash (event_id);

create index idx_event_contacts_contact_id
    on event_contacts using hash (contact_id);

create index idx_event_contacts_created_at
    on event_contacts (created_at desc);

create index idx_event_contacts_created_by
    on event_contacts using hash (created_by);

create index idx_event_contacts_updated_at
    on event_contacts (updated_at desc);

create index idx_event_contacts_updated_by
    on event_contacts using hash (updated_by);
