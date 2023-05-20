create table roles
(
    role_id      uuid      default uuid_generate_v4() not null
        primary key,
    company_id   uuid                                 not null,
    display_name text      default '':: text not null,
    name         text      default '':: text not null,
    is_enable    bool      default true               not null,
    is_deleted   bool      default false              not null,
    created_at   timestamp default now()              not null,
    created_by   uuid                                 not null,
    updated_at   timestamp default now()              not null,
    updated_by   uuid                                 not null
);

create index idx_roles_role_id
    on roles using hash (role_id);

create index idx_roles_display_name
    on roles using gin (display_name gin_trgm_ops);

create index idx_roles_name
    on roles using gin (name gin_trgm_ops);

create index idx_roles_created_at
    on roles (created_at desc);

create index idx_roles_created_by
    on roles using hash (created_by);

create index idx_roles_updated_at
    on roles (updated_at desc);

create index idx_roles_updated_by
    on roles using hash (updated_by);
