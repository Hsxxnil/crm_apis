insert into roles(role_id, company_id, display_name, name, created_by, created_at, updated_by, updated_at)
values ('d56fc184-9441-4396-be6c-d48580650171', '00000000-0000-4000-a000-000000000000', '管理員', 'admin',
        '00000000-0000-4000-a000-000000000000', now(), '00000000-0000-4000-a000-000000000000', now());

insert into users(user_id, company_id, user_name, name, password, created_by, created_at, updated_by, updated_at,
                  role_id)
values ('a1bb0141-68e3-420c-8a92-9332fc21bd25', '00000000-0000-4000-a000-000000000000', 'admin', '管理員',
        '9HXSglPqDWrOyA29croTTu8O8ahmj2EMHhxrsfzrEpJBVykaIkDJ211tJ03aq25Q2iHvkECACPDI/yJXiDsRQDojG1iLqTMQp3nUSmfV/9Yhc3i+ovXLuiRoapCluqw4oxkiuLtqlQMivNTnphmOF+iHnu6sz8N6aouA3mOS89aSoPpHwbWbo4ilh3sPIyEnwLT9npq3ICQwP7FxXPFxaw==',
        '00000000-0000-4000-a000-000000000000', now(), '00000000-0000-4000-a000-000000000000', now(),
        'd56fc184-9441-4396-be6c-d48580650171');

