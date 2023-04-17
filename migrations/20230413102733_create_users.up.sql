create table users
(
    user_id      uuid      default uuid_generate_v4() not null
        primary key,
    company_id   uuid                                 not null,
    user_name    text      default '':: text not null,
    name         text      default '':: text not null,
    password     text      default '':: text not null,
    is_deleted   bool      default false              not null,
    phone_number text,
    email        text,
    created_at   timestamp default now()              not null,
    created_by   uuid                                 not null,
    updated_at   timestamp default now()              not null,
    updated_by   uuid                                 not null
);

create index idx_users_user_id
    on users using hash (user_id);

create index idx_users_user_name
    on users using gin (user_name gin_trgm_ops);

create unique index uidx_users_user_name
    on users (user_name);

create index idx_users_name
    on users using gin (name gin_trgm_ops);

create index idx_users_phone_number
    on users using gin (phone_number gin_trgm_ops);

create index idx_users_email
    on users using gin (email gin_trgm_ops);

create index idx_users_created_at
    on users (created_at desc);

create index idx_users_created_by
    on users using hash (created_by);

create index idx_users_updated_at
    on users (updated_at desc);

create index idx_users_updated_by
    on users using hash (updated_by);
