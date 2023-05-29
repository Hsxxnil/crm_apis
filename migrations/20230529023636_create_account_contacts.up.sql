create table account_contacts
(
    account_contact_id uuid      default uuid_generate_v4() not null
        primary key,
    account_id          uuid                                 not null,
    contact_id             uuid                                 not null,
    created_at              timestamp default now()              not null,
    created_by              uuid                                 not null,
    updated_at              timestamp default now()              not null,
    updated_by              uuid                                 not null
);

create index idx_account_contacts_account_contact_id
    on account_contacts using hash (account_contact_id);

create index idx_account_contacts_account_id
    on account_contacts using hash (account_id);

create index idx_account_contacts_contact_id
    on account_contacts using hash (contact_id);

create index idx_account_contacts_created_at
    on account_contacts (created_at desc);

create index idx_account_contacts_created_by
    on account_contacts using hash (created_by);

create index idx_account_contacts_updated_at
    on account_contacts (updated_at desc);

create index idx_account_contacts_updated_by
    on account_contacts using hash (updated_by);
