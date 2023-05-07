alter table contacts
    add salesperson_id uuid not null;

create index idx_contacts_salesperson_id
    on contacts using hash (salesperson_id);