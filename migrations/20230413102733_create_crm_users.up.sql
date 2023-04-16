create table crm_users
(
    user_id      uuid      default uuid_generate_v4() not null
        primary key,
    company_id   uuid                                 not null,
    user_name    varchar   default '':: varchar not null,
    name         varchar   default '':: varchar not null,
    password     varchar   default '':: varchar not null,
    is_deleted   bool      default false              not null,
    phone_number text,
    email        text,
    created_at   timestamp default now()              not null,
    created_by   uuid                                 not null,
    updated_at   timestamp default now()              not null,
    updated_by   uuid                                 not null
);

create index idx_crm_users_user_id
    on crm_users using hash (user_id);

create index idx_crm_users_user_name
    on crm_users (user_name);

create unique index uidx_crm_users_user_name
    on crm_users (user_name);

create index idx_crm_users_name
    on crm_users (name);

create index idx_crm_users_phone_number
    on crm_users using gin (phone_number gin_trgm_ops);

create index idx_crm_users_email
    on crm_users using gin (email gin_trgm_ops);

create index idx_crm_users_created_at
    on crm_users (created_at desc);

create index idx_crm_users_created_by
    on crm_users using hash (created_by);

create index idx_crm_users_updated_at
    on crm_users (updated_at desc);

create index idx_crm_users_updated_by
    on crm_users using hash (updated_by);
