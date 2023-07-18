alter table historical_records
drop column content;

alter table historical_records
    add source_type text not null;

alter table historical_records
    add field text not null;

alter table historical_records
    add value text not null;