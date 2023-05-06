alter table quotes
    add code text default quote_code() not null;

create index idx_quotes_code
    on quotes using gin (code gin_trgm_ops);