drop index idx_lead_contacts_lead_id;

alter table lead_contacts
    drop column lead_id;

