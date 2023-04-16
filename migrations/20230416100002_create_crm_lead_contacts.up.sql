create table crm_lead_contacts
(
    lead_contact_id uuid      default uuid_generate_v4() not null
        primary key,
    name            text      default '':: text not null,
    title           text,
    phone_number    text      default '':: text not null,
    cell_phone      text,
    email           text,
    line            text,
    created_at      timestamp default now()              not null,
    created_by      uuid                                 not null,
    updated_at      timestamp default now()              not null,
    updated_by      uuid                                 not null
);

create index idx_crm_lead_contacts_lead_contact_id
    on crm_lead_contacts using hash (lead_contact_id);

create index idx_crm_lead_contacts_name
    on crm_lead_contacts using gin (name gin_trgm_ops);

create index idx_crm_lead_contacts_title
    on crm_lead_contacts using gin (title gin_trgm_ops);

create index idx_crm_lead_contacts_phone_number
    on crm_lead_contacts using gin (phone_number gin_trgm_ops);

create index idx_crm_lead_contacts_cell_phone
    on crm_lead_contacts using gin (cell_phone gin_trgm_ops);

create index idx_crm_lead_email
    on crm_lead_contacts using gin (email gin_trgm_ops);

create index idx_crm_lead_line
    on crm_lead_contacts using gin (line gin_trgm_ops);

create index idx_crm_lead_contacts_created_at
    on crm_lead_contacts (created_at desc);

create index idx_crm_lead_contacts_created_by
    on crm_lead_contacts using hash (created_by);

create index idx_crm_lead_contacts_updated_at
    on crm_lead_contacts (updated_at desc);

create index idx_crm_lead_contacts_updated_by
    on crm_lead_contacts using hash (updated_by);
