alter table campaigns
    alter column start_date type timestamp using start_date::timestamp;

alter table campaigns
    alter column end_date type timestamp using end_date::timestamp;