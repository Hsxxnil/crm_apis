alter table campaigns
    alter column start_date type date using start_date::date;

alter table campaigns
    alter column end_date type date using end_date::date;
