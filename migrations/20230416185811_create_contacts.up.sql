create table contacts
(
    contact_id   uuid      default uuid_generate_v4() not null
        primary key,
    name         text      default '':: text not null,
    title        text,
    phone_number text      default '':: text not null,
    cell_phone   text,
    email        text,
    salutation   text,
    department   text,
    supervisor_id   uuid                                 not null,
    account_id   uuid                                 not null,
    created_at   timestamp default now()              not null,
    created_by   uuid                                 not null,
    updated_at   timestamp default now()              not null,
    updated_by   uuid                                 not null
);

create index idx_contacts_contact_id
    on contacts using hash (contact_id);

create index idx_contacts_name
    on contacts using gin (name gin_trgm_ops);

create index idx_contacts_phone_number
    on contacts using gin (phone_number gin_trgm_ops);

create index idx_contacts_cell_phone
    on contacts using gin (cell_phone gin_trgm_ops);

create index idx_contacts_email
    on contacts using gin (email gin_trgm_ops);

create index idx_contacts_supervisor_id
    on contacts using hash (supervisor_id);

create index idx_contacts_account_id
    on contacts using hash (account_id);

create index idx_contacts_created_at
    on contacts (created_at desc);

create index idx_contacts_created_by
    on contacts using hash (created_by);

create index idx_contacts_updated_at
    on contacts (updated_at desc);

create index idx_contacts_updated_by
    on contacts using hash (updated_by);
