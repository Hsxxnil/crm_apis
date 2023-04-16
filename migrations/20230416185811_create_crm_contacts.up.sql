create table crm_contacts
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
    manager_id   uuid                                 not null,
    account_id   uuid                                 not null,
    created_at   timestamp default now()              not null,
    created_by   uuid                                 not null,
    updated_at   timestamp default now()              not null,
    updated_by   uuid                                 not null
);

create index idx_crm_contacts_contact_id
    on crm_contacts using hash (contact_id);

create index idx_crm_contacts_name
    on crm_contacts using gin (name gin_trgm_ops);

create index idx_crm_contacts_phone_number
    on crm_contacts using gin (phone_number gin_trgm_ops);

create index idx_crm_contacts_cell_phone
    on crm_contacts using gin (cell_phone gin_trgm_ops);

create index idx_crm_contacts_email
    on crm_contacts using gin (email gin_trgm_ops);

create index idx_crm_contacts_manager_id
    on crm_contacts using hash (manager_id);

create index idx_crm_contacts_account_id
    on crm_contacts using hash (account_id);

create index idx_crm_contacts_created_at
    on crm_contacts (created_at desc);

create index idx_crm_contacts_created_by
    on crm_contacts using hash (created_by);

create index idx_crm_contacts_updated_at
    on crm_contacts (updated_at desc);

create index idx_crm_contacts_updated_by
    on crm_contacts using hash (updated_by);
