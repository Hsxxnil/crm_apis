alter table quote_products
    add code text not null;

alter table quote_products
    add description text;

create index idx_quote_products_code
    on quote_products using gin (code gin_trgm_ops);

create index idx_quote_products_description
    on quote_products using gin (description gin_trgm_ops);