alter table quotes
alter column expiration_date type timestamp using expiration_date::timestamp;