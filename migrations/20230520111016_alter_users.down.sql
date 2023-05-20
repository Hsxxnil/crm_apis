drop index idx_users_role_id;

alter table users
drop column role_id;