create table crm_customers
(
    c_id          uuid      default uuid_generate_v4() not null
        primary key,
    short_name    text      default ''::text not null,
    eng_name      text,
    name          text      default ''::text not null,
    zip_code      text      default ''::text not null,
    address       text      default ''::text not null,
    tel           text      default ''::text not null,
    fax           text,
    map           text,
    liaison       text,
    mail          text,
    liaison_phone text,
    tax_id_number text      default ''::text not null,
    remark        text,
    created_at    timestamp default now()              not null,
    created_by    uuid,
    updated_at    timestamp default now()              not null
);

create index idx_crm_customers_c_id
    on crm_customers using hash (c_id);

create index idx_crm_customers_short_name
    on crm_customers using gin (short_name gin_trgm_ops);

create index idx_crm_customers_eng_name
    on crm_customers using gin (eng_name gin_trgm_ops);

create index idx_crm_customers_name
    on crm_customers using gin (name gin_trgm_ops);

create index idx_crm_customers_zip_code
    on crm_customers using gin (zip_code gin_trgm_ops);

create index idx_crm_customers_address
    on crm_customers using gin (address gin_trgm_ops);

create index idx_crm_customers_tel
    on crm_customers using gin (tel gin_trgm_ops);

create index idx_crm_customers_fax
    on crm_customers using gin (fax gin_trgm_ops);

create index idx_crm_customers_map
    on crm_customers using gin (map gin_trgm_ops);

create index idx_crm_customers_liaison
    on crm_customers using gin (liaison gin_trgm_ops);

create index idx_crm_customers_mail
    on crm_customers using gin (mail gin_trgm_ops);

create index idx_crm_customers_liaison_phone
    on crm_customers using gin (liaison_phone gin_trgm_ops);

create index idx_crm_customers_tax_id_number
    on crm_customers using gin (tax_id_number gin_trgm_ops);

create index idx_crm_customers_remark
    on crm_customers using gin (remark gin_trgm_ops);

create index idx_crm_customers_created_at
    on crm_customers (created_at desc);

create index idx_crm_customers_created_by
    on crm_customers using hash (created_by);

create index idx_crm_customers_updated_at
    on crm_customers (updated_at desc);
