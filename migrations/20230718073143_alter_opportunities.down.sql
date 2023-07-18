alter table opportunities
alter column close_date type date using close_date::date;

