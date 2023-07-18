alter table quotes
alter column expiration_date type date using expiration_date::date;

