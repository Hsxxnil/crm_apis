alter table users
    add role_id uuid not null;

create index idx_users_role_id
    on users using hash (role_id);