alter table orders
alter column start_date type date using start_date::date;

