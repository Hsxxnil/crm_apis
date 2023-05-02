alter table quotes
    add code serial;

create index idx_quotes_code
    on quotes (code);