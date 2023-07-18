alter table historical_records
    add content text not null;

alter table historical_records
drop column source_type;

alter table historical_records
drop column field;

alter table historical_records
drop column value;