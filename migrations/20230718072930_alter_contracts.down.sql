alter table contracts
alter column start_date type date using start_date::date;

alter table contracts
alter column end_date type date using end_date::date;