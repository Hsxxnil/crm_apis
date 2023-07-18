alter table orders
alter column start_date type timestamp using start_date::timestamp;

