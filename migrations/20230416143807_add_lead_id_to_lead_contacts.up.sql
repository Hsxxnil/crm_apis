alter table lead_contacts
    add lead_id uuid not null;

create index idx_lead_contacts_lead_id
    on lead_contacts using hash (lead_id);