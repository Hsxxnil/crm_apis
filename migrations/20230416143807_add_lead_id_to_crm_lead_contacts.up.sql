alter table crm_lead_contacts
    add lead_id uuid not null;

create index idx_crm_lead_contacts_lead_id
    on crm_lead_contacts using hash (lead_id);