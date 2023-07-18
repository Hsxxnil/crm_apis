alter table opportunities
alter column close_date type timestamp using close_date::timestamp;